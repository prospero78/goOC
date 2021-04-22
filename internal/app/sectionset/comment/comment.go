package comment

/*
	Пакет предоставляет тип секции модуля для комментариев.
*/

import (
	// "fmt"
	"log"

	"github.com/prospero78/goOC/internal/types"
)

// TComment -- операции с секцией коментариев
type TComment struct {
	listWord []types.IWord
}

// New -- возвращает новый *TComment
func New() *TComment {
	return &TComment{
		listWord: make([]types.IWord, 0),
	}
}

// Set -- выделяет из пула слов комментарии и возвращает остаток (кроме тех, что за концом модуля).
func (sf *TComment) Set(poolWord []types.IWord) (poolRes []types.IWord) {
	poolRes = make([]types.IWord, 0)
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
				sf.listWord = append(sf.listWord, word)
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
func (sf TComment) AddPoolWord(poolWord []types.IWord) int {
	if poolWord == nil {
		log.Panicf("TComment.AddPoolWord(): poolWord==nil\n")
	}
	sf.listWord = append(sf.listWord, poolWord...)
	// log.Printf("TComment.AddPoolWord(): len=%v", len(sf.poolWord))
	return len(sf.listWord)
}
