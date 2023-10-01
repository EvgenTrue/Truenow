// Задача: Учет студентов в университете
//
// Создайте структуру Student, которая будет содержать информацию о студенте
// (например, имя, возраст, специальность и т.д.).
// Создайте структуру University, которая будет содержать слайс или мапу студентов.
// Добавьте методы к структуре University для добавления и удаления студентов,
// а также для поиска студентов по определенным критериям (например, по специальности или возрасту).
//
// Надо найти 3 ошибки

package main

import "fmt"

type Student struct {
	Name      string
	Age       int
	Specialty string
}

type University struct {
	Students map[string]Student
}

func (u *University) AddStudent(id string, student Student) {
	u.Students[id] = student
}

func (u *University) RemoveStudent(id string) {
	delete(u.Students, id)
}

func (u *University) FindStudentsBySpecialty(specialty string) []Student {
	var result []Student
	for _, student := range u.Students {
		if student.Specialty == specialty{
		result = append(result, student)	
		} 
	}
	return result
}

func main() {
	university := University{
		Students: make(map[string]Student),
	}

	student1 := Student{Name: "John Doe", Age: 20, Specialty: "Computer Science"}
	student2 := Student{Name: "Jane Smith", Age: 21, Specialty: "Mathematics"}

	university.AddStudent("001", student1)
	university.AddStudent("002", student2)

	fmt.Println(university.Students)

	university.RemoveStudent("001")

	fmt.Println(university.Students)

	studentsBySpecialty := university.FindStudentsBySpecialty("Mathematics")
	fmt.Println(studentsBySpecialty)
}