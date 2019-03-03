package пакМодуль

/*
	Предоставляет возможность сканеру секций определеить правильно ли
	описан модуль.
*/

import (
	мКонс "../../../пакКонсоль"
	мСлово "../../пакСущность/пакСлово"
	мСекция "../пакСекция"
	мФмт "fmt"
)

//ТМодуль -- предоставляет тип, содержащий слова модуля
type ТМодуль struct {
	секция *мСекция.ТСекция
	имя    *мСлово.ТСлово //Имя модуля из слов модуля
}

//Новый -- возвращает новый экземпляр типа для выделения слов модуля
func Новый() (модуль *ТМодуль, ош error) {
	мКонс.Конс.Отладить("пакСекции.пакМодуль.Новый()")
	секция := мСекция.Новый("МОДУЛЬ")
	модуль = &ТМодуль{секция: секция}
	return модуль, ош
}

//Обработать -- обеспечивает выделение слов модуля
func (сам *ТМодуль) Обработать(пСловаМодуля []*мСлово.ТСлово) (ош error) {

	мКонс.Конс.Отладить("ТМодуль.Обработать()")
	сам.секция.СловаУст(пСловаМодуля)
	мФмт.Printf("Всего слов в модуле после отбрасывания комментариев1 > [%v]\n", len(сам.секция.Слова()))

	if ош := сам._МодульНачалоПроверить(); ош != nil {
		ош = мФмт.Errorf("ТМодуль.Обработать(): ошибка при поиске начала модуля\n\t%v", ош)
		return ош
	}
	if ош := сам._ИмяПроверить(); ош != nil {
		ош = мФмт.Errorf("ТМодуль.Обработать(): ошибка при поиске имени модуля\n\t%v", ош)
		return ош
	}
	if ош := сам._РазделительПроверить(); ош != nil {
		ош = мФмт.Errorf("ТМодуль.Обработать(): ошибка при поиске разделителя заголовка модуля\n\t%v", ош)
		return ош
	}
	if ош := сам._МодульКонецНайти(); ош != nil {
		ош = мФмт.Errorf("ТМодуль.Обработать(): ошибка при поиске конца модуля\n\t%v", ош)
		return ош
	}
	if ош := сам._МодульОдинПроверить(); ош != nil {
		ош = мФмт.Errorf("ТМодуль.Обработать(): ошибка при поиске единственного MODULE\n\t%v", ош)
		return ош
	}
	мФмт.Printf("Всего слов в модуле после обработки > [%v]\n", len(сам.секция.Слова()))
	return ош
}

// Проверяет имя модуля в тексте
func (сам *ТМодуль) _ИмяПроверить() (ош error) {
	мКонс.Конс.Отладить("ТМодуль._ИмяПроверить()")
	слово := сам.секция.Слова()[0]
	//мКонс.Конс.Отладить("Проверка имени модуля: \"" + слово.Строка() + "\"")
	if слово.ЕслиИмяСтрого() {
		сам.имя = слово
	} else {
		стрСтрока, _ := слово.Строка()
		стрИсх, _ := слово.СтрИсх()
		ош = мФмт.Errorf("ТуМодуль._ИмяПроверить(): такое имя модуля запрещено\n\tимя=%v\n\tСтрока=%v", стрСтрока, стрИсх)
		return ош
	}
	сам.секция.СловаОбрезать()

	return ош
}

// Проверяет разделитель после имени модуля в начале
func (сам *ТМодуль) _РазделительПроверить() (ош error) {
	мКонс.Конс.Отладить("ТуМодуль._РазделительПроверить()")
	слово := сам.секция.Слова()[0]
	if стрСтрока, ош := слово.Строка(); ош == nil {
		if стрСтрока != ";" {
			стрИсх, _ := слово.СтрИсх()
			ош = мФмт.Errorf("ТуМодуль._РазделительПроверить(): ошибка в окончании названия модуля\n\t%v %v",
				стрСтрока, стрИсх)
			return ош
		}
	} else {
		стрИсх, _ := слово.СтрИсх()
		ош = мФмт.Errorf("ТуМодуль._РазделительПроверить(): ошибка при проверке разделителя\n\t%v %v",
			стрСтрока, стрИсх)
		return ош
	}
	сам.секция.СловаОбрезать()
	return nil
}

func (сам *ТМодуль) _МодульКонецНайти() (ош error) {
	мКонс.Конс.Отладить("ТуМодуль._МодульКонецНайти()")
	цСчётОбр := len(сам.секция.Слова()) - 1
	for цСчётОбр >= 0 {
		словоТочка := сам.секция.Слова()[цСчётОбр]
		// Нашли конечную точку?
		стрТочка, _ := словоТочка.Строка()
		if стрТочка == "." {
			// Попытка найти END. КсМодуль содержит множество слов
			словоКонец := сам.секция.Слова()[цСчётОбр-2]
			for _, строкаКонец := range мСлово.КсКонец {
				стрКонец, _ := словоКонец.Строка()
				if стрКонец == строкаКонец {
					// Попытка проверить совпадение имя модуля и конца модуля
					словоИмя := сам.секция.Слова()[цСчётОбр-1]
					стрИмя, _ := сам.имя.Строка()
					стрИмя2, _ := словоИмя.Строка()
					if стрИмя2 == стрИмя {
						сам._ХвостОтбросить(цСчётОбр)
						return
					}
					// Это гарантированная ошибка, так как не было возврата
					ош = мФмт.Errorf("ТуМодуль._КонецМодульНайти(): имя модуля в начале и в конце модуля не совпадают\n\t%v %v",
						стрИмя, стрИмя2)
					return ош
				}
			}
		}
		цСчётОбр--
	}
	if цСчётОбр == 0 { // Гарантированная ошибка. Так быть не может.
		ош = мФмт.Errorf("ТуМодуль.__КонецМодуль_Найти(): нет завершающей точки в модуле")
		return ош
	}
	return ош
}

// Отбросить всё, что за END <name_module>.
func (сам *ТМодуль) _ХвостОтбросить(пСчётОбр int) {
	//мКонс.Конс.Отладить("Хвост отбросить")
	индекс := 0
	var слова []*мСлово.ТСлово
	for индекс < (пСчётОбр - 2) {
		слова = append(слова, сам.секция.Слова()[индекс])
		индекс++
	}
	сам.секция.СловаУст(слова)
	//мКонс.Конс.Отладить("Последнее слово перед концом модуля:" + мФмт.Sprintf("%v", пСчётОбр))
}

// Проверяет, что MODULE реально один в модуле, первый уже отброшен
func (сам *ТМодуль) _МодульОдинПроверить() (ош error) {
	мКонс.Конс.Отладить("ТуМодуль._МодульОдинПроверить()")
	for индекс := range сам.секция.Слова() {
		слово := сам.секция.Слова()[индекс]
		for индекс2 := range мСлово.КсМодуль {
			стрМодуль, _ := слово.Строка()
			if стрМодуль == мСлово.КсМодуль[индекс2] {
				ош = мФмт.Errorf("ТуМодуль._МодульОдинПроверить(): MODULE встречается больше одного раза")
				return ош
			}
		}
	}
	return ош
}

// Проверяет, что модуль начинается правильно
func (сам *ТМодуль) _МодульНачалоПроверить() (ош error) {
	мКонс.Конс.Отладить("ТуМодуль.__МодульНачало_Проверить()")
	бРез := false
	слово := сам.секция.Слова()[0]
	for индекс := range мСлово.КсМодуль {
		стрМодуль, _ := слово.Строка()
		if стрМодуль == мСлово.КсМодуль[индекс] {
			бРез = true
		}
	}
	if !бРез {
		ош = мФмт.Errorf("ТуМодуль._МодульНачалоПроверить(): модуль не начинается с MODULE")
	}
	сам.секция.СловаОбрезать()
	return ош
}

// СловаСекцииУст -- устанавливает слова после обрезки секции модуля (это и есть глобальная секция)
func (сам *ТМодуль) СловаСекцииУст(слова []*мСлово.ТСлово) {
	сам.секция.СловаУст(слова)
}

// СловаМодуля -- возвращает список слов модуля
func (сам *ТМодуль) СловаМодуля() (слова []*мСлово.ТСлово) {
	return сам.секция.Слова()
}
