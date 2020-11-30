package srcproc

import "oc/internal/app/scanner/word"

/*
	Пакет предоставляет тип для хранения слов отдельной процедуры.
*/

// TSrcProc --
type TSrcProc struct {
	poolWord []*word.TWord
	name     *word.TWord
}

// New -- возвращает новый *TSrcProc
func New(name *word.TWord) *TSrcProc {
	return &TSrcProc{
		poolWord: make([]*word.TWord, 0),
		name:     name,
	}
}

// AddWord -- добавляет слово в пул слов процедуры
func (sf *TSrcProc)AddWord(word *word.TWord){
	sf.poolWord = append(sf.poolWord, word)
}

// Name -- возвращает хранимое имя процедуры
func (sf *TSrcProc)Name()string{
	return sf.name.Word()
}
