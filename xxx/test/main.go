//На вход дается строка, из нее нужно сделать другую строку,
//оставив только нечетные символы (считая с нуля)

package main

import (
	"fmt"
	"strings"
	
)

func main() {
	 
	a:= "ihgewlqlkot" 
	var b string 
    var s string

	 
	for j,v := range a {
		if j%2 != 0 {
			s+=a
			b=string(v)
			 s:=strings.Join([]string a,b)
			fmt.Println(s)

		}
		 
	}
	

 
}
