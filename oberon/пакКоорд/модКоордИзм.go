package пакКоорд

/*
	Модуль предоставляет изменяемую координату
*/

import (
	мФмт "fmt"
	мСн "../пакСтрНомер"
	мСп "../пакСтрПоз"
	мСинх "sync"
)

//ИКоордИзм -- интерфейс для изменяемых координат
type ИКоордИзм interface {
	ИКоордФикс
	СтрНомерУст(мСн.ССтрНомер) error
	СтрНомерСброс()
	СтрНомерДоб()
	СтрПозУст(мСп.ССтрПоз) error
	СтрПозСброс()
	СтрПозДоб()
}


//тКоордИзм -- тип для хранения координат
type тКоордИзм struct {
	стр  мСн.ИСтрНомерИзм
	поз  мСп.ИСтрПозИзм
	блок мСинх.RWMutex
}


//КоордИзмНов -- возвращает новый экземпляр изменяемых координат
func КоордИзмНов(пСтр мСн.ССтрНомер, пПоз мСп.ССтрПоз) (коорд ИКоордИзм, ош error) {
	_коорд := тКоордИзм{}
	if _коорд.стр, ош = мСн.СтрНомерИзмНов(пСтр); ош != nil {
		return nil, мФмт.Errorf("КоордИзмНов(): ОШИБКА при создании номера строки\n\t%v", ош)
	}
	if _коорд.поз, ош = мСп.СтрПозИзмНов(пПоз); ош != nil {
		return nil, мФмт.Errorf("КоордИзмНов(): ОШИБКА при создании позиции в строке\n\t%v", ош)
	}
	return &_коорд, nil
}

//Возвращает строковое представление координаты
func (сам *тКоордИзм) String() string {
	return мФмт.Sprintf("Коорд: стр=%v поз=%v", сам.стр, сам.поз)
}

//СтрНомер -- возвращает ссылку на объект номера строки
func (сам *тКоордИзм) СтрНомер() мСн.ССтрНомер {
	return сам.стр.Получ()
}

//СтрНомерСброс -- сбрасывает номер строки
func (сам *тКоордИзм) СтрНомерСброс() {
	сам.стр.Сброс()
}

//СтрНомерДоб -- добавлет +1 номер строки
func (сам *тКоордИзм) СтрНомерДоб() {
	сам.стр.Доб()
}

//СтрНомерУст -- устанавливает номер строки
func (сам *тКоордИзм) СтрНомерУст(пНомер мСн.ССтрНомер) (ош error) {
	if ош = сам.стр.Уст(пНомер); ош != nil {
		return мФмт.Errorf("тКоордИзм.СтрНомерУст(): ошибка в установке номера строки\n\t%v", ош)
	}
	return nil
}

//СтрПоз -- возвращает ссылку на объект позиции в строке
func (сам *тКоордИзм) СтрПоз() мСп.ССтрПоз {
	return сам.поз.Получ()
}

//СтрПозСброс -- сбрасывает позицию строки
func (сам *тКоордИзм) СтрПозСброс() {
	сам.поз.Сброс()
}

//СтрПозДоб -- добавлет +1 позицию в строке
func (сам *тКоордИзм) СтрПозДоб() {
	сам.поз.Доб()
}

//СтрПозУст -- устанавливает позицию строки
func (сам *тКоордИзм) СтрПозУст(пПоз мСп.ССтрПоз) (ош error) {
	if ош = сам.поз.Уст(пПоз); ош != nil {
		return мФмт.Errorf("тКоордИзм.СтрПозУст(): ошибка в установке позиции в строке\n\t%v", ош)
	}
	return nil
}
