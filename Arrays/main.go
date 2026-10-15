package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1
	fmt.Println("Задание 1")
	var in string
	var input string
	var forRunning [10]string
	var forWalking [10]string
	for index := range forWalking {
		fmt.Println("Введите упражнение для ходьбы (или слово done чтобы остановить цикл):")
		fmt.Scanln(&input)
		if input == "done" || input == "" {
			fmt.Printf("Вы ввели %s, цикл остановлен\n", input)
			break
		}
		forWalking[index] = input
	}

	for _, value := range forWalking {
		fmt.Printf("Вывод массива с упражнениями для ходьбы: %s\n", value)
	}

	for index := range forRunning {
		fmt.Println("Введите упражнение для бега (или слово done чтобы остановить цикл):")

		fmt.Scanln(&in)
		if in == "done" {
			fmt.Printf("Вы ввели %s, цикл остановлен\n", in)
			break
		}

		forRunning[index] = in
	}

	for _, value := range forRunning {
		fmt.Printf("Вывод массива с упражнениями для бега: %s\n", value)
	}

	// 2
	fmt.Println("Задание 2")

	subjectsList := [3]string{"Физика", "Химия", "География"}

	fmt.Printf("Первый элемент массива subjectsList: %s, последний элемент данного массива: %s\n", subjectsList[0], subjectsList[len(subjectsList)-1])
	fmt.Printf("Весь массив subjectsList: %v\n", subjectsList)

	subjectsList[0] = "Биология"

	fmt.Printf("Обновлённый массив subjectsList: %v\n", subjectsList)

	// 3
	fmt.Println("Задание 3")

	var name string
	var age string
	var city string

	fmt.Println("Введите имя:")
	fmt.Scanln(&name)
	fmt.Printf("Имя: %s\n", name)

	fmt.Println("Введите возраст:")
	fmt.Scanln(&age)
	fmt.Printf("Возраст: %s\n", age)

	fmt.Println("Введите город проживания:")
	fmt.Scanln(&city)
	fmt.Printf("Город: %s\n", city)

	if name == "" || age == "" || city == "" {
		name = "Неизвестное имя"
		age = "Неизвестный возраст"
		city = "Неизвестный город"
	}

	personData := [3]string{name, age, city}

	fmt.Printf("Имя: %v, возраст: %v, город проживания: %v\n", personData[0], personData[1], personData[2])

	// 4
	fmt.Println("Задание 4")

	var num int
	var number int
	var number3 int
	var num4 int
	var num5 int

	fmt.Println("Введите 1 число:")
	fmt.Scanln(&num)

	fmt.Println("Введите 2 число:")
	fmt.Scanln(&number)

	fmt.Println("Введите 3 число:")
	fmt.Scanln(&number3)

	fmt.Println("Введите 4 число:")
	fmt.Scanln(&num4)

	fmt.Println("Введите 5 число:")
	fmt.Scanln(&num5)

	numbersList := [5]int{num, number, number3, num4, num5}

	foundNum := false
	for index, value := range numbersList {
		if value == 3 {
			fmt.Printf("В массиве numbersList есть, число 3, его индекс: %d\n", index)
			foundNum = true
			break
		}
	}

	if !foundNum {
		fmt.Println("В массиве numbersList нет числа 3")
	}

	// 5
	fmt.Println("Задание 5")

	var friend string
	var friend2 string
	var friend3 string
	var friend4 string

	fmt.Println("Введите имя первого друга:")
	fmt.Scanln(&friend)

	fmt.Println("Введите имя второго друга:")
	fmt.Scanln(&friend2)

	fmt.Println("Введите имя третьего друга:")
	fmt.Scanln(&friend3)

	fmt.Println("Введите имя четвертого друга:")
	fmt.Scanln(&friend4)

	friendsList := [4]string{strings.ToLower(friend), strings.ToLower(friend2), strings.ToLower(friend3), strings.ToLower(friend4)}

	ifBekbolat := false

	for _, value := range friendsList {
		if value == "bekbolat" || value == "бекболат" {
			ifBekbolat = true
			break
		}
	}

	if ifBekbolat {
		fmt.Printf("В моём списке друзей есть друг по имени Бекболат, мне очень повезло!\n")
	} else {
		fmt.Printf("В моём списке друзей нет друга по имени Бекболат, мне не повезло, но я всё равно люблю своих друзей!\n")
	}

	// 6
	fmt.Println("Задание 6")

	var number1 int
	var number2 int
	var num3 int

	var n1 int
	var n2 int
	var n3 int

	fmt.Println("Введите 1 число:")
	fmt.Scanln(&number1)

	fmt.Println("Введите 2 число:")
	fmt.Scanln(&number2)

	fmt.Println("Введите 3 число:")
	fmt.Scanln(&num3)

	fmt.Println("Введите 4 число:")
	fmt.Scanln(&n1)

	fmt.Println("Введите 5 число:")
	fmt.Scanln(&n2)

	fmt.Println("Введите 6 число:")
	fmt.Scanln(&n3)

	firstList := [3]int{number1, number2, num3}

	secondList := [3]int{n1, n2, n3}

	ifTheSame := true
	for index := range firstList {
		if firstList[index] != secondList[index] || len(firstList) != len(secondList) {
			ifTheSame = false
			break
		}
	}

	if ifTheSame {
		fmt.Println("Списки одинаковы.")
	} else {
		fmt.Println("Списки различны.")
	}

	// 7
	fmt.Println("Задание 7")

	myWishList := [5]string{"Выздоровить", "Поменять стекло на телефоне", "НОРМАЛЬНЫЕ наушники", "Хорошую мышку", "Другую клавиатуру"}

	friendWishList := [5]string{"Не ходить в школу", "Хороший телефон", "Хороший интернет", "Поменять стекло на планшете", "Компьютер"}

	var resultList [10]string

	for index := range myWishList {
		resultList[index] = myWishList[index]
	}

	for index := range friendWishList {
		resultList[len(myWishList)+index] = friendWishList[index]
	}

	fmt.Println("Итоговый виш лист:", resultList)

	// 8
	fmt.Println("Задание 8")

	var toy1 string
	var toy2 string
	var toy3 string
	var toy4 string

	fmt.Println("Введите первую игрушку:")
	fmt.Scanln(&toy1)

	fmt.Println("Введите вторую игрушку:")
	fmt.Scanln(&toy2)

	fmt.Println("Введите третью игрушку:")
	fmt.Scanln(&toy3)

	fmt.Println("Введите четвертую игрушку:")
	fmt.Scanln(&toy4)

	toyList := [3]string{strings.ToLower(toy1), strings.ToLower(toy2), strings.ToLower(toy3)}

	testToyList := toyList

	fmt.Printf("Второй массив с игрушками: %v\n", testToyList)

	testToyList[1] = toy4

	fmt.Printf("Первый массив с игрушками: %v\n", toyList)
	fmt.Printf("Второй массив с игрушками после изменения: %v\n", testToyList)
}
