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

//Нов -- возвращает ссылку на TStringPos
func Нов(пПоз types.UStringPos) (поз *TStringPos, ош error) {
	_поз := TStringPos{}
	if ош = _поз.Уст(пПоз); ош != nil {
		return nil, fmt.Errorf("stringpos.go/Нов(): ОШИБКА при установке позиции\n\t%v", ош)
	}
	return &_поз, nil
}

//Доб -- добавляет +1 к значению позиции строки
func (сам *TStringPos) Доб() {
	сам.val++
}

//Сброс -- сбрасывает значение позиции строки
func (сам *TStringPos) Сброс() {
	сам.val = 0
}

//Получ -- возвращает значение позиции в строке
func (сам *TStringPos) Получ() types.UStringPos {
	return сам.val
}

func (сам *TStringPos) String() string {
	return fmt.Sprint(сам.val)
}

//Уст -- установка значения позиции строки
func (сам *TStringPos) Уст(пПоз types.UStringPos) (ош error) {
	if пПоз < 0 {
		return fmt.Errorf("TStringPos.Уст(): ОШИБКА  пПоз(%v)<0", пПоз)
	}
	сам.val = пПоз
	return nil
}
