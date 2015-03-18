// 把百度百科词条内容里面提取到的书名号实体到IDOL里面搜索
// 查找对应的法规，只在标题精确搜索
package baidubaike

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/kevindragon/lexiscnexec/config"
	"github.com/kevindragon/lexiscnexec/idol"
	"github.com/tealeg/xlsx"
)

const (
	INPUT_FILENAME  = "baidubaike/data/baike.csv"
	OUTPUT_FILENAME = "baidubaike/data/baike_analyzed.xlsx"
)

var re = regexp.MustCompile("(（[^）]+）)?（[^）]+）$")

func AnalyzeEntity() {
	file := xlsx.NewFile()
	sheet := file.AddSheet("Sheet1")
	sheet.SetColWidth(4, 4, 60)
	sheet.SetColWidth(5, 5, 15)

	highlightStyle := xlsx.NewStyle()
	highlightStyle.Fill = xlsx.Fill{"", "", "FFFF0000"}
	highlightStyle.ApplyFill = true
	highlightStyle.ApplyFont = true
	highlightStyle.Font.Size = 11
	highlightStyle.Font.Name = "Calibri"
	highlightStyle.Font.Family = 2
	highlightStyle.Font.Bold = true
	highlightStyle.Font.Color = "FFFF0000"

	normalStyle := xlsx.NewStyle()
	normalStyle.Font.Size = 11
	normalStyle.Font.Name = "Calibri"
	normalStyle.Font.Family = 2

	csvWalk(func(num int, record []string, err error) {
		for i, _ := range record {
			removeBookTitleMark(&record[i])
		}

		// fmt.Println(record)

		sheetRowSlice := make([]string, 0)

		row := sheet.AddRow()

		// 判断有没有书名号的实体
		if num > 0 && len(record) >= 3 {
			results := searchInIDOL(record[2])
			fmt.Println("results", results)

			sheetRowSlice = append(sheetRowSlice, record[:4]...)

			originTitle := record[1]

			titles := make([]string, 0)
			pureTitles := make([]string, 0)
			ids := make([]string, 0)
			for index, result := range results {
				titles = append(titles,
					fmt.Sprintf("%d，%s", index+1, result.title))
				pureTitles = append(pureTitles, result.pureTitle)
				ids = append(ids, fmt.Sprintf("%d，%d", index+1, result.id))
			}
			sheetRowSlice = append(sheetRowSlice, strings.Join(titles, "\n"))
			sheetRowSlice = append(sheetRowSlice, strings.Join(ids, "\n"))

			for index, cellContent := range sheetRowSlice {
				cell := row.AddCell()
				if index == 2 && len(pureTitles) > 0 && originTitle == pureTitles[0] {
					cell.SetStyle(highlightStyle)
				} else {
					cell.SetStyle(normalStyle)
				}
				cell.Value = cellContent
			}
		} else {
			row.WriteSlice(&record, -1)
		}
	})

	err := file.Save(OUTPUT_FILENAME)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

// 读取csv每一行，调用回调函数进行每一行的处理
func csvWalk(callback func(int, []string, error)) {
	file, err := os.Open(INPUT_FILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	indexBreak := 0

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		callback(indexBreak, record, err)

		indexBreak += 1
		if indexBreak > 3 {
			break
		}
	}
}

func removeBookTitleMark(entity *string) string {
	entityCopy := *entity
	if strings.HasPrefix(entityCopy, "《") {
		entityCopyRune := []rune(entityCopy)
		entityCopy = string(entityCopyRune[1:])
	}
	if strings.HasSuffix(entityCopy, "》") {
		entityCopyRune := []rune(entityCopy)
		entityCopy = string(entityCopyRune[:len(entityCopyRune)-1])
	}

	*entity = entityCopy

	return entityCopy
}

type rows struct {
	title     string
	pureTitle string
	id        int
}

func searchInIDOL(entity string) []rows {
	keyword := url.QueryEscape(entity)
	queryString := fmt.Sprintf(`action=Query&AnyLanguage=true&Combine=simple&DatabaseMatch=law+lawpic&FieldRestriction=DRETITLE&MaxResults=10&PrintFields=ID+DRETITLE+ISSUE_DATE+POWER_LEVEL+EFFECT_ID+EFFECT_STATUS&Sort=Relevance+power_level:numberincreasing+Date&Start=1&Text=("%s")+OR+("%s":TAGS)&fieldtext=EQUAL{1}:EFFECT_ID`, keyword, keyword)
	uri := config.AUTN_HOST + "/" + queryString
	fmt.Println(uri)

	bytes, _ := idol.Query(uri)

	var autnResponse idol.Response
	xml.Unmarshal(bytes, &autnResponse)

	results := make([]rows, 0)

	if autnResponse.ResponseData.Numhits > 0 {
		for _, row := range autnResponse.ResponseData.Hits {
			pureTitle := re.ReplaceAllString(row.Title, "")
			results = append(results, rows{
				row.Title,
				pureTitle,
				row.Id,
			})
		}
	}

	return results
}
