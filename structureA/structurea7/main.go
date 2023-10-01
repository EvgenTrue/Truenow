//Создайте структуру Queue, представляющую очередь строк.
//Напишите методы Enqueue(value string) и Dequeue() string для добавления и удаления элементов из очереди.
//Используйте слайс для хранения элементов очереди.
package main

import "fmt"

type Queue struct {
	Queue []string
}

func (q *Queue) Enqueue(value string) {
	q.Queue = append(q.Queue, value)
}
func (q *Queue) Dequeue() string {
	if len(q.Queue) > 0 {
		a:=q.Queue[0]
		q.Queue = q.Queue[1:]
		return a 
	}
	return ""
}

func main() {

}
