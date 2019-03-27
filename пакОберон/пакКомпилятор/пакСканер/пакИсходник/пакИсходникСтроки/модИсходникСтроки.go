package пакИсходникСтроки

// модИсходникСтроки

/*
	Пакет предоставляет тип, который хранит построчно исходник.
	К строке можно обратиться по номеру, или полностью получить массив строк
*/
import (
	мКонс "../../../../пакКонсоль"
	//пакФмт "fmt"
	мСтр "strings"
)

//ТИсхСтрока -- специальный строковый тип для хранения итроки исходного кода
type ТИсхСтрока string

//ТИсхСтроки -- тип хранит список исходных строк
type ТИсхСтроки struct {
	СтрСписок []string
}

//Новый -- возвращает экземпляр тимпа для хранения строк исходника
func Новый() (исхСтр *ТИсхСтроки) {
	исхСтр = new(ТИсхСтроки)
	return исхСтр
}

//НаСтрокиРазбить -- разбивает на строки содержимое строки
func (сам *ТИсхСтроки) НаСтрокиРазбить(пИсх string) {
	мКонс.Конс.Отладить("ТИсхСтроки.НаСтроки_Разбить()")
	сам.СтрСписок = мСтр.Split(пИсх, "\n")

	for итер, стр := range сам.СтрСписок {
		if len(стр) > 1 {
			стр = стр[:len(стр)-1]
			сам.СтрСписок[итер] = стр
		}
		//сам.конс.Печать(пакФмт.Sprintf("%v: %v", итер, сам.СтрСписок[итер]))
	}
}

//Строка -- возвращает строку по указанному номеру
func (сам *ТИсхСтроки) Строка(пНомер int) (стр string) {
	if пНомер < 1 || пНомер > len(сам.СтрСписок) {
		panic("ТИсхСтроки.Строка(пНомер): пНомер за пределами разрешённого диапазона строк")
	}
	return сам.СтрСписок[пНомер-1]
}
