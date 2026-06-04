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

func recovery(rec string) string {
	runes := []rune(rec)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	po1, po2 := polz()
	fmt.Println("Вот колличество гласных в первом", vowel(po1))
	fmt.Println("Вот слова", po1, po2)
	fmt.Println("Вот перевернутое первое слово", recovery(po1))
}