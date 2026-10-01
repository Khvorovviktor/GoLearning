package main

import "fmt"

func main() {
	// Т З №1

	const BaseTariff = 0.45    // цена за 1 кВт/ч
	const HighLoadTax = 0.15   // налог на выскокое потребление, 15%
	const NightDiscount = 0.30 // ночная скидка, 30%

	name := "none"           // название прибора
	power := 0               // сколько ватт, т.е. мощности в телевизоре
	timeOfWork := 0          // время работы, часов
	NightmareRegime := false // true or false, ночной режим
	category := "none"       // категория товара

	// Part 2, Функционал

	for {
		fmt.Println("Введите название прибора/девайса, или, если хотите остановитть работу цикла, то введите done: ")
		fmt.Scan(&name)

		if name == "done" {
			return
			fmt.Println("Расчёт завершен.")
		}

		fmt.Println("Введите мощность прибора/девайса в ватт (W): ")
		fmt.Scan(&power)

		fmt.Println("Введите время работы прибора/девайса в часах (h): ")
		fmt.Scan(&timeOfWork)

		fmt.Println("Введите есть ли у прибора/девайса ночной режим (true/false): ")
		fmt.Scan(&NightmareRegime)

		expdtOfkWpH := (power * timeOfWork) / 1000
		StockCosts := float64(expdtOfkWpH) * BaseTariff

		if NightmareRegime == true {
			StockCosts = StockCosts * (1 - NightDiscount)
		} else {
			fmt.Println("В вашем девайсе нету ночного режима, поэтому скидка не прибавляется.")
		}

		if expdtOfkWpH >= 10 {
			StockCosts = StockCosts + HighLoadTax
		}

		switch {
		case power < 100:
			category = "Экономный"
		case power >= 100 && power <= 1000:
			category = "Стандартный"
		case power > 1000:
			category = "Мощный"
		}

		// Часть 3, вывод

		fmt.Println("--Отчёт по прибору--")
		fmt.Printf("Название прибора, или слово done: %s, категория: %s\n", name, category)
		fmt.Printf("Расход: %.2f кВт*ч\n", float64(expdtOfkWpH))
		fmt.Printf("Итоговая стоимость: %.2f (с учётом налогов)\n", float64(StockCosts))
		fmt.Printf("Ночной режим (true/false): %t\n", NightmareRegime)

		fmt.Println("---------")
	}
}
