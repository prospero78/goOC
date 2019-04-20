package пакКомпилятор

/*
	Предоставляет тип компилятора для всего Оберона
*/

import (
	мАбс "../пакАбстракция"
	мКонс "../пакКонсоль"
	мМод "../пакМодуль"
	мФмт "fmt"
)

//ТКомпилятор -- предоставляет тип компилятора
type ТКомпилятор struct {
	ИмяФайла мАбс.СИсхФайл
	модуль   мАбс.АМодуль
	файлИмя  мАбс.СИсхФайл
}

//КомпиляторНов -- возвращает ссылку на новый ТКомпилятор
func КомпиляторНов() (компилятор *ТКомпилятор, ош error) {
	мКонс.Конс.Отладить("пакКомплиятор.Новый()")
	компилятор = &ТКомпилятор{}
	if компилятор == nil {
		return nil, мФмт.Errorf("КомпиляторНов(): нет памяти для компилятора?\n")
	}
	if компилятор.модуль, ош = мМод.МодульНов(); ош != nil {
		return nil, мФмт.Errorf("КомпиляторНов(): ошибка при создании модуля\n\t%v", ош)
	}
	return компилятор, nil
}

//Обработать -- начинает обработку предоставляемого модуля
func (сам *ТКомпилятор) Обработать(пИмяФайла мАбс.СИсхФайл) (ош error) {
	мКонс.Конс.Отладить("ТКомпилятор.Обработать()")
	сам.файлИмя = пИмяФайла
	if ош = сам.модуль.Обработать(сам.файлИмя); ош != nil {
		return мФмт.Errorf("ТКомпилятор.Обработать(): ошибка при работе обработке модулей\n\t%v", ош)
	}
	return nil
}
