package пакСекция

/*
	Предоставляет возможность сканеру секций определеить правильно ли
	описан модуль.
*/

import (
	мИнт "../пакИнтерфейсы"
	мКлючи "../пакКлючи"
	мКонс "../пакКонсоль"
	мСкан "../пакСканер"
	мСлово "../пакСлово"
	мФмт "fmt"
)

//ТСекцияМодуль -- предоставляет тип, содержащий слова модуля
type ТСекцияМодуль struct {
	*ТСекция
	модульИмя     мИнт.СМодульИмя
	слово         мИнт.ИСлово // слово в котором хранится имя модуля
	сканер        мИнт.ИСканер
	секцияКоммент мИнт.ИСекцияКоммент //Секция комментариев
}

//МодульНов -- возвращает на новый ТСекцияМодуль
func МодульНов() (модуль *ТСекцияМодуль, ош error) {
	мКонс.Конс.Отладить("МодульНов()")
	секция, ош := СекцияНов("МОДУЛЬ")
	if ош != nil {
		return nil, мФмт.Errorf("МодульНов(): ошибка при создании базовой секции\n\t%v", ош)
	}
	модуль = &ТСекцияМодуль{ТСекция: секция}
	if модуль == nil {
		return nil, мФмт.Errorf("МодульНов(): нет памяти на модуль?\n")
	}
	if модуль.сканер, ош = мСкан.СканерНов(); ош != nil {
		return nil, мФмт.Errorf("МодульНов(): ошибка при создании сканера модуля\n\t%v", ош)
	}
	if модуль.секцияКоммент, ош = СекцияКомментНов(); ош != nil {
		return nil, мФмт.Errorf("МодульНов(): ошибка при создании секции комментариев\n\t%v", ош)
	}
	return модуль, nil
}

//Обработать -- обеспечивает выделение слов модуля
func (сам *ТСекцияМодуль) Обработать(пФайлИмя мИнт.СИсхФайл) (ош error) {
	мКонс.Конс.Отладить("ТСекцияМодуль.Обработать()")
	if ош = сам.сканер.Обработать(пФайлИмя); ош != nil {
		return мФмт.Errorf("ТСекцияМодуль.Обработать(): ошибка при работе сканера\n\t%v", ош)
	}
	сам.словаМодуля = сам.сканер.Слова()
	мФмт.Printf("Всего слов в модуле [%v]\n", len(сам.СловаМодуля()))

	//Вывести все комментарии в отдельную секцию. Там могут быть различные опции
	if ош = сам.секцияКоммент.СловаУст(сам.СловаМодуля()); ош != nil {
		return мФмт.Errorf("ТСекцияМодуль.Обработать(): ошибка при установке слов секции комментариев\n\t%v", ош)
	}
	if ош = сам.секцияКоммент.Обработать(); ош != nil {
		return мФмт.Errorf("ТСекцияМодуль.Обработать(): ошибка при обработке секции комментариев\n\t%v", ош)
	}
	сам.СловаУст(сам.секцияКоммент.СловаМодуля())
	мФмт.Printf("\tВсего слов в комментариях: [%v]\n", len(сам.секцияКоммент.СловаСекции()))

	if ош = сам._МодульНачалоПроверить(); ош != nil {
		return мФмт.Errorf("ТСекцияМодуль.Обработать(): ошибка при поиске начала модуля\n\t%v", ош)
	}
	if ош = сам._ИмяПроверить(); ош != nil {
		return мФмт.Errorf("ТСекцияМодуль.Обработать(): ошибка при поиске имени модуля\n\t%v", ош)
	}
	if ош = сам._РазделительПроверить(); ош != nil {
		return мФмт.Errorf("ТСекцияМодуль.Обработать(): ошибка при поиске разделителя заголовка модуля\n\t%v", ош)
	}
	if ош = сам._МодульКонецНайти(); ош != nil {
		return мФмт.Errorf("ТСекцияМодуль.Обработать(): ошибка при поиске конца модуля\n\t%v", ош)
	}
	if ош = сам._МодульОдинПроверить(); ош != nil {
		return мФмт.Errorf("ТСекцияМодуль.Обработать(): ошибка при поиске единственного MODULE\n\t%v", ош)
	}
	мФмт.Printf("Всего слов в модуле после обработки > [%v]\n", len(сам.СловаМодуля()))
	return nil
}

// Проверяет имя модуля в тексте
func (сам *ТСекцияМодуль) _ИмяПроверить() (ош error) {
	мКонс.Конс.Отладить("ТСекцияМодуль._ИмяПроверить()")
	слово := сам.СловаМодуля()[0]
	//мКонс.Конс.Отладить("Проверка имени модуля: \"" + слово.Строка() + "\"")
	if слово.ЕслиИмяСтрого() {
		сам.модульИмя = мИнт.СМодульИмя(слово.Слово())
	} else {
		стрСлово := слово.Слово()
		стрИсх := слово.Строка()
		return мФмт.Errorf("ТСекцияМодуль._ИмяПроверить(): такое имя модуля запрещено, имя=[%v], строка=[%v]\n", стрСлово, стрИсх)
	}
	if сам.словаМодуля, ош = мИнт.СловаОбрезать(сам.СловаМодуля()); ош != nil {
		return мФмт.Errorf("ТСекцияМодуль._ИмяПроверить(): ошибка при обрезке слов модуля\n\t%v", ош)
	}
	return nil
}

// Проверяет разделитель после имени модуля в начале
func (сам *ТСекцияМодуль) _РазделительПроверить() (ош error) {
	мКонс.Конс.Отладить("ТСекцияМодуль._РазделительПроверить()")
	слово := сам.СловаМодуля()[0]
	стрСлово := слово.Слово()
	стрИсх := слово.Строка()
	ок, ош := мКлючи.Ключи.Проверить(";", мИнт.СКлюч(стрСлово))
	if ош != nil {
		return мФмт.Errorf("ТСекцияМодуль._РазделительПроверить(): ошибка при проверке разделителя модуля\n\t%v", ош)
	}
	if !ок {
		return мФмт.Errorf("ТСекцияМодуль._РазделительПроверить(): ошибка в окончании названия модуля, слово=[%v], строка=[%v]\n",
			стрСлово, стрИсх)
	}
	if сам.словаМодуля, ош = мИнт.СловаОбрезать(сам.СловаМодуля()); ош != nil {
		return мФмт.Errorf("ТСекцияМодуль._РазделительПроверить(): ошибка пи обрезании слов модуля\n\t%v", ош)
	}
	//Теперь первое слово указывает на следующую интсрукцию
	return nil
}

func (сам *ТСекцияМодуль) _МодульКонецНайти() (ош error) {
	мКонс.Конс.Отладить("ТСекцияМодуль._МодульКонецНайти()")
	цСчётОбр := len(сам.СловаМодуля()) - 1
	for цСчётОбр := len(сам.СловаМодуля()); цСчётОбр >= 0; цСчётОбр-- {
		словоТочка := сам.СловаМодуля()[цСчётОбр]
		// Нашли конечную точку?
		стрТочка := словоТочка.Слово()
		ок, ош := мКлючи.Ключи.Проверить(".", стрТочка)
		if ош!=nil{
			return мФмт.Errorf("ТСекцияМодуль._КонецМодульНайти(): ошибка при поиске конца модуля\n\t%v",
						стрИмя, стрИмя2)
		}
		if стрТочка == "." {
			// Попытка найти END. КсМодуль содержит множество слов
			словоКонец := сам.СловаМодуля()[цСчётОбр-2]
			for _, строкаКонец := range мСлово.КсКонец {
				стрКонец := словоКонец.Слово()
				if стрКонец == строкаКонец {
					// Попытка проверить совпадение имя модуля и конца модуля
					словоИмя := сам.СловаМодуля()[цСчётОбр-1]
					стрИмя := сам.слово.Слово()
					стрИмя2 := словоИмя.Слово()
					if стрИмя2 == стрИмя {
						сам._ХвостОтбросить(цСчётОбр)
						return nil
					}
					// Это гарантированная ошибка, так как не было возврата
					return мФмт.Errorf("ТСекцияМодуль._КонецМодульНайти(): имя модуля в начале и в конце модуля не совпадают\n\t%v %v",
						стрИмя, стрИмя2)
				}
			}
		}
	}
	if цСчётОбр == 0 { // Гарантированная ошибка. Так быть не может.
		return мФмт.Errorf("ТСекцияМодуль._КонецМодуль_Найти(): нет завершающей точки в модуле")
	}
	return nil
}

// Отбросить всё, что за END <name_module>.
func (сам *ТСекцияМодуль) _ХвостОтбросить(пСчётОбр int) {
	//мКонс.Конс.Отладить("Хвост отбросить")
	for индекс := 0; индекс < (пСчётОбр - 2); индекс++ {
		сам.словаМодуля[индекс] = сам.СловаМодуля()[индекс+1]
	}
	//мКонс.Конс.Отладить("Последнее слово перед концом модуля:" + мФмт.Sprintf("%v", пСчётОбр))
}

// Проверяет, что MODULE реально один в модуле, первый уже отброшен
func (сам *ТСекцияМодуль) _МодульОдинПроверить() (ош error) {
	мКонс.Конс.Отладить("ТМодТСекцияМодульуль._МодульОдинПроверить()")
	for _, ключ := range сам.СловаМодуля() {
		ок, ош := мКлючи.Ключи.Проверить("MODULE", мИнт.СКлюч(ключ.Слово()))
		if ош != nil {
			return мФмт.Errorf("ТСекцияМодуль._МодульОдинПроверить(): ошибка при проверке единственного MODULE\n\t%v", ош)
		}
		if !ок {
			return мФмт.Errorf("ТСекцияМодуль._МодульОдинПроверить(): MODULE встречается больше одного раза")
		}
	}
	return nil
}

// Проверяет, что модуль начинается правильно
func (сам *ТСекцияМодуль) _МодульНачалоПроверить() (ош error) {
	мКонс.Конс.Отладить("ТСекцияМодуль._МодульНачало_Проверить()")
	словоМОДУЛЬ := сам.СловаМодуля()[0]
	//Проверить, что это ключевое слово МОДУЛЬ
	ок, ош := мКлючи.Ключи.Проверить("MODULE", мИнт.СКлюч(словоМОДУЛЬ.Слово()))
	if ош != nil {
		return мФмт.Errorf("ТСекцияМодуль._МодульНачалоПроверить(): ошибка при проверке слова\n\t%v", ош)
	}
	if !ок {
		return мФмт.Errorf("ТСекцияМодуль._МодульНачалоПроверить(): нет начала модуля, MODULE=\"%v\"\n\t%v", словоМОДУЛЬ, ош)
	}
	сам.модульИмя = мИнт.СМодульИмя(словоМОДУЛЬ.Слово())
	if сам.словаМодуля, ош = мИнт.СловаОбрезать(сам.СловаМодуля()); ош != nil {
		return мФмт.Errorf("ТСекцияМодуль._МодульНачалоПроверить(): ошибка при обрезании слов модуля\n\t%v", ош)
	}
	//Теперь следующее слово имя. обрабатывается уровнем выше.
	return nil
}

//МодульИмя -- возвращает список слов модуля
func (сам *ТСекцияМодуль) МодульИмя() мИнт.СМодульИмя {
	return сам.модульИмя
}
