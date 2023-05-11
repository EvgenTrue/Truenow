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
	for i := 0; i <= 5; i++ {
		j := rand.Intn(max-min) + min
		k := rand.Intn(max-min) + min
		a = append(a, j)
		b = append(b, k)
		fmt.Println(a, b)
	}
	c := make([]int, 0, 5)
	for i, v := range a {
		c = append(c, v+b[i])
	}
	fmt.Println(c)

	// i=0 j=13, a=13 , b=13  a(13,13)
	// i=1 j=99, a=13, 13, 99 ; b=13, 99 a(13,13, 99, 13,99)
	// i=2 j=55, a=13,13,99,13,99,55 ; b=13,99,55  a(13,13,99,13,99, 55,13,99,55)

}
