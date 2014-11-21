// 把存在于数据库的不统一的法院名称统一
// 输出csv格式的输出
package main

import (
	"fmt"
	"github.com/kevindragon/lexiscnexec/court"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	convert()
	//test()
}

func convert() {
	analyzer := court.NewAnalyzer()
	analyzer.LoadDict("court/dict.txt")
	analyzer.LoadStandard("court/standard.txt")
	analyzer.LoadMapping("court/manual_mapping.txt")

	lines, err := readFile("court/courts_in_db.txt")
	if err != nil {
		fmt.Println("read file court/courts_in_db.txt failed")
		os.Exit(1)
	}

	var standards []string
	var unStandards []string
	for index, l := range lines {
		line := l
		line = strings.Replace(strings.Trim(line, " "), " ", "", -1)
		if line == "" {
			continue
		}

		standardName := analyzer.GetFromMapping(line)
		if standardName == "" {
			if analyzer.IsStandard(line) {
				standardName = line
			} else {
				standardName = analyzer.GetAncestor(line)
			}
			sameTop := underSameTop(analyzer, standardName, line)
			if !sameTop {
				unStandards = append(unStandards, fmt.Sprintf(`"%d","%s",`, index+1, l))
				standardName = ""
			} else {
				standards = append(standards, fmt.Sprintf(`"%d","%s","%s"`, index+1, l, standardName))
			}
		} else {
			standards = append(standards, fmt.Sprintf(`"%d","%s","%s"`, index+1, l, standardName))
		}

	}

	fmt.Printf("%s,%s,%s\n", "#", "原始名称", "转换后的名称")
	for _, line := range standards {
		fmt.Printf("%s\n", line)
	}
	fmt.Println(`"","",""`)
	for _, line := range unStandards {
		fmt.Printf("%s\n", line)
	}
}

func test() {
	analyzer := court.NewAnalyzer()

	analyzer.LoadDict("court/dict.txt")
	names := []string{
		"鼎城区人民法院",
	}
	for _, name := range names {
		fmt.Println(name, "terms:", analyzer.ToTerms(name))
	}
	analyzer.LoadStandard("court/standard.txt")

	fmt.Println("")

	for _, name := range names {
		standardName := analyzer.GetAncestor(name)
		sameTop := underSameTop(analyzer, standardName, name)
		if sameTop {
			fmt.Println(name, "-->", standardName, sameTop)
		} else {
			fmt.Println(name, "<-?->", standardName, sameTop)

		}
	}
}

func underSameTop(analyzer *court.Analyzer, src, dist string) bool {
	src, dist = strings.Trim(src, " "), strings.Trim(dist, " ")
	if src == "" || dist == "" {
		return false
	}

	srcTerms := analyzer.ToTerms(src)
	distTerms := analyzer.ToTerms(dist)

	if srcTerms[0] == distTerms[0] {
		return true
	}

	for _, distTerm := range distTerms {
		if !analyzer.IsTop(distTerm) {
			distTop := analyzer.GetTop(distTerm)
			if distTop != "" {
				distTerms = append(distTerms, distTop)
			}
		}
	}

	srcTop := srcTerms[0]
	for _, srcTerm := range srcTerms {
		if !analyzer.IsTop(srcTerm) {
			tmpTop := analyzer.GetTop(srcTerm)
			if tmpTop != "" {
				srcTop = tmpTop
				break
			}
		}
	}

	for _, distTerm := range distTerms {
		if srcTop == distTerm {
			return true
		}
	}

	return false
}

func readFile(filepath string) ([]string, error) {
	if _, err := os.Stat(filepath); err != nil {
		return []string{}, err
	}

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return []string{}, err
	}

	content := strings.Replace(string(bytes), "\r", "\n", -1)
	content = strings.Replace(content, "\n\n", "\n", -1)
	lines := strings.Split(content, "\n")

	return lines, nil
}
