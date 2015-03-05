package baidubaike

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kevindragon/html"
)

var entry_url_tpl string = "http://baike.baidu.com/search/word?word=%s"

type Page struct {
	Body []byte
}

type Result struct {
	Entities []string
	alias    string
}

func search(keyword string) ([]byte, error) {
	uri := fmt.Sprintf(entry_url_tpl, url.QueryEscape(keyword))
	pageBytes, err := get(uri)
	if err != nil {
		return pageBytes, err
	}
	return pageBytes, nil
}

func GetRelatedLaws(keyword string) (Result, error) {
	bytes, err := search(keyword)

	result := Result{[]string{""}, ""}

	r := strings.NewReader(string(bytes))
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return result, err
	}

	cardContent := doc.Find("#card-container")
	cardText := html.StripTags(cardContent.Text())
	cardEntities := getEntities(cardText)

	biTitle := doc.Find(".biTitle")
	alias := ""
	biTitle.Each(func(i int, selection *goquery.Selection) {
		other := selection.Text()
		if strings.HasPrefix(other, "别") && strings.HasSuffix(other, "称") {
			alias = selection.Next().Text()
		}
	})

	content := doc.Find("#lemmaContent-0")
	pureText := html.StripTags(content.Text())

	mainEntities := getEntities(pureText)

	entities := append(cardEntities, mainEntities...)

	result.Entities = entities
	result.alias = alias

	return result, nil
}

// 获取url指定的内容
func get(url string) ([]byte, error) {
	body := []byte{}

	res, err := http.Get(url)
	if err != nil {
		return body, err
	}

	body, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return body, err
	}

	return body, nil
}

func getEntities(s string) []string {
	re := regexp.MustCompile("《[^》]+》")
	entities := re.FindAllString(s, -1)

	return entities
}
