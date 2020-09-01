package sourcefile

/*
	Модуль предоставляет тип для чтения файла и записи его же.
*/

import (
	"fmt"
	мВв "io/ioutil"
	мЛог "oc/internal/log"
)

//ТФайлИсх -- тип для работы с исходным файлом
type ТФайлИсх struct {
	размер int //Размер исходника в рунах
	текст  []rune
	режим  int
	лог    *мЛог.ТЛог
}

//Нов -- возвращает указатель на новый ТФайлИсх
func Нов(файлИмя string, режим int) (файл *ТФайлИсх, ош error) {
	_файл := ТФайлИсх{
		лог:   мЛог.Нов("ТФайлИсх", режим),
		режим: режим,
	}
	_файл.лог.Отладка("Нов()")
	if ош = _файл._Считать(файлИмя); ош != nil {
		return nil, fmt.Errorf("sourcefile.go/Нов(): ОШИБКА чтения файла %q\n\t%v", файлИмя, ош)
	}
	return &_файл, nil
}

//_Считать -- читает исходный файл
func (сам *ТФайлИсх) _Считать(файлИмя string) (ош error) {
	сам.лог.Отладка("_Считать", файлИмя)
	if файлИмя == "" {
		файлИмя = "./Hello.o7"
	}
	байты, ош := мВв.ReadFile(файлИмя)
	if ош != nil {
		return fmt.Errorf("ТФайлИсх.Считать(): ОШИБКА при попытке прочитать файл\n")
	}

	// Строковое представление байтов
	сам.текст = []rune(string(байты))

	сам.размер = len([]rune(сам.текст))
	//сам.лог.Отладка("_Считать", fmt.Sprintf("Текст:\n%v\nДлина: %v\n", сам.текст, сам.размер))
	return nil
}

//Лит -- Возвращает литеру по номеру руны
func (сам *ТФайлИсх) Лит(пПоз int) (лит string, ош error) {
	if пПоз < 0 {
		return "", fmt.Errorf("ТФайлИсх.Лит(): указатель литеры пПоз не может быть < 0\n")
	}
	if пПоз > сам.размер-1 {
		return "", fmt.Errorf("ТФайлИсх.Лит(): указатель литеры пПоз больше последней литеры, пПоз=%v, размер=[%v]\n", пПоз, сам.размер)
	}
	лит = string(сам.текст[пПоз])
	return лит, nil
}

//Исходник -- возвращает полностью исходный текст в отдельном срезе рун
func (сам *ТФайлИсх) Исходник() (текст []rune) {
	return сам.текст
}

//Размер -- возвращает размер исходника в рунах
func (сам *ТФайлИсх) Размер() int {
	return сам.размер
}
