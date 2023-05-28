package main

import "fmt"

func main(){

	a:=make([]string, 0, 10)
	a = append(a, "g", "h", "r", "h", "g", "r", "e")
	b:=make(map[string]int, 10)

	 for _,v:=range a{
		
		b[v]++
		fmt.Println(b)
	 }
	
	}
		
	
