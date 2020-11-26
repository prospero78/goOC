package coord

/*
	Модуль предоставляет потокобезопасный тип координаты.
*/

import (
	"fmt"
	"oc/internal/app/module/scannerword/coord/num"
	"oc/internal/app/module/scannerword/coord/pos"
)

//TCoord -- тип для хранения координат
type TCoord struct {
	num *num.TNum
	pos *pos.TPos
}

//New -- возвращает указатель на новый TCoord
func New(numStr мТип.UStringNum, posStr мТип.UStringPos) (coord *TCoord, err error) {
	coord := &TCoord{}
	if coord.num, err = num.New(numStr); err != nil {
		return nil, fmt.Errorf("coord.go/New(): ERROR in create num source string\n\t%v", err)
	}
	if coord.pos, err = pos.New(posStr); err != nil {
		return nil, fmt.Errorf("coord.go/New(): ERROR in create pos source string\n\t%v", err)
	}
	return coord, nil
}

//Возвращает строковое представление координаты
func (sf *TCoord) String() string {
	return fmt.Sprintf("Coord=%v:%v", sf.num, sf.pos)
}

//Num -- возвращает хранимый номера строки
func (sf *TCoord) Num() мТип.UStringNum {
	return sf.num.Get()
}

//NumInc -- добавлет +1 номер строки
func (sf *TCoord) NumInc() {
	sf.num.Inc()
}

//NumSet -- устанавливает номер строки
func (sf *TCoord) NumSet(num мТип.UStringNum) (err error) {
	if err = sf.num.Set(num); err != nil {
		return fmt.Errorf("TCoord.NumSet(): ERROR in set num source string\n\t%v", err)
	}
	return nil
}

//NumReset -- сброасывает номер строки
func (sf *TCoord) NumReset() {
	sf.num.Reset()
}

//Pos -- возвращает хранимую позицию в строке
func (sf *TCoord) Pos() types.UStringPos {
	return sf.pos.Get()
}

//PosReset -- сбрасывает позицию строки
func (sf *TCoord) PosReset() {
	sf.pos.Reset()
}

//PosInc -- добавлет +1 позицию в строке
func (sf *TCoord) PosInc() {
	sf.pos.Inc()
}

//PosSet -- устанавливает позицию строки
func (sf *TCoord) PosSet(pos мТип.UStringPos) (err error) {
	if err = sf.pos.Уст(pos); err != nil {
		return fmt.Errorf("TCoord.PosSet(): ERROR in set pos in source string\n\t%v", err)
	}
	return nil
}
