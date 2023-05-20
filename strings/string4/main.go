package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "hello world"
	result:=strings.Replace(a, "world", "clown", -1)
	fmt.Println(result)
}
