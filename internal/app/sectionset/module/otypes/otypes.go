package otypes

/*
	Пакет предоставляет тип для формирования секции Оберон-типов
*/

import (
	"log"
	"oc/internal/app/scanner/word"
	"oc/internal/app/sectionset/module/keywords"

	"oc/internal/app/sectionset/module/otypes/srctype.go"
)

// TOtypes -- операци ис секцией типов
type TOtypes struct {
	keywords *keywords.TKeywords
}

// New -- возвращает новый *TOtypes
func New() *TOtypes {
	return &TOtypes{
		keywords: keywords.New(),
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
		name := pool[0]
		// Проверить на пустую секцию
		if sf.keywords.IsKey("VAR", name.Word()) {
			return pool
		}
		if sf.keywords.IsKey("PROCEDURE", name.Word()) {
			return pool
		}
		if sf.keywords.IsKey("BEGIN", name.Word()) {
			return pool
		}
		if !name.IsName() {
			log.Panicf("TOtypes.Split(): bad name for type(%v)\n", name.Word())
		}
		pool = pool[1:]

		// Проверить на разделитель
		otp := srctype.New(name)
		term := pool[0]
		switch term.Word() {
		case "*": // Признак экспорта типа
			otp.SetExport()
			pool = pool[1:]
			term := pool[0]
			if term.Word() != "=" {
				log.Panicf("TOtypes.Split(): bad term(%v) for type after export\n", term.Word())
			}
			pool = pool[1:]
		case "=": //Признак определения типа
			// Получить описатель типа
			pool = pool[1:]
		default:
			log.Panicf("TOtypes.Split(): bad term(%v) for type\n", term.Word())
		}

		// Получить слова типа
		for len(pool) > 0 {
			word = pool[0]
			if word.Word() == ";" {
				break
			}
			// Возможны варианты продолжения типа, надо проверять
			if sf.keywords.IsKey("POINTER", word.Word()) {
				pool = pool[1:]
				sf.checkPointer(pool, otp)
			}
			panic("доделать!!!")
			otp.AddWord(word)
			pool = pool[1:]
		}
		if len(pool) == 0 {
			log.Panicf("TOtypes.Split(): not have end for TYPE %q\n", name.Word())
		}
		term = pool[0]
		if term.Word() != ";" {
			log.Panicf("TOtypes.Split(): bad term(%v) for type after defenicion\n", term.Word())
		}
		pool = pool[1:]
	}
	log.Panicf("TOtypes.Split(): not have end for TYPE\n")
	return nil
}

// Проверяет определение типа, если встретилось "POINTER"
func (sf TOtypes) checkPointer(pool []*word.TWord, otp *srctype.TSrcType) []*word.TWord {
	// Первым должно идти слово "TO"
	word := pool[0]
	if !sf.keywords.IsKey("TO", word.Word()) {
		log.Panicf("TOtypes.checkPointer(): word(%v) must by 'TO'\n", word.Word())
	}
	pool=pool[1:]
	// Дальше должно идти: или имя или составное имя

	return nil
}
