package main

import "fmt"

func main() {
	a := "hello"
	for i:=len(a)- 1 ; i>=0; i--{
		b:=a[i]
//	for i,b:= range a {
			fmt.Println(i,string(b))
	}
}
