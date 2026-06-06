package main

import (
	"fmt"
)

func main() {
	mape := map[int]string{
		1: "sergey",
		2: "sergay",
	}
	fmt.Println("Введите свое имя")
	var pol string
	fmt.Scanln(&pol)
	mape[len(mape)+1] = pol
	fmt.Println("Ваше имя:",mape[len(mape)])
	fmt.Println("Введите число значение каторого хотите увидеть")
	var i int
	fmt.Scan(&i)
	value, ok := mape[i]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("значение",i,"не найдено, значение поиска:",ok)
	}
}