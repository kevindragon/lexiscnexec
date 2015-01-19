// 处理indexergetstatus的结果，做一些统计
package idol

import (
	"encoding/xml"
	"io/ioutil"
	"strings"
	"time"
)

type IndexerGetStatus struct {
	Items []IndexerGetStatusItem `xml:"responsedata>item"`
}
type IndexerGetStatusItem struct {
	StartTime           customTime `xml:"start_time"`
	EndTime             customTime `xml:"end_time"`
	Description         string     `xml:"description"`
	IndexCommand        string     `xml:"index_command"`
	PercentageProcessed int        `xml:"percentage_processed"`
	DocumentsProcessed  int        `xml:"documents_processed"`
}
type customTime struct {
	time.Time
}

func (c *customTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = "2006/01/02 15:04:05"
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return nil
	}
	*c = customTime{parse}
	return nil
}

func IndexPerformance(file string) float64 {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var status IndexerGetStatus
	xml.Unmarshal(bytes, &status)

	var duration int64
	var documents int64
	for _, item := range status.Items {
		if item.Description != "Finished" {
			continue
		}
		if !strings.HasPrefix(item.IndexCommand, "/DREADD") {
			continue
		}

		timeDuration := item.EndTime.Unix() - item.StartTime.Unix()
		duration += timeDuration

		docCount := int64(item.DocumentsProcessed)
		documents += docCount
	}

	return float64(documents) / float64(duration)
}
