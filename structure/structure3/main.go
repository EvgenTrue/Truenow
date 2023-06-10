// Задача: Учет заказов в интернет-магазине
//
// Создайте структуру Order, которая будет представлять информацию о заказе
// (например, номер заказа, список товаров, статус и т.д.).
// Создайте структуру OnlineStore, которая будет содержать слайс или мапу заказов.
// Добавьте методы к структуре OnlineStore для создания нового заказа,
// обновления статуса заказа и получения списка всех заказов.
//
// Надо найти 3 ошибки
package main

import "fmt"

type Order struct {
	OrderNumber int
	Items       []string
	Status      string
}

type OnlineStore struct {
	Orders []Order
}

func (os *OnlineStore) CreateOrder(order Order) {
	os.Orders = append(os.Orders, order)
}

func (os *OnlineStore) UpdateOrderStatus(orderNumber int, status string) {
	for i, _ := range os.Orders {
		if os.Orders[i].OrderNumber == orderNumber {
			os.Orders[i].Status = status
			break
		}
	}
}

func (os *OnlineStore) GetAllOrders(Order string) {
	allbook:=make([]Order,0,len(os.Orders))
	for_,v:=range os.Orders{
		fmt.Println(allallbook)
		return  
	}
	
}

func main() {
	onlineStore := OnlineStore{}

	order1 := Order{OrderNumber: 1, Items: []string{"Item 1", "Item 2"}, Status: "Pending"}
	order2 := Order{OrderNumber: 2, Items: []string{"Item 3", "Item 4"}, Status: "Shipped"}

	onlineStore.CreateOrder(order1)
	onlineStore.CreateOrder(order2)

	fmt.Println(onlineStore.Orders)

	onlineStore.UpdateOrderStatus(1, "Completed")

	fmt.Println(onlineStore.Orders)

	allOrders := onlineStore.GetAllOrders()
	fmt.Println(allOrders)
}
