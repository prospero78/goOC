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

var (
	запрет_имя = []string{КсМодуль, КсИмпорт, КсКонст, КсТипы, КсБулево, КсБайт, КсЦелое, КсЛит, КсНабор,
		КсВещ, КсПер, КсУказ, КсДо, КсМассив, КсИз, КсНачало, КсКонец, КсПроцедура, КсДля,
		КсПока, КсПовтор, КсЗапись, КсЯвляется}
)

var слов_всего = 0

type ТуСлово struct {
	стрИсх   пакТипы.ТСтрИсх       // Строка исходника
	стрСлово пакТипы.ТСтрСлово     // Строка слова
	цРод     пакТипы.ТРод          // род слова
	конс     *пакКонс.ТуКонсоль    // Системная консоль
	Коорд    *пакКоорд.ТуКоордФикс // Координаты слова в исходном тексте
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

	род_проверить := func(сам *ТуСлово, пСтрока string) (ош error) { //Устанавливает род слова.
		var род int
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
		case !НеЦифра(пСтрока) || (string(пСтрока[0]) == "."):
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
		сам.цРод = пакТипы.ТРод(род)
		return ош
	}

	if ош1 := слово_проверить(пСтрока); ош != nil {
		ош1 = пакФмт.Errorf("ТуСлово.Новое(): при вызове слово_проверить()\n\t%v", ош1)
	} else {
		if ош2 := род_проверить(слово, пСтрока); ош2 != nil {
			ош2 = пакФмт.Errorf("пакСлово.Новое(): при вызове род_проверить()\n\t%v", ош1)
		} else {
			if коорд, ош3 := пакКоорд.НовыйФикс(пакКоорд.ТЦелСтр(пНомСтр), пакКоорд.ТЦелПоз(пНомПоз)); ош3 != nil {
				ош3 = пакФмт.Errorf("пакСлово.Новое(): ошибка при создании ТуКоорд\n %v", ош2)
			} else {
				слово = &ТуСлово{
					стрСлово: пакТипы.ТСтрСлово(пСтрока),
					конс:     пакКонс.Конс,
					Коорд:    коорд,
				}
			}
		}
	}
	return слово, ош
}

// Проверяет наличие первой буквы в строке
func ЕслиПерваяБуква(пСтрока string) (бРез bool) {
	руна := []rune(пСтрока)
	стрНач := string(руна[0])
	//пакКонс.Конс.Печать("пакСлово.ЕслиПерваяБуква(): лит=" + стрНач)
	if strings.Contains(стрБуквыВсе, стрНач) {
		бРез = true
	}
	return бРез
}

// Проверяет наличие буквы в литере
func ЕслиБуква(пЛит string) (бРез bool) {
	if strings.Contains(пЛит, стрБуквыВсе) {
		бРез = true
	}
	return бРез
}

func ЕслиЦифра(пЛит string) (бРез bool) {
	стрЦифры := "0123456789."
	if strings.Contains(пЛит, стрЦифры) {
		бРез = true
	}
	return бРез
}

func НеЦифра(пСтрока string) (бРез bool) {
	стрЦифры := "0123456789."
	for лит := range пСтрока {
		if !(strings.Contains(string(лит), стрЦифры)) {
			бРез = true
		}
	}
	return бРез
}
