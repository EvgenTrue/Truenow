//Создайте структуру Stack, представляющую стек целых чисел. Напишите методы Push() и Pop() для добавления
// и удаления элементов из стека. Используйте слайс для хранения элементов стека.

package main

import (
	"fmt"

 
)

type Stack struct{
	Stack []int
}
func (s *Stack)Push(new int){
	s.Stack=append(s.Stack, new)	 
}

func (s *Stack)Pop(){
	 s.Stack=s.Stack[0: len(s.Stack)-1]
}

func main(){

}