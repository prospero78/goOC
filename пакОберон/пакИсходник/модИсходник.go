package пакИсходник

/*
	Пакет предоставляет тип для работы с исходным файлом.
	Позволяет текст считывать целиком, в другой части позволяет работать с
	отдельными строками по номерам.
*/

import (
	мАбс "../пакАбстракция"
	мИсхСтр "../пакИсходникСтроки"
	мИсхФайл "../пакИсходникФайл"
	мКонс "../пакКонсоль"
	мКоорд "../пакКоорд"
	мЛит "../пакЛитера"
	мРес "../пакРесурс"
	мСлово "../пакСлово"
	мФмт "fmt"
)

//ТИсходник -- тип для считывания исходника и разбиения его на строки
type ТИсходник struct {
	ИсхФайл      мАбс.АИсхФайл       // содержимое файла исходника целиком
	ИсхСтр       мАбс.АИсхСтроки     // исходник построчно
	Коорд        мАбс.АКоордИзм      // изменяемые координаты
	слова        map[int]мАбс.АСлово // справочник слов
	бСловаГотово bool                // признак готовности на разделения на слова
	цСтр         мАбс.СКоордСтр      // текущая строка в исходном тексте
	лит          мАбс.АЛит           //текущая литера
	литНомер     мАбс.СЛитНомер      // текущая позиция в исходном тексте
}

//ИсходникНов -- возвращает ссылку на новый ТИсходник
func ИсходникНов() (исх *ТИсходник, ош error) {
	мКонс.Конс.Отладить("пакИсходник.Новый()")
	исх = &ТИсходник{}
	if исх == nil {
		return nil, мФмт.Errorf("ИсходникНов(): нет памяти?\n")
	}
	if исх.ИсхФайл, ош = мИсхФайл.ИсхФайлНов(); ош != nil {
		return nil, мФмт.Errorf("ИсходникНов(): ошибка при создании ТИсхФайл\n\t%v", ош)
	}
	исх.ИсхСтр = мИсхСтр.Новый()
	if исх.Коорд, ош = мКоорд.КоордИзмНов(1, 0); ош != nil {
		return nil, мФмт.Errorf("пакТсходник.Новый(): ошибка при создании координат\n\t%v", ош)
	}
	исх.слова = make(map[int]мАбс.АСлово)
	if исх.лит, ош = мЛит.ЛитераНов(); ош != nil {
		return nil, мФмт.Errorf("пакТсходник.Новый(): ошибка при создании объекта литеры\n\t%v", ош)
	}
	return исх, nil
}

//Обработать -- главный цикл чтения и разбиения на строки исходника
func (сам *ТИсходник) Обработать(пИмяФайла мАбс.СИсхФайл) (ош error) {
	мКонс.Конс.Отладить("ТИсходник.Обработать()")
	if ош = сам.ИсхФайл.Считать(пИмяФайла); ош != nil {
		return мФмт.Errorf("ТИсходник.Обработать(): ошибка при чтении файа исходника\n\t%v", ош)
	}
	//сам.конс.Отладить("ffff @" + string(сам.ИсхФайл.Исходник()) + "@")
	сам.ИсхСтр.НаСтрокиРазбить(сам.ИсхФайл.Исходник())
	if ош = сам._НаСловаРазделить(); ош != nil {
		return мФмт.Errorf("ТИсходник.Обработать(): ошибка при обработке исходника\n\t%v", ош)
	}
	return nil
}

// Добавляет слово с атрибутами положения в исходном тексте
// Строки исходника остаются как были.
func (сам *ТИсходник) _СловоДобав(пСлово мАбс.ССлово) (ош error) {
	мФмт.Printf("%v-%v %v\n", сам.Коорд.Стр(), сам.Коорд.Поз(), пСлово)
	крдФикс, ош := мКоорд.КоордНов(сам.Коорд.Стр(), сам.Коорд.Поз())
	if ош != nil {
		return мФмт.Errorf("ТИсходник._СловоДобав(): ошибка при создании координат слова\n\t%v", ош)
	}
	слово, ош := мСлово.СловоНов(крдФикс, пСлово, сам.ИсхСтр.Строка(сам.Коорд.Стр()))
	if ош != nil {
		return мФмт.Errorf("ТИсходник._СловоДобав(): ошибка при добавлении ТуСлово\n\t%v", ош)
	}
	сам.Коорд.ПозДоб()
	сам.слова[len(сам.слова)] = слово
	//сам.конс.Отладить("Слово № " + мФмт.Sprintf("%d", слово.Номер()) + "=\"" + пСлово + "\"")

	return nil
}

// Разделяет исходник на слова
func (сам *ТИсходник) _НаСловаРазделить() (ош error) {
	мКонс.Конс.Отладить("ТИсходник._НаСловаРазделить()")
	for !сам.бСловаГотово {
		if ош = сам._СловоВыделить(); ош != nil {
			return мФмт.Errorf("ТИсходник._НаСловаРазделить(): ошибка при выделении слова\n\t%v", ош)
		}
	}
	мКонс.Конс.Отладить(мФмт.Sprintf("ТИсходник._НаСловаРазделить(): всего слов: [%v]\n", len(сам.слова)))
	return nil
}

func (сам *ТИсходник) _ЛитШаг() { //Получает следующую позицию литеры
	сам.Коорд.ПозДоб()
	сам.литНомер++
}

// Выделяет слово из исходного текста
func (сам *ТИсходник) _СловоВыделить() (ош error) {
	_Пробел := func(пЛит мАбс.СЛит) {
		if пЛит == " " {
			//сам.Коорд.ПозДоб()
			сам._ЛитШаг()
		} else if пЛит == "\t" {
			for i := 0; i < мРес.ШирТаба; i++ {
				//сам.Коорд.ПозДоб()
				сам._ЛитШаг()
			}
		}
	}
	_ТочкаЗапятая := func(пЛит мАбс.СЛит) (ош error) {
		if пЛит == ";" {
			if ош = сам._СловоДобав(";"); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._ТочкаЗапятая(): ошибка при добавлении точки с запятой\n\t%v", ош)
			}
			сам._ЛитШаг()
		}
		return nil
	}
	_НоваяСтрока := func(пЛит мАбс.СЛит) {
		if пЛит == "\n" {
			сам.Коорд.СтрДоб()
			сам.Коорд.ПозСброс()
			сам.литНомер++
		}
	}
	_Запятая := func(пЛит мАбс.СЛит) (ош error) {
		if пЛит == "," {
			if ош = сам._СловоДобав(","); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._Запятая(): ошибка при добавлении запятой\n\t%v", ош)
			}
			сам._ЛитШаг()
		}
		return nil
	}
	_Плюс := func(пЛит мАбс.СЛит) (ош error) {
		if пЛит == "+" {
			if ош = сам._СловоДобав("+"); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._Запятая(): ошибка при добавлении плюса\n\t%v", ош)
			}
			сам._ЛитШаг()
		}
		return nil
	}
	_Минус := func(пЛит мАбс.СЛит) (ош error) {
		if пЛит == "-" {
			if ош = сам._СловоДобав("-"); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._Минус(): ошибка при добавлении минуса\n\t%v", ош)
			}
			сам._ЛитШаг()
		}
		return nil
	}
	_Деление := func(пЛит мАбс.СЛит) (ош error) {
		if пЛит == "/" {
			if ош = сам._СловоДобав("/"); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._Деление(): ошибка при добавлении деления\n\t%v", ош)
			}
			сам._ЛитШаг()
		}
		return nil
	}
	_ЛеваяСкобка := func(пЛит мАбс.СЛит) (ош error) {
		if пЛит == "(" {
			лит1, ош := сам.ИсхФайл.Лит(сам.литНомер + 1)
			if ош != nil {
				return мФмт.Errorf("ЛеваяСкобка(): ошибка при получении литеры\n\t%v", ош)
			}
			if пЛит+лит1 == "(*" {
				if ош = сам._СловоДобав("(*"); ош != nil {
					return мФмт.Errorf("ТИсходник._СловоВыделить()._ЛеваяСкобка(): ошибка при добавлении открытия комментария\n\t%v", ош)
				}
				сам._ЛитШаг()
				сам._ЛитШаг()
				return nil
			}
			if ош = сам._СловоДобав("("); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._ЛеваяСкобка(): ошибка при добавлении левой скобки\n\t%v", ош)
			}
			сам._ЛитШаг()
		}
		return nil
	}
	_ПраваяСкобка := func(пЛит мАбс.СЛит) (ош error) {
		if пЛит == ")" {
			if ош = сам._СловоДобав(")"); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._ПраваяСкобка(): ошибка при добавлении правой скобки\n\t%v", ош)
			}
			сам._ЛитШаг()
		}
		return nil
	}
	_Умножить := func(пЛит мАбс.СЛит) (ош error) {
		if пЛит == "*" {
			лит1, ош := сам.ИсхФайл.Лит(сам.литНомер + 1)
			if ош != nil {
				return мФмт.Errorf("Умножить(): ошибка при получении литеры\n\t%v", ош)
			}
			if пЛит+лит1 == "*)" {
				if ош = сам._СловоДобав("*)"); ош != nil {
					return мФмт.Errorf("ТИсходник._СловоВыделить()._Умножить(): ошибка при добавлении закрытия комментария\n\t%v", ош)
				}
				сам._ЛитШаг()
				сам._ЛитШаг()
				return nil
			}
			if ош = сам._СловоДобав("*"); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._Умножить(): ошибка при добавлении умножить\n\t%v", ош)
			}
			сам._ЛитШаг()
		}
		return nil
	}
	_Двоеточие := func(пЛит мАбс.СЛит) (ош error) {
		if пЛит == ":" {
			лит1, ош := сам.ИсхФайл.Лит(сам.литНомер + 1)
			if ош != nil {
				return мФмт.Errorf("Двоеточие(): ошибка при получении литеры\n\t%v", ош)
			}
			if пЛит+лит1 == ":=" {
				if ош = сам._СловоДобав(":="); ош != nil {
					return мФмт.Errorf("ТИсходник._СловоВыделить()._Двоеточие(): ошибка при добавлении присовить\n\t%v", ош)
				}
				сам._ЛитШаг()
				сам._ЛитШаг()
				return nil
			}
			if ош = сам._СловоДобав(":"); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._Двоеточие(): ошибка при добавлении двоеточия\n\t%v", ош)
			}
			сам._ЛитШаг()
		}
		return nil
	}
	_ПереводКаретки := func(пЛит мАбс.СЛит) {
		if пЛит == "\r" {
			сам._ЛитШаг()
		}
	}
	_ЕслиСлово := func(пЛит мАбс.СЛит) (ош error) {
		слово := мАбс.ССлово("")
		// это что-то строковое
		начИмя := пЛит == "_" || сам.лит.ЕслиБуква(пЛит)
		if начИмя {
			var бДальше bool
			for {
				продИмя := пЛит == "_" || сам.лит.ЕслиБуква(пЛит) || сам.лит.ЕслиЦифра(пЛит)
				бДальше = !(пЛит == ".") && продИмя
				if бДальше {
					слово += мАбс.ССлово(пЛит)
					сам.литНомер++
					пЛит, ош = сам.ИсхФайл.Лит(сам.литНомер)
					if ош != nil {
						return мФмт.Errorf("ЕслиСущность(): ошибка при получении литеры\n\t%v", ош)
					}
					continue
				}
				break
			}
			// откат на одну позицию
			//сам.литНомер--
			if ош = сам._СловоДобав(слово); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._ЕслиСлово(): ошибка при добавлении имени\n\t%v", ош)
			}
			сам.Коорд.ПозУст(сам.Коорд.Поз() + мАбс.СКоордПоз(len(слово)) - 1)
			// Это что-то числовое
		} else if сам.лит.ЕслиЦифра(пЛит) || пЛит == "." { //Это число
			for {
				if сам.лит.ЕслиЦифра(пЛит) || пЛит == "." {
					слово += мАбс.ССлово(пЛит)
					сам.литНомер++
					пЛит, ош = сам.ИсхФайл.Лит(сам.литНомер)
					if ош != nil {
						return мФмт.Errorf("ЕслиСущность(): ошибка при получении литеры\n\t%v", ош)
					}
					continue
				}
				break
			}
			// откат на одну позицию
			//сам.литНомер--
			if ош = сам._СловоДобав(слово); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._ЕслиСлово(): ошибка при добавлении числа\n\t%v", ош)
			}
			сам.Коорд.ПозУст(сам.Коорд.Поз() + мАбс.СКоордПоз(len(слово)) )
		}
		return nil
	}
	_Равно := func(пЛит мАбс.СЛит) (ош error) {
		if пЛит == "=" {
			if ош = сам._СловоДобав("="); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._Равно(): ошибка при добавлении равно\n\t%v", ош)
			}
			сам._ЛитШаг()
		}
		return nil
	}
	_Точка := func(пЛит мАбс.СЛит) (ош error) {
		if пЛит == "." {
			if ош = сам._СловоДобав("."); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._Точка(): ошибка при добавлении равно\n\t%v", ош)
			}
			сам._ЛитШаг()
		}
		return nil
	}
	_Кавычка2 := func(пЛит мАбс.СЛит) (ош error) { // вычисляет строки
		if пЛит == "\"" {
			слово := мАбс.ССлово("\"")
			пЛит = ""
			for пЛит != "\"" {
				сам._ЛитШаг()
				if пЛит, ош = сам.ИсхФайл.Лит(сам.литНомер); ош != nil {
					return мФмт.Errorf("Кавычка2(): ошибка при получении литеры\n\t%v", ош)
				}
				слово += мАбс.ССлово(пЛит)
			}
			if ош = сам._СловоДобав(слово); ош != nil {
				return мФмт.Errorf("ТИсходник._СловоВыделить()._Кавычка2(): ошибка при добавлении кавычек\n\t%v", ош)
			}
			сам._ЛитШаг()
		}
		return nil
	}

	//сам.конс.Отладить("ТИсходник.__Слово_Выделить()")
	лит := мАбс.СЛит("")
	if сам.литНомер < мАбс.СЛитНомер(сам.ИсхФайл.Размер()-1) {
		лит, ош = сам.ИсхФайл.Лит(сам.литНомер)
		if ош != nil {
			return мФмт.Errorf("_СловоВыделить(): ошибка при получении слова\n\t%v", ош)
		}
		//мФмт.Printf(" Лит=[%v] ", лит)
		_Пробел(лит)
		_НоваяСтрока(лит)
		_ПереводКаретки(лит)
		if ош = _Запятая(лит); ош != nil {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении запятой\n\t%v", ош)
		}
		if ош = _ТочкаЗапятая(лит); ош != nil {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении точки с запятой\n\t%v", ош)
		}
		if ош = _Плюс(лит); ош != nil {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении плюса\n\t%v", ош)
		}
		if ош = _Минус(лит); ош != nil {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении минуса\n\t%v", ош)
		}
		if ош = _Деление(лит); ош != nil {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении деления\n\t%v", ош)
		}
		if ош = _ЛеваяСкобка(лит); ош != nil {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении левой скобки\n\t%v", ош)
		}
		if ош = _ПраваяСкобка(лит); ош != nil {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении правой скобки\n\t%v", ош)
		}
		if ош = _Умножить(лит); ош != nil {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении умножить\n\t%v", ош)
		}
		if ош = _Двоеточие(лит); ош != nil {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении двоеточия\n\t%v", ош)
		}
		if ош = _Равно(лит); ош != nil {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении равно\n\t%v", ош)
		}
		if ош = _Точка(лит); ош != nil {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении точки\n\t%v", ош)
		}
		if ош = _ЕслиСлово(лит); ош != nil {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении слова\n\t%v", ош)
		}
		if ош = _Кавычка2(лит); ош != nil {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении кавычек\n\t%v", ош)
		}
		//сам.литНомер++
	} else {
		сам.бСловаГотово = true
	}
	return nil
}

//Слова -- возвращает все распознанные слова в исходнике
func (сам *ТИсходник) Слова() map[int]мАбс.АСлово {
	return сам.слова
}
