// Метод "CalculateDiscount" рассчитывает скидку на товар.
func (dp DiscountedProduct) CalculateDiscount() float64 {
	return dp.Price * dp.DiscountRate * float64(dp.Quantity)
  }
  