package пакИнтерфейсы

/*
	Модуль предоставляет интерфейс типа литеры
*/

//СЛитНомер -- специальный целочисленный тип для хранения номера литеры
type СЛитНомер int

//СЛит -- специальный строковый тип для хранения литеры исходного текста
type СЛит string

//ИЛит -- интерфейс тип для литеры
type ИЛит interface {
	ЕслиБуква() bool
	ЕслиЦифра() bool
	ЕслиЗнаки() bool
	Уст(СЛит) error
	Лит() СЛит
}
