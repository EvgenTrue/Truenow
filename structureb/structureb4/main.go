//Рассмотрим структуру TodoList, представляющую список задач, со следующими полями:
//Tasks (мап) - хранит задачи, где ключами являются идентификаторы задач, а значениями - их описания.
//Задача: Напишите метод CompleteTask(taskID int), который помечает задачу с указанным идентификатором 
//как выполненную.

package main

import "fmt"

type TodoList struct{
	Tasks map[int]string
}
func (t *TodoList)CompleteTask(taskID int){
 for k:=range t.Tasks{
	
	
fmt.Println(k)
 }
}
