package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
 
	a := "retert 232 f  45 dddf"
	b := strconv.Itoa(a)
	for _,v:=range a{
		
			fmt.Println(string, v)
		}

	}
	


// я думаю надо сделать так:
// 1 привести строку к типу стринг,
// 2 завернуть в массив 
// 3 с помощью range перебрать по словам.
// 4 если слово является интом удалить.
// 5 вывести новый срез
 