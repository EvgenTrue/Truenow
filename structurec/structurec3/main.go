//	на интерфейсы: Есть структура которая представляет магазин - Shop, у нее есть метод AddModel который на вход принимает
//
// интерфейс Transport и добавляет машину на склад в магазин, склад это мапа или слайс.
// В магазин можно добавить машину, мотоцикл, мопед и должна быть возможность в будущем добавить любой транспорт.
// Это разные структуры у которых есть поле кол-во колес и мощность.
// В структуре Shop сделать метод который возвращает транспорт у которого колес больше чем n, n передается на вход функции
// и метод который возвращает транспорт с самой большой мощностью.
// И еще тест написать на две функции
package main

import (
	"fmt"
	
)

type Transport interface{
	GetPower()int
	GetWheel()int
	GetName()string
}
type Shop struct {
	Transport map[string]Transport
}
type Avto struct{
	Name string
	Wheel int
	Power int
}
func(a *Avto)GetPower()int{
	return a.Power
}
func(a *Avto)GetWheel()int{
	return a.Wheel
}
func(a *Avto)GetName()string{
	return a.Name
}
type Moto struct{
	Name string
	Wheel int
	Power int
}
func(m *Moto)GetPower()int{
	return m.Power
}
func(m *Moto)GetWheel()int{
	return m.Wheel
}
func(m *Moto)GetName()string{
	return m.Name
}
type Scooter struct{
	Name string
	Wheel int
	Power int
}
func (s *Scooter)GetPower()int{
	return s.Power
}
func(s *Scooter)GetWheel()int{
	return s.Wheel
}
func (s *Scooter)GetName()string{
	return s.Name
}

func (s *Shop)AddModel(transport Transport){
s.Transport[transport.GetName()]=transport
}

func (s *Shop)BackWheel(wheel int)[]string{
    result:=make([]string,0,10)
	for _,item:= range s.Transport{
	if item.GetWheel()>=wheel{
		result=append(result, item.GetName())
	}
}
return result
}

func (s *Shop)BackPower()Transport{
     maxPower:=0
	var result Transport
	for _,item:= range s.Transport{
	if item.GetPower()>=maxPower{
		result=item
		maxPower=item.GetPower()
	}
}
return result
}
func main(){
	shop:=Shop{
		Transport:make(map[string]Transport),
}
 item1:=&Scooter{Name: "Suzuki", Wheel:2, Power:150}
 item2:=&Moto{Name: "toyota", Wheel:3, Power:210}
 item3:=&Avto{Name: "infinity", Wheel:4, Power:250}
  
 shop.AddModel(item1)
 shop.AddModel(item2)
 shop.AddModel(item3)  
 fmt.Println(shop.BackPower())
}