package main

import "fmt"

func main() {
	// 1
	fmt.Println("Задание 1")

	//true
	fmt.Println(5 == 5) //1

	//true
	fmt.Println(10 != 3) //2

	//true
	fmt.Println(7 > 3) //3

	//true
	fmt.Println(15 < 20) //4

	//true
	fmt.Println(8 >= 8) //5

	//false
	fmt.Println(6 <= 4) //6

	//false
	fmt.Println((10 > 5) && (3 < 1)) //7

	//true
	fmt.Println((10 > 5) || (3 < 1)) //8

	//false
	fmt.Println(!(5 == 5)) //9

	//true
	fmt.Println(!(7 < 3)) //10

	//false
	fmt.Println(true && false) //11

	//false
	fmt.Println(false || false) //12

	//true
	fmt.Println(true || false) //13

	//true
	fmt.Println((4+6 == 10) && (9 > 2)) //14

	//true
	fmt.Println((12/3 == 4) || (8 < 5)) //15

	// 2
	fmt.Println("Задание 2")

	var name string
	fmt.Println("Введите ваше имя:")
	fmt.Scanln(&name)

	var age int
	fmt.Println("Введите ваш возраст:")
	fmt.Scanln(&age)

	var hasTicket bool
	fmt.Println("У вас есть билет (true или false):")
	fmt.Scanln(&hasTicket)

	canEnter := (age >= 18) && (hasTicket == true)
	fmt.Printf("Имя: %s, может ли войти: %t\n", name, canEnter)

	// 3
	fmt.Println("Задание 3")

	var namee string
	fmt.Println("Введите ваше имя:")
	fmt.Scanln(&namee)

	var isLoggedIn bool
	fmt.Println("Вы зарегестрированы? (ответ только true или false):")
	fmt.Scanln(&isLoggedIn)

	var isAdmin bool
	fmt.Println("Вы являетесь админом? (ответ только true или false):")
	fmt.Scanln(&isAdmin)

	hasAccess := (isLoggedIn == true) && (isAdmin == true) || (isLoggedIn == true) && (isAdmin == false)
	fmt.Printf("Имя пользователя: %s, имеет ли доступ: %t", namee, hasAccess)
}
