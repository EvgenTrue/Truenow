package main
import "fmt"

func main(){
	var a int
	var c = 0
	fmt.Scan(&a)
	for i:=0 ; i<=a ; i++{
     c = i + c
	}
	fmt.Println(c)
}
