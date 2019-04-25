package пакСекция

// модСекция

/*
	Модуль предоставляет тип секции (общий для всех секций)
*/

import (
	мИнт "../пакИнтерфейсы"
	мФмт "fmt"
)

//ТСекция -- базовый тип для всех секций
type ТСекция struct {
	имя         мИнт.ССекцияИмя //заданное имя секции
	словаМодуля map[int]мИнт.ИСлово
	словаСекции map[int]мИнт.ИСлово
	стрТип      string //Хранит имя секциии (для модуля -- "МОДУЛЬ")
}

//СекцияНов -- создаёт и возвращает новый объект типа секции
func СекцияНов(пИмя string) (секция *ТСекция, ош error) {
	секция = &ТСекция{стрТип: пИмя}
	if секция == nil {
		return nil, мФмт.Errorf("СекцияНов(): нет памяти под новую секцию?\n")
	}
	секция.словаМодуля = make(map[int]мИнт.ИСлово)
	секция.словаСекции = make(map[int]мИнт.ИСлово)
	return секция, nil
}

//Имя -- возвращает имя секции
func (сам *ТСекция) Имя() мИнт.ССекцияИмя {
	return сам.имя
}

// СловаСекции -- возвращает список слов секции
func (сам *ТСекция) СловаСекции() map[int]мИнт.ИСлово {
	return сам.словаСекции
}

// СловаМодуля -- возвращает список слов модуля
func (сам *ТСекция) СловаМодуля() map[int]мИнт.ИСлово {
	return сам.словаМодуля
}

// СловаУст -- устанавливает список слов секции (модуль тоже своего рода секция)
func (сам *ТСекция) СловаУст(пСлова map[int]мИнт.ИСлово) {
	сам.словаМодуля = пСлова
}

// СловаСекцииПечать -- Печатает все слова секции
func (сам *ТСекция) СловаСекцииПечать() (ош error) {
	итер := 0
	стр := ""
	for _, слово := range сам.словаСекции {
		стрСлово := слово.Строка()
		с := мФмт.Sprintf("%v) %10.12v    ", слово.Номер(), стрСлово)
		стр = стр + с
		итер++
		if итер == 3 {
			стр += "\n"
			мФмт.Printf(стр)
			стр = ""
			итер = 0
		}
	}
	return nil
}
