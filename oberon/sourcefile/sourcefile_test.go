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
			if исхФайл.Размер() != 2593 { //Здесь проверяется количество рун
				тест.Errorf("п3.1 ОШИБКА при проверке размера текста исходника(2593), размер=[%v]\n", исхФайл.Размер())
			}
			if len(исхФайл.Исходник()) != 3400 { //Здесь проверяется количество байт
				тест.Errorf("п3.2 ОШИБКА при проверке размера текста исходника(3400), размер=[%v]\n", len(исхФайл.Исходник()))
			}
		}
		{ //4 Проверка литеры в указанной позиции
			тест.Logf("п4 Проверка литеры в указанной позиции\n")
			if лит := исхФайл.Лит(36); лит != "н" {
				тест.Errorf("п4.2 ОШИБКА при проверке литеры в позиции(36, н), лит=[%v]\n", лит)
			}
		}
		{ //5 Печать исходника
			тест.Logf("п5 Печать исходника\n")
			исхФайл.Печать()
		}
	}
	_Негатив := func() {
		_ПустоеИмя := func() {
			//1 Создание файла исходника с пустым именем
			тест.Logf("н1 Создание ИсходникФайл\n")
			defer func() {
				if паника := recover(); паника == nil {
					тест.Errorf("н1.1 ОШИБКА при генерации паники\n")
				}
			}()
			исхФайл = ИсхФайлНов("")
		}
		_ПлохойНомерПоз := func() {
			//3 Проверка литеры с несуществующей позицией
			тест.Logf("н3 Проверка литеры в указанной позиции\n")
			defer func() {
				if паника := recover(); паника == nil {
					тест.Errorf("н3.2 ОШИБКА при генерации паники\n")
				}
			}()
			if лит := исхФайл.Лит(8500); лит != "" {
				тест.Errorf("н3.2 ОШИБКА при проверке литеры в указанной позиции\n")
			}
		}
		_ОтрицНомерПоз := func() {
			//4 Проверка литеры с отрицательной позицией
			тест.Logf("н4 Проверка литеры в указанной позиции\n")
			defer func() {
				if паника := recover(); паника == nil {
					тест.Errorf("н4.2 ОШИБКА при генерации паники\n")
				}
			}()
			исхФайл.Лит(-5)
		}
		_ПустоеИмя()
		_ПлохойНомерПоз()
		_ОтрицНомерПоз()
		{ //2 Считывание файла исходника с пустым именем и отсутствующим на диске
			тест.Logf("н2 Получени размера\n")
			if исхФайл = ИсхФайлНов(стрФайл); исхФайл.Размер() == 0 {
				тест.Errorf("н2.2 ОШИБКА при проверке размера текста исходника\n")
			}
		}
	}
	_Позитив()
	_Негатив()
}