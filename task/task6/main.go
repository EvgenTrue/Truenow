package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)
	for i := 4; i <= a; i++ {
		for j := 2; j <= i; j++ {
			if i%j == 0 && i != j {
				break

			}
			fmt.Println(i)
		}

	}
}

// i j a=7
// i=4  j=2  i%j =0 ==0 true

// i=5 j=2   i%J !=0 false
// i=5 j=3   i%j  !=0 false
