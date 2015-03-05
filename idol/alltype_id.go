// http://10.123.4.210:9002/action=query&fieldtext=NOT%20EXISTS%7B%7D:alltype_id&maxresults=50000&printfields=id
// 用上面的命令查询出没有alltype_id节点的数据
// 用下面的程序找到这些数据对应的库以供fetch
package idol

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func AlltypeId() {
	bytes, err := ioutil.ReadFile("idol/data/alltype_id.xml")
	if err != nil {
		panic(err)
	}

	var response Response
	xml.Unmarshal(bytes, &response)

	skeepDB := map[string]bool{
		"ctax_topic":        true,
		"mini_book_chapter": true,
		"mini_book":         true,
		"topic_taxonomy":    true,
		"mini_bbs":          true,
		"book":              true,
	}

	all := make(map[string][]int)

	for _, hit := range response.ResponseData.Hits {
		if _, ok := skeepDB[hit.AutnDatabase]; ok {
			continue
		}

		if _, ok := all[hit.AutnDatabase]; !ok {
			all[hit.AutnDatabase] = make([]int, 0)
		}
		all[hit.AutnDatabase] = append(all[hit.AutnDatabase], hit.Id)
	}

	for k, v := range all {
		fmt.Println(k)
		for _, id := range v {
			fmt.Printf("%d,", id)
		}
		fmt.Println()
	}
}
