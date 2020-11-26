package litera

/*
	Модуль предоставляет потокобезопасный тип для работы с отдельной литерой.
*/

import (
	"fmt"
	мТип "oc/internal/types"
	"strings"
)

//ТЛит -- тип для работы с отдельной литерой
type ТЛит struct {
	лит   мТип.ULit
	класс мТип.ULitType //Хранит класс литеры
}

const (
	//наборы букв для перебора
	стрБуквыРус = "абвгдеёжзийклмнопрстуфхцчшщьыъэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЬЫЪЭЮЯ"
	стрБуквыАнг = "abcdefghjiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	стрБуквыВсе = стрБуквыРус + стрБуквыАнг

	//CтрЦифры -- список цифр, что могут встречаться в числах
	стрЦифры = "0123456789."

	//Различные специальные литеры, не имеющие смысла в коде
	стрСпецЛит = "`~!@#№$%^&?\\_"
	//Литеры-разделители
	стрРазделЛит = "\n\t "
	//Литеры-скобки
	стрЛитСкобки = "(){}[]"
	//Литеры-операции
	стрЛитОпер               = "*/+-"
	//LETTER -- буквы
	LETTER     мТип.ULitType = iota + 1
	//SPECLETTER -- специальные литеры
	SPECLETTER
	//DIGIT -- цифры
	DIGIT
	//TERMINAL -- разделительные литеры
	TERMINAL
	//BRACKET -- скобки всех видов
	BRACKET
	//OPERATION -- литеры операций
	OPERATION
)

//Нов -- возвращает ссылку на новый ТЛит
func Нов(пЛит мТип.ULit) (лит *ТЛит, ош error) {
	_лит := ТЛит{}
	if ош = _лит.Уст(пЛит); ош != nil {
		return nil, fmt.Errorf("litera.go/Нов(): ERROR при установке литеры(%v)\n\t%v", пЛит, ош)
	}
	return &_лит, nil
}

// IsLetter -- проверяет наличие буквы в литере
func (сам *ТЛит) IsLetter() bool {
	return сам.класс == LETTER
}

// IsSpecLetter -- проверяет наличие специальных литер не имеющих смысла в коде
func (сам *ТЛит) IsSpecLetter() bool {
	return сам.класс == SPECLETTER
}

//IsDigit -- проверяет, что литера цифра
func (сам *ТЛит) IsDigit() bool {
	return сам.класс == DIGIT
}

//ЕслиРазделит -- проверяет, что литера разделитель
func (сам *ТЛит) ЕслиРазделит() bool {
	return сам.класс == TERMINAL
}

//ЕслиСкобки -- проверяет, что литера скобка
func (сам *ТЛит) ЕслиСкобки() bool {
	return сам.класс == BRACKET
}

//ЕслиОпер -- проверяет, что литера операция
func (сам *ТЛит) ЕслиОпер() bool {
	return сам.класс == OPERATION
}

//Уст -- устанавливает хранимую литеру
func (сам *ТЛит) Уст(пЛит мТип.ULit) error {
	if пЛит == "" {
		return fmt.Errorf("ТЛит.Уст(): пЛит не может быть пустой")
	}
	switch{
	case strings.Contains(стрБуквыВсе, string(пЛит)):
		сам.класс = LETTER
	case strings.Contains(стрСпецЛит, string(пЛит)):
		сам.класс = SPECLETTER
	case strings.Contains(стрЦифры, string(пЛит)):
		сам.класс = DIGIT
	case strings.Contains(стрРазделЛит, string(пЛит)):
		сам.класс = TERMINAL
	case strings.Contains(стрЛитСкобки, string(пЛит)):
		сам.класс = BRACKET
	case strings.Contains(стрЛитОпер, string(пЛит)):
		сам.класс = OPERATION
	default:
		return fmt.Errorf("ТЛит.Уст(): ERROR неизвестный класс литеры, пЛит=[%v]", пЛит)
	}
	сам.лит = пЛит
	return nil
}

//Получ -- возвращает хранимую литеру
func (сам *ТЛит) Получ() мТип.ULit {
	return сам.лит
}

//Класс -- возвращает класс литеры
func (сам *ТЛит) Класс() мТип.ULitType {
	return сам.класс
}

func (сам *ТЛит) String() string {
	return string(сам.лит)
}
