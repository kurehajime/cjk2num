package cjk2num

import (
	"fmt"
	"testing"
)

// test Convert
func TestConvert(t *testing.T) {
	check(t, "一億三千万二百十五", 130000215.0)
	check(t, "一億", 100000000.0)
	check(t, "千拾", 1010.0)
	check(t, "百", 100.0)
	check(t, "二百三十一兆五十五億二千万千五百一", 231005520001501.0)
	check(t, "二百三十一兆五十五億二千万千五〇一", 231005520001501.0)
	check(t, "二十三人", 23.0)
	check(t, "７十1", 71.0)

	// expect error
	if _, err := Convert("一億万"); err == nil {
		t.Errorf("passed invalid format \n")
	}
}

//Check
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

// Example 1
func Example_case1() {
	res, _ := Convert("一九八四")
	fmt.Printf("%.0f", res)
	//Output:1984
}

// Example 2
func Example_case2() {
	res, _ := Convert("一億二千三百四十五万六千七百八十九")
	fmt.Printf("%.0f", res)
	//Output:123456789
}

// Example 3
func Example_case3() {
	res, _ := Convert("壱萬弐千参百")
	fmt.Printf("%.0f", res)
	//Output:12300
}

// Example 4
func Example_case4() {
	res, _ := Convert("一〇九")
	fmt.Printf("%.0f", res)
	//Output:109
}

// Example 5
func Example_case5() {
	res, _ := Convert("壹億貳仟叁佰肆拾伍萬陸仟柒佰捌拾玖")
	fmt.Printf("%.0f", res)
	//Output:123456789
}

// Example 6
func Example_case6() {
	res, _ := Convert("오만육천칠백팔십구")
	fmt.Printf("%.0f", res)
	//Output:56789
}
