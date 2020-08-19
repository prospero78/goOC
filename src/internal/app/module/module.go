package module

/*
	Модуль предоставляет тип для компиляции отдельного Оберон-модуля с исходным кодом
*/

import (
	"fmt"
	мСкан "oc/internal/app/module/scaner"
	мФис "oc/internal/app/module/sourcefile"
	мЛог "oc/internal/log"
)

//ТМодуль -- операции с Оберон-модулем
type ТМодуль struct {
	файлИмя string //Имя модуля для компиляции
	лог     *мЛог.ТЛог
	файлИсх *мФис.ТФайлИсх //Операции с файлом исходника
	сканер  *мСкан.ТСканер
	код     string //Полученный код после компиляции
}

//Нов -- возвращает указатель на новый ТМодуль
func Нов(файлИмя string, режим int) (мод *ТМодуль, ош error) {
	_мод := ТМодуль{
		файлИмя: файлИмя,
		лог:     мЛог.Нов("ТМодуль", режим),
		сканер:  мСкан.Нов(режим),
	}
	if _мод.файлИсх, ош = мФис.Нов(файлИмя, режим); ош != nil {
		return nil, fmt.Errorf("modul.go/Нов(): ошибка при чтении файла модуля %q\n\t%v", ош)
	}
	_мод.лог.Отладка("Нов", "Создание модуля", файлИмя)
	if ош = _мод._Обработать(); ош != nil {
		return nil, fmt.Errorf("modul.go/Нов(): ошибка при обработке модуля %q", файлИмя)
	}
	return &_мод, nil
}

//Запускает непосредственную обработку модуля
func (сам *ТМодуль) _Обработать() (ош error) {
	сам.лог.Отладка("_Обработать", сам.файлИмя)
	if ош = сам.сканер.Обработать(сам.файлИмя); ош != nil {
		return fmt.Errorf("ТМодуль._Обработать(): ошибка при работе сканера\n\t%v", ош)
	}
	return nil
}

//КодПолуч -- возвращает скомпилированный код
func (сам *ТМодуль) КодПолуч() string {
	return сам.код
}
