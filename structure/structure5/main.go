// Задача: Учет товаров на складе
//
// Создайте структуру Product, которая будет представлять информацию о товаре
// (например, название, цена, количество и т.д.).
// Создайте структуру Warehouse, которая будет содержать мапу товаров,
// где ключом будет уникальный идентификатор товара.
// Добавьте методы к структуре Warehouse для добавления и удаления товаров,
// а также для получения общего количества товаров и обновления информации о товаре.
//
// Надо найти 3 ошибки

package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Product struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

type Warehouse struct {
	Products map[string]Product
}

func (w *Warehouse) AddProduct(product Product) {
	w.Products[product.ID] = product
}

func (w *Warehouse) RemoveProduct(id string) {
	delete(w.Products, id)
}

func (w *Warehouse) UpdateProductInfo(id, name, price, quantity string) {
	product, exists := w.Products[id]
	if exists {
		product.Name = name
		product.Price = price
		product.Quantity = quantity
		w.Products[id] = product
	}
}

func (w *Warehouse) GetTotalQuantity() {
	totalQuantity := 0
	for _, product := range w.Products {
		totalQuantity += product.Quantity
	}
	return totalQuantity
}

func main() {
	warehouse := Warehouse{
		Products: make(map[string]Product),
	}

	product1 := Product{ID: "001", Name: "Product 1", Price: 9.99, Quantity: 10}
	product2 := Product{ID: "002", Name: "Product 2", Price: 19.99, Quantity: 5}

	warehouse.AddProduct(product1)
	warehouse.AddProduct(product2)

	fmt.Println(warehouse.Products)

	warehouse.RemoveProduct("001")

	fmt.Println(warehouse.Products)

	warehouse.UpdateProductInfo("002", "New Product Name", "24.99", "8")

	fmt.Println(warehouse.Products)

	totalQuantity := warehouse.GetTotalQuantity(Product)
	fmt.Println("Total quantity:", totalQuantity)
}
