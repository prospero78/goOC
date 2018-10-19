package пакСлово

// модСлово

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

var словВсего = 0

//ТуСлово -- тип хранящий слово со всеми необходимыми атрибутами
type ТуСлово struct {
	_стрИсх string                // Строка исходника
	_строка string                // Строка слова
	род     пакТипы.ТРод          // род слова
	конс    *пакКонс.ТуКонсоль    // Системная консоль
	Коорд   *пакКоорд.ТуКоордФикс // Координаты слова в исходном тексте
	номер   int                   // Текущий номер слова
}

//Новое -- возвращает нвое слово
func Новое(пКоорд пакКоорд.ИКоордФикс, пСтрока, пСтрИсх string) (слово *ТуСлово, ош error) {
	//Проверяет входящую строку на минимальную длину.
	словоПроверить := func(пСтрока string) (ош error) { // Вспомогательная функция
		if len(пСтрока) == 0 {
			ош = пакФмт.Errorf(" Строка имеет длину 0")
			return ош
		}
		return ош
	}

	родПроверить := func(сам *ТуСлово, пСтрока string) (род пакТипы.ТРод, ош error) { //Устанавливает род слова.
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

	if ош1 := словоПроверить(пСтрока); ош != nil {
		ош1 = пакФмт.Errorf("ТуСлово.Новое(): при вызове слово_проверить()\n\t%v", ош1)
		panic(ош1)
	} else {
		if род, ош2 := родПроверить(слово, пСтрока); ош2 != nil {
			ош2 = пакФмт.Errorf("пакСлово.Новое(): при вызове род_проверить()\n\t%v", ош1)
			panic(ош2)
		} else {
			if коорд, ош3 := пакКоорд.НовыйФикс(пКоорд.Стр(), пКоорд.Поз()); ош3 != nil {
				ош3 = пакФмт.Errorf("пакСлово.Новое(): ошибка при создании ТуКоорд\n %v", ош2)
				panic(ош3)
			} else {
				словВсего++
				слово = &ТуСлово{
					_строка: пСтрока,
					_стрИсх: пСтрИсх,
					конс:    пакКонс.Конс,
					Коорд:   коорд,
					род:     род,
					номер:   словВсего,
				}
				//слово.конс.Отладить("ТуСлово:" + пСтрока + пакФмт.Sprintf(" %d", слово.номер))
			}
		}
	}
	return слово, ош
}

//Род -- возвращает род слова
func (сам *ТуСлово) Род() int {
	return int(сам.род)
}

//Номер -- возвращает номер слова в исходнике
func (сам *ТуСлово) Номер() int {
	return сам.номер
}

// ЕслиПерваяБуква -- проверяет наличие первой буквы в строке
func ЕслиПерваяБуква(пСтрока string) (бРез bool) {
	руна := []rune(пСтрока)
	стрЛит := string(руна[0])
	//пакКонс.Конс.Печать("пакСлово.ЕслиПерваяБуква(): лит=" + стрЛит)
	if ЕслиБуква(стрЛит) {
		бРез = true
	}
	return бРез
}

// ЕслиБуква -- проверяет наличие буквы в литере
func ЕслиБуква(пЛит string) (бРез bool) {
	if strings.Contains(стрБуквыВсе, пЛит) {
		бРез = true
	}
	return бРез
}

//ЕслиЦифра -- проверяет является ли литера цифрой
func ЕслиЦифра(пЛит string) (бРез bool) {
	стрЦифры := "0123456789."
	if strings.Contains(стрЦифры, пЛит) {
		бРез = true
	}
	return бРез
}

//ЕслиНеЦифра -- проверяет, что строка не цифра
func ЕслиНеЦифра(пСтрока string) (бРез bool) {
	стрЦифры := "0123456789."
	for лит := range пСтрока {
		if !(strings.Contains(стрЦифры, string(лит))) {
			бРез = true
		}
	}
	return бРез
}

//Строка -- возвращает слово, которое хранит тип
func (сам *ТуСлово) Строка() (строка string) {
	if сам._строка == "" {
		panic(пакФмт.Sprintf("ТуСлово.Строка(): ТуСтрока._строка не может быть пустой"))
	}
	return сам._строка
}

// Проверяет, что строка не находится в ключевых словах
func _ЕслиКлючевоеСлово(пСтрока string) (бРез bool) {
	бРез = false // По умолчанию -- имя разрешено
	for _, группа := range КлючСлово {
		for _, ключевое := range группа {
			if пСтрока == ключевое {
				бРез = true
				return бРез
			}
		}
	}
	return бРез
}

// Проверяет, что литера не находится в списке запрещённых (для имён сущностей)
func _ЕслиЛитЗапрещена(пЛит rune) (бРез bool) {
	//стрЗапрет:=[]rune("")
	стрЗапрет := []rune("\"~`!@$%^&*()-=+{}[]|\\<,>?/№;:\t\n'\r ")
	for _, лит := range стрЗапрет {
		if пЛит == лит {
			бРез = true
			return бРез
		}
	}
	return бРез
}

// ЕслиИмяСтрого -- проверяет, что строка обладает строгим соответствием, чтобы быть именем сущности
func (сам *ТуСлово) ЕслиИмяСтрого() (бРез bool) {
	// имя сущности должно начинаться либо с "_", либо с буквы
	бРез = true // По умолчанию -- имя строго соответствует
	строка := []rune(сам._строка)
	лит := string(строка[0])
	if лит == "_" || ЕслиБуква(лит) {
		if лит != "_" && _ЕслиКлючевоеСлово(сам._строка) {
			бРез = false
			return бРез
		}
		for _, лит := range строка {
			// Точка в имени -- допустимо, но здесь её не будет.
			if _ЕслиЛитЗапрещена(лит) {
				//
				бРез = false
				return бРез
			}
		}
	}
	return бРез
}

// СтрИсх -- возвращает строку исходника, содержащую подстроку
func (сам *ТуСлово) СтрИсх() (стр string) {
	if сам._стрИсх == "" {
		panic("ТуСлово.СтрИсх(): строка исходника не может быть пустой")
	}
	return сам._стрИсх
}
