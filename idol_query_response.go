package main

import (
	"fmt"
	"github.com/kevindragon/lexiscnexec/idol"
	"os"
)

func main() {
	//filename := "idol/data/idol212-ip_hottopic.xml"
	filename := "idol/data/idol212-topic_taxonomy.xml"

	response, err := idol.ParseXMLFile(filename)
	if err != nil {
		os.Exit(1)
	}

	if len(response.ResponseData.Hits) > 0 {
		fmt.Println("len(response.ResponseData.Hits)", len(response.ResponseData.Hits))
		for _, v := range response.ResponseData.Hits {
			if v.Id > 0 {
				fmt.Println(v)
			}
		}
	}
	//fmt.Println(response, err)
}
