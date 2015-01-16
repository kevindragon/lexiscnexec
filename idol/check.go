package idol

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/kevindragon/lexiscnexec/config"
)

func getCustomKeywords() []string {
	bytes, err := ioutil.ReadFile("idol/data/custom_words.txt")
	if err != nil {
		panic(err)
	}

	c := string(bytes)
	c = strings.Replace(strings.Replace(c, "\r", "\n", -1), "\n\n", "\n", -1)

	lines := strings.Split(c, "\n")
	return lines
}

type TermInfoResponse struct {
	Response string   `xml:"response"`
	Terms    []string `xml:"responsedata>term"`
}

func getTermInfo(text string) []string {
	reqUrl := fmt.Sprintf("%s/Action=TermGetInfo&Text=%s&OnlyExisting=False&stemming=false",
		config.AUTN_HOST, url.QueryEscape(text))

	fmt.Println(reqUrl)

	resp, err := http.Get(reqUrl)
	if err != nil {
		panic(err)
	}
	xmlBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}

	var response TermInfoResponse
	xml.Unmarshal(xmlBytes, &response)

	return response.Terms
}

func CustomDict() {
	words := getCustomKeywords()
	for _, word := range words {
		// getTermInfo(word)
		terms := getTermInfo(word)
		if len(terms) == 0 || len(terms) > 1 ||
			strings.ToUpper(word) != strings.ToUpper(terms[0]) {
			fmt.Println(word, "分词不对", terms)
			continue
		}
		fmt.Println(word, "分词正确", terms)
	}
}
