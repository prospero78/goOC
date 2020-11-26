package stringpos

/*
	Модуль предоставляет тест для позиции в строке исходника
*/

import (
	"fmt"
	"oc/internal/types"
	"testing"
)

var (
	pos  *TStringPos
	err  error
	isOk bool
)

func TestStringPos(test *testing.T) {
	createNegativePos(test)
	create(test)
	add(test)
	reset(test)
}

func add(test *testing.T) {
	test.Logf("add()\n")
	pos.Inc()
	check(test, 11)
	pos.Inc()
	check(test, 12)
}

func reset(test *testing.T) {
	test.Logf("reset()\n")
	pos.Reset()
	check(test, 0)
	pos.Inc()
	pos.Reset()
	check(test, 0)
}

func createNegativePos(test *testing.T) {
	test.Logf("createNegativePos()\n")
	if pos, err = New(-1); err == nil {
		test.Errorf("createNegativePos(): ERROR err==nil\n")
	}
	if pos != nil {
		test.Errorf("createNegativePos(): ERROR pos!=nil\n")
	}
}

func create(test *testing.T) {
	test.Logf("create()\n")
	if pos, err = New(10); err != nil {
		test.Errorf("create(): ERROR err!=nil\n\t%v", err)
	}
	if pos == nil {
		test.Errorf("п1.1 ERROR поз не может быть nil\n")
	}
	check(test, 10)
}

func check(test *testing.T, val types.UStringPos) {
	test.Logf("check(): val=%v\n", val)
	if pos.Get() != val {
		test.Errorf("check(): ERROR in save default value(%v), pos=%v\n", val, pos)
	}
	if pos.String() != fmt.Sprint(val) {
		test.Errorf("check(): ERROR in save strings value(%v), pos=%s", val, pos)
	}
}
