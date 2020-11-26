package scannerword

/*
	Модуль предоставляет сканер слов в исходном модуле
*/

import (
	"fmt"
	"oc/internal/app/module/scannerword/coord"

	//мЛит "oc/internal/app/module/scannerword/litera"
	мПулСтр "oc/internal/app/module/scannerword/pullsourcestr"
	мПс "oc/internal/app/module/scannerword/pullword"
	мСлово "oc/internal/app/module/scannerword/word"
	мЛог "oc/internal/log"
	мТип "oc/internal/types"
)

//TScannerWord -- операции сканера с исходным текстом Оберон-файла
type TScannerWord struct {
	лог      *мЛог.ТЛог
	режим    int
	пулСтр   мТип.IPoolStringSource //Пул исходных строк для выведения координат
	poolWord мТип.IPoolWord         //Пул для хранения слов исходника
	текстИсх мТип.UTextSource       //Исходный текст для анализа
	коорд    *coord.ТКоорд          //Текущие координаты в тексте
}

//Нов -- возвращает указатель на новый TScannerWord
func Нов(текстИсх мТип.UTextSource, режим int) (скан *TScannerWord, ош error) {
	_скан := TScannerWord{
		лог:      мЛог.Нов("TScannerWord", режим),
		режим:    режим,
		пулСтр:   мПулСтр.Нов(текстИсх, режим),
		текстИсх: текстИсх,
		poolWord: мПс.Нов(),
	}
	if _скан.коорд, ош = мКоорд.Нов(1, 0); ош != nil {
		return nil, fmt.Errorf("scannerword.go/Нов(): ERROR при создании координат сканера\n\t%v", ош)
	}
	if ош = _скан._Обработать(текстИсх); ош != nil {
		return nil, fmt.Errorf("scannerword.go/Нов(): ERROR при работе сканера\n\t%v", ош)
	}
	_скан.лог.Отладка("scaner.go/Нов", "Создание сканера")
	_скан.лог.Отладка("scaner.go/Нов", "Всего слов", _скан.poolWord.Len())
	return &_скан, nil
}

//Добавляет новое слово в список слов
func (sf *TScannerWord) _СловоДобав(пСлово мТип.UWord) (ош error) {
	коорд, ош := мКоорд.Нов(sf.коорд.Стр(), sf.коорд.Поз())
	if ош != nil {
		return fmt.Errorf("TScannerWord._СловоДобав(): ERROR при создании координат слова\n\t%v", ош)
	}
	строка, ош := sf.пулСтр.Строка(sf.коорд.Стр())
	if ош != nil {
		return fmt.Errorf("TScannerWord._СловоДобав(): ERROR при получении строки\n\t%v", ош)
	}
	слово, ош := мСлово.Нов(коорд, пСлово, строка)
	if ош != nil {
		return fmt.Errorf("TScannerWord._СловоДобав(): ERROR при создании слова(%v)\n\t%v", пСлово, ош)
	}
	sf.poolWord.Доб(слово)
	return nil
}

//_Обработать -- обрабатывает исходный текст на наличие слов, выделяет слова, расставляет словам
//   номера строк, позицию в строке. Тип не заполняет так как нужен семантический анализ (может это
//   тупо комментарии или строки).
func (sf *TScannerWord) _Обработать(текстИсх мТип.UTextSource) (ош error) {
	sf.лог.Отладка("_Обработать")
	словоНов := мТип.UWord("")
	_СловоНов := func() error {
		if словоНов != "" {
			if ош = sf._СловоДобав(словоНов); ош != nil {
				return fmt.Errorf("TScannerWord._Обработать()._СловоНов(): ERROR при добавлении слова %q\n\t%v", словоНов, ош)
			}
		}
		словоНов = ""
		return nil
	}
	//перебор всех литер в исходнике на предмет получить все слова
	for _, руна := range sf.текстИсх {
		//Контролировать формирование слов по разделителям слов
		switch _лит := мТип.UWord(руна); _лит {
		case "\n": //Увеличить номер строки, сбросить позицию
			_СловоНов()
			sf.коорд.СтрДоб()
			sf.коорд.ПозСброс()
		case "\t", " ", "\r": //Начало нового слова
			_СловоНов()
			sf.коорд.ПозДоб()
		default:
			словоНов += _лит
		}
	}
	return nil
}

// PoolWord -- возвращает хранимый пул слов
func (sf *TScannerWord) PoolWord() мТип.IPoolWord {
	return sf.poolWord
}
