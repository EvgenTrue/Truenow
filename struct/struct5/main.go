//Расширьте структуру "Person", добавив поле "Address" (структура с полями "City" и "Street").
//Напишите функцию "PrintAddress", которая выводит на экран полный адрес человека.

package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address Address
}
type Address struct {
	City   string
	Street string
}

func (p Person) PrintAddress() string {
	return fmt.Sprintf("Живет по адресу:%s,%s", p.Address.City, p.Address.Street)
}

func main() {

}
