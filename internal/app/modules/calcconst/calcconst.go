// Package calcconst -- вычислитель констант
package calcconst

import (
	"log"
	"github.com/prospero78/goOC/internal/app/modules/calcexp"
	"github.com/prospero78/goOC/internal/app/modules/calcword"
	"github.com/prospero78/goOC/internal/app/scanner/word"
	"github.com/prospero78/goOC/internal/app/sectionset/module/consts/srcconst"
)

// TCalcConst -- операци ипо вычислению констант
type TCalcConst struct {
	wordCurrent *word.TWord
	calcWord    *calcword.TCalcWord // Вычисление слова
	calcExp     *calcexp.TCalcExp   // Вычисление выражения
}

// New -- возвращает новый *TCalcConst
func New() *TCalcConst {
	return &TCalcConst{
		calcWord: calcword.New(),
		calcExp:  calcexp.New(),
	}
}

// Calc -- рассчитывает константу
func (sf *TCalcConst) Calc(cons *srcconst.TConst) {
	pool := cons.GetWords()
	if len(pool) == 0 { // У константы нет имени. Теоретически, это невозможно
		log.Panicf("TModules.processConstant(): const(%v) not have type\n", cons.Name())
	}
	if cons.Name() == "\"цЯблоки\"" {
		log.Print("")
	}
	lenPool := len(pool)
	fnCheckWord := func() bool {
		adr := 0
		for {
			pool = cons.GetWords()
			if adr >= len(pool) {
				sf.setType(cons)
				return false
			}
			word := pool[adr]
			adr++
			sf.calcWord.RecognizeType(word)
			_lenPool := len(pool)
			if _lenPool < lenPool {
				lenPool = _lenPool
				return true
			}
		}
	}

	if len(pool) == 1 {
		word := pool[0]
		sf.calcWord.RecognizeType(word)
		sf.setType(cons)
		return
	}
	for fnCheckWord() {
	}
}

// После обработки всех слов константы -- устанавливает её тип
func (sf *TCalcConst) setType(cons *srcconst.TConst) {
	pool := cons.GetWords()
	switch len(pool) {
	case 0: // Нет слов у константы (теоретически такого быть не может)
		log.Panicf("TModules.processConstant(): const(%v.%v) not have type\n", cons.Module(), cons.Name())
	case 1: // Тип константы определяется единственным словом
		cons.SetType(pool[0].GetType())
	default: // Тип имеет выражение и его надо вычислить
		// exp := sf.consCurrent.GetExpres()
		// sf.exprConstCalc(exp)
		poolWord := cons.GetWords()
		poolWord = poolWord[1:] // Откинуть открывающую скобку
		for len(poolWord) > 0 {
			word := poolWord[0]
			sf.calcExp.AddWord(word)
			poolWord = poolWord[1:]
			if word.Word() == ")" {
				break
			}
		}
		
		// sf.exprConstCalc()
		sf.calcExp.RecognizeType()
		
		
		
		// После передачи слов в выражение -- надо сформировать новый словарь слов
		poolNew := make([]*word.TWord, 0)
		poolNew = append(poolNew, poolWord...)
		cons.SetPoolWord(poolNew)
		sf.calcExp.Calc()
	}
}
