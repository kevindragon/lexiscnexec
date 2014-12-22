// 把存在excel的层级的省、法院转换成php的数组
// go run court_to_php_array.go > court_definition.php
package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"strings"
)

func main() {
	excelFilename := "court/法院名称整理总表20141128.xlsx"
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
		for i := 0; i < 4; i++ {
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

	fmt.Print("<?php\n\nreturn array(\n")
	print(lines, inverseNames)
	fmt.Print(");")

	/*
		for {
			var i int
			fmt.Print("Enter a number: ")
			fmt.Scanf("%d\n", &i)
			if i < 0 || i > len(lines) {
				break
			}

			line := lines[i]
			fmt.Println(i, line, inverseNames[line])
			fmt.Println(sheet1.Rows[i])
		}
	*/
}

func print(order []string, data map[string]string) {
	sortOrder := 1
	for _, key := range order {
		if key != "" {
			nolink := ""
			if parent := data[key]; parent == "" && key != "最高人民法院" {
				nolink = `, "nolink" => true`
			}
			outputLine := fmt.Sprintf(`    "%s" => array("parent" => "%s"%s, "order" => %d), `,
				key, data[key], nolink, sortOrder)
			sortOrder += 1

			fmt.Println(outputLine)
			//fmt.Println(`"`+key+`" =>`, `"`+data[key]+`", `)
		}
	}
}
