package types

/*
	Файл предоставляет интерфейс к словам.
*/

// IStrWord -- интерфейс к строковому представлению слова
type IStrWord interface {
	// Get -- возвращает хранимое значение слова
	Get() AWord
	// Len -- возвращает длину слова
	Len() int
}

// IWord -- интерфейс слова
type IWord interface {
	// SetModule -- устанавливает модуль, к коорому принадлежит слово
	SetModule(*AModule) error
	// Word -- возвращает хранимое слово
	Word() AWord
	// IsName -- проверяет слово на строгое соответствие требованиям к имени
	IsName() bool
	// IsInt -- проверяет, что слово является целым числом
	IsInt() bool
	// IsReal -- проверяет, что слово является вещественным числом
	IsReal() bool
	// IsString -- проверяет, что слово является строкой (должно быть в кавычках)
	IsString() bool
	// IsBool -- проверяет, что слово является булевым числом
	IsBool() bool
	// IsCompoundName -- проверяет, что имя является составным
	IsCompoundName() bool
	// NumStr -- возвращает номер строки
	NumStr() ANumStr
	// GetType -- возвращает хранимое значение типа
	GetType() string
	// SetType -- устанавливает значение типа слова
	SetType(strType string)
	// Module -- возвращает хранимое имя модуля
	Module() *AModule
	// Pos -- возвращает позицию в строке
	Pos() APos
}
