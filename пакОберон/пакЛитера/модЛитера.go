package пакЛитера

/*
	Модуль предоставляет тип для работы с отдельной литерой.
*/

import (
	мИнт "../пакИнтерфейсы"
	мФмт "fmt"
	мСтр "strings"
)

//ТЛит -- тип для работы с отдельной литерой
type ТЛит struct {
	лит мИнт.СЛит
}

const (
	//наборы букв для перебора
	стрБуквыРус = "абвгдеёжзийклмнопрстуфхцчшщьыъэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЬЫЪЭЮЯ"
	стрБуквыАнг = "abcdefghjiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	стрБуквыВсе = стрБуквыРус + стрБуквыАнг

	//CтрЦифры -- список цифр, что могут встречаться в числах
	стрЦифры = "0123456789."
)

//ЛитераНов -- возвращает ссылку на новый ТЛит
func ЛитераНов() (лит *ТЛит, ош error) {
	лит = &ТЛит{}
	if лит == nil {
		return nil, мФмт.Errorf("ЛитераНов(): нет памяти для литеры?\n")
	}
	return лит, nil
}

// ЕслиБуква -- проверяет наличие буквы в литере
func (сам *ТЛит) ЕслиБуква() bool {
	if мСтр.Contains(стрБуквыВсе, string(сам.лит)) {
		return true
	}
	return false
}

//ЕслиЦифра -- проверяет, что литера цифра
func (сам *ТЛит) ЕслиЦифра() bool {
	if мСтр.Contains(стрЦифры, string(сам.лит)) {
		return true
	}
	return false
}

//Уст -- устанавливает хранимую литеру
func (сам *ТЛит) Уст(пЛит мИнт.СЛит) error {
	if пЛит == "" {
		return мФмт.Errorf("ТЛит.Уст(): пЛит не может быть пустой\n")
	}
	сам.лит = пЛит
	return nil
}

//Лит -- возвращает хранимую литеру
func (сам *ТЛит) Лит() мИнт.СЛит {
	return сам.лит
}
