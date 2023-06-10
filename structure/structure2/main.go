// Задача: Учет сотрудников в компании
//
// Создайте структуру Employee, которая будет содержать информацию о сотруднике
// (например, имя, возраст, должность и т.д.).
// Создайте структуру Company, которая будет содержать мапу сотрудников,
// где ключом будет идентификатор сотрудника (например, уникальный номер или имя).
// Добавьте методы к структуре Company для добавления, удаления и поиска сотрудников,
// а также для получения списка всех сотрудников.
//
//
// Надо найти 3 ошибки

package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Position string
}

type Company struct {
	Employees map[string]Employee
}

func (c *Company) AddEmployee(id string, employee Employee) {
	c.Employees[id] = employee
}

func (c *Company) RemoveEmployee(id string) {
	delete(c.Employees, id)
}

func (c *Company) GetEmployee(id string) (Employee, bool) {
	employee, exists := c.Employees[id]
	return employee, exists
}

func (c *Company) GetAllEmployees() []Employee {
	var employees []Employee
	for _, employee := range c.Employees {
		employees = append(employees, employee)
	}
	return employees
}

func main() {
	company := Company{}

	employee1 := Employee{Name: "John Doe", Age: 30, Position: "Manager"}
	employee2 := Employee{Name: "Jane Smith", Age: 25, Position: "Developer"}

	company.AddEmployee("001", employee1)
	company.AddEmployee("002", employee2)

	fmt.Println(company.Employees)

	company.RemoveEmployee("001")

	fmt.Println(company.Employees)

	employee, exists := company.GetEmployee("002")
	if exists {
		fmt.Println(employee)
	}

	allEmployees := company.GetAllEmployees()
	fmt.Println(allEmployees)
}
