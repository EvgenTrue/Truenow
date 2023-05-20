package main

import (
	"fmt"
	"strings"
)

func main(){

	a:= "sdad j 2132  4 4grfrf f4 423 frv4f43"
	ss:=strings.Split(a, " ")
	for _,v:=range ss{
		fmt.Println(string(v))
	}
}