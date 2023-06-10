package main

import (
	"fmt"
	"strconv"
	 
)

func main() {

	a := "sdad j 2132 grfrf 4 423 frvkjh"

	for _, v := range a {

		b, err := strconv.Atoi(i,int(string(v)))
		if err != nil {
			continue

		}
		fmt.Println(b)

	}
}
