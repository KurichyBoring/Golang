package main

import "fmt"

func main() {
	name := "hz"
	var age int = 30
	country := "Msk"
	var salary float64 = 500000.00
	isGo := true

	fmt.Printf("Имя: %s, Возраст: %d, Город: %s, Зарплата: %.0f$, Go: %t\n", name, age, country, salary, isGo)
}
