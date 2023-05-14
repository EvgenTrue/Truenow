package main

import (
	"fmt"

)
func main(){
	var n int
	sum:=3
	fmt.Print("введи число:")
	fmt.Scan(&n)
	for i:=0; i<=n; i++{
		sum=(n-1)+(n-2)
		fmt.Println(sum)
	}
}