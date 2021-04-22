package srcproc

import (
	"log"

	"github.com/prospero78/goOC/internal/types"
)

/*
	Пакет предоставляет тип для хранения слов отдельной процедуры.
*/

// TSrcProc --
type TSrcProc struct {
	listWord []types.IWord
	name     types.IWord
}

// New -- возвращает новый *TSrcProc
func New(name types.IWord) *TSrcProc {
	return &TSrcProc{
		listWord: make([]types.IWord, 0),
		name:     name,
	}
}

// AddWord -- добавляет слово в пул слов процедуры
func (sf *TSrcProc) AddWord(word types.IWord) {
	if word == nil {
		log.Panicf("TSrcProc.AddWord(): word==nil\n")
	}
	sf.listWord = append(sf.listWord, word)
}

// Name -- возвращает хранимое имя процедуры
func (sf *TSrcProc) Name() types.AWord {
	return sf.name.Word()
}
