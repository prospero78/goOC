// модСловоКонстанты
package пакСлово

/*
Содержит все константы, которые могут содержаться в словах
*/

const (
	// константы сущностей
	КПусто            = iota // ""
	КМодульИмя               // Строковое имя модуля
	КМодульСиноним           // Локальное имя модуля
	КЗапятая                 // ","
	КТочкаЗапятая            // ";"
	КИмя                     // Имя пользовательской сущности
	ККомментНачать           // "(*"
	ККоммент                 // Слово внутри комментария
	ККомментЗакончить        // "*)"
	КОпределить              // ":"
	КПрисвоить               // ":="
	КСкобкаОткрКругл         // "("
	КСкобкаЗакрКругл         // ")"
	КДеление
	КУмножить
	КМинус
	КПлюс
	КЧисло
	КСтрока
	КРавно
	КТочка

	//наборы букв для перебора
	стрБуквыРус = "абвгдеёжзийклмнопрстуфхцчшщьыъэюяАБВШДЕЙЖЗИЙКЛМНОПРСТУФХЦЧШЩЬЫЪЭЮЯ"
	стрБуквыАнг = "abcdefghjiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	стрБуквыВсе = стрБуквыРус + стрБуквыАнг
)

var (
	// ключевые слова
	КсМодуль    = []string{"MODULE", "МОДУЛЬ"}
	КсИмпорт    = []string{"IMPORT", "ИМПОРТ"}
	КсКонст     = []string{"CONST", "КОНСТ"}
	КсТипы      = []string{"TYPE", "ТИПЫ"}
	КсБулево    = []string{"BOOLEAN", "БУЛЕВО"}
	КсБайт      = []string{"BYTE", "БАЙТ"}
	КсЦелое     = []string{"INTEGER", "ЦЕЛОЕ"}
	КсЛит       = []string{"CHAR", "ЛИТ"}
	КсНабор     = []string{"SET", "НАБОР"}
	КсВещ       = []string{"REAL", "ДРОБ", "ВЕЩ"}
	КсПер       = []string{"VAR", "ПЕРЕМ", "ССЫЛКА"}
	КсУказ      = []string{"POINTER", "УКАЗ", "УКАЗАТЕЛЬ"}
	КсДо        = []string{"TO", "НА"}
	КсМассив    = []string{"ARRAY", "МАС", "МАССИВ"}
	КсИз        = []string{"OF", "ИЗ"}
	КсНачало    = []string{"BEGIN", "НАЧ", "НАЧАЛО"}
	КсКонец     = []string{"END", "КОНЕЦ"}
	КсПроцедура = []string{"PROCEDURE", "ПРОЦ", "ПРОЦЕДУРА"}
	КсДля       = []string{"FOR", "ДЛЯ"}
	КсПока      = []string{"WHILE", "ПОКА"}
	КсПовтор    = []string{"DO", "ДЕЛАТЬ", "ВЫП", "ВЫПОЛНЯТЬ"}
	КсЗапись    = []string{"RECORD", "ЗАПИСЬ"}
	КсЯвляется  = []string{"IS", "ЕСТЬ"}
	запрет_имя  = [][]string{КсМодуль, КсИмпорт, КсКонст, КсТипы, КсБулево, КсБайт, КсЦелое, КсЛит, КсНабор,
		КсВещ, КсПер, КсУказ, КсДо, КсМассив, КсИз, КсНачало, КсКонец, КсПроцедура, КсДля,
		КсПока, КсПовтор, КсЗапись, КсЯвляется}
)
