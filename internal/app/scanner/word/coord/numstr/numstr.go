// Package numstr -- номер строки слова
package numstr

import (
	"fmt"

	"github.com/prospero78/goOC/internal/types"
)

// TNumStr -- операции с номером строки
type TNumStr struct {
	val types.ANumStr
}

// New -- возвращает новый INumStr
func New(val types.ANumStr) (types.INumStr, error) {
	if val < 1 {
		return nil, fmt.Errorf("numstr.go/New(): val(%v)<1", val)
	}
	ns := &TNumStr{
		val: val,
	}
	return ns, nil
}

// Get -- возвращает хранимое значение номера строки
func (sf *TNumStr) Get() types.ANumStr {
	return sf.val
}

// Set -- устанавливает значение хранимой строки
func (sf *TNumStr) Set(val types.ANumStr) error {
	if val < 1 {
		return fmt.Errorf("TNumStr.Set(): val(%v)<1", val)
	}
	sf.val = val
	return nil
}
