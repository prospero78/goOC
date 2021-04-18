package otypes

/*
	Пакет предоставляет тип для формирования секции Оберон-типов
*/

import (
	"log"

	"github.com/prospero78/goOC/internal/app/scanner/word"
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
func (sf *TOtypes) Split(pool []*word.TWord) []*word.TWord {
	// Убедиться, что есть секция типов
	word := pool[0]
	if !sf.keywords.IsKey("TYPE", word.Word()) {
		return pool
	}
	pool = pool[1:]
	for len(pool) > 4 {
		// Получить слова типа, признак: <; name = >
		for len(pool) > 0 {
			word := pool[0]
			if sf.keywords.IsKey("VAR", word.Word()) {
				return pool
			}
			if sf.keywords.IsKey("PROCEDURE", word.Word()) {
				return pool
			}
			if sf.keywords.IsKey("BEGIN", word.Word()) {
				return pool
			}
			pool = sf.fillTypes(pool)
		}
	}
	log.Panicf("TOtypes.Split(): not have end for TYPE\n")
	return nil
}

// Заполняет один тип, возвращает остаток
func (sf *TOtypes) fillTypes(pool []*word.TWord) []*word.TWord {
	name := pool[0] // Проверить на разделитель
	otp := srctype.New(name)
	sf.poolType = append(sf.poolType, otp)
	if !name.IsName() {
		log.Panicf("TOtypes.fillTypes(): bad name(%v) for type\n", name.Word())
	}
	pool = pool[1:]
	isRec := false
	for {
		pool = sf.checkExport(otp, pool)
		pool = sf.checkAsign(pool)
		pool, isRec = sf.checkRecord(otp, pool)
		if isRec {
			return pool
		}
		term := pool[1]
		if term.Word() != ";" {
			otp.AddWord(pool[0])
			pool = pool[1:]
			continue
		}
		nameNext := pool[2]
		if !nameNext.IsName() {
			log.Panicf("TOtypes.fiilTypes(): nameNext=%q\n", nameNext.Word())
		}
		term = pool[3]
		if !(term.Word() == "=" || term.Word() == "*") {
			// log.Panicf("TOtypes.fiilTypes(): term=%q\n", term.Word())
			otp.AddWord(pool[0])
			otp.AddWord(pool[1])
			otp.AddWord(pool[2])
			otp.AddWord(pool[3])
			pool = pool[4:]
			continue
		}
		otp.AddWord(pool[0])
		pool = pool[2:] // отбросить слово типа вместе с окончательным разделителем
		return pool
	}
}

// проверяет наличие экспорта в списке слов
func (sf *TOtypes) checkExport(otp *srctype.TSrcType, pool []*word.TWord) []*word.TWord {
	exp := pool[0]
	if exp.Word() == "*" {
		otp.SetExport()
		pool = pool[1:]
	}
	return pool
}

// Проверяет наличие присвоения в указанной позиции
func (sf *TOtypes) checkAsign(pool []*word.TWord) (pl []*word.TWord) {
	assign := pool[0]
	if assign.Word() != "=" { // Признак определения типа
		// log.Panicf("TOtypes.checkAsign(): bad assign(%v) for type\n", assign.Word())
		return pool
	}
	// Получить описатель типа
	pool = pool[1:]
	return pool
}

// Проверяет наличие слова RECORD (если есть -- добавляет всё до <END>)
func (sf *TOtypes) checkRecord(otp *srctype.TSrcType, pool []*word.TWord) (pl []*word.TWord, res bool) {
	record := pool[0]
	if !sf.keywords.IsKey("RECORD", record.Word()) {
		return pool, false
	}
	otp.AddWord(pool[0])
	pool = pool[1:]
	// Это запись, добавляем все до <END;>, если опять встретится RECORD
	// вызовем рекурсивно себя
	isRec := false
	for {
		word := pool[0]
		if sf.keywords.IsKey("RECORD", word.Word()) { // Рекурсивный вызов
			pool, isRec = sf.checkRecord(otp, pool)
			if isRec {
				word = pool[0]
			}
		}
		if !sf.keywords.IsKey("END", word.Word()) {
			otp.AddWord(pool[0])
			pool = pool[1:]
			continue
		}
		// Если конец -- проверить разделитель
		term := pool[1]
		if term.Word() == ";" {
			otp.AddWord(pool[0])
			pool = pool[2:]
			return pool, true
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
