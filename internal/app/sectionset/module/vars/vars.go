package vars

import (
	"log"

	"github.com/prospero78/goOC/internal/app/scanner/word"
	"github.com/prospero78/goOC/internal/app/sectionset/module/keywords"
	"github.com/prospero78/goOC/internal/app/sectionset/module/vars/srcvar"
	"github.com/prospero78/goOC/internal/types"
)

/*
	Пакет предоставляет тип для вырезания переменных из пула слов модуля.
*/

// TVars -- операции со словами секции переменных
type TVars struct {
	poolVar  []*srcvar.TSrcVar
	keywords types.IKeywords
}

// New -- возвращает новый *TVars
func New() *TVars {
	return &TVars{
		keywords: keywords.GetKeys(),
		poolVar:  make([]*srcvar.TSrcVar, 0),
	}
}

// Split -- вырезает слова из секции переменных
func (sf *TVars) Split(pool []*word.TWord) []*word.TWord {
	// Убедиться, что это секция переменных
	keyVar := pool[0]
	if !sf.keywords.IsKey("VAR", keyVar.Word()) {
		return pool
	}
	// В цикле перебрать все переменные
	pool = pool[1:]
	isRec := false
	for len(pool) > 0 {
		name := pool[0]
		if !name.IsName() {
			log.Panicf("TVars.Split(): word(%v) must by name\n", name.Word())
		}
		if sf.keywords.IsKey("PROCEDURE", name.Word()) {
			return pool
		}
		if sf.keywords.IsKey("BEGIN", name.Word()) {
			return pool
		}
		pool = pool[1:]
		svar := srcvar.New(name)
		pool = sf.checkExport(svar, pool)
		pool = sf.checkDefine(pool)
		pool, isRec = sf.checkRecord(svar, pool)
		if isRec {
			sf.poolVar = append(sf.poolVar, svar)
			continue
		}
		isSimple := false
		if pool, isSimple = sf.findVarWord(svar, pool); isSimple {
			if pool[1].Word() == ";" {
				pool = pool[2:]
				sf.poolVar = append(sf.poolVar, svar)
				continue
			}
		}
		sf.poolVar = append(sf.poolVar, svar)
		// Могут просто закончиться переменные и начаться процедуры
		if pool[0].Word() == ";" {
			pool = pool[1:]
		}
	}
	return pool
}

// Ищет все слова переменной, если выражение простое -- возвращает признак
func (sf *TVars) findVarWord(svar *srcvar.TSrcVar, pool []*word.TWord) (pl []*word.TWord, isSimple bool) {
	term := pool[1]
	isRec := false
	if term.Word() == ";" {
		svar.AddWord(pool[0])
		pool = pool[1:]
		return pool, true
	}
	for term.Word() != ";" {
		pool, isRec = sf.checkRecord(svar, pool)
		if isRec {
			return pool, false
		}
		svar.AddWord(pool[0])
		pool = pool[1:]
		term = pool[0]
		continue
	}
	return pool, false
}

// проверяет наличие экспорта в списке слов
func (sf *TVars) checkExport(svar *srcvar.TSrcVar, pool []*word.TWord) []*word.TWord {
	exp := pool[0]
	if exp.Word() == "*" {
		svar.SetExport()
		pool = pool[1:]
	}
	return pool
}

// Проверяет наличие присвоения в указанной позиции
func (sf *TVars) checkDefine(pool []*word.TWord) (pl []*word.TWord) {
	define := pool[0]
	if define.Word() != ":" { // Признак определения типа
		// log.Panicf("TVars.checkAsign(): bad assign(%v) for type\n", assign.Word())
		return pool
	}
	// Получить описатель типа
	pool = pool[1:]
	return pool
}

// Проверяет наличие слова RECORD (если есть -- добавляет всё до <END>)
func (sf *TVars) checkRecord(svar *srcvar.TSrcVar, pool []*word.TWord) (pl []*word.TWord, res bool) {
	record := pool[0]
	if !sf.keywords.IsKey("RECORD", record.Word()) {
		return pool, false
	}
	svar.AddWord(pool[0])
	pool = pool[1:]
	// Это запись, добавляем все до <END;>, если опять встретится RECORD
	// вызовем рекурсивно себя
	isRec := false
	for {
		word := pool[0]
		if sf.keywords.IsKey("RECORD", word.Word()) { // Рекурсивный вызов
			pool, isRec = sf.checkRecord(svar, pool)
			if isRec {
				word = pool[0]
			}
		}
		if !sf.keywords.IsKey("END", word.Word()) {
			svar.AddWord(pool[0])
			pool = pool[1:]
			continue
		}
		// Если конец -- проверить разделитель
		term := pool[1]
		if term.Word() == ";" {
			svar.AddWord(pool[0])
			pool = pool[2:]
			return pool, true
		}
		log.Panicf("TVars.checkRecord(): unknown word=%+v\n", term)
	}
}

// Len -- возвращает количество типов
func (sf *TVars) Len() int {
	// for _, vr := range sf.poolVar {
	// 	fmt.Printf("%v = ", vr.Name())
	// 	for _, word := range vr.Words() {
	// 		fmt.Printf("%v ", word.Word())
	// 	}
	// 	fmt.Print("\n")
	// }
	return len(sf.poolVar)
}
