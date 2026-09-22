package main

import (
	"fmt"
)

func main() {
	// 1 часть
	const BaseRate = 5.50    // обычная цена за 1 кг груза
	const TaxRate = 0.12     // налог 12%
	const DistanceRate = 2.0 // цена за 1 км доставки
	const FragileFee = 0.2   // коэффициент хрупкости, т.е. наценка 20%

	// 2 часть
	var userName string
	fmt.Print("Введите имя:")
	fmt.Scanln(&userName)

	var weight float64
	fmt.Print("Введите вес груза (кг):")
	fmt.Scanln(&weight)

	var distance float64
	fmt.Print("Введите дистанцию (км):")
	fmt.Scanln(&distance)

	var quantityFragile float64
	fmt.Print("Введите кол-во хрупких упаковок:")
	fmt.Scanln(&quantityFragile)

	baseRate := (weight*BaseRate)*(1+FragileFee*quantityFragile) + (distance * DistanceRate)
	result := baseRate + TaxRate

	fmt.Println("--Отчет о доставке--")
	fmt.Printf("Имя отправителя: %s\n", userName)
	fmt.Printf("Итоговая стоимость: %.2f\n", result)
}
