// Package litpos -- тип для позиции литеры в строке
package litpos

import (
	"fmt"

	"github.com/prospero78/goOC/internal/types"
)

// TPos -- операции с координатой литеры в строке
type TPos struct {
	val types.APos
}

// New -- возвращает новый IPos
func New(pos types.APos) (_pos types.IPos, err error) {
	if pos < 0 {
		return nil, fmt.Errorf("litpos.go/New(): pos(%v)<0", pos)
	}
	_pos = &TPos{
		val: pos,
	}
	return _pos, nil
}

// Get -- возвращает хранимое значение кординаты
func (sf *TPos) Get() types.APos {
	return sf.val
}

// Set -- устанавливает хранимое значение
func (sf *TPos) Set(val types.APos) error {
	if val < 0 {
		return fmt.Errorf("litpos.go/New(): pos(%v)<0", val)
	}
	sf.val = val
	return nil
}
