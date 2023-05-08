package main
import "fmt"

func main(){
var a, b int
fmt.Println("введите число:")
fmt.Scan(&a)
fmt.Scan(&b)
for i:=1; i<=a; i++{
	for j:=1; j<=i; j++{
		fmt.Print(b)
		b++	
	}
	fmt.Println("")}
	 
	for i:=1; i<=a; i++{
		for j:=1; j>=i; j++{
			fmt.Print(b)
		}
		fmt.Println("")}
	}
