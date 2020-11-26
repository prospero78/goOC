package scannerword

/*
	Модуль предоставляет сканер слов в исходном модуле
*/

import (
	"fmt"
	"oc/internal/app/module/scannerword/coord"
	"oc/internal/app/module/scannerword/poolword/word"
	"oc/internal/log"

	//мЛит "oc/internal/app/module/scannerword/litera"
	"oc/internal/types"
)

//TScannerWord -- операции сканера с исходным текстом Оберон-файла
type TScannerWord struct {
	log      *log.TLog
	режим    int
	пулСтр   types.IPoolStringSource //Пул исходных строк для выведения координат
	poolWord types.IPoolWord         //Пул для хранения слов исходника
	текстИсх types.UTextSource       //Исходный текст для анализа
	коорд    *coord.TCoord           //Текущие координаты в тексте
}

//Нов -- возвращает указатель на новый TScannerWord
func Нов(текстИсх types.UTextSource, режим int) (скан *TScannerWord, ош error) {
	_скан := TScannerWord{
		log:      log.New("TScannerWord", режим),
		режим:    режим,
		пулСтр:   poolsourcestr.New(текстИсх, режим),
		текстИсх: текстИсх,
		poolWord: poolword.New(),
	}
	if _скан.коорд, ош = coord.New(1, 0); ош != nil {
		return nil, fmt.Errorf("scannerword.go/Нов(): ERROR при создании координат сканера\n\t%v", ош)
	}
	if ош = _скан._Обработать(текстИсх); ош != nil {
		return nil, fmt.Errorf("scannerword.go/Нов(): ERROR при работе сканера\n\t%v", ош)
	}
	_скан.log.Debugf("scaner.go/Нов", "Создание сканера")
	_скан.log.Debugf("scaner.go/Нов", "Всего слов", _скан.poolWord.Len())
	return &_скан, nil
}

//Добавляет новое слово в список слов
func (sf *TScannerWord) _СловоДобав(пСлово types.UWord) (ош error) {
	коорд, ош := coord.New(sf.коорд.Num(), sf.коорд.Pos())
	if ош != nil {
		return fmt.Errorf("TScannerWord._СловоДобав(): ERROR при создании координат слова\n\t%v", ош)
	}
	строка, ош := sf.пулСтр.String(sf.коорд.Num())
	if ош != nil {
		return fmt.Errorf("TScannerWord._СловоДобав(): ERROR при получении строки\n\t%v", ош)
	}
	слово, ош := word.Нов(коорд, пСлово, строка)
	if ош != nil {
		return fmt.Errorf("TScannerWord._СловоДобав(): ERROR при создании слова(%v)\n\t%v", пСлово, ош)
	}
	sf.poolWord.Add(слово)
	return nil
}

//_Обработать -- обрабатывает исходный текст на наличие слов, выделяет слова, расставляет словам
//   номера строк, позицию в строке. Тип не заполняет так как нужен семантический анализ (может это
//   тупо комментарии или строки).
func (sf *TScannerWord) _Обработать(текстИсх types.UTextSource) (ош error) {
	sf.log.Debugf("_Обработать")
	словоНов := types.UWord("")
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
		switch _лит := types.UWord(руна); _лит {
		case "\n": //Увеличить номер строки, сбросить позицию
			_СловоНов()
			sf.коорд.NumInc()
			sf.коорд.PosReset()
		case "\t", " ", "\r": //Начало нового слова
			_СловоНов()
			sf.коорд.PosInc()
		default:
			словоНов += _лит
		}
	}
	return nil
}

// PoolWord -- возвращает хранимый пул слов
func (sf *TScannerWord) PoolWord() types.IPoolWord {
	return sf.poolWord
}
