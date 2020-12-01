package begin

/*
	Пакет предоставляет тип для хранения слов секции BEGIN модуля.
*/

import (
	"log"
	"oc/internal/app/scanner/word"
	"oc/internal/app/sectionset/module/keywords"
)

// TBegin -- операци ис секцией BEGIN модуля
type TBegin struct {
	keywords *keywords.TKeywords
	poolWord []*word.TWord
}

// New -- возвращает новый *TBegin
func New() *TBegin {
	return &TBegin{
		keywords: keywords.Keys,
		poolWord: make([]*word.TWord, 0),
	}
}

// Split -- вырезает слова модуля секции BEGIN. Слов остаться не должно.
func (sf *TBegin) Split(pool []*word.TWord) []*word.TWord {
	// Проверить, что есть секция BEGIN
	sec := pool[0]
	if !sf.keywords.IsKey("BEGIN", sec.Word()) {
		return pool
	}
	pool = pool[1:]
	for len(pool) > 1 {
		sf.poolWord = append(sf.poolWord, pool[0])
		pool = pool[1:]
	}
	// Проверить, что последнее слово "END"
	end := pool[0]
	if !sf.keywords.IsKey("END", end.Word()) {
		log.Panicf("TBegin.Split(): word(%v)!='end'\n", end.Word())
	}
	pool = pool[1:]
	return pool
}
