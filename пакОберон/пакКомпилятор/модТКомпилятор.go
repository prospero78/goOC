package пакКомпилятор

/*
	Предоставляет тип компилятора для всего Оберона
*/

import (
	мФмт "fmt"
	мИнт "github.com/prospero78/goOC/пакОберон/пакИнтерфейсы"
	мКонс "github.com/prospero78/goOC/пакОберон/пакКонсоль"
	мСек "github.com/prospero78/goOC/пакОберон/пакСекция"
)

//ТКомпилятор -- предоставляет тип компилятора
type ТКомпилятор struct {
	ИмяФайла мИнт.СИсхФайл
	модуль   мИнт.ИМодуль
	файлИмя  мИнт.СИсхФайл
}

//КомпиляторНов -- возвращает ссылку на новый ТКомпилятор
func КомпиляторНов() (компилятор *ТКомпилятор, ош error) {
	мКонс.Конс.Отладить("пакКомплиятор.Новый()")
	компилятор = &ТКомпилятор{}
	if компилятор == nil {
		return nil, мФмт.Errorf("КомпиляторНов(): нет памяти для компилятора?\n")
	}
	компилятор._КлючевыеСловаЗаполнить()
	if компилятор.модуль, ош = мСек.МодульНов(); ош != nil {
		return nil, мФмт.Errorf("КомпиляторНов(): ошибка при создании модуля\n\t%v", ош)
	}
	return компилятор, nil
}

//Обработать -- начинает обработку предоставляемого модуля
func (сам *ТКомпилятор) Обработать(пИмяФайла мИнт.СИсхФайл) (ош error) {
	мКонс.Конс.Отладить("ТКомпилятор.Обработать()")
	сам.файлИмя = пИмяФайла
	if ош = сам.модуль.Обработать(сам.файлИмя); ош != nil {
		return мФмт.Errorf("ТКомпилятор.Обработать(): ошибка при работе обработке модулей\n\t%v", ош)
	}
	return nil
}

func (сам *ТКомпилятор) _КлючевыеСловаЗаполнить() {
	//Заполняет список ключевых слов

}
