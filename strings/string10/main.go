package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	b := "world"
	fmt.Println(strings.Join([]string{a, b}, "-"))
}
