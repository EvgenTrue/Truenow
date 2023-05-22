package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){

	a:= "sdad j 2132 grfrf 4 423 frvkjh"
	
	for _, v:=range a{
	}
	i, err := strconv.Atoi(a)
	if err!= nil {
		fmt.Println(a)
	
	}
	 ss := strings.Split(a, "")
     fmt.Println(ss,i)
	 
	}
	
