// модСлово
package пакСлово

/*
Предоставляет тип слова для построения AST.
Слово -- кусочек текста в исходнике.
Обладает несколькими свойствами:
1. Группа литер (само слово, может быть из одной литеры)
2. Координаты.
3. Строка, в которой он находится.
*/

import (
	пакКонс "../../../пакКонсоль"
	пакКоорд "../../пакСущность/пакКоорд"
	пакТипы "../пакТипы"
	пакФмт "fmt"
	"strings"
)

var слов_всего = 0

type ТуСлово struct {
	_стрИсх *string               // Строка исходника
	_строка string                // Строка слова
	род     пакТипы.ТРод          // род слова
	конс    *пакКонс.ТуКонсоль    // Системная консоль
	Коорд   *пакКоорд.ТуКоордФикс // Координаты слова в исходном тексте
	номер   int                   // Текущий номер слова
}

func Новое(пНомСтр, пНомПоз int, пСтрока string) (слово *ТуСлово, ош error) {
	//Проверяет входящую строку на минимальную длину.
	слово_проверить := func(пСтрока string) (ош error) { // Вспомогательная функция
		if len(пСтрока) == 0 {
			ош = пакФмт.Errorf(" Строка имеет длину 0")
			return ош
		}
		return ош
	}

	род_проверить := func(сам *ТуСлово, пСтрока string) (род пакТипы.ТРод, ош error) { //Устанавливает род слова.
		род = -1
		switch {
		case string(пСтрока) == ";":
			род = КТочкаЗапятая
		case string(пСтрока) == ",":
			род = КЗапятая
		case string(пСтрока) == "+":
			род = КПлюс
		case string(пСтрока) == "-":
			род = КМинус
		case string(пСтрока) == "/":
			род = КДеление
		case string(пСтрока) == "(":
			род = КСкобкаОткрКругл
		case string(пСтрока) == "(*":
			род = ККомментНачать
		case string(пСтрока) == ")":
			род = КСкобкаЗакрКругл
		case string(пСтрока) == "*)":
			род = ККомментЗакончить
		case string(пСтрока) == "*":
			род = КУмножить
		case string(пСтрока) == ":=":
			род = КПрисвоить
		case string(пСтрока) == ":":
			род = КОпределить
		case strings.HasPrefix(пСтрока, "_") || ЕслиПерваяБуква(пСтрока):
			род = КИмя
		case ЕслиЦифра(string(пСтрока[0])):
			род = КЧисло
		case string(пСтрока) == "=":
			род = КРавно
		case string(пСтрока) == ".":
			род = КТочка
		case пСтрока[0] == '"' && пСтрока[len(пСтрока)-1] == '"':
			род = КСтрока
		default:
			ош = пакФмт.Errorf("Не могу классифицировать строку, строка=" + string(пСтрока))
		}
		//сам.Род = пакТипы.ТРод(род)
		return пакТипы.ТРод(род), ош
	}

	if ош1 := слово_проверить(пСтрока); ош != nil {
		ош1 = пакФмт.Errorf("ТуСлово.Новое(): при вызове слово_проверить()\n\t%v", ош1)
		panic(ош1)
	} else {
		if род, ош2 := род_проверить(слово, пСтрока); ош2 != nil {
			ош2 = пакФмт.Errorf("пакСлово.Новое(): при вызове род_проверить()\n\t%v", ош1)
			panic(ош2)
		} else {
			if коорд, ош3 := пакКоорд.НовыйФикс(пакКоорд.ТЦелСтр(пНомСтр), пакКоорд.ТЦелПоз(пНомПоз)); ош3 != nil {
				ош3 = пакФмт.Errorf("пакСлово.Новое(): ошибка при создании ТуКоорд\n %v", ош2)
				panic(ош3)
			} else {
				слов_всего++
				слово = &ТуСлово{
					_строка: пСтрока,
					конс:    пакКонс.Конс,
					Коорд:   коорд,
					род:     род,
					номер:   слов_всего,
				}
				//слово.конс.Отладить("ТуСлово:" + пСтрока + пакФмт.Sprintf(" %d", слово.номер))
			}
		}
	}
	return слово, ош
}

func (сам *ТуСлово) Род() int {
	return int(сам.род)
}

func (сам *ТуСлово) Номер() int {
	return сам.номер
}

// Проверяет наличие первой буквы в строке
func ЕслиПерваяБуква(пСтрока string) (бРез bool) {
	руна := []rune(пСтрока)
	стрЛит := string(руна[0])
	//пакКонс.Конс.Печать("пакСлово.ЕслиПерваяБуква(): лит=" + стрЛит)
	if ЕслиБуква(стрЛит) {
		бРез = true
	}
	return бРез
}

// Проверяет наличие буквы в литере
func ЕслиБуква(пЛит string) (бРез bool) {
	if strings.Contains(стрБуквыВсе, пЛит) {
		бРез = true
	}
	return бРез
}

func ЕслиЦифра(пЛит string) (бРез bool) {
	стрЦифры := "0123456789."
	if strings.Contains(стрЦифры, пЛит) {
		бРез = true
	}
	return бРез
}

func ЕслиНеЦифра(пСтрока string) (бРез bool) {
	стрЦифры := "0123456789."
	for лит := range пСтрока {
		if !(strings.Contains(стрЦифры, string(лит))) {
			бРез = true
		}
	}
	return бРез
}

// Возвращает слово, которое хранит тип
func (сам *ТуСлово) Строка() (строка string) {
	if сам._строка == "" {
		panic(пакФмт.Sprintf("ТуСлово.Строка(): ТуСтрока._строка не может бытьпустой"))
	}
	return сам._строка
}

// Проверяет, что строка не находится в запрещённых словах (ключевые)
func __ЕслиИмя_Запрещено(пСтрока string) (бРез bool) {
	бРез = false
	for _, запр := range запрет_имя {
		for _, запр1 := range запр {
			if пСтрока == запр1 {
				бРез = true
				return бРез
			}
		}
	}
	return бРез
}

// Проверяет, что литера не находится в списке запрещённых (для имён сущностей)
func __ЕслиЛит_Запрещена(пЛит rune) (бРез bool) {
	//стрЗапрет:=[]rune("")
	стрЗапрет := []rune("\"~`!@$%^&*()-_=+{}[]|\\<,>?/№;:\t\n'\r ")
	for _, лит := range стрЗапрет {
		if пЛит == лит {
			бРез = true
			return бРез
		}
	}
	return бРез
}

// Проверяет, что строка обладает строгим соответствием ,чтобы быть именем сущности
func (сам *ТуСлово) ЕслиИмя_Строго() (бРез bool) {
	// имя сущности должно начинаться либо с "_", либо с буквы
	бРез = true
	строка := []rune(сам._строка)
	лит := string(строка[0])
	if лит == "_" || ЕслиБуква(лит) {
		if __ЕслиИмя_Запрещено(сам._строка) {
			бРез = false
			return бРез
		}
		for _, лит := range строка {
			// Точка в имени -- допустимо, но здесь её не будет.
			if __ЕслиЛит_Запрещена(лит) {
				//
				бРез = false
				return бРез
			}
		}
	}
	return бРез
}

// Возвращает строку исходника, содержащую подстроку
func (сам *ТуСлово) СтрИсх() (стр string) {
	if *сам._стрИсх == "" {
		panic("ТуСлово.СтрИсх(): строка исходника не может бытьпустой")
	}
	return *сам._стрИсх
}
