package main

import "fmt"

func main() {
	// 1
	fmt.Println("Задание 1")

	var temperature int
	fmt.Println("Введите температуру от -20 до 40 градусов:")
	fmt.Scan(&temperature)

	if temperature < 0 {
		fmt.Printf("Температура: %d - это холодно!!\n", temperature)
	} else if temperature >= 0 && temperature <= 20 {
		fmt.Printf("Температура: %d - это тепло!!\n", temperature)
	} else {
		fmt.Printf("Температура: %d - это жарко!!\n", temperature)
	}

	// 2
	fmt.Println("Задание 2")

	var score int
	fmt.Println("Введите любое значение счёта от 0 до 100:")
	fmt.Scan(&score)

	if score >= 90 {
		fmt.Printf("Ваш счёт: %d. Это ОТЛИЧНЫЙ результат!!", score) // ничего страшного, что я изменил слово по падежам?
	} else if score >= 70 && score <= 89 {
		fmt.Printf("Ваш счёт: %d. Это хороший результат!", score)
	} else if score >= 50 && score <= 69 {
		fmt.Printf("Ваш счёт: %d. Это удовлетворительный результат.", score)
	} else {
		fmt.Printf("Ваш счёт: %d. Вы не сдали!\n", score)
	}

	// 3
	fmt.Println("Задание 3")

	var hour int
	fmt.Println("Введите любое значение от 0 до 23:")
	fmt.Scan(&hour)

	switch {
	case hour >= 0 && hour <= 5:
		fmt.Println("Это ночь.")
	case hour >= 6 && hour <= 11:
		fmt.Println("Это утро.")
	case hour >= 12 && hour <= 17:
		fmt.Println("Это день.")
	case hour >= 18 && hour <= 23:
		fmt.Println("Это вечер.")
	default:
		fmt.Println("Вы написали не то время, попробуйте ввести другое!!")
	}

	// 4
	fmt.Println("Задание 4")

	var number int
	fmt.Println("Введите любое целое число:")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Число, которое вы ввели чётное")
	} else {
		fmt.Println("Число, которое вы ввели не чётное")
	}

	// 5
	fmt.Println("Задание 5")

	var day string
	fmt.Println("Введите любой день недели на английском в верхнем регистре:")
	fmt.Scan(&day)

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday":
		fmt.Printf("Вы ввели: %s. Это будний день!\n", day)
	case "Saturday", "Sunday":
		fmt.Printf("Вы ввели: %s. Это выходной день!\n", day)
	default:
		fmt.Printf("Вы ввели: %s. Вы ввели некорректный день, попробуйте ввести другой!!\n", day)
	}

	// 6
	fmt.Println("Задание 6")

	var balance int
	fmt.Println("Введите любой счёт баланса:")
	fmt.Scan(&balance)

	if balance >= 0 {
		fmt.Printf("Баланс, который вы ввели: %d. Это положительный баланс!!!\n", balance)
	} else {
		fmt.Printf("Баланс, который вы ввели: %d. Это отрицательный баланс!!!\n", balance)
	}

	// 7
	fmt.Println("Задание 7")

	var age int
	fmt.Println("Введите ваш возраст:")
	fmt.Scan(&age)

	if age < 13 {
		fmt.Printf("Ваш возраст, который вы ввели: %d. Вы еще ребёнок.\n", age)
	} else if age >= 13 && age <= 17 {
		fmt.Printf("Ваш возраст, который вы ввели: %d. Вы подросток.\n", age)
	} else {
		fmt.Printf("Ваш возраст, который вы ввели: %d. Вы уже взрослый человек!\n", age)
	}

	// 8
	fmt.Println("Задание 8")

	var command string
	fmt.Println("Здравствуйте! Введите любую команду в нижнем регистре:")
	fmt.Scan(&command)

	switch command {
	case "start":
		fmt.Printf("Команда, которую вы ввели: %s! Start.\n", command)
	case "stop":
		fmt.Printf("Команда, которую вы ввели: %s! Stop.\n", command)
	case "restart":
		fmt.Printf("Команда, которую вы ввели: %s! Restart.\n", command)
	default:
		fmt.Printf("Команда, которую вы ввели: %s! Это неизвестная команда, попробуйте другую!\n", command)
	}

	// 9
	fmt.Println("Задание 9")

	var grade int
	fmt.Println("Введите любую оценку от 1 до 5 включительно:")
	fmt.Scan(&grade)

	switch grade {
	case 3:
		fmt.Printf("Вы ввели оценку: %d. Это оценка <<C>>.\n", grade)
	case 5:
		fmt.Printf("Вы ввели оценку: %d. Это оценка <<A>>!\n", grade)
	case 2:
		fmt.Printf("Вы ввели оценку: %d. Это оценка <<D>>.\n", grade)
	case 1:
		fmt.Printf("Вы ввели оценку: %d. Это оценка <<F>>!\n", grade)
	case 4:
		fmt.Printf("Вы ввели оценку: %d. Это оценка <<E>>\n", grade)
	default:
		fmt.Printf("Вы ввели оценку: %d. Это неккоректная оценка, введите другую!\n", grade)
	}
}
