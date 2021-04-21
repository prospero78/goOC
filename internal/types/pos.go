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

// INumStrFix -- фиксированный номер строки исходника
type INumStrFix interface {
	// Get -- возвращает хранимый номер строки
	Get() ANumStr
}

// INumStr -- номер строки исходника
type INumStr interface {
	INumStrFix
	// Set -- устанавливает номер строки
	Set(ANumStr) error
}

// ICoordFix -- фиксированные координаты в исходнике
type ICoordFix interface {
	// Pos -- возвращает позицию литеры в строке
	Pos() APos
	// NumStr -- возвращает номер строки исходника
	NumStr() ANumStr
}

// ICoord -- изменяемые координаты в исходнике
type ICoord interface {
	ICoordFix
	// Set -- устанавливает координаты
	Set(APos, ANumStr) error
	// SetPos -- устанавливает позицию литеры
	SetPos(APos) error
	// SetNumStr -- устанавливает номер строки
	SetNumStr(ANumStr) error
}
