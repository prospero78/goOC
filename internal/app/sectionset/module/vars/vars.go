package vars

import (
	"log"

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
func (sf *TVars) Split(listWord []types.IWord) []types.IWord {
	// Убедиться, что это секция переменных
	keyVar := listWord[0]
	if !sf.keywords.IsKey("VAR", keyVar.Word()) {
		return listWord
	}
	// В цикле перебрать все переменные
	listWord = listWord[1:]
	isRec := false
	for len(listWord) > 0 {
		name := listWord[0]
		if !name.IsName() {
			log.Panicf("TVars.Split(): word(%v) must by name\n", name.Word())
		}
		if sf.keywords.IsKey("PROCEDURE", name.Word()) {
			return listWord
		}
		if sf.keywords.IsKey("BEGIN", name.Word()) {
			return listWord
		}
		listWord = listWord[1:]
		svar := srcvar.New(name)
		listWord = sf.checkExport(svar, listWord)
		listWord = sf.checkDefine(listWord)
		listWord, isRec = sf.checkRecord(svar, listWord)
		if isRec {
			sf.poolVar = append(sf.poolVar, svar)
			continue
		}
		isSimple := false
		if listWord, isSimple = sf.findVarWord(svar, listWord); isSimple {
			if listWord[1].Word() == ";" {
				listWord = listWord[2:]
				sf.poolVar = append(sf.poolVar, svar)
				continue
			}
		}
		sf.poolVar = append(sf.poolVar, svar)
		// Могут просто закончиться переменные и начаться процедуры
		if listWord[0].Word() == ";" {
			listWord = listWord[1:]
		}
	}
	return listWord
}

// Ищет все слова переменной, если выражение простое -- возвращает признак
func (sf *TVars) findVarWord(svar *srcvar.TSrcVar, listWord []types.IWord) (pl []types.IWord, isSimple bool) {
	term := listWord[1]
	isRec := false
	if term.Word() == ";" {
		svar.AddWord(listWord[0])
		listWord = listWord[1:]
		return listWord, true
	}
	for term.Word() != ";" {
		listWord, isRec = sf.checkRecord(svar, listWord)
		if isRec {
			return listWord, false
		}
		svar.AddWord(listWord[0])
		listWord = listWord[1:]
		term = listWord[0]
		continue
	}
	return listWord, false
}

// проверяет наличие экспорта в списке слов
func (sf *TVars) checkExport(svar *srcvar.TSrcVar, listWord []types.IWord) []types.IWord {
	exp := listWord[0]
	if exp.Word() == "*" {
		svar.SetExport()
		listWord = listWord[1:]
	}
	return listWord
}

// Проверяет наличие присвоения в указанной позиции
func (sf *TVars) checkDefine(listWord []types.IWord) (pl []types.IWord) {
	define := listWord[0]
	if define.Word() != ":" { // Признак определения типа
		// log.Panicf("TVars.checkAsign(): bad assign(%v) for type\n", assign.Word())
		return listWord
	}
	// Получить описатель типа
	listWord = listWord[1:]
	return listWord
}

// Проверяет наличие слова RECORD (если есть -- добавляет всё до <END>)
func (sf *TVars) checkRecord(svar *srcvar.TSrcVar, listWord []types.IWord) (pl []types.IWord, res bool) {
	record := listWord[0]
	if !sf.keywords.IsKey("RECORD", record.Word()) {
		return listWord, false
	}
	svar.AddWord(listWord[0])
	listWord = listWord[1:]
	// Это запись, добавляем все до <END;>, если опять встретится RECORD
	// вызовем рекурсивно себя
	isRec := false
	for {
		word := listWord[0]
		if sf.keywords.IsKey("RECORD", word.Word()) { // Рекурсивный вызов
			listWord, isRec = sf.checkRecord(svar, listWord)
			if isRec {
				word = listWord[0]
			}
		}
		if !sf.keywords.IsKey("END", word.Word()) {
			svar.AddWord(listWord[0])
			listWord = listWord[1:]
			continue
		}
		// Если конец -- проверить разделитель
		term := listWord[1]
		if term.Word() == ";" {
			svar.AddWord(listWord[0])
			listWord = listWord[2:]
			return listWord, true
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
