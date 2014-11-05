package main

import (
	"fmt"
	"github.com/kevindragon/lexiscnexec/court"
)

func main() {
	analysis := court.NewAnalysis()

	analysis.LoadDict("court/dict.txt")

	//fmt.Println("玄武人民法院 terms:", analysis.ToTerms("玄武人民法院"))
	//fmt.Println("中华人民共和国淄博市中级人民法院 terms:", analysis.ToTerms("中华人民共和国淄博市中级人民法院"))

	analysis.LoadStandard("court/standard.txt")

	names := []string{
		//"江苏省南京市玄武区人民法院",
		"玄武人民法院",
		"苏州市姑苏区人民法院",
		//"中华人民共和国江苏省高级人民法院",
		//"中华人民共和国淄博市中级人民法院",
	}
	for _, name := range names {
		standardName := analysis.GetAncestor(name)
		fmt.Println(name, "-->", standardName, "\n")
	}
}
