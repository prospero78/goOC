// Package srcconst -- пакет для константы исходника с именем и значением.
package srcconst

import (
	"log"

	"github.com/prospero78/goOC/internal/app/scanner/word"
	"github.com/prospero78/goOC/internal/app/sectionset/module/consts/srcconst/constexpres"
	"github.com/prospero78/goOC/internal/types"
)

// TConst --  операции с константой секции CONST
type TConst struct {
	name     *word.TWord
	isExport bool
	poolWord []*word.TWord
	strType  string                        // Строковое представление типа
	exp      *constexpres.TConstExpression // Выражение для константы

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
		exp:      constexpres.New(),
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
	if word == nil {
		log.Panicf("TConst.AddWord(): word==nil\n")
	}
	sf.poolWord = append(sf.poolWord, word)
}

// GetWords -- возвращает слова константы
func (sf *TConst) GetWords() []*word.TWord {
	return sf.poolWord
}

// Name -- возвращает имя константы
func (sf *TConst) Name() types.AWord {
	return sf.name.Word()
}

// GetType -- возвращает тип константы
func (sf *TConst) GetType() string {
	return sf.strType
}

// SetType -- устанавливает тип константы
func (sf *TConst) SetType(strType string) {
	if strType == "" {
		log.Panicf("TConst.SetType(): strType==''\n")
	}
	if sf.strType != "" {
		if sf.strType != strType {
			log.Panicf("TConst.SetType(): type(%v)!=strType(%v)\n", sf.strType, strType)
		}
		return
	}
	if sf.strType != "" {
		log.Panicf("TConst.SetType(): type(%v) already set, strType=%v\n", sf.strType, strType)
	}
	sf.strType = strType
}

// GetExpres -- возвращает выражение для константы
func (sf *TConst) GetExpres() *constexpres.TConstExpression {
	return sf.exp
}

// SetPoolWord -- устанавливает пул слов после обработки выражения
func (sf *TConst) SetPoolWord(pool []*word.TWord) {
	if pool == nil {
		log.Panicf("TConst.SetPoolWord(): pool==nil\n")
	}
	sf.poolWord = pool
}

// Module -- возвращает имя модуля, к которой относится константа
func (sf *TConst) Module() types.AModule {
	return *sf.name.Module()
}
