package main

import "fmt"

// Структура "Product" представляет информацию о товаре.
type Product struct {
  ID       int
  Name     string
  Price    float64
  Quantity int
}

// Метод "CalculateTotalPrice" рассчитывает общую стоимость товара.
func (p Product) CalculateTotalPrice() float64 {
  return p.Price * float64(p.Quantity)
}

// Структура "Order" представляет информацию о заказе.
type Order struct {
  ID       int
  Products []Product
}

// Метод "AddProduct" добавляет товар в заказ.
func (o *Order) AddProduct(product Product) {
  o.Products = append(o.Products, product)
}

// Метод "CalculateOrderTotal" рассчитывает общую стоимость заказа.
func (o Order) CalculateOrderTotal() float64 {
  total := 0.0
  for _, product := range o.Products {
    total += product.CalculateTotalPrice()
  }
  return total
}

// Интерфейс "Discountable" определяет метод для расчета скидки.
type Discountable interface {
  CalculateDiscount() float64
}

// Структура "DiscountedProduct" представляет информацию о товаре со скидкой.
type DiscountedProduct struct {
  Product
  DiscountRate float64
}

// Метод "CalculateDiscount" рассчитывает скидку на товар.
func (dp DiscountedProduct) CalculateDiscount() float64 {
  return dp.Price * dp.DiscountRate * float64(dp.Quantity)
}
func main(){
	product1:=Product{ID: 1, Name: "Product 1", Price:10.00, Quantity: 2}
	product2:=Product{ID: 2, Name: "Product 2", Price: 24.50, Quantity: 1}
	order:=Order{ID: 1,Products: []Product{product1}}
	
	order.AddProduct(product1)
  order.AddProduct(product2)
	fmt.Println(order.CalculateOrderTotal())
	
	product3:=DiscountedProduct{DiscountRate: 4.50, Product: product1}
	
	fmt.Println(product3.CalculateDiscount())
}