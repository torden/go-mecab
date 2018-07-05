package mecab

//#cgo CFLAGS: -I/usr/local/mecab-ko/include/ -I./
//#cgo LDFLAGS: -L/usr/local/mecab-ko/lib/ -Wl,-rpath,/usr/local/mecab-ko/lib/ -lmecab
//#include <stdio.h>
//#include <stdlib.h>
//#include <string.h>
//#include <mecab.h>
import "C"
import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"unsafe"
)

const (
	// NormalNode is Normal node defined in the dictionary.
	NormalNode = C.MECAB_NOR_NODE

	// UnKnownNode is Unknown node not defined in the dictionary.
	UnKnownNode = C.MECAB_UNK_NODE

	// BOSNode is Virtual node representing a beginning of the sentence.
	BOSNode = C.MECAB_BOS_NODE

	// EOSNode is Virtual node representing a end of the sentence.
	EOSNode = C.MECAB_EOS_NODE

	// EONNode is Virtual node representing a end of the N-best enumeration.
	EONNode = C.MECAB_EON_NODE

	// SysDic is This is a system dictionary.
	SysDic = C.MECAB_SYS_DIC

	// UsrDic is This is a user dictionary.
	UsrDic = C.MECAB_USR_DIC

	// UnKnownDic is This is a unknown word dictionary.
	UnKnownDic = C.MECAB_UNK_DIC

	// OneBest is One best result is obtained (default mode)
	OneBest = C.MECAB_ONE_BEST

	// NBest is Set this flag if you want to obtain N best results.
	NBest = C.MECAB_NBEST

	// Partial is Set this flag if you want to enable a partial parsing mode.
	// When this flag is set, the input |sentence| needs to be written in partial parsing format.
	Partial = C.MECAB_PARTIAL

	// MarginalProb is Set this flag if you want to obtain marginal probabilities.
	// Marginal probability is set in MeCab::Node::prob.
	// The parsing speed will get 3-5 times slower than the default mode.
	MarginalProb = C.MECAB_MARGINAL_PROB

	// Alternative is Set this flag if you want to obtain alternative results.
	// Not implemented.
	Alternative = C.MECAB_ALTERNATIVE

	// AllMorphs is When this flag is set, the result linked-list (Node::next/prev) traverses all nodes in the lattice.
	AllMorphs = C.MECAB_ALL_MORPHS

	// AllocateSentence is When this flag is set, tagger internally copies the body of passed sentence into internal buffer.
	AllocateSentence = C.MECAB_ALLOCATE_SENTENCE

	// AnyBoundary is The token boundary is not specified.
	AnyBoundary = C.MECAB_ANY_BOUNDARY

	// The position is a strong token boundary.
	TokenBoundary = C.MECAB_TOKEN_BOUNDARY

	// The position is not a token boundary.
	InsideToken = C.MECAB_INSIDE_TOKEN
)

// Tag v2.0 : https://docs.google.com/spreadsheets/d/1-9blXKjtjeKZqsf4NzHeYJCrr49-nXeRF6D80udfcwY/edit#gid=589544265
var tagList = []string{"NNG", "NNP", "NNB", "NNBC", "NR", "NP", "VV", "VA", "VX", "VCP", "VCN", "MM", "MAG", "MAJ", "IC", "JKS", "JKC", "JKG", "JKO", "JKB", "JKV", "JKQ", "JX", "JC", "EP", "EF", "EC", "ETN", "ETM", "XPN", "XSN", "XSV", "XSA", "XR", "SF", "SE", "SSO", "SSC", "SC", "SY", "SL", "SH", "SN"}

// DicInfo represents the mecab_dictionary_info_t structure
type DicInfo struct {
	FileName string // filename of dictionary On Windows, filename is stored in UTF-8 encoding
	CharSet  string // character set of the dictionary. e.g., "SHIFT-JIS", "UTF-8"
	Size     uint   // How many words are registered in this dictionary.
	Type     int    // dictionary type this value should be MECAB_USR_DIC, MECAB_SYS_DIC, or MECAB_UNK_DIC.
	Lsize    uint   // left attributes size
	Rsize    uint   // right attributes size
	Version  uint8  // version of this dictionary
}

// NLPAnalyzed represents the List of NLP Analyzed Data
type NLPAnalyzed struct {
	Result []NLPAnalyzedItem
}

// NLPAnalyzedItem represents the NLP Analyzed Item Including Value and Tag
type NLPAnalyzedItem struct {
	Value string
	Tag   string
}

// GetValues returns string slice of List of NLP Analyzed Data
func (n *NLPAnalyzed) GetValues() []string {

	var retval []string

	for _, val := range n.Result {
		retval = append(retval, val.Value)
	}

	return retval
}

// MeCab is mecab-ko NLP library methods, All operations on this object
type MeCab struct {
	version     string
	dicPath     string
	isInit      bool
	isInitModel bool
	pmecab      *C.struct_mecab_t
	pmodel      *C.struct_mecab_model_t
	ptagger     *C.struct_mecab_t
	plattice    *C.struct_mecab_lattice_t
	mutx        sync.RWMutex
}

// NewMeCab Creates and returns a MeCab Library methods's pointer.
func NewMeCab(dicpath string) *MeCab {

	if len(strings.TrimSpace(dicpath)) < 1 {
		fmt.Fprintf(os.Stderr, "`dicpath` was empty")
		return nil
	}

	if _, err := os.Stat(dicpath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "`%s' does not exists", dicpath)
		return nil
	}

	obj := &MeCab{}
	obj.version = C.GoString(C.mecab_version())
	obj.dicPath = ""

	obj.pmecab = nil
	obj.pmodel = nil
	obj.ptagger = nil
	obj.plattice = nil

	obj.isInit = false
	obj.isInitModel = false

	obj.dicPath = fmt.Sprintf("-d %s", dicpath)

	return obj
}

// checkInit does check the Initilization, if not inited, returns error
func (m *MeCab) checkInit() error {

	if !m.isInit {
		return errors.New("not initialized(Init()) go-mecab-ko (MeCab-Ko)")
	}

	return nil
}

// checkInit does check the Mecab's Model Initilization, if not inited, returns error
func (m *MeCab) checkInitModel() error {

	if !m.isInitModel {
		return errors.New("not initialized(InitModel()) go-mecab-ko (MeCab-Ko)")
	}

	return nil
}

// Version returns MeCab-Ko's Version
func (m *MeCab) Version() string {
	return m.version
}

// LastError returns last error
func (m *MeCab) LastError() error {

	retval := C.GoString(C.mecab_strerror(nil))
	if len(retval) < 1 {
		return nil
	}

	return errors.New(retval)
}

// Init does Initialization the mecab, if an error occurs, returns error
func (m *MeCab) Init() error {

	if len(m.dicPath) < 1 {
		return fmt.Errorf("failed to Init, after set the dictionary path")
	}

	tmpDicPath := m.dicPath
	if m.checkInit() == nil {
		m.Destroy()
		m.dicPath = tmpDicPath
	}

	dicArgv := C.CString(m.dicPath)
	defer C.free(unsafe.Pointer(dicArgv))

	m.pmecab = C.mecab_new2(dicArgv)
	if m.pmecab == nil {
		return errors.New(C.GoString(C.mecab_strerror(m.pmecab)))
	}

	m.isInit = true

	return nil
}

// InitModel does Initialization the mecab's model, if an error occurs, returns error
func (m *MeCab) InitModel() error {

	if len(m.dicPath) < 1 {
		return fmt.Errorf("failed to Init, after set the dictionary path")
	}

	tmpDicPath := m.dicPath
	if m.checkInitModel() == nil {
		m.Destroy()
		m.dicPath = tmpDicPath
	}

	dicArgv := C.CString(m.dicPath)
	//defer C.free(unsafe.Pointer(dicArgv))

	//mecab_model_t
	m.pmodel = C.mecab_model_new2(dicArgv)

	if m.pmodel == nil {
		return errors.New(C.GoString(C.mecab_strerror(m.pmecab)))
	}

	//mecab_t
	m.ptagger = C.mecab_model_new_tagger(m.pmodel)
	if m.ptagger == nil {
		return errors.New(C.GoString(C.mecab_strerror(m.pmecab)))
	}

	//mecab_lattice_t
	m.plattice = C.mecab_model_new_lattice(m.pmodel)
	if m.plattice == nil {
		return errors.New(C.GoString(C.mecab_strerror(m.pmecab)))
	}

	m.isInitModel = true

	return nil
}

// Destroy closes the MeCab-Ko's underlying file descriptor and delete objects
func (m *MeCab) Destroy() {

	if m.checkInit() != nil {
		C.mecab_destroy(m.pmecab)
	}

	if m.checkInitModel() != nil {
		C.mecab_lattice_destroy(m.plattice)
		C.mecab_model_destroy(m.pmodel)
		C.mecab_destroy(m.ptagger)
	}

	m.plattice = nil
	m.pmodel = nil
	m.pmecab = nil
	m.pmodel = nil

	m.isInit = false
	m.isInitModel = false

	m.version = ""
	m.dicPath = ""
}

// SParseToStr Gets tagged result in string, if an error occurs, returns error
//
// Parse given sentence and return parsed result as string.
// You should not delete the returned string.
// The returned buffer is overwritten when parse method is called again.
func (m *MeCab) SParseToStr(input string) (string, error) {

	if err := m.checkInit(); err != nil {
		return "", err
	}

	cInputString := C.CString(input)
	defer C.free(unsafe.Pointer(cInputString))

	m.mutx.Lock()
	defer m.mutx.Unlock()

	result := C.mecab_sparse_tostr(m.pmecab, cInputString)
	if result == nil {
		return "", errors.New(C.GoString(C.mecab_strerror(m.pmecab)))
	}

	return C.GoString(result), nil
}

// NBestSparseToStr Gets N best results, if an error occurs, returns error
//
// Parse given sentence and obtain N-best results as a string format.
// Currently, N must be 1 <= N <= 512 due to the limitation of the buffer size.
// You should not delete the returned string.
// The returned buffer is overwritten when parse method is called again.
func (m *MeCab) NBestSparseToStr(input string, size int) (string, error) {

	if err := m.checkInit(); err != nil {
		return "", err
	}

	cInputString := C.CString(input)
	defer C.free(unsafe.Pointer(cInputString))

	m.mutx.Lock()
	defer m.mutx.Unlock()

	result := C.mecab_nbest_sparse_tostr(m.pmecab, C.size_t(size), cInputString)
	if result == nil {
		return "", errors.New(C.GoString(C.mecab_strerror(m.pmecab)))
	}

	return C.GoString(result), nil
}

// NBestInit Does Initialize N-best enumeration with a sentence.
// if an error occurs, returns error
func (m *MeCab) NBestInit(input string) error {

	if err := m.checkInit(); err != nil {
		return err
	}

	cInputString := C.CString(input)
	defer C.free(unsafe.Pointer(cInputString))

	m.mutx.Lock()
	defer m.mutx.Unlock()

	ret := C.mecab_nbest_init(m.pmecab, cInputString)
	if ret == C.int(0) {
		return errors.New(C.GoString(C.mecab_strerror(m.pmecab)))
	}

	return nil
}

// GetPartial Returns true if partial parsing mode is on.
// if an error occurs, returns error
func (m *MeCab) GetPartial() (bool, error) {

	if err := m.checkInit(); err != nil {
		return false, err
	}

	retval := C.mecab_get_partial(m.pmecab)
	if retval == C.int(0) {
		return false, nil
	}

	return true, nil
}

// SetPartial Sets partial parsing mode.
// if an error occurs, returns error
func (m *MeCab) SetPartial(partial bool) error {

	if err := m.checkInit(); err != nil {
		return err
	}

	setpartial := C.int(0)
	if partial {
		setpartial = C.int(1)
	}

	C.mecab_set_partial(m.pmecab, setpartial)
	return nil
}

// NBestNextToStr return next-best result
// if an error occurs, returns error
//
// Obtain next-best result. The internal linked list structure is updated.
// You should set MECAB_NBEST reques_type in advance.
func (m *MeCab) NBestNextToStr() (string, error) {

	if err := m.checkInit(); err != nil {
		return "", err
	}

	result := C.mecab_nbest_next_tostr(m.pmecab)
	if result == nil {
		return "", errors.New(C.GoString(C.mecab_strerror(m.pmecab)))
	}

	return C.GoString(result), nil
}

// GetLatticeLevel Returns lattice level.
// if an error occurs, returns error
func (m *MeCab) GetLatticeLevel() (int, error) {

	if err := m.checkInit(); err != nil {
		return int(-1), err
	}

	retval := C.mecab_get_lattice_level(m.pmecab)
	return int(retval), nil
}

// SetLatticeLevel Sets lattice level.
// if an error occurs, returns error
func (m *MeCab) SetLatticeLevel(level int) error {

	if err := m.checkInit(); err != nil {
		return err
	}

	C.mecab_set_lattice_level(m.pmecab, C.int(level))
	return nil
}

// GetAllMorphs Returns true if all morphs output mode is on.
// if an error occurs, returns error
func (m *MeCab) GetAllMorphs() (bool, error) {

	if err := m.checkInit(); err != nil {
		return false, err
	}

	retval := C.mecab_get_all_morphs(m.pmecab)
	if retval == C.int(0) {
		return false, nil
	}

	return true, nil
}

// SetAllMorphs Sets all-morphs output mode.
// if an error occurs, returns error
func (m *MeCab) SetAllMorphs(allmorphs bool) error {

	if err := m.checkInit(); err != nil {
		return err
	}

	setallmorphs := C.int(0)
	if allmorphs {
		setallmorphs = C.int(1)
	}

	C.mecab_set_all_morphs(m.pmecab, setallmorphs)
	return nil
}

// GetTheTA Returns temparature parameter theta.
// if an error occurs, returns error
func (m *MeCab) GetTheTA() (float32, error) {

	if err := m.checkInit(); err != nil {
		return float32(-1), err
	}

	retval := C.mecab_get_theta(m.pmecab)
	return float32(retval), nil
}

// SetTheTA Sets temparature parameter theta.
// if an error occurs, returns error
func (m *MeCab) SetTheTA(theta float32) error {

	if err := m.checkInit(); err != nil {
		return err
	}

	C.mecab_set_theta(m.pmecab, C.float(theta))
	return nil
}

// GetDictionaryInfo Returns Dictionary Information.
// if an error occurs, returns error
func (m *MeCab) GetDictionaryInfo() ([]DicInfo, error) {

	var retval []DicInfo

	if err := m.checkInit(); err != nil {
		return retval, err
	}

	dicinfo := C.mecab_dictionary_info(m.pmecab)
	for {

		var tmpret DicInfo

		tmpret.FileName = C.GoString((*dicinfo).filename)
		tmpret.CharSet = C.GoString((*dicinfo).charset)
		tmpret.Size = uint((dicinfo).size)
		tmpret.Type = int((dicinfo)._type)
		tmpret.Lsize = uint((dicinfo).lsize)
		tmpret.Rsize = uint((dicinfo).rsize)
		tmpret.Version = uint8((dicinfo).version)

		retval = append(retval, tmpret)

		dicinfo = (*dicinfo).next
		if dicinfo == nil {
			break
		}
	}

	return retval, nil
}

// LatticeSetSentence Sets sentence.
// if an error occurs, returns error
// This method does not take the ownership of the object.
func (m *MeCab) LatticeSetSentence(input string) error {

	if err := m.checkInitModel(); err != nil {
		return err
	}

	cInputString := C.CString(input)

	C.mecab_lattice_set_sentence(m.plattice, cInputString)

	return nil

}

// ParseLattice Does Parse lattice object.
// if an error occurs, returns error
func (m *MeCab) ParseLattice() error {

	if err := m.checkInitModel(); err != nil {
		return err
	}

	m.mutx.Lock()
	defer m.mutx.Unlock()

	C.mecab_parse_lattice(m.ptagger, m.plattice)

	return nil
}

// LatticeToSTR Returns string representation of the lattice.
// if an error occurs, returns error
// Returned object is managed by this instance.
// When clear/LatticeSetSentence() method is called, the returned buffer is initialized.
func (m *MeCab) LatticeToSTR() (string, error) {

	if err := m.checkInitModel(); err != nil {
		return "", err
	}

	retval := C.mecab_lattice_tostr(m.plattice)
	return C.GoString(retval), nil
}

// LatticeGetBosNode Returns bos (begin of sentence) node.
// if an error occurs, returns error
func (m *MeCab) LatticeGetBosNode() (*C.struct_mecab_node_t, error) {

	if err := m.checkInitModel(); err != nil {
		return nil, err
	}

	retval := C.mecab_lattice_get_bos_node(m.plattice)
	return retval, nil
}

// LatticeGetEosNode Returns bos (end of sentence) node.
// if an error occurs, returns error
func (m *MeCab) LatticeGetEosNode() (*C.struct_mecab_node_t, error) {

	if err := m.checkInitModel(); err != nil {
		return nil, err
	}

	retval := C.mecab_lattice_get_eos_node(m.plattice)
	return retval, nil
}

// parser does parsing the nlp analyzed data
// if an error occurs, returns error
func (m *MeCab) parser(args ...string) (NLPAnalyzed, error) {

	var err error
	var retval NLPAnalyzed
	var text, feature string

	argc := len(args)
	if argc < 1 {
		return retval, errors.New("required! input-data")
	}

	text = args[0]
	if argc > 1 {
		feature = args[1]
	}

	if !m.isInitModel {
		err = m.InitModel()
		if err != nil {
			return retval, err
		}
	}

	err = m.LatticeSetSentence(text)
	if err != nil {
		return retval, err
	}

	err = m.ParseLattice()
	if err != nil {
		return retval, err

	}

	tmpres, err := m.LatticeToSTR()
	if err != nil {
		return retval, err
	}

	reslist := strings.Split(tmpres, "\n")
	if len(reslist) < 1 {
		return retval, errors.New("result was empty")
	}

	for _, tmpline := range reslist {

		if tmpline == "EOS" || len(tmpline) < 1 {
			break
		}

		tmpitem := strings.Split(tmpline, "\t")
		tmpval := strings.Split(tmpitem[1], ",")

		switch argc {

		case 1:
			retval.Result = append(retval.Result, NLPAnalyzedItem{Value: tmpitem[0], Tag: tmpval[0]})

		case 2:
			if tmpval[0] == "NR" { //NR = rhetoric (수사)
				continue
			}

			if strings.HasPrefix(tmpval[0], feature) {
				retval.Result = append(retval.Result, NLPAnalyzedItem{Value: tmpitem[0], Tag: tmpval[0]})
			}
		}
	}

	return retval, nil
}

// NounsWithTagInfo returns the nouns with tag id from nlp analyzed data
// if an error occurs, returns error
// Obtains all nouns (NNG, NNP, NNB, NNBC, NP, without NR)
func (m *MeCab) NounsWithTagInfo(text string) (NLPAnalyzed, error) {
	return m.parser(text, "N")
}

// Pos returns return the whole nlp analyzed data
// if an error occurs, returns error
func (m *MeCab) Pos(text string) (NLPAnalyzed, error) {
	return m.parser(text)
}

// Nouns returns the only nouns without tag from nlp analyzed data
// if an error occurs, returns error
func (m *MeCab) Nouns(text string) ([]string, error) {
	result, err := m.parser(text, "N")
	return result.GetValues(), err
}

// Morphs returns the whole values of nlp analyzed data
// if an error occurs, returns error
func (m *MeCab) Morphs(text string) ([]string, error) {
	result, err := m.parser(text)
	return result.GetValues(), err
}

// ByTagWithTagInfo returns the nlp analyzed data by tag-id (begins with prefix)
// if an error occurs, returns error
func (m *MeCab) ByTagWithTagInfo(text string, tag string) (NLPAnalyzed, error) {
	return m.parser(text, tag)
}

// ByTag returns the nlp analyzed data by tag-id (begins with prefix) without tag-id
// if an error occurs, returns error
func (m *MeCab) ByTag(text string, tag string) ([]string, error) {

	var retval []string
	var err error

	ok := false
	for _, val := range tagList {
		if val == tag {
			ok = true
			break
		}
	}

	if !ok {
		return retval, fmt.Errorf("not allow tag(=%s)", tag)
	}

	result, err := m.parser(text, tag)
	retval = result.GetValues()
	return retval, err
}
