package sourcestr

/*
	Модуль предоставляет тип исходной строки.
	Можно получить строку исходника, её номер.
	Но в строке исходника ничего изменит ьнельзя
*/

import (
	"fmt"
	мСн "oc/internal/app/module/scannerword/pullsourcestr/sourcestr/stringnum"
	мСткс "oc/internal/app/module/scannerword/pullsourcestr/sourcestr/stringtext"
	"oc/internal/types"
)

//ТСтрокаИсх -- тип для операций со строкой исходника
type ТСтрокаИсх struct {
	strSource *types.UStringSource // Исходная строка
	num       *types.UStringNum    // Номер строки
}

//Нов -- возвращает указатель но новый ТСтрокаИсх
func Нов(пНом types.UStringNum, пСтр types.UStringSource) (стр *ТСтрокаИсх, ош error) {
	_стр := ТСтрокаИсх{}

	if _стр.num, ош = мСн.Нов(пНом); ош != nil {
		return nil, fmt.Errorf("sourcestr.go/Нов(): ERROR при установке номера исходной строки\n\t%v", ош)
	}
	if _стр.strSource, ош = мСткс.Нов(пСтр); ош != nil {
		return nil, fmt.Errorf("sourcestr.go/Нов(): ERROR при установке текста исходной строки\n\t%v", ош)
	}
	return &_стр, nil
}

//Строка -- возвращает хранимое значение исходной строки
func (сам *ТСтрокаИсх) Строка() types.UStringSource {
	return сам.strSource.Получ()
}

//Номер -- возвращает хранимое значение номера исходной строки
func (сам *ТСтрокаИсх) Номер() types.UStringNum {
	return сам.num.Получ()
}
