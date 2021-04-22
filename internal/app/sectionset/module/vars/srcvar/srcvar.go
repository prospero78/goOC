package srcvar

import (
	"log"

	"github.com/prospero78/goOC/internal/types"
)

/*
	Пакет предоставляет тип по хранению слов переменных.
*/

// TSrcVar -- операции о словами переменной
type TSrcVar struct {
	listWord []types.IWord
	name     types.IWord
	isExport bool
}

// New -- возвращает новый *TScrVar
func New(name types.IWord) *TSrcVar {
	return &TSrcVar{
		listWord: make([]types.IWord, 0),
		name:     name,
	}
}

// AddWord -- добавляет слово в словарь переменных
func (sf *TSrcVar) AddWord(word types.IWord) {
	if word == nil {
		log.Panicf("TSrcVar.AddWord(): word==nil\n")
	}
	sf.listWord = append(sf.listWord, word)
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
func (sf *TSrcVar) Words() []types.IWord {
	return sf.listWord
}
