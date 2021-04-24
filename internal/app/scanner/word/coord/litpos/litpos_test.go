package litpos

/*
	Тест для позиции литеры в строке
*/

import (
	"testing"

	"github.com/prospero78/goOC/internal/types"
)

const (
	_pos = 7
)

var (
	iPos types.IPos
	pos  *TPos
	err  error
)

func TestPos(test *testing.T) {
	create(test)
	set(test)
	inc(test)
	reset(test)
}

func reset(test *testing.T) {
	test.Logf("reset()\n")
	iPos.Reset()
	if val := iPos.Get(); val != 0 {
		test.Errorf("reset(): val(%v)!=0\n", val)
	}
}

func inc(test *testing.T) {
	test.Logf("inc()\n")
	iPos.Inc()
	if val := iPos.Get(); val != 9 {
		test.Errorf("inc(): val(%v)!=9\n", val)
	}
}

func set(test *testing.T) {
	test.Logf("test()\n")
	{ // Неправильное присовение
		if err = iPos.Set(-1); err == nil {
			test.Errorf("set(): err==nil\n")
		}
		if val := iPos.Get(); val != 7 {
			test.Errorf("set(): val(%v)!=7\n", val)
		}
	}
	{ // Правильное присовение
		if err = iPos.Set(8); err != nil {
			test.Errorf("set(): err=%v\n", err)
		}
		if val := iPos.Get(); val != 8 {
			test.Errorf("set(): val(%v)!=8\n", val)
		}
	}
}

func create(test *testing.T) {
	test.Logf("create()\n")
	{ // Неправильное создание
		iPos, err = New(-1)
		if err == nil {
			test.Errorf("create(): err==nil\n")
		}
		if iPos != nil {
			test.Errorf("create(): pos!=nil\n")
		}
	}
	{ // Правильное создание
		iPos, err = New(_pos)
		if err != nil {
			test.Errorf("create(): err=%v\n", err)
		}
		if iPos == nil {
			test.Errorf("create(): pos==nil\n")
		}
		var ok bool
		pos, ok = iPos.(*TPos)
		if !ok {
			test.Errorf("create(): нельзя конвертировать из iPos в *TPos\n")
		}
		if pos == nil {
			test.Errorf("create(): pos==nil")
		}
	}

}
