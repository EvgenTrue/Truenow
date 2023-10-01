//Рассмотрим структуру Dictionary, представляющую словарь с переводами, со следующими полями:
//Words (мап) - хранит переводы слов, где ключами являются исходные слова, а значениями - их переводы.
//Задача: Напишите метод Translate(word string) string, который возвращает перевод указанного слова.
//Если перевод отсутствует, метод должен вернуть пустую строку.

package main

import (
	"fmt"
	
)

type Dictionary struct{
	Words map[string]string
}
func (d *Dictionary)Translate(word string)string{
	for key:=range d.Words{
		return  
	}
	return ""
}
func main(){
	   
	  m:=Dictionary{}
	  m:=make(map[string]string)
	  m["aplle"] = "яблоко"
	  m["hello"] = "привет"
	fmt.Println(m["hello"])
}
