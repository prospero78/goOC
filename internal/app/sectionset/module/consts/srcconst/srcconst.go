// Package srcconst -- пакет для константы исходника с именем и значением.
package srcconst

import (
	"log"

	"github.com/prospero78/goOC/internal/app/sectionset/module/consts/srcconst/constexpres"
	"github.com/prospero78/goOC/internal/types"
)

// TConst --  операции с константой секции CONST
type TConst struct {
	name     types.IWord
	isExport bool
	listWord []types.IWord
	strType  string                        // Строковое представление типа
	exp      *constexpres.TConstExpression // Выражение для константы

}

// New -- возвращает новый *TConst
func New(name types.IWord) *TConst {
	{ // Предусловия
		if name == nil {
			log.Panicf("srcconst.go/New(): name==nil\n")
		}
	}
	return &TConst{
		name:     name,
		listWord: make([]types.IWord, 0),
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
func (sf *TConst) AddWord(word types.IWord) {
	if word == nil {
		log.Panicf("TConst.AddWord(): word==nil\n")
	}
	sf.listWord = append(sf.listWord, word)
}

// GetWords -- возвращает слова константы
func (sf *TConst) GetWords() []types.IWord {
	return sf.listWord
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
func (sf *TConst) SetPoolWord(listWord []types.IWord) {
	if listWord == nil {
		log.Panicf("TConst.SetPoolWord(): pool==nil\n")
	}
	sf.listWord = listWord
}

// Module -- возвращает имя модуля, к которой относится константа
func (sf *TConst) Module() types.AModule {
	return *sf.name.Module()
}
