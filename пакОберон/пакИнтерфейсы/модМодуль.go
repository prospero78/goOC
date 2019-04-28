package пакИнтерфейсы

/*
	Модуль предоставляет интерфейс модуля.
*/

import (
	мФмт "fmt"
)

//СМодульИмя -- специальный строковый тип для хранения имени модуля
type СМодульИмя string

//ИМодуль -- интерфейс интерфейса
type ИМодуль interface {
	ИСекция
	МодульИмя() СМодульИмя
	Обработать(СИсхФайл) error
}

//СловаМодуляОбрезать -- обрезает слова в секции
func СловаМодуляОбрезать(пСлова map[ССловоНомерМодуль]ИСлово) (map[ССловоНомерМодуль]ИСлово, error) {
	if len(пСлова) == 0 {
		return nil, мФмт.Errorf("СловаОбрезать(): справочник слов не может быть пустым\n")
	}
	адр := ССловоНомерМодуль(1)
	for адр = 1; int(адр) < len(пСлова); адр++ {
		пСлова[адр-1] = пСлова[адр]
	}
	delete(пСлова, ССловоНомерМодуль(len(пСлова))-1)
	return пСлова, nil
}
