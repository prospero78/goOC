// модСекции
package пакСекции

import (
	пакСлово "../пакСущность/пакСлово"
	пакИмпорт "./пакИмпорт"
	пакФмт "fmt"
)

type ТуСекции struct {
	импорт   *пакИмпорт.ТуИмпорт
	СлИмпорт []*пакСлово.ТуСлово
	СлМодуль []*пакСлово.ТуСлово
}

func Новый() (секции *ТуСекции, ош error) {
	секции = &ТуСекции{}
	if секции.импорт, ош = пакИмпорт.Новый(); ош != nil {
		ош = пакФмт.Errorf("ТуСекции.Новый(): ошибка при создании ТуИмпорт\n\t%v", ош)
		return секции, ош
	}
	return секции, ош
}

// Главный цикл обработки слов модуля -- разбитие на секции
func (сам *ТуСекции) Обработать(пСловаМодуля []*пакСлово.ТуСлово) (ош error) {
	if слИмпорт, слМодуль, ош := сам.импорт.Обработать(пСловаМодуля); ош != nil {
		return ош
	} else {
		сам.СлИмпорт = make([]*пакСлово.ТуСлово, len(слИмпорт))
		copy(слИмпорт, сам.СлИмпорт)

		сам.СлМодуль = make([]*пакСлово.ТуСлово, len(слМодуль))
		copy(слМодуль, сам.СлМодуль)
	}
	return ош
}
