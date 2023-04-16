package main

import "fmt"

func main() {

	a := "ahba jdocd"

	s := []string{"ab", "ba", "cd", "dc"}

	for i := 0; i <= len(a); i++ {
		for _, j := range s {
			if string(a[i]) == j{
			fmt.Println(i, string(a[i]))
			break

		}
	}

}
}