// модИсходник
package пакИсходник

/*
Пакет предоставляет тип для работы с исходным файлом.
Позволяет текст считывать целиком, в другой части позволяет работать с
отдельными строками по номерам.
*/
import (
	пакКонс "../../../пакКонсоль"
	пакРес "../../../пакРесурс"
	пакКоорд "../../пакСущность/пакКоорд"
	пакСлово "../../пакСущность/пакСлово"
	пакИсхСтр "./пакИсходникСтроки"
	пакИсхФайл "./пакИсходникФайл"
	пакФмт "fmt"
)

type ТуИсходник struct {
	конс         *пакКонс.ТуКонсоль    // ссылка на глобальную консоль
	ИсхФайл      *пакИсхФайл.ТуИсхФайл // содержимое файла исходника целиком
	ИсхСтр       *пакИсхСтр.ТуИсхСтр   // исходник построчно
	Коорд        *пакКоорд.ТуКоордИзм  // изменяемые координаты
	СловаМодуля  []*пакСлово.ТуСлово   // срез слов
	бСловаГотово bool                  // признак готовности на разделения на слова
	__цСтр       int                   // текущая строка в исходном тексте
	__цПоз       int                   // текущая позиция в исходном тексте
}

func Новый() (исх *ТуИсходник, ош error) {
	пакКонс.Конс.Отладить("пакИсходник.Новый()")
	исх = new(ТуИсходник)
	исх.конс = пакКонс.Конс
	исх.ИсхФайл = пакИсхФайл.Новый()
	исх.ИсхСтр = пакИсхСтр.Новый()
	исх.__цСтр = 1 // Строки начинаются с 1
	исх.Коорд, ош = пакКоорд.НовыйИзм(пакКоорд.ТЦелСтр(исх.__цСтр), пакКоорд.ТЦелПоз(исх.__цПоз))
	return исх, ош
}

func (сам *ТуИсходник) Обработать(пИмяФайла string) {
	сам.конс.Отладить("ТуИсходник.Обработать()")
	сам.ИсхФайл.Считать(пИмяФайла)
	//сам.конс.Отладить("ffff @" + string(сам.ИсхФайл.Исходник()) + "@")
	сам.ИсхСтр.НаСтроки_Разбить(сам.ИсхФайл.Исходник())
	сам.__НаСлова_Разделить()
}

// Добавляет слово с атрибутами положения в исходном тексте
// Строки исходника остаются как были.
func (сам *ТуИсходник) __Слово_Добав(пСлово string) (ош error) {
	сам.конс.Отладить("Слово=" + пСлово)
	слово, _ош := пакСлово.Новое(сам.__цСтр, сам.__цПоз, пСлово)
	if _ош != nil {
		ош = пакФмт.Errorf("ТуИсходник.__Слово_Добав(): ошибка при добавлении ТуСлово\n%v", _ош)
	}
	сам.СловаМодуля = append(сам.СловаМодуля, слово)
	сам.Коорд.Поз_Доб()
	return ош
}

// Разделяет исходник на слова
func (сам *ТуИсходник) __НаСлова_Разделить() {
	сам.конс.Отладить("ТуИсходник.__НаСлова_Разделить()")
	for {
		if сам.бСловаГотово {
			break
		} else {
			сам.__Слово_Выделить()
		}
	}
}

// Выделяет слово из исходного текста
func (сам *ТуИсходник) __Слово_Выделить() {
	Пробел := func(пЛит string) {
		if пЛит == " " {
			сам.Коорд.Поз_Доб()
		} else if пЛит == "\t" {
			for i := 0; i < пакРес.ШирТаба; i++ {
				сам.Коорд.Поз_Доб()
			}
		}
	}
	Запятая := func(пЛит string) {
		if пЛит == "," {
			сам.__Слово_Добав(",")
		}
	}
	ТочкаЗапятая := func(пЛит string) {
		if пЛит == ";" {
			сам.__Слово_Добав(";")
		}
	}
	Плюс := func(пЛит string) {
		if пЛит == "+" {
			сам.__Слово_Добав("+")
		}
	}
	Минус := func(пЛит string) {
		if пЛит == "-" {
			сам.__Слово_Добав("-")
		}
	}
	Деление := func(пЛит string) {
		if пЛит == "/" {
			сам.__Слово_Добав("/")
		}
	}
	ЛеваяСкобка := func(пЛит string) {
		if пЛит == "(" {
			if пЛит+сам.ИсхФайл.Лит(сам.__цПоз+1) != "(*" {
				сам.__Слово_Добав("(")
			} else {
				сам.__Слово_Добав("(*")
				сам.Коорд.Поз_Доб()
				сам.__цПоз++
			}
		}
	}
	ПраваяСкобка := func(пЛит string) {
		if пЛит == ")" {
			сам.__Слово_Добав(")")
		}
	}
	НоваяСтрока := func(пЛит string) {
		if пЛит == "\n" {
			сам.Коорд.Стр_Доб()
			сам.Коорд.Поз_Сброс()
		}
	}
	Умножить := func(пЛит string) {
		if пЛит == "*" {
			if пЛит+сам.ИсхФайл.Лит(сам.__цПоз) != "*)" {
				сам.__Слово_Добав("*")
			} else {
				сам.__Слово_Добав("*)")
				сам.Коорд.Поз_Доб()
				сам.__цПоз++
			}
		}
	}
	Двоеточие := func(пЛит string) {
		if пЛит == ":" {
			if пЛит+сам.ИсхФайл.Лит(сам.__цПоз) != ":=" {
				сам.__Слово_Добав(":")
			} else {
				сам.__Слово_Добав(":=")
				сам.Коорд.Поз_Доб()
				сам.__цПоз++
			}
		}
	}
	ПереводКаретки := func(пЛит string) {
		if пЛит == "\r" {
			сам.Коорд.Поз_Доб()
			сам.__цПоз++
		}
	}
	ЕслиСущность := func(пЛит string) {
		сущн := ""
		if пЛит == "_" || пакСлово.ЕслиБуква(пЛит) {
			var бДальше bool
			for {
				бДальше = пакСлово.ЕслиБуква(пЛит) || пакСлово.ЕслиЦифра(пЛит) || пЛит == "_"
				if бДальше {
					сущн += пЛит
					сам.__цПоз++
					сам.Коорд.Поз_Доб()
					пЛит = сам.ИсхФайл.Лит(сам.__цПоз)
				} else {
					// откат на одну позиция
					сам.__цПоз--
					сам.Коорд.Поз_Уст(сам.Коорд.Поз() - пакКоорд.ТЦелПоз(len(сущн)))
					сам.__Слово_Добав(сущн)
					сам.Коорд.Поз_Уст(сам.Коорд.Поз() + пакКоорд.ТЦелПоз(len(сущн)) - 1)
					break
				}
			}
		} else if пакСлово.ЕслиЦифра(пЛит) {
			var бДальше bool
			for {
				бДальше = пакСлово.ЕслиЦифра(пЛит)
				if бДальше {
					сущн += пЛит
					сам.__цПоз++
					сам.Коорд.Поз_Доб()
					пЛит = сам.ИсхФайл.Лит(сам.__цПоз)
				} else {
					// откат на одну позицию
					сам.Коорд.Поз_Уст(сам.Коорд.Поз() - пакКоорд.ТЦелПоз(len(сущн)))
					сам.__Слово_Добав(сущн)
					сам.Коорд.Поз_Уст(сам.Коорд.Поз() + пакКоорд.ТЦелПоз(len(сущн)) - 1)
					break
				}
			}
		}
	}
	Равно := func(пЛит string) {
		if пЛит == "=" {
			сам.__Слово_Добав("=")
		}
	}
	Точка := func(пЛит string) {
		if пЛит == "." {
			сам.__Слово_Добав(".")
		}
	}
	Кавычка2 := func(пЛит string) { // вычисляет строки
		if пЛит == "\"" {
			стр := "\""
			пЛит = ""
			for {
				if пЛит != "\"" {
					сам.__цПоз++
					сам.Коорд.Поз_Доб()
					пЛит = сам.ИсхФайл.Лит(сам.__цПоз)
					стр += пЛит
				} else {
					break
				}
			}
			сам.__Слово_Добав(стр)
			сам.Коорд.Поз_Доб()
			сам.__цПоз++
		}
	}

	//сам.конс.Отладить("ТуИсходник.__Слово_Выделить()")
	цИсхДлина := len(сам.ИсхФайл.Исходник()) - 1
	лит := ""
	if сам.__цПоз < цИсхДлина {
		лит = сам.ИсхФайл.Лит(сам.__цПоз)
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
		ЕслиСущность(лит)
		Равно(лит)
		Точка(лит)
		Кавычка2(лит)
		сам.__цПоз++
	} else {
		сам.бСловаГотово = true
	}
}
