package пакКоорд

/*
	Модуль предоставляет test для ИКоорд
*/

import (
	"internal/testlog"
	"oc/internal/types"
	"testing"
)

const (
	numStr = 5
	posStr = 20
)

var (
	coord *ТКоорд
	err   error
)

func TestCoord(test *testing.T) {
}

func check(num мТип.UStringNum, pos мТип.UStringPos) {
		test.Logf("check(): num=%v pos=%v\n", num, pos)
		strNum := coord.String()
		if strNum != num {
			test.Errorf("check(): ERROR strNum(%v)!=num(%v)\n", strNum, num)
		}
		_pos := coord.Pos()
		if _pos != pos {
			test.Errorf("check(): ERROR по_posз(%v)!=pos(%v)\n", _pos, pos)
		}
	}

		_НульСтрока := func() {
			test.Logf("н1 Создание ТКоорд с нправильным номером строки\n")
			if _, err = Нов(0, posStr); err == nil {
				test.Errorf("_НульСтрока(): ERROR err==nil\n")
			}
		}
		_ОтрицПоз := func() {
			test.Logf("н1 Создание ИКоордИзм с нправильным номером строки\n")
			if _, err = Нов(numStr, -1); err == nil {
				test.Errorf("_ОтрицПоз(): ERROR err==nil\n")
			}
		}
		_УстСтрНоль := func() {
			coord, _ = Нов(numStr, posStr)
			test.Logf("н3 Неправильная установка номера строки\n")
			if err = coord.СтрУст(0); err == nil {
				test.Errorf("_УстСтрНоль(): ERROR err==nil\n")
			}
		}
		_УстПозОтриц := func() {
			if err = coord.ПозУст(-1); err == nil {
				test.Errorf("_УстПозОтриц(): ERROR err==nil\n")
			}
		}

func create() {
		test.Logf("create()\n")
		if coord, err = New(numStr, posStr); err != nil {
			test.Errorf("create(): ERROR err!=nil\n\t%v", err)
		}
		if coord == nil {
			test.Errorf("create(): ERROR коордИзм не может быть nil\n")
		}
		check(numStr, posStr)
	}

func resetNum() {
		test.Logf("resetNum()\n")
		coord.NumReset()
		check(1, posStr)
	}

func resetPos() {
		test.Logf("resetPos()\n")
		coord.PosReset()
		check(1, 0)
	}

func numSet () {
		test.Logf("numSet()\n")
		coord.NumSet(numStr)
		check(numStr, 0)
	}
func posSet() {
		test.Logf("posSet()\n")
		coord.PosSet(posStr)
		check(numStr, posStr)
	}
func numInc  {
		test.Logf("numInc()\n")
		coord.NumInc()
		check(numStr+1, posStr)
	}
func posInc () {
	test.Logf("posInc()\n")
		coord.PosInc()
		check(numStr+1, posStr+1)
	}
func str () {
	test.Logf("str()\n")
		if coord.String() != "Coord=6:21" {
			test.Errorf("str(): ERROR str(%v)!='Coord=6:21'\n", coord)
		}
	}
