package srcconst

import (
	"log"
	"oc/internal/app/scanner/word"
)

/*
	Тип предоставляет пакет для константы исходника с именем и значением.
*/

// TConst --  операции с константой секции CONST
type TConst struct {
	name     *word.TWord
	isExport bool
	poolWord []*word.TWord
}

// New -- возвращает новый *TConst
func New(name *word.TWord) *TConst {
	{ // Предусловия
		if name == nil {
			log.Panicf("srcconst.go/New(): name==nil\n")
		}
	}
	return &TConst{
		name:     name,
		poolWord: make([]*word.TWord, 0),
	}
}

// SetExport -- установки признака на экспорт константы
func (sf *TConst) SetExport() {
	if sf.isExport {
		log.Panicf("TConst.SetExport(): export already set!\n")
	}
	sf.isExport = true
}

// AddWord -- добавляет слова в константу
func (sf *TConst) AddWord(word *word.TWord) {
	sf.poolWord = append(sf.poolWord, word)
}

// GetWords -- возвращает слова константы
func (sf *TConst) GetWords() []*word.TWord {
	return sf.poolWord
}

// Name -- возвращает имя константы
func (sf *TConst) Name() string {
	return sf.name.Word()
}
