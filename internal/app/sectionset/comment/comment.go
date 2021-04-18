package comment

/*
	Пакет предоставляет тип секции модуля для комментариев.
*/

import (
	// "fmt"
	"log"

	"github.com/prospero78/goOC/internal/app/scanner/word"
)

// TComment -- операции с секцией коментариев
type TComment struct {
	poolWord []*word.TWord
}

// New -- возвращает новый *TComment
func New() *TComment {
	return &TComment{
		poolWord: make([]*word.TWord, 0),
	}
}

// Set -- выделяет из пула слов комментарии и возвращает остаток (кроме тех, что за концом модуля).
func (sf *TComment) Set(poolWord []*word.TWord) (poolRes []*word.TWord) {
	poolRes = make([]*word.TWord, 0)
	level := 0
	for len(poolWord) > 0 {
		word := poolWord[0]
		switch word.Word() {
		case "(*": // Начало комментариев
			poolWord = poolWord[1:]
			level++
			continue
		case "*)": // Конец комментариев
			poolWord = poolWord[1:]
			level--
			if level < 0 {
				log.Panicf("TComment.Set(): level(%v)<0\n", level)
			}
			continue
		default:
			if level > 0 {
				sf.poolWord = append(sf.poolWord, word)
				poolWord = poolWord[1:]
				continue
			}
		}
		poolWord = poolWord[1:]
		poolRes = append(poolRes, word)
	}
	// for _, word := range poolRes{
	// 	fmt.Printf("%v\t", word.Word())
	// }
	return poolRes
}

// AddPoolWord -- добавляет слова комментариев за концом модуля
func (sf TComment) AddPoolWord(poolWord []*word.TWord) int {
	if poolWord==nil{
		log.Panicf("TComment.AddPoolWord(): poolWord==nil\n")
	}
	sf.poolWord = append(sf.poolWord, poolWord...)
	// log.Printf("TComment.AddPoolWord(): len=%v", len(sf.poolWord))
	return len(sf.poolWord)
}
