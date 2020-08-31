package stringpos

/*
	Модуль предоставляет тип для операций с позицией в строке исходника
*/

import (
	"fmt"
	мТип "oc/internal/types"
)

//ТСтрокаПоз -- тип для операций с позицией в строке исходника
type ТСтрокаПоз struct {
	знач мТип.ССтрокаПоз
}

//Нов -- возвращает ссылку на ТСтрокаПоз
func Нов(пПоз мТип.ССтрокаПоз) (поз *ТСтрокаПоз, ош error) {
	_поз := ТСтрокаПоз{}
	if ош = _поз.Уст(пПоз); ош != nil {
		return nil, fmt.Errorf("stringpos.go/Нов(): ОШИБКА при установке позиции\n\t%v", ош)
	}
	return &_поз, nil
}

//Доб -- добавляет +1 к значению позиции строки
func (сам *ТСтрокаПоз) Доб() {
	сам.знач++
}

//Сброс -- сбрасывает значение позиции строки
func (сам *ТСтрокаПоз) Сброс() {
	сам.знач = 0
}

//Получ -- возвращает значение позиции в строке
func (сам *ТСтрокаПоз) Получ() мТип.ССтрокаПоз {
	return сам.знач
}

func (сам *ТСтрокаПоз) String() string {
	return fmt.Sprint(сам.знач)
}

//Уст -- установка значения позиции строки
func (сам *ТСтрокаПоз) Уст(пПоз мТип.ССтрокаПоз) (ош error) {
	if пПоз < 0 {
		return fmt.Errorf("ТСтрокаПоз.Уст(): ОШИБКА  пПоз(%v)<0", пПоз)
	}
	сам.знач = пПоз
	return nil
}
