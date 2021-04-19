package types

/*
	Файл предоставляет различные виды позиции в строке
*/

// IPosFix -- фиксированная позиция в строке
type IPosFix interface {
	// Get -- возвращает позицию в строке
	Get() APos
}

// IPos -- изменяемая позиция в строке
type IPos interface {
	IPosFix
	// Set -- устанавливает позицию литеру в строке
	Set(APos) error
}
