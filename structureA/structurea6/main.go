//Создайте структуру Cache, имеющую поле data типа map[string]interface{}.
// Напишите метод Set(key string, value interface{}), который сохраняет значение по ключу в кэше,
// и метод Get(key string) interface{}, который возвращает значение по ключу из кэша.
//Используйте interface{} для работы с различными типами значений.
package main

import "fmt"

type Cache struct {
	data map[string]interface{}
}

func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *Cache) Get(key string) interface{} {
	return c.data[key]
}
func main() {
	m := Cache{data: make(map[string]interface{})}
	m.Set("a", "b")
	fmt.Println(m.Get("a"))
}
