//Напишите программу, которая принимает на вход map типа string-int
// и удаляет все элементы с четными значениями.

package main

import "fmt"

func main() {

	a := make(map[string]int, 5)
	a["r"] = 15
	a["g"] = 14
	a["j"] = 13
	a["l"] = 18
	a["m"] = 1
	for k, v := range a {
		if v%2 != 0 {
			delete(a, k)
			fmt.Println(v)
		}
	}
}
