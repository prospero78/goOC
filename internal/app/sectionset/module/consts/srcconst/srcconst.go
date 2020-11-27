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
	name     string
	word     *word.TWord
	isExport bool
	poolExpr []string
}

// New -- возвращает новый *TConst
func New(word *word.TWord) *TConst {
	{ // Предусловия
		if word == nil {
			log.Panicf("srcconst.go/New(): word==nil\n")
		}
	}
	return &TConst{
		name:     word.Word(),
		word:     word,
		poolExpr: make([]string, 0),
	}
}

// SetExport -- установки признака на экспорт константы
func (sf *TConst) SetExport() {
	if sf.isExport {
		log.Panicf("TConst.SetExport(): export already set!\n")
	}
	sf.isExport = true
}

// AddExpr -- добавляет часть выражения в константу
func (sf *TConst) AddExpr(exp string) {
	sf.poolExpr = append(sf.poolExpr, exp)
}
