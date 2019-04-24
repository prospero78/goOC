package пакАбстракция

/*
	Модуль предоставляет абсракцию типа литеры
*/

//СЛитНомер -- специальный целочисленный тип для хранения номера литеры
type СЛитНомер int

//СЛит -- специальный строковый тип для хранения литеры исходного текста
type СЛит string

//АЛит -- абстрактный тип для литеры
type АЛит interface {
	ЕслиБуква() bool
	ЕслиЦифра() bool
	Уст(СЛит) error
	Лит() СЛит
}
