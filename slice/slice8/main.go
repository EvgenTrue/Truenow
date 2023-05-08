package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main(){
	a:= make([]int,0, 10)
	max:=100
	min:=0
	for i:=0; i<10;i++{
		j:=rand.Intn(max-min)-min
		a=append(a, j)
		
	}
	fmt.Println(a)
	for i:=len(a)-1; i>0;i++{
	if a[i]%2==0{
		sort.Ints(a)
		a=append(a[:9], a[9:]...) 
		fmt.Println(a)
	}
	
} 

}