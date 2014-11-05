package court_test

import (
	//"fmt"
	"github.com/kevindragon/lexiscnexec/court"
	"testing"
)

func TestLoadDict(t *testing.T) {
	analysis := court.NewAnalysis()

	load := analysis.LoadDict("dict.txt")
	if !load {
		t.Error("load dict.txt error")
	}
}

func TestLoadStandard(t *testing.T) {
	analysis := court.NewAnalysis()

	analysis.LoadDict("dict.txt")
	load := analysis.LoadStandard("standard.txt")
	if !load {
		t.Error("load standard.txt error")
	}
}

func TestInverseChain(t *testing.T) {
	analysis := court.NewAnalysis()

	analysis.LoadDict("dict.txt")
	analysis.LoadStandard("standard.txt")

	testData := map[string]string{
		"中华人民共和国江苏省高级人民法院": "江苏省高级人民法院",
		"江苏省高级人民法院":        "江苏省高级人民法院",
		"徐州市中级人民法院":        "江苏省徐州市中级人民法院",
		"玄武区人民法院":          "江苏省南京市玄武区人民法院",
		"中华人民共和国淄博市中级人民法院": "山东省淄博市中级人民法院",
	}
	for mussy, standard := range testData {
		standardName := analysis.GetAncestor(mussy)
		if standardName != standard {
			t.Error(mussy, "can't get standard name", standard,
				"by returned", standardName)
		}
	}
}

func TestToTerms(t *testing.T) {
	analysis := court.NewAnalysis()

	analysis.LoadDict("dict.txt")

	testData := map[string][]string{
		"江苏省南京中级人民法院":      []string{"江苏省", "南京", "中级人民法院"},
		"南京市江宁县人民法院":       []string{"南京市", "江宁县", "人民法院"},
		"玄武人民法院":           []string{"玄武", "人民法院"},
		"中华人民共和国淄博市中级人民法院": []string{"中华人民共和国", "淄博市", "中级人民法院"},
		"重庆市铜梁县人民法院":       []string{"重庆市", "铜梁县", "人民法院"},
	}
	for word, terms := range testData {
		actualTerms := analysis.ToTerms(word)
		//fmt.Println(word, terms, actualTerms)
		if len(actualTerms) != len(terms) {
			t.Error("terms analysis failed", word, terms, actualTerms)
		}
		for index := range actualTerms {
			if actualTerms[index] != terms[index] {
				t.Error("terms analysis failed", word, terms, actualTerms)
			}
		}
	}
}
