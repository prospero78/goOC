package sourcefile

/*
	Модуль предоставляет тест для типа исходного файла
*/

import (
	мТест "testing"
)

const (
	стрФайл = "./test.o7"
)

var (
	исхФайл *ТИсхФайл
	ош      error
)

func TestИсходникФайл(тест *мТест.T) {
	_Позитив := func() {
		{ //1 Создание файла исходника
			тест.Logf("п1 Создание ИсходникФайл\n")
			if исхФайл = ИсхФайлНов(стрФайл); исхФайл == nil {
				тест.Errorf("п1.2 ОШИБКА ИсхФайл не должен быть nil\n")
			}
		}
		{ //3 Считывание файла исходника
			тест.Logf("п3 Считывание файла исходника\n")
			if исхФайл.Размер() != 2593 {
				тест.Errorf("п3.1 ОШИБКА при проверке размера текста исходника(2593), размер=[%v]\n", исхФайл.Размер())
			}
		}
		{ //4 Проверка литеры в указанной позиции
			тест.Logf("п4 Проверка литеры в указанной позиции\n")
			лит, ош := исхФайл.Лит(36)
			if ош != nil {
				тест.Errorf("п4.1 ОШИБКА при проверке литеры в позиции 36(ж), \n\t%v", ош)
			}
			if лит != "н" {
				тест.Errorf("п4.2 ОШИБКА при проверке литеры в позиции(36, н), лит=[%v]\n", лит)
			}
		}
		{ //5 Печать исходника
			тест.Logf("п5 Печать исходника\n")
			исхФайл.Печать()
		}
	}
	_Негатив := func() {
		{ //1 Создание файла исходника с пустым именем
			тест.Logf("н1 Создание ИсходникФайл\n")
			if исхФайл = ИсхФайлНов(""); исхФайл != nil {
				тест.Errorf("н1.2 ОШИБКА ИсхФайл должен быть nil\n")
			}
		}
		{ //2 Считывание файла исходника с пустым именем и отсутствующим на диске
			тест.Logf("н2 Получени размера\n")
			исхФайл = ИсхФайлНов(стрФайл)
			if исхФайл.Размер() == 0 {
				тест.Errorf("н2.2 ОШИБКА при проверке размера текста исходника\n")
			}
		}
		{ //3 Проверка литеры с несуществующей позицией
			тест.Logf("н3 Проверка литеры в указанной позиции\n")
			лит, ош := исхФайл.Лит(8500)
			if ош == nil {
				тест.Errorf("н3.1 ОШИБКА при проверке литеры в указанной позиции\n")
			}
			if лит != "" {
				тест.Errorf("н3.2 ОШИБКА при проверке литеры в указанной позиции\n")
			}
		}
		{ //4 Проверка литеры с отрицательной позицией
			тест.Logf("н4 Проверка литеры в указанной позиции\n")
			лит, ош := исхФайл.Лит(-5)
			if ош == nil {
				тест.Errorf("н4.1 ОШИБКА при проверке литеры в указанной позиции\n")
			}
			if лит != "" {
				тест.Errorf("н4.2 ОШИБКА при проверке литеры в указанной позиции\n")
			}
		}
	}
	_Позитив()
	_Негатив()
}
