package main

import (
	"fmt"
	"github.com/kevindragon/lexiscnexec/court"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	analysis := court.NewAnalysis()
	analysis.LoadDict("court/dict.txt")
	analysis.LoadStandard("court/standard.txt")

	lines, err := readFile("court/courts_in_db.txt")
	if err != nil {
		fmt.Println("read file court/courts_in_db.txt failed")
	}

	for _, line := range lines {
		line = strings.Trim(line, " ")
		if line == "" {
			continue
		}
		standardName := analysis.GetAncestor(line)
		fmt.Println(line, "-->", standardName)
	}

	//test()
}

func test() {
	analysis := court.NewAnalysis()

	analysis.LoadDict("court/dict.txt")

	//fmt.Println("江苏省南京中级人民法院 terms:", analysis.ToTerms("江苏省南京中级人民法院"))
	//fmt.Println("南京市江宁县人民法院 terms:", analysis.ToTerms("南京市江宁县人民法院"))
	//fmt.Println("玄武人民法院 terms:", analysis.ToTerms("玄武人民法院"))
	//fmt.Println("中华人民共和国淄博市中级人民法院 terms:", analysis.ToTerms("中华人民共和国淄博市中级人民法院"))

	analysis.LoadStandard("court/standard.txt")

	names := []string{
		"淄博市淄川区人民法院",
		/*
			"重庆市第二中级人民法院",
			"重庆市铜梁县人民法院",
			"江苏省南京市玄武区人民法院",
			"玄武人民法院",
			"苏州市姑苏区人民法院",
			"南京市江宁县人民法院",
			"中华人民共和国江苏省高级人民法院",
			"中华人民共和国淄博市中级人民法院",
			"中华人民共和国江苏省高级人民法院",
			"江苏南京市中级人民法院",
			"江苏省南京中级人民法院",
			"中华人民共和国南京市人民法院",
			"中华人民共和国南京市中级人民法院",
			"南京市玄武区人民法院",
			"南京市玄武区某某法院",
			"玄武区人民法院",
			"南京市白下区人民法院",
			"南京市秦淮区人民法院",
			"秦淮区人民法院 ",
			"中华人民共和国南京市秦淮区人民法院",
			"南京市下关区人民法院",
			"南京市鼓楼区人民法院",
			"南京市鼓楼区法院",
			"南京市下关区人民法院",
			"南京浦口区人民法院",
			"南京市浦口区人民法院",
			"南京市栖霞区人民法院",
			"某市栖霞区人民法院",
			"南京市雨花台区人民法院 ",
			"雨花台区人民法院",
			"江苏省南京市江宁区（县）人民法院",
			"南京市江宁县人民法院",
			"江宁区人民法院",
			"南京市江宁区人民法院 ",
			"南京市江宁县人民法院",
			"江宁区人民法院",
		*/
	}
	for _, name := range names {
		standardName := analysis.GetAncestor(name)
		fmt.Println(name, "-->", standardName)
	}

}

func readFile(filepath string) ([]string, error) {
	if _, err := os.Stat(filepath); err != nil {
		return []string{}, err
	}

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return []string{}, err
	}

	content := strings.Replace(string(bytes), "\r", "\n", -1)
	content = strings.Replace(content, "\n\n", "\n", -1)
	lines := strings.Split(content, "\n")

	return lines, nil
}
