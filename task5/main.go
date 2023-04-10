package main
 
import "fmt"
 
func main() {

var a = 1
var b int
fmt.Scan(&b)
for i:=1; i<=b; i++ {
	a= i * a
}
fmt.Println(a)
}