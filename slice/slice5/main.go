package main

import (
	"fmt"
	"math/rand"
)
func main(){
	a:=make([]int,0, 10)
	max:=100
	min:=0
	for i:=0; i<=10;i++{
		j:=rand.Intn(max - min) + min
		a=append(a,j)
	}
	fmt.Println(a)
	for _,v:=range a{
		if v%3 ==0{
			fmt.Println(v)
		}
	}
}