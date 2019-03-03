package пакИсходник

/*
	Пакет предоставляет тип для работы с исходным файлом.
	Позволяет текст считывать целиком, в другой части позволяет работать с
	отдельными строками по номерам.
*/

import (
	мКонс "../../../пакКонсоль"
	мРес "../../../пакРесурс"
	мКоорд "../../пакСущность/пакКоорд"
	мСлово "../../пакСущность/пакСлово"
	мИсхСтр "./пакИсходникСтроки"
	мИсхФайл "./пакИсходникФайл"
	мФмт "fmt"
)

//ТИсходник -- тип для считывания исходника и разбиения его на строки
type ТИсходник struct {
	ИсхФайл      *мИсхФайл.ТИсхФайл // содержимое файла исходника целиком
	ИсхСтр       *мИсхСтр.ТИсхСтр   // исходник построчно
	Коорд        *мКоорд.ТКоордИзм  // изменяемые координаты
	СловаМодуля  []*мСлово.ТСлово   // срез слов
	бСловаГотово bool               // признак готовности на разделения на слова
	цСтр         int                // текущая строка в исходном тексте
	цПоз         int                // текущая позиция в исходном тексте
}

//Новый -- возвращает новый экземпляр типа для чтения исходника
func Новый() (исх *ТИсходник, ош error) {
	мКонс.Конс.Отладить("пакИсходник.Новый()")
	исх = new(ТИсходник)
	исх.ИсхФайл = мИсхФайл.Новый()
	исх.ИсхСтр = мИсхСтр.Новый()
	исх.цСтр = 1 // Строки начинаются с 1
	исх.Коорд, ош = мКоорд.НовыйИзм(мКоорд.ТЦелСтр(исх.цСтр), мКоорд.ТЦелПоз(исх.цПоз))
	if ош != nil {
		return nil, мФмт.Errorf("пакТсходник.Новый(): ошибка при добавлении координат\n\t%v", ош)
	}
	return исх, nil
}

//Обработать -- главный цикл чтения и разбиения на строки исходника
func (сам *ТИсходник) Обработать(пИмяФайла string) (ош error) {
	мКонс.Конс.Отладить("ТИсходник.Обработать()")
	сам.ИсхФайл.Считать(пИмяФайла)
	//сам.конс.Отладить("ffff @" + string(сам.ИсхФайл.Исходник()) + "@")
	сам.ИсхСтр.НаСтрокиРазбить(сам.ИсхФайл.Исходник())
	if ош = сам._НаСловаРазделить(); ош != nil {
		return мФмт.Errorf("ТИсходник.Обработать(): ошибка при обработке исходника\n\t%v", ош)
	}
	return nil
}

// Добавляет слово с атрибутами положения в исходном тексте
// Строки исходника остаются как были.
func (сам *ТИсходник) _СловоДобав(пСлово string) (ош error) {
	крдФикс, _ := мКоорд.НовыйФикс(сам.Коорд.Стр(), сам.Коорд.Поз())
	if слово, ош := мСлово.Новое(крдФикс, пСлово, сам.ИсхСтр.Строка(int(сам.Коорд.Стр()))); ош == nil {
		if слово == nil {
			мКонс.Конс.Отладить("________  ТуСлово =   nil     !!!")
			return мФмт.Errorf("ТИсходник._СловоДобав(): ошибка при создании ТуСлово\n\t%v", ош)
		}
		сам.Коорд.ПозДоб()
		сам.СловаМодуля = append(сам.СловаМодуля, слово)
		//сам.конс.Отладить("Слово № " + мФмт.Sprintf("%d", слово.Номер()) + "=\"" + пСлово + "\"")
	} else {
		return мФмт.Errorf("ТуИсходник._СловоДобав(): ошибка при добавлении ТуСлово\n\t%v", ош)
	}
	return nil
}

// Разделяет исходник на слова
func (сам *ТИсходник) _НаСловаРазделить() (ош error) {
	мКонс.Конс.Отладить("ТуИсходник._НаСловаРазделить()")
	for {
		if сам.бСловаГотово {
			мКонс.Конс.Отладить("ТИсходник._СловоДобав(): всего слов: " + мФмт.Sprintf("%d", len(сам.СловаМодуля)))
			break
		} else {
			сам._СловоВыделить()
		}
	}
	return nil
}

// Выделяет слово из исходного текста
func (сам *ТИсходник) _СловоВыделить() {
	Пробел := func(пЛит string) {
		if пЛит == " " {
			сам.Коорд.ПозДоб()
		} else if пЛит == "\t" {
			for i := 0; i < мРес.ШирТаба; i++ {
				сам.Коорд.ПозДоб()
			}
		}
	}
	Запятая := func(пЛит string) {
		if пЛит == "," {
			сам._СловоДобав(",")
		}
	}
	ТочкаЗапятая := func(пЛит string) {
		if пЛит == ";" {
			сам._СловоДобав(";")
		}
	}
	Плюс := func(пЛит string) {
		if пЛит == "+" {
			сам._СловоДобав("+")
		}
	}
	Минус := func(пЛит string) {
		if пЛит == "-" {
			сам._СловоДобав("-")
		}
	}
	Деление := func(пЛит string) {
		if пЛит == "/" {
			сам._СловоДобав("/")
		}
	}
	ЛеваяСкобка := func(пЛит string) {
		if пЛит == "(" {
			if пЛит+сам.ИсхФайл.Лит(сам.цПоз+1) != "(*" {
				сам._СловоДобав("(")
			} else {
				сам._СловоДобав("(*")
				сам.Коорд.ПозДоб()
				сам.цПоз++
			}
		}
	}
	ПраваяСкобка := func(пЛит string) {
		if пЛит == ")" {
			сам._СловоДобав(")")
		}
	}
	НоваяСтрока := func(пЛит string) {
		if пЛит == "\n" {
			сам.Коорд.СтрДоб()
			сам.Коорд.ПозСброс()
		}
	}
	Умножить := func(пЛит string) {
		if пЛит == "*" {
			if пЛит+сам.ИсхФайл.Лит(сам.цПоз+1) != "*)" {
				сам._СловоДобав("*")
			} else {
				сам._СловоДобав("*)")
				сам.Коорд.ПозДоб()
				сам.цПоз++
			}
		}
	}
	Двоеточие := func(пЛит string) {
		if пЛит == ":" {
			if пЛит+сам.ИсхФайл.Лит(сам.цПоз+1) != ":=" {
				сам._СловоДобав(":")
			} else {
				сам._СловоДобав(":=")
				сам.Коорд.ПозДоб()
				сам.цПоз++
			}
		}
	}
	ПереводКаретки := func(пЛит string) {
		if пЛит == "\r" {
			сам.Коорд.ПозДоб()
			//сам.__цПоз++
		}
	}
	ЕслиСущность := func(пЛит string) {
		сущн := ""
		// это что-то строковое
		if пЛит == "_" || мСлово.ЕслиБуква(пЛит) {
			var бДальше bool
			for {
				бДальше = (пЛит != ".") && (мСлово.ЕслиБуква(пЛит) || пЛит == "_" || мСлово.ЕслиЦифра(пЛит))
				if бДальше {
					сущн += пЛит
					сам.цПоз++
					сам.Коорд.ПозДоб()
					пЛит = сам.ИсхФайл.Лит(сам.цПоз)
				} else {
					// откат на одну позицию
					сам.цПоз--
					сам.Коорд.ПозУст(сам.Коорд.Поз() - мКоорд.ТЦелПоз(len(сущн)))
					if сущн == "." {
						мКонс.Конс.Отладить("________  сущн_стр =   ТОЧКА     !!!")
					}
					//сам.конс.Отладить("Сущность=\"" + сущн + "\"")
					сам._СловоДобав(сущн)
					сам.Коорд.ПозУст(сам.Коорд.Поз() + мКоорд.ТЦелПоз(len(сущн)) - 1)
					break
				}
			}
			// Это что-то числовое
		} else if мСлово.ЕслиЦифра(пЛит) {
			var бДальше bool
			for {
				бДальше = мСлово.ЕслиЦифра(пЛит)
				if бДальше {
					сущн += пЛит
					сам.цПоз++
					сам.Коорд.ПозДоб()
					пЛит = сам.ИсхФайл.Лит(сам.цПоз)
				} else {
					// откат на одну позицию
					сам.цПоз--
					if сущн != "." {
						сам.Коорд.ПозУст(сам.Коорд.Поз() - мКоорд.ТЦелПоз(len(сущн)))
						if сущн == "." {
							мКонс.Конс.Отладить("________  сущн_цифра =   ТОЧКА     !!!")
							//сам.конс.Отладить("Сущность=\"" + сущн + "\"")
						}
						сам._СловоДобав(сущн)
						сам.Коорд.ПозУст(сам.Коорд.Поз() + мКоорд.ТЦелПоз(len(сущн)) - 1)
					}
					break
				}
			}
		}
	}
	Равно := func(пЛит string) {
		if пЛит == "=" {
			сам._СловоДобав("=")
		}
	}
	Точка := func(пЛит string) {
		if пЛит == "." {
			сам._СловоДобав(".")
		}
	}
	Кавычка2 := func(пЛит string) { // вычисляет строки
		if пЛит == "\"" {
			стр := "\""
			пЛит = ""
			for {
				if пЛит != "\"" {
					сам.цПоз++
					сам.Коорд.ПозДоб()
					пЛит = сам.ИсхФайл.Лит(сам.цПоз)
					стр += пЛит
				} else {
					break
				}
			}
			сам._СловоДобав(стр)
			сам.Коорд.ПозДоб()
			сам.цПоз++
		}
	}

	//сам.конс.Отладить("ТуИсходник.__Слово_Выделить()")
	цИсхДлина := len([]rune(сам.ИсхФайл.Исходник())) - 1
	лит := ""
	if сам.цПоз < цИсхДлина {
		лит = сам.ИсхФайл.Лит(сам.цПоз)
		//сам.конс.Отладить("Лит=" + лит)
		Пробел(лит)
		Запятая(лит)
		ТочкаЗапятая(лит)
		Плюс(лит)
		Минус(лит)
		Деление(лит)
		ЛеваяСкобка(лит)
		ПраваяСкобка(лит)
		НоваяСтрока(лит)
		Умножить(лит)
		Двоеточие(лит)
		ПереводКаретки(лит)
		Равно(лит)
		Точка(лит)
		ЕслиСущность(лит)
		Кавычка2(лит)
		сам.цПоз++
	} else {
		сам.бСловаГотово = true
	}
}
