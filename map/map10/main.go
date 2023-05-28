//Напишите функцию, которая принимает на вход slice типа string и возвращает map,
//где ключи - это элементы из slice,
//а значения - булевые значения, указывающие на то, есть ли данный элемент в исходном slice

package main

import "fmt"

func main() {

	s := make([]string, 7, 10)
	s = append(s, "one", "two", "tree", "four", "five", " ")
	m := make(map[string]bool, 10)
	for _,v := range s {
		m[v] = true
		fmt.Println(s, m																												)
	}

}
