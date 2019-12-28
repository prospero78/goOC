package пакИсходник

/*
	Пакет предоставляет тип для работы с исходным файлом.
	Позволяет текст считывать целиком, в другой части позволяет работать с
	отдельными строками по номерам.
*/

import (
	мФмт "fmt"

	мИф "../sourcefile"
	мКонс "../пакКонсоль"
	мКоорд "../coord"
	мЛит "../liter"
	мПстр "../пакПулИсхСтроки"
	мРес "../пакРесурс"
	мСлово "../пакСлово"
	мПоз "../coord/coordx"
)

//ИИсходник -- интерфейс для исходника
type ИИсходник interface {
	Обработать(мИф.СФайлИсхИмя) error
	Слова() map[мСлово.ССловоНомер]мСлово.ИСлово
}

//ТИсходник -- тип для считывания исходника и разбиения его на строки
type ТИсходник struct {
	исхФайл      мИф.ИИсхФайл                         // содержимое файла исходника целиком
	пулСтр       мПстр.ИПулИсхСтроки                  // исходник построчно
	Коорд        мКоорд.ИКоордИзм                     // изменяемые координаты
	слова        map[мСлово.ССловоНомер]мСлово.ИСлово // справочник слов
	бСловаГотово bool                                 // признак готовности на разделения на слова
	//стрНомер         мНс.ССтрНомер                         // текущая строка в исходном тексте
	лит      мЛит.ИЛит    //текущая литера
	литНомер мПоз.ССтрПоз // текущая позиция в исходном тексте
}

//ИсходникНов -- возвращает ссылку на новый ТИсходник
func ИсходникНов(пИмяФайла мИф.СФайлИсхИмя) (исх ИИсходник, ош error) {
	мКонс.Конс.Отладить("пакИсходник.Новый()")
	_исх := ТИсходник{}
	if _исх.исхФайл, ош = мИф.ИсхФайлНов(пИмяФайла); ош != nil {
		return nil, мФмт.Errorf("ИсходникНов(): ОШИБКА при создании ТИсхФайл\n\t%v", ош)
	}
	if _исх.пулСтр, ош = мПстр.ПулИсхСтрокиНов(); ош != nil {
		return nil, мФмт.Errorf("пакТсходник.Новый(): ОШИБКА при создании пула строк\n\t%v", ош)
	}
	if _исх.Коорд, ош = мКоорд.КоордИзмНов(1, 0); ош != nil {
		return nil, мФмт.Errorf("пакТсходник.Новый(): ОШИБКА при создании координат\n\t%v", ош)
	}
	_исх.слова = make(map[мСлово.ССловоНомер]мСлово.ИСлово)
	if _исх.лит, ош = мЛит.ЛитераНов(); ош != nil {
		return nil, мФмт.Errorf("пакТсходник.Новый(): ОШИБКА при создании объекта литеры\n\t%v", ош)
	}
	return &_исх, nil
}

//Обработать -- главный цикл чтения и разбиения на строки исходника
func (сам *ТИсходник) Обработать(пИмяФайла мИф.СФайлИсхИмя) (ош error) {
	мКонс.Конс.Отладить("ТИсходник.Обработать()")
	//сам.конс.Отладить("ffff @" + string(сам.исхФайл.Исходник()) + "@")
	сам.пулСтр.НаСтрокиРазбить(сам.исхФайл.Исходник())
	if ош = сам._НаСловаРазделить(); ош != nil {
		return мФмт.Errorf("ТИсходник.Обработать(): ОШИБКА при обработке исходника\n\t%v", ош)
	}
	return nil
}

// Добавляет слово с атрибутами положения в исходном тексте
// Строки исходника остаются как были.
func (сам *ТИсходник) _СловоДобав(пСлово мСлово.ССлово) (ош error) {
	//мФмт._rintf("%v %v-%v %v\n", len(сам.слова), сам.Коорд.Стр(), сам.Коорд.Поз(), пСлово)
	крдФикс, ош := мКоорд.КоордНов(сам.Коорд.СтрНомер(), сам.Коорд.СтрПоз())
	if ош != nil {
		return мФмт.Errorf("ТИсходник._СловоДобав(): ОШИБКА при создании координат слова\n\t%v", ош)
	}
	строка, ош := сам.пулСтр.Строка(сам.Коорд.СтрНомер())
	if ош != nil {
		return мФмт.Errorf("ТИсходник._СловоДобав(): ОШИБКА при получении строки\n\t%v", ош)
	}
	слово, ош := мСлово.СловоНов(крдФикс, пСлово, строка)
	if ош != nil {
		return мФмт.Errorf("ТИсходник._СловоДобав(): ОШИБКА при добавлении ТСлово\n\t%v", ош)
	}
	сам.Коорд.СтрПозДоб()
	сам.слова[мСлово.ССловоНомер(len(сам.слова))] = слово
	//сам.конс.Отладить("Слово № " + мФмт.S_rintf("%d", слово.Номер()) + "=\"" + пСлово + "\"")

	return nil
}

// Разделяет исходник на слова
func (сам *ТИсходник) _НаСловаРазделить() (ош error) {
	мКонс.Конс.Отладить("ТИсходник._НаСловаРазделить()")
	for !сам.бСловаГотово {
		if ош = сам._СловоВыделить(); ош != nil {
			return мФмт.Errorf("ТИсходник._НаСловаРазделить(): ОШИБКА при выделении слова\n\t%v", ош)
		}
	}
	мКонс.Конс.Отладить(мФмт.Sprintf("ТИсходник._НаСловаРазделить(): всего слов: [%v]\n", len(сам.слова)))
	return nil
}

func (сам *ТИсходник) _ЛитШаг() { //Получает следующую позицию литеры
	сам.Коорд.СтрПозДоб()
	сам.литНомер++
}

// Выделяет слово из исходного текста
func (сам *ТИсходник) _СловоВыделить() (ош error) {
	_Пробел := func() (ок bool) {
		if сам.лит.Получ() == " " {
			сам._ЛитШаг()
			return true
		} else if сам.лит.Получ() == "\t" { //Позицию добавить на столько, сколько пробелов считается таб
			сам.литНомер++
			for i := 0; i < мРес.ШирТаба; i++ {
				сам.Коорд.СтрПозДоб()
			}
			return true
		}
		return false
	}
	_ТочкаЗапятая := func() (ок bool, ош error) {
		if сам.лит.Получ() == ";" {
			if ош = сам._СловоДобав(";"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._ТочкаЗапятая(): ОШИБКА при добавлении точки с запятой\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_НоваяСтрока := func() (ок bool) {
		if сам.лит.Получ() == "\n" {
			сам.Коорд.СтрНомерДоб()
			сам.Коорд.СтрПозСброс()
			сам.литНомер++
			return true
		}
		return false
	}
	_Запятая := func() (ок bool, ош error) {
		if сам.лит.Получ() == "," {
			if ош = сам._СловоДобав(","); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Запятая(): ОШИБКА при добавлении запятой\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Плюс := func() (ок bool, ош error) {
		if сам.лит.Получ() == "+" {
			if ош = сам._СловоДобав("+"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Запятая(): ОШИБКА при добавлении плюса\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Минус := func() (ок bool, ош error) {
		if сам.лит.Получ() == "-" {
			if ош = сам._СловоДобав("-"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Минус(): ОШИБКА при добавлении минуса\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Деление := func() (ок bool, ош error) {
		if сам.лит.Получ() == "/" {
			if ош = сам._СловоДобав("/"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Деление(): ОШИБКА при добавлении деления\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_ЛеваяСкобка := func() (ок bool, ош error) {
		if сам.лит.Получ() == "(" {
			лит1, ош := сам.исхФайл.Лит(сам.литНомер + 1)
			if ош != nil {
				return false, мФмт.Errorf("ЛеваяСкобка(): ОШИБКА при получении литеры\n\t%v", ош)
			}
			if сам.лит.Получ()+лит1 == "(*" {
				if ош = сам._СловоДобав("(*"); ош != nil {
					return false, мФмт.Errorf("ТИсходник._СловоВыделить()._ЛеваяСкобка(): ОШИБКА при добавлении открытия комментария\n\t%v", ош)
				}
				сам._ЛитШаг()
				сам._ЛитШаг()
				return true, nil
			}
			if ош = сам._СловоДобав("("); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._ЛеваяСкобка(): ОШИБКА при добавлении левой скобки\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_ПраваяСкобка := func() (ок bool, ош error) {
		if сам.лит.Получ() == ")" {
			if ош = сам._СловоДобав(")"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._ПраваяСкобка(): ОШИБКА при добавлении правой скобки\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Умножить := func() (ок bool, ош error) {
		if сам.лит.Получ() == "*" {
			лит1, ош := сам.исхФайл.Лит(сам.литНомер + 1)
			if ош != nil {
				return false, мФмт.Errorf("Умножить(): ОШИБКА при получении литеры\n\t%v", ош)
			}
			if сам.лит.Получ()+лит1 == "*)" {
				if ош = сам._СловоДобав("*)"); ош != nil {
					return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Умножить(): ОШИБКА при добавлении закрытия комментария\n\t%v", ош)
				}
				сам._ЛитШаг()
				сам._ЛитШаг()
				return true, nil
			}
			if ош = сам._СловоДобав("*"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Умножить(): ОШИБКА при добавлении умножить\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Двоеточие := func() (ок bool, ош error) {
		if сам.лит.Получ() == ":" {
			лит1, ош := сам.исхФайл.Лит(сам.литНомер + 1)
			if ош != nil {
				return false, мФмт.Errorf("Двоеточие(): ОШИБКА при получении литеры\n\t%v", ош)
			}
			if сам.лит.Получ()+лит1 == ":=" {
				if ош = сам._СловоДобав(":="); ош != nil {
					return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Двоеточие(): ОШИБКА при добавлении присовить\n\t%v", ош)
				}
				сам._ЛитШаг()
				сам._ЛитШаг()
				return true, nil
			}
			if ош = сам._СловоДобав(":"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Двоеточие(): ОШИБКА при добавлении двоеточия\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_ПереводКаретки := func() bool {
		if сам.лит.Получ() == "\r" {
			сам._ЛитШаг()
			return true
		}
		return false
	}
	_ЕслиСлово := func() (ок bool, ош error) {
		_Имя := func() (ок bool, ош error) {
			слово := мСлово.ССлово("")
			for {
				пЛит, ош := сам.исхФайл.Лит(сам.литНомер)
				if ош != nil {
					return false, мФмт.Errorf("ЕслиСущность(): ОШИБКА при получении литеры\n\t%v", ош)
				}
				if ош = сам.лит.Уст(пЛит); ош != nil {
					return false, мФмт.Errorf("ЕслиСущность(): ОШИБКА при установке литеры\n\t%v", ош)
				}
				бПродИмя := (сам.лит.ЕслиБуква() || сам.лит.ЕслиЦифра() || сам.лит.ЕслиСпецЛит()) && !(сам.лит.Получ() == ".")
				if !бПродИмя {
					break
				}
				слово += мСлово.ССлово(сам.лит.Получ())
				сам._ЛитШаг()
			}
			// откат на одну позицию
			//сам.литНомер--
			if ош = сам._СловоДобав(слово); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._ЕслиСлово(): ОШИБКА при добавлении имени\n\t%v", ош)
			}
			сам.Коорд.СтрПозУст(сам.Коорд.СтрПоз() + мПоз.ССтрПоз(len(слово)))
			return true, nil
		}
		_Число := func() (ок bool, ош error) {
			слово := мСлово.ССлово("")
			for {
				пЛит, ош := сам.исхФайл.Лит(сам.литНомер)
				if ош != nil {
					return false, мФмт.Errorf("ЕслиСущность(): ОШИБКА при получении литеры\n\t%v", ош)
				}
				if ош = сам.лит.Уст(пЛит); ош != nil {
					return false, мФмт.Errorf("ЕслиСущность(): ОШИБКА при установке литеры\n\t%v", ош)
				}
				if !(сам.лит.ЕслиЦифра() || сам.лит.Получ() == ".") {
					break
				}
				слово += мСлово.ССлово(сам.лит.Получ())
				сам._ЛитШаг()
			}
			// откат на одну позицию
			//сам.литНомер--
			if ош = сам._СловоДобав(слово); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._ЕслиСлово(): ОШИБКА при добавлении числа\n\t%v", ош)
			}
			сам.Коорд.СтрПозУст(сам.Коорд.СтрПоз() + мПоз.ССтрПоз(len(слово)))
			return true, nil
		}
		// это что-то строковое
		начИмя := (сам.лит.Получ() == "_" || сам.лит.ЕслиБуква()) && !(сам.лит.Получ() == ".")
		if начИмя {
			return _Имя()
		}
		if сам.лит.ЕслиЦифра() || сам.лит.Получ() == "." { //Это число
			return _Число()
		}
		return false, nil
	}
	_Равно := func() (ок bool, ош error) {
		if сам.лит.Получ() == "=" {
			if ош = сам._СловоДобав("="); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Равно(): ОШИБКА при добавлении равно\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Точка := func() (ок bool, ош error) {
		if сам.лит.Получ() == "." {
			if ош = сам._СловоДобав("."); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Точка(): ОШИБКА при добавлении равно\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Кавычка2 := func() (ок bool, ош error) { // вычисляет строки
		if сам.лит.Получ() == "\"" {
			слово := мСлово.ССлово("\"")
			for сам.лит.Получ() != "\"" { //Перебирать, пока не встретится вторая кавычка
				сам._ЛитШаг()
				пЛит, ош := сам.исхФайл.Лит(сам.литНомер)
				if ош != nil {
					return false, мФмт.Errorf("Кавычка2(): ОШИБКА при получении литеры\n\t%v", ош)
				}
				слово += мСлово.ССлово(пЛит)
			}
			слово += мСлово.ССлово(сам.лит.Получ())
			сам._ЛитШаг()
			if ош = сам._СловоДобав(слово); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Кавычка2(): ОШИБКА при добавлении кавычек\n\t%v", ош)
			}
			return true, nil
		}
		return false, nil
	}

	//сам.конс.Отладить("ТИсходник.__Слово_Выделить()")
	лит := мЛит.СЛит("")
	ок := false
	if сам.литНомер < мПоз.ССтрПоз(сам.исхФайл.Размер()-1) {
		лит, ош = сам.исхФайл.Лит(сам.литНомер)
		if ош != nil {
			return мФмт.Errorf("_СловоВыделить(): ОШИБКА при получении слова\n\t%v", ош)
		}
		if ош = сам.лит.Уст(лит); ош != nil {
			return мФмт.Errorf("_СловоВыделить(): ОШИБКА при получении слова\n\t%v", ош)
		}
		//мФмт._rintf(" Лит=[%v] ", лит)
		if _Пробел() {
			return nil
		}
		if _НоваяСтрока() {
			return nil
		}
		if _ПереводКаретки() {
			return nil
		}
		if ок, ош = _Запятая(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ОШИБКА при выделении запятой\n\t%v", ош)
		}

		if ок, ош = _ТочкаЗапятая(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ОШИБКА при выделении точки с запятой\n\t%v", ош)
		}

		if ок, ош = _Плюс(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ОШИБКА при выделении плюса\n\t%v", ош)
		}

		if ок, ош = _Минус(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ОШИБКА при выделении минуса\n\t%v", ош)
		}

		if ок, ош = _Деление(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ОШИБКА при выделении деления\n\t%v", ош)
		}

		if ок, ош = _ЛеваяСкобка(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ОШИБКА при выделении левой скобки\n\t%v", ош)
		}

		if ок, ош = _ПраваяСкобка(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ОШИБКА при выделении правой скобки\n\t%v", ош)
		}

		if ок, ош = _Умножить(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ОШИБКА при выделении умножить\n\t%v", ош)
		}

		if ок, ош = _Двоеточие(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ОШИБКА при выделении двоеточия\n\t%v", ош)
		}

		if ок, ош = _Равно(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ОШИБКА при выделении равно\n\t%v", ош)
		}
		if ок, ош = _Точка(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ОШИБКА при выделении точки\n\t%v", ош)
		}

		if ок, ош = _ЕслиСлово(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ОШИБКА при выделении слова\n\t%v", ош)
		}

		if ок, ош = _Кавычка2(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ОШИБКА при выделении кавычек\n\t%v", ош)
		}
		//сам._ЛитШаг()
	} else {
		сам.бСловаГотово = true
	}
	return nil
}

//Слова -- возвращает все распознанные слова в исходнике
func (сам *ТИсходник) Слова() map[мСлово.ССловоНомер]мСлово.ИСлово {
	return сам.слова
}
