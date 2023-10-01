// Необходимо спроектировать архитектуру ресторана, которая позволит вести учет смен
//персонала. В структуре "Ресторан" должны быть методы, позволяющие назначить на смену
//как повара, так и официанта. Однако, следует обратить внимание, что официанта
//нельзя назначить на роль повара, так как их обязанности существенно различаются,
// но повара можно назначать на обе роли.Каждый раз, когда сотрудник выходит на смену
//или принимает заказ, в специальный журнал ресторана записывается информация о событии,
//указывая, кто вышел на смену, кто принял заказ, а также, при необходимости,
//кто будет готовить соответствующее блюдо.

package main

import "fmt"

type Staff interface {
	GetCook() string
	GetStaff() string
	GetEat() string
}

type Restorant struct {
	Staff map[int]Staff
}
type Waiter struct {
	Name string
}

func (w *Waiter) GetStaff() string {
	return w.Name
}

type Chief struct {
	Name string
}

func (c *Chief) GetStaff() string {
	return c.Name
}
func (c *Chief) GetCook() string {
	return c.Name
}

type Menu struct {
	Name string
	Eat map[string]Eat
}

func (r *Restorant) AddCook(staff Staff) {
	r.Staff[staff.GetCook()] = staff
}
func (r *Restorant) AddStaff(staff Staff) {
	r.Staff[staff.GetStaff()] = staff
}
func (r *Restorant) AddStaff(staff Staff) {
	r.Staff[staff.GetEat()] = eat
}
func (z *Zakaz) AddZakaz(Eat string) []string {
	result := make([]string, 0, 10)
	for _, item := range menu {
		result = append(result, item.GetEat())

	}
	return z.Eat
}

func main() {
	restorant:=Restorant{
		Staff: make(map[int]Staff),
	}
	chief1:=&Chief{Name: "Arbuzov"}
	chief2:=&Chief{Name: "Baragozov"}
	waiter1:=&Waiter{Name: "Kartoshkin"}
	product1:=&Menu{}
	menu:=Menu{Eat: []Menu{product1}}
	fmt.Println()
}
