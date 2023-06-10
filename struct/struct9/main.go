//1. Определите структуру "Книга" с полями "Название", "Автор" и "Год издания".


//Напишите структуру "Библиотека", которая содержит следующие элементы:
//* Поле "книги" типа слайс, представляющее коллекцию книг.
//* Функцию "ПолучитьКнигиПозднееГода", которая принимает год в качестве аргумента и выводит все книги изданные позже этого года.
//3. Создайте структуру "Журнал" с полями "Название", "Год издания" и "Дата выхода".
package main

import "fmt"

type Paper interface{
	Getdate()int
} 
func (b book)Getdate()int{
	return b.date
}
func (j jornal)Getdate()int{
	return j.Date
}

type book struct{
	name string
	author string
	date int
}
type jornal struct{
	Name string
	Date int
	Pdate int

}

type biblioteka struct{
	Books []Paper 

}

func (b biblioteka)AddBook(a Paper){
 b.Books=append(b.Books, a)
}

func (b biblioteka)GetBooks(date int)[]Paper{
	mybook:=make([]Paper,0, len(b.Books))
	for _,v:=range b.Books{
		if v.Getdate()>=date{
			mybook = append(mybook, v)
			
		}
	}
	 
	return mybook
}

func main(){

}