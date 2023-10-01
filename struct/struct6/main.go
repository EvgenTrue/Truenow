//Создайте структуру "Rectangle" с полями "Width" (целое число) и "Height" (целое число). 
//Напишите метод "Area", который принимает указатель на структуру "Rectangle" и
//возвращает площадь прямоугольника (умножение ширины на высоту).

package main

import "fmt"


type Rectangle struct{
	Width int
	Height int
}

func (r Rectangle) Area()int{
	s:=r.Width*r.Height
	return s
}



func main(){
	s:=Rectangle{Width: 10, Height: 5}
	fmt.Println(s)
}
