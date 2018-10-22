package пакСекция

// модСекция

/*
	Модуль предоставляет тип секции (общий для всех секций)
*/

import (
	пакСлово "../../пакСущность/пакСлово"
	пакФмт "fmt"
)

//ТуСекция -- базовый тип для всех секций
type ТуСекция struct {
	слова  []*пакСлово.ТуСлово
	стрТип string //Хранит имя секциии (для модуля -- "МОДУЛЬ")
}

//Новый -- создаёт и возвращает новый объект типа секции
func Новый(пИмя string) (секция *ТуСекция) {
	секция = &ТуСекция{стрТип: пИмя}
	return секция
}

// Тип -- возвращает тип секции
func (сам *ТуСекция) Тип() string {
	return сам.стрТип
}

// Слова -- возвращает список слов секции (модуль тоже своего рода секция)
func (сам *ТуСекция) Слова() []*пакСлово.ТуСлово {
	return сам.слова
}

// СловаУст -- устанавливает список слов секции (модуль тоже своего рода секция)
func (сам *ТуСекция) СловаУст(пСлова []*пакСлово.ТуСлово) {
	сам.слова = пСлова
}

// СловаОбрезать -- обрезает слова в секции
func (сам *ТуСекция) СловаОбрезать() {
	слова := сам.слова
	слова = слова[1:]
	сам.слова = слова
}

// СловаПечать -- Печатает все слова секции
func (сам *ТуСекция) СловаПечать() (ош error) {
	итер := 0
	стр := ""
	for _, слово := range сам.слова {
		if стрСлово, ош := слово.Строка(); ош == nil {
			с := пакФмт.Sprintf("%v) %10.12v    ", слово.Номер(), стрСлово)
			стр = стр + с
			итер++
			if итер == 3 {
				стр += "\n"
				пакФмт.Printf(стр)
				стр = ""
				итер = 0
			}
		} else {
			ош = пакФмт.Errorf("пакСекция.СловаПечать(): ошибка при получении значения слова\n\t%v", ош)
			return ош
		}

	}
	return nil
}

//КонецУст -- Устанавливает слово-маркер -- конец секции
func (сам *ТуСекция) КонецУст() {

}

//СловаСекцииРазделить -- разделяет слова своей секции и слова модуля (оставшиеся)
func (сам *ТуСекция) СловаСекцииРазделить() {

}
