// Создайте структуру "Person" с полями "Name" (строка) и "Age" (целое число).
// Напишите функцию "NewPerson", которая принимает имя и возраст в качестве аргументов
// и возвращает новый экземпляр структуры "Person".
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}

}
func main() {
	fedor := NewPerson("Fedya", 15)
	nastya := NewPerson("Nastya", 13)
	fmt.Println(fedor)
	fmt.Println(nastya)
}
