// Задача: Создание структуры, представляющей информацию о книжной библиотеке,
// включая список книг, список авторов и метод для добавления новой книги с автором.

package main

import "fmt"

type Author struct {
  Name  string
  Books []string
}

type Library struct {
  Books   []string
  Authors []Author
}

func (l *Library) AddBookWithAuthor(book string, author string) {
  l.Books = append(l.Books, book)

  for i := range l.Authors {
    if l.Authors[i].Name == author {
      l.Authors[i].Books = append(l.Authors[i].Books, book)
      return
    }
  }

  l.Authors = append(l.Authors, Author{Name: author, Books: []string{book}})
}

func main(){
	author:=Author{Name: "gogol", Books: []string{"nos", "rot"} }
	library:=Library{Books: []string{"nos","rot"}, Authors: []Author{author}}
	library.AddBookWithAuthor("golova", "gogol")
	library.AddBookWithAuthor("derovo", "fet") 
	fmt.Println(library)
}