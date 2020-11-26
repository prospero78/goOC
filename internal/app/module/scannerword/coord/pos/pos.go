package pos

/*
	Модуль предоставляет тип для операций с позицией в строке исходника
*/

import (
	"fmt"
	"oc/internal/types"
)

//TPos -- тип для операций с позицией в строке исходника
type TPos struct {
	val types.UStringPos
}

//New -- возвращает ссылку на TPos
func New(strPos types.UStringPos) (pos *TPos, ош error) {
	pos = &TPos{}
	if ош = pos.Set(strPos); ош != nil {
		return nil, fmt.Errorf("stringpos.go/New(): ERROR при установке позиции\n\t%v", ош)
	}
	return pos, nil
}

//Inc -- добавляет +1 к значению позиции строки
func (sf *TPos) Inc() {
	sf.val++
}

//Reset -- сбрасывает значение позиции строки
func (sf *TPos) Reset() {
	sf.val = 0
}

//Get -- возвращает значение позиции в строке
func (sf *TPos) Get() types.UStringPos {
	return sf.val
}

func (sf *TPos) String() string {
	return fmt.Sprint(sf.val)
}

//Set -- установка значения позиции строки
func (sf *TPos) Set(pos types.UStringPos) (ош error) {
	if pos < 0 {
		return fmt.Errorf("TPos.Set(): ERROR pos(%v)<0", pos)
	}
	sf.val = pos
	return nil
}
