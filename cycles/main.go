package main

import "fmt"

func main() {
	// 1
	fmt.Println("Задание 1")

	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}

	// 2
	fmt.Println("Задание 2")

	sum := 0

	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

	// 3
	fmt.Println("Задание 3")

	var number int
	fmt.Println("Введите любое число:")
	fmt.Scan(&number)

	for i := 1; i <= 10; i++ {
		fmt.Print(number*i, " ")
	}

	// 4
	fmt.Println("Задание 4")

	var n int
	fmt.Println("Введите любое число:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Printf("Вы ввели число: %d, числа, которые делятся без остатка на 3 до этого числа: %d\n", n, i)
		}
	}

	// 5
	fmt.Println("Задание 5")

	var numberr int
	fmt.Println("Введите любое число:")
	fmt.Scan(&numberr)

	count := 0

	for number != 0 {
		number /= 10
		count++
	}

	fmt.Println("Количество цифр:", count)

	// 6
	fmt.Println("Задание 6")

	var text string
	fmt.Println("Введите любой текст:")
	fmt.Scan(&text)

	for index, value := range text {
		fmt.Printf("Index: %d, value: %c\n", index, value)
	}

	// 7

	var balance int
	fmt.Println("Введите любой баланс:")
	fmt.Scan(&balance)

	for {
		var n int
		fmt.Println("Введите любую цифру от 0 до 3 включительно:")
		fmt.Scan(&n)
		if n == 0 {
			fmt.Println("Выход из программы")
			break
		} else if n == 1 {
			fmt.Printf("Ваш баланс равен: %d\n", balance)
		} else if n == 2 {
			balance += 500
			fmt.Printf("увеличение баланса на 500: %d\n", balance)
		} else if n == 3 {
			balance -= 200
			fmt.Printf("уменьшение баланса на 200: %d\n", balance)
		} else {
			fmt.Println("Вы ввели не ту цифру, попробуйте еще раз")
		}
	}
}
