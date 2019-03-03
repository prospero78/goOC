package пакКоорд

import (
	пакТест "testing"
)

// Тест на сохранение правильных значений и создание ТуКоордФикс
func TestСтр1(тест *пакТест.T) {
	кф, ош := КоордФиксНов(1, 0)
	if ош != nil {
		тест.Errorf("Ошибка при создании *ТуКоордФикс\n %v", ош)
	}
	if кф.Стр() != 1 {
		тест.Errorf("Значение кф.Стр(1) не равно 1 (минимальное для номера строки)")
	}
}

// Тест на сопротивление 0-строке (должна быть минимум 1)
func TestСтр2(тест *пакТест.T) {
	кф, ош := КоордФиксНов(0, 0)
	if ош == nil {
		тест.Errorf("Неверное создание *ТуКоордФикс: строка не может иметь № 0")
	}
	if ош == nil && кф.Стр() == 0 { // При правильной работе прокатывать не должно
		тест.Errorf("Значение кф.Стр() не может быть равно 0")
	}
}

// Тест на сопротивление отрицательному номеру строки (должна быть минимум 1)
func TestСтр3(тест *пакТест.T) {
	кф, ош := КоордФиксНов(-1, 0)
	if ош == nil {
		тест.Errorf("Неверное создание *ТуКоордФикс: строка не может иметь № -1")
	}
	if ош == nil && кф.Стр() == -1 { // При правильной работе прокатывать не должно
		тест.Errorf("Значение кф.Стр() не может быть равно 0")
	}
}

// Тест на сохранении номера позиции (должен быть минимум 0)
func TestПоз1(тест *пакТест.T) {
	кф, ош := КоордФиксНов(1, 10)
	if ош != nil {
		тест.Errorf("Неверное создание *ТуКоордФикс: позиция в строке")
	}
	if кф.Поз.Знач() != 10 { // При правильной работе прокатывать не должно
		тест.Errorf("Значение кф.Поз()=10 не равно 10")
	}
}

// Тест на сопротивление отрицательного номера позиции (должен быть минимум 0)
func TestПоз2(тест *пакТест.T) {
	кф, ош := КоордФиксНов(1, -5)
	if ош == nil {
		тест.Errorf("Неверное создание *ТуКоордФикс: позиция в строке не может быть отрицательной")
	}
	if ош == nil && кф.Поз() == -5 { // При правильной работе прокатывать не должно
		тест.Errorf("Значение кф.Поз()=-5 как отрицательное недопустимо")
	}
}
