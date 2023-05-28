package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := "retert 232 f  45 dddf"

	for _, v := range a {
		_, err := strconv.Atoi(string(v))
		if err == nil {
			continue

		}

		fmt.Println(string(v))
	}
}

// я думаю надо сделать так:
// 1 привести строку к типу стринг,
// 2 завернуть в массив
// 3 с помощью range перебрать по словам.
// 4 если слово является интом удалить.
// 5 вывести новый срез
