package coord

/*
	Модуль предоставляет test для ИКоорд
*/

import (
	"fmt"
	"oc/internal/types"
	"testing"
)

const (
	numStr = 5
	posStr = 20
)

var (
	coord *TCoord
	err   error
)

func TestCoord(test *testing.T) {
}

func check(test *testing.T, num types.UStringNum, pos types.UStringPos) {
	test.Logf("check(): num=%v pos=%v\n", num, pos)
	strNum := coord.String()
	if strNum != fmt.Sprint(num) {
		test.Errorf("check(): ERROR strNum(%v)!=num(%v)\n", strNum, num)
	}
	_pos := coord.Pos()
	if _pos != pos {
		test.Errorf("check(): ERROR по_posз(%v)!=pos(%v)\n", _pos, pos)
	}
}

func _НульСтрока(test *testing.T) {
	test.Logf("н1 Создание ТКоорд с нправильным номером строки\n")
	if _, err = New(0, posStr); err == nil {
		test.Errorf("_НульСтрока(): ERROR err==nil\n")
	}
}
func _ОтрицПоз(test *testing.T) {
	test.Logf("н1 Создание ИКоордИзм с нправильным номером строки\n")
	if _, err = New(numStr, -1); err == nil {
		test.Errorf("_ОтрицПоз(): ERROR err==nil\n")
	}
}
func _УстСтрНоль(test *testing.T) {
	coord, _ = New(numStr, posStr)
	test.Logf("н3 Неправильная установка номера строки\n")
	if err = coord.NumSet(0); err == nil {
		test.Errorf("_УстСтрНоль(): ERROR err==nil\n")
	}
}
func _УстПозОтриц(test *testing.T) {
	if err = coord.NumSet(-1); err == nil {
		test.Errorf("_УстПозОтриц(): ERROR err==nil\n")
	}
}

func create(test *testing.T) {
	test.Logf("create()\n")
	if coord, err = New(numStr, posStr); err != nil {
		test.Errorf("create(): ERROR err!=nil\n\t%v", err)
	}
	if coord == nil {
		test.Errorf("create(): ERROR коордИзм не может быть nil\n")
	}
	check(test, numStr, posStr)
}

func resetNum(test *testing.T) {
	test.Logf("resetNum()\n")
	coord.NumReset()
	check(test, 1, posStr)
}

func resetPos(test *testing.T) {
	test.Logf("resetPos()\n")
	coord.PosReset()
	check(test, 1, 0)
}

func numSet(test *testing.T) {
	test.Logf("numSet()\n")
	coord.NumSet(numStr)
	check(test, numStr, 0)
}
func posSet(test *testing.T) {
	test.Logf("posSet()\n")
	coord.PosSet(posStr)
	check(test, numStr, posStr)
}
func numInc(test *testing.T) {
	test.Logf("numInc()\n")
	coord.NumInc()
	check(test, numStr+1, posStr)
}
func posInc(test *testing.T) {
	test.Logf("posInc()\n")
	coord.PosInc()
	check(test, numStr+1, posStr+1)
}
func str(test *testing.T) {
	test.Logf("str()\n")
	if coord.String() != "Coord=6:21" {
		test.Errorf("str(): ERROR str(%v)!='Coord=6:21'\n", coord)
	}
}
