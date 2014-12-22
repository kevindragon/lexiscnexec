// 从court/mapping.txt文件中取出对应关系，生成php脚本以做案例的法院迁移
package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"strings"
)

type Table struct {
	Old, New string
}

func readTable() []Table {
	xlFile, err := xlsx.OpenFile("court/Copy of 法院名称第二次转换 (3).xlsx")
	if err != nil {
		panic(err)
	}

	if len(xlFile.Sheets) < 1 {
		fmt.Println("There are sheets here.")
		os.Exit(1)
	}

	sheet1 := xlFile.Sheets[0]
	tables := make([]Table, 0)
	for _, row := range sheet1.Rows[1:] {
		if len(row.Cells) < 3 {
			continue
		}
		oldName := strings.Replace(row.Cells[1].String(), " ", "", -1)
		newName := strings.Replace(row.Cells[2].String(), " ", "", -1)
		if oldName != "" && newName != "" {
			table := Table{
				Old: oldName,
				New: newName,
			}
			tables = append(tables, table)
		}
	}

	return tables
}

func readStandard() map[string]bool {
	bytes, err := ioutil.ReadFile("court/standard.txt")
	if err != nil {
		panic(err)
	}
	c := string(bytes)
	c = strings.Replace(c, "\r", "\n", -1)
	c = strings.Replace(c, "\n\n", "\n", -1)
	lines := strings.Split(c, "\n")

	standard := make(map[string]bool)
	for _, line := range lines {
		standard[strings.Trim(line, " ")] = true
	}

	return standard
}

func print(order []string, data map[string]string) string {
	phpArray := ""
	sortOrder := 1
	for _, key := range order {
		if key != "" {
			outputLine := fmt.Sprintf(`    "%s" => array("parent" => "%s", "order" => %d), `,
				key, data[key], sortOrder)
			sortOrder += 1

			phpArray += outputLine + "\n"
		}
	}

	return phpArray
}

func getPHPCode() string {
	return `
<?php

$argv = $_SERVER['argv'] ;
if ($argv[1] != "LexisNexis") {
    exit("auth failed\n");
}

include '../main.inc.php';
define('COMM_PATH', R_P . '/topic/');
require_once R_P . 'topic/libs/db.class.php';

$db = db::getInstance();
$stgdb = db::getInstance('stg');

$mapping = array(
{{ARRAY}}
);

foreach ($mapping as $old => $new) {
    if ($old == $new) {
        continue;
    }

    $sql = sprintf("update cases set issue_party = '%s' where issue_party = '%s';", $new, $old);
    echo $sql . "\n";

    $db->update($sql);
    $stgdb->update($sql);
}
`
}

func main() {
	standard := readStandard()
	tables := readTable()
	phpArray := ""
	for _, table := range tables {
		if _, ok := standard[table.New]; !ok {
			fmt.Println(table.Old, table.New)
			continue
		}
		if table.Old != table.New {
			phpArray += fmt.Sprintf("    \"%s\" => \"%s\", \n",
				table.Old, table.New)
		}
	}

	phpCode := getPHPCode()

	fmt.Print(strings.Replace(phpCode, "{{ARRAY}}", phpArray, -1))
}
