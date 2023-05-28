package main

import "fmt"

func main() {
	a := make(map[string]int)
	a["r"] = 15
	a["g"] = 14
	a["j"] = 13
	a["l"] = 18
	a["m"] = 1
	max := 0
	for _, v := range a {
		if v > max {
			max = v

		}

	}
	fmt.Println(max)
}
