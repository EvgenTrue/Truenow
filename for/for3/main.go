package main
import "fmt"
func main(){
	sum:=0
	for i:=0;i<=100;i++{
		sum+=i
		fmt.Println(sum)
	}
}

// первая итерация начинаем с 0 вывод 0
/// вторая итерация 1 суммируем с предыдущей 0 получаем 1. вывод 1
/// третья итерация 2 суммируем с предыдущими 0 и 1 . вывод 3
/// четвертая итерация 3 суммируем с предыдущими 0 1 2. вывод 6
/// пятая итерация 4 суммируеем с предыдущим 0 1 2 3 . вывод 10
//...