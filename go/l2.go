package main
import (
	"fmt"
)
func polz() {
	fmt.Println("Первое число ")
	var x1 int
	fmt.Scan(&x1)
	fmt.Println("Второе число ")
	var x2 int
	fmt.Scan(&x2)
	x3 := x1+x2
	fmt.Println(x1 , "+" , x2 , "=" , x3)
}
func main() {
	for {
		fmt.Println("1. Начать\n2. Закончить")
		var choise int
		fmt.Scan(&choise)
		if choise == 1 {
			polz()
		} else if choise == 2 {
			return
		}
	}
}
