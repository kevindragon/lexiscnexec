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
	Title         string `xml:"title"`
	Issue_date    string `xml:"content>DOCUMENT>ISSUE_DATE"`
	Effect_id     int    `xml:"content>DOCUMENT>EFFECT_ID"`
	Effect_status string `xml:"content>DOCUMENT>EFFECT_STATUS"`
	Power_level   int    `xml:"content>DOCUMENT>POWER_LEVEL"`
}
