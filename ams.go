package main

import "fmt"

func main() {
	var a, b int = 0, 0
	fmt.Print("Введи 1 число: ")
	fmt.Scanln(&a)
	fmt.Print("Введи 2 число: ")
	fmt.Scanln(&b)

	c := a - b
	fmt.Printf("Разница: %d\n", c)
}
