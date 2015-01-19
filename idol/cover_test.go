// 对IDOL做一比较全面的测试
package idol

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/kevindragon/lexiscnexec/config"
)

// 一般的检查测试
//func TestSearchFunc() {

//}

// 压力并发测试
var benchmarkQueryKeywords []string = []string{
	"公司法", "企业所得税", "合同法", "营业税", "上市", "增值税",
	"劳动合同法", "房产税", "税收征管法", "外商投资企业法",
}

// landing page上面的搜索命令
func BenchmarkQuery(b *testing.B) {
	for n := 0; n < b.N; n++ {
		query()
	}
}
func query() {
	autnTpl := `/action=query&Combine=simple&MaxResults=10&start=1` +
		`&summary=context&Characters=400&totalresults=true` +
		`&PRINTFIELDS=ID+TITLE+DRETITLE+ISSUE_DATE+ISSUE_DATE_STR` +
		`+TAXONOMY+isEnglish+effect_id+ch_eng_counter_id+provider_id+` +
		`elearning_att+TYPE_ID+ALLTYPE_ID+ALLTYPE+FILE_NAME+TEMPFILENAME+` +
		`SOURCE+ARTICLEID+PHASE+POWER_LEVEL+SELECTED_ID+CHECK_DATE+isExclusive+` +
		`CSUMMARY+REF_COUNTS+REF_LINKS+HL_LINK+HL_LINKS+SELECTED` +
		`&databasematch=law+case+ip_hottopic+module_guide+module_contract+` +
		`ep_elearning+ex_questions+ep_news+ep_news_law+ep_news_case+overview+` +
		`casecourt+newlaw+commentary+dtt+journal+profnewsletter+contract+` +
		`pgl_content+expert+hotnews+foreignlaw+newsletter+pg_template+` +
		`pg_checklists+pg_gov_form+pg_smart_chart+lawpic&SORT=Relevance+` +
		`power_level:numberincreasing+Date&Text=(({KEYWORD}))+OR+(("{KEYWORD}":TAGS))` +
		`&FieldRestriction=DRETITLE:DRECONTENT:ARTICLEID:SOURCE:EFFECT_STATUS:` +
		`COUNTRYED:AUTHOR:AUTHORSOURCE:FIRM_NAME&FieldText=(EQUAL{0,2,3,4,5,6,7,8,9,13,14,51,1}:ALLTYPE_ID` +
		`+OR+EMPTY{}:ALLTYPE_ID)+AND+(EQUAL{1,2,3,4,5,6,7,8,9,10,11,12,13,14,` +
		`15,16,17,18,19,20,21,22,23,24,25,26,27}:CHAPTER_ID+OR+EMPTY{}:CHAPTER_ID)` +
		`&sertype=sim&anylanguage=true`
	for i, keyword := range benchmarkQueryKeywords {
		fmt.Println(i)
		autnCmd := fmt.Sprintf(config.AUTN_HOST+"%s",
			strings.Replace(autnTpl, "{KEYWORD}", url.QueryEscape(keyword), -1))
		http.Get(autnCmd)
	}
}
