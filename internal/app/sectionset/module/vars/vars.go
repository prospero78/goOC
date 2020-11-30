package vars

import "oc/internal/app/scanner/word"

/*
	Пакет предоставляет тип для вырезания переменных из пула слов модуля.
*/

// TVars -- операции со словами секции переменных
type TVars struct {
	poolWord []*word.TWord
}

// New -- возвращае тновый *TVars
func New() *TVars {
	return &TVars{}
}

// Split -- вырезает слова из секции переменных
func (sf *TVars) Split(pool []*word.TWord) []*word.TWord {
	return pool
}
