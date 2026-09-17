package main

import (
	"fmt"
	"math"
)

func main() {
	// 1
	fmt.Println("Задание 1")

	bannerWidth := 12
	bannerHeight := 8
	bannerArea := bannerWidth * bannerHeight
	fmt.Println("Площадь баннера:", bannerArea)

	halfBannerArea := bannerArea / 2
	fmt.Println("Половина площади баннера:", halfBannerArea)
	bannerBorderLength := (bannerWidth + bannerHeight) * 2
	fmt.Println("Длина бордера баннера:", bannerBorderLength)

	// 2
	fmt.Println("Задание 2")

	boxCount := 29
	leftoverBoxes := boxCount % 5
	fmt.Println("кол-во коробок, которые остались между делешками на 5 курьеров:", leftoverBoxes)

	// 3
	fmt.Println("Задание 3")

	tempMorning := 17
	tempAfternoon := 21
	tempEvening := 16
	totalTemp := tempMorning + tempAfternoon + tempEvening
	averageTemp := totalTemp / 3
	fmt.Println("Среднее кол-во температуры за день:", averageTemp)

	// 4
	fmt.Println("Задание 4")

	knownWords := 47.0
	wordsGoal := 120.0
	progressPercent := (knownWords / wordsGoal) * 100
	fmt.Println("Процент, который выполнен:", int(progressPercent), "%!")

	// 5
	fmt.Println("Задание 5")

	coins := 0
	fmt.Println("Начальный счёт:", coins)
	coins += 500
	fmt.Println("Счёт после выполнения задания:", coins)
	coins += 1200
	fmt.Println("Счёт после получения бонуса:", coins)
	coins /= 2
	fmt.Println("Счёт после траты половина монет в магазине:", coins)
	coins *= 2
	fmt.Println("Счёт после события в игре:", coins)
	coins -= 300
	fmt.Println("Счёт после траты 300 монет:", coins)

	// 6
	fmt.Println("Задание 6")

	participants := 42
	groupCount := 8

	participantsPerGroup := participants / groupCount
	fmt.Println("Кол-во студентов на каждую группу:", participantsPerGroup)

	// 7
	fmt.Println("Задание 7")

	fmt.Println(20 - 4*3) // результат 8, т.к. 20-4 не в скобках. а значит первым выполняется умножение

	fmt.Println((20 - 4) * 3) // результат 48, т.к. 20-4 в скобках. а значит первым выполняется это действие

	// 8
	squareValue := 81
	fmt.Println("Значение квадратного корня 81:", math.Sqrt(float64(squareValue)))

	multiplier := 5
	exponent := 2
	fmt.Println("Степень числа", multiplier, ":", math.Pow(float64(multiplier), float64(exponent)))
}
