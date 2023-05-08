package main

import (
	"fmt"
	"math/rand"
)

func main() {

	a := make([]int, 0, 5)
	b := make([]int, 0, 5)
	max := 100
	min := 0
	for i := 0; i < 5; i++ {
		j := rand.Intn(max-min) + min
		a = append(a, j)
		b = append(b, j)
	

		for _, j := range b{
			a = append(a, j)
		}
	}
		fmt.Println(a)
	

}
