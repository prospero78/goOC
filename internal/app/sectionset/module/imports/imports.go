package imports

/*
	пакет предоставляет тип для разбора секции импорта
*/

import (
	"log"

	"github.com/prospero78/goOC/internal/app/scanner/word"
	"github.com/prospero78/goOC/internal/app/sectionset/module/imports/alias"
	"github.com/prospero78/goOC/internal/app/sectionset/module/keywords"
	"github.com/prospero78/goOC/internal/types"
)

// TImports -- операции с секцией импорта
type TImports struct {
	poolAlias []*alias.TAlias
	keywords  types.IKeywords
}

// New -- возвращает новый *tImports
func New() *TImports {
	return &TImports{
		poolAlias: make([]*alias.TAlias, 0),
		keywords:  keywords.GetKeys(),
	}
}

// Split -- выделяет слова импорта и возвращает что осталось
func (sf *TImports) Split(pool []*word.TWord) []*word.TWord {
	imp := pool[0]
	if !sf.keywords.IsKey("IMPORT", imp.Word()) { // Если нет импорта -- сраз всё вернуть
		return pool
	}
	pool = pool[1:]
	for len(pool) > 2 {
		wordName := pool[0]
		if !wordName.IsName() {
			log.Panicf("TImports.Split(): module wordName (%q) not valid\n", wordName.Word())
		}
		pool = pool[1:]
		term := pool[0]
		switch term.Word() {
		case ",": // Прямое имя модуля
			alais := alias.New(wordName.Word(), "")
			sf.poolAlias = append(sf.poolAlias, alais)
			pool = pool[1:]
		case ";": // Окончание импорта модулей
			alais := alias.New(wordName.Word(), "")
			sf.poolAlias = append(sf.poolAlias, alais)
			pool = pool[1:]
			return pool
		case ":=": // Алиас к имени модуля
			pool = pool[1:]
			wordTrueName := pool[0]
			if !wordTrueName.IsName() {
				log.Panicf("TImports.Split(): module wordTrueName (%q) not valid\n", wordTrueName.Word())
			}
			alias := alias.New(wordTrueName.Word(), wordName.Word())
			sf.poolAlias = append(sf.poolAlias, alias)
			pool = pool[1:]
			term := pool[0]
			if !(term.Word() == "," || term.Word() == ";") {
				log.Panicf("TImports.Split(): invalid term(%q)\n", term.Word())
			}
			pool = pool[1:]
			if term.Word() == ";" {
				return pool
			}
		}
	}
	// log.Panicf("TImports.Split(): not have IMPORTS\n")
	return nil
}

// Imports -- возращает все импроты в модуле
func (sf *TImports) Imports() []*alias.TAlias {
	return sf.poolAlias
}

// Len -- возвращает общее число импортов
func (sf *TImports) Len() int {
	return len(sf.poolAlias)
}
