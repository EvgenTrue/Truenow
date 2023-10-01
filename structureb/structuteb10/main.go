// Задача: Создание структуры, представляющей информацию о фильме, включая название,
// год выпуска, список жанров и список актеров.

package main

import "fmt"

type Movie struct {
  Title    string
  Year     int
  Genres   []string
  Actors   []string
  Duration int
}
func main(){
	movie:=Movie{Title: "zhara", Year: 1322, Genres: []string{"comedy", "tragedy"}, Actors: []string{"Van-Damme","Stalone"},Duration: 134}
	fmt.Println(movie)
}