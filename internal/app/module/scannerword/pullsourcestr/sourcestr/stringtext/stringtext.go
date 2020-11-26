package stringtext

/*
	Модуль предоставляет тип для безопасного хранения
	строки исходного кода
*/

import (
	"fmt"
	"oc/internal/types"
)

//TSourceString -- тип для хранения строки исходного кода
type TSourceString struct {
	val types.UStringSource
}

//New -- возвращает ссылку на новый TSourceString
func New(strSource types.UStringSource) (src *TSourceString, err error) {
	if strSource == "" {
		return nil, fmt.Errorf("stringtext.go/New(): ERROR strSource empty")
	}
	src = &TSourceString{
		val: strSource,
	}
	return src, nil
}

//Get -- возвращает хранимое значение строки исходника TSourceString
func (sf *TSourceString) Get() types.UStringSource {
	return sf.val
}
