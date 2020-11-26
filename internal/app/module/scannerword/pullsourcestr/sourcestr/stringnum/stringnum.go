package stringnum

/*
	Модуль предоставляет потокобезопасный тип для операций с номером строки.
*/

import (
	"fmt"
	"oc/internal/types"
)

//TStringNum -- тип для операций с номером строки
type TStringNum struct {
	val types.UStringNum // Хранимый номер строки
}

//New -- возвращает указатель  на новый TStringNum
func New(numStr types.UStringNum) (num *TStringNum, err error) {
	num = &TStringNum{}
	if err = num.Set(numStr); err != nil {
		return nil, fmt.Errorf("stringnum.go/New(): ERROR in set number string\n\t%v", err)
	}
	return num, nil
}

//Set -- установка значения номера строки
func (sf *TStringNum) Set(numStr types.UStringNum) (ош error) {
	if numStr <= 0 {
		return fmt.Errorf("TStringNum.Set(): ERROR numStr(%v)<1\n", numStr)
	}
	sf.val = numStr
	return nil
}

//Get -- возвращает хранимое значение номера строки
func (sf *TStringNum) Get() types.UStringNum {
	return sf.val
}

//String -- возвращает строковое представление номера строки
func (sf *TStringNum) String() string {
	return fmt.Sprint(sf.val)
}

//Inc -- увеличивает номер строки на +1
func (sf *TStringNum) Inc() {
	sf.val++
}

//Reset -- сбрасывает значение в "1"
func (sf *TStringNum) Reset() {
	sf.val = 1
}
