package srcvar

import (
	"log"

	"github.com/prospero78/goOC/internal/app/scanner/word"
	"github.com/prospero78/goOC/internal/types"
)

/*
	Пакет предоставляет тип по хранению слов переменных.
*/

// TSrcVar -- операции о словами переменной
type TSrcVar struct {
	poolWord []*word.TWord
	name     *word.TWord
	isExport bool
}

// New -- возвращает новый *TScrVar
func New(name *word.TWord) *TSrcVar {
	return &TSrcVar{
		poolWord: make([]*word.TWord, 0),
		name:     name,
	}
}

// AddWord -- добавляет слово в словарь переменных
func (sf *TSrcVar) AddWord(word *word.TWord) {
	if word == nil {
		log.Panicf("TSrcVar.AddWord(): word==nil\n")
	}
	sf.poolWord = append(sf.poolWord, word)
}

// Name -- возвращает хранимое имя переменной
func (sf *TSrcVar) Name() types.AWord {
	return sf.name.Word()
}

// SetExport -- устанавливает признак экспорта переменной
func (sf *TSrcVar) SetExport() {
	if sf.isExport {
		log.Panicf("TSrcVar.SetExport(): export already set\n")
	}
	sf.isExport = true
}

// Words -- возвращает хранимые слова типа
func (sf *TSrcVar) Words() []*word.TWord {
	return sf.poolWord
}
