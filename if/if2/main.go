//Задача "Определение времени суток": Напишите программу, которая запрашивает текущий час (число от 0 до 23)
//и выводит сообщение "Доброе утро", "Добрый день",
//"Добрый вечер" или "Доброй ночи" в зависимости от времени суток.

package main

import "fmt"

func main() {
	a := 0
	fmt.Scan(&a)
	if a >= 21 && a < 5 {
		fmt.Println("Доброй ночи")
	}
	if a >= 5 && a < 10 {
		fmt.Println("Доброе утро")
	}
	if a >= 10 && a < 18 {
		fmt.Println("Добрый день")
	}
	if a >= 18 && a < 21 {
		fmt.Println("Добрый вечер")
	}
}
