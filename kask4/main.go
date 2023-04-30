package main
import "fmt"
func main(){

	var a   int
	fmt.Print("введите число:")
	fmt.Scan(&a)
	for a>0{
		fmt.Print(a%10)
		a/=10
	}
	

	
} 
//a/1000/100+a/1000/100/10+a%10

//a/1000

func findDigitSum(num int) int {
	res := 0
	for num>0 {
	res += num % 10
	num /= 10
	}
	return res
   }
   
   func main(){
	fmt.Println(findDigitSum(168))
   }