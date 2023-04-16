package main

import "fmt"

func main(){
	var a rune
	fmt.Print("введите число:")
	fmt.Scan(&a)
	fmt.Printf("%b", a)

}
