package keywords

/*
	Пакет предоставляет тип с ключевыми словами
*/

// TKeywords -- операции с ключевыми словами
type TKeywords struct {
	poolKey map[string][]string
}

var(
	// Keys -- глобальный объект ключей
	Keys *TKeywords
)

func init() {
	Keys = &TKeywords{
		poolKey: make(map[string][]string),
	}
	Keys.addKeyword("MODULE", "МОДУЛЬ")
	Keys.addKeyword("IMPORT", "ИМПОРТ")
	Keys.addKeyword("CONST", "КОНСТ")
	Keys.addKeyword("TYPE", "ТИПЫ")
	Keys.addKeyword("VAR", "ПЕРЕМ", "ПРМ", "УКАЗ")
	Keys.addKeyword("PROCEDURE", "ПРОЦЕДУРА", "ПРОЦ")
	Keys.addKeyword("BEGIN", "НАЧАЛО", "НАЧ")
	Keys.addKeyword("RECORD", "ЗАПИСЬ")
	Keys.addKeyword("END", "КОНЕЦ", "КНЦ", "КОН")
	Keys.addKeyword("RECORD", "ЗАПИСЬ")
	Keys.addKeyword("TRUE", "ИСТИНА")
	Keys.addKeyword("FALSE", "ЛОЖЬ")
}

// IsKey -- проверяет ключевое слово с необходимым образцом
func (sf *TKeywords) IsKey(sample, key string) bool {
	keyword, ok := sf.poolKey[sample]
	if !ok {
		return false
	}
	for _, val := range keyword {
		if val == key {
			return true
		}
	}
	return false
}

// Конструирует срез допустимых ключевых слов по позиции
func (sf *TKeywords) addKeyword(key string, key1 ...string) {
	res := make([]string, 0)
	res = append(res, key)
	res = append(res, key1...)
	sf.poolKey[key] = res
}
