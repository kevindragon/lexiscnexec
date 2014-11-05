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
