package main
import "fmt"
func main(){

	var a   int
	fmt.Print("введите число:")
	fmt.Scan(&a)
	s := []int{5, 4, 3, 2, 1}
for i := len(s)-1; i >= 0; i-- {
   fmt.Print(s[i])
}
}