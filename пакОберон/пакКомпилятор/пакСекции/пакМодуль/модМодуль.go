// модМодуль
package пакМодуль

/*
Предоставляет возможность сканеру секций определеить правильно ли
описан модуль.
*/

import (
	пакКонс "../../../пакКонсоль"
	пакСлово "../../пакСущность/пакСлово"
	пакСекция "../пакСекция"
	пакФмт "fmt"
)

type ТуМодуль struct {
	_слова []*пакСлово.ТуСлово // Сохраняет все слова модуля
	_имя   *пакСлово.ТуСлово   // Сохраняет имя модуля
	_бИмя  bool                // Сохраняет состояние присовенности имени
}

func Новый() (модуль *ТуМодуль, ош error) {
	пакКонс.Конс.Отладить("пакСекции.пакМодуль.Новый()")
	модуль = &ТуМодуль{}
	return модуль, ош
}

func (сам *ТуМодуль) Обработать(пСловаМодуля []*пакСлово.ТуСлово) (ош error) {
	сам._слова = пСловаМодуля
	пакКонс.Конс.Отладить("ТуМодуль.Обработать()")
	return ош
}

// Проверяет имя модуля в тексте
func (сам *ТуМодуль) __Имя_Проверить() (ош error) {
	слово := сам._слова[0]
	if !слово.ЕслиИмя_Строго() {
		ош = пакФмт.Errorf("ТуМодуль.__Имя_Проверить(): такое имя модуля запрещено\n\tимя=%v,\n\t%v", слово.Строка(), слово.СтрИсх())
		return ош
	} else {
		сам._имя = слово
	}
	пакСекция.Слова_Обрезать(сам)
	сам._бИмя = true

	return ош
}

// Проверяет разделитель после имени модуля в начале
func (сам *ТуМодуль) __Разделитель_Проверить() (ош error) {
	слово := сам._слова[0]
	if слово.Строка() != ";" {
		ош = пакФмт.Errorf("ТуМодуль.__Разделитель_Проверить(): ошибка в окончании названия модуля\n\t%v %v",
			слово.Строка(), слово.СтрИсх())
		return ош
	}
	пакСекция.Слова_Обрезать(сам)
	return ош
}

func (сам *ТуМодуль) __КонецМодуль_Найти() (ош error) {
	цСчётОбр := 0
	for цСчётОбр = len(сам._слова); цСчётОбр >= 0; цСчётОбр-- {
		слово_точка := сам._слова[цСчётОбр]
		if слово_точка.Строка() == "." {
			слово_конец := сам._слова[цСчётОбр-2]
			for _, строка := range пакСлово.КсМодуль {
				if слово_конец.Строка() == строка {
					// Проверить на совадение имя модуля между ними
					слово_имя := сам._слова[цСчётОбр-1]
					if слово_имя.Строка() == сам._имя.Строка() {
						return ош
					} else {
						ош = пакФмт.Errorf("ТуМодуль.__КонецМодуль_Найти(): имя моудля в начале и в конце модуля не совпадают\n\t%v %v",
							сам._имя.Строка(), слово_имя.Строка())
					}
				}
			}

		}
	}
	if цСчётОбр == 0 {
		ош = пакФмт.Errorf("ТуМодуль.__КонецМодуль_Найти(): нет завершающей точки в модуле")
		return ош
	}
	return ош
}

// Устанавлиает слова после обрезки секции слов
func (сам *ТуМодуль) Слова_Уст(слова []*пакСлово.ТуСлово) {
	сам._слова = слова
}

// Возвращает список слов
func (сам *ТуМодуль) Слова() (слова *[]*пакСлово.ТуСлово) {
	return &сам._слова
}

func (сам *ТуМодуль) Секция() {
	// Ничего не делает, просто введена для совместимости. Проверка интерфейса.
}
