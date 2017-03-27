package swan

import (
	"regexp"

	"github.com/wangbin/jiebago"
)

var (
	chineseMatch = regexp.MustCompile(`[\p{Han}]+`)
)

type stopWordChinese struct {
	seg jiebago.Segmenter
}

func NewStopWordChinese(dict string) (StopWord, error) {
	st := new(stopWordChinese)
	err := st.seg.LoadDictionary(dict)
	if err != nil {
		return nil, err
	}
	return st, nil
}

func (st *stopWordChinese) Match(text string) bool {
	return chineseMatch.MatchString(text)
}

func (st *stopWordChinese) Split(text string) []string {
	val := make([]string, 0)
	ch := st.seg.Cut(text, false)
	for v := range ch {
		val = append(val, v)
	}
	return val
}
