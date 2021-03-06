package court_test

import (
	//"fmt"
	"github.com/kevindragon/lexiscnexec/court"
	"testing"
)

func TestLoadDict(t *testing.T) {
	analyzer := court.NewAnalyzer()

	load := analyzer.LoadDict("dict.txt")
	if !load {
		t.Error("load dict.txt error")
	}
}

func TestLoadStandard(t *testing.T) {
	analyzer := court.NewAnalyzer()

	analyzer.LoadDict("dict.txt")
	load := analyzer.LoadStandard("standard.txt")
	if !load {
		t.Error("load standard.txt error")
	}
}

func TestInverseChain(t *testing.T) {
	analyzer := court.NewAnalyzer()

	analyzer.LoadDict("dict.txt")
	analyzer.LoadStandard("standard.txt")

	testData := map[string]string{
		"中华人民共和国江苏省高级人民法院": "江苏省高级人民法院",
		"江苏省高级人民法院":        "江苏省高级人民法院",
		"徐州市中级人民法院":        "江苏省徐州市中级人民法院",
		"玄武区人民法院":          "江苏省南京市玄武区人民法院",
		"中华人民共和国淄博市中级人民法院": "山东省淄博市中级人民法院",
	}
	for mussy, standard := range testData {
		standardName := analyzer.GetAncestor(mussy)
		if standardName != standard {
			t.Error(mussy, "can't get standard name", standard,
				"by returned", standardName)
		}
	}
}

func TestToTerms(t *testing.T) {
	analyzer := court.NewAnalyzer()

	analyzer.LoadDict("dict.txt")

	testData := map[string][]string{
		"江苏省南京中级人民法院":      []string{"江苏省", "南京", "中级人民法院"},
		"南京市江宁县人民法院":       []string{"南京市", "江宁县", "人民法院"},
		"玄武人民法院":           []string{"玄武", "人民法院"},
		"中华人民共和国淄博市中级人民法院": []string{"中华人民共和国", "淄博市", "中级人民法院"},
		"重庆市铜梁县人民法院":       []string{"重庆市", "铜梁县", "人民法院"},
	}
	for word, terms := range testData {
		actualTerms := analyzer.ToTerms(word)
		//fmt.Println(word, terms, actualTerms)
		if len(actualTerms) != len(terms) {
			t.Error("terms analyzer failed", word, terms, actualTerms)
		}
		for index := range actualTerms {
			if actualTerms[index] != terms[index] {
				t.Error("terms analyzer failed", word, terms, actualTerms)
			}
		}
	}
}

func TestManualMapping(t *testing.T) {
	analyzer := court.NewAnalyzer()
	analyzer.LoadMapping("manual_mapping.txt")

	words := map[string]string{
		"鼎城区人民法院":                               "湖南省常德市鼎城区人民法院",
		"湖滨区人民法院":                               "河南省三门峡市湖滨区人民法院",
		"湖南省资兴市人民法院":                            "湖南省资兴市人民法院",
		"重庆市第一中级人民法院":                           "重庆市第一中级人民法院",
		"北京市崇文区人民法院":                            "北京市崇文区人民法院",
		"北京市宣武区人民法院":                            "北京市宣武区人民法院",
		"大足县人民法院":                               "重庆市大足区人民法院",
		"永定区人民法院":                               "湖南省张家界市永定区人民法院",
		"重庆市綦江县人民法院":                            "重庆市綦江区人民法院",
		"阿拉善中级人民法院":                             "内蒙古自治区阿拉善盟中级人民法院",
		"安徽省蚌埠铁路运输法院":                           "蚌埠铁路运输法院",
		"安徽省高级人民法院赔偿委员会":                        "安徽省高级人民法院",
		"安徽省合肥市郊区人民法院":                          "安徽省合肥市郊区人民法院",
		"安徽省芜湖市新芜区人民法院":                         "安徽省芜湖市新芜区人民法院",
		"安徽省宣城地区中级人民法院":                         "安徽省宣城市中级人民法院",
		"安徽省宣城市宣州区（市）人民法院":                      "安徽省宣城市宣州区人民法院",
		"安徽省宣州市人民法院":                            "安徽省宣城市宣州区人民法院",
		"安顺地区中级人民法院":                            "贵州省安顺市中级人民法院",
		"巴盟中级人民法院":                              "内蒙古自治区巴彦淖尔市中级人民法院",
		"北京第一中级人民法院":                            "北京市第一中级人民法院",
		"北京市昌平区（县）人民法院":                         "北京市昌平区人民法院",
		"北京市大兴区（县）人民法院":                         "北京市大兴区人民法院",
		"北京市怀柔区（县）人民法院":                         "北京市怀柔区人民法院",
		"北京市怀柔县人民法院":                            "北京市怀柔区人民法院",
		"北京市平谷区（县）人民法院":                         "北京市平谷区人民法院",
		"北京市平谷县人民法院":                            "北京市平谷区人民法院",
		"北京市铁路运输法院":                             "北京铁路运输法院",
		"北京市铁路运输中级法院":                           "北京铁路运输中级法院",
		"北京市中级人民法院（原）":                          "北京市中级人民法院",
		"北京铁路运输法院":                              "北京铁路运输法院",
		"长沙市望城区人民法院":                            "湖南省长沙市望城区人民法院",
		"常德鼎城区人民法院":                             "湖南省常德市鼎城区人民法院",
		"常州市郊区人民法院":                             "湖南省常州市郊区人民法院",
		"巢湖市居巢区人民法院":                            "安徽省巢湖市人民法院",
		"成都高新区人民法院":                             "四川省成都高新技术产业开发区人民法院",
		"川汇区人民法院":                               "河南省周口市川汇区人民法院",
		"达州市达县人民法院":                             "四川省达州市达川区人民法院",
		"定西地区中级人民法院":                            "甘肃省定西市中级人民法院",
		"恩施州中级人民法院":                             "湖北省恩施土家族苗族自治州中级人民法院",
		"扶余县人民法院":                               "吉林省扶余市人民法院",
		"福建省福州市晋安区（郊区）人民法院":                     "福建省福州市晋安区人民法院",
		"福州铁路运输法院":                              "福州铁路运输法院",
		"甘肃省定西地区中级人民法院":                         "甘肃省定西市中级人民法院",
		"甘肃省高级人民法院赔偿委员会":                        "甘肃省高级人民法院",
		"甘肃省酒泉地区（市）中级人民法院":                      "甘肃省酒泉市中级人民法院",
		"甘肃省天水市秦城区人民法院":                         "甘肃省天水市秦州区人民法院",
		"甘肃省天水市秦州区人民法院（原甘肃省天水市秦城区人民法院）":         "甘肃省天水市秦州区人民法院",
		"甘肃省张掖地区（市）中级人民法院":                      "甘肃省张掖市中级人民法院",
		"甘肃省张掖地区中级人民法院":                         "甘肃省张掖市中级人民法院",
		"汩罗市人民法院":                               "湖南省汨罗市人民法院",
		"鼓楼区人民法院":                               "江苏省南京市鼓楼区人民法院",
		"广东省从市人民法院":                             "广东省从化市人民法院",
		"广东省番禺区市（区）人民法院":                        "广东省广州市番禺区人民法院",
		"广东省佛山区中级人民法院":                          "广东省佛山市中级人民法院",
		"广东省高明市人民法院":                            "广东省佛山市高明区人民法院",
		"广东省曲江县人民法院":                            "广东省韶关市曲江区人民法院",
		"广东省三水区（市）人民法院":                         "广东省佛山市三水区人民法院",
		"广东省汕头市潮阳市（区）人民法院":                      "广东省汕头市潮阳区人民法院",
		"广东省韶关市北江区人民法院":                         "广东省韶关市北江区人民法院",
		"广东省顺德市人民法院":                            "广东省佛山市顺德区人民法院",
		"广东省珠海市香洲区人民法院":                         "广东省珠海市香洲区人民法院",
		"广西壮族自治区百色地区（市）中级人民法院":                  "广西壮族自治区百色市中级人民法院",
		"广西壮族自治区河池市（地区）中级人民法院":                  "广西壮族自治区河池市中级人民法院",
		"广西壮族自治区江州区人民法院":                        "广西壮族自治区崇左市江州区人民法院",
		"广西壮族自治区柳州地区中级人民法院":                     "广西壮族自治区柳州市中级人民法院",
		"广西壮族自治区南宁地区中级人民法院":                     "广西壮族自治区南宁市中级人民法院",
		"广西壮族自治区南宁市新城区人民法院":                     "广西壮族自治区南宁市新城区人民法院",
		"广西壮族自治区南宁市邕宁区（县）人民法院":                  "广西壮族自治区南宁市邕宁区人民法院",
		"广西壮族自治区梧州市蝶山区人民法院":                     "广西壮族自治区梧州市万秀区人民法院",
		"广州市海事法院":                               "广州海事法院",
		"广州市铁路运输中级法院":                           "广州铁路运输中级法院",
		"贵州省安顺地区中级人民法院":                         "贵州省安顺市中级人民法院",
		"贵州省毕节地区中级人民法院":                         "贵州省毕节市中级人民法院",
		"贵州省东南苗族侗族自治州中级人民法院":                    "贵州省黔东南苗族侗族自治州中级人民法院",
		"贵州省高级人民法院赔偿委员会":                        "贵州省高级人民法院",
		"贵州省贵阳市小河区人民法院":                         "贵州省贵阳市小河区人民法院",
		"贵州省六盘水市中级人民法院赔偿委员会":                    "贵州省六盘水市中级人民法院",
		"贵州省黔南布依族苗族治州中级人民法院":                    "贵州省黔南布依族苗族自治州中级人民法院",
		"贵州省铜仁地区中级人民法院":                         "贵州省铜仁市中级人民法院",
		"海东地区中级人民法院":                            "青海省海东市中级人民法院",
		"海口市海南中级人民法院":                           "海南省海口市中级人民法院",
		"海口市海事法院":                               "海口海事法院",
		"海口市新华区人民法院":                            "海南省海口市龙华区人民法院",
		"海南省海口市海事法院":                            "海口海事法院",
		"海南省海口市新华区人民法院":                         "海南省海口市龙华区人民法院",
		"海南省海南中级人民法院":                           "海南省海南中级人民法院",
		"海南省琼山市人民法院":                            "海南省海口市琼山区人民法院",
		"海南省琼中黎族苗族自治县人民法院":                      "海南省琼中黎族苗族自治县人民法院",
		"合川市人民法院":                               "重庆市合川区人民法院",
		"河北省保定地区（市）中级人民法院":                      "河北省保定市中级人民法院",
		"河南省常垣县人民法院":                            "河南省长垣县人民法院",
		"河南省开封市顺河区人民法院":                         "河南省开封市顺河回族区人民法院",
		"河南省洛阳市廛河回族区人民法院":                       "河南省洛阳市瀍河回族区人民法院",
		"河南省洛阳市洛龙区人民法院（原河南省洛阳市郊区人民法院）":          "河南省洛阳市洛龙区人民法院",
		"河南省漯河市郾城区（县）人民法院":                      "河南省漯河市郾城区人民法院",
		"河南省牟县人民法院":                             "河南省中牟县人民法院",
		"河南省荥阳县人民法院":                            "河南省荥阳市人民法院",
		"河南省郑州市某区人民法院":                          "河南省郑州市中原区人民法院",
		"河南省郑州市中级人民法":                           "河南省郑州市中级人民法院",
		"河南省周口地区中级人民法院":                         "河南省周口市中级人民法院",
		"河南省驻马店地区中级人民法院":                        "河南省驻马店市中级人民法院",
		"河南郑州铁路运输中级法院":                          "郑州铁路运输法院",
		"鹤城区人民法院":                               "湖南省怀化市鹤城区人民法院",
		"黑龙江省阿城市人民法院":                           "黑龙江省哈尔滨市阿城区人民法院",
		"黑龙江省哈尔滨市（松花江地区）中级人民法院":                 "黑龙江省哈尔滨市中级人民法院",
		"洪江区人民法院":                               "湖南省洪江人民法院",
		"呼和浩特市铁路运输中级法院":                         "内蒙古自治区呼和浩特铁路运输中级法院",
		"湖北省恩施州中级人民法院":                          "湖北省恩施土家族苗族自治州中级人民法院",
		"湖北省恩施自治州中级人民法院":                        "湖北省恩施土家族苗族自治州中级人民法院",
		"湖北省汉江市中级人民法院":                          "湖北省汉江中级人民法院",
		"湖北省武汉市海事法院":                            "武汉海事法院",
		"湖北省武某市中级人民法院":                          "湖北省武汉市中级人民法院",
		"湖北省襄樊市中级人民法院":                          "湖北省襄阳市中级人民法院",
		"湖北省襄阳县人民法院":                            "湖北省襄阳市襄州区人民法院",
		"湖北省宜昌县人民法院":                            "湖北省宜昌市夷陵区人民法院",
		"湖北省宜市中级人民法院":                           "湖北省宜昌市中级人民法院",
		"湖南省长沙市望城区人民法院":                         "湖南省长沙市望城区人民法院",
		"湖南省鼎城区人民法院":                            "湖南省常德市鼎城区人民法院",
		"湖南省湖阳市赫山区人民法院":                         "湖南省益阳市赫山区人民法院",
		"湖南省靖州苗族侗族自治县人民法院（原靖州县人民法院）":            "湖南省靖州苗族侗族自治县人民法院",
		"湖南省娄底地区（市）中级人民法院":                      "湖南省娄底市中级人民法院",
		"湖南省娄底地区中级人民法院":                         "湖南省娄底市中级人民法院",
		"湖南省娄底市娄底区人民法院":                         "湖南省娄底市娄星区人民法院",
		"湖南省永州市零陵区（芝山区）人民法院":                    "湖南省永州市零陵区人民法院",
		"湖南省永州市芝山区人民法院":                         "湖南省永州市零陵区人民法院",
		"怀化市洪江区人民法院":                            "湖南省洪江市人民法院",
		"淮安市淮阴县人民法院":                            "江苏省淮安市淮阴区人民法院",
		"吉林省白山市八道江区人民法院":                        "吉林省白山市浑江区人民法院",
		"吉林省德惠县人民法院":                            "吉林省德惠市人民法院",
		"吉林省扶余县人民法院":                            "吉林省扶余市人民法院",
		"吉林省江源县人民法院":                            "吉林省白山市江源区人民法院",
		"吉林省延边朝鲜族自治州中级人民法院分院":                   "吉林省延边林业中级法院",
		"吉林市高新产业开发区人民法院":                        "吉林省吉林高新技术产业开发区人民法院",
		"济南铁路运输中级法院":                            "济南铁路运输中级法院",
		"嘉峪关市人法院":                               "甘肃省嘉峪关市中级人民法院",
		"江苏省常州市楼区人民法院":                          "江苏省常州市钟楼区人民法院",
		"江苏省常州市武进区（市）人民法院":                      "江苏省常州市武进区人民法院",
		"江苏省高级人民法院法院":                           "江苏省高级人民法院",
		"江苏省海门市人民法院　":                           "江苏省海门市人民法院",
		"江苏省淮安市中级人民法院（原江苏省淮阴市中级人民法院）":           "江苏省淮安市中级人民法院",
		"江苏省淮阴市清河区人民法院（原江苏省淮安市清河区人民法院）":         "江苏省淮阴市清河区人民法院",
		"江苏省金坛县人民法院":                            "江苏省金坛市人民法院",
		"江苏省连云港市云台区人民法院":                        "江苏省连云港市云台区人民法院",
		"江苏省某市人民法院":                             "江苏省常熟市人民法院",
		"江苏省某市中级人民法院":                           "江苏省泰州市中级人民法院",
		"江苏省某县人民法院":                             "江苏省宝应县人民法院",
		"江苏省某中级人民法院":                            "江苏省泰州市中级人民法院",
		"江苏省南京市建邺区人民法院":                         "江苏省南京市建邺区人民法院",
		"江苏省南京市江宁区（县）人民法院":                      "江苏省南京市江宁区人民法院",
		"江苏省南京市中级人法院":                           "江苏省南京市中级人民法院",
		"江苏省南通市崇山区人民法院":                         "江苏省南通市崇川区人民法院",
		"江苏省邳州市人法院":                             "江苏省邳州市人民法院",
		"江苏省沭阳县人民法院":                            "江苏省沭阳县人民法院",
		"江苏省睢眙县人民法院":                            "江苏省盱眙县人民法院",
		"江苏省泰州市高港区人民法院":                         "江苏省泰州市高港区人民法院",
		"江苏省通州市人民法院":                            "江苏省南通市通州区人民法院",
		"江苏省通州市通州区人民法院":                         "江苏省南通市通州区人民法院",
		"江苏省铜山市人民法院":                            "江苏省铜山县人民法院",
		"江苏省无锡高新技术开发区人民法院":                      "江苏省无锡市高新技术产业开发区人民法院",
		"江苏省吴江市人民法院　":                           "江苏省吴江市人民法院",
		"江苏省吴县市人民法院":                            "江苏省苏州市吴中区人民法院",
		"江苏省武进市人民法院":                            "江苏省常州市武进区人民法院",
		"江苏省锡山市人民法院":                            "江苏省无锡市锡山区人民法院",
		"江苏省新沂县人民法院":                            "江苏省新沂市人民法院",
		"江苏省徐州市九里区人民法院":                         "江苏省徐州市九里区人民法院",
		"江苏省扬州市郊区人民法院":                          "江苏省扬州市维扬区人民法院",
		"江西九江中级人民法院":                            "江西省九江市中级人民法院",
		"江西省波阳县人民法院":                            "江西省鄱阳县人民法院",
		"江西省抚州地区中级人民法院":                         "江西省抚州市中级人民法院",
		"江西省赣州地区（市）中级人民法院":                      "江西省赣州市中级人民法院",
		"江西省吉安地区中级人民法院":                         "江西省吉安市中级人民法院",
		"江西省九江市xx人民法院":                          "江西省九江市中级人民法院",
		"江西省南昌市郊区人民法院":                          "江西省南昌市青山湖区人民法院",
		"江西省南康市（县）人民法院":                         "江西省南康市人民法院",
		"江西省瑞金市（县）人民法院":                         "江西省瑞金市人民法院",
		"江西省上饶地区（市）中级人民法院":                      "江西省上饶市中级人民法院",
		"江西省上饶地区中级人民法院":                         "江西省上饶市中级人民法院",
		"金城江区人民法院":                              "广西壮族自治区河池市金城江区人民法院",
		"荆门市中级人法院":                              "湖北省荆门市中级人民法院",
		"靖州县人民法院":                               "湖南省靖州苗族侗族自治县人民法院",
		"开封市顺河区人民法院":                            "河南省开封市顺河回族区人民法院",
		"科尔沁区人民法院":                              "内蒙古自治区阿鲁科尔沁旗人民法院",
		"昆明市铁路运输中级法院":                           "昆明铁路运输中级法院",
		"兰州市铁路运输中级法院":                           "兰州铁路运输中级法院",
		"冷水滩区人民法院":                              "湖南省永州市冷水滩区人民法院",
		"连山区人民法院":                               "辽宁省葫芦岛市连山区人民法院",
		"连云港市云台区人民法院":                           "江苏省连云港市云台区人民法院",
		"辽宁省北宁市人民法院":                            "辽宁省北镇市人民法院",
		"辽宁省高级人民法院赔偿委员会":                        "辽宁省高级人民法院",
		"辽宁省辽河油田中级法院":                           "辽宁省辽河中级人民法院",
		"辽宁省沈阳市中级人民法院赔偿委员会":                     "辽宁省沈阳市中级人民法院",
		"辽宁省铁法市人民法院":                            "辽宁省调兵山市人民法院",
		"零陵区人民法院":                               "湖南省永州市零陵区人民法院",
		"龙胜县人民法院":                               "广西壮族自治区龙胜各族自治县人民法院",
		"陇南地区中级人民法院":                            "甘肃省陇南市中级人民法院",
		"吕梁地区中级人民法院":                            "山西省吕梁市中级人民法院",
		"洛阳市铁路运输法院":                             "洛阳铁路运输法院",
		"马坝人民法庭":                                "江苏省盱眙县人民法院",
		"内蒙古高级人民法院":                             "内蒙古自治区高级人民法院",
		"内蒙古自治区包头市昆都伦区人民法院":                     "内蒙古自治区包头市昆都仑区人民法院",
		"内蒙古自治区鄂尔多斯市中级人民法院（原内蒙古自治区伊克昭盟中级人民法院）":  "内蒙古自治区鄂尔多斯市中级人民法院",
		"内蒙古自治区乌兰察布盟中级人民法院":                     "内蒙古自治区乌兰察布市中级人民法院",
		"内蒙古自治区伊克昭盟中级人民法院":                      "内蒙古自治区伊克昭盟中级人民法院",
		"南昌市郊区人民法院":                             "江西省南昌市青山湖区人民法院",
		"南川市人民法院":                               "重庆市南川区人民法院",
		"南京市大厂区人民法院":                            "江苏省南京市六合区人民法院",
		"南京市建邺区民法院":                             "江苏省南京市建邺区人民法院",
		"南京市建邺区人民法院":                            "江苏省南京市建邺区人民法院",
		"南京市江宁县人民法院":                            "江苏省南京市江宁区人民法院",
		"南京市玄武区某某法院":                            "江苏省南京市玄武区人民法院",
		"南宁市新城区人民法院":                            "广西壮族自治区南宁市新城区人民法院",
		"南山区人民法院":                               "广东省深圳市南山区人民法院",
		"南通市崇州区人民法院":                            "江苏省南通市崇川区人民法院",
		"宁夏高级法院":                                "宁夏回族自治区高级人民法院",
		"宁夏高级人民法院":                              "宁夏回族自治区高级人民法院",
		"宁夏回族自治区中卫市（县）人民法院":                     "宁夏回族自治区中卫市中级人民法院",
		"宁夏回族自治区中卫市城区人民法院":                      "宁夏回族自治区中卫市沙坡头区人民法院",
		"普陀区人民法院":                               "上海市普陀区人民法院",
		"前郭县人民法院":                               "吉林省前郭尔罗斯蒙古族自治县人民法院",
		"青岛海事法院":                                "青岛海事法院",
		"青岛市海事法院":                               "青岛海事法院",
		"青岛铁路运输法院":                              "青岛铁路运输法院",
		"青海省高级人民法院赔偿委员会":                        "青海省高级人民法院",
		"青海省海东地区中级人民法院":                         "青海省海东市中级人民法院",
		"青海省海南州中级人民法院":                          "青海省海南藏族自治州中级人民法院",
		"青海中级人民法院":                              "青海省高级人民法院",
		"庆阳地区中级人民法院":                            "甘肃省庆阳市中级人民法院",
		"衢州市衢县人民法院":                             "浙江省衢县人民法院",
		"如皋市民法院":                                "江苏省如皋市人民法院",
		"三中级人民法院":                               "北京市第三中级人民法院",
		"山东省滨州地区（市）中级人民法院":                      "山东省滨州市中级人民法院",
		"山东省德州地区（市）中级人民法院":                      "山东省德州市中级人民法院",
		"山东省高级人民法院赔偿委员会":                        "山东省高级人民法院",
		"山东省济南铁路运输中级法院":                         "济南铁路运输法院",
		"山东省济宁市中级人民法院赔偿委员会":                     "山东省济宁市中级人民法院",
		"山东省青岛市中院":                              "山东省青岛市中级人民法院",
		"山东省荣城市人民法院":                            "山东省荣成市人民法院",
		"山东省烟台市中级人某法院":                          "山东省烟台市中级人民法院",
		"山东市东营县人民法院":                            "山东省东营市东营区人民法院",
		"山西省运城地区中级人民法院":                         "山西省运城市中级人民法院",
		"陕西省高级民法院":                              "陕西省高级人民法院",
		"陕西省高级人民法院赔偿委员会":                        "陕西省高级人民法院",
		"陕西省商洛市（地区）中级人民法院":                      "陕西省商洛市中级人民法院",
		"陕西省铜川市郊区人民法院":                          "陕西省铜川市郊区人民法院",
		"陕西省铜川市耀州区人民法院（原陕西省耀县人民法院）":             "陕西省铜川市耀州区人民法院",
		"陕西省吴旗县人民法院":                            "陕西省吴起县人民法院",
		"陕西省西安市长安区（县）人民法院":                      "陕西省西安市长安区人民法院",
		"上海第二中级人民法院":                            "上海市第二中级人民法院",
		"上海第一中级人民法院":                            "上海市第一中级人民法院",
		"上海高级人民法院":                              "上海市高级人民法院",
		"上海海市法院":                                "上海海事法院",
		"上海市奉贤区（县）人民法院":                         "上海市奉贤区人民法院",
		"上海市海事法院":                               "上海海事法院",
		"上海市闵行区（上海县）人民法院":                       "上海市闵行区人民法院",
		"上海市南汇县人民法院":                            "上海市南汇区人民法院",
		"上海市浦东区人民法院":                            "上海市浦东新区人民法院",
		"上海市浦东人民法院":                             "上海市浦东新区人民法院",
		"上海市浦东新区人民检察院":                          "上海市浦东新区人民法院",
		"上海市青浦区（县）人民法院":                         "上海市青浦区人民法院",
		"上海市松江县人民法院":                            "上海市松江区人民法院",
		"上海市铁路运输中级法院":                           "上海铁路运输中级法院",
		"上饶地区中级人民法院":                            "江西省上饶市中级人民法院",
		"上饶市波阳县人民法院":                            "江西省鄱阳县人民法院",
		"韶关市北江区人民法院":                            "广东省韶关市北江区人民法院",
		"韶关市曲江县人民法院":                            "广东省韶关市曲江区人民法院",
		"邵阳城步县人民法院":                             "湖南省城步苗族自治县人民法院",
		"邵阳市城步县人民法院":                            "湖南省城步苗族自治县人民法院",
		"邵阳市郊区人民法院":                             "湖南省邵阳市郊区人民法院",
		"邵阳市中级人民法院裁定书":                          "湖南省邵阳市中级人民法院",
		"沈阳是中级人民法院":                             "辽宁省沈阳市中级人民法院",
		"石家庄铁路运输法院":                             "石家庄铁路运输法院",
		"石嘴山市中级人民法院赔偿委员会":                       "宁夏回族自治区石嘴山市中级人民法院",
		"沭阳县人民法院":                               "江苏省沭阳县人民法院",
		"思茅地区中级人民法院":                            "云南省普洱市中级人民法院",
		"四川成都高新区人民法院":                           "四川省成都高新技术产业开发区人民法院",
		"四川省XX县人民法院":                            "四川省安岳县人民法院",
		"四川省达县人民法院":                             "四川省达州市达川区人民法院",
		"四川省高级人民法院赔偿委员会":                        "四川省高级人民法院",
		"四川省广元市市中区人民法院":                         "四川省广元市利州区人民法院",
		"四川省泸洲市中级人民法院":                          "四川省泸州市中级人民法院",
		"四川省眉山地区中级人民法院":                         "四川省眉山市中级人民法院",
		"四川省遂宁市市中区人民法院":                         "四川省遂宁市市中区人民法院",
		"四川省新都县人民法院":                            "四川省成都市新都区人民法院",
		"四川省雅安地区中级人民法院":                         "四川省雅安市中级人民法院",
		"四川省资阳地区中级人民法院":                         "四川省资阳市中级人民法院",
		"松江县人民法院":                               "上海市松江区人民法院",
		"苏州市江区人民法院":                             "江苏省苏州市吴江区人民法院",
		"苏州工业园人民法院":                             "江苏省苏州市工业园区人民法院",
		"苏州市金阊区民法院":                             "江苏省苏州市金阊区人民法院",
		"苏州市金阊区人民院":                             "江苏省苏州市金阊区人民法院",
		"苏州市吴江区江人民法院":                           "江苏省苏州市吴江区人民法院",
		"苏州市吴江市江人民法院":                           "江苏省吴江市人民法院",
		"宿迁市沭阳县人民法院":                            "江苏省沭阳县人民法院",
		"遂宁市市中区人民法院":                            "四川省遂宁市市中区人民法院",
		"泰州市高港区人民法院":                            "江苏省泰州市高港区人民法院",
		"泰州市海陵区某某法院":                            "江苏省泰州市海陵区人民法院",
		"天津高院":                                  "天津市高级人民法院",
		"天津市宝坻县人民法院":                            "天津市宝坻区人民法院",
		"天津市第二中级人民法院赔偿委员会":                      "天津市第二中级人民法院",
		"天津市二中级人民法院":                            "天津市第二中级人民法院",
		"天津市海事法院":                               "天津海事法院",
		"天津市塘沽区人民法院":                            "天津市滨海新区人民法院",
		"天津市武清县人民法院":                            "天津市武清区人民法院",
		"天津市中级人民法院（原）":                          "天津市第一中级人民法院",
		"天水市中极人民法院":                             "甘肃省天水市中级人民法院",
		"通州市人民法院":                               "江苏省南通市通州区人民法院",
		"通州市通州区人民法院":                            "江苏省南通市通州区人民法院",
		"铜仁地区万山特区人民法院":                          "贵州省铜仁市万山区人民法院",
		"威海市中级法院":                               "山东省威海市中级人民法院",
		"威海中级法院":                                "山东省威海市中级人民法院",
		"渭南市白水县人民法院　":                           "陕西省白水县人民法院",
		"温州市中级人民法院民三庭":                          "浙江省温州市中级人民法院",
		"乌鲁木齐市铁路运输法院":                           "乌鲁木齐铁路运输法院",
		"无锡市xx人民法院":                             "江苏省无锡市锡山区人民法院",
		"无锡市湖滨区人民法院":                            "江苏省无锡市滨湖区人民法院",
		"吴旗县人民法院":                               "陕西省吴起县人民法院",
		"梧州市蝶山区人民法院":                            "广西壮族自治区梧州市万秀区人民法院",
		"武汉市海事法院":                               "武汉海事法院",
		"武汉市黄坡区人民法院":                            "湖北省武汉市黄陂区人民法院",
		"武汉中院":                                  "湖北省武汉市中级人民法院",
		"武陵区人民法院":                               "湖南省常德市武陵区人民法院",
		"武陵源区人民法院":                              "湖南省张家界市武陵源区人民法院",
		"武威市铁路运输法院":                             "武威铁路运输法院",
		"舞钢市X区人民法院":                             "河南省舞钢市人民法院",
		"锡山市人民法院":                               "江苏省无锡市锡山区人民法院",
		"厦门海事法院":                                "厦门海事法院",
		"厦门市海事法院":                               "厦门海事法院",
		"厦门市开元区人民法院":                            "福建省厦门市开元区人民法院",
		"襄城县人民":                                 "河南省襄城县人民法院",
		"襄城县人民法":                                "河南省襄城县人民法院",
		"襄樊市中级人民法院":                             "湖北省襄阳市中级人民法院",
		"孝南区人民法院":                               "湖北省孝感市孝南区人民法院",
		"新疆生产建设兵团农二师中级人民法院":                     "新疆生产建设兵团第二师中级人民法院",
		"新疆生产建设兵团农六师中级人民法院":                     "新疆生产建设兵团第六师中级人民法院",
		"新疆生产建设兵团农七师中级人民法院":                     "新疆生产建设兵团第七师中级人民法院",
		"新疆生产建设兵团农三师中级人民法院":                     "新疆生产建设兵团第三师中级人民法院",
		"新疆生产建设兵团农十师中级人民法院":                     "新疆生产建设兵团第十师中级人民法院",
		"新疆生产建设兵团农四师中级人民法院":                     "新疆生产建设兵团第四师中级人民法院",
		"新疆生产建设兵团农五师中级人民法院":                     "新疆生产建设兵团第五师中级人民法院",
		"新疆生产建设兵团农一师中级人民法院":                     "新疆生产建设兵团第一师中级人民法院",
		"新疆维吾尔自治区克拉玛依市中级人民法院民事判决书":              "新疆维吾尔自治区克拉玛依市中级人民法院",
		"新疆维吾尔自治区克孜勒苏柯尔克孜自治州中级人法院":              "新疆维吾尔自治区克孜勒苏柯尔克孜自治州中级人民法院",
		"新疆维吾尔自治区农八师中级人民法院":                     "新疆生产建设兵团第八师中级人民法院",
		"新疆维吾尔自治区农六师中级人民法院":                     "新疆生产建设兵团第六师中级人民法院",
		"新疆维吾尔自治区生产建设兵团农八师中级人民法院":               "新疆生产建设兵团第八师中级人民法院",
		"新疆维吾尔自治区乌鲁本齐市中级入民法院":                   "新疆维吾尔自治区乌鲁木齐市中级人民法院",
		"新疆维吾尔自治区新疆生产建设兵团农十师中级人民法院":             "新疆生产建设兵团第十师中级人民法院",
		"新疆维吾尔自治区叶城县人民法院":                       "新疆维吾尔自治区叶城县人民法院",
		"新疆维吾尔自治区伊犁地区中级人民法院":                    "新疆维吾尔自治区高级人民法院伊犁哈萨克自治州分院",
		"新疆维吾尔族自治区高级人民法院":                       "新疆维吾尔自治区高级人民法院",
		"新疆自治区高级人民法院":                           "新疆维吾尔自治区高级人民法院",
		"新乡市常垣县人民法院":                            "河南省长垣县人民法院",
		"新乡市原野县人民法院":                            "河南省原阳县人民法院",
		"刑事裁定书（2009）皖刑终字第0351号原公诉机关安徽省亳州市人民检察院": "安徽省高级人民法院",
		"刑事判决书":                     "河南省济源市人民法院",
		"徐州市九里区人民法院":                "江苏省徐州市九里区人民法院",
		"雅安地区名山县人民法院":               "四川省雅安市名山区人民法院",
		"扬州市郊区人民法院":                 "江苏省扬州市维扬区人民法院",
		"宜昌市宜昌县人民法院":                "湖北省宜昌市夷陵区人民法院",
		"玉树藏族自治州中级人法院":              "青海省玉树藏族自治州中级人民法院",
		"玉溪中院":                      "云南省玉溪市中级人民法院",
		"云南省保山地区中级人民法院":             "云南省保山市中级人民法院",
		"云南省呈贡县人民法院":                "云南省昆明市呈贡区人民法院",
		"云南省丽江地区中级人民法院":             "云南省丽江市中级人民法院",
		"云南省临沧地区中级人民法院":             "云南省临沧市中级人民法院",
		"云南省潞西市人民法院":                "云南省芒市人民法院",
		"云南省石林（路南）彝族自治县人民法院":        "云南省石林彝族自治县人民法院",
		"云南省思茅地区中级人民法院":             "云南省普洱市中级人民法院",
		"云南省文山县人民法院":                "云南省文山壮族苗族自治州中级人民法院",
		"云南省昭通地区（市）中级人民法院":          "云南省昭通市中级人民法院",
		"浙江省杭州市萧山区（市）人民法院":          "浙江省杭州市萧山区人民法院",
		"浙江省杭州市余杭区（市）人民法院":          "浙江省杭州市余杭区人民法院",
		"浙江省杭州市中级人民法":               "浙江省杭州市中级人民法院",
		"浙江省湖州市城郊人民法院":              "浙江省湖州市城郊人民法院",
		"浙江省金华市级人民法院":               "浙江省金华市中级人民法院",
		"浙江省金华市中级人民院":               "浙江省金华市中级人民法院",
		"浙江省丽水地区中级人民法院":             "浙江省丽水市中级人民法院",
		"浙江省丽水市中级人民法院（原浙江省丽水市人民法院）": "浙江省丽水市中级人民法院",
		"浙江省宁波市海事法院":                "宁波海事法院",
		"浙江省衢县人民法院":                 "浙江省衢县人民法院",
		"浙江省萧山市人民法院":                "浙江省杭州市萧山区人民法院",
		"浙江省鄞县人民法院":                 "浙江省宁波市鄞州区人民法院",
		"镇江经济开发区人民法院　":              "江苏省镇江市经济开发区人民法院",
		"郑州市郑州矿区人民法院":               "河南省郑州市郑州矿区人民法院",
		"郑州中院":                      "河南省郑州市中级人民法院",
		"中国人民共和国最高人民法院":             "最高人民法院",
		"中华人民共和国海南中级人民法院":           "海南省第一中级人民法院",
		"中华人民共和国青岛海事法院":             "青岛海事法院",
		"中华人民共和国上海海市法院":             "上海海事法院",
		"中华人民共和国上海市第二中级人民院":         "上海市第二中级人民法院",
		"中华人民共和国厦门海事法院":             "厦门海事法院",
		"中华人民共和国重庆市第一中级人民法院":        "重庆市中级人民法院",
		"中华人民共和国最高人民人民法院":           "最高人民法院",
		"重庆南川市人民法院":                 "重庆市南川区人民法院",
		"重庆市城口县某法院":                 "重庆市城口县人民法院",
		"重庆市大足县人民法院":                "重庆市大足区人民法院",
		"重庆市第四中级人民法院赔偿委员会":          "重庆市第四中级人民法院",
		"重庆市第四中级人民法院网":              "重庆市第四中级人民法院",
		"重庆市第五中级人民法院国家赔偿委员会":        "重庆市第五中级人民法院",
		"重庆市第五中级人民法院赔偿委员会":          "重庆市第五中级人民法院",
		"重庆市合川市人民法院":                "重庆市合川区人民法院",
		"重庆市某县人民法院":                 "重庆市大足区人民法院",
		"重庆市南川区（市）人民法院":             "重庆市南川区人民法院",
		"重庆市彭水县人民法院":                "重庆市彭水苗族土家族自治县人民法院",
		"重庆市石柱县土家族自治县人民法院":          "重庆市石柱土家族自治县人民法院",
		"重庆市秀山土家苗族自治区人民法院":          "重庆市秀山土家族苗族自治县人民法院",
		"重庆市秀山县人民法院":                "重庆市秀山土家族苗族自治县人民法院",
		"重庆市永川区（县）人民法院":             "重庆市永川区人民法院",
		"重庆市永川市人民法院":                "重庆市永川区人民法院",
		"重庆市酉阳县人民法院":                "重庆市酉阳土家族苗族自治县人民法院",
		"重庆市渝某区人民法院":                "重庆市渝北区人民法院",
		"重庆市云阳县人民":                  "重庆市云阳县人民法院",
		"重庆市运输法院":                   "重庆铁路运输法院",
		"重庆永川市人民法院":                 "重庆市永川区人民法院",
		"株洲市元区人民法院":                 "湖南省株洲市天元区人民法院",
		"珠海市香洲区人民法院":                "广东省珠海市香洲区人民法院",
		"资阳地区中级人民法院":                "四川省资阳市中级人民法院",
		"最高人民法院;中国法学应用研究所":          "最高人民法院",
		"最高人民法院;中国应用法学研究所":          "最高人民法院",
		"最高人民法院人民法院":                "最高人民法院",
		"最高人民法院中国应用法学研究所":           "最高人民法院",
	}

	for name, word := range words {
		manualStand := analyzer.GetFromMapping(name)
		if manualStand != word {
			t.Errorf("Can't found from mapping. Name: %s, word: %s, found: %s\n", name, word, manualStand)
		}
	}
}
