package consts

/*
	Пакет предоставляет тип для разбора секции констант
*/

import (
	"log"
	"oc/internal/app/scanner/word"
	"oc/internal/app/sectionset/module/consts/srcconst"
	"oc/internal/app/sectionset/module/keywords"
)

// TConsts -- операции с секцией констант
type TConsts struct {
	keywords  *keywords.TKeywords
	poolConst []*srcconst.TConst
}

// New -- возвращает новый *TConsts
func New() *TConsts {
	return &TConsts{
		keywords:  keywords.Keys,
		poolConst: make([]*srcconst.TConst, 0),
	}
}

// Split -- вырезает секцию констант, если она есть
func (sf *TConsts) Split(pool []*word.TWord) []*word.TWord {
	word := pool[0]
	if !sf.keywords.IsKey("CONST", word.Word()) {
		return pool
	}
	pool = pool[1:]
	for len(pool) >= 4 {
		name := pool[0]
		if !name.IsName() {
			log.Panicf("TConsts.Split(): not valid name(%v)\n", name.Word())
		}
		// Фильтр на следующие секции
		if sf.keywords.IsKey("TYPE", name.Word()) { // Началась секция типов
			return pool
		}
		if sf.keywords.IsKey("VAR", name.Word()) { // Началась секция типов
			return pool
		}
		if sf.keywords.IsKey("PROCEDURE", name.Word()) { // Началась секция типов
			return pool
		}
		if sf.keywords.IsKey("BEGIN", name.Word()) { // Началась секция типов
			return pool
		}
		pool = pool[1:]
		term := pool[0]
		cons := srcconst.New(name)
		switch term.Word() {
		case "*": // Экспорт константы
			cons.SetExport()
			pool = pool[1:]
			term := pool[0]
			if term.Word() != "=" {
				log.Panicf("TConsts.Split(): not valid term(%v)\n", term.Word())
			}
			pool = pool[1:]
		case "=": // Правильный терминал
			pool = pool[1:]
		default:
			log.Panicf("TConsts.Split(): not valid term(%v)\n", term.Word())
		}

		// Теперь надо установить выражение константы
		for len(pool) > 0 {
			val := pool[0]
			if val.Word() == ";" { // Конец выражения
				break
			}
			cons.AddWord(val) // Продолжается выражение
			pool = pool[1:]
		}
		if len(pool) == 0 {
			log.Panicf("TConsts.Split(): not end CONST\n")
		}
		delim := pool[0]
		if delim.Word() != ";" { // Следующий элемент секции констант
			log.Panicf("TConsts.Split(): not valid delimeter const(%v)\n", delim.Word())
		}
		sf.poolConst = append(sf.poolConst, cons)
		pool = pool[1:]
	}

	log.Panicf("TConsts.Split(): not have CONSTS\n")
	return nil
}

// Len -- возвращает количество констант.
func (sf *TConsts) Len() int {
	return len(sf.poolConst)
}

// Get -- возвращает пул констант
func (sf *TConsts) Get() []*srcconst.TConst {
	return sf.poolConst
}
