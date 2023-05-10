package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := make([]int, 0, 10)
	max := 100
	min := 0

	for i := 0; i < 10; i++ {
		j := rand.Intn(max-min) - min
		a = append(a, j)

	}
	fmt.Println(a)

	for i := len(a) - 1; i > 0; i++ {
		if a[i]%2 == 0 {

			b := make([]int, 10)
			for _, j := range b {

				b = append(b, j)

				fmt.Println(j)
			}

		}
	}
}
