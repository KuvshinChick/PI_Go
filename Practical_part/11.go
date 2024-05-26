package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Print("Введите значение a: ")
	fmt.Scan(&a)
	fmt.Print("Введите значение b: ")
	fmt.Scan(&b)

	volume := (4.0 / 3.0) * math.Pi * a * b * b
	surfaceArea := 2 * math.Pi * b * (b + a*math.Erf(math.Pi/2))

	fmt.Printf("Объем: %.2f\n", volume)
	fmt.Printf("Площадь поверхности: %.2f\n", surfaceArea)
}
