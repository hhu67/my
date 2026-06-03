package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Ты кто?")
	var pol string
	fmt.Scanln(&pol)
	ser := strings.ToLower(pol)
	if strings.HasPrefix(ser, "серг") {
		fmt.Println("Обнаружен сергей")
	} else {
		fmt.Println("Ты не сергей")
	}
}
