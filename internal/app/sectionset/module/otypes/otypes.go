package otypes

/*
	Пакет предоставляет тип для формирования секции Оберон-типов
*/

import (
	"log"

	"github.com/prospero78/goOC/internal/app/sectionset/module/keywords"
	"github.com/prospero78/goOC/internal/types"

	"github.com/prospero78/goOC/internal/app/sectionset/module/otypes/srctype.go"
)

// TOtypes -- операци ис секцией типов
type TOtypes struct {
	keywords types.IKeywords
	poolType []*srctype.TSrcType
}

// New -- возвращает новый *TOtypes
func New() *TOtypes {
	return &TOtypes{
		keywords: keywords.GetKeys(),
		poolType: make([]*srctype.TSrcType, 0),
	}
}

// Split -- вырезает слова секции типов, остаток возвращает
func (sf *TOtypes) Split(listWord []types.IWord) []types.IWord {
	// Убедиться, что есть секция типов
	word := listWord[0]
	if !sf.keywords.IsKey("TYPE", word.Word()) {
		return listWord
	}
	listWord = listWord[1:]
	for len(listWord) > 4 {
		// Получить слова типа, признак: <; name = >
		for len(listWord) > 0 {
			word := listWord[0]
			if sf.keywords.IsKey("VAR", word.Word()) {
				return listWord
			}
			if sf.keywords.IsKey("PROCEDURE", word.Word()) {
				return listWord
			}
			if sf.keywords.IsKey("BEGIN", word.Word()) {
				return listWord
			}
			listWord = sf.fillTypes(listWord)
		}
	}
	log.Panicf("TOtypes.Split(): not have end for TYPE\n")
	return nil
}

// Заполняет один тип, возвращает остаток
func (sf *TOtypes) fillTypes(listWord []types.IWord) []types.IWord {
	name := listWord[0] // Проверить на разделитель
	otp := srctype.New(name)
	sf.poolType = append(sf.poolType, otp)
	if !name.IsName() {
		log.Panicf("TOtypes.fillTypes(): bad name(%v) for type\n", name.Word())
	}
	listWord = listWord[1:]
	isRec := false
	for {
		listWord = sf.checkExport(otp, listWord)
		listWord = sf.checkAsign(listWord)
		listWord, isRec = sf.checkRecord(otp, listWord)
		if isRec {
			return listWord
		}
		term := listWord[1]
		if term.Word() != ";" {
			otp.AddWord(listWord[0])
			listWord = listWord[1:]
			continue
		}
		nameNext := listWord[2]
		if !nameNext.IsName() {
			log.Panicf("TOtypes.fiilTypes(): nameNext=%q\n", nameNext.Word())
		}
		term = listWord[3]
		if !(term.Word() == "=" || term.Word() == "*") {
			// log.Panicf("TOtypes.fiilTypes(): term=%q\n", term.Word())
			otp.AddWord(listWord[0])
			otp.AddWord(listWord[1])
			otp.AddWord(listWord[2])
			otp.AddWord(listWord[3])
			listWord = listWord[4:]
			continue
		}
		otp.AddWord(listWord[0])
		listWord = listWord[2:] // отбросить слово типа вместе с окончательным разделителем
		return listWord
	}
}

// проверяет наличие экспорта в списке слов
func (sf *TOtypes) checkExport(otp *srctype.TSrcType, listWord []types.IWord) []types.IWord {
	exp := listWord[0]
	if exp.Word() == "*" {
		otp.SetExport()
		listWord = listWord[1:]
	}
	return listWord
}

// Проверяет наличие присвоения в указанной позиции
func (sf *TOtypes) checkAsign(listWord []types.IWord) (pl []types.IWord) {
	assign := listWord[0]
	if assign.Word() != "=" { // Признак определения типа
		// log.Panicf("TOtypes.checkAsign(): bad assign(%v) for type\n", assign.Word())
		return listWord
	}
	// Получить описатель типа
	listWord = listWord[1:]
	return listWord
}

// Проверяет наличие слова RECORD (если есть -- добавляет всё до <END>)
func (sf *TOtypes) checkRecord(otp *srctype.TSrcType, listWord []types.IWord) (pl []types.IWord, res bool) {
	record := listWord[0]
	if !sf.keywords.IsKey("RECORD", record.Word()) {
		return listWord, false
	}
	otp.AddWord(listWord[0])
	listWord = listWord[1:]
	// Это запись, добавляем все до <END;>, если опять встретится RECORD
	// вызовем рекурсивно себя
	isRec := false
	for {
		word := listWord[0]
		if sf.keywords.IsKey("RECORD", word.Word()) { // Рекурсивный вызов
			listWord, isRec = sf.checkRecord(otp, listWord)
			if isRec {
				word = listWord[0]
			}
		}
		if !sf.keywords.IsKey("END", word.Word()) {
			otp.AddWord(listWord[0])
			listWord = listWord[1:]
			continue
		}
		// Если конец -- проверить разделитель
		term := listWord[1]
		if term.Word() == ";" {
			otp.AddWord(listWord[0])
			listWord = listWord[2:]
			return listWord, true
		}
		log.Panicf("TOtypes.checkRecord(): unknown word=%+v\n", term)
	}
}

// Len -- возвращает количество типов
func (sf *TOtypes) Len() int {
	// for _, typ := range sf.poolType {
	// 	fmt.Printf("%v = ", typ.Name())
	// 	for _, word := range typ.Words() {
	// 		fmt.Printf("%v ", word.Word())
	// 	}
	// 	fmt.Print("\n")
	// }
	return len(sf.poolType)
}
