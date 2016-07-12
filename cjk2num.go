// Package cjk2num : Convert /漢数字|中文数字|한자 숫자/  to number
package cjk2num

import (
	"fmt"
	"strings"
)

//Symbol 数字や単位など有効なすべての文字------------------
type Symbol interface {
	//Key 対象となる文字
	Key() string
	//与えられたパラメータを元に数字や単位を反映させた結果を返す
	Calc(stage1, stage2, stage3 int64) (int64, int64, int64, error)
}

//BreakSymbol ex:兆億万...------------------
type BreakSymbol struct {
	key   string
	value int64
}

//Key :get key
func (sym BreakSymbol) Key() string { return sym.key }

//Calc 直左のBreakSymbolから現在位置までの数字に対し係数をかける
func (sym BreakSymbol) Calc(stage1, stage2, stage3 int64) (int64, int64, int64, error) {
	if (stage2 + stage1) == 0 {
		return 0, 0, 0, fmt.Errorf("Error:Invalid value.[ ???%s]", sym.key)
	}
	return 0, 0, stage3 + (stage1+stage2)*sym.value, nil
}

//NonBreakSymbol ex:千百十...------------------
type NonBreakSymbol struct {
	key   string
	value int64
}

//Key :get key
func (sym NonBreakSymbol) Key() string { return sym.key }

//Calc 直左のBreakSymbolから現在位置までの数字に対し係数をかける
func (sym NonBreakSymbol) Calc(stage1, stage2, stage3 int64) (int64, int64, int64, error) {
	if stage1 == 0 {
		return 0, stage2 + 1*sym.value, stage3, nil
	}
	return 0, stage2 + (stage1 * sym.value), stage3, nil
}

//AllBreakSymbol ex:ダース...------------------
type AllBreakSymbol struct {
	key   string
	value int64
}

//Key :get key
func (sym AllBreakSymbol) Key() string { return sym.key }

//Calc 左端から現在位置までの数字に対し係数をかける
func (sym AllBreakSymbol) Calc(stage1, stage2, stage3 int64) (int64, int64, int64, error) {
	return 0, 0, (stage3 + stage2 + stage1) * sym.value, nil
}

//NumberSymbol ex:一二三壱弐参...------------------
type NumberSymbol struct {
	key   string
	value int64
}

//Key :get key
func (sym NumberSymbol) Key() string { return sym.key }

//Calc 直左のNumberSymbolを10倍し係数を足す
func (sym NumberSymbol) Calc(stage1, stage2, stage3 int64) (int64, int64, int64, error) {
	return (stage1 * 10) + sym.value, stage2, stage3, nil
}

// Convert /漢数字|中文数字|한자 숫자/  to number------------------
func Convert(word string) (result int64, err error) {
	return ConvertBy(word, presetSymbols)
}

// ConvertBy :オリジナルの桁定義を指定して変換 ------------------
func ConvertBy(word string, symbols []Symbol) (result int64, err error) {
	runes := []rune(word)
	var stage1, stage2, stage3 int64
	defer func() { //オーバーフローでコケるかも
		r := recover()
		if r != nil {
			result, err = 0, fmt.Errorf("%v", r)
		}
	}()
L:
	for len(runes) > 0 {
		for i := range symbols {
			if strings.Index(string(runes), symbols[i].Key()) == 0 {
				runes = runes[len([]rune(symbols[i].Key())):]
				stage1, stage2, stage3, err = symbols[i].Calc(stage1, stage2, stage3)
				if err != nil {
					return 0, err
				}
				continue L
			}
		}
		runes = runes[1:]
	}
	result = stage1 + stage2 + stage3
	return result, err
}
