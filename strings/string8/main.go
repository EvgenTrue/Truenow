package main

import (
	"fmt"
	"net"
)

func main() {

	a := "95.220.197.1"
	address := net.ParseIP(a)
	if address != nil {
		fmt.Println("Правильный", a)
	} else {
		fmt.Println("Неправильный", a)
	}

}
