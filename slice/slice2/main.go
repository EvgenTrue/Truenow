package main

import (
	"fmt"
)

func main() {
	a := make([]string, 0, 5)
	a = append(a, "sdfdsf", "hgfhfg", "sdfsdf", "sdfds", "sdfsf")
	for _, v := range a {
		for _, b := range v {
			fmt.Print(string(b), " ")
		}
		fmt.Println(" ")
	}
}
