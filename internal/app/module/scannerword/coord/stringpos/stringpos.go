package stringpos

/*
	Модуль предоставляет тип для операций с позицией в строке исходника
*/

import (
	"fmt"
	"oc/internal/types"
)

//TStringPos -- тип для операций с позицией в строке исходника
type TStringPos struct {
	val types.UStringPos
}

//New -- возвращает ссылку на TStringPos
func New(strPos types.UStringPos) (pos *TStringPos, ош error) {
	pos = &TStringPos{}
	if ош = pos.Set(strPos); ош != nil {
		return nil, fmt.Errorf("stringpos.go/New(): ERROR при установке позиции\n\t%v", ош)
	}
	return pos, nil
}

//Inc -- добавляет +1 к значению позиции строки
func (sf *TStringPos) Inc() {
	sf.val++
}

//Reset -- сбрасывает значение позиции строки
func (sf *TStringPos) Reset() {
	sf.val = 0
}

//Get -- возвращает значение позиции в строке
func (sf *TStringPos) Get() types.UStringPos {
	return sf.val
}

func (sf *TStringPos) String() string {
	return fmt.Sprint(sf.val)
}

//Set -- установка значения позиции строки
func (sf *TStringPos) Set(pos types.UStringPos) (ош error) {
	if pos < 0 {
		return fmt.Errorf("TStringPos.Set(): ERROR pos(%v)<0", pos)
	}
	sf.val = pos
	return nil
}
