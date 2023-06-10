//Напишите функцию "AddContact", которая принимает тип контакта и значение контакта
//в качестве аргументов и добавляет их в карту "Contacts".
//Затем напишите функцию "GetContact", которая принимает тип контакта
//и возвращает значение соответствующего контакта из карты "Contacts".
package main

import "fmt"

type Person struct {
	Name string
	Age  int

	Friends map[string]Contact
}
type Contact struct {
	Phone string
	Email string
	Name  string
}

func (p Person) SayHello() string {
	return fmt.Sprintf("Привет, меня зовут %s, мне %d лет", p.Name, p.Age)
}
func (p Person) AddContact(friend Contact) {
	p.Friends[friend.Name] = friend
}
func (p Person) GetContact(name string) Contact {
	return p.Friends[name]
}
func NewContact(name, email, phone string) Contact {
	return Contact{
		Phone: phone,
		Email: email,
		Name:  name,
	}
}
func NewPerson(name string, age int) Person {
	return Person{
		Name:    name,
		Age:     age,
		Friends: make(map[string]Contact),
	}

}
func main() {
	fedor := NewPerson("Fedya", 15)
	nastya := NewPerson("Nastya", 34)
	fmt.Println(fedor.SayHello())
	fmt.Println(nastya.SayHello())
	fedor.AddContact(NewContact("Mama", "mama@mail.ru", "8213213213"))
	fedor.AddContact(NewContact("Ded", "ded@mail.ru", "8213214533"))
	fmt.Println(fedor.GetContact("baba")) 
}
