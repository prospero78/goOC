package keywords

import "github.com/prospero78/goOC/internal/types"

/*
	Пакет предоставляет тип с ключевыми словами
*/

// TKeywords -- операции с ключевыми словами
type TKeywords struct {
	dictKey map[types.AWord][]types.AWord
}

var (
	// keys -- глобальный объект ключей
	keys *TKeywords
)

// GetKeys -- возвращает набор ключей для языка
func GetKeys() types.IKeywords {
	return keys
}

func init() {
	keys = &TKeywords{
		dictKey: make(map[types.AWord][]types.AWord),
	}
	keys.addKeyword("MODULE", "МОДУЛЬ")
	keys.addKeyword("IMPORT", "ИМПОРТ")
	keys.addKeyword("CONST", "КОНСТ")
	keys.addKeyword("TYPE", "ТИПЫ")
	keys.addKeyword("VAR", "ПЕРЕМ", "ПРМ", "УКАЗ")
	keys.addKeyword("PROCEDURE", "ПРОЦЕДУРА", "ПРОЦ")
	keys.addKeyword("BEGIN", "НАЧАЛО", "НАЧ")
	keys.addKeyword("RECORD", "ЗАПИСЬ")
	keys.addKeyword("END", "КОНЕЦ", "КНЦ", "КОН")
	keys.addKeyword("RECORD", "ЗАПИСЬ")
	keys.addKeyword("TRUE", "ИСТИНА")
	keys.addKeyword("FALSE", "ЛОЖЬ")
}

// IsKey -- проверяет ключевое слово с необходимым образцом
func (sf *TKeywords) IsKey(sample, key types.AWord) bool {
	keyword, ok := sf.dictKey[sample]
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
func (sf *TKeywords) addKeyword(primarKey types.AWord, otherkeys ...types.AWord) {
	res := make([]types.AWord, 0)
	res = append(res, primarKey)
	res = append(res, otherkeys...)
	sf.dictKey[primarKey] = res
}
