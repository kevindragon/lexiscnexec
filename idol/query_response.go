package idol

import (
	"encoding/xml"
	"io/ioutil"
)

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
