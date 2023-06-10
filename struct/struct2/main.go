//Расширьте структуру "Person" методом "SayHello", который выводит на экран приветствие
// с именем
//и возрастом человека. Пример вывода: "Привет, меня зовут Иван, мне 25 лет."

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() string {
	return fmt.Sprintf("Привет, меня зовут %s, мне %d лет", p.Name, p.Age)
}

func (p Person) SSS() (string, int) {
	return p.Name, p.Age
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}

}
func main() {
	fedor := NewPerson("Fedya", 15)
	nastya := NewPerson("Nastya", 34)
	fmt.Println(fedor.SayHello())
	fmt.Println(nastya.SayHello())
	a, b := fedor.SSS()
	fmt.Println(a, b)
}
