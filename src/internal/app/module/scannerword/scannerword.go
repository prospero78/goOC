package scannerword

/*
	Модуль предоставляет сканер слов в исходном модуле
*/

import (
	"fmt"
	мПулСтр "oc/internal/app/module/scannerword/pullsourcestr"
	мЛог "oc/internal/log"
	мТип "oc/internal/types"
)

//ТСканерСлов -- операции сканера с исходным текстом Оберон-файла
type ТСканерСлов struct {
	лог      *мЛог.ТЛог
	режим    int
	пулСтр   мТип.ИПулИсхСтр //Пул исходных строк для выведения координат
	пулСлова мТип.ИПулСлова  //Пул для хранения слов исходника
}

//Нов -- возвращает указатель на новый ТСканерСлов
func Нов(текстИсх мТип.СТекстИсх, режим int) (скан *ТСканерСлов, ош error) {
	_скан := ТСканерСлов{
		лог:    мЛог.Нов("ТСканерСлов", режим),
		режим:  режим,
		пулСтр: мПулСтр.Нов(текстИсх, режим),
	}
	if ош = _скан._Обработать(текстИсх); ош != nil {
		return nil, fmt.Errorf("scannerword.go/Нов(): ОШИБКА при работе сканера\n\t%v")
	}
	_скан.лог.Отладка("scaner.go/Нов", "Создание сканера")
	return &_скан, nil
}

//_Обработать -- обрабатывает исходный текст на наличие слов
func (сам *ТСканерСлов) _Обработать(текстИсх мТип.СТекстИсх) (ош error) {
	сам.лог.Отладка("_Обработать")
	return nil
}
