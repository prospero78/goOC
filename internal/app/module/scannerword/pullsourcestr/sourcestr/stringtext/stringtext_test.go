package stringtext

/*
	Модуль предоставляет весьма условный тест для типа исходной строки
*/

import (
	"testing"
)

const (
	strSource = "МОДУЛЬ модТест;"
)

var (
	str *TSourceString
	err error
)

func TestSourceString(test *testing.T) {
	createEmpty(test)
	create(test)
	check(test)
}

func createEmpty(test *testing.T) {
	test.Logf("createEmpty()\n")
	if str, err = New(""); err == nil {
		test.Errorf("createEmpty(): ERROR err==nil\n")
	}
	if str != nil {
		test.Errorf("createEmpty(): ERROR str!=nil\n")
	}
}
func create(test *testing.T) {
	test.Logf("create()\n")
	if str, err = New(strSource); err != nil {
		test.Errorf("create(): ERROR err!=nil\n\t%v", err)
	}
	if str == nil {
		test.Errorf("create(): ERROR str==nil\n")
	}
}

func check(test *testing.T) {
	test.Logf("check()\n")
	if str.Get() != strSource {
		test.Errorf("check(): ERROR strSource(%v)!=val(%v)", strSource, str.Get())
	}
}
