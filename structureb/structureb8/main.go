// Задача: Создание структуры, представляющей информацию о музыкальном альбоме,
// включая название, год выпуска, список песен и продолжительность альбома.
package main

import "fmt"

type Song struct {
  Title    string
  Duration int
}

type Album struct {
  Title    string
  Year     int
  Songs    []Song
  Duration int
}
func main(){
	song:=Song{Title: "kyky", Duration: 35}
	// song1:=Song{Title: "tik-tak", Duration: 311}
	album:=Album{
		Title: "kykyshka",Year: 1993,Songs: []Song{song},Duration: 35}
	fmt.Println(album)

}