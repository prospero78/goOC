package sourcestr

/*
	Модуль предоставляет тип исходной строки.
	Можно прочитать, получить исхНомер, но нельзя изменить
*/

import (
	"fmt"
	мСн "oc/internal/app/module/scannerword/pullsourcestr/sourcestr/stringnum"
	мСткс "oc/internal/app/module/scannerword/pullsourcestr/sourcestr/stringtext"
	мТип "oc/internal/types"
)

//ТСтрокаИсх -- тип для операций со строкой исходника
type ТСтрокаИсх struct {
	строка *мСткс.ТСтрока
	номер  *мСн.ТСтрокаНом
}

//Нов -- возвращает указатель но новый ТСтрокаИсх
func Нов(пНом мТип.UStringNum, пСтр мТип.UStringSource) (стр *ТСтрокаИсх, ош error) {
	_стр := ТСтрокаИсх{}

	if _стр.номер, ош = мСн.Нов(пНом); ош != nil {
		return nil, fmt.Errorf("sourcestr.go/Нов(): ОШИБКА при установке номера исходной строки\n\t%v", ош)
	}
	if _стр.строка, ош = мСткс.Нов(пСтр); ош != nil {
		return nil, fmt.Errorf("sourcestr.go/Нов(): ОШИБКА при установке текста исходной строки\n\t%v", ош)
	}
	return &_стр, nil
}

//Строка -- возвращает хранимое значение исходной строки
func (сам *ТСтрокаИсх) Строка() мТип.UStringSource {
	return сам.строка.Получ()
}

//Номер -- возвращает хранимое значение номера исходной строки
func (сам *ТСтрокаИсх) Номер() мТип.UStringNum {
	return сам.номер.Получ()
}
