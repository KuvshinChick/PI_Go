package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a) // читаем переменную 'a' с консоли
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b) // читаем переменную 'b' с консоли
	a = a * a
	b = b * b
	c := a + b
	fmt.Println("Результат:", c)
}
