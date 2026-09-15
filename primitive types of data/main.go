package main

import (
	"fmt"
)

func main() {
	// 1
	fmt.Println("Задание 1")

	age := 14
	fmt.Print("Мне", age, "лет!")

	age = 15
	fmt.Print("Мне сегодня исполянется", age, "лет!")

	// 2
	fmt.Println("Задание 2")

	height := 167 //соврал))
	const snm = "сантиметров!"
	fmt.Println("Мой рост", height, snm)

	height_in_meters := 1.67 // ЫЫЫ СИКС-СЕВЕН
	const mtrs = "метров!"
	fmt.Println("Мой рост", height_in_meters, mtrs)

	// 3
	fmt.Println("Задание 3") // если вам пофиг на это и надо только чтобы в коде было показана цифра задания, то я могу не писать

	isStudent := true                     // я итак учусь..
	fmt.Println("Учусь ли я:", isStudent) // а чо так много принтов за три дз?

	// 4
	fmt.Println("Задание 4")

	temperature := 17
	fmt.Println("Температура:", temperature, "прохладная!")

	// 5
	fmt.Println("Задание 5")

	favoriteQuote := "Veeni, vidi, vici" // Гай Юлий Цезарь, любимка !!!
	fmt.Println("Моя любимая цитата:", favoriteQuote)

	// 6
	fmt.Println("Задание 6")
	const PI = 3.14
	fmt.Println("Приблизительное значение числа pi:", PI)

	// А ЧО НИ ОДНОГО ЗАДАНИЕ НА СМЕНУ ТИПА ДАННЫХ!?!??!?!
}
