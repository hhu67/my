package main

import (
	"fmt"
	"os"
)

type Person struct {
    name string
    age  int
}

func polz() (string, int) {
	fmt.Println("Введите ваше имя")
	var pol1 string
	fmt.Scanln(&pol1)
	if pol1 == "Сергей" || pol1 == "сергей" {
		fmt.Println("Такая херь здесь не нужна")
		os.Exit(0)
	}
	fmt.Println("Введите ваш возраст")
	var pol2 int
	fmt.Scan(&pol2)
	return pol1, pol2
}

func main() {
	pol1, pol2 := polz()
	pers := Person{
		name: pol1,
		age: pol2,
	}
	if pers.age == 200 {
		fmt.Println("Ваше имя:", pers.name, "\n"+"Ты", pers.age)
	} else {
		fmt.Println("Ваше имя:", pers.name, "\n"+"Твой возраст:", pers.age)
	}
}