package types

/*
	Модуль предоставляет специальные и интерфейсные типы
	для работы компилятора.
*/

//СТекстИсх -- специальный строковый тип для хранения всего исходника
type СТекстИсх string

//ССтрокаНом -- специальный целочисленный тип для хранения номера строки исходника
type ССтрокаНом int

//ССтрокаИсх -- специальный строковый тип для хранения строки исходника
type ССтрокаИсх string

//СФайлИсхИмя -- специальный строковый тип для хранения имени файла исходника
type СФайлИсхИмя string

//ССловоНомер -- специальный строковый тип для хранения имени файла исходника
type ССловоНомер string

//ССлово -- специальный строковый тип для хранения слова исходника
type ССлово string

//ССловоКлюч -- специальный строковый тип для хранения ключевых слов
type ССловоКлюч string

//ИСтрокаИсх -- интерфейс к исходной строке
type ИСтрокаИсх interface {
	//Строка -- возвращат хранимую строку исходника
	//Номер -- возвращает номер исходной строки
	Строка() ССтрокаИсх
	Номер() ССтрокаНом
}

//ИПулИсхСтр -- интерфейс для исходника разбитого построчно
type ИПулИсхСтр interface {
	//Строка -- возвращает строку исходника по номеру
	Строка(ССтрокаНом) (ИСтрокаИсх, error)
	//СтрокиВсе -- возвращает пул всех строк исходника
	СтрокиВсе() map[int]ИСтрокаИсх
}

//СЛит -- специальный строковый тип для хранения литеры исходника
type СЛит string

//СЛитКласс -- специальный целочисленный класс для хранения признака класса литеры
type СЛитКласс int

//ИЛит -- интерфейс тип для литеры
type ИЛит interface {
	ЕслиБуква() bool
	ЕслиЦифра() bool
	ЕслиСпецЛит() bool
	Уст(СЛит) error
	Получ() СЛит
	String() string
}

//ССтрокаПоз -- специальный тип для значения позиции литеры в строке
type ССтрокаПоз int

//ИКоордФикс -- интерфейс к неизменяемой координате
type ИКоордФикс interface {
	//Стр - -возвращает хранимый номер строки
	Стр() ССтрокаНом
	//Поз -- возвращает позицию в строке
	Поз() ССтрокаПоз
}

//ИСлово -- интерфейс для типа слова
type ИСлово interface {
	//ЕслиИмяСтрого -- проверяет сторго слово на допустимост ьимени (не ключевое слово)
	ЕслиИмяСтрого() bool
	//Слово -- возвращает хранимое слово
	Слово() ССлово
	//Строка -- возвращает строку хранимого слова
	Строка() ССтрокаИсх
}

//ИПулСлова -- интерфейс для пула слов. Можно добавлять слова, получать, удалять
type ИПулСлова interface {
	//Доб -- добавляет новое слово в пул
	Доб(ИСлово) error
	//Len -- возвращает размер словаря
	Len() int
}

//ИСканер -- интерфейс сканера
type ИСканер interface {
	Обработать(СФайлИсхИмя) error
	Слова() map[ССловоНомер]ИСлово
}
