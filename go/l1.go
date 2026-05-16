package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	fmt.Println("Your name ")
	name, _ := read.ReadString('\n')
	fmt.Println("Your name %s", name)
}
