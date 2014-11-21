// 从court/mapping.txt文件中取出对应关系，生成php脚本以做案例的法院迁移
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("court/mapping.txt")
	if err != nil {
		os.Exit(1)
	}
	content := strings.Replace(string(bytes), "\r", "", -1)
	lines := strings.Split(content, "\n")

	phpArray := ""
	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Split(line, ",")
		if len(fields) < 2 || fields[0] == fields[1] {
			continue
		}

		phpArray += fmt.Sprintf("    \"%s\" => \"%s\", \n", fields[0], fields[1])
	}

	phpCode := getPHPCode()

	fmt.Print(strings.Replace(phpCode, "{{ARRAY}}", phpArray, -1))
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
