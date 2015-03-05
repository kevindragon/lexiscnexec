package main

import (
	"fmt"

	"github.com/kevindragon/lexiscnexec/baidubaike"
)

func main() {
	laws, _ := baidubaike.GetRelatedLaws("股权转让")
	fmt.Println(laws)

	haha, _ := baidubaike.GetRelatedLaws("年终奖")
	fmt.Println(haha)
}
