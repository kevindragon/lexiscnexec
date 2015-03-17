// 把overview里面提到的法规全部加上overview文章的标题作为tag
package main

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kevindragon/html"
)

const (
	allid string = "11825,11826,11843,11955,11962,11973,11985,11991,12004,13698,13697,13702,13701,14562,12918,12921,12923,12927,13125,13134,13141,13156,13161,13164,13170,13172,13688,13195,13198,13200,13202,13207,13686,13687,13689,13690,13691,13692,13849,13850,13852,14042,14046,14048,14050,14056,14248,14548,14549,14793,14794,14795,14796,14797,14949,14950,14951,14952,14953,14954,14955,14956,15044,15045,15046,15047,15054,15055,15056,15744,15745,15746,15816,15817,15818,15819,15820,15821,15822,15823,15975,15976,15977,15978,16535,16536,16537,16538,16539,16542,18886,18887,18888,18889,18890,18891,18892,18893,18894,18895,18896,18897,18898,18899,18900,18901,19097,19098,19102,19104,19106,19107,19109,19111,19364,19365,19366,19367,19368,19370,19369,19371,19372,19392,19393,19394,19395,19396,19397,19398,19582,19599,19609,19632,19634,19638,19642,19661,19670,19699,19704,19707,19709,19710,19711,19779,19780,19788,19792,19811,20318,20322,20331,20334,20337,20642,20643,20644,20647,20654,20676,20677,20701,20702,20703,20704,20705,20706,20707,20755,20757,20758,20759,20760,20761,20762,26921,26922,26923,26924,26962,27232,27234,27316,27317,27318,27320,27321,27779,27783,27788,27789,27790,27867,27868,27869,27870,27871,27872,27873,28397,28629,28630,28633,28634,28994,28995,28996,28997,28998,36044,36063,36155,36156,36157,36274,36275,36276,36277,36278,36294,36305,36306,36308,36341,36342,36343,36344,36345,36346,36347,36348,36349,36350,36351,36352,36353,36354,36359,36360,36361,36362,36363,36401,36402,36403,36404,36405,36434,36436,36437,36438,36439,36440,36442,36453,36454,36455,36456,36590,36591,36592,36609,36613,38533,38613,38614,38615,38616,38617,38618,38619,38620,38621,38622,38623,38624,38625,38626,38627,38628,38629,38630,38631,38632,38643,38633,38635,38636,38637,38639,38641,38642,38644,38645,38646,38647,38648,38649,38650,38651,38652,38653,38654,38926,38928,38929,38930,38935,38936,38939,38940,38943,38944,38967,38945,38946,38947,38948,38949,38950,38951,38952,38953,38954,38955,38956,38957,38958,38959,38960,38961,38962,38963,38964,38965,38966,38968,41009,41010,41011,41012,41013,41014,41015,41016,41017,41018,41019,41020,41021,41022,41023,41024,41025,41026,41027,41028,41029,41030,41031,41032,41033,41034,41035,41036,41037,41038,41039,41040,41041,41042,41043,41044,41045,41047,41057,41058,41059,41060,41061,41062,41063,41064,41065,41066,41067,41068,41069,41078,41077,41075,41074,41076,41070,41071,41072,41073,41079,41080,41081,41082,41083,41084,41085,41086,41087,41088,41089,41090,41091,41092,41093,41094,41095,41096,41097,41098,41099,41100,41101,41102,41103,41104,41105,41106,41107,41108,41109,41110,41111,41112,41113,41114,41115,41116,41117,41118,41119,41120,41343,43849,44133,43851,43859,43861,43862,43863,43864,43865,43866,43878,43879,43880,43881,43905,50306,43907,43909,43910,43911,43917,43919,43924,43925,43955,43956,43957,43958,43972,43973,43974,44025,44026,44027,44028,44030,44031,44032,44033,44034,44035,44036,44037,44038,44039,44116,44117,44118,44119,44120,44121,44124,44125,44126,44134,44135,44136,44138,44170,44171,44177,44178,44179,44180,44181,44182,44183,44184,44185,44186,44187,44188,44189,44190,44195,44255,44256,44257,44265,44304,44305,44306,44307,44312,44315,44320,44321,44322,44323,44590,44593,44594,44595,44596,44597,44645,44646,44863,44864,44866,44954,44957,44987,44989,44990,44991,45017,45018,45019,45020,45024,45025,45026,45027,45028,45037,45094,45097,45099,45100,45201,45203,45204,45206,45211,45212,47449,47450,47451,47452,47522,47523,50173,48825,47983,47987,47991,47992,47993,48050,48051,48052,48053,48055,48061,48065,48233,48236,48237,48238,48239,48279,48280,48281,48282,48283,48284,48285,48286,48287,48289,48290,48291,48293,48294,48295,48297,48298,48299,48300,48301,48403,48404,48456,48457,48458,48459,48485,48486,48601,48602,48603,48604,48606,48608,48610,48611,48614,48615,48616,48617,48618,48713,48714,48724,48729,48842,48917,48969,48970,49082,49208,49328,49557,49558,49561,49565,49567,49571,49583,49585,49652,49653,49654,49664,49699,49725,49760,49803,49805,49806,49809,49834,49835,49836,50067,50125,50126,50127,50129,50130,50170,50174,50175,50176,50177,50194,50205,50207,50208,50209,50210,50211,50212,50213,50215,50216,50217,50218,50244,50245,50246,50247,50248,50251,50252,50253,50254,50273,50274,52850,52899,52966,52977,53042,53043,53044,53045,53855,54160,54302,54670,54671,54937,54950,60412,60389,60390,60391,60392,60393,60394,60395,60396,60397,60398,60399,60400,60401,60402,60405,60407,60410,60413,60415,60416,60417,60418,60419,60420,60421,60422,60423,60424,60425,60426,60427,60428,60430,60431,60432,60433,60464,60465,60466,60467,60468,60469,60508,60509,60510,60511,60560,60561,60562,60563,60564,60565,60566,60567,60568,60569,60570,60571,60572,60573,60644,60647,60648,60649,60651,60653,60654,60656,60658,60659,60661,60662,60668,60673,60695,60697,60698,60699,60700,60701,60702,60703,60704,60705,60707,60708,60709,60710,60711,60712,60713,60714,60715,60716,60717,60718,60719,60720,60733,60734,60735,60736,60737,60738,60739,60740,60741,60742,60743,60744,60745,60746,60747,60748,60749,60751,60753,60755,60756,60757,60758,60773,60774,60775,60776,60777,60778,60779,60780,60781,60782,60783,60784,60785,60786,60787,60799,60807,60813,60814,60815,60817,60818,60819,60820,60821,60822,60823,60824,60825,60826,61043,61044,61045,61046,61047,61048,61049,61050,61194,61195,61196,61197,61198,61199,61200,61201,61205,61206,61207,61208,61209,61210,61211,61212,61213,61214,61215,61216,61217,61218,61219,61220,61221,61222,61223,61224,61225,61226,61227,61228,61246,61247,61248,61249,61250,61251,61252,61253,61272,61273,61280,61281,61284,61285,61286,61287,61291,61293,61297,61298,61299,61300,61301,61302,61303,61304,61305,61306,61307,61308,61309,61310,61311,61312,61313,61314,61315,61316,61317,61318,61319,61320,61321,61322,61323,61324,61325,61326,61327,61328,61329,61330,61331,61332,61333,61334,61335,61336,61337,61338,61339,61340,61341,61342,61343,61344,61362,61363,61364,61365,61368,61369,61370,61371,61372,61373,61374,61375,61376,61377,61378,61379,61380,61381,61384,61385,61396,61397,61398,61399,61400,61401,61402,61403,61404,61405,61406,61407,61408,61409,61410,61411,61412,61413,61414,61415,61416,61417,61418,61419,61420,61421,61422,61423,61424,61425,61426,61427,61428,61429,61443,61444,61445,61446,61447,61448,61449,61450,61451,61452,61453,61454,61455,61456,61457,61458,61459,61460,61461,61462,61463,61464,61465,61466,61477,61478,61479,61480,61481,61482,61502,61508,61509,61510,61511,61547,61548,61558,61559,61560,61562,61563,61564,61565,61566,61567,61568,61569,61570,61586,61587,61588,61589,61590,61591,61599,61600,61603,61604,61618,61619,61634,61635,61636,61637,61690,61691,61692,61693,61694,61695,61696,61697,61698,61699,61700,61819,61820,61831,61832,61833,61834,61835,61836,61837,61838,61839,61840,61841,61842,61843,61878,61879,61919,61938,61944,61948,61955,61956,61963,61964,61965,61966,61967,61968,61969,61970,61971,61972,61973,61974,61985,61986,61987,61988,61989,61990,61991,61992,61993,61994,61995,61996,61997,61998,61999,62000,62001,62002,62003,62004,62005,62006,62007,62008,62009,62010,62011,62012,62013,62014,62443,62444,62580,62581,62582,62583,62584,62585,62586,62587,62588,62593,62633,62634,62637,62639,62653,62654,62655,62656,62657,62658,62659,62660,62661,62662,62673,62675,62678,62680,62684,62687,64859,64860,65310,65311,65322,65324,65325,65326,65532,65534,65535,65536,65549,65551,65597,65598,65599,65600,65615,65616,65623,65625,65693,65694,65731,65732,65733,65734,65744,65807,65808,65809,65810,65811,65812,65813,65838,65841,65844,65845,65846,65847,65987,65988,65989,65990,66017,66018,66019,66020,66021,66026,66027,66061,66066,66068,66071,66074,66075,66076,66077,66078,66079,66080,66081,66082,66085,66086,66087,66088,66205,66206,66207,66208,66209,66210,66211,66212,66213,66214,66215,66216,66217,66218,66219,66220,66222,66223,66224,66225,66226,66227,66228,66229,66232,66233,66234,66235,66285,66296,66297,66298,66299,66318,66322,66323,66324,66325,66327,66331,66334,66335,66336,66337,66338,66339,66430,66431,66432,66433,66434,66435,66436,66437,66438,66439,66450,66451,66456,66538,66539,66541,66542,66543,66548,66553,66555,66558,66560,66561,66563,66564,66572,66585,66590,66597,66599,66600,66602,66610,66624,66627,66628,66633,66636,66640,66642,66643,66658,66666,66667,66670,66673,66680,66686,66700,66703,66704,66714,66715,66793,66794,66802,66803,66848,66849,66862,66864,66867,66868,66869,66871,66873,66874,66877,66881,66882,66888,66892,66905,67009,67011,67243,67244,67248,67249,67250,67251,67252,67253,67254,67255,67256,67257,67270,67279,67280,67281,67282,67283,67284,67285,67286,67287,67288,67289,67290,67291,67292,67293,67297,67298,67300,67302,67342,67343,67344,67345,67346,67347,67355,67356,67359,67360,67396,67397,67725,67730,69221,69244,69254,69460,69463,69469,69484,69600,69601,69602,69603,69604,69605,69660,72880,69903,69957,70002,70003,70059,70470,70535,72343,72345,72359,72364,72368,72384,72877,72878,72879,72884,72885,72886,72887,72888,72889,72901,72904,72917,72919,72922,72926,73165,73167,73171,73173,73211,73221,73223,73230,73232,73234,73266,73267,73268,73269,73270,73271,73272,73273,73274,73275,73276,73279,73288,73289,73291,73296,73301,73304,73309,73310,73311,73312,73348,73349,73353,73362,73363,73365,73367,73370,73441,73442,73730,73444,73447,73448,73449,73450,73451,73452,73474,73475,73497,73512,73515,73522,73523,73524,73526,73527,73528,73529,73530,73531,73533,73555,73556,73571,73576,73579,73592,73593,73595,73597,73598,73605,73608,73721,73722,73723,73724,73726,73727,73732,73738,73739,73750,73759,73768,73769,73774,73776,73838,73839,73840,73841,73842,73843,73845,73851,73877,73878,73879,73880,73881,73882,73883,73884,73945,73946,73947,73948,73949,73950,73951,73952,73969,73974,73976,73977,73996,73997,73998,73999,74000,74001,74002,74003,74004,74005,74006,74007,74008,74009,74010,74011,74012,74013,74014,74020,74025,74026,74028,74032,74035,74036,74037,74038,74039,74040,74041,74042,74045,74054,74063,74064,74067,74069,74070,74072,74074,74075,74076,74080,74081,74082,74086,74087,74088,74089,74092,74094,74096,74097,74125,74126,74154,74155,74156,74157,74158,74159,74160,74161,74162,74163,74164,74165,74183,74184,74189,74191,74197,74198,74359,74360,74382,74383,74384,74385,74386,74387,74388,74389,74390,74391,74392,74393,74394,74395,74396,74397,74398,74399,74400,74401,74402,74403,74404,74405,74406,74407,74409,74410,74411,74412,74413,74414,74415,74416,74417,74418,74429,74430,74441,74442,74443,74444,74445,74447,74450,74453,74455,74456,74457,74458,74459,74460,74468,74469,74470,74471,74472,74473,74474,74475,74477,74484,75684,75686,75687,75698,74489,74490,74493,74494,74498,74500,74502,74503,74505,74506,74507,74509,74511,74512,74513,74515,74517,74519,74544,74546,74548,74549,74553,74554,74560,74562,74564,74566,74569,74570,74579,74586,74587,74589,74603,74604,74616,74619,74623,74625,74636,74639,74645,74646,74647,74648,74709,74713,74716,74717,74718,74719,74720,74721,74722,74723,74724,74725,74727,74728,74733,74734,74735,74736,74737,74738,74739,74741,74753,74754,74759,74764,74765,74766,74770,74771,74773,74774,74781,74791,74798,74799,74801,74802,74803,74804,74945,74946,74947,74948,74949,74950,74951,74952,74953,74954,74955,74956,74957,74958,74959,74960,74969,74970,74971,74972,74973,74974,74975,74977,74978,74979,74980,74981,74982,74983,75240,75241,75248,75249,75312,75313,75314,75315,75316,75317,75319,75320,75321,75322,75386,75387,75389,75388,75390,75391,75589,75590,75591,75592,75643,75644,75682,75683,75685,75688,75689,75699,75700,75704,75713,75714,88851,88852,88853,88854,88855,88856,88857,88858,88859,88860,88861,88862,88863,88864,88865,88866,88867,88868,88869,88870,88871,88872,88873,88874,88875,88876,88877,88878,88879,88880,88881,88882,88883,88884,88885,88886,88887,88888,88889,88890,88891,88892,88893,88894,88895,88896,88897,88898,88899,88900,88901,88902,88903,88904,88905,88906,88907,88908,88909,88910,88911,88912,88913,88914,88915,88916,88917,88918,88919,88920,88921,88922,88923,88924,88925,88926,88927,88928,88929,88930,88931,88932,88933,88934,88935,88936,88937,88938,88939,88940,88941,88942,88943,88944,88945,88946,88947,88948,88949,89125,89127,89178,89179,89180,89181,89189,89190,89191,89192,89224,89225,89226,89227,89238,89239,89240,89241,89242,89243,89339,89345,89346,89347,89351,89354,89361,89367,89381,89382,89390,89391,89401,89408,89426,89427,89428,89429,89436,89437,89438,89439,89440,89441,89442,89443,89445,89448,89479,89481,89654,89655,89656,89657,89658,89659,89660,89661,89662,89663,89664,89665,89666,89667,89668,89669,89670,89671,89672,89673,89674,89675,89676,89677,89678,89679,89680,89681,89682,89683,89684,89685,89686,89687,89688,89689,89690,89691,89692,89693,89694,89695,89696,89697,89698,89699,89700,89701,89702,89703,89704,89705,89706,89707,89708,89709,89710,89711,89712,89713,89714,89715,89716,89717,89718,89719,89720,89721,89722,89723,89724,89725,89726,89727,89728,89729,89730,89731,89732,89733,89778,89779,89780,89781,89782,89783,89784,89785,89786,89787,89788,89789,89790,89791,89792,89793,89794,89795,89796,89797,89798,89799,89800,89801,89802,89803,89805,89806,89807,89808,89811,89812,89813,89814,89815,89822,89824,89834,89839,89841,89860,89867,89871,89872,89874,89875,89884,89885,89886,89887,89888,89889,89890,89891,89892,89893,89894,89895,89911,89913,89927,89931,89938,89939,89965,89966,89972,89977,90005,90008,90009,90010,90011,90016,90017,90018,90019,90020,90052,90088,90089,90348,90349,90350,90351,98443,98444,98445,151859,151860,150744,150745,150760,150822,150981,150983,150987,150988,150989,150990,151023,151377,151840,162536,162567,152376,152379,152380,152381,161537,161538,161539,161540,162069,162107,162109,162110,162111,162112,162113,163341,162616,162614,162550,162535,162539,162553,162560,162561,162573,162581,162615,162540,162543,162545,162547,162548,162549,162552,162554,162555,162558,162559,162562,162563,162564,162566,162569,162570,162576,162589,162590,162592,162593,162595,162597,162600,162602,162613,162617,162618,162619,162620,162621,162622,162623,162624,162625,162627,162628,162629,162630,162631,162632,162633,162634,162635,163343,163346,163347,163350,163353,163354,163355,163356,163407,163408,163411,163416,163713,163722,163723,164040,164041,164042,164043,164046,164047,164062,164063,165714,166206"
)

func main() {
	db, err := sql.Open("mysql",
		"prcinvestment:lu007bond008#!@tcp(10.123.4.215:3306)/newlaw")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("select id, title, content from ex_news left join ex_news_contents on ex_news.id = ex_news_contents.ex_new_id where ex_news.id = ? and alltype like '%?%'")
	if err != nil {
		panic(err)
	}

	ids := strings.Split(allid, ",")
	for _, oneId := range ids {
		var id int
		var title, content string

		err = stmt.QueryRow(oneId, 13).Scan(&id, &title, &content)
		fmt.Println(err)

		noHtmlContent := html.StripTags(content)
		entities := getEntities(noHtmlContent)
		uniqueSlice(&entities)
		fmt.Printf("%d\t%s\t%s\n", id, title, strings.Join(entities, ","))
	}
	/*
		rows, err := db.Query("select id, title, content from ex_news left join ex_news_contents on ex_news.id = ex_news_contents.ex_new_id where ex_news.id in (?)", allid)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var title, content string
			rows.Scan(&id, &title, &content)

			noHtmlContent := html.StripTags(content)
			entities := getEntities(noHtmlContent)
			uniqueSlice(&entities)
			fmt.Printf("%d\t%s\t%s\n", id, title, strings.Join(entities, ","))
		}
	*/
}

func getEntities(s string) []string {
	re := regexp.MustCompile("《[^》]+》")
	entities := re.FindAllString(s, -1)

	for i, _ := range entities {
		entities[i] = strings.Replace(entities[i], "《", "", -1)
		entities[i] = strings.Replace(entities[i], "》", "", -1)
	}

	return entities
}

func uniqueSlice(slice *[]string) {
	found := make(map[string]bool)
	total := 0
	for i, val := range *slice {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slice)[total] = (*slice)[i]
			total++
		}
	}
	*slice = (*slice)[:total]
}
