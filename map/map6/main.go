// Напишите функцию, которая принимает на вход map типа string-int и значение типа int.
// Функция должна удалить все элементы map,значения которых меньше заданного значения.
package main

import (
	"fmt"

	
)

func main() {
	var q int
	a := make(map[string]int, 10)
	a["r"] = 15
	a["g"] = 14
	a["j"] = 13
	a["l"] = 18
	a["m"] = 1
	fmt.Sprintln("vvedi chislo:")
	fmt.Scan(&q)

	for k, v := range a {
		if v > q {
			delete(a, k) 
			fmt.Println(v)
		}
	}

}
