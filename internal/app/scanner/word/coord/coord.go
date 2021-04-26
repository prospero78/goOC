// Package coord -- координаты в исходном тексте
package coord

import (
	"fmt"

	"github.com/prospero78/goOC/internal/app/scanner/word/coord/litpos"
	"github.com/prospero78/goOC/internal/app/scanner/word/coord/numstr"
	"github.com/prospero78/goOC/internal/types"
)

// TCoord -- операции с координатами исходного текста
type TCoord struct {
	pos    types.IPos
	numStr types.INumStr
}

// New -- возвращает новый ICoord
func New(numStr types.ANumStr, pos types.APos) (types.ICoord, error) {
	_pos, err := litpos.New(pos)
	if err != nil {
		return nil, fmt.Errorf("coord.go/New(): in create pos\n\t%w", err)
	}
	_numStr, err := numstr.New(numStr)
	if err != nil {
		return nil, fmt.Errorf("coord.go/New(): in create numStr\n\t%w", err)
	}
	crd := &TCoord{
		pos:    _pos,
		numStr: _numStr,
	}
	return crd, nil
}

// SetPos -- устанавливает номер строки
func (sf *TCoord) SetPos(pos types.APos) error {
	if err := sf.pos.Set(pos); err != nil {
		return fmt.Errorf("TCoord.SetPos(): in set pos\n\t%w", err)
	}
	return nil
}

// SetNumStr -- устанавливает номер строки
func (sf *TCoord) SetNumStr(numStr types.ANumStr) error {
	if err := sf.numStr.Set(numStr); err != nil {
		return fmt.Errorf("TCoord.SetNumStr(): in set numStr\n\t%w", err)
	}
	return nil
}

// Set -- устанавливает координаты
func (sf *TCoord) Set(numStr types.ANumStr, pos types.APos) error {
	if err := sf.pos.Set(pos); err != nil {
		return fmt.Errorf("TCoord.Set(): in set pos\n\t%w", err)
	}
	if err := sf.numStr.Set(numStr); err != nil {
		return fmt.Errorf("TCoord.Set(): in set numStr\n\t%w", err)
	}
	return nil
}

// Pos -- возвращает хранимую позицию литеры в строке
func (sf *TCoord) Pos() types.APos {
	return sf.pos.Get()
}

// NumStr -- возвращает хранимый номер строки
func (sf *TCoord) NumStr() types.ANumStr {
	return sf.numStr.Get()
}
