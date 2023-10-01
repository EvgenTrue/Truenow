//Создайте структуру Calculator, имеющую поле result типа float64. Напишите метод Add(),
//который принимает два числа и добавляет их к result. Затем напишите метод GetResult(),
// который возвращает текущее значение result.

package main

import "fmt"

type Calculator struct {
	result float64
}

func (c *Calculator) Add(a float64, b float64) {
	c.result += a+b
	 
}
func (c *Calculator) GetResult() float64 {
	return c.result
}
func main() {
	calc:=Calculator{}
	calc.Add(24, 6)
	
	fmt.Println(calc.GetResult())
}
