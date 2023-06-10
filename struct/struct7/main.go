//Создайте структуру "Car" с полями "Brand" (строка) и "Engine" (структура с полями "Horsepower"
//и "FuelType"). Напишите функцию "PrintCarDetails", которая выводит на экран марку автомобиля, 
//мощность двигателя и тип топлива.

package main

import "fmt"


type Car struct{
	Brand string
	Engine Engine
}
type Engine struct{
	Horsepower int
	FuelType string
}
 	func (c Car)PrintCarDetails()string{
		
	return fmt.Sprintf("Марка автомобиля:%s, лошадиные силы %d, топливо %s", c.Brand, c.Engine.Horsepower,c.Engine.FuelType)
}


func main(){
	
	}
	