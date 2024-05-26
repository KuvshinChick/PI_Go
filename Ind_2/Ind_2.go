package main

import "fmt"

func main() {
	var priceCandy, priceCookie, priceApple, x, y, z float64

	fmt.Println("Введите стоимость 1 кг конфет:")
	fmt.Scan(&priceCandy)
	fmt.Println("Введите стоимость 1 кг печенья:")
	fmt.Scan(&priceCookie)
	fmt.Println("Введите стоимость 1 кг яблок:")
	fmt.Scan(&priceApple)
	fmt.Println("Введите количество купленных килограмм конфет, печенья и яблок соответственно:")
	fmt.Scan(&x, &y, &z)

	totalCost := priceCandy*x + priceCookie*y + priceApple*z
	fmt.Printf("Стоимость покупки: %.2f\n", totalCost)
}
