package main

import (
	"fmt"
)

func main() {
	grades := []int{}
	grades = append(grades, 4, 3, 5, 2, 4, 5, 3)
	good := []int{}
	for i := len(grades) - 1; i >= 0; i-- {
		if grades[i] >= 4 {
			good = append(good, grades[i])
			grades = append(grades[:i], grades[i+1:]...)
		}
	}
	fmt.Println(grades[0:3])
}