package пакСекциИмпорт

/*
	Модуль предоставляет тип секции импорта
*/
import (
	мКлюч "../../пакКлючи"
	мСлово "../../пакСлово"
	мСек "../пакСекция"
	мФмт "fmt"
)

//ИСекцияИмпорт -- интерфейс к секции импорта
type ИСекцияИмпорт interface {
	мСек.ИСекция
	//Обработать() error //Нужно ли это -- под сомнением
}

//тСекцияИмпорт -- тип выделяети хранит слова импорта
type тСекцияИмпорт struct {
	слова      map[мСек.UWordNum]мСлово.IWord
	словаВсе   map[мСек.UWordNum]мСлово.IWord
	словоКонец мСек.UWordNum //номер слова для ограничния секции импорта
	класс      мСек.СКласс
}

//СекцияИмпортНов -- Создаёт и возвращает новый экземпляр для выделения слов импорта модуля
func СекцияИмпортНов() (импорт ИСекцияИмпорт, ош error) {
	_импорт := &тСекцияИмпорт{}
	if _импорт == nil {
		return nil, мФмт.Errorf("СекцияИмпортНов(): нет памяти под секцию импорта?\n")
	}
	_импорт.слова = make(map[мСек.UWordNum]мСлово.IWord)
	return _импорт, ош
}

//Возвращает класс секции (МОДУЛЬ, ИМПОРТ, КОММЕНТ и т.д.)
func (сам *тСекцияИмпорт) Класс() мСек.СКласс {
	return сам.класс
}

// возвращает список слов модуля
func (сам *тСекцияИмпорт) Слова() map[мСек.UWordNum]мСлово.IWord {
	return сам.слова
}

//  устанавливает список слов секции
func (сам *тСекцияИмпорт) СловаУст(пСлова map[мСек.UWordNum]мСлово.IWord) error {
	if пСлова == nil {
		return мФмт.Errorf("тСекцияИмпорт.СловаУст(): пСлова не может быть nil\n")
	}
	сам.словаВсе = пСлова
	return nil
}

//главная функция обработки секции импорт
func (сам *тСекцияИмпорт) _Обработать() (ош error) {
	окИмп := false //Признак наличи импорта
	{              //Проверить есть ли импорт
		if окИмп, ош = сам._ЕслиИмпорт(); ош != nil {
			return мФмт.Errorf("тСекцияИмпорт.Обработать(): ошибка при поиске начала секции импорта\n\t%v", ош)
		}
	}
	//Разобрать секцию импорта
	if окИмп { //ИМПОРТ есть, разобрать секцию импорта
		//Проверить на окончание импорта
		if ош = сам._ЕслиИмпортОграничен(); ош != nil {
			return мФмт.Errorf("тСекцияИмпорт.Обработать(): ошибка при проверке ограничения секции импорта\n\t%v", ош)
		}
		if ош = сам._СловаСекцииРазделить(); ош != nil {
			return мФмт.Errorf("тСекцияИмпорт.Обработать(): ошибка при разделении слов секции ИМПОРТ и модуля\n\t%v", ош)
		}
		if ош := сам._ЕслиИмпортОдин(); ош != nil {
			return мФмт.Errorf("тСекцияИмпорт.Обработать(): ошибка при проверке одной секции импорт\n\t%v", ош)
		}
		//Теперь нужно разобрать все слова импорта по именам модулей и алиасам
		if ош = сам._ИмпортРазобрать(); ош != nil {
			return мФмт.Errorf("тСекцияИмпорт.Обработать(): ошибка при разборе секции импорта\n\t%v", ош)
		}
	}
	return nil
}

//Проверяет первое ключевое слово IMPORT
func (сам *тСекцияИмпорт) _ЕслиИмпорт() (рез bool, ош error) {
	//Проверить наличие слова IMPORT
	словоИМПОРТ := сам.СловаДругие()[0]
	окИмп, ош := мКлюч.Ключи.Проверить("IMPORT", мКлюч.СКлюч(словоИМПОРТ.Слово()))
	if ош != nil {
		return false, мФмт.Errorf("тСекцияИмпорт._ЕслиИмпорт(): ошибка пи поиске ключевого слова IMPORT\n\t%v", ош)
	}
	if !окИмп {
		return false, nil
	}
	//Слово IMPORT найдено
	слова, ош := мСек.СловаОбрезать(сам.СловаДругие())
	if ош != nil {
		return false, мФмт.Errorf("тСекцияИмпорт._ЕслиИмпорт(): ошибка при обрезке слов\n\t%v", ош)
	}
	сам.СловаУст(слова)
	return true, nil
}

//Ищет разделитель в секции импорта
func (сам *тСекцияИмпорт) _ЕслиИмпортОграничен() (ош error) {
	_КонецСекции := func(слово мСлово.IWord) bool {
		стр := слово.Слово()
		return стр == ";"
	}
	for счётСловоИмпорт := мСек.UWordNum(0); int(счётСловоИмпорт) < len(сам.СловаДругие()); счётСловоИмпорт++ {
		слово := сам.СловаДругие()[счётСловоИмпорт]
		if _КонецСекции(слово) {
			сам.словоКонец = счётСловоИмпорт
			return nil
		}
	}
	return мФмт.Errorf("ТуИмпорт._ЕслиИмпортОграничен(): секция импорта ничем не ограничена\n")
}

//По поледнему слову секции импорта вырезает слова для импорта
func (сам *тСекцияИмпорт) _СловаСекцииРазделить() (ош error) {
	словаИмпорт := make(map[мСек.UWordNum]мСлово.IWord)
	//Первое слово ИМПОРТ уже отброшено. Разделитель секции нужен -- подскажет, где конец
	for адр := мСек.UWordNum(0); адр <= сам.словоКонец; адр++ {
		словаИмпорт[мСек.UWordNum(адр)] = сам.СловаДругие()[адр]
		delete(сам.СловаДругие(), адр)
	}
	сам.слова = словаИмпорт
	//Отладочный вывод для понять, что получилось
	мФмт.Printf("тСекцияИмпорт._СловаСекцииРазделить(): всего слов секции ИМПОРТ=[%v]\n", len(сам.слова))
	//Теперь надо изменить нумерацию слов в секции модуля
	//Получим все номера секции МОДУЛЬ
	номера := make(map[int]мСек.UWordNum)
	ном := 0
	for адр := range сам.СловаДругие() {
		номера[ном] = адр
		ном++
	}
	//Теперь перенумеруем слова модуля
	словаМодуля := make(map[мСек.UWordNum]мСлово.IWord)
	for адр := 0; адр < len(номера); адр++ {
		словаМодуля[мСек.UWordNum(адр)] = сам.СловаДругие()[номера[(адр)]]
	}
	if ош = сам.СловаУст(словаМодуля); ош != nil {
		return мФмт.Errorf("тСекцияИмпорт._СловаСекцииРазделить(): ошибка при обрезке слов модуля в секции импорта\n\t%v", ош)
	}
	//Отладочный вывод для понять, что получилось
	мФмт.Printf("тСекцияИмпорт._СловаСекцииРазделить(): всего слов секции МОДУЛЬ=[%v]\n", len(сам.СловаДругие()))
	for адр := мСек.UWordNum(0); адр < 15; адр++ {
		мФмт.Printf("\tадр=[%v]\t[%v]\n", адр, сам.Слова()[адр].Слово())
	}
	return nil
}

func (сам *тСекцияИмпорт) _ЕслиИмпортОдин() (ош error) {
	счётИмпорт := 0
	словоНом := мСек.UWordNum(0)
	словаМодуляВсего := len(сам.СловаДругие())
	for словоНом < мСек.UWordNum(словаМодуляВсего) {
		слово := сам.СловаДругие()[словоНом]
		for индекс := range мСлово.КсИмпорт {
			стрИмпорт := слово.Слово()
			if стрИмпорт == мСлово.КсИмпорт[индекс] {
				счётИмпорт++
			}
			if счётИмпорт > 1 {
				стрИсх := слово.Строка()
				ош = мФмт.Errorf("тСекцияИмпорт._ЕслиИмпортОдин(): IMPORT два раза в одном модуле запрещён\n\t%v", стрИсх)
				return ош
			}
		}
		словоНом++
	}
	return nil
}

//Разбирает секцию импорта, хранит имена импортируемых модулей, их алиасов и номеров
func (сам *тСекцияИмпорт) _ИмпортРазобрать() (ош error) {
	_ЕслиПуть := func(пИмя мСлово.IWord) (ок bool) {
		//Возвращае ток, если переданное слово, либо ИмяСтрого, либо точка
		бДа := пИмя.IsName() || пИмя.Слово() == "."
		бНет := пИмя.Слово() == "," || пИмя.Слово() == ";" || пИмя.Слово() == ":="
		if бДа && !бНет {
			return true
		}
		return false
	}
	if len(сам.Слова()) < 2 {
		return мФмт.Errorf("тСекцияИмпорт._ИмпортРазобрать(): Объявленная секция импорта не может иметь меньше 2 слов\n\t%v", ош)
	}
	var (
		словоРаздел мСлово.IWord
		ок          bool
	)
	for адр := мСек.UWordNum(0); int(адр) < len(сам.Слова()); адр++ {
		/* Слово +1 может быть:
		- запятая (0 -- реальное имя модуля)
		- ";" ( окончание секции импорта)
		- ":=" (0 -- алиас, +1 -- реальное имя модуля)
		если не эти три, то это -- сбой.
		*/
		//1. Контроль, что не вышли за предела словаря импорта
		if int(адр) < len(сам.Слова())-1 {
			словоРаздел = сам.Слова()[адр+1]

		}
		мФмт.Printf("Слово: %v.[%v]\n", адр, словоРаздел.Слово())
		{ //2. Проверим, а не конец ли это секции
			if ок, ош = мКлюч.Ключи.Проверить(";", мКлюч.СКлюч(словоРаздел.Слово())); ош != nil {
				return мФмт.Errorf("тСекцияИмпорт._ИмпортРазобрать(): ошибка при проверке разделителя `;`\n\t%v", ош)
			}
			if ок { //Достигнут конец секции
				return nil
			}
		}
		{ //3. Проверим, а не запятая ли это
			if ок, ош = мКлюч.Ключи.Проверить(",", мКлюч.СКлюч(словоРаздел.Слово())); ош != nil {
				return мФмт.Errorf("тСекцияИмпорт._ИмпортРазобрать(): ошибка при проверке разделителя `,`\n\t%v", ош)
			}
			if ок { //Предыдущее слово -- имя модуля
				словоМодуль := сам.Слова()[адр]
				if ош = сам._МодульДоб(словоМодуль); ош != nil {
					return мФмт.Errorf("тСекцияИмпорт._ИмпортРазобрать(): ошибка при добавлении модуля [%v]\n\t%v", словоМодуль.Слово(), ош)
				}
				//3.1 Теперь изменим адрес на следующий разделитель
				адр++
				continue
			}
		}
		{ //4. Проверим, а не разделитель алиасов ли это
			if ок, ош = мКлюч.Ключи.Проверить(":=", мКлюч.СКлюч(словоРаздел.Слово())); ош != nil {
				return мФмт.Errorf("тСекцияИмпорт._ИмпортРазобрать(): ошибка при проверке разделителя `:=`\n\t%v", ош)
			}
			if ок { //Предыдущее слово -- алиас модуля
				алиасМодуль := сам.Слова()[адр]
				//4.1 Соберём имя модуля с его пакетом из частей (если оно такое)
				адр += 2
				путьМодуль := сам.Слова()[адр]
				стрПуть := мСлово.UWord("")
				for _ЕслиПуть(путьМодуль) {
					стрПуть += путьМодуль.Слово()
					адр++
					путьМодуль = сам.Слова()[адр]
				}
				//4.2 Убедиться, что последняя литера не точка (а такое может быть)
				if путьМодуль.Слово() == "." {
					return мФмт.Errorf("тСекцияИмпорт._ИмпортРазобрать(): ошибка при проверке имени пути модуля, путь=[%v]\n", стрПуть)
				}
				if ош = сам._АлиасДоб(алиасМодуль, стрПуть); ош != nil {
					return мФмт.Errorf("тСекцияИмпорт._ИмпортРазобрать(): ошибка при добавлении алиаса и пути модуля [%v] [%v]\n\t%v", алиасМодуль.Слово(), стрПуть, ош)
				}
				//4.3 Поскольку счётчик адреса уже прибавлен -- его не меняем. Либо ",", либо ";"
			}
		}

	}
	return nil
}

func (сам *тСекцияИмпорт) СловаДругие() map[мСек.UWordNum]мСлово.IWord {
	return сам.словаВсе
}

//Добавляет имя модуля в общий список
func (сам *тСекцияИмпорт) _МодульДоб(пСлово мСлово.IWord) (ош error) {
	/*
		if ош := мСмод.Модули.МодульДоб(пСлово, ""); ош != nil {
			return мФмт.Errorf("тСекцияИмпорт._МодульДоб(): ошибка при добавлении сущности модуля [%v]\n\t%v", пСлово.Слово(), ош)
		}
	*/
	return nil
}

//Добавляет имя модуля и его алиас в общий список
func (сам *тСекцияИмпорт) _АлиасДоб(пСлово мСлово.IWord, путьМодуль мСлово.UWord) (ош error) {
	/*
		if ош := мСмод.Модули.МодульДоб(пСлово, путьМодуль); ош != nil {
			return мФмт.Errorf("тСекцияИмпорт._МодульДоб(): ошибка при добавлении сущности алиаса и пути модуля [%v] [%v]\n\t%v", пСлово.Слово(), путьМодуль, ош)
		}
	*/
	return nil
}
