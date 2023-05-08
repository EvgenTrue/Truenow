package main
import "fmt"

func main(){

	var a int
	fmt.Println("введите число:")
	fmt.Scan(&a)

	for i:=1; i<=a; i++{
		
		for j:=1; j<=a/2-i; j++{
			fmt.Print(" ")
		}
		
		for j:=1; j>=a/2-i; j++{
			fmt.Print(" ")

		}
		
		for k:=1; k<=i*2-1; k++{
			
			fmt.Print("*")
		}
		if i>a/2{
			k--
		}
			
		fmt.Println("")
		 

		}
	}


