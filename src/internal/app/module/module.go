package module

/*
	Модуль предоставляет тип для компиляции отдельного Оберон-модуля с исходным кодом
*/

import (
	"fmt"
	мСкан "oc/internal/app/module/scannerword"
	мФис "oc/internal/app/module/sourcefile"
	мЛог "oc/internal/log"
	мТип "oc/internal/types"
)

//ТМодуль -- операции с Оберон-модулем
type ТМодуль struct {
	файлИмя string //Имя модуля для компиляции
	лог     *мЛог.ТЛог
	файлИсх *мФис.ТФайлИсх //Операции с файлом исходника
	сканер  *мСкан.ТСканерСлов
	код     string //Полученный код после компиляции
	режим   int    //Режим логирования
}

//Нов -- возвращает указатель на новый ТМодуль
func Нов(файлИмя string, режим int) (мод *ТМодуль, ош error) {
	_мод := ТМодуль{
		файлИмя: файлИмя,
		лог:     мЛог.Нов("ТМодуль", режим),
		режим:   режим,
	}
	if _мод.файлИсх, ош = мФис.Нов(файлИмя, режим); ош != nil {
		return nil, fmt.Errorf("modul.go/Нов(): ОШИБКА при чтении файла модуля %q\n\t%v", ош)
	}
	_мод.лог.Отладка("Нов", "Создание модуля", файлИмя)
	if ош = _мод._Обработать(); ош != nil {
		return nil, fmt.Errorf("modul.go/Нов(): ОШИБКА при обработке модуля %q\n\t%v", файлИмя, ош)
	}
	return &_мод, nil
}

//Запускает непосредственную обработку модуля
func (сам *ТМодуль) _Обработать() (ош error) {
	сам.лог.Отладка("_Обработать", сам.файлИмя)
	руны := сам.файлИсх.Исходник()
	if сам.сканер, ош = мСкан.Нов(мТип.СТекстИсх(руны), сам.режим); ош != nil {
		return fmt.Errorf("ТМодуль._Обработать(): ОШИБКА при создании сканера\n\t%v", ош)
	}
	return nil
}

//КодПолуч -- возвращает скомпилированный код
func (сам *ТМодуль) КодПолуч() string {
	return сам.код
}
