/*	Своеобразный аналог игры в пинг-понг или спора философов (яйцо-курица),
	который имитирует столкновение двух армий.
	
	Сперва задаётся имя полководца и  численность каждой из армий
	(Левый и правый фланги, центр, знаменосец),
	после же полководцы независимо друг от друга атакуют армию противника,
	нанося случайным типом воинов определённый урон.
*/
package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Army struct {
	Warlord	string	// Имя полководца
	FlangL	int		// Левый фланг
	FlangR	int		// Правый фланг
	Center	int		// Центр
	Pennant	int		// Знаменосец
	TotalCount int	// Общий размер войска
}

func main(){
	first	:= &Army{}
	second	:= &Army{}
	oldF	:= Army{}
	oldS	:= Army{}
	
	fmt.Println("Состав первого войска:")
	first = setArmy()
	first.TotalCount = totalCount(first)
	oldF = *first
	
	fmt.Println("Состав второго войска:")
	second = setArmy()
	second.TotalCount = totalCount(second)
	oldS = *second	
	
	var win int = 0
	
	go battle(1, oldS, oldF, second, first, &win)
	go battle(2, oldF, oldS, first, second, &win)
	
	<-time.After(time.Second * 5)
}

// Генерация ходов до победы или поражения. ar - противник, my - кто ходит
func battle(num int, old, oldMy Army, ar, my *Army, win *int){
	for ; *win == 0; {
		if ar.TotalCount <= 0 {
			*win = 1
			
			fmt.Printf("\n\n\t\t\tПобеду одержал полководец %s! В его войске осталось %d воинов из %d.\n \tЛевый фланг: %d [из %d]\n \tПравый фланг: %d [из %d]\n \tЦентр: %d [из %d]\n \tЗнаменосец: %d [из %d]\n",
			my.Warlord, my.TotalCount, oldMy.TotalCount, my.FlangL, oldMy.FlangL, my.FlangR, oldMy.FlangR, my.Center, oldMy.Center, my.Pennant, oldMy.Pennant)
			fmt.Printf("\n\n\t\t\tПолководец %s потерпел поражение. В его войске осталось %d воинов из %d.\n \tЛевый фланг: %d [из %d]\n \tПравый фланг: %d [из %d]\n \tЦентр: %d [из %d]\n \tЗнаменосец: %d [из %d]\n",
			ar.Warlord, ar.TotalCount, old.TotalCount, ar.FlangL, old.FlangL, ar.FlangR, old.FlangR, ar.Center, old.Center, ar.Pennant, old.Pennant)
		} else {
			punch(num, old, ar, my)
		}
	}
}

// Генерация одного хода
func punch(num int, old Army, ar, oth *Army) {
	if (oth.TotalCount>0)&&(ar.TotalCount>0){
		attack	:= []string{"Лучники", "Лёгкая кавалерия", "Тяжёлая кавалерия", "Копейщики", "Пехота"}
		line	:= []string{"по левому флангу", "по првому флангу", "по центру", "по знаменосцу"}
		
		var dmg		int
		var before	int
		var now		int
		
		randAttack	:= rand.Intn(len(attack))
		
		var randLine int
		for ok := 0; ok<1; {
			randLine	= rand.Intn(len(line))
			switch randLine{
				case 0:
					if (ar.FlangL > 0)&&(ar.FlangL > 0) {
						ok = 1
					} else {
						ok--
					}
				case 1:
					if (ar.FlangR > 0)&&(ar.FlangR > 0) {
						ok = 1
					} else {
						ok--
					}
				case 2:
					if (ar.Center > 0)&&(ar.Center > 0) {
						ok = 1
					} else {
						ok--
					}
				case 3:
					if (ar.Pennant > 0)&&(ar.Pennant > 0) {
						ok = 1
					} else {
						ok--
					}
			}
			if ok < -3 {
				fmt.Println("Сражение закончилось!")
				randLine = -1
				ok = 1
			}
		}
		
		switch randAttack{
			case 0:
				dmg = 30
			case 1:
				dmg = 50
			case 2:
				dmg = 90
			case 3:
				dmg = 45
			case 4:
				dmg = 10
		}
		
		switch randLine{
			case -1:
				fmt.Println("Армия победителей чуть не пошла добивать проигравших")
			case 0:
				before = old.FlangL
				now = calcNow(ar.FlangL, dmg)
			case 1:
				before = old.FlangR
				now = calcNow(ar.FlangR, dmg)
			case 2:
				before = old.Center
				now = calcNow(ar.Center, dmg)
			case 3:
				before = old.Pennant
				now = calcNow(ar.Pennant, dmg)
		}
		
		updArmy(randLine, ar, dmg)
		ar.TotalCount = totalCount(ar)
		
		if randLine != -1 {
			fmt.Printf("Полководец %s (%d): \t %s, бей %s!\n[-%d]\n\t Часть: [%d/%d]\n\t Всего: [%d/%d]\n\n",
			oth.Warlord, num, attack[randAttack], line[randLine], dmg, now, before, ar.TotalCount, old.TotalCount)
		}
	}
}

// Заполнение данных об одном войске
func setArmy () *Army{
	var name string
	var l,r,c,p int
	for i := 0; i < 5; {
		switch i {
			case 0:
				fmt.Println("\tИмя полководца:")
				fmt.Scanf("%s\n",&name)
				i = 1
			case 1:
				fmt.Println("\tВоинов на левом фланге:")
				fmt.Scanf("%d\n",&l)
				i = 2
			case 2:
				fmt.Println("\tВоинов на правом фланге:")
				fmt.Scanf("%d\n",&r)
				i = 3
			case 3:
				fmt.Println("\tВоинов по центру:")
				fmt.Scanf("%d\n",&c)
				i = 4
			case 4:
				fmt.Println("\tВоинов вокруг знаменосца:")
				fmt.Scanf("%d\n",&p)
				i = 5
		}
	}
	
	ar := &Army{
		Warlord:	name,
		FlangL:		l,
		FlangR:		r,
		Center: 	c,
		Pennant:	p,
	}
	return ar
}

// Подсчёт количества воинов в войске
func totalCount(ar *Army) int {
	tc := ar.FlangL + ar.FlangR + ar.Center + ar.Pennant
	if tc < 0 {
		return 0
	} else {
		return tc
	}
}

// Обновление значения конкретной части войска по её номеру
func updArmy(num int, ar *Army, val int) {

	switch num{
		case 0:
			ar.FlangL -= val
			if ar.FlangL < 0 {
				ar.FlangL = 0
			}
		case 1:
			ar.FlangR -= val
			if ar.FlangR < 0 {
				ar.FlangR = 0
			}
		case 2:
			ar.Center -= val
			if ar.Center < 0 {
				ar.Center = 0
			}
		case 3:
			ar.Pennant -= val
			if ar.Pennant < 0 {
				ar.Pennant = 0
			}
	}
}

// Вычисление количества воинов после атаки
func calcNow(before, dmg int) int {
	temp := before - dmg
	if temp < 0 {
		return 0
	} else {
		return temp
	}
}
