package module

/*
	Модуль предоставляет тип для компиляции отдельного Оберон-модуля с исходным кодом
*/

import (
	"fmt"
	mScan "oc/internal/app/module/scannerword"
	мФис "oc/internal/app/module/sourcefile"
	"oc/internal/log"
	мТип "oc/internal/types"
)

//ТМодуль -- операции с Оберон-модулем
type ТМодуль struct {
	файлИмя string //Имя модуля для компиляции
	log     *log.TLog
	файлИсх *мФис.TFileSource //Операции с файлом исходника
	сканер  *mScan.TScannerWord
	код     string //Полученный код после компиляции
	режим   int    //Режим логирования
}

//Нов -- возвращает указатель на новый ТМодуль
func Нов(файлИмя string, режим int) (мод *ТМодуль, ош error) {
	_мод := ТМодуль{
		файлИмя: файлИмя,
		log:     log.New("ТМодуль", режим),
		режим:   режим,
	}
	if _мод.файлИсх, ош = мФис.New(файлИмя, режим); ош != nil {
		return nil, fmt.Errorf("modul.go/Нов(): ERROR при чтении файла модуля %q\n\t%v", файлИмя, ош)
	}
	_мод.log.Debugf("Нов", "Создание модуля ", файлИмя)
	if ош = _мод._Обработать(); ош != nil {
		return nil, fmt.Errorf("modul.go/Нов(): ERROR при обработке модуля %q\n\t%v", файлИмя, ош)
	}
	return &_мод, nil
}

//Запускает непосредственную обработку модуля
func (сам *ТМодуль) _Обработать() (ош error) {
	сам.log.Debugf("_Обработать", сам.файлИмя)
	руны := сам.файлИсх.Source()
	if сам.сканер, ош = mScan.Нов(мТип.UTextSource(руны), сам.режим); ош != nil {
		return fmt.Errorf("ТМодуль._Обработать(): ERROR при создании сканера\n\t%v", ош)
	}
	//Теперь надо немножечко разбить по секциям
	if ош = сам._РазбитьНаСекции(); ош != nil {
		return fmt.Errorf("ТМодуль._Обработать(): ERROR при разбитии на секции\n\t%v", ош)
	}
	return nil
}

//Разбивает все слова на секции (type, var, const, procedure, module, comment)
func (сам *ТМодуль) _РазбитьНаСекции() (ош error) {
	//Выбрать все слова секции comment
	пулСлова := сам.сканер.PoolWord()
	пулСлова.Add(nil)
	panic("доделать")
}

//КодПолуч -- возвращает скомпилированный код
func (сам *ТМодуль) КодПолуч() string {
	return сам.код
}
