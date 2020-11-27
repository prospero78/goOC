package imports

/*
	пакет предоставляет тип для разбора секции импорта
*/

import (
	"log"
	"oc/internal/app/scanner/word"
	"oc/internal/app/sectionset/module/imports/alias"
	"oc/internal/app/sectionset/module/keywords"
)

// TImports -- операции с секцией импорта
type TImports struct {
	poolAlias []*alias.TAlias
	keywords  *keywords.TKeywords
}

// New -- возвращает новый *tImports
func New() *TImports {
	return &TImports{
		poolAlias: make([]*alias.TAlias, 0),
		keywords:  keywords.New(),
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
		nameWord := pool[0]
		if !nameWord.IsName() {
			log.Panicf("TImports.Split(): module name (%q) not valid\n", nameWord.Word())
		}
		pool = pool[1:]
		term := pool[0]
		switch term.Word() {
		case ",": // Прямое имя модуля
			alais := alias.New(nameWord.Word(), "")
			sf.poolAlias = append(sf.poolAlias, alais)
			pool = pool[1:]
		case ":=": // Алиас к имени модуля
			pool = pool[1:]
			name := pool[0]
			if !name.IsName() {
				log.Panicf("TImports.Split(): module name (%q) not valid\n", name.Word())
			}
			alias := alias.New(name.Word(), nameWord.Word())
			sf.poolAlias = append(sf.poolAlias, alias)
			pool = pool[1:]
			term:=pool[0]
			if !(term.Word()=="," ||  term.Word()==";"){
				log.Panicf("TImports.Split(): invalid term(%q)\n", term.Word())
			}
			pool = pool[1:]
		case ";": // Окончание импорта модулей
			alais := alias.New(nameWord.Word(), "")
			sf.poolAlias = append(sf.poolAlias, alais)
			pool = pool[1:]
			log.Printf("TImports.Split(): all imports=%v\n", len(sf.poolAlias))
			return pool
		}
	}
	return nil
}
