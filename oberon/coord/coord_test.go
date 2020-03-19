package пакКоорд

/*
	Модуль предоставляет тест для ИКоорд
*/

import (
	мСн "../stringnum"
	мПоз "../stringpos"
	мТест "testing"
)

const (
	номСтр = 5
	позСтр = 20
)

var (
	коорд *ТКоорд
	ош    error
)

func TestКоорд(тест *мТест.T) {
	_Провер := func(пНом мСн.ССтрНомер, пПоз мПоз.ССтрПоз) {
		тест.Logf("_Провер(): пПСтр=%v пПоз=%v\n", пНом, пПоз)
		номер := коорд.СтрНомер()
		if номер != пНом {
			тест.Errorf("_Провер(): ОШИБКА номер(%v)!=пНом(%v)\n", номер, пНом)
		}
		поз := коорд.СтрПоз()
		if поз != пПоз {
			тест.Errorf("_Провер(): ОШИБКА поз(%v)!=пПоз(%v)\n", поз, пПоз)
		}
	}
	_Создать := func() {
		тест.Logf("_Создать(): ТКоорд\n")
		if коорд = КоордНов(номСтр, позСтр); коорд == nil {
			тест.Errorf("п1.2 ОШИБКА коордИзм не может быть nil\n")
		}
		_Провер(номСтр, позСтр)
	}
	_СбросСтр := func() {
		тест.Logf("_СбросСтр()\n")
		коорд.СтрНомерСброс()
		_Провер(1, позСтр)
	}
	_СбросПоз := func() {
		тест.Logf("_СбросПоз()\n")
		коорд.СтрПозСброс()
		_Провер(1, 0)
	}
	_НомерУст := func() {
		тест.Logf("_НомерУст()\n")
		коорд.СтрНомерУст(номСтр)
		_Провер(номСтр, 0)
	}
	_ПозУст := func() {
		тест.Logf("_ПозУст()\n")
		коорд.СтрПозУст(позСтр)
		_Провер(номСтр, позСтр)
	}
	_СтрНомДоб := func() {
		тест.Logf("_СтрНомДоб()\n")
		коорд.СтрНомерДоб()
		_Провер(номСтр+1, позСтр)
	}
	_ПозДоб := func() {
		коорд.СтрПозДоб()
		_Провер(номСтр+1, позСтр+1)
	}
	_Стр := func() {
		тест.Logf("п9 Проверка на строку\n")
		if коорд.String() != "Коорд: стр=6 поз=21" {
			тест.Errorf("п9.1 ОШИБКА при получении строкового представления, знач=[%v]\n", коорд)
		}
	}
	_Позитив := func() {
		_Создать()
		_СбросСтр()
		_СбросПоз()
		_НомерУст()
		_ПозУст()
		_СтрНомДоб()
		_ПозДоб()
		_Стр()
	}
	_Негатив := func() {
		_НульСтрока := func() {
			тест.Logf("н1 Создание ТКоорд с нправильным номером строки\n")
			defer func() {
				if паника := recover(); паника == nil {
					тест.Errorf("_НульСтрока(): ОШИБКА в генерации пустой паники\n")
				}
			}()
			_ = КоордНов(0, позСтр)
		}
		_ОтрицПоз := func() {
			тест.Logf("н1 Создание ИКоордИзм с нправильным номером строки\n")
			defer func() {
				if паника := recover(); паника == nil {
					тест.Errorf("_ОтрицПоз(): ОШИБКА в генерации пустой паники\n")
				}
			}()
			_ = КоордНов(номСтр, -1)
		}
		_УстСтрНоль := func() {
			коорд = КоордНов(номСтр, позСтр)
			тест.Logf("н3 Неправильная установка номера строки\n")
			defer func() {
				if паника := recover(); паника == nil {
					тест.Errorf("_УстСтрНоль(): ОШИБКА в генерации пустой паники\n")
				}
			}()
			коорд.СтрНомерУст(0)
		}
		_УстПозОтриц := func() {
			тест.Logf("н4 Неправильная установка позиции в строке\n")
			defer func() {
				if паника := recover(); паника == nil {
					тест.Errorf("_УстПозОтриц(): ОШИБКА в генерации пустой паники\n")
				}
			}()
			коорд.СтрПозУст(-1)
		}
		_НульСтрока()
		_ОтрицПоз()
		_УстСтрНоль()
		_УстПозОтриц()
	}
	_Позитив()
	_Негатив()
}
