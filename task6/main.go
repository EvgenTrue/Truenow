package main
 
import "fmt"
 
func main() {
    var number, i, j int
    var flag bool
 
    fmt.Print("Введите число для вывода всех простых чисел до данного числа: ")
    
    fmt.Scan(&number)
    for i = 2; i < number; i++ {
        
        flag = false
        for j = 2; j <= i /2; j++ {
            
            if i  %j == 0 {
                flag = true
            }
        }
        if flag == false {
            fmt.Println(i)
        }
    }
 
}