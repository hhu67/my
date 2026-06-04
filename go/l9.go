package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Введите фразу")
	var pol string
	fmt.Scanln(&pol)
	poll := strings.ToLower(pol)
	if strings.Contains(poll, "go") {
		x := strings.Count(poll, "go")
		fmt.Println(x)
	} else {
		x2 := poll + " go"
		fmt.Println(x2)
	}
}