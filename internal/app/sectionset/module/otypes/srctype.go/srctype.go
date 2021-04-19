package srctype

import (
	"log"

	"github.com/prospero78/goOC/internal/app/scanner/word"
	"github.com/prospero78/goOC/internal/types"
)

/*
	Пакет предоставляет тип для хранения типа, определяемом в исходнике.
*/

// TSrcType -- операции со словами типа из исходника
type TSrcType struct {
	word     *word.TWord
	isExport bool
	poolWord []*word.TWord // Слова описателя типа
}

// New -- возвращает новый *TSrcType
func New(wrd *word.TWord) *TSrcType {
	if wrd == nil {
		log.Panicf("srctype.go/New(): wrd==nil\n")
	}
	st := &TSrcType{
		word:     wrd,
		poolWord: make([]*word.TWord, 0),
	}
	return st
}

// SetExport -- устанавливает признак экспортирования типа
func (sf *TSrcType) SetExport() {
	if sf.isExport {
		log.Panicf("srctype.go/New(): export already set\n")
	}
}

// AddWord -- добавляет слово в описатель типа
func (sf *TSrcType) AddWord(word *word.TWord) {
	if word == nil {
		log.Panicf("TSrcType.AddWord(): word==nil\n")
	}
	sf.poolWord = append(sf.poolWord, word)
}

// Words -- возвращает хранимые слова типа
func (sf *TSrcType) Words() []*word.TWord {
	return sf.poolWord
}

// Name -- возвращает имя типа
func (sf *TSrcType) Name() types.AWord {
	return sf.word.Word()
}
