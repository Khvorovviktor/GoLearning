package main

import "fmt"

func main() {
	// 1
	fmt.Println("1 Задание") //буду так делать первое время, чтобы запомнить как писать принт

	schooling := 8
	fmt.Println("Сколько лет обучения в школе:", schooling)

	schooling = 9
	fmt.Println("Обновленное кол-во лет обучения в школе:", schooling)

	// 2
	fmt.Println("2 Задание")

	name := "Vladislav"
	fmt.Println("Hello,", name, "!")
	name = "Artur"
	fmt.Println("Разговор в метро: -Мужчина, вы мне на ногу наступаете! -Я знаю. -Мужчина, как вас зовут?! -Мое имя:", name, "-Так,", name, "УБЕРИТЕ НОГУУ!!!!") //это не смешно, но я старался

	// 3
	fmt.Println("3 Задание")

	steps := 0
	fmt.Println("Кол-во шагов в течении дня:", steps, "!")

	steps = 2000
	fmt.Println("Обновлённое кол-во шагов в течении дня:", steps, "!")
	fmt.Println("Хорошая работа! Вы уже на пути к ежедневной цели!")

	// 4
	fmt.Println("4 Задание")

	largeNumber := 5000001
	fmt.Println("Число больше 5 миллионов:", largeNumber)

	// 5
	fmt.Println("5 Задание")

	const breakTime = 15
	fmt.Println("Перемена длится:", breakTime, "минут!")

	// breakTime = 20 - Будет ошибка, т.к. нельзя изменить константу, потому что она постоянна.
}
