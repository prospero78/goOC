// модКомпилятор
package пакКомпилятор

/*
	Предоставляет тип компилятора для всего Оберона
*/

import (
	пакКонс "../пакКонсоль"
	пакСекции "./пакСекции"
	пакСканер "./пакСканер"
	пакФмт "fmt"
)

type ТуКомпилятор struct {
	сканер   *пакСканер.ТуСканер
	конс     *пакКонс.ТуКонсоль
	ИмяФайла string
	секции   *пакСекции.ТуСекции
}

func Новый() (компилятор *ТуКомпилятор, ош error) {
	пакКонс.Конс.Отладить("пакКомплиятор.Новый()")
	компилятор = new(ТуКомпилятор)
	if компилятор.сканер, ош = пакСканер.Новый(); ош != nil {
		ош = пакФмт.Errorf("ПакКомпилятор.Новый(): ошибка присоздании ТуСканер\n\t%v", ош)
		return компилятор, ош
	}
	if компилятор.секции, ош = пакСекции.Новый(); ош != nil {
		ош = пакФмт.Errorf("ПакКомпилятор.Новый(): ошибка присоздании ТуСканер\n\t%v", ош)
		return компилятор, ош
	}
	компилятор.конс = пакКонс.Конс
	return компилятор, ош
}

func (сам *ТуКомпилятор) Выполнить(пИмяФайла string) (ош error) {
	сам.конс.Отладить("ТуКомпилятор.Выполнить()")
	сам.ИмяФайла = пИмяФайла
	if _ош := сам.сканер.Обработать(пИмяФайла); _ош != nil {
		ош = пакФмт.Errorf("ТуКомпилятор.Выполнить(): ошибка при работе сканера\n\t%v", _ош)
	}
	return ош
}
