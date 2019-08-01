package пакИсходник

/*
	Пакет предоставляет тип для работы с исходным файлом.
	Позволяет текст считывать целиком, в другой части позволяет работать с
	отдельными строками по номерам.
*/

import (
	мФмт "fmt"

	мИнт "github.com/prospero78/goOC/пакОберон/пакИнтерфейсы"
	мИсхСтр "github.com/prospero78/goOC/пакОберон/пакИсходникСтроки"
	мИсхФайл "github.com/prospero78/goOC/пакОберон/пакИсходникФайл"
	мКонс "github.com/prospero78/goOC/пакОберон/пакКонсоль"
	мКоорд "github.com/prospero78/goOC/пакОберон/пакКоорд"
	мЛит "github.com/prospero78/goOC/пакОберон/пакЛитера"
	мРес "github.com/prospero78/goOC/пакОберон/пакРесурс"
	мСлово "github.com/prospero78/goOC/пакОберон/пакСлово"
)

//ИИсходник -- интерфейс для исходника
type ИИсходник interface {
	Обработать(СИсхФайл) error
	Слова() map[ССловоНомерМодуль]ИСлово
}

//ТИсходник -- тип для считывания исходника и разбиения его на строки
type ТИсходник struct {
	ИсхФайл      мИнт.ИИсхФайл                          // содержимое файла исходника целиком
	ИсхСтр       мИнт.ИИсхСтроки                        // исходник построчно
	Коорд        мКоорд.ИКоордИзм                       // изменяемые координаты
	слова        map[мИнт.ССловоНомерМодуль]мИнт.ИСлово // справочник слов
	бСловаГотово bool                                   // признак готовности на разделения на слова
	цСтр         мИнт.СКоордСтр                         // текущая строка в исходном тексте
	лит          мИнт.ИЛит                              //текущая литера
	литНомер     мИнт.СЛитНомер                         // текущая позиция в исходном тексте
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
	исх.слова = make(map[мИнт.ССловоНомерМодуль]мИнт.ИСлово)
	if исх.лит, ош = мЛит.ЛитераНов(); ош != nil {
		return nil, мФмт.Errorf("пакТсходник.Новый(): ошибка при создании объекта литеры\n\t%v", ош)
	}
	return исх, nil
}

//Обработать -- главный цикл чтения и разбиения на строки исходника
func (сам *ТИсходник) Обработать(пИмяФайла мИнт.СИсхФайл) (ош error) {
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
func (сам *ТИсходник) _СловоДобав(пСлово мИнт.ССлово) (ош error) {
	//мФмт._rintf("%v %v-%v %v\n", len(сам.слова), сам.Коорд.Стр(), сам.Коорд.Поз(), пСлово)
	крдФикс, ош := мКоорд.КоордНов(сам.Коорд.СтрНомер(), сам.Коорд.СтрПоз())
	if ош != nil {
		return мФмт.Errorf("ТИсходник._СловоДобав(): ошибка при создании координат слова\n\t%v", ош)
	}
	слово, ош := мСлово.СловоНов(крдФикс, пСлово, сам.ИсхСтр.Строка(сам.Коорд.СтрНомер()))
	if ош != nil {
		return мФмт.Errorf("ТИсходник._СловоДобав(): ошибка при добавлении ТСлово\n\t%v", ош)
	}
	сам.Коорд.ПозДоб()
	сам.слова[мИнт.ССловоНомерМодуль(len(сам.слова))] = слово
	//сам.конс.Отладить("Слово № " + мФмт.S_rintf("%d", слово.Номер()) + "=\"" + пСлово + "\"")

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
	_Пробел := func() (ок bool) {
		if сам.лит.Лит() == " " {
			сам._ЛитШаг()
			return true
		} else if сам.лит.Лит() == "\t" { //Позицию добавить на столько, сколько пробелов считается таб
			сам.литНомер++
			for i := 0; i < мРес.ШирТаба; i++ {
				сам.Коорд.ПозДоб()
			}
			return true
		}
		return false
	}
	_ТочкаЗапятая := func() (ок bool, ош error) {
		if сам.лит.Лит() == ";" {
			if ош = сам._СловоДобав(";"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._ТочкаЗапятая(): ошибка при добавлении точки с запятой\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_НоваяСтрока := func() (ок bool) {
		if сам.лит.Лит() == "\n" {
			сам.Коорд.СтрДоб()
			сам.Коорд.ПозСброс()
			сам.литНомер++
			return true
		}
		return false
	}
	_Запятая := func() (ок bool, ош error) {
		if сам.лит.Лит() == "," {
			if ош = сам._СловоДобав(","); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Запятая(): ошибка при добавлении запятой\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Плюс := func() (ок bool, ош error) {
		if сам.лит.Лит() == "+" {
			if ош = сам._СловоДобав("+"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Запятая(): ошибка при добавлении плюса\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Минус := func() (ок bool, ош error) {
		if сам.лит.Лит() == "-" {
			if ош = сам._СловоДобав("-"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Минус(): ошибка при добавлении минуса\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Деление := func() (ок bool, ош error) {
		if сам.лит.Лит() == "/" {
			if ош = сам._СловоДобав("/"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Деление(): ошибка при добавлении деления\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_ЛеваяСкобка := func() (ок bool, ош error) {
		if сам.лит.Лит() == "(" {
			лит1, ош := сам.ИсхФайл.Лит(сам.литНомер + 1)
			if ош != nil {
				return false, мФмт.Errorf("ЛеваяСкобка(): ошибка при получении литеры\n\t%v", ош)
			}
			if сам.лит.Лит()+лит1 == "(*" {
				if ош = сам._СловоДобав("(*"); ош != nil {
					return false, мФмт.Errorf("ТИсходник._СловоВыделить()._ЛеваяСкобка(): ошибка при добавлении открытия комментария\n\t%v", ош)
				}
				сам._ЛитШаг()
				сам._ЛитШаг()
				return true, nil
			}
			if ош = сам._СловоДобав("("); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._ЛеваяСкобка(): ошибка при добавлении левой скобки\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_ПраваяСкобка := func() (ок bool, ош error) {
		if сам.лит.Лит() == ")" {
			if ош = сам._СловоДобав(")"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._ПраваяСкобка(): ошибка при добавлении правой скобки\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Умножить := func() (ок bool, ош error) {
		if сам.лит.Лит() == "*" {
			лит1, ош := сам.ИсхФайл.Лит(сам.литНомер + 1)
			if ош != nil {
				return false, мФмт.Errorf("Умножить(): ошибка при получении литеры\n\t%v", ош)
			}
			if сам.лит.Лит()+лит1 == "*)" {
				if ош = сам._СловоДобав("*)"); ош != nil {
					return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Умножить(): ошибка при добавлении закрытия комментария\n\t%v", ош)
				}
				сам._ЛитШаг()
				сам._ЛитШаг()
				return true, nil
			}
			if ош = сам._СловоДобав("*"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Умножить(): ошибка при добавлении умножить\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Двоеточие := func() (ок bool, ош error) {
		if сам.лит.Лит() == ":" {
			лит1, ош := сам.ИсхФайл.Лит(сам.литНомер + 1)
			if ош != nil {
				return false, мФмт.Errorf("Двоеточие(): ошибка при получении литеры\n\t%v", ош)
			}
			if сам.лит.Лит()+лит1 == ":=" {
				if ош = сам._СловоДобав(":="); ош != nil {
					return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Двоеточие(): ошибка при добавлении присовить\n\t%v", ош)
				}
				сам._ЛитШаг()
				сам._ЛитШаг()
				return true, nil
			}
			if ош = сам._СловоДобав(":"); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Двоеточие(): ошибка при добавлении двоеточия\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_ПереводКаретки := func() bool {
		if сам.лит.Лит() == "\r" {
			сам._ЛитШаг()
			return true
		}
		return false
	}
	_ЕслиСлово := func() (ок bool, ош error) {
		_Имя := func() (ок bool, ош error) {
			слово := мИнт.ССлово("")
			for {
				пЛит, ош := сам.ИсхФайл.Лит(сам.литНомер)
				if ош != nil {
					return false, мФмт.Errorf("ЕслиСущность(): ошибка при получении литеры\n\t%v", ош)
				}
				if ош = сам.лит.Уст(пЛит); ош != nil {
					return false, мФмт.Errorf("ЕслиСущность(): ошибка при установке литеры\n\t%v", ош)
				}
				бПродИмя := (сам.лит.ЕслиБуква() || сам.лит.ЕслиЦифра() || сам.лит.ЕслиЗнаки()) && !(сам.лит.Лит() == ".")
				if !бПродИмя {
					break
				}
				слово += мИнт.ССлово(сам.лит.Лит())
				сам._ЛитШаг()
			}
			// откат на одну позицию
			//сам.литНомер--
			if ош = сам._СловоДобав(слово); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._ЕслиСлово(): ошибка при добавлении имени\n\t%v", ош)
			}
			сам.Коорд.ПозУст(сам.Коорд.Поз() + мИнт.СКоордПоз(len(слово)))
			return true, nil
		}
		_Число := func() (ок bool, ош error) {
			слово := мИнт.ССлово("")
			for {
				пЛит, ош := сам.ИсхФайл.Лит(сам.литНомер)
				if ош != nil {
					return false, мФмт.Errorf("ЕслиСущность(): ошибка при получении литеры\n\t%v", ош)
				}
				if ош = сам.лит.Уст(пЛит); ош != nil {
					return false, мФмт.Errorf("ЕслиСущность(): ошибка при установке литеры\n\t%v", ош)
				}
				if !(сам.лит.ЕслиЦифра() || сам.лит.Лит() == ".") {
					break
				}
				слово += мИнт.ССлово(сам.лит.Лит())
				сам._ЛитШаг()
			}
			// откат на одну позицию
			//сам.литНомер--
			if ош = сам._СловоДобав(слово); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._ЕслиСлово(): ошибка при добавлении числа\n\t%v", ош)
			}
			сам.Коорд.ПозУст(сам.Коорд.Поз() + мИнт.СКоордПоз(len(слово)))
			return true, nil
		}
		// это что-то строковое
		начИмя := (сам.лит.Лит() == "_" || сам.лит.ЕслиБуква()) && !(сам.лит.Лит() == ".")
		if начИмя {
			return _Имя()
		}
		if сам.лит.ЕслиЦифра() || сам.лит.Лит() == "." { //Это число
			return _Число()
		}
		return false, nil
	}
	_Равно := func() (ок bool, ош error) {
		if сам.лит.Лит() == "=" {
			if ош = сам._СловоДобав("="); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Равно(): ошибка при добавлении равно\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Точка := func() (ок bool, ош error) {
		if сам.лит.Лит() == "." {
			if ош = сам._СловоДобав("."); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Точка(): ошибка при добавлении равно\n\t%v", ош)
			}
			сам._ЛитШаг()
			return true, nil
		}
		return false, nil
	}
	_Кавычка2 := func() (ок bool, ош error) { // вычисляет строки
		if сам.лит.Лит() == "\"" {
			слово := мИнт.ССлово("\"")
			for сам.лит.Лит() != "\"" { //Перебирать, пока не встретится вторая кавычка
				сам._ЛитШаг()
				пЛит, ош := сам.ИсхФайл.Лит(сам.литНомер)
				if ош != nil {
					return false, мФмт.Errorf("Кавычка2(): ошибка при получении литеры\n\t%v", ош)
				}
				слово += мИнт.ССлово(пЛит)
			}
			слово += мИнт.ССлово(сам.лит.Лит())
			сам._ЛитШаг()
			if ош = сам._СловоДобав(слово); ош != nil {
				return false, мФмт.Errorf("ТИсходник._СловоВыделить()._Кавычка2(): ошибка при добавлении кавычек\n\t%v", ош)
			}
			return true, nil
		}
		return false, nil
	}

	//сам.конс.Отладить("ТИсходник.__Слово_Выделить()")
	лит := мИнт.СЛит("")
	ок := false
	if сам.литНомер < мИнт.СЛитНомер(сам.ИсхФайл.Размер()-1) {
		лит, ош = сам.ИсхФайл.Лит(сам.литНомер)
		if ош != nil {
			return мФмт.Errorf("_СловоВыделить(): ошибка при получении слова\n\t%v", ош)
		}
		if ош = сам.лит.Уст(лит); ош != nil {
			return мФмт.Errorf("_СловоВыделить(): ошибка при получении слова\n\t%v", ош)
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
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении запятой\n\t%v", ош)
		}

		if ок, ош = _ТочкаЗапятая(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении точки с запятой\n\t%v", ош)
		}

		if ок, ош = _Плюс(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении плюса\n\t%v", ош)
		}

		if ок, ош = _Минус(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении минуса\n\t%v", ош)
		}

		if ок, ош = _Деление(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении деления\n\t%v", ош)
		}

		if ок, ош = _ЛеваяСкобка(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении левой скобки\n\t%v", ош)
		}

		if ок, ош = _ПраваяСкобка(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении правой скобки\n\t%v", ош)
		}

		if ок, ош = _Умножить(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении умножить\n\t%v", ош)
		}

		if ок, ош = _Двоеточие(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении двоеточия\n\t%v", ош)
		}

		if ок, ош = _Равно(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении равно\n\t%v", ош)
		}
		if ок, ош = _Точка(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении точки\n\t%v", ош)
		}

		if ок, ош = _ЕслиСлово(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении слова\n\t%v", ош)
		}

		if ок, ош = _Кавычка2(); ош == nil {
			if ок {
				return nil
			}
		} else {
			return мФмт.Errorf("ТИсходник._СловоВыделить(): ошибка при выделении кавычек\n\t%v", ош)
		}
		//сам._ЛитШаг()
	} else {
		сам.бСловаГотово = true
	}
	return nil
}

//Слова -- возвращает все распознанные слова в исходнике
func (сам *ТИсходник) Слова() map[мИнт.ССловоНомерМодуль]мИнт.ИСлово {
	return сам.слова
}
