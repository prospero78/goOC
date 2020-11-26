package stringnum

/*
	Модуль предоставляет test для номера строки
*/

import (
	"oc/internal/types"
	"testing"
)

var (
	num *TStringNum
	err error
)

func TestStringNum(test *testing.T) {
	create0(test)
	create(test)
	inc(test)
	reset(test)
}

func check(test *testing.T, val types.UStringNum) {
	test.Logf("check(): пЗНач=%v\n", val)
	_val := num.Get()
	if _val != val {
		test.Errorf("check(): ERROR val(%v)!=%v\n", _val, val)
	}
}

func create(test *testing.T) {
	test.Logf("create()\n")
	if num, err = New(1); err != nil {
		test.Errorf("create(): ERROR err!=nil\n\t%v", err)

	}
	if num == nil {
		test.Errorf("create(): ERROR num==nil\n")
	}
	check(test, 1)
}
func create0(test *testing.T) {
	test.Logf("create0()\n")
	if _, err = New(0); err == nil {
		test.Errorf("create0(): ERROR err==nil\n")
	}
}

func inc(test *testing.T) {
	test.Logf("inc()\n")
	num.Inc()
	check(test, 2)
	num.Inc()
	num.Inc()
	check(test, 4)
}

func reset(test *testing.T) {
	test.Logf("reset()\n")
	стр := num.String()
	if стр != "4" {
		test.Errorf("reset(): ERROR стр(%q)!=4\n", стр)
	}
	num.Reset()
	check(test, 1)
	num.Inc()
	num.Reset()
	check(test, 1)
}
