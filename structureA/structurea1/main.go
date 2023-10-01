//Создайте структуру Person, имеющую поля Name и Age. Напишите метод IsAdult(),
//который возвращает true, если возраст больше или равен 18, и false в противном случае.

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) IsAdult(Age int) bool {
	 
		if p.Age >= 18 {
			return true
		}
		return false
	}
	


func main() {
	
	fmt.Println()
}
