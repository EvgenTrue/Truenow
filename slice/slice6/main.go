package main

import "fmt"

func main() {

	a := make([]string, 0, 5)
	a = append(a, "hello", "man", "a", "kenny", "carhart")
	for _, v := range a {
		{

			fmt.Print(len(v))
		}
		fmt.Print(" ")
	}
}
