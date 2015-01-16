// 生成发布的脚本
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filelistFile := "release/files.txt"

	file, err := os.Open(filelistFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	cmds := make([]string, 0)

	for _, line := range lines {
		fileExt := filepath.Ext(line)

		userGroup := "www:www"
		if fileExt == ".php" {
			userGroup = "root:root"
		}
		cmd := fmt.Sprintf(`chown %s "%s"`, userGroup, line)
		cmds = append(cmds, cmd)
	}

	fmt.Println("#!/bin/sh\n\ncd /home/www/alpha/\n\ntar xzvf update.tar.gz\n\n")
	fmt.Println(strings.Join(cmds, "\n"))
	fmt.Printf(`ll "%s"`, strings.Join(lines, `" "`))
	fmt.Println("\n\nrm update.tar.gz\n")
}
