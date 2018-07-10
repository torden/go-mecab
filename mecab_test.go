package mecab_test

import (
	"encoding/json"
	"testing"

	mecab "github.com/torden/go-mecab"
	"github.com/torden/go-strutil"
)

var assert = strutils.NewAssert()

var defaultDicPath = "/usr/local/mecab-ko/lib/mecab/dic/mecab-ko-dic/"
var gMeCab = mecab.NewMeCab(defaultDicPath)

var str_kr = `mecab-ko는 은전한닢 프로젝트에서 사용하기 위한 MeCab의 fork 프로젝트 입니다.
최소한의 변경으로 한국어의 특성에 맞는 기능을 추가하는 것이 목표입니다.`

var str_jp = `MeCabはオープンソースの形態素解析エンジンで、奈良先端科学技術大学院大学出身、
現GoogleソフトウェアエンジニアでGoogle 日本語入力開発者の一人である工藤拓によって開発されている。
名称は開発者の好物「和布蕪（めかぶ）」から取られた。`

var tagList = []string{"NNG", "NNP", "NNB", "NNBC", "NR", "NP",
	"VV", "VA", "VX", "VCP", "VCN",
	"MM", "MAG", "MAJ",
	"IC",
	"JKS", "JKC", "JKG", "JKO", "JKB", "JKV", "JKQ", "JX", "JC",
	"EP", "EF", "EC", "ETN", "ETM",
	"XPN", "XSN", "XSV", "XSA", "XR",
	"SF", "SE", "SSO", "SSC", "SC", "SY", "SL", "SH", "SN"}

func Test_mecab_NewMeCab(t *testing.T) {

	err1 := mecab.NewMeCab("            ")
	assert.AssertNotNil(t, err1, "failed to testing about empty dicpath")

	err2 := mecab.NewMeCab("/dev/TEST")
	assert.AssertNotNil(t, err2, "failed to testing about wrong dicpath")

	test1 := mecab.NewMeCab(defaultDicPath)
	test1.Destroy()

	test2 := mecab.NewMeCab(defaultDicPath)
	test2.Init()
	test2.Destroy()

	test3 := mecab.NewMeCab(defaultDicPath)
	test3.InitModel()
	test3.Destroy()

	test4 := mecab.NewMeCab(defaultDicPath)
	test4.Init()
	test4.InitModel()
	test4.Destroy()

	test5 := mecab.NewMeCab(defaultDicPath)
	test5.Init()
	test5.Init()
	test5.Init()
}

func Test_mecab_SParseToStr(t *testing.T) {

	var err error

	str := `太郎は次郎が持っている本を花子に渡した。`
	str_ok := "太\tNNG,*,F,태,*,*,*,*\n郎\tSH,*,*,*,*,*,*,*\nは\tSL,*,*,*,*,*,*,*\n次\tXSN,*,F,차,*,*,*,*\n郎\tSH,*,*,*,*,*,*,*\nが\tSL,*,*,*,*,*,*,*\n持\tSH,*,*,*,*,*,*,*\nっている\tSL,*,*,*,*,*,*,*\n本\tXSN,*,T,본,*,*,*,*\nを\tSL,*,*,*,*,*,*,*\n花子\tNNG,*,F,화자,*,*,*,*\nに\tSL,*,*,*,*,*,*,*\n渡\tSH,*,*,*,*,*,*,*\nした\tSL,*,*,*,*,*,*,*\n。\tSY,*,*,*,*,*,*,*\nEOS\n"

	_, err = gMeCab.SParseToStr(str)
	assert.AssertNotNil(t, err, "failed to testing about not inited.")

	gMeCab.Init()
	result, err := gMeCab.SParseToStr(str)
	assert.AssertNil(t, err, "failed to NLP analyze")
	assert.AssertEquals(t, str_ok, result, "Return Value mismatch.\nExpected: %v\nActual: %v", str_ok, result)
}

func Test_mecab_NBestSparseToStr(t *testing.T) {

	str := `太郎は次郎が持っている本を花子に渡した。`
	str_ok := "太\tNNG,*,F,태,*,*,*,*\n郎\tSH,*,*,*,*,*,*,*\nは\tSL,*,*,*,*,*,*,*\n次\tXSN,*,F,차,*,*,*,*\n郎\tSH,*,*,*,*,*,*,*\nが\tSL,*,*,*,*,*,*,*\n持\tSH,*,*,*,*,*,*,*\nっている\tSL,*,*,*,*,*,*,*\n本\tXSN,*,T,본,*,*,*,*\nを\tSL,*,*,*,*,*,*,*\n花子\tNNG,*,F,화자,*,*,*,*\nに\tSL,*,*,*,*,*,*,*\n渡\tSH,*,*,*,*,*,*,*\nした\tSL,*,*,*,*,*,*,*\n。\tSY,*,*,*,*,*,*,*\nEOS\n太\tNNG,*,F,태,*,*,*,*\n郎\tSH,*,*,*,*,*,*,*\nは\tSL,*,*,*,*,*,*,*\n次\tXSN,*,F,차,*,*,*,*\n郎\tSH,*,*,*,*,*,*,*\nが\tSL,*,*,*,*,*,*,*\n持\tSH,*,*,*,*,*,*,*\nっている\tSL,*,*,*,*,*,*,*\n本\tNNG,*,T,본,*,*,*,*\nを\tSL,*,*,*,*,*,*,*\n花子\tNNG,*,F,화자,*,*,*,*\nに\tSL,*,*,*,*,*,*,*\n渡\tSH,*,*,*,*,*,*,*\nした\tSL,*,*,*,*,*,*,*\n。\tSY,*,*,*,*,*,*,*\nEOS\n太\tNNG,*,F,태,*,*,*,*\n郎\tSH,*,*,*,*,*,*,*\nは\tSL,*,*,*,*,*,*,*\n次\tXSN,*,F,차,*,*,*,*\n郎\tSH,*,*,*,*,*,*,*\nが\tSL,*,*,*,*,*,*,*\n持\tSH,*,*,*,*,*,*,*\nっている\tSL,*,*,*,*,*,*,*\n本\tXPN,*,T,본,*,*,*,*\nを\tSL,*,*,*,*,*,*,*\n花子\tNNG,*,F,화자,*,*,*,*\nに\tSL,*,*,*,*,*,*,*\n渡\tSH,*,*,*,*,*,*,*\nした\tSL,*,*,*,*,*,*,*\n。\tSY,*,*,*,*,*,*,*\nEOS\n"

	gMeCab.Init()
	result, err := gMeCab.NBestSparseToStr(str, 3)
	assert.AssertNil(t, err, "failed to NLP analyze")
	assert.AssertEquals(t, str_ok, result, "Return Value mismatch.\nExpected: %v\nActual: %v", str_ok, result)
}

func Test_mecab_Partial(t *testing.T) {

	var retval bool
	var err error
	gMeCab.Init()

	retval, err = gMeCab.GetPartial()
	assert.AssertNil(t, err, "failed to testing about GetPartial")
	assert.AssertFalse(t, retval, "failed to testing about GetPartial(=%t)", retval)

	err = gMeCab.SetPartial(true)
	assert.AssertNil(t, err, "failed to testing about SetPartial")

	retval, err = gMeCab.GetPartial()
	assert.AssertNil(t, err, "failed to testing about GetPartial")
	assert.AssertTrue(t, retval, "failed to testing about GetPartial(=%t)", retval)

	//roll-back
	err = gMeCab.SetPartial(false)
	assert.AssertNil(t, err, "failed to testing about SetPartial")
}

func Test_mecab_AllMorphs(t *testing.T) {

	var retval bool
	var err error
	gMeCab.Init()

	retval, err = gMeCab.GetAllMorphs()
	assert.AssertNil(t, err, "failed to testing about GetAllMorphs")
	assert.AssertFalse(t, retval, "failed to testing about GetAllMorphs(=%t)", retval)

	err = gMeCab.SetAllMorphs(true)
	assert.AssertNil(t, err, "failed to testing about SetAllMorphs")

	retval, err = gMeCab.GetAllMorphs()
	assert.AssertNil(t, err, "failed to testing about GetAllMorphs")
	assert.AssertTrue(t, retval, "failed to testing about GetAllMorphs(=%t)", retval)

	//roll-back
	err = gMeCab.SetAllMorphs(false)
	assert.AssertNil(t, err, "failed to testing about SetAllMorphs")
}

func Test_mecab_TheTA(t *testing.T) {

	var retval float32
	var err error
	gMeCab.Init()

	retval, err = gMeCab.GetTheTA()
	assert.AssertNil(t, err, "failed to testing about GetTheTA")
	assert.AssertEquals(t, retval, float32(0.75), "Return Value mismatch.\nExpected: %v\nActual: %v", retval, float32(0.75))

	err = gMeCab.SetTheTA(0.95)
	assert.AssertNil(t, err, "failed to testing about SetTheTA")

	retval, err = gMeCab.GetTheTA()
	assert.AssertNil(t, err, "failed to testing about GetTheTA")
	assert.AssertEquals(t, retval, float32(0.95), "Return Value mismatch.\nExpected: %v\nActual: %v", retval, float32(0.95))

	//roll-back
	err = gMeCab.SetTheTA(0.75)
	assert.AssertNil(t, err, "failed to testing about SetTheTA")
}

func Test_mecab_GetDictionaryInfo(t *testing.T) {

	test1 := mecab.NewMeCab("/usr/local/mecab-ko/lib/mecab/dic/mecab-ko-dic/")
	test1.GetDictionaryInfo()

	retval, err := gMeCab.GetDictionaryInfo()
	assert.AssertNil(t, err, "failed to testing about GetDictionaryInfo")

	if len(retval) > 0 {

		assert.AssertGreaterThan(t, len((retval[0]).FileName), 1, "failed to testing about GetDictionaryInfo.FileName")
		assert.AssertGreaterThan(t, len((retval[0]).CharSet), 1, "failed to testing about GetDictionaryInfo.CharSet")
		assert.AssertGreaterThan(t, (retval[0]).Size, 1, "failed to testing about GetDictionaryInfo.Size")
		assert.AssertGreaterThanEqualTo(t, (retval[0]).Type, 0, "failed to testing about GetDictionaryInfo.Type")
		assert.AssertGreaterThanEqualTo(t, (retval[0]).Lsize, 1, "failed to testing about GetDictionaryInfo.Lsize")
		assert.AssertGreaterThanEqualTo(t, (retval[0]).Rsize, 1, "failed to testing about GetDictionaryInfo.Rsize")
		assert.AssertGreaterThanEqualTo(t, (retval[0]).Version, 100, "failed to testing about GetDictionaryInfo.Version")
	}
}

func Test_mecab_NounsWithTagInfo_KR(t *testing.T) {

	ok_nouns_kr_with_tag := map[string]string{"은전한닢": "NNG+NR+NNG",
		"프로젝트": "NNG",
		"사용":   "NNG",
		"최소한":  "NNG",
		"변경":   "NNG",
		"한국어":  "NNG",
		"특성":   "NNG",
		"기능":   "NNG",
		"추가":   "NNG",
		"것":    "NNB",
		"목표":   "NNG",
	}

	result, err := gMeCab.NounsWithTagInfo(str_kr)
	assert.AssertNil(t, err, "failed to NLP Analyze, err=%v", err)

	for _, item := range result.Result {
		tag, ok := ok_nouns_kr_with_tag[item.Value]
		assert.AssertTrue(t, ok, "invalid tag(Value=%s,Tag=`%s`)", item.Value, item.Tag)
		assert.AssertEquals(t, tag, item.Tag, "Return Value mismatch.\nExpected: %v\nActual: %v", tag, item.Tag)
	}
}

func Test_mecab_NounsWithTagInfo_JP(t *testing.T) {

	ok_nouns_jp_with_tag := map[string]string{"形態": "NNG",
		"素":  "NNG",
		"解析": "NNG",
		"奈良": "NNG",
		"先端": "NNG",
		"科":  "NNG",
		"技術": "NNG",
		"大":  "NNG",
		"院":  "NNG",
		"出身": "NNG",
		"現":  "NNG",
		"日本": "NNG",
		"入力": "NNG",
		"工":  "NNG",
		"藤":  "NNG",
		"名":  "NNG",
		"好物": "NNG",
		"和":  "NNG",
		"取":  "NNG",
	}

	result, err := gMeCab.NounsWithTagInfo(str_jp)
	assert.AssertNil(t, err, "failed to NLP Analyze, err=%v", err)

	for _, item := range result.Result {
		tag, ok := ok_nouns_jp_with_tag[item.Value]
		assert.AssertTrue(t, ok, "invalid tag(Value=%s,Tag=`%s`)", item.Value, item.Tag)
		assert.AssertEquals(t, tag, item.Tag, "Return Value mismatch.\nExpected: %v\nActual: %v", tag, item.Tag)
	}
}

func Test_mecab_Pos_KR(t *testing.T) {

	ok_full_kr_with_tag_json := `{"Result":[{"Value":"mecab","Tag":"SL"},{"Value":"-","Tag":"SY"},{"Value":"ko","Tag":"SL"},{"Value":"는","Tag":"JX"},{"Value":"은전한닢","Tag":"NNG+NR+NNG"},{"Value":"프로젝트","Tag":"NNG"},{"Value":"에서","Tag":"JKB"},{"Value":"사용","Tag":"NNG"},{"Value":"하","Tag":"XSV"},{"Value":"기","Tag":"ETN"},{"Value":"위한","Tag":"VV+ETM"},{"Value":"MeCab","Tag":"SL"},{"Value":"의","Tag":"JKG"},{"Value":"fork","Tag":"SL"},{"Value":"프로젝트","Tag":"NNG"},{"Value":"입니다","Tag":"VCP+EF"},{"Value":".","Tag":"SF"},{"Value":"최소한","Tag":"NNG"},{"Value":"의","Tag":"JKG"},{"Value":"변경","Tag":"NNG"},{"Value":"으로","Tag":"JKB"},{"Value":"한국어","Tag":"NNG"},{"Value":"의","Tag":"JKG"},{"Value":"특성","Tag":"NNG"},{"Value":"에","Tag":"JKB"},{"Value":"맞","Tag":"VV"},{"Value":"는","Tag":"ETM"},{"Value":"기능","Tag":"NNG"},{"Value":"을","Tag":"JKO"},{"Value":"추가","Tag":"NNG"},{"Value":"하","Tag":"XSV"},{"Value":"는","Tag":"ETM"},{"Value":"것","Tag":"NNB"},{"Value":"이","Tag":"JKS"},{"Value":"목표","Tag":"NNG"},{"Value":"입니다","Tag":"VCP+EF"},{"Value":".","Tag":"SF"}]}`

	result, err := gMeCab.Pos(str_kr)
	assert.AssertNil(t, err, "failed to NLP Analyze, err=%v", err)

	tmpbuf, err := json.Marshal(result)
	assert.AssertNil(t, err, "failed to NLP Analyze, err=%v", err)

	result_str := string(tmpbuf)

	assert.AssertEquals(t, result_str, ok_full_kr_with_tag_json, "Return Value mismatch.\nExpected: %v\nActual: %v", result_str, ok_full_kr_with_tag_json)
}

func Test_mecab_Pos_JP(t *testing.T) {

	ok_full_jp_with_tag_json := `{"Result":[{"Value":"MeCab","Tag":"SL"},{"Value":"は","Tag":"SL"},{"Value":"オープンソース","Tag":"SL"},{"Value":"の","Tag":"SL"},{"Value":"形態","Tag":"NNG"},{"Value":"素","Tag":"NNG"},{"Value":"解析","Tag":"NNG"},{"Value":"エンジン","Tag":"SL"},{"Value":"で","Tag":"SL"},{"Value":"、","Tag":"SY"},{"Value":"奈良","Tag":"NNG"},{"Value":"先端","Tag":"NNG"},{"Value":"科","Tag":"NNG"},{"Value":"学","Tag":"SH"},{"Value":"技術","Tag":"NNG"},{"Value":"大","Tag":"NNG"},{"Value":"学","Tag":"SH"},{"Value":"院","Tag":"NNG"},{"Value":"大","Tag":"NNG"},{"Value":"学","Tag":"SH"},{"Value":"出身","Tag":"NNG"},{"Value":"、","Tag":"SY"},{"Value":"現","Tag":"NNG"},{"Value":"Google","Tag":"SL"},{"Value":"ソフトウェアエンジニア","Tag":"SL"},{"Value":"で","Tag":"SL"},{"Value":"Google","Tag":"SL"},{"Value":"日本","Tag":"NNG"},{"Value":"語","Tag":"XSN"},{"Value":"入力","Tag":"NNG"},{"Value":"開","Tag":"SH"},{"Value":"発","Tag":"SH"},{"Value":"者","Tag":"XSN"},{"Value":"の","Tag":"SL"},{"Value":"一人","Tag":"SH"},{"Value":"である","Tag":"SL"},{"Value":"工","Tag":"NNG"},{"Value":"藤","Tag":"NNG"},{"Value":"拓","Tag":"SH"},{"Value":"によって","Tag":"SL"},{"Value":"開","Tag":"SH"},{"Value":"発","Tag":"SH"},{"Value":"されている","Tag":"SL"},{"Value":"。","Tag":"SY"},{"Value":"名","Tag":"NNG"},{"Value":"称","Tag":"SH"},{"Value":"は","Tag":"SL"},{"Value":"開","Tag":"SH"},{"Value":"発","Tag":"SH"},{"Value":"者","Tag":"XSN"},{"Value":"の","Tag":"SL"},{"Value":"好物","Tag":"NNG"},{"Value":"「","Tag":"SSO"},{"Value":"和","Tag":"NNG"},{"Value":"布","Tag":"SH"},{"Value":"蕪","Tag":"SH"},{"Value":"（","Tag":"SY"},{"Value":"めかぶ","Tag":"SL"},{"Value":"）」","Tag":"SY"},{"Value":"から","Tag":"SL"},{"Value":"取","Tag":"NNG"},{"Value":"られた","Tag":"SL"},{"Value":"。","Tag":"SY"}]}`

	result, err := gMeCab.Pos(str_jp)
	assert.AssertNil(t, err, "failed to NLP Analyze, err=%v", err)

	tmpbuf, err := json.Marshal(result)
	assert.AssertNil(t, err, "failed to NLP Analyze, err=%v", err)

	result_str := string(tmpbuf)

	assert.AssertEquals(t, result_str, ok_full_jp_with_tag_json, "Return Value mismatch.\nExpected: %v\nActual: %v", result_str, ok_full_jp_with_tag_json)
}

func Test_mecab_Nouns_KR(t *testing.T) {

	ok_nouns_kr := map[string]bool{
		"은전한닢": true,
		"프로젝트": true,
		"사용":   true,
		"최소한":  true,
		"변경":   true,
		"한국어":  true,
		"특성":   true,
		"기능":   true,
		"추가":   true,
		"것":    true,
		"목표":   true,
	}

	result, err := gMeCab.Nouns(str_kr)
	assert.AssertNil(t, err, "failed to NLP Analyze, err=%v", err)

	for _, item := range result {
		_, exists := ok_nouns_kr[item]
		assert.AssertTrue(t, exists, "invalid noun(=%s)", item)
	}
}

func Test_mecab_Nouns_JP(t *testing.T) {

	ok_nouns_jp := map[string]bool{
		"形態": true,
		"素":  true,
		"解析": true,
		"奈良": true,
		"先端": true,
		"科":  true,
		"技術": true,
		"大":  true,
		"院":  true,
		"出身": true,
		"現":  true,
		"日本": true,
		"入力": true,
		"工":  true,
		"藤":  true,
		"名":  true,
		"好物": true,
		"和":  true,
		"取":  true,
	}

	result, err := gMeCab.Nouns(str_jp)
	assert.AssertNil(t, err, "failed to NLP Analyze, err=%v", err)

	for _, item := range result {
		_, exists := ok_nouns_jp[item]
		assert.AssertTrue(t, exists, "invalid noun(=%s)", item)
	}
}

func Test_mecab_Morphs_KR(t *testing.T) {

	ok_full_kr_without_tag := map[string]bool{
		"mecab": true,
		"-":     true,
		"ko":    true,
		"는":     true,
		"은전한닢":  true,
		"프로젝트":  true,
		"에서":    true,
		"사용":    true,
		"하":     true,
		"기":     true,
		"위한":    true,
		"MeCab": true,
		"fork":  true,
		"입니다":   true,
		".":     true,
		"최소한":   true,
		"의":     true,
		"변경":    true,
		"으로":    true,
		"한국어":   true,
		"특성":    true,
		"에":     true,
		"맞":     true,
		"기능":    true,
		"을":     true,
		"추가":    true,
		"것":     true,
		"이":     true,
		"목표":    true,
	}

	result, err := gMeCab.Morphs(str_kr)
	assert.AssertNil(t, err, "failed to NLP Analyze, err=%v", err)

	for _, item := range result {
		_, exists := ok_full_kr_without_tag[item]
		assert.AssertTrue(t, exists, "invalid noun(=%s)", item)
	}
}

func Test_mecab_Morphs_JP(t *testing.T) {

	ok_full_jp_without_tag := map[string]bool{
		"MeCab": true,
		"は":     true,
		"オープンソース": true,
		"の":      true,
		"形態":     true,
		"素":      true,
		"解析":     true,
		"エンジン":   true,
		"で":      true,
		"、":      true,
		"奈良":     true,
		"先端":     true,
		"科":      true,
		"学":      true,
		"技術":     true,
		"大":      true,
		"院":      true,
		"出身":     true,
		"現":      true,
		"Google": true,
		"ソフトウェアエンジニア": true,
		"日本":    true,
		"語":     true,
		"入力":    true,
		"開":     true,
		"発":     true,
		"者":     true,
		"一人":    true,
		"である":   true,
		"工":     true,
		"藤":     true,
		"拓":     true,
		"によって":  true,
		"されている": true,
		"。":     true,
		"名":     true,
		"称":     true,
		"好物":    true,
		"「":     true,
		"和":     true,
		"布":     true,
		"蕪":     true,
		"（":     true,
		"めかぶ":   true,
		"）」":    true,
		"から":    true,
		"取":     true,
		"られた":   true,
	}

	result, err := gMeCab.Morphs(str_jp)
	assert.AssertNil(t, err, "failed to NLP Analyze, err=%v", err)

	for _, item := range result {
		_, exists := ok_full_jp_without_tag[item]
		assert.AssertTrue(t, exists, "invalid noun(=%s)", item)
	}
}

func Test_mecab_ByTagWithTagInfo_KR(t *testing.T) {

	var ok_kr map[string]string

	ok_kr_NNG := map[string]string{"은전한닢": "NNG+NR+NNG", "프로젝트": "NNG", "사용": "NNG", "최소한": "NNG", "변경": "NNG", "한국어": "NNG", "특성": "NNG", "기능": "NNG", "추가": "NNG", "목표": "NNG"}
	ok_kr_NNB := map[string]string{"것": "NNB"}
	ok_kr_VV := map[string]string{"위한": "VV+ETM", "맞": "VV"}
	ok_kr_VCP := map[string]string{"입니다": "VCP+EF"}
	ok_kr_JKS := map[string]string{"이": "JKS"}
	ok_kr_JKG := map[string]string{"의": "JKG"}
	ok_kr_JKO := map[string]string{"을": "JKO"}
	ok_kr_JKB := map[string]string{"에서": "JKB", "으로": "JKB", "에": "JKB"}
	ok_kr_JX := map[string]string{"는": "JX"}
	ok_kr_ETN := map[string]string{"기": "ETN"}
	ok_kr_ETM := map[string]string{"는": "ETM"}
	ok_kr_XSV := map[string]string{"하": "XSV"}
	ok_kr_SF := map[string]string{".": "SF"}
	ok_kr_SY := map[string]string{"-": "SY"}
	ok_kr_SL := map[string]string{"mecab": "SL", "ko": "SL", "MeCab": "SL", "fork": "SL"}

	for _, tag := range tagList {

		result, err := gMeCab.ByTagWithTagInfo(str_kr, tag)
		assert.AssertNil(t, err, "failed to NLP Analyze, err=%v", err)

		switch tag {

		case "NNP", "NNBC", "NR", "NP", "VA", "VX", "VCN", "MM", "MAG", "MAJ", "IC", "JKC", "JKV", "JKQ", "JC", "EP", "EF", "EC", "XPN", "XSN", "XSA", "XR", "SE", "SSO", "SSC", "SC", "SH", "SN":
			assert.AssertEquals(t, len(result.Result), 0, "invalid result of NLP Analyzed Data")
			continue

		case "NNG":

			ok_kr = ok_kr_NNG
		case "NNB":
			ok_kr = ok_kr_NNB
		case "VV":
			ok_kr = ok_kr_VV
		case "VCP":
			ok_kr = ok_kr_VCP
		case "JKS":
			ok_kr = ok_kr_JKS
		case "JKG":
			ok_kr = ok_kr_JKG
		case "JKO":
			ok_kr = ok_kr_JKO
		case "JKB":
			ok_kr = ok_kr_JKB
		case "JX":
			ok_kr = ok_kr_JX
		case "ETN":
			ok_kr = ok_kr_ETN
		case "ETM":
			ok_kr = ok_kr_ETM
		case "XSV":
			ok_kr = ok_kr_XSV
		case "SF":
			ok_kr = ok_kr_SF
		case "SY":
			ok_kr = ok_kr_SY
		case "SL":
			ok_kr = ok_kr_SL

		}

		for _, item := range result.GetValues() {
			_, exists := ok_kr[item]
			assert.AssertTrue(t, exists, "invalid noun(Value=%s, Tag=%s)", item, tag)
		}
	}
}

func Test_mecab_ByTag_JP(t *testing.T) {

	var ok_jp map[string]string

	ok_jp_NNG := map[string]string{
		"形態": "NNG",
		"素":  "NNG",
		"解析": "NNG",
		"奈良": "NNG",
		"先端": "NNG",
		"科":  "NNG",
		"技術": "NNG",
		"大":  "NNG",
		"院":  "NNG",
		"出身": "NNG",
		"現":  "NNG",
		"日本": "NNG",
		"入力": "NNG",
		"工":  "NNG",
		"藤":  "NNG",
		"名":  "NNG",
		"好物": "NNG",
		"和":  "NNG",
		"取":  "NNG",
	}

	ok_jp_XSN := map[string]string{"語": "XSN", "者": "XSN"}
	ok_jp_SSO := map[string]string{"「": "SSO"}
	ok_jp_SY := map[string]string{"、": "SY", "。": "SY", "（": "SY", "）」": "SY"}
	ok_jp_SL := map[string]string{
		"MeCab": "SL",
		"は":     "SL",
		"オープンソース": "SL",
		"の":      "SL",
		"エンジン":   "SL",
		"で":      "SL",
		"Google": "SL",
		"ソフトウェアエンジニア": "SL",
		"である":         "SL",
		"によって":        "SL",
		"されている":       "SL",
		"めかぶ":         "SL",
		"から":          "SL",
		"られた":         "SL",
	}

	ok_jp_SH := map[string]string{
		"学":  "SH",
		"開":  "SH",
		"発":  "SH",
		"一人": "SH",
		"拓":  "SH",
		"称":  "SH",
		"布":  "SH",
		"蕪":  "SH",
	}

	for _, tag := range tagList {
		result, err := gMeCab.ByTagWithTagInfo(str_jp, tag)
		assert.AssertNil(t, err, "failed to NLP Analyze, err=%v", err)

		switch tag {

		case "NNP", "NNB", "NNBC", "NR", "NP", "VV", "VA", "VX", "VCP", "VCN", "MM", "MAG", "MAJ", "IC", "JKS", "JKC", "JKG", "JKO", "JKB", "JKV", "JKQ", "JX", "JC", "EP", "EF", "EC", "ETN", "ETM", "XPN", "XSV", "XSA", "XR", "SF", "SE", "SSC", "SC", "SN":
			assert.AssertEquals(t, len(result.Result), 0, "invalid result of NLP Analyzed Data")
			continue

		case "NNG":
			ok_jp = ok_jp_NNG
		case "XSN":
			ok_jp = ok_jp_XSN
		case "SSO":
			ok_jp = ok_jp_SSO
		case "SY":
			ok_jp = ok_jp_SY
		case "SL":
			ok_jp = ok_jp_SL
		case "SH":
			ok_jp = ok_jp_SH
		}

		for _, item := range result.GetValues() {
			_, exists := ok_jp[item]
			assert.AssertTrue(t, exists, "invalid noun(Value=%s, Tag=%s)", item, tag)
		}
	}
}
