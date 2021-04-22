package begin

/*
	Пакет предоставляет тип для хранения слов секции BEGIN модуля.
*/

import (
	"log"

	"github.com/prospero78/goOC/internal/app/sectionset/module/keywords"
	"github.com/prospero78/goOC/internal/types"
)

// TBegin -- операци ис секцией BEGIN модуля
type TBegin struct {
	keywords types.IKeywords
	listWord []types.IWord
}

// New -- возвращает новый *TBegin
func New() *TBegin {
	return &TBegin{
		keywords: keywords.GetKeys(),
		listWord: make([]types.IWord, 0),
	}
}

// Split -- вырезает слова модуля секции BEGIN. Слов остаться не должно.
func (sf *TBegin) Split(listWord []types.IWord) []types.IWord {
	// Проверить, что есть секция BEGIN
	sec := listWord[0]
	if !sf.keywords.IsKey("BEGIN", sec.Word()) {
		return listWord
	}
	listWord = listWord[1:]
	for len(listWord) > 1 {
		sf.listWord = append(sf.listWord, listWord[0])
		listWord = listWord[1:]
	}
	// Проверить, что последнее слово "END"
	end := listWord[0]
	if !sf.keywords.IsKey("END", end.Word()) {
		log.Panicf("TBegin.Split(): word(%v)!='end'\n", end.Word())
	}
	listWord = listWord[1:]
	return listWord
}
