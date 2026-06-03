package main

import (
	"fmt"
)

func logic() string {
	fmt.Println("Твой возраст")
	var x int
	_, err := fmt.Scan(&x)
	if err != nil {
		return "еблан"
	}
	var phrase string
	switch {
	case x < 0:
		phrase = "Ты еще не родился"
	case x < 12:
		phrase = "Малолетка ебаная"
	case x >= 12 && x <= 19:
		phrase = "Ты подросток"
	case x >= 20 && x <= 65:
		phrase = "Ебать ты средний"
	case x > 65:
		phrase = "Все давай в могилу"
	default:
		phrase = "не"
	}
	return phrase
}

func main() {
	fmt.Println("Как звать тебя")
	var name string
	fmt.Scanln(&name)
	for {
		fmt.Println("1. Начать\n2. Закончить")
		var pol string
		fmt.Scanln(&pol)
		if pol == "1" {
			phr := logic()
			fmt.Println("Тебя зовут", name, ",", phr)
		} else if pol == "2" {
			break
		} else {
			fmt.Println("Не то")
		}
	}
}
