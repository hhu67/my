package main

import (
	"fmt"
)

func main() {
	fmt.Println("Сколько")
	var pol string
	fmt.Scanln(&pol)
	fmt.Println("Кому?")
	var pol2 string
	fmt.Scanln(&pol2)
	
	var sergay []string
	sergay = append(sergay, pol2, pol)
	if sergay[1] == "31" || sergay[1] == "21" || sergay[1] == "41" || sergay[1] == "51" || sergay[1] == "61" || sergay[1] == "71" || sergay[1] == "81" || sergay[1] == "91" || sergay[1] == "101" {
		sergay = append(sergay, "год")
	} else {
		sergay = append(sergay, "лет")
	}
	fmt.Println(sergay[0], sergay[1], sergay[2])
}
