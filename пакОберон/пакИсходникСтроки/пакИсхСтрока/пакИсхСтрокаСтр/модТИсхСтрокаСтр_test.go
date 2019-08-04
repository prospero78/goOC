package пакИсхСтрокаСтр

/*
	Модуль предоставляет весьма условный тест для типа исходной строки
*/

import (
	мТест "testing"
)

const (
	стрСтрока = "МОДУЛЬ модТест;"
)

var (
	строка ИИсхСтрокаСтр
	ош     error
)

func TestИсхСтрокаСтр(тест *мТест.T) {
	_Позитив := func() {
		{ //1 Создание строки
			тест.Logf("п1 Создание строки\n")
			if строка, ош = ИсхСтрокаСтрНов(стрСтрока); ош != nil {
				тест.Errorf("п1.1 ОШИБКА при создании строки исходника\n\t%v", ош)
			}
			if строка == nil {
				тест.Errorf("п1.2 ОШИБКА строка не должна быть nil\n")
			}
		}
		{ //2 Проверка начальных значений
			тест.Logf("п2 Проверка начальных значений\n")
			if строка.Получ() != стрСтрока {
				тест.Errorf("п2.1 ОШИБКА при хранении начального згачения строки(%v), стр=[%v]", стрСтрока, строка.Получ())
			}
		}
	}
	_Негатив := func() {
		{ //1 Создание строки с устым значением
			тест.Logf("н1 Создание строки с пустым значением\n")
			if строка, ош = ИсхСтрокаСтрНов(""); ош == nil {
				тест.Errorf("н1.1 ОШИБКА при создании строки исходника\n")
			}
			if строка != nil {
				тест.Errorf("н1.2 ОШИБКА строка должна быть nil\n")
			}
		}
	}
	_Позитив()
	_Негатив()
}
