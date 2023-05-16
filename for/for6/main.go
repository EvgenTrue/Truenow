package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("vvedi chislo:")
	fmt.Scan(&a)
	fmt.Print("vvedi chislo:")
	fmt.Scan(&b)
	for i := b; i >= a; i-- {

		fmt.Println(i)
	}
}
//вводим число 5 ,10 . начинаем с 10 т.к i=b . 10>5   выводим 10.
//т.к i-- след число 9 . 9>5. выводим 9. 
// 8 > 5 выводим 8
// 7 > 5 выводим 7
// 6> 5 выводим 6
// 5>=5 выводим 5
//
