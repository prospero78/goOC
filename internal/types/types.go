package types

/*
	Модуль предоставляет специальные и интерфейсные типы
	для работы компилятора.
*/

//UTextSource -- специальный строковый тип для хранения всего исходника
type UTextSource string

//UStringNum -- специальный целочисленный тип для хранения номера строки исходника
type UStringNum int

//UStringSource -- специальный строковый тип для хранения строки исходника
type UStringSource string

//UFileSourceName -- специальный строковый тип для хранения имени файла исходника
type UFileSourceName string

//UWordNum -- специальный строковый тип для хранения имени файла исходника
type UWordNum string

//UWord -- специальный строковый тип для хранения слова исходника
type UWord string

//UWordKey -- специальный строковый тип для хранения ключевых слов
type UWordKey string

//ISourceString -- интерфейс к исходной строке
type ISourceString interface {
	//String -- возвращат хранимую строку исходника
	String() UStringSource
	//Num -- возвращает номер исходной строки
	Num() UStringNum
}

//IPoolStringSource -- интерфейс для исходника разбитого построчно
type IPoolStringSource interface {
	//String -- возвращает строку исходника по номеру
	String(UStringNum) (ISourceString, error)
	//PoolSourceString -- возвращает пул всех строк исходника
	PoolSourceString() map[int]ISourceString
}

//ULit -- специальный строковый тип для хранения литеры исходника
type ULit string

//ULitType -- специальный целочисленный класс для хранения признака класса литеры
type ULitType int

//ILit -- интерфейс тип для литеры
type ILit interface {
	IsLetter() bool
	IsDigit() bool
	IsSpecLetter() bool
	Set(ULit) error
	Get() ULit
	String() string
}

//UStringPos -- специальный тип для значения позиции литеры в строке
type UStringPos int

//ICoordFix -- интерфейс к неизменяемой координате
type ICoordFix interface {
	//Num -- возвращает хранимый номер строки
	Num() UStringNum
	//Pos -- возвращает позицию в строке
	Pos() UStringPos
}

//IWord -- интерфейс для типа слова
type IWord interface {
	//IsName -- проверяет строго слово на допустимость имени (не ключевое слово)
	IsName() bool
	//Word -- возвращает хранимое слово
	Word() UWord
	//SourceString -- возвращает строку хранимого слова
	SourceString() UStringSource
}

//IPoolWord -- интерфейс для пула слов. Можно добавлять слова, получать, удалять
type IPoolWord interface {
	//Add -- добавляет новое слово в пул
	Add(IWord) error
	//Len -- возвращает размер словаря
	Len() int
}

//IScanner -- интерфейс сканера
type IScanner interface {
	//Process -- обрабатывает переданный исходник
	Process(UFileSourceName) error
	//PoolWord -- возвращает полученный пул слов
	PoolWord() map[UWordNum]IWord
}
