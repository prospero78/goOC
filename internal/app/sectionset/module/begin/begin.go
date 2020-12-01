package begin

/*
	Пакет предоставляет тип для хранения слов секции BEGIN модуля.
*/

import (
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
		keywords: keywords.New(),
		poolWord: make([]*word.TWord, 0),
	}
}

// Split -- вырезает слова модуля секции BEGIN. Слов остаться не должно.
func (sf *TBegin) Split(pool []*word.TWord) []*word.TWord {
	return nil
}
