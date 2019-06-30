package пакСекция

/*
	Модуль предоставляет тип для анализа секции импорта
*/

import (
	мФмт "fmt"
	мИнт "github.com/prospero78/goOC/пакОберон/пакИнтерфейсы"
	мКлюч "github.com/prospero78/goOC/пакОберон/пакКлючи"
	мКонс "github.com/prospero78/goOC/пакОберон/пакКонсоль"
	мСмод "github.com/prospero78/goOC/пакОберон/пакМодули"
	мСлово "github.com/prospero78/goOC/пакОберон/пакСлово"
)

//ТСекцияИмпорт -- тип выделяети хранит слова импорта
type ТСекцияИмпорт struct {
	*ТСекция
	словаИмпорт map[мИнт.ССловоНомерИмпорт]мИнт.ИСлово
	словоКонец  мИнт.ССловоНомерМодуль //номер слова для ограничния секции импорта
}

//СекцияИмпортНов -- Создаёт и возвращает новый экземпляр для выделения слов импорта модуля
func СекцияИмпортНов() (импорт *ТСекцияИмпорт, ош error) {
	мКонс.Конс.Отладить("СекцияИмпортНов()")
	секция, ош := СекцияНов("ИМПОРТ")
	if ош != nil {
		return nil, мФмт.Errorf("СекцияИмпортНов(): ошибка при создании секции импорта\n\t%v", ош)
	}
	импорт = &ТСекцияИмпорт{
		ТСекция: секция,
	}
	импорт.словаИмпорт = make(map[мИнт.ССловоНомерИмпорт]мИнт.ИСлово)
	return импорт, ош
}

//Обработать -- главная функция обработки секции импорт
func (сам *ТСекцияИмпорт) Обработать() (ош error) {
	мКонс.Конс.Отладить("ТСекцияИмпорт.Обработать()")
	окИмп := false //Признак наличи импорта
	{              //Проверить есть ли импорт
		if окИмп, ош = сам._ЕслиИмпорт(); ош != nil {
			return мФмт.Errorf("ТСекцияИмпорт.Обработать(): ошибка при поиске начала секции импорта\n\t%v", ош)
		}
	}
	//Разобрать секцию импорта
	if окИмп { //ИМПОРТ есть, разобрать секцию импорта
		//Проверить на окончание импорта
		if ош = сам._ЕслиИмпортОграничен(); ош != nil {
			return мФмт.Errorf("ТСекцияИмпорт.Обработать(): ошибка при проверке ограничения секции импорта\n\t%v", ош)
		}
		if ош = сам._СловаСекцииРазделить(); ош != nil {
			return мФмт.Errorf("ТСекцияИмпорт.Обработать(): ошибка при разделении слов секции ИМПОРТ и модуля\n\t%v", ош)
		}
		if ош := сам._ЕслиИмпортОдин(); ош != nil {
			return мФмт.Errorf("ТСекцияИмпорт.Обработать(): ошибка при проверке одной секции импорт\n\t%v", ош)
		}
		//Теперь нужно разобрать все слова импорта по именам модулей и алиасам
		if ош = сам._ИмпортРазобрать(); ош != nil {
			return мФмт.Errorf("ТСекцияИмпорт.Обработать(): ошибка при разборе секции импорта\n\t%v", ош)
		}
	}
	return nil
}

//По поледнему слову секции импорта вырезает слова для импорта
func (сам *ТСекцияИмпорт) _СловаСекцииРазделить() (ош error) {
	словаИмпорт := make(map[мИнт.ССловоНомерИмпорт]мИнт.ИСлово)
	//Первое слово ИМПОРТ уже отброшено. Разделитель секции нужен -- подскажет, где конец
	for адр := мИнт.ССловоНомерМодуль(0); адр <= сам.словоКонец; адр++ {
		словаИмпорт[мИнт.ССловоНомерИмпорт(адр)] = сам.СловаМодуля()[адр]
		delete(сам.СловаМодуля(), адр)
	}
	сам.словаИмпорт = словаИмпорт
	//Отладочный вывод для понять, что получилось
	мФмт.Printf("ТСекцияИмпорт._СловаСекцииРазделить(): всего слов секции ИМПОРТ=[%v]\n", len(сам.словаИмпорт))
	//Теперь надо изменить нумерацию слов в секции модуля
	//Получим все номера секции МОДУЛЬ
	номера := make(map[int]мИнт.ССловоНомерМодуль)
	ном := 0
	for адр := range сам.СловаМодуля() {
		номера[ном] = адр
		ном++
	}
	//Теперь перенумеруем слова модуля
	словаМодуля := make(map[мИнт.ССловоНомерМодуль]мИнт.ИСлово)
	for адр := 0; адр < len(номера); адр++ {
		словаМодуля[мИнт.ССловоНомерМодуль(адр)] = сам.СловаМодуля()[номера[(адр)]]
	}
	if ош = сам.СловаУст(словаМодуля); ош != nil {
		return мФмт.Errorf("ТСекцияИмпорт._СловаСекцииРазделить(): ошибка при обрезке слов модуля в секции импорта\n\t%v", ош)
	}
	//Отладочный вывод для понять, что получилось
	мФмт.Printf("ТСекцияИмпорт._СловаСекцииРазделить(): всего слов секции МОДУЛЬ=[%v]\n", len(сам.СловаМодуля()))
	for адр := мИнт.ССловоНомерИмпорт(0); адр < 15; адр++ {
		мФмт.Printf("\tадр=[%v]\t[%v]\n", адр, сам.СловаИмпорт()[адр].Слово())
	}
	return nil
}

func (сам *ТСекцияИмпорт) _ЕслиИмпортОдин() (ош error) {
	мКонс.Конс.Отладить("ТСекцияИмпорт._ЕслиИмпортОдин()")
	счётИмпорт := 0
	словоНом := мИнт.ССловоНомерМодуль(0)
	словаМодуляВсего := len(сам.СловаМодуля())
	for словоНом < мИнт.ССловоНомерМодуль(словаМодуляВсего) {
		слово := сам.СловаМодуля()[словоНом]
		for индекс := range мСлово.КсИмпорт {
			стрИмпорт := слово.Слово()
			if стрИмпорт == мСлово.КсИмпорт[индекс] {
				счётИмпорт++
			}
			if счётИмпорт > 1 {
				стрИсх := слово.Строка()
				ош = мФмт.Errorf("ТСекцияИмпорт._ЕслиИмпортОдин(): IMPORT два раза в одном модуле запрещён\n\t%v", стрИсх)
				return ош
			}
		}
		словоНом++
	}

	return nil
}

//Проверяет первое ключевое слово IMPORT
func (сам *ТСекцияИмпорт) _ЕслиИмпорт() (рез bool, ош error) {
	//Проверить наличие слова IMPORT
	словоИМПОРТ := сам.СловаМодуля()[0]
	окИмп, ош := мКлюч.Ключи.Проверить("IMPORT", мИнт.СКлюч(словоИМПОРТ.Слово()))
	if ош != nil {
		return false, мФмт.Errorf("ТСекцияИмпорт._ЕслиИмпорт(): ошибка пи поиске ключевого слова IMPORT\n\t%v", ош)
	}
	if !окИмп {
		return false, nil
	}
	//Слово IMPORT найдено
	слова, ош := мИнт.СловаМодуляОбрезать(сам.СловаМодуля())
	if ош != nil {
		return false, мФмт.Errorf("ТСекцияИмпорт._ЕслиИмпорт(): ошибка при обрезке слов\n\t%v", ош)
	}
	сам.СловаУст(слова)
	return true, nil
}

//Ищет разделитель в секции импорта
func (сам *ТСекцияИмпорт) _ЕслиИмпортОграничен() (ош error) {
	_КонецСекции := func(слово мИнт.ИСлово) bool {
		стр := слово.Слово()
		return стр == ";"
	}
	for счётСловоИмпорт := мИнт.ССловоНомерМодуль(0); int(счётСловоИмпорт) < len(сам.СловаМодуля()); счётСловоИмпорт++ {
		слово := сам.СловаМодуля()[счётСловоИмпорт]
		if _КонецСекции(слово) {
			сам.словоКонец = счётСловоИмпорт
			return nil
		}
	}
	return мФмт.Errorf("ТуИмпорт._ЕслиИмпортОграничен(): секция импорта ничем не ограничена\n")
}

// СловаИмпорт -- Возвращает слова импорта
func (сам *ТСекцияИмпорт) СловаИмпорт() (слова map[мИнт.ССловоНомерИмпорт]мИнт.ИСлово) {
	return сам.словаИмпорт
}

//Разбирает секцию импорта, хранит имена импортируемых модулей, их алиасов и номеров
func (сам *ТСекцияИмпорт) _ИмпортРазобрать() (ош error) {
	_ЕслиПуть := func(пИмя мИнт.ИСлово) (ок bool) {
		//Возвращае ток, если переданное слово, либо ИмяСтрого, либо точка
		бДа := пИмя.ЕслиИмяСтрого() || пИмя.Слово() == "."
		бНет := пИмя.Слово() == "," || пИмя.Слово() == ";" || пИмя.Слово() == ":="
		if бДа && !бНет {
			return true
		}
		return false
	}
	if len(сам.СловаИмпорт()) < 2 {
		return мФмт.Errorf("ТСекцияИмпорт._ИмпортРазобрать(): Объявленная секция импорта не может иметь меньше 2 слов\n\t%v", ош)
	}
	var (
		словоРаздел мИнт.ИСлово
		ок          bool
	)
	for адр := мИнт.ССловоНомерИмпорт(0); int(адр) < len(сам.СловаИмпорт()); адр++ {
		/* Слово +1 может быть:
		- запятая (0 -- реальное имя модуля)
		- ";" ( окончание секции импорта)
		- ":=" (0 -- алиас, +1 -- реальное имя модуля)
		если не эти три, то это -- сбой.
		*/
		//1. Контроль, что не вышли за предела словаря импорта
		if int(адр) < len(сам.СловаИмпорт())-1 {
			словоРаздел = сам.СловаИмпорт()[адр+1]

		}
		мФмт.Printf("Слово: %v.[%v]\n", адр, словоРаздел.Слово())
		{ //2. Проверим, а не конец ли это секции
			if ок, ош = мКлюч.Ключи.Проверить(";", мИнт.СКлюч(словоРаздел.Слово())); ош != nil {
				return мФмт.Errorf("ТСекцияИмпорт._ИмпортРазобрать(): ошибка при проверке разделителя `;`\n\t%v", ош)
			}
			if ок { //Достигнут конец секции
				return nil
			}
		}
		{ //3. Проверим, а не запятая ли это
			if ок, ош = мКлюч.Ключи.Проверить(",", мИнт.СКлюч(словоРаздел.Слово())); ош != nil {
				return мФмт.Errorf("ТСекцияИмпорт._ИмпортРазобрать(): ошибка при проверке разделителя `,`\n\t%v", ош)
			}
			if ок { //Предыдущее слово -- имя модуля
				словоМодуль := сам.СловаИмпорт()[адр]
				if ош = сам._МодульДоб(словоМодуль); ош != nil {
					return мФмт.Errorf("ТСекцияИмпорт._ИмпортРазобрать(): ошибка при добавлении модуля [%v]\n\t%v", словоМодуль.Слово(), ош)
				}
				//3.1 Теперь изменим адрес на следующий разделитель
				адр++
				continue
			}
		}
		{ //4. Проверим, а не разделитель алиасов ли это
			if ок, ош = мКлюч.Ключи.Проверить(":=", мИнт.СКлюч(словоРаздел.Слово())); ош != nil {
				return мФмт.Errorf("ТСекцияИмпорт._ИмпортРазобрать(): ошибка при проверке разделителя `:=`\n\t%v", ош)
			}
			if ок { //Предыдущее слово -- алиас модуля
				алиасМодуль := сам.СловаИмпорт()[адр]
				//4.1 Соберём имя модуля с его пакетом из частей (если оно такое)
				адр += 2
				путьМодуль := сам.СловаИмпорт()[адр]
				стрПуть := мИнт.ССлово("")
				for _ЕслиПуть(путьМодуль) {
					стрПуть += путьМодуль.Слово()
					адр++
					путьМодуль = сам.СловаИмпорт()[адр]
				}
				//4.2 Убедиться, что последняя литера не точка (а такое может быть)
				if путьМодуль.Слово() == "." {
					return мФмт.Errorf("ТСекцияИмпорт._ИмпортРазобрать(): ошибка при проверке имени пути модуля, путь=[%v]\n", стрПуть)
				}
				if ош = сам._АлиасДоб(алиасМодуль, стрПуть); ош != nil {
					return мФмт.Errorf("ТСекцияИмпорт._ИмпортРазобрать(): ошибка при добавлении алиаса и пути модуля [%v] [%v]\n\t%v", алиасМодуль.Слово(), стрПуть, ош)
				}
				//4.3 Поскольку счётчик адреса уже прибавлен -- его не меняем. Либо ",", либо ";"
			}
		}

	}
	return nil
}

//Добавляет имя модуля в общий список
func (сам *ТСекцияИмпорт) _МодульДоб(пСлово мИнт.ИСлово) (ош error) {
	if ош := мСмод.Модули.МодульДоб(пСлово, ""); ош != nil {
		return мФмт.Errorf("ТСекцияИмпорт._МодульДоб(): ошибка при добавлении сущности модуля [%v]\n\t%v", пСлово.Слово(), ош)
	}
	return nil
}

//Добавляет имя модуля и его алиас в общий список
func (сам *ТСекцияИмпорт) _АлиасДоб(пСлово мИнт.ИСлово, путьМодуль мИнт.ССлово) (ош error) {
	if ош := мСмод.Модули.МодульДоб(пСлово, путьМодуль); ош != nil {
		return мФмт.Errorf("ТСекцияИмпорт._МодульДоб(): ошибка при добавлении сущности алиаса и пути модуля [%v] [%v]\n\t%v", пСлово.Слово(), путьМодуль, ош)
	}
	return nil
}
