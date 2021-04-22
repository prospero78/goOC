package procs

import (
	"log"

	"github.com/prospero78/goOC/internal/app/sectionset/module/keywords"
	"github.com/prospero78/goOC/internal/app/sectionset/module/procs/srcproc"
	"github.com/prospero78/goOC/internal/types"
)

/*
	Пакет предоставляет тип для вырезания слов относящихся к процедурам.
*/

// TProcedures -- операции с вырезаним слов модуля для процедур
type TProcedures struct {
	keywords types.IKeywords
	poolProc []*srcproc.TSrcProc // Пул вычисленных процедур
}

// New -- возвращает новый *TProcedures
func New() *TProcedures {
	return &TProcedures{
		keywords: keywords.GetKeys(),
		poolProc: make([]*srcproc.TSrcProc, 0),
	}
}

// Split -- вырезает все слова процедур, формирует словарь
func (sf *TProcedures) Split(listWord []types.IWord) []types.IWord {
	for {
		// Проверить, что впереди реально процедура
		word := listWord[0]
		if !sf.keywords.IsKey("PROCEDURE", word.Word()) {
			return listWord
		}
		listWord = listWord[1:]
		name := listWord[0]
		if !name.IsName() {
			log.Panicf("TProcessing.getProcedure(): name(%v) must be strong\n", name.Word())
		}
		listWord = listWord[1:]
		proc := srcproc.New(name)
		listWord = sf.getProcedure(proc, listWord)
		sf.poolProc = append(sf.poolProc, proc)
	}
}

// Рекурсивная функция для вычисления процедур.
// Внутри процедуры могут быть ещё процедуры
func (sf *TProcedures) getProcedure(proc *srcproc.TSrcProc, listWord []types.IWord) []types.IWord {
	for {
		word := listWord[0]
		if sf.keywords.IsKey("PROCEDURE", word.Word()) {
			listWord = listWord[1:]
			listWord = sf.getProcedure(proc, listWord)
			word = listWord[0]
		}
		if !sf.keywords.IsKey("END", word.Word()) {
			proc.AddWord(word)
			listWord = listWord[1:]
			continue
		}
		name := listWord[1]
		term := listWord[2]
		if name.Word() != proc.Name() && term.Word() != ";" {
			proc.AddWord(word)
			listWord = listWord[1:]
			continue
		}
		// Это точно конец процедуры.
		listWord = listWord[3:]
		return listWord
	}
}

// Len -- возвращает число процедур в модуле
func (sf *TProcedures) Len() int {
	return len(sf.poolProc)
}
