package пакИсхСтрокаЗначение

/*
	Модуль предоставляет тип для безопасного хранения
	строки исходного кода
*/

import (
	мФмт "fmt"
)

//ТИсхСтрокаЗначениеСтр -- специальный строковый тип для хранения строки исходника
type ТИсхСтрокаЗначениеСтр string

//ТИсхСтрокаЗначение -- тип для хранения строки исходного кода
type ТИсхСтрокаЗначение struct {
	знач ТИсхСтрокаЗначениеСтр
}

//ИсхСтрокаЗначениеНов -- возвращает ссылку на новый ТИсхСтрокаЗначение
func ИсхСтрокаЗначениеНов(пСтр ТИсхСтрокаЗначениеСтр) (стр *ТИсхСтрокаЗначение, ош error) {
	стр = &ТИсхСтрокаЗначение{}
	if стр == nil {
		return nil, мФмт.Errorf("ИсхСтрокаЗначНов(): нет памяти для типа строки исходника?\n")
	}
	стр.знач = пСтр
	return стр, nil
}

//Знач -- возвращает хранимое значение строки исходника ТИсхСтрокаЗначение
func (сам *ТИсхСтрокаЗначение) Знач() ТИсхСтрокаЗначениеСтр {
	return сам.знач
}
