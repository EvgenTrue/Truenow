package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "kolombo"
	result := strings.HasPrefix(a, "ko")
	fmt.Println(result)
}
