package main

import "fmt"

func main() {

	a := make([]string, 0, 5)
	a = append(a, "hello", "man", "a", "kenny", "carhart")
	for _, v := range a {
		for i := 0; i < len(v); i++ {

			fmt.Print(len(v))
		}
		fmt.Print(" ")
	}
}
