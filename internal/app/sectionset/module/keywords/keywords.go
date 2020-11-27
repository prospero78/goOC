package keywords

/*
	Пакет предоставляет тип с ключевыми словами
*/

// TKeywords -- операции с ключевыми словами
type TKeywords struct {
	poolKey map[string][]string
}

// New -- возвращает новый *TKeywords
func New() (kws *TKeywords) {
	kws = &TKeywords{
		poolKey: make(map[string][]string),
	}
	kws.addKeyword("MODULE", "МОДУЛЬ")
	kws.addKeyword("IMPORT", "ИМПОРТ")
	return kws
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
func (sf *TKeywords) addKeyword(key, key1 string) {
	res := make([]string, 0)
	res = append(res, key)
	res = append(res, key1)
	sf.poolKey[key] = res
}
