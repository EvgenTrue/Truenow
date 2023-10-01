//Рассмотрим структуру BankAccount, представляющую банковский счет, со следующими полями:
//Transactions (слайс) - хранит информацию о транзакциях на счете.
//Задача: Напишите метод GetBalance() float64, который возвращает текущий баланс счета на основе информации о транзакциях.

package main

import "fmt"

type BankAccount struct{
	Transactions []float64
}
 
func (b *BankAccount)GetBalance()float64{
for _,v:=range b.Transactions{	
return +v 
}
}

func main(){

	a:= BankAccount{}
	
	fmt.Println(a.GetBalance)
}

