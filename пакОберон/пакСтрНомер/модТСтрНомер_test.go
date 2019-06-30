package пакСтрНомер

/*
	Модуль предоставляет тест для номера строки исходника
*/

import (
	мТест "testing"
)

var (
	номер  ИСтрНомер
	номИзм ИСтрНомерИзм
	ош     error
	ок     bool
)

func TestНомерСтр(тест *мТест.T) {
	_Позитив := func() {
		{ //1 Создание номера строки
			тест.Logf("п1 Создание номера строки\n")
			if номер, ош = СтрНомерНов(); ош != nil {
				тест.Errorf("п1.1 ОШИБКА при создании номера строки\n\t%v", ош)
			}
			if номер == nil {
				тест.Errorf("п1.2 ОШИБКА номер строки не может быть nil\n")
			}
		}
		{ //2
			тест.Logf("п2 Проверка начальных значений\n")
			if номер.Знач() != 0 {
				тест.Errorf("п2.1 ОШИБКА номер строки должен быть 0, знач=[%v]\n", номер.Знач())
			}
			if номер.String() != "0" {
				тест.Errorf("п2.2 ОШИБКА строковое представление номера строки должен быть (0), стр=[%v]\n", номер)
			}
		}
		{ //3
			тест.Logf("п3 Установка значения номера строки\n")
			if ош = номер.Уст(11); ош != nil {
				тест.Errorf("п3.1 ОШИБКА при установке номера строки\n\t%v", ош)
			}
			if номер.Знач() != 11 {
				тест.Errorf("п3.2 ОШИБКА номер строки должен быть 11, знач=[%v]\n", номер)
			}
			if номер.String() != "11" {
				тест.Errorf("п3.3 ОШИБКА строковое представление номера строки должен быть (11), стр=[%v]\n", номер)
			}
		}
		{ //4
			тест.Logf("п4 Повторнач установка значения номера строки\n")
			if ош = номер.Уст(21); ош == nil {
				тест.Errorf("п4.1 ОШИБКА при установке номера строки\n\t%v", ош)
			}
			if номер.Знач() != 11 {
				тест.Errorf("п4.2 ОШИБКА номер строки должен быть 11, знач=[%v]\n", номер)
			}
			if номер.String() != "11" {
				тест.Errorf("п4.3 ОШИБКА строковое представление номера строки должен быть (11), стр=[%v]\n", номер)
			}
		}
		{ //5
			тест.Logf("п5 Преобразование типа к изменяемому\n")
			if номИзм, ок = номер.(ИСтрНомерИзм); !ок {
				тест.Errorf("п5.1 ОШИБКА при приведении фиксированного номера строки к изменяемому\n")
			}
			тест.Logf("н3 результат кастинга, ок=[%v]\n", ок)
			номИзм.Доб()
			if номИзм.Знач() != 12 {
				тест.Errorf("п5.2 ОШИБКА номер строки должен быть 12, знач=[%v]\n", номИзм)
			}
			if номИзм.String() != "12" {
				тест.Errorf("п5.3 ОШИБКА строковое представление номера строки должен быть (12), стр=[%v]\n", номИзм)
			}
			тест.Logf("п5а Проверка значения исходного типа\n")
			if номер.Знач() != 12 {
				тест.Errorf("п5.4 ОШИБКА номер строки должен быть 12, знач=[%v]\n", номер)
			}
			if номер.String() != "12" {
				тест.Errorf("п5.5 ОШИБКА строковое представление номера строки должен быть (12), стр=[%v]\n", номер)
			}
		}
		{ //6
			тест.Logf("п6 Сброс номера строки\n")
			номИзм.Сброс()
			if номер.Знач() != 1 {
				тест.Errorf("п3.2 ОШИБКА номер строки должен быть 1, знач=[%v]\n", номер)
			}
			if номер.String() != "1" {
				тест.Errorf("п3.3 ОШИБКА строковое представление номера строки должен быть (1), стр=[%v]\n", номер)
			}
		}
	}
	_Негатив := func() {
		{ //1 Создание номера строки
			тест.Logf("н1 Создание номера строки\n")
			if номер, ош = СтрНомерНов(); ош != nil {
				тест.Errorf("н1.1 ОШИБКА при создании номера строки\n\t%v", ош)
			}
			if номер == nil {
				тест.Errorf("н1.2 ОШИБКА номер строки не может быть nil\n")
			}
		}
		{ //2
			тест.Logf("н2 Установка нулевого значения номера строки\n")
			if ош = номер.Уст(0); ош == nil {
				тест.Errorf("н2.1 ОШИБКА при установке номера строки\n")
			}
			if номер.Знач() != 0 {
				тест.Errorf("н2.2 ОШИБКА номер строки должен быть 0, знач=[%v]\n", номер.Знач())
			}
			if номер.String() != "0" {
				тест.Errorf("н2.3 ОШИБКА строковое представление номера строки должен быть (0), стр=[%v]\n", номер)
			}
		}
	}
	_Позитив()
	_Негатив()
}
