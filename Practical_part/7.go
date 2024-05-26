package main

import "fmt"

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	tens := (num / 10) % 10
	fmt.Println("Число десятков:", tens)
}
