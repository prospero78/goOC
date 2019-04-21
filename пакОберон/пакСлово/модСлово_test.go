package пакСлово

/*
	Модуль предоставляет тест для проверки правильности работы ТуСлово
*/

import (
	мАбс "../пакАбстракция"
	//мТест "testing"
)

var (
	слХор  = []мАбс.ССлово{"ТестСлово", "Упсь", "Ням", "почка", "тоже", "алё", "ягодка", "металл", "рок-енд-ролл", "страховка"}
	слПлох = []мАбс.ССлово{".ТестСлово", "2Упсь", "#Ням", "!почка", "?тоже", "+алё", "*ягодка", "%металл", "~рок-енд-ролл", "$страховка"}
)

var (
	ЦПрав   = []rune("1.23233435.5464356.657.67.567.8.768.56785678.56")
	ЦНеПрав = []rune("ываргролрдывбюячсмкп!\"№reterkfdlkfdqiesapdfksdfgsdpfgksdfpgo")
)

var (
	ЦЧислаПрав   = []мАбс.ССлово{"1.23", "233", ".435", "0.0", "0.005464356"}
	ЦЧислаНеПрав = []мАбс.ССлово{"a.str", "0.a", "123.6b6", ".567d555"}
)
