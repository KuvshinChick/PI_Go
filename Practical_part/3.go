package main

import "fmt"

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	num = num * 2
	num += 100
	fmt.Println("Результат:", num)
}
