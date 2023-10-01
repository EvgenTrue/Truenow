// Задача: Создание структуры, представляющей информацию о компании,
// включая список сотрудников и метод для добавления нового сотрудника.
package main

import "fmt"

type Employee struct {
  Name     string
  Age      int
  Position string
}

type Company struct {
  Name      string
  Location  string
  Employees []Employee
}
func NewCompany(name, location string)*Company{
	return &Company{Name: name, Location: location}


}

func NewEmployee(name, position string, age int)*Employee{
	return &Employee{Name: name, Age: age, Position: position }
}


func (c *Company) AddEmployee(employee Employee) {
  c.Employees = append(c.Employees, employee)
}
func main(){
	company:=NewCompany("coca-cola", "kipr")
	employee:=NewEmployee("alesha", "taxist", 48)
	employee2:=NewEmployee("vitya", "vodoprod", 55)
	company.AddEmployee(*employee)
	company.AddEmployee(*employee2)
	fmt.Println(company)
}