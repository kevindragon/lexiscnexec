// 把存在excel的层级的省、法院转换成php的数组
package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"strings"
)

func main() {
	excelFilename := "court/courts.xlsx"
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

	inverseNames := make(map[string]string)
	level := make([]string, 4)
	current := make([]string, 4)
	for _, row := range sheet1.Rows[1:] {
		current = []string{"", "", "", ""}
		for i := 0; i < 4; i++ {
			if len(row.Cells) < 1 {
				break
			}
			cell := strings.Trim(row.Cells[i].String(), " ")
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
				lines = append(lines, content)
				break
			}
			if content != "" {
				inverseNames[content] = level[i-1]
				lines = append(lines, content)
			}
		}
	}

	print(lines, inverseNames)
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
	for _, key := range order {
		if key != "" {
			fmt.Println(`"`+key+`" =>`, `"`+data[key]+`", `)
		}
	}
}
