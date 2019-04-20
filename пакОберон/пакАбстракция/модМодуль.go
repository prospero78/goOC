package пакАбстракция

/*
	Модуль предоставляет абстракцию модуля.
*/

import (
	мФмт "fmt"
)

//СМодуль -- специальный строковый тип для хранения имени модуля
type СМодуль string

//АМодуль -- абстракция интерфейса
type АМодуль interface {
	Имя() СМодуль
	Обработать(СИсхФайл) error
}

//СловаОбрезать -- обрезает слова в секции
func СловаОбрезать(пСлова map[int]АСлово) (map[int]АСлово, error) {
	if len(пСлова) == 0 {
		return nil, мФмт.Errorf("СловаОбрезать(): справочник слов не может быть пустым\n")
	}
	for адр := 1; адр < len(пСлова); адр++ {
		пСлова[адр-1] = пСлова[адр]
	}
	delete(пСлова, len(пСлова)-1)
	return пСлова, nil
}
