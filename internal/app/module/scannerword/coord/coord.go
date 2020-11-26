package coord

/*
	Модуль предоставляет потокобезопасный тип координаты.
*/

import (
	"fmt"
	мСп "oc/internal/app/module/scannerword/coord/stringpos"
	мСн "oc/internal/app/module/scannerword/pullsourcestr/sourcestr/stringnum"
	мТип "oc/internal/types"
)

//TCoord -- тип для хранения координат
type TCoord struct {
	стр *мСн.ТСтрокаНом
	поз *мСп.ТСтрокаПоз
}

//Нов -- возвращает указатель на новый TCoord
func Нов(пСтр мТип.UStringNum, пПоз мТип.UStringPos) (коорд *TCoord, ош error) {
	_коорд := TCoord{}
	if _коорд.стр, ош = мСн.Нов(пСтр); ош != nil {
		return nil, fmt.Errorf("coord.go/Нов(): ERROR при создании номера строки\n\t%v", ош)
	}
	if _коорд.поз, ош = мСп.Нов(пПоз); ош != nil {
		return nil, fmt.Errorf("coord.go/Нов(): ERROR при создании позиции в строке строки\n\t%v", ош)
	}
	return &_коорд, nil
}

//Возвращает строковое представление координаты
func (сам *TCoord) String() string {
	return fmt.Sprintf("Коорд: стр=%v поз=%v", сам.стр, сам.поз)
}

//Стр -- возвращает хранимый номера строки
func (сам *TCoord) Стр() мТип.UStringNum {
	return сам.стр.Получ()
}

//СтрДоб -- добавлет +1 номер строки
func (сам *TCoord) СтрДоб() {
	сам.стр.Доб()
}

//СтрУст -- устанавливает номер строки
func (сам *TCoord) СтрУст(пНомер мТип.UStringNum) (ош error) {
	if ош = сам.стр.Уст(пНомер); ош != nil {
		return fmt.Errorf("TCoord.СтрУст(): ERROR при установке номера строки\n\t%v", ош)
	}
	return nil
}

//СтрСброс -- сброасывает номер строки
func (сам *TCoord) СтрСброс() {
	сам.стр.Сброс()
}

//Поз -- возвращает хранимую позицию в строке
func (сам *TCoord) Поз() мТип.UStringPos {
	return сам.поз.Получ()
}

//ПозСброс -- сбрасывает позицию строки
func (сам *TCoord) ПозСброс() {
	сам.поз.Сброс()
}

//СтрПозДоб -- добавлет +1 позицию в строке
func (сам *TCoord) ПозДоб() {
	сам.поз.Доб()
}

//СтрПозУст -- устанавливает позицию строки
func (сам *TCoord) ПозУст(пПоз мТип.UStringPos) (ош error) {
	if ош = сам.поз.Уст(пПоз); ош != nil {
		return fmt.Errorf("TCoord.ПозУст(): ERROR при установке позиции в строке\n\t%v", ош)
	}
	return nil
}
