package main

import (
	"fmt"
)

func main() {
	s := make([]string, 6, 10)
	s = append(s, "a", "r", "h", "d", "e", "r")
	m := make(map[string]int, 10)
	for _, v := range s {
		m[v]++
		fmt.Println(m)
	}

}

//Напишите программу, которая создает slice типа string и затем создает map,
//где ключи - это уникальные элементы из slice, а значения - количество их вхождений.
