// Package iscomment -- проверяет комментарии
package iscomment

import (
	"github.com/prospero78/goOC/internal/types"
)

// Выбирает открытие комментария
func CheckOpen(lit1, lit2 string, pos types.IPos) (strWord string, res bool) {
	if lit1 != "(" {
		return "", false
	}
	// pos.Inc()
	switch lit2 {
	case "*":
		// pos.Inc()
		return "(*", true
	default:
		return "(", true
	}
	// return "", false
}

// Выбирает закрытие комментария
func CheckClose(lit, lit2 string) bool {
	if lit != "*" {
		return false
	}
	if lit2 != ")" {
		// sf.addWord("*")
		return true
	}
	// sf.pos++
	// sf.addWord("*)")
	// sf.listRune = sf.listRune[1:]
	return true
}
