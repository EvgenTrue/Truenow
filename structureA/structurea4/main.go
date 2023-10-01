// Создайте структуру Counter, имеющую поле count типа int. Напишите метод Increment(),
// который увеличивает значение count на 1, и метод GetValue(),
//который возвращает текущее значение count. Используйте замыкание для доступа к полю count.

package main

import "fmt"

type Counter struct {
	Count int
}

func (c *Counter) Increment() {
	c.Count = c.Count + 1

}
func (c *Counter) GetValue() int {
	return c.Count
}

func main() {
	cou := Counter{}
	cou.Increment()
	fmt.Println(cou.GetValue())
}
