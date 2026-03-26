package main

import "fmt"

// Задание: Калькулятор — возьми 2 числа от юзера (fmt.Scanln), посчитай +,-,*,/,%, сравни >/<. Обработай деление на 0 (if).
// Выведи в красивом формате. 3 примера тестов. Commit: "Day2 calculator"

func main() {
	var a, b int
	fmt.Print("Введи 1 число: ")
	fmt.Scanln(&a)
	fmt.Print("Введи 2 число: ")
	fmt.Scanln(&b)

	sum := a + b
	difference := a - b
	multy := a * b

	if b == 0 {
		fmt.Println("Деление на 0 невозможно")
		return
	}

	delit := a / b
	mod := a % b

	fmt.Printf("Сумма ваших чисел равна: %d\n", sum)
	fmt.Printf("Разница ваших чисел равна: %d\n", difference)
	fmt.Printf("Умножение ваших чисел равно: %d\n", multy)
	fmt.Printf("Деление ваших чисел равно: %d\n", delit)
	fmt.Printf("Остаток: %d\n", mod)
}
