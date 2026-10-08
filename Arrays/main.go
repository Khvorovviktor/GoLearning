package main

import "fmt"

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
}
