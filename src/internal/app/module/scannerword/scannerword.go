package scannerword

/*
	Модуль предоставляет сканер слов в исходном модуле
*/

import (
	"fmt"
	мКоорд "oc/internal/app/module/scannerword/coord"
	//мЛит "oc/internal/app/module/scannerword/litera"
	мПулСтр "oc/internal/app/module/scannerword/pullsourcestr"
	мПс "oc/internal/app/module/scannerword/pullword"
	мСлово "oc/internal/app/module/scannerword/word"
	мЛог "oc/internal/log"
	мТип "oc/internal/types"
)

//ТСканерСлов -- операции сканера с исходным текстом Оберон-файла
type ТСканерСлов struct {
	лог      *мЛог.ТЛог
	режим    int
	пулСтр   мТип.ИПулИсхСтр //Пул исходных строк для выведения координат
	пулСлова мТип.ИПулСлова  //Пул для хранения слов исходника
	текстИсх мТип.СТекстИсх  //Исходный текст для анализа
	коорд    *мКоорд.ТКоорд  //Текущие координаты в тексте
}

//Нов -- возвращает указатель на новый ТСканерСлов
func Нов(текстИсх мТип.СТекстИсх, режим int) (скан *ТСканерСлов, ош error) {
	_скан := ТСканерСлов{
		лог:      мЛог.Нов("ТСканерСлов", режим),
		режим:    режим,
		пулСтр:   мПулСтр.Нов(текстИсх, режим),
		текстИсх: текстИсх,
		пулСлова: мПс.Нов(),
	}
	if _скан.коорд, ош = мКоорд.Нов(1, 0); ош != nil {
		return nil, fmt.Errorf("scannerword.go/Нов(): ОШИБКА при создании координат сканера\n\t%v", ош)
	}
	if ош = _скан._Обработать(текстИсх); ош != nil {
		return nil, fmt.Errorf("scannerword.go/Нов(): ОШИБКА при работе сканера\n\t%v", ош)
	}
	_скан.лог.Отладка("scaner.go/Нов", "Создание сканера")
	_скан.лог.Отладка("scaner.go/Нов", "Всего слов", _скан.пулСлова.Len())
	return &_скан, nil
}

//Добавляет новое слово в список слов
func (сам *ТСканерСлов) _СловоДобав(пСлово мТип.ССлово) (ош error) {
	коорд, ош := мКоорд.Нов(сам.коорд.Стр(), сам.коорд.Поз())
	if ош != nil {
		return fmt.Errorf("ТСканерСлов._СловоДобав(): ОШИБКА при создании координат слова\n\t%v", ош)
	}
	строка, ош := сам.пулСтр.Строка(сам.коорд.Стр())
	if ош != nil {
		return fmt.Errorf("ТСканерСлов._СловоДобав(): ОШИБКА при получении строки\n\t%v", ош)
	}
	слово, ош := мСлово.Нов(коорд, пСлово, строка)
	if ош != nil {
		return fmt.Errorf("ТСканерСлов._СловоДобав(): ОШИБКА при создании слова(%v)\n\t%v", пСлово, ош)
	}
	сам.пулСлова.Доб(слово)
	return nil
}

//_Обработать -- обрабатывает исходный текст на наличие слов, выделяет слова, расставляет словам
//   номера строк, позицию в строке. Тип не заполняет так как нужен семантический анализ (может это
//   тупо комментарии или строки).
func (сам *ТСканерСлов) _Обработать(текстИсх мТип.СТекстИсх) (ош error) {
	сам.лог.Отладка("_Обработать")
	словоНов := мТип.ССлово("")
	_СловоНов := func() error {
		if словоНов != "" {
			if ош = сам._СловоДобав(словоНов); ош != nil {
				return fmt.Errorf("ТСканерСлов._Обработать()._СловоНов(): ОШИБКА при добавлении слова %q\n\t%v", словоНов, ош)
			}
		}
		словоНов = ""
		return nil
	}
	//перебор всех литер в исходнике на предмет получить все слова
	for _, руна := range сам.текстИсх {
		//Контролировать формирование слов по разделителям слов
		switch _лит := мТип.ССлово(руна); _лит {
		case "\n": //Увеличить номер строки, сбросить позицию
			_СловоНов()
			сам.коорд.СтрДоб()
			сам.коорд.ПозСброс()
		case "\t", " ", "\r": //Начало нового слова
			_СловоНов()
			сам.коорд.ПозДоб()
		default:
			словоНов += _лит
		}
	}
	return nil
}

//Слова -- возвращает хранимый пул слов
func (сам *ТСканерСлов) ПулСлова() мТип.ИПулСлова {
	return сам.пулСлова
}
