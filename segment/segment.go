// 对比两个IDOL的分词结果是否一致
package segment

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Run() {
	keywords := readKeywords()

	if len(keywords) <= 0 {
		fmt.Println("No keywords.")
		os.Exit(1)
	}

	f, err := os.Create("segment/terms.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	w.Write([]string{"keyword", "original terms", "new terms", "same"})

	for _, keyword := range keywords {
		oriTerms, newTerms := getTerms(keyword)

		oriTermStr := strings.Join(oriTerms.Terms, " / ")
		newTermStr := strings.Join(newTerms.Terms, " / ")
		flagColumn := "Y"
		if oriTermStr != newTermStr {
			flagColumn = "N"
		}
		w.Write([]string{keyword, oriTermStr, newTermStr, flagColumn})
	}
	w.Flush()
}

type TermsRoot struct {
	Response string   `xml:"response"`
	Number   int      `xml:"responsedata>number_of_terms"`
	Terms    []string `xml:"responsedata>term"`
}

func getTerms(keyword string) (TermsRoot, TermsRoot) {
	escapedKeyword := url.QueryEscape(keyword)
	cmdTmpl := "a=termgetinfo&stemming=false&type=NONE&onlyexisting=false&text=%s"
	oriUri := fmt.Sprintf("http://10.123.4.210:9002/"+cmdTmpl, escapedKeyword)
	newUri := fmt.Sprintf("http://10.123.4.215:9002/"+cmdTmpl, escapedKeyword)
	oriRespXml := query(oriUri)
	newRespXml := query(newUri)

	var oriResp, newResp TermsRoot
	xml.Unmarshal([]byte(oriRespXml), &oriResp)
	xml.Unmarshal([]byte(newRespXml), &newResp)
	return oriResp, newResp
}
func query(uri string) string {
	resp, err := http.Get(uri)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(bb)
}

func readKeywords() []string {
	kwFile := "segment/keywords_hits_30.txt"
	bb, err := ioutil.ReadFile(kwFile)
	if err != nil {
		fmt.Println("Read keywords error.", err)
		os.Exit(1)
	}

	content := strings.Replace(string(bb), "\r", "\n", -1)
	content = strings.Replace(content, "\n\n", "\n", -1)
	keywords := strings.Split(content, "\n")

	return keywords
}
