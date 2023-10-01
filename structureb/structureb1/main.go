//Рассмотрим структуру Inventory, представляющую инвентарь магазина, со следующими полями:
//Items (мап) - хранит информацию о наличии товаров в магазине. Ключами являются названия товаров, 
//значениями - их количество.
//Задача: Напишите метод AddItem(name string, quantity int), который добавляет товары в инвентарь. 
//Если товар уже существует, то количество обновляется. Если товара нет в инвентаре, он добавляется.

package main
import "fmt"

type Inventory struct{
	Items map[string]int
}

func (i *Inventory)AddItem(name string, quantity int){
	 i.Items[name]+=quantity
}
func main(){
	fmt.Println()
}

 