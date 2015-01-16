// 测试bias是否会对性能有影响
package idol

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func TestBias() {
	keywords := []string{
		"公司法", "劳动法", "劳动合同法", "婚姻法", "劳务派遣", "个人所得税", "刑法",
	}

	autnCmdTpl := "http://10.123.4.210:9003/action=Query&databasematch=law&text=%s"

	//biasTpl := fmt.Sprintf("&FieldText=BIAS{1,3,2}:POWER_LEVEL+AND+BIAS{%d,%d,2}:DREDATE",
	//	time.Now().Unix(), 10*12*31*24*3600)
	biasTpl := "&FieldText=BIAS{1,3,2}:POWER_LEVEL"

	noBiasDuration := make([]int64, 0)
	haveBiasDuration := make([]int64, 0)

	for _, keyword := range keywords {
		haveBiasCmd := fmt.Sprintf(autnCmdTpl, url.QueryEscape(keyword)) + biasTpl
		noBiasCmd := fmt.Sprintf(autnCmdTpl, url.QueryEscape(keyword))

		noBiasSt := time.Now().UnixNano()
		http.Get(noBiasCmd)
		noBiasEt := time.Now().UnixNano()
		noBiasDuration = append(noBiasDuration, noBiasEt-noBiasSt)

		haveBiasSt := time.Now().UnixNano()
		http.Get(haveBiasCmd)
		haveBiasEt := time.Now().UnixNano()
		haveBiasDuration = append(haveBiasDuration, haveBiasEt-haveBiasSt)
	}

	var noBiasSum, haveBiasSum int64
	for _, duration := range noBiasDuration {
		noBiasSum += duration
	}
	for _, duration := range haveBiasDuration {
		haveBiasSum += duration
	}

	fmt.Println(noBiasSum, haveBiasSum)
}
