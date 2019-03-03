package пакКомпилятор

/*
	Предоставляет тип компилятора для всего Оберона
*/

import (
	мКонс "../пакКонсоль"
	мСекции "./пакСекции"
	мСканер "./пакСканер"
	мФмт "fmt"
)

//ТКомпилятор -- предоставляет тип компилятора
type ТКомпилятор struct {
	сканер   *мСканер.ТСканер
	ИмяФайла string
	секции   *мСекции.ТСекции
}

//КомпиляторНов -- возвращает ссылку на новый ТКомпилятор
func КомпиляторНов() (компилятор *ТКомпилятор, ош error) {
	мКонс.Конс.Отладить("пакКомплиятор.Новый()")
	компилятор = &ТКомпилятор{}
	if компилятор.сканер, ош = мСканер.Новый(); ош != nil {
		return nil, мФмт.Errorf("ПакКомпилятор.Новый(): ошибка присоздании ТуСканер\n\t%v", ош)
	}
	if компилятор.секции, ош = мСекции.Новый(); ош != nil {
		return nil, мФмт.Errorf("ПакКомпилятор.Новый(): ошибка присоздании ТуСканер\n\t%v", ош)
	}
	return компилятор, nil
}

//Обработать -- начинает обработку предоставляемого модуля
func (сам *ТКомпилятор) Обработать(пИмяФайла string) (ош error) {
	мКонс.Конс.Отладить("ТКомпилятор.Обработать()")
	сам.ИмяФайла = пИмяФайла
	if ош = сам.сканер.Обработать(пИмяФайла); ош != nil {
		return мФмт.Errorf("ТКомпилятор.Обработать(): ошибка при работе сканера\n\t%v", ош)
	}
	if ош = сам.секции.Обработать(сам.сканер.Исх.СловаМодуля); ош != nil {
		return мФмт.Errorf("ТКомпилятор.Обработать(): ошибка при работе разбиения секций\n\t%v", ош)
	}
	return nil
}
