// Задача: Создание структуры, представляющей информацию о ресторане,
// включая название, адрес, список блюд и метод для добавления нового блюда.
package main

import "fmt"

type Dish struct {
  Name  string
  Price float64
}

type Restaurant struct {
  Name    string
  Address string
  Menu    []Dish
}

func (r *Restaurant) AddDish(dish Dish) {
  r.Menu = append(r.Menu, dish)
}
func main(){
	dish:=Dish{Name: "salad", Price: 34.21}
	dish1:=Dish{Name: "myaso", Price: 34.34}
	restaurant:=Restaurant{Name: "drova", Address: "stop street, 12", Menu: []Dish{dish}}
	restaurant.AddDish(dish1)
	fmt.Println(restaurant)

}