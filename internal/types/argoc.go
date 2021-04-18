package types

// IArgOc -- интерфейс к параметрам запуска компилятора
type IArgOc interface {
	// FileName -- возвращает имя комплируемого файла
	FileName() string
}
