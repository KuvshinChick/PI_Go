package main

import "fmt"

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	lastDigit := num % 10
	fmt.Println("Последняя цифра:", lastDigit)
}
