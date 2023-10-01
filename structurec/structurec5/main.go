// Метод "CalculateOrderTotal" рассчитывает общую стоимость заказа.
func (o Order) CalculateOrderTotal() float64 {
	total := 0.0
	for _, product := range o.Products {
	  total += product.CalculateTotalPrice()
	}
	return total
  }
  