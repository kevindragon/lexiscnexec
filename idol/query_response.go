package idol

import (
	"encoding/xml"
	"io/ioutil"
)

type Response struct {
	Action       string   `xml:"action"`
	Response     string   `xml:"response"`
	ResponseData RespData `xml:"responsedata"`
}

type RespData struct {
	Numhits int   `xml:"numhits"`
	Hits    []Hit `xml:"hit"`
}

type Hit struct {
	AutnReference string `xml:"reference"`
	Id            int    `xml:"content>DOCUMENT>ID"`
}

func ParseXMLFile(filename string) (*Response, error) {
	var response Response

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return &response, err
	}

	err = xml.Unmarshal(bytes, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}
