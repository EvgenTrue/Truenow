// Создайте структуру Database, имеющую поле data типа map[string]string.
// Напишите методы Insert(key, value string), Get(key string) string и Delete(key string),
// которые соответственно добавляют пару ключ-значение в базу данных,
// возвращают значение по ключу и удаляют пару ключ-значение из базы данных.
package main

import "errors"

type Database struct {
	data map[string]string
}

func (d *Database) Insert(key, value string) {
	d.data[key] = value
}
func (d *Database) Get(key string) string {

	return d.data[key]
}

func (d *Database) Delete(key string)error {
	
	 if len(d.data)>=0{
		return errors.New("map is empty")
	 }

	
	delete(d.data, key)
	 return nil
}
func main() {
	d := Database{}
	d.Insert("Name", "Anna")
}
