package main
import (
	"fmt"
	"strings"
)
func main(){
	a:=make([]string, 0, 10)
	a=append(a, "hello", "moon", "box", "kenny", "goat")
	for i:= range a {
		fmt.Println(strings.Count(a[i], "o"))
	}
}