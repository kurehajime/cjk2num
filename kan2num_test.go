package cjk2num

import (
	"testing"
)

func TestConvert(t *testing.T) {
	check(t, "一億三千万二百十五", 130000215.0)
	check(t, "一億", 100000000.0)
	check(t, "千拾", 1010.0)
	check(t, "百", 100.0)
	check(t, "二百三十一兆五十五億二千万千五百一", 231005520001501.0)
	check(t, "二百三十一兆五十五億二千万千五〇一", 231005520001501.0)
	check(t, "二十三人", 23.0)
	check(t, "７十1", 71.0)

}
func check(t *testing.T, input string, ans float64) {
	res, err := Convert(input)
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if res != ans {
		t.Errorf("%s =\n%f\n, want \n%f", input, res, ans)
		return
	}
}
