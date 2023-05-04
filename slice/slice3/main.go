package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	a := make([]int, 0, 10)
	max := 100
	min := 0

	for i := 0; i <= 10; i++ {
		j := rand.Intn(max-min) + min
		a = append(a, j)
	}
	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)
}
