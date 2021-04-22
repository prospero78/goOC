package types

// IScanner -- интерфейс к сканеру
type IScanner interface {
	// ListWord -- возвращает список слов
	ListWord() []IWord
	// Scan -- сканирует исходник, разбивает на необходимые структуры
	Scan(nameMod AModule, strSource string)
}
