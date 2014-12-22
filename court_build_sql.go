// 把存在excel的层级的省、法院转换成php的数组
// go run court_to_php_array.go > court_definition.php
package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	genCnSql()
	//cnenmapsql()
}

func genCnSql() {
	excelFilename := "court/法院名称整理总表20141128 - Update.xlsx"
	xlFile, err := xlsx.OpenFile(excelFilename)
	if err != nil {
		panic(err)
	}

	if len(xlFile.Sheets) < 1 {
		fmt.Println("There are sheets here.")
		os.Exit(1)
	}

	sheet1 := xlFile.Sheets[0]

	lines := make([]string, 0)
	lineUniqMap := make(map[string]string)

	inverseNames := make(map[string]string)
	level := make([]string, 4)
	current := make([]string, 4)

	for _, row := range sheet1.Rows[1:] {
		current = []string{"", "", "", ""}
		// 读取中文名称
		for i := 0; i < 3; i++ {
			if len(row.Cells) < 1 {
				break
			}

			cell := strings.Replace(row.Cells[i].String(), " ", "", -1)
			if cell != "" {
				level[i] = cell
				current[i] = cell
				break
			}
		}

		if strings.Join(level, "") == "" || strings.Join(current, "") == "" {
			continue
		}

		for i, content := range current {
			if content != "" && i == 0 {
				inverseNames[content] = ""
				if _, ok := lineUniqMap[content]; !ok {
					lines = append(lines, content)
					lineUniqMap[content] = "1"
				}
				break
			}
			if content != "" {
				inverseNames[content] = level[i-1]
				if _, ok := lineUniqMap[content]; !ok {
					lines = append(lines, content)
					lineUniqMap[content] = "1"
				}
				break
			}
		}
	}

	print(lines, inverseNames)
}

func print(order []string, data map[string]string) {
	sortOrder := 1
	id := 1
	nameIdMapping := make(map[string]int)
	for _, key := range order {
		if key != "" {
			isDistrict := 0
			if parent := data[key]; parent == "" && key != "最高人民法院" {
				isDistrict = 1
			}

			nameIdMapping[key] = id

			var sqlStr string
			if data[key] == "" {
				sqlStr = fmt.Sprintf("insert into case_court_name(`id`, `name_cn`, `is_district`, `parent`, `order`) "+
					"values(%d, '%s', '%d', %d, %d);",
					id, key, isDistrict, 0, sortOrder)
			} else {
				sqlStr = fmt.Sprintf("insert into case_court_name(`id`, `name_cn`, `is_district`, `parent`, `order`) "+
					"values(%d, '%s', '%d', %d, %d);",
					id, key, isDistrict, nameIdMapping[data[key]], sortOrder)
			}

			sortOrder += 1
			id += 1

			fmt.Println(sqlStr)
		}
	}
}

func cnenmapsql() {
	cnenmapFile := "court/cnenmap.txt"
	bytes, err := ioutil.ReadFile(cnenmapFile)
	if err != nil {
		panic(err)
	}

	content := strings.Replace(string(bytes), "\r", "\n", -1)
	content = strings.Replace(content, "\n\n", "\n", -1)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		fields := strings.Split(line, "@")
		if len(fields) < 2 {
			continue
		}
		fields[0] = strings.Replace(fields[0], " ", "", -1)
		sqlStr := fmt.Sprintf("update case_court_name set name_en = '%s' where name_cn = '%s' limit 1;",
			strings.Replace(fields[1], "'", "\\'", -1),
			strings.Replace(fields[0], "'", "\\'", -1))
		fmt.Println(sqlStr)
	}
}
