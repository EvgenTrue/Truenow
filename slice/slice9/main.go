package main
import (
	"fmt"
	"strings"
)
func main(){
	a:=make([]string, 0, 10)
	a=append(a, "hello", "moon", "box", "kenny", "goat")
	for _,v:= range a {
		fmt.Println(strings.Count(v, "o"))
	}
}