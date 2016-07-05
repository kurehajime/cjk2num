// Package kan2num is 漢数字を数字に変えるやつ
package kan2num

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Knetic/govaluate"
)

//BreakSymbol : 10000の倍数の単位
var BreakSymbol = map[string]int64{"万": 10000,
	"億": 10000 * 10000,
	"兆": 10000 * 10000 * 10000,
	"京": 10000 * 10000 * 10000 * 10000,
}

//NonBreakSymbol :10000の倍数以外の単位
var NonBreakSymbol = map[string]int64{"十": 10, "拾": 10,
	"百": 100,
	"千": 1000,
}

//Numbers :数字と互換性のある文字列
var Numbers = map[string]int64{"零": 0, "〇": 0, "○": 0,
	"一": 1, "壱": 1,
	"二": 2, "弐": 2,
	"三": 3, "参": 3,
	"四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9,
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"０": 0, "１": 1, "２": 2, "３": 3, "４": 4, "５": 5, "６": 6, "７": 7, "８": 8, "９": 9,
}

//Kan2num :漢数字を数字に変換する。
func Kan2num(_word string) (float64, error) {
	var word = _word
	word = clean(word)
	word = makeFormula(word)
	word = transNum(word)

	//eval
	expression, err := govaluate.NewEvaluableExpression(word)
	parameters := make(map[string]interface{}, 0)
	result, err := expression.Evaluate(parameters)
	if err != nil {
		return 0, err
	}

	defer func() {
		err2 := recover()
		if err2 != nil {
			err = fmt.Errorf("%s", err2)
		}
	}()

	return result.(float64), err
}

//定義済みの文字以外は削除
func clean(_word string) string {
	var key string
	var targets = ""
	var word = _word
	var re *regexp.Regexp

	for key = range BreakSymbol {
		targets += key
	}
	for key = range NonBreakSymbol {
		targets += key
	}
	for key = range Numbers {
		targets += key
	}

	re = regexp.MustCompile("[^" + targets + "]")
	word = re.ReplaceAllString(word, "")
	return word
}

func transNum(_word string) string {
	var key string
	var value int64
	var word = _word
	var re *regexp.Regexp
	for key, value = range Numbers {
		re = regexp.MustCompile(key)
		word = re.ReplaceAllString(word, strconv.FormatInt(value, 10))
	}

	for key, value = range BreakSymbol {
		re = regexp.MustCompile(key)
		word = re.ReplaceAllString(word, strconv.FormatInt(value, 10))
	}

	for key, value = range NonBreakSymbol {
		re = regexp.MustCompile(key)
		word = re.ReplaceAllString(word, strconv.FormatInt(value, 10))
	}

	return word
}

func makeFormula(_word string) string {
	var key string
	var targets = ""
	var word = _word
	var re *regexp.Regexp

	//BreakSymbol
	targets = ""
	for key = range BreakSymbol {
		targets += key
	}
	re = regexp.MustCompile("([" + targets + "])")
	word = re.ReplaceAllString(word, ")*$1)+((")

	//NonBreakSymbol
	targets = ""
	for key = range NonBreakSymbol {
		targets += key
	}
	re = regexp.MustCompile("([" + targets + "])")
	word = re.ReplaceAllString(word, "*$1+")

	word = "((" + word + "))"

	//replace *+ -> +
	re = regexp.MustCompile("\\+\\*")
	word = re.ReplaceAllString(word, "+")

	//replace +) -> +0)
	re = regexp.MustCompile("\\+\\)")
	word = re.ReplaceAllString(word, "+0)")

	//replace (* -> (1*
	re = regexp.MustCompile("\\(\\*")
	word = re.ReplaceAllString(word, "(1*")

	//replace +(()) -> +((0))
	re = regexp.MustCompile("\\(\\)")
	word = re.ReplaceAllString(word, "(0)")

	return word
}
