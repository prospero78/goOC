package num

/*
	Модуль предоставляет тип для операций с номером строки.
	(нужен для координат слов)
*/

import (
	"fmt"
	"oc/internal/types"
)

//TNum -- тип для операций с номером строки
type TNum struct {
	val types.UStringNum // Хранимый номер строки
}

//New -- возвращает указатель  на новый TNum
func New(numStr types.UStringNum) (num *TNum, err error) {
	num = &TNum{}
	if err = num.Set(numStr); err != nil {
		return nil, fmt.Errorf("stringnum.go/New(): ERROR in set number string\n\t%v", err)
	}
	return num, nil
}

//Set -- установка значения номера строки
func (sf *TNum) Set(numStr types.UStringNum) (ош error) {
	if numStr <= 0 {
		return fmt.Errorf("TNum.Set(): ERROR numStr(%v)<1\n", numStr)
	}
	sf.val = numStr
	return nil
}

//Get -- возвращает хранимое значение номера строки
func (sf *TNum) Get() types.UStringNum {
	return sf.val
}

//String -- возвращает строковое представление номера строки
func (sf *TNum) String() string {
	return fmt.Sprint(sf.val)
}

//Inc -- увеличивает номер строки на +1
func (sf *TNum) Inc() {
	sf.val++
}

//Reset -- сбрасывает значение в "1"
func (sf *TNum) Reset() {
	sf.val = 1
}
