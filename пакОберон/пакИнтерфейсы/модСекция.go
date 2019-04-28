package пакИнтерфейсы

/*
	Модуль предоставляет интерфейсы для секций
*/

//ССекцияИмя -- специальный строковый тип для хранения имени секции
type ССекцияИмя string

//ССловоНомерСекция -- специадбный целочисленный тип для хранения номера слова секции
type ССловоНомерСекция int

//ССловоНомерМодуль -- специальный целочисленный тип для хранения номера слова модуля
type ССловоНомерМодуль int

//ИСекция -- базовый тип секции
type ИСекция interface {
	СловаМодуля() map[ССловоНомерМодуль]ИСлово
	//СловаСекции() map[ССловоНомерСекция]ИСлово // У каждого типа секции свой метод
	СловаУст(map[ССловоНомерМодуль]ИСлово) error
	Имя() ССекцияИмя
}

//ССловоНомерКоммент -- специальный целочисленный тип для хранения номера слова секции комментариев
type ССловоНомерКоммент int

//ИСекцияКоммент -- интерфейс к псевдосекции комментариев
type ИСекцияКоммент interface {
	ИСекция
	СловаКоммент() map[ССловоНомерКоммент]ИСлово
	Обработать() error //Нужно ли это -- под сомнением
}

//ССловоНомерИмпорт -- специальный целочисленный тип для хранения номера слова секции импорта
type ССловоНомерИмпорт int

//ИСекцияИмпорт -- интерфейс к секции импорта
type ИСекцияИмпорт interface {
	ИСекция
	СловаИмпорт() map[ССловоНомерИмпорт]ИСлово
	Обработать() error //Нужно ли это -- под сомнением
}
