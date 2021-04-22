// Package isletter -- проверка на литеру
package isletter

import (
	"strings"
)

const (
	litEng   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	litRu    = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	litDigit = "1234567890"
	litOther = "_.@!"
)

// Check -- проверяет, что литера -- буква
func Check(lit string) bool {
	if strings.Contains(litEng, lit) { // en_En.UTF-8
		return true
	}
	if strings.Contains(litRu, lit) { // ru_RU.UTF-8
		return true
	}
	if strings.Contains(litOther, lit) { // Если допустимые литеры
		return true
	}
	if strings.Contains(litDigit, lit) { // Если цифры
		return true
	}
	return false
}
