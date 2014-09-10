// IDOL10升级测试。测试IDOL10的分词，搜索。
//
// usage: idoltesting -t task options
// task: segment

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kevindragon/lexiscnexec/segment"
)

var task string

func init() {
	flag.StringVar(&task, "t", "", "Specify your task")
}

func main() {
	flag.Parse()
	if task == "" {
		flag.Usage()
		exit("Specify your task")
	}

	switch task {
	case "segment":
		segment.Run()
	default:
		exit("task name incorrect")
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
