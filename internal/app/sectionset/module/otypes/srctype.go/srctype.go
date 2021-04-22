package srctype

import (
	"log"

	"github.com/prospero78/goOC/internal/types"
)

/*
	Пакет предоставляет тип для хранения типа, определяемом в исходнике.
*/

// TSrcType -- операции со словами типа из исходника
type TSrcType struct {
	word     types.IWord
	isExport bool
	listWord []types.IWord // Слова описателя типа
}

// New -- возвращает новый *TSrcType
func New(wrd types.IWord) *TSrcType {
	if wrd == nil {
		log.Panicf("srctype.go/New(): wrd==nil\n")
	}
	st := &TSrcType{
		word:     wrd,
		listWord: make([]types.IWord, 0),
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
func (sf *TSrcType) AddWord(word types.IWord) {
	if word == nil {
		log.Panicf("TSrcType.AddWord(): word==nil\n")
	}
	sf.listWord = append(sf.listWord, word)
}

// Words -- возвращает хранимые слова типа
func (sf *TSrcType) Words() []types.IWord {
	return sf.listWord
}

// Name -- возвращает имя типа
func (sf *TSrcType) Name() types.AWord {
	return sf.word.Word()
}
