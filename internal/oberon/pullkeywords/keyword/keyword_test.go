package keyword

/*
	Модуль предоставляет тест для кейворда
*/

import (
	мТест "testing"
)

var (
	ключ *ТКлюч
)

func TestКлюч(тест *мТест.T) {
	_Провер := func(пКлюч СКлюч, пДлина int) {
		тест.Logf("_Провер(): пКлюч=%q пДлина=%v\n", пКлюч, пДлина)
		длин := len(ключ.пул)
		if длин != пДлина {
			тест.Errorf("_Провер(): ОШИБКА длин(%v)!=пДлин(%v)\n", длин, пДлина)
		}
		if !ключ.ЕслиСовпал(пКлюч) {
			тест.Errorf("_Провер(): ОШИБКА пКлюч(%v) не найден\n", пКлюч)
		}
	}
	_СоздатьПусто := func() {
		тест.Logf("_СоздатьПусто()\n")
		defer func() {
			if паника := recover(); паника == nil {
				тест.Errorf("_СоздатьПусто(): ОШИБКА при генерации пустой пники\n")
			}
		}()
		_ = КлючНов("")
	}
	_Создать := func() {
		тест.Logf("_Создать()\n")
		if ключ = КлючНов("тест"); ключ == nil {
			тест.Errorf("_Создать(): ОШИБКА ключ не может быть nil\n")
		}
		_Провер("тест", 1)
	}
	_Доб := func() {
		тест.Logf("_Доб()\n")
		ключ.Доб("тест")
		_Провер("тест", 1)

		ключ.Доб("тест1")
		_Провер("тест", 2)
		_Провер("тест1", 2)
	}
	_СоздатьПусто()
	_Создать()
	_Доб()
}