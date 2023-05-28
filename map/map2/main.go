package main

import "fmt"

func main() {

	a := make(map[string]int)
	a["r"] = 123
	a["e"] = 434
	for i := range a {
		fmt.Println(i)
	}
}
