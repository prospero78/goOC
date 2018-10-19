package пакИмпорт

// модИмпорт

/*
	Модуль предоставляет тип для анализа секции импорта
*/

import (
	пакКонс "../../../пакКонсоль"
	пакСлово "../../пакСущность/пакСлово"
	пакСекция "../пакСекция"
	пакФмт "fmt"
)

//ТуИмпорт -- тип выделяети хранит слова импорта
type ТуИмпорт struct {
	_слова       *[]*пакСлово.ТуСлово
	_словаМодуля *[]*пакСлово.ТуСлово
	бИмпорт      bool // Указатель на то, что есть секция импорта
}

//Новый -- Создаёт и возвращает новый экземпляр для выделения слов импорта модуля
func Новый() (импорт *ТуИмпорт, ош error) {
	пакКонс.Конс.Отладить("пакСекции.пакИмпорт.Новый()")
	импорт = &ТуИмпорт{}
	return импорт, ош
}

//Обработать -- главная функция обработки секции импорт
func (сам *ТуИмпорт) Обработать(пСловаМодуля []*пакСлово.ТуСлово) (ош error) {
	пакКонс.Конс.Отладить("ТуИмпорт.Обработать()")
	сам._слова = &пСловаМодуля
	сам._словаМодуля = &пСловаМодуля
	сам._ЕслиИмпорт()
	if ош := сам._ЕслиИмпортОграничен(); ош != nil {
		ош = пакФмт.Errorf("ТуИмпорт.Обработать(): ошибка при проверке ограничения секции импорта\n\t%v", ош)
		return ош
	}
	return ош
}

func (сам *ТуИмпорт) _ЕслиИмпорт() {
	слово := (*сам._словаМодуля)[0]
	//Цикл по словарю слова IMPORT
	for _, слИмпорт := range пакСлово.КсИмпорт {
		//Проверка на все возможные значения
		if слово.Строка() == слИмпорт {
			сам.бИмпорт = true
			пакСекция.СловаОбрезать(сам)
		}
	}
}

func (сам *ТуИмпорт) _ЕслиИмпортОграничен() (ош error) {
	Дальше := func(слово *пакСлово.ТуСлово) bool {
		return слово.Строка() != ";"
	}
	цСловоНом := 0
	слово := (*сам._словаМодуля)[0]
	for Дальше(слово) && (цСловоНом < len(*сам._словаМодуля)) {
		цСловоНом++
		слово = (*сам._словаМодуля)[цСловоНом]
	}
	if Дальше(слово) {
		ош = пакФмт.Errorf("ТуИмпорт._ЕслиИмпортОграничен(): секция импорта ничем не ограничена")
		return ош
	}
	пакСекция.КонецУст(сам)
	return ош
}

//Слова -- возвращает слова модуля
func (сам *ТуИмпорт) Слова() (слова *[]*пакСлово.ТуСлово) {
	return сам._словаМодуля
}

// СловаИмпорт -- Возвращает слова импорта
func (сам *ТуИмпорт) СловаИмпорт() (слова *[]*пакСлово.ТуСлово) {
	return сам._слова
}

// СловаУст -- устанавлиает слова после обрезки секции слов
func (сам *ТуИмпорт) СловаУст(слова []*пакСлово.ТуСлово) {
	сам._слова = &слова
}
