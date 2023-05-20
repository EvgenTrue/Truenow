package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "sada  asdsa  asdsa sadsa"
	ss := strings.Split(a, " ")
	fmt.Println(ss)

	for _, v := range ss {
		fmt.Println(len(v))
	}

}
