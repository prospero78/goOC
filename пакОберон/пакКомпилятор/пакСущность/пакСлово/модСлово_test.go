// модСлово_test
package пакСлово

/*
	Модуль предоставляет тест для проверки правильности работы ТуСлово
*/

import (
	пакТест "testing"
)

var (
	слХор  = []string{"ТестСлово", "Упсь", "Ням", "почка", "тоже", "алё", "ягодка", "металл", "рок-енд-ролл", "страховка"}
	слПлох = []string{".ТестСлово", "2Упсь", "#Ням", "!почка", "?тоже", "+алё", "*ягодка", "%металл", "~рок-енд-ролл", "$страховка"}
)

// Проверяет, что первая литера реально буква
func Test_ЕслиПерваяБуква(тест *пакТест.T) {
	for _, стрСлово := range слХор {
		if !ЕслиПерваяБуква(стрСлово) {
			тест.Errorf("ПакСлово.ЕслиПерваяБуква(): ошибка в правильном слове = \"" + стрСлово + "\"")
		}
	}
	for _, стрСлово := range слПлох {
		if ЕслиПерваяБуква(стрСлово) {
			тест.Errorf("ПакСлово.ЕслиПерваяБуква(): ошибка в правильном слове = \"" + стрСлово + "\"")
		}
	}

}

// Проверяет, что литера реально буква
func Test_ЕслиБуква(тест *пакТест.T) {
	for _, стрСлово := range слХор {
		стрЛит := string([]rune(стрСлово)[0])
		if !ЕслиБуква(стрЛит) {
			тест.Errorf("ПакСлово.ЕслиБуква(): ошибка в правильном слове = \"" + стрСлово + "\" " + стрЛит)
		}
	}
	for _, стрСлово := range слПлох {
		стрЛит := string([]rune(стрСлово)[0])
		if ЕслиБуква(стрЛит) {
			тест.Errorf("ПакСлово.ЕслиБуква(): ошибка в правильном слове = \"" + стрСлово + "\" " + стрЛит)
		}
	}
}

var (
	ЦПрав   = []rune("1.23233435.5464356.657.67.567.8.768.56785678.56")
	ЦНеПрав = []rune("ываргролрдывбюячсмкп!\"№reterkfdlkfdqiesapdfksdfgsdpfgksdfpgo")
)

// Проверяет, что первая литера реально цифра
func Test_ЕслиЦифра(тест *пакТест.T) {
	for _, стрЛит := range ЦПрав {
		if !ЕслиЦифра(string(стрЛит)) {
			тест.Errorf("ПакСлово.ЕслиЦифра(): ошибка в правильном слове = \"" + string(стрЛит))
		}
	}
	for _, стрЛит := range ЦНеПрав {
		if ЕслиЦифра(string(стрЛит)) {
			тест.Errorf("ПакСлово.ЕслиЦифра(): ошибка в правильном слове = \"" + string(стрЛит))
		}
	}

}

var (
	ЦЧислаПрав   = []string{"1.23", "233", ".435", "0.0", "0.005464356"}
	ЦЧислаНеПрав = []string{"a.str", "0.a", "123.6b6", ".567d555"}
)

// Проверяет, что строка не число
func Test_ЕслиНеЦифра(тест *пакТест.T) {
	for _, стрЧисло := range ЦЧислаПрав {
		if ЕслиЦифра(стрЧисло) {
			тест.Errorf("ПакСлово.ЕслиНеЦифра(): ошибка в правильном слове БЕЗ ОШИБОК = \"" + стрЧисло)
		}
	}
	for _, стрЧисло := range ЦЧислаНеПрав {
		if !ЕслиНеЦифра(стрЧисло) {
			тест.Errorf("ПакСлово.ЕслиНеЦифра(): ошибка в правильном ОШИБОЧНОМ слове = \"" + стрЧисло)
		}
	}

}
