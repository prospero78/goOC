package пакМодуль

/*
	Предоставляет возможность сканеру секций определеить правильно ли
	описан модуль.
*/

import (
	мИнт "../пакИнтерфейсы"
	мКонс "../пакКонсоль"
	мСкан "../пакСканер"
	мСлово "../пакСлово"
	мФмт "fmt"
)

//ТМодуль -- предоставляет тип, содержащий слова модуля
type ТМодуль struct {
	имя    мИнт.СМодуль        //Имя модуля
	слово  мИнт.ИСлово         // слово в котором хранится имя модуля
	слова  map[int]мИнт.ИСлово //Справочник всех слов модуля
	сканер мИнт.ИСканер
}

//МодульНов -- возвращает на новый ТМодуль
func МодульНов() (модуль *ТМодуль, ош error) {
	мКонс.Конс.Отладить("МодульНов()")
	модуль = &ТМодуль{}
	if модуль == nil {
		return nil, мФмт.Errorf("МодульНов(): нет памяти на модуль?\n")
	}
	модуль.слова = make(map[int]мИнт.ИСлово)
	if модуль.сканер, ош = мСкан.СканерНов(); ош != nil {
		return nil, мФмт.Errorf("МодульНов(): ошибка при создании сканера модуля\n\t%v", ош)
	}
	return модуль, nil
}

//Обработать -- обеспечивает выделение слов модуля
func (сам *ТМодуль) Обработать(пФайлИмя мИнт.СИсхФайл) (ош error) {
	мКонс.Конс.Отладить("ТМодуль.Обработать()")
	if ош = сам.сканер.Обработать(пФайлИмя); ош != nil {
		return мФмт.Errorf("ТМодуль.Обработать(): ошибка при работе сканера\n\t%v", ош)
	}
	сам.слова = сам.сканер.Слова()
	мФмт.Printf("Всего слов в модуле [%v]\n", len(сам.слова))

	if ош = сам._МодульНачалоПроверить(); ош != nil {
		return мФмт.Errorf("ТМодуль.Обработать(): ошибка при поиске начала модуля\n\t%v", ош)
	}
	if ош = сам._ИмяПроверить(); ош != nil {
		return мФмт.Errorf("ТМодуль.Обработать(): ошибка при поиске имени модуля\n\t%v", ош)
	}
	if ош = сам._РазделительПроверить(); ош != nil {
		return мФмт.Errorf("ТМодуль.Обработать(): ошибка при поиске разделителя заголовка модуля\n\t%v", ош)
	}
	if ош = сам._МодульКонецНайти(); ош != nil {
		return мФмт.Errorf("ТМодуль.Обработать(): ошибка при поиске конца модуля\n\t%v", ош)
	}
	if ош = сам._МодульОдинПроверить(); ош != nil {
		return мФмт.Errorf("ТМодуль.Обработать(): ошибка при поиске единственного MODULE\n\t%v", ош)
	}
	мФмт.Printf("Всего слов в модуле после обработки > [%v]\n", len(сам.слова))
	return nil
}

// Проверяет имя модуля в тексте
func (сам *ТМодуль) _ИмяПроверить() (ош error) {
	мКонс.Конс.Отладить("ТМодуль._ИмяПроверить()")
	слово := сам.слова[0]
	//мКонс.Конс.Отладить("Проверка имени модуля: \"" + слово.Строка() + "\"")
	if слово.ЕслиИмяСтрого() {
		сам.имя = мИнт.СМодуль(слово.Слово())
	} else {
		стрСлово := слово.Слово()
		стрИсх := слово.Строка()
		return мФмт.Errorf("ТуМодуль._ИмяПроверить(): такое имя модуля запрещено, имя=[%v], строка=[%v]\n", стрСлово, стрИсх)
	}
	if сам.слова, ош = мИнт.СловаОбрезать(сам.слова); ош != nil {
		return мФмт.Errorf("ТуМодуль._ИмяПроверить(): ошибка при обрезке слов модуля\n\t%v", ош)
	}

	return nil
}

// Проверяет разделитель после имени модуля в начале
func (сам *ТМодуль) _РазделительПроверить() (ош error) {
	мКонс.Конс.Отладить("ТуМодуль._РазделительПроверить()")
	слово := сам.слова[0]
	стрСлово := слово.Слово()
	стрИсх := слово.Строка()
	if стрСлово != ";" {
		return мФмт.Errorf("ТуМодуль._РазделительПроверить(): ошибка в окончании названия модуля, слово=[%v], строка=[%v]\n",
			стрСлово, стрИсх)
	}
	if сам.слова, ош = мИнт.СловаОбрезать(сам.слова); ош != nil {
		return мФмт.Errorf("ТуМодуль._РазделительПроверить(): ошибка пи обрезании слов модуля\n\t%v", ош)
	}
	return nil
}

func (сам *ТМодуль) _МодульКонецНайти() (ош error) {
	мКонс.Конс.Отладить("ТуМодуль._МодульКонецНайти()")
	цСчётОбр := len(сам.слова) - 1
	for цСчётОбр >= 0 {
		словоТочка := сам.слова[цСчётОбр]
		// Нашли конечную точку?
		стрТочка := словоТочка.Строка()
		if стрТочка == "." {
			// Попытка найти END. КсМодуль содержит множество слов
			словоКонец := сам.слова[цСчётОбр-2]
			for _, строкаКонец := range мСлово.КсКонец {
				стрКонец := словоКонец.Слово()
				if стрКонец == строкаКонец {
					// Попытка проверить совпадение имя модуля и конца модуля
					словоИмя := сам.слова[цСчётОбр-1]
					стрИмя := сам.слово.Слово()
					стрИмя2 := словоИмя.Слово()
					if стрИмя2 == стрИмя {
						сам._ХвостОтбросить(цСчётОбр)
						return nil
					}
					// Это гарантированная ошибка, так как не было возврата
					return мФмт.Errorf("ТуМодуль._КонецМодульНайти(): имя модуля в начале и в конце модуля не совпадают\n\t%v %v",
						стрИмя, стрИмя2)
				}
			}
		}
		цСчётОбр--
	}
	if цСчётОбр == 0 { // Гарантированная ошибка. Так быть не может.
		return мФмт.Errorf("ТуМодуль.__КонецМодуль_Найти(): нет завершающей точки в модуле")
	}
	return nil
}

// Отбросить всё, что за END <name_module>.
func (сам *ТМодуль) _ХвостОтбросить(пСчётОбр int) {
	//мКонс.Конс.Отладить("Хвост отбросить")
	for индекс := 0; индекс < (пСчётОбр - 2); индекс++ {
		сам.слова[индекс] = сам.слова[индекс+1]
	}
	//мКонс.Конс.Отладить("Последнее слово перед концом модуля:" + мФмт.Sprintf("%v", пСчётОбр))
}

// Проверяет, что MODULE реально один в модуле, первый уже отброшен
func (сам *ТМодуль) _МодульОдинПроверить() (ош error) {
	мКонс.Конс.Отладить("ТуМодуль._МодульОдинПроверить()")
	for индекс := range сам.слова {
		слово := сам.слова[индекс]
		for индекс2 := range мСлово.КсМодуль {
			стрМодуль := слово.Слово()
			if стрМодуль == мСлово.КсМодуль[индекс2] {
				return мФмт.Errorf("ТуМодуль._МодульОдинПроверить(): MODULE встречается больше одного раза")
			}
		}
	}
	return nil
}

// Проверяет, что модуль начинается правильно
func (сам *ТМодуль) _МодульНачалоПроверить() (ош error) {
	мКонс.Конс.Отладить("ТуМодуль.__МодульНачало_Проверить()")
	бРез := false
	слово := сам.слова[0]
	for индекс := range мСлово.КсМодуль {
		стрМодуль := слово.Слово()
		if стрМодуль == мСлово.КсМодуль[индекс] {
			бРез = true
		}
	}
	if !бРез {
		return мФмт.Errorf("ТуМодуль._МодульНачалоПроверить(): модуль не начинается с MODULE")
	}
	if сам.слова, ош = мИнт.СловаОбрезать(сам.слова); ош != nil {
		return мФмт.Errorf("ТуМодуль._МодульНачалоПроверить(): ошибка при обрезании слов модуля\n\t%v", ош)
	}
	return nil
}

//Слова -- возвращает список слов модуля
func (сам *ТМодуль) Слова() map[int]мИнт.ИСлово {
	return сам.слова
}

//Имя -- возвращает список слов модуля
func (сам *ТМодуль) Имя() мИнт.СМодуль {
	return сам.имя
}
