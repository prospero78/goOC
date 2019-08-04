package пакЛитера

/*
	Модуль предоставляет тест для интерфейса литеры
*/

import (
	мТест "testing"
)

var (
	лит ИЛит
	ош  error
)

func TestЛит(тест *мТест.T) {
	_Позитив := func() {
		{ //1 Создание литеры
			тест.Logf("п1 Сощдание литеры\n")
			if лит, ош = ЛитераНов(); ош != nil {
				тест.Errorf("п1.1 ОШИБКА при создании литеры\n\t%v", ош)
			}
			if лит == nil {
				тест.Errorf("п1.2 ОШИБКА лит не должна быть nil\n")
			}
		}
		{ //2 Проверка начальных значений
			тест.Logf("п2 Проверка начальных значений\n")
			if лит.Получ() != "" {
				тест.Errorf("п2.1 ОШИБКА при хранении пустой литеры, пЛит=[%v]\n", лит)
			}
			if лит.ЕслиЦифра() {
				тест.Errorf("п2.2 ОШИБКА при хранении пустой литеры\n")
			}
			if лит.ЕслиСпецЛит() {
				тест.Errorf("п2.2 ОШИБКА при хранении пустой литеры\n")
			}
			if лит.ЕслиБуква() {
				тест.Errorf("п2.3 ОШИБКА при хранении пустой литеры\n")
			}
		}
		{ //3 Установка буквы
			тест.Logf("п3 Установка буквы\n")
			if ош = лит.Уст("Б"); ош != nil {
				тест.Errorf("п3.1 ОШИБКА в установке литеры\n\t%v", ош)
			}
			if лит.ЕслиЦифра() {
				тест.Errorf("п3.2 ОШИБКА при хранении буквы\n")
			}
			if лит.ЕслиСпецЛит() {
				тест.Errorf("п3.3 ОШИБКА при хранении буквы\n")
			}
			if !лит.ЕслиБуква() {
				тест.Errorf("п3.4 ОШИБКА при хранении буквы \"Б\", знач=[%v]\n", лит)
			}
		}
		{ //4 Установка цифры
			тест.Logf("п4 Установка цифры\n")
			if ош = лит.Уст("8"); ош != nil {
				тест.Errorf("п4.1 ОШИБКА в установке литеры\n\t%v", ош)
			}
			if !лит.ЕслиЦифра() {
				тест.Errorf("п4.2 ОШИБКА при хранении цифры \"8\", знач=[%v]\n", лит)
			}
			if лит.ЕслиСпецЛит() {
				тест.Errorf("п4.3 ОШИБКА при хранении цифры\n")
			}
			if лит.ЕслиБуква() {
				тест.Errorf("п4.4 ОШИБКА при хранении цифры\n")
			}
		}
		{ //5 Установка спецлитеры
			тест.Logf("п5 Установка спецлитеры\n")
			if ош = лит.Уст("_"); ош != nil {
				тест.Errorf("п5.1 ОШИБКА в установке литеры\n\t%v", ош)
			}
			if лит.ЕслиЦифра() {
				тест.Errorf("п5.2 ОШИБКА при хранении спецлитеры\n")
			}
			if !лит.ЕслиСпецЛит() {
				тест.Errorf("п5.3 ОШИБКА при хранении спецлитеры\"_\", знач=[%v]\n", лит)
			}
			if лит.ЕслиБуква() {
				тест.Errorf("п5.4 ОШИБКА при хранении цифры\n")
			}
		}
		{ //6 Проверка строкового представления
			тест.Logf("п6 Проверка строкового представления\n")
			if лит.String() != "_" {
				тест.Errorf("п6 ОШИБКА в строковом представлении, знач=[%v]\n", лит)
			}
		}
	}
	_Негатив := func() {
		{ //1 Создание литеры
			тест.Logf("н1 Сощдание литеры\n")
			if лит, ош = ЛитераНов(); ош != nil {
				тест.Errorf("н1.1 ОШИБКА при создании литеры\n\t%v", ош)
			}
			if лит == nil {
				тест.Errorf("н1.2 ОШИБКА лит не должна быть nil\n")
			}
		}
		{ //2  Пустое присвоение
			тест.Logf("н2 Пустое присвоение\n")
			if ош = лит.Уст(""); ош == nil {
				тест.Errorf("н2.1 ОШИБКА в установке литеры\n\t%v", ош)
			}
			if лит.ЕслиЦифра() {
				тест.Errorf("н2.2 ОШИБКА при хранении буквы\n")
			}
			if лит.ЕслиСпецЛит() {
				тест.Errorf("н2.3 ОШИБКА при хранении буквы\n")
			}
			if лит.ЕслиБуква() {
				тест.Errorf("н2.4 ОШИБКА при хранении буквы\n")
			}
		}
		{ //3  Запрещённое присвоение
			тест.Logf("н3 Запрещённое присвоение\n")
			if ош = лит.Уст("\t"); ош == nil {
				тест.Errorf("н3.1 ОШИБКА в установке литеры\n\t%v", ош)
			}
			if лит.ЕслиЦифра() {
				тест.Errorf("н3.2 ОШИБКА при хранении буквы\n")
			}
			if лит.ЕслиСпецЛит() {
				тест.Errorf("н3.3 ОШИБКА при хранении буквы\n")
			}
			if лит.ЕслиБуква() {
				тест.Errorf("н3.4 ОШИБКА при хранении буквы \n")
			}
		}
	}
	_Позитив()
	_Негатив()
}
