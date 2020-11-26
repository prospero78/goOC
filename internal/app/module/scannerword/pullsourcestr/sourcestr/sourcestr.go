package sourcestr

/*
	Модуль предоставляет тип исходной строки.
	Можно получить строку исходника, её номер.
	Но в строке исходника ничего изменит ьнельзя
*/

import (
	"fmt"
	"oc/internal/types"
)

//ТСтрокаИсх -- тип для операций со строкой исходника
type ТСтрокаИсх struct {
	strSource types.UStringSource // Исходная строка
	num       types.UStringNum    // Номер строки
}

//Нов -- возвращает указатель но новый ТСтрокаИсх
func Нов(num types.UStringNum, strSource types.UStringSource) (стр *ТСтрокаИсх, ош error) {
	if num < 1 {
		return nil, fmt.Errorf("sourcestr.go/Нов(): ERROR num(%v)<1", num)
	}
	_стр := ТСтрокаИсх{
		strSource: strSource,
		num:       num,
	}
	return &_стр, nil
}

//String -- возвращает хранимое значение исходной строки
func (сам *ТСтрокаИсх) String() types.UStringSource {
	return сам.strSource
}

//Num -- возвращает хранимое значение номера исходной строки
func (сам *ТСтрокаИсх) Num() types.UStringNum {
	return сам.num
}
