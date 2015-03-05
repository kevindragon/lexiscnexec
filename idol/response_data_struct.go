package idol

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
	AutnDatabase  string `xml:"database"`
	Id            int    `xml:"content>DOCUMENT>ID"`
}
