package procs

import (
	"log"
	"oc/internal/app/scanner/word"
	"oc/internal/app/sectionset/module/keywords"
	"oc/internal/app/sectionset/module/procs/srcproc"
)

/*
	Пакет предоставляет тип для вырезания слов относящихся к процедурам.
*/

// TProcedures -- операции с вырезаним слов модуля для процедур
type TProcedures struct {
	keywords *keywords.TKeywords
	poolProc []*srcproc.TSrcProc // Пул вычисленных процедур
}

// New -- возвращает новый *TProcedures
func New() *TProcedures {
	return &TProcedures{
		keywords: keywords.New(),
		poolProc: make([]*srcproc.TSrcProc, 0),
	}
}

// Split -- вырезает все слова процедур, формирует словарь
func (sf *TProcedures) Split(pool []*word.TWord) []*word.TWord {
	for {
		// Проверить, что впереди реально процедура
		word := pool[0]
		if !sf.keywords.IsKey("PROCEDURE", word.Word()) {
			return pool
		}
		pool = pool[1:]
		name := pool[0]
		if !name.IsName() {
			log.Panicf("TProcessing.getProcedure(): name(%v) must be strong\n", name.Word())
		}
		pool = pool[1:]
		proc := srcproc.New(name)
		pool = sf.getProcedure(proc, pool)
		sf.poolProc = append(sf.poolProc, proc)
	}
}

// Рекурсивная функция для вычисления процедур.
// Внутри процедуры могут быть ещё процедуры
func (sf *TProcedures) getProcedure(proc *srcproc.TSrcProc, pool []*word.TWord) []*word.TWord {
	for {
		word := pool[0]
		if sf.keywords.IsKey("PROCEDURE", word.Word()) {
			pool = sf.getProcedure(proc, pool)
			word = pool[0]
		}
		if !sf.keywords.IsKey("END", word.Word()) {
			proc.AddWord(word)
			pool = pool[1:]
			continue
		}
		name := pool[1]
		term := pool[2]
		if name.Word() != proc.Name() && term.Word() != ";" {
			proc.AddWord(word)
			pool = pool[1:]
			continue
		}
		// Это точно конец процедуры.
		pool = pool[3:]
		return pool
	}
}
