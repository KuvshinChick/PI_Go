package main

import "fmt"

func main() {
	var a int = 8
	const b int = 10
	a = a + b
	// b = b + a нельзя изменить значение константы, удаляем эту строку
	fmt.Println(a)
}
