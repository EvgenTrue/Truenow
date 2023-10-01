//Создайте структуру URLShortener, имеющую поле urls типа map[string]string.
//Напишите методы Shorten(url string) string и Expand(shortURL string) string,
//которые соответственно сокращают длинный URL до короткого и расшифровывают короткий URL до оригинального.
//Используйте мапу для хранения соответствий между длинными и короткими URL.

package main

import "fmt"

type URLShortener struct {
	urls map[string]string
}

func (u *URLShortener) Shorten(url string) string {
	 a:="44455"
	 b:="44"
	 a=b
	return  fmt.Println()
}
func (u *URLShortener) Expand(shortURL string) string {
	return u.urls[key]

}
func main() {
adress:=URLShortener{}
adress.Shorten("1.34.6")
fmt.Println(adress.Expand(adress.Shorten()))
}

 