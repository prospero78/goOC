package coord

/*
	Исходник предоставляет тест для координат литеры
*/

import (
	"testing"

	"github.com/prospero78/goOC/internal/types"
)

var (
	iCoord types.ICoord
	err    error
)

func TestCoord(test *testing.T) {
	create(test)
	setpos(test)
	setnum(test)
	set(test)
}

func set(test *testing.T) {
	test.Logf("set()\n")
	if err = iCoord.Set(0, 1); err == nil {
		test.Errorf("set(): err==nil")
	}
	if err = iCoord.Set(1, -1); err == nil {
		test.Errorf("set(): err==nil")
	}
	if err = iCoord.Set(4, 6); err != nil {
		test.Errorf("set(): err=%v", err)
	}
}

func setnum(test *testing.T) {
	test.Logf("setnum()\n")
	{ // Неправильная позиция
		if err = iCoord.SetNumStr(0); err == nil {
			test.Errorf("setnum(): err==nil\n")
		}
		if val := iCoord.NumStr(); val != 1 {
			test.Errorf("setnum(): val(%v)!=1\n", val)
		}
	}
	{ // Правильная позиция
		if err = iCoord.SetNumStr(2); err != nil {
			test.Errorf("setnum(): err=%v\n", err)
		}
		if val := iCoord.NumStr(); val != 2 {
			test.Errorf("setnum(): val(%v)!=2\n", val)
		}
	}
}

func setpos(test *testing.T) {
	test.Logf("setpos()\n")
	{ // Неправильная позиция
		if err = iCoord.SetPos(-1); err == nil {
			test.Errorf("setpos(): err==nil\n")
		}
		if val := iCoord.Pos(); val != 0 {
			test.Errorf("setpos(): val(%v)!=0\n", val)
		}
	}
	{ // Правильная позиция
		if err = iCoord.SetPos(4); err != nil {
			test.Errorf("setpos(): err=%v\n", err)
		}
		if val := iCoord.Pos(); val != 4 {
			test.Errorf("setpos(): val(%v)!=4\n", val)
		}
	}
}

func create(test *testing.T) {
	test.Logf("create()\n")
	{ // Неправильная позиция
		if iCoord, err = New(4, -1); err == nil {
			test.Errorf("create(): err==nil\n")
		}
		if iCoord != nil {
			test.Errorf("create(): iCoord!=nil\n")
		}
	}
	{ // Неправильный номер строки
		if iCoord, err = New(0, 0); err == nil {
			test.Errorf("create(): err==nil\n")
		}
		if iCoord != nil {
			test.Errorf("create(): iCoord!=nil\n")
		}
	}
	{ // Правильное создание
		if iCoord, err = New(1, 0); err != nil {
			test.Errorf("create(): err=%v\n", err)
		}
		if iCoord == nil {
			test.Errorf("create(): iCoord==nil\n")
		}
	}
	if val := iCoord.Pos(); val != 0 {
		test.Errorf("create(): val(%v)!=0\n", val)
	}
	if val := iCoord.NumStr(); val != 1 {
		test.Errorf("create(): val(%v)!=1\n", val)
	}
}
