package numstr

/*
	Файл предоставляет тест для номера строки исходника
*/

import (
	"testing"

	"github.com/prospero78/goOC/internal/types"
)

var (
	iNum types.INumStr
	num  *TNumStr
	err  error
)

func TestNumStr(test *testing.T) {
	create(test)
	set(test)
}

func set(test *testing.T) {
	test.Logf("set()\n")
	{ // Неправильное присвоение
		if err = iNum.Set(-3); err == nil {
			test.Errorf("set(): err==nil\n")
		}
		if val := iNum.Get(); val != 1 {
			test.Errorf("set(): val(%v)!=1\n", val)
		}
	}
	{ // Правильное присвоение
		if err = iNum.Set(7); err != nil {
			test.Errorf("set(): err=%v\n", err)
		}
		if val := iNum.Get(); val != 7 {
			test.Errorf("set(): val(%v)!=7\n", val)
		}
	}
}

func create(test *testing.T) {
	test.Logf("create()\n")
	{ // Неправильный номер строки
		if iNum, err = New(0); err == nil {
			test.Errorf("create(): err==nil\n")
		}
		if iNum != nil {
			test.Errorf("create(): iNum!=nil\n")
		}
	}
	{ // Правильный номер строки
		if iNum, err = New(1); err != nil {
			test.Errorf("create(): err=%v\n", err)
		}
		if iNum == nil {
			test.Errorf("create(): iNum==nil\n")
		}
		var ok bool
		num, ok = iNum.(*TNumStr)
		if !ok {
			test.Errorf("create(): iNum not convert to *TNumStr\n")
		}
		if num == nil {
			test.Errorf("create(): num==nil\n")
		}
		if val := iNum.Get(); val != 1 {
			test.Errorf("create(): val(%v)!=1\n", val)
		}
	}
}
