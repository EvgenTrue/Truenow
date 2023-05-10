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
		a = append(a, j)
		b = append(b, j)
		fmt.Println(a,b)

		for i,_= range b{
		a = append(a, b...)
		}
		fmt.Println(a[i])
	}
		
	

}
