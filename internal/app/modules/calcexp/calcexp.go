// Package calcexp -- вычисляет выражение  (тип и значение)
package calcexp

import (
	"fmt"
	"log"
	"strconv"

	"github.com/prospero78/goOC/internal/app/modules/calcword"
	"github.com/prospero78/goOC/internal/app/scanner/word"
	"github.com/prospero78/goOC/internal/types"
)

// TCalcExp -- операции с выражением
type TCalcExp struct {
	val      types.IWord
	listWord []types.IWord
	calcWord *calcword.TCalcWord
}

// New -- возвращает новый *TCalcExp
func New() *TCalcExp {
	return &TCalcExp{
		listWord: make([]types.IWord, 0),
		calcWord: calcword.New(),
	}
}

// AddWord -- добавляет слово в выражение
func (sf *TCalcExp) AddWord(word types.IWord) {
	if word == nil {
		log.Panicf("TCalcExp.AddWord(): word==nil\n")
	}
	sf.listWord = append(sf.listWord, word)
}

// RecognizeType -- распознаёт тип выражения
func (sf *TCalcExp) RecognizeType() (name, operation string) {
	// log.Panicf("TCalcExp.RecognizeType(): доделать\n")
	for _, word := range sf.listWord {
		if word.GetType() == "" {
			name, operation = sf.calcWord.RecognizeType(word)
			if name != "" || operation != "" {
				return name, operation
			}
		}
		if sf.val == nil { // Первое присовение типа выражения
			sf.val = word
			continue
		}
		if sf.val.GetType() != word.GetType() {
			log.Panicf("TCalcExp.RecognizeType(): %q(%v:%v) not union type(exp=%q word=%q)\n", *sf.val.Module(), sf.val.NumStr(), sf.val.Pos(), sf.val.GetType(), word.GetType())
		}
	}
	return "", ""
}

// Calc -- рассчитывает выражение
func (sf *TCalcExp) Calc() {
	// log.Panicf("TCalcExp.Calc(): доделать\n")
	op := ""
	for _, wrd := range sf.listWord {
		// выполнить операции в соответствии со своим типом
		switch wrd.GetType() {
		case "INTEGER":
			switch op {
			case "+":
				sf.intPlus(wrd)
			case "-":
				sf.intMinus(wrd)
			}
		case "STRING":
			switch op {
			case "+":
				log.Panicf("TCalcExp.Calc(): доделать\n")
			case "-":
				log.Panicf("TCalcExp.Calc(): доделать\n")
			}
		}
	}
}

// Складывает целые числа
func (sf *TCalcExp) intPlus(wrd types.IWord) {
	res := sf.val.Word()
	resNum, err := strconv.Atoi(string(res))
	if err != nil {
		log.Panicf("TCalcExp.Calc(): %q(%v:%v) val exp(%v) not number\n", *sf.val.Module(), sf.val.NumStr(), sf.val.Pos(), sf.val.Word())
	}
	res = wrd.Word()
	resNum2, err := strconv.Atoi(string(res))
	if err != nil {
		log.Panicf("TCalcExp.Calc(): %q(%v:%v) word(%v) not number\n", *sf.val.Module(), sf.val.NumStr(), sf.val.Pos(), wrd.Word())
	}
	resNum += resNum2
	wrd, err = word.New(wrd.NumStr(), wrd.Pos(), types.AWord(fmt.Sprint(resNum)))
	if err != nil {
		log.Panicf("TCalcExp.Calc(): in create word\n\t%v", err)
	}
	wrd.SetModule(sf.val.Module())
	sf.val = wrd
}

// Вычитает целые числа
func (sf *TCalcExp) intMinus(word types.IWord) {
	log.Panicf("TCalcExp.intMinus(): доделать\n")
}
