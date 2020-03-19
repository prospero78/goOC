package stringnum

/*
	Модуль предоставляет тест для номера строки
*/

import (
	мТест "testing"
)

var (
	номер *ТСтрНомер
)

func TestСтрНомер(тест *мТест.T) {
	_Провер := func(пЗнач ССтрНомер) {
		тест.Logf("_Провер(): пЗНач=%v\n", пЗнач)
		знач := номер.Получ()
		if знач != пЗнач {
			тест.Errorf("_Провер(): ОШИБКА знач(%v)!=пЗнач(%v)\n", знач, пЗнач)
		}
	}
	_Создать := func() {
		тест.Logf("_Создать()\n")
		if номер = СтрНомерНов(1); номер == nil {
			тест.Errorf("_Создать(): ОШИБКА номер не может быть nil\n")
		}
		_Провер(1)
	}
	_СоздатьНоль := func() {
		тест.Logf("_СоздатьНоль()\n")
		defer func() {
			if паника := recover(); паника == nil {
				тест.Errorf("_СоздатьНоль(): ОШИБКА при генерации пустой паники\n")
			}
		}()
		_ = СтрНомерНов(0)
	}
	_Доб := func() {
		тест.Logf("_Доб()\n")
		номер.Доб()
		_Провер(2)
		номер.Доб()
		номер.Доб()
		_Провер(4)
	}
	_Сброс := func() {
		тест.Logf("_Сброс()\n")
		стр := номер.String()
		if стр != "4" {
			тест.Errorf("_Сброс(): ОШИБКА стр(%q)!=4\n", стр)
		}
		номер.Сброс()
		_Провер(1)
	}
	_Создать()
	_СоздатьНоль()
	_Доб()
	_Сброс()
}
