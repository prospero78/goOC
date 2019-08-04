package пакСловоРод

/*
	Модуль предоставляет тест для рода слова
*/

import (
	мТест "testing"
)

var (
	род ИСловоРод
	ош  error
)

func TestСловоРод(тест *мТест.T) {
	_Позитив := func() {
		{ //1 Создание точки с запятой
			тест.Logf("п1 Создание точки с запятой\n")
			if род, ош = СловоРодНов(";"); ош != nil {
				тест.Errorf("п1.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КТочкаЗапятая {
				тест.Errorf("п1.2 ОШИБКА в хранении рода(;), знач=[%v]\n", род.Получ())
			}
		}
		{ //2 Создание  запятой
			тест.Logf("п2 Создание запятой\n")
			if род, ош = СловоРодНов(","); ош != nil {
				тест.Errorf("п2.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КЗапятая {
				тест.Errorf("п2.2 ОШИБКА в хранении рода(,), знач=[%v]\n", род.Получ())
			}
		}
		{ //3 Создание  плюса
			тест.Logf("п3 Создание плюса\n")
			if род, ош = СловоРодНов("+"); ош != nil {
				тест.Errorf("п3.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КПлюс {
				тест.Errorf("п3.2 ОШИБКА в хранении рода(+), знач=[%v]\n", род.Получ())
			}
		}
		{ //4 Создание  минуса
			тест.Logf("п4 Создание минуса\n")
			if род, ош = СловоРодНов("-"); ош != nil {
				тест.Errorf("п4.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КМинус {
				тест.Errorf("п4.2 ОШИБКА в хранении рода(-), знач=[%v]\n", род.Получ())
			}
		}
		{ //5 Создание  делить
			тест.Logf("п5 Создание делить\n")
			if род, ош = СловоРодНов("/"); ош != nil {
				тест.Errorf("п5.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КДелить {
				тест.Errorf("п5.2 ОШИБКА в хранении рода(/), знач=[%v]\n", род.Получ())
			}
		}
		{ //6 Создание скобка открыть круглая
			тест.Logf("п6 Создание скобка открыть круглая\n")
			if род, ош = СловоРодНов("("); ош != nil {
				тест.Errorf("п6.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КСкобкаОткрКругл {
				тест.Errorf("п6.2 ОШИБКА в хранении рода\"(\", знач=[%v]\n", род.Получ())
			}
		}
		{ //7 Создание коммент начать
			тест.Logf("п7 Создание коммент начать\n")
			if род, ош = СловоРодНов("(*"); ош != nil {
				тест.Errorf("п7.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != ККомментНачать {
				тест.Errorf("п7.2 ОШИБКА в хранении рода\"(*\", знач=[%v]\n", род.Получ())
			}
		}
		{ //8 Создание коммент закончить
			тест.Logf("п7 Создание коммент закончить\n")
			if род, ош = СловоРодНов("*)"); ош != nil {
				тест.Errorf("п7.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != ККомментЗакончить {
				тест.Errorf("п7.2 ОШИБКА в хранении рода\"*)\", знач=[%v]\n", род.Получ())
			}
		}
		{ //9 Создание скобка закрыть круглая
			тест.Logf("п9 Создание скобка закрыть круглая\n")
			if род, ош = СловоРодНов(")"); ош != nil {
				тест.Errorf("п9.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КСкобкаЗакрКругл {
				тест.Errorf("п9.2 ОШИБКА в хранении рода\")\", знач=[%v]\n", род.Получ())
			}
		}
	}
	_Позитив2 := func() {
		{ //10 Создание умноижить
			тест.Logf("п10 Создание умножить\n")
			if род, ош = СловоРодНов("*"); ош != nil {
				тест.Errorf("п10.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КУмножить {
				тест.Errorf("п10.2 ОШИБКА в хранении рода\")\", знач=[%v]\n", род.Получ())
			}
		}
		{ //11 Создание присвоить
			тест.Logf("п11 Создание присвоить\n")
			if род, ош = СловоРодНов(":="); ош != nil {
				тест.Errorf("п11.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КПрисвоить {
				тест.Errorf("п11.2 ОШИБКА в хранении рода\":=\", знач=[%v]\n", род.Получ())
			}
		}
		{ //12 Создание определить
			тест.Logf("п12 Создание определить\n")
			if род, ош = СловоРодНов(":"); ош != nil {
				тест.Errorf("п12.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КОпределить {
				тест.Errorf("п12.2 ОШИБКА в хранении рода\":\", знач=[%v]\n", род.Получ())
			}
		}
		{ //13 Создание равно
			тест.Logf("п13 Создание равно\n")
			if род, ош = СловоРодНов("="); ош != nil {
				тест.Errorf("п13.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КРавно {
				тест.Errorf("п13.2 ОШИБКА в хранении рода\"=\", знач=[%v]\n", род.Получ())
			}
		}
		{ //14 Создание точка
			тест.Logf("п14 Создание точка\n")
			if род, ош = СловоРодНов("."); ош != nil {
				тест.Errorf("п14.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КТочка {
				тест.Errorf("п14.2 ОШИБКА в хранении рода\".\", знач=[%v]\n", род.Получ())
			}
		}
		{ //15 Создание строка
			тест.Logf("п15 Создание строка\n")
			if род, ош = СловоРодНов("\"строка\""); ош != nil {
				тест.Errorf("п15.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КСтрока {
				тест.Errorf("п15.2 ОШИБКА в хранении рода(\"строка\"), знач=[%v]\n", род.Получ())
			}
		}
		{ //16 Создание имени
			тест.Logf("п16 Создание имени\n")
			if род, ош = СловоРодНов("ФункСтрока"); ош != nil {
				тест.Errorf("п16.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КИмя {
				тест.Errorf("п16.2 ОШИБКА в хранении рода \"ФункСтрока\", знач=[%v]\n", род.Получ())
			}
		}
		{ //17 Создание числа
			тест.Logf("п17 Создание числа\n")
			if род, ош = СловоРодНов("012345"); ош != nil {
				тест.Errorf("п17.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род.Получ() != КЧисло {
				тест.Errorf("п17.2 ОШИБКА в хранении рода \"012345\", знач=[%v]\n", род.Получ())
			}
		}
	}
	_Негатив := func() {
		{ //1 Передача пустой строки
			тест.Logf("н1 Передача пустой строки\n")
			if род, ош = СловоРодНов(""); ош == nil {
				тест.Errorf("н1.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род != nil {
				тест.Errorf("н1.2 ОШИБКА род должен быть nil\n")
			}
		}
		{ //2 Передача неизвестной литеры
			тест.Logf("н2 Передача неизвестной литеры\n")
			if род, ош = СловоРодНов("%"); ош == nil {
				тест.Errorf("н2.1 ОШИБКА при создании рода\n\t%v", ош)
			}
			if род != nil {
				тест.Errorf("н2.2 ОШИБКА род должен быть nil\n")
			}
		}
	}
	_Позитив()
	_Позитив2()
	_Негатив()
}
