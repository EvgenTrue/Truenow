// Задача: Учет книг в библиотеке
//
// Создайте структуру Book, которая будет представлять информацию о книге
// (например, заголовок, автор, год издания и т.д.).
// Создайте структуру Library, которая будет содержать слайс или мапу книг.
// Добавьте методы к структуре Library для добавления и удаления книг,
// а также для поиска книг по определенным критериям (например, по автору или году издания).
//
// Надо найти 3 ошибки
package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func (l *Library) RemoveBook(title string) {
	for i, book := range l.Books {
		if book.Title == title {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			break
		}
	}
}

func (l *Library) FindBooksByAuthor(author string) []Book {
	result := make([]Book, 0, len(l.Books))
	for _, book := range l.Books {
		if book.Author == author {
			result = append(result, book)
		}
	}
	return result
}

func main() {
	library := Library{}

	book1 := Book{Title: "Book 1", Author: "Author 1", Year: 2021}
	book2 := Book{Title: "Book 2", Author: "Author 2", Year: 2022}

	library.AddBook(book1)
	library.AddBook(book2)

	fmt.Println(library.Books)

	library.RemoveBook("Book 1")

	fmt.Println(library.Books)

	booksByAuthor := library.FindBooksByAuthor("Author 2")
	fmt.Println(booksByAuthor)
}
