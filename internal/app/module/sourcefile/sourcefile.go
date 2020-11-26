package sourcefile

/*
	Модуль предоставляет тип для чтения файла и записи его же.
*/

import (
	"fmt"
	"io/ioutil"
	"oc/internal/log"
)

//TFileSource -- тип для работы с исходным файлом
type TFileSource struct {
	size int //Размер исходника в рунах
	text []rune
	mode int
	log  *log.TLog
}

// New -- возвращает указатель на новый ТФайлИсх
func New(fileName string, mode int) (filesource *ТФайлИсх, err error) {
	filesource := &TFileSource{
		log:  log.Нов("TFileSource", mode),
		mode: mode,
	}
	filesource.log.Debugf("New()")
	if err = _файл.read(fileName); err != nil {
		return nil, fmt.Errorf("sourcefile.go/New(): ОШИБКА чтения файла %q\n\t%v", fileName, err)
	}
	return filesource, nil
}

// readFile -- читает исходный файл
func (sf *TFileSource) readFile(файлИмя string) (ош error) {
	sf.log.Отладка("readFile", файлИмя)
	if файлИмя == "" {
		файлИмя = "./Hello.o7"
	}
	байты, ош := мВв.ReadFile(файлИмя)
	if ош != nil {
		return fmt.Errorf("TFileSource.readFile(): ОШИБКА при попытке прочитать файл")
	}

	// Строковое представление байтов
	sf.text = []rune(string(байты))

	sf.size = len([]rune(sf.text))
	//sf.лог.Отладка("read", fmt.Sprintf("Текст:\n%v\nДлина: %v\n", sf.текст, sf.размер))
	return nil
}

// PosLit -- Возвращает литеру по номеру руны
func (sf *TFileSource) PosLit(пПоз int) (lit string, err error) {
	if пПоз < 0 {
		return "", fmt.Errorf("TFileSource.PosLit(): указатель литеры пПоз не может быть < 0")
	}
	if пПоз > sf.size-1 {
		return "", fmt.Errorf("TFileSource.PosLit(): указатель литеры пПоз больше последней литеры, пПоз=%v, размер=[%v]", пПоз, sf.size)
	}
	lit = string(sf.text[пПоз])
	return lit, nil
}

// Source -- возвращает полностью исходный текст в отдельном срезе рун
func (sf *TFileSource) Source() (текст []rune) {
	return sf.text
}

// Size -- возвращает размер исходника в рунах
func (sf *TFileSource) Size() int {
	return sf.size
}
