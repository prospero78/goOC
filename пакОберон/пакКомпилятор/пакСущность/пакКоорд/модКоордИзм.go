package пакКоорд

// модКоордИзм

/*
	Модуль описывает тип с изменяемыми координатами
*/

import (
	пакКонс "../../../пакКонсоль"
	пакФмт "fmt"
)

//ТуКоордИзм -- тип для изменяемых координат
type ТуКоордИзм struct {
	стр  ТЦелСтр
	поз  ТЦелПоз
	конс *пакКонс.ТуКонсоль
}

//НовыйИзм -- возвращает новый экземпляр изменяемых координат
func НовыйИзм(пСтр ТЦелСтр, пПоз ТЦелПоз) (ки *ТуКоордИзм, ош error) {
	if пСтр >= 1 {
		if пПоз >= 0 {
			ки = &ТуКоордИзм{
				стр:  пСтр,
				поз:  пПоз,
				конс: пакКонс.Конс,
			}
		} else {
			ош = пакФмт.Errorf("пакКоорд.НовыйИзм(): пПоз не может быть меньше 0")
			return nil, ош
		}
	} else {
		ош = пакФмт.Errorf("пакКоорд.НовыйИзм(): пСтр не может быть меньше 1")
		return nil, ош
	}
	return ки, nil
}

//Поз -- возвращает позицию литеры или слова в строке (глядя в составе чего использовано)
func (сам *ТуКоордИзм) Поз() (поз ТЦелПоз) {
	return сам.поз
}

//ПозДоб -- увеличение координаты Поз на 1
func (сам *ТуКоордИзм) ПозДоб() {
	сам.поз += 1
}

func (сам *ТуКоордИзм) Поз_Сброс() {
	сам.поз = 0
}

func (сам *ТуКоордИзм) Поз_Уст(пПоз ТЦелПоз) (ош error) {
	if пПоз < 0 {
		ош = пакФмт.Errorf("ТуКоордИзм.Поз_Уст(): отрицательный пПоз запрещён, пПоз=%v", пПоз)
		return ош
	}
	сам.поз = пПоз
	return ош
}

func (сам *ТуКоордИзм) Стр() (стр ТЦелСтр) {
	return сам.стр
}

func (сам *ТуКоордИзм) Стр_Доб() {
	сам.стр++
}

// Для соответствия интерфейсу ИКоордФикс
func (сам *ТуКоордИзм) String() string {
	return "*ТуКоордИзм{стр:" + пакФмт.Sprintf("%v", сам.стр) + ", поз:" + пакФмт.Sprintf("%v", сам.поз) + "}"
}
