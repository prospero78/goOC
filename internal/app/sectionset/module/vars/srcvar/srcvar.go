package srcvar

import (
	"log"
	"oc/internal/app/scanner/word"
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
	sf.poolWord = append(sf.poolWord, word)
}

// Name -- возвращает хранимое имя переменной
func (sf *TSrcVar) Name() string {
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
