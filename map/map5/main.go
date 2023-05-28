
//Напишите программу, которая объединяет два map типа string-int в один,
// при этом суммируя значения для общих ключей.

package main

import "fmt"
func main(){
	a:=make(map[string]int, 10)
	b:=make(map[string]int, 10)

	a["r"] = 15
	a["g"] = 14
	a["j"] = 13
	a["l"] = 18
	a["m"] = 1

	b["q"] = 1
	b["w"] = 4
	b["r"] = 3
	b["l"] = 5
	b["m"] = 1
	for k, v:=range a{
	   
	   b[k]+=v
	   fmt.Println(b)

}
}