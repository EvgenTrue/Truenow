//Напишите функцию, которая принимает на вход map типа string-int
// и возвращает отсортированный по значениям map

package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[string]int, 5)
	m["one"] = 5
	m["two"] = 1
	m["three"] = 77
	m["four"] = 45
	m["five"] = 52

	keys := make([]string, 0, len(m))
	for k:= range m {
		keys = append(keys, k)
		fmt.Println(m)
		fmt.Println(keys)

		sort.SliceStable(keys, func(i, j int) bool {
			return m[keys[i]] > m[keys[j]]
		})
	}
	fmt.Println(keys)
}
