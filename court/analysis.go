// 简单的使用trie对案例的发文机关做分词
package court

import (
	"fmt"
	"github.com/kevindragon/trie"
	"io/ioutil"
	"os"
	"strings"
)

type Analysis struct {
	tree  *trie.Trie
	chain map[string][]string
}

func NewAnalysis() *Analysis {
	analysis := &Analysis{
		tree:  trie.Create(),
		chain: make(map[string][]string),
	}
	return analysis
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

func (analysis *Analysis) LoadDict(file string) bool {
	lines, err := readFile(file)
	if err != nil {
		return false
	}

	for _, line := range lines {
		analysis.tree.Add(line)
	}

	return true
}

func (analysis *Analysis) LoadStandard(file string) bool {
	lines, err := readFile(file)
	if err != nil {
		return false
	}

	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}
		analysis.AddInverseChain(line)
	}

	return true
}

func (analysis *Analysis) ToTerms(str string) []string {
	var terms []string

	chars := []rune(str)

	if len(chars) == 0 {
		return terms
	}
	/*
		// 正向最小匹配
		var match string = ""
		for i := 0; i < len(chars); i++ {
			match += string(chars[i])
			if analysis.tree.Find(match) {
				terms = append(terms, match)
				match = ""
			}
		}
		if match != "" {
			terms = append(terms, match)
		}
	*/
	// 逆向最大匹配
	var inverseMatch string
	var lastTerm string
	for i := len(chars) - 1; i >= 0; i-- {
		inverseMatch = string(chars[i]) + inverseMatch

		// 把当前词加上最后一个分词看是否是一个词
		if len(terms) > 0 {
			lastTerm = terms[len(terms)-1]
			newMatch := inverseMatch + lastTerm
			if analysis.tree.Find(newMatch) {
				terms[len(terms)-1] = newMatch
				inverseMatch = ""
				continue
			}
		}

		if analysis.tree.Find(inverseMatch) {
			terms = append([]string{inverseMatch}, terms...)
			inverseMatch = ""
		}
	}
	if inverseMatch != "" {
		terms = append([]string{inverseMatch}, terms...)
	}

	return terms
}

var suffix map[string]bool = map[string]bool{
	"高级人民法院": true,
	"中级人民法院": true,
	"人民法院":   true,
	"法院":     true,
	"分院":     true,
}

func removeSuffix(analysis *Analysis, str string) string {
	if len(str) == 0 {
		return str
	}

	terms := analysis.ToTerms(str)

	lastTerm := terms[len(terms)-1]
	if _, ok := suffix[lastTerm]; ok {
		terms = terms[:len(terms)-1]
	}

	return strings.Join(terms, "")
}

func (analysis *Analysis) AddInverseChain(str string) {
	terms := analysis.ToTerms(removeSuffix(analysis, str))

	for i := len(terms) - 1; i >= 0; i-- {
		var parentTerm string
		if i == 0 {
			parentTerm = ""
		} else {
			parentTerm = terms[i-1]
		}

		if parent, ok := analysis.chain[terms[i]]; ok {
			exists := false
			for index := range parent {
				if parent[index] == parentTerm {
					exists = true
					break
				}
			}
			if !exists {
				analysis.chain[terms[i]] = append(analysis.chain[terms[i]],
					parentTerm)
			}
		} else {
			analysis.chain[terms[i]] = append(analysis.chain[terms[i]],
				parentTerm)
		}
	}
}

var districtSuffix []string = []string{
	"县",
	"区",
	"市",
}

func (analysis *Analysis) GetAncestor(str string) string {
	shortStr := removeSuffix(analysis, str)
	last := strings.Replace(str, shortStr, "", -1)

	// 先分词
	terms := analysis.ToTerms(shortStr)
	fmt.Println("terms", terms)

	var ancestor string
	for i := len(terms) - 1; i >= 0; i-- {
		term := terms[i]

		fmt.Println("find", terms[i], "parent")

		// 第一次如果没有找到对应的父级，是不是最后一个字不是县区市呢
		for _, distSuffix := range districtSuffix {
			if strings.HasSuffix(term, distSuffix) {
				continue
			}
			ancestor = getAncestor(analysis, term+distSuffix, "")

			if ancestor != "" {
				return ancestor + last
			}
		}

		ancestor = getAncestor(analysis, term, "")

		if ancestor != "" {
			return ancestor + last
		}

		if ancestor == "" {
			return ""
		}
	}

	return ancestor + last
}

// 查找str的父级，如果有suffix，一起拼接上返回
// 如果没有找到返回空字符串
// 从analysis.chain中递归查找
func getAncestor(analysis *Analysis, str, suffix string) string {
	if parents, ok := analysis.chain[str]; ok {
		//fmt.Println(str, "parents", parents, "with suffix", suffix)
		if len(parents) > 1 {
			return ""
		}
		for _, parent := range parents {
			if parent == "" {
				return str + suffix
			}
			return getAncestor(analysis, parent, str+suffix)
		}
	}

	return ""
}

func (analysis *Analysis) Print() {
	fmt.Println(analysis.chain)
}
