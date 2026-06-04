package main

import (
	"fmt"
	"strings"
)

func polz() (string, string) {
	fmt.Println("Введите первое")
	var pol1 string
	fmt.Scanln(&pol1)
	fmt.Println("Введите второе")
	var pol2 string
	fmt.Scanln(&pol2)
	return pol1, pol2
}

func vowel(vow string) int {
	tolower := strings.ToLower(vow)
	wovel := 0
	for _, r := range tolower {
		if strings.ContainsRune("aeiyou", r) {
			wovel++
		}
	}
	return wovel
}

func main() {
	po1, po2 := polz()
	fmt.Println(vowel(po1))
	fmt.Println(po1, po2)
}