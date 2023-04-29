package main
import "fmt"

func main() {
	var n, m, sum int
	fmt.Println("введите число:")
	fmt.Scan(&n)
    fmt.Scan(&m)
	sum=0
	for i:=n; i<=m; i++{
		sum=sum+i
	}
	fmt.Println(sum)
}