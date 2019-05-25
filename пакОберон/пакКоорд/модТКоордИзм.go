package пакКоорд

// модКоордИзм

/*
	Модуль описывает тип с изменяемыми координатами
*/

import (
	//пакКонс "../../../пакКонсоль"
	мАбс "../пакИнтерфейсы"
	мФмт "fmt"
)

//ТКоордИзм -- тип для изменяемых координат
type ТКоордИзм struct {
	*ТКоорд
}

//КоордИзмНов -- возвращает новый экземпляр изменяемых координат
func КоордИзмНов(пСтр мАбс.СКоордСтр, пПоз мАбс.СКоордПоз) (ки *ТКоордИзм, ош error) {
	_коорд, ош := КоордНов(пСтр, пПоз)
	if ош != nil {
		return nil, мФмт.Errorf("КоордИзмНов(): ошибка при создании ТКоордИзм\n\t%v", ош)
	}
	ки = &ТКоордИзм{
		ТКоорд: _коорд}
	if ки == nil {
		return nil, мФмт.Errorf("КоордИзмНов(): ошибка при создании ТКоордИзм\n\t%v", ош)
	}
	if ош = ки._ПозУст(пПоз); ош != nil {
		return nil, мФмт.Errorf("КоордИзмНов(): ошибка при создании ТКоордПозИзм\n\t%v", ош)
	}
	if ош := ки._СтрУст(пСтр); ош != nil {
		return nil, мФмт.Errorf("КоордИзмНов(): ошибка при создании ТКоордСтрИзм\n\t%v", ош)
	}
	return ки, nil
}

//String -- для соответствия интерфейсу Stringer
func (сам *ТКоордИзм) String() string {
	return "*ТКоордИзм{стр:" + мФмт.Sprintf("%v", сам.стр) + ", поз:" + мФмт.Sprintf("%v", сам.поз) + "}"
}

//ПозУст -- устанавливает позицию встроке
func (сам *ТКоордИзм) ПозУст(пПоз мАбс.СКоордПоз) error {
	return сам._ПозУст(пПоз)
}

//СтрУст -- устанавливает номер строки
func (сам *ТКоордИзм) СтрУст(пСтр мАбс.СКоордСтр) error {
	return сам._СтрУст(пСтр)
}

//ПозДоб -- увеличивает позициию в строке на +1
func (сам *ТКоордИзм) ПозДоб() {
	сам.поз++
}

//СтрДоб -- увеличивает строку на +1
func (сам *ТКоордИзм) СтрДоб() {
	сам.стр++
}

//ПозСброс -- сбрасывает позицию в строке
func (сам *ТКоордИзм) ПозСброс() {
	сам.поз = 0
}
