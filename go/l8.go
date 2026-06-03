package main

import (
	"fmt"
)

func polz() []int {
	slice := []int{}
	for i := 0; i < 10; i++ {
		fmt.Println("Введите числа")
		var pol int
		fmt.Scan(&pol)
		slice = append(slice, pol)
	}
	return slice
}

func main() {
	packages := polz()
	lenp := len(packages) -1
	for i := 0; i < lenp; i++ {
		for j := 1; j <= lenp - i; j++ {
			if packages[j-1] > packages[j] {
				packages[j-1], packages[j] = packages[j], packages[j-1]
			}
		}
	}
	fmt.Println(packages)
}