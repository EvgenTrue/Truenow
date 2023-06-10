//Создайте функцию "AddFriend", которая принимает имя друга в качестве аргумента 
//и добавляет его в срез "Friends". Затем напишите функцию "GetFriendsCount", 
//которая возвращает количество друзей у данного человека.

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
func (p Person) GetFriendsCount() int{
	return len(p.Friends)
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
	fmt.Println(fedor.GetFriendsCount())
	fmt.Println(nastya.GetFriendsCount())

}
