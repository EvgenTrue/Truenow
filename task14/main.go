package main

import "fmt"

func main() {


	s := []string{"a", "b", "c", "d"}

	for i:=0; i<4; i++{
		for j:=0; j<4; j++{
			for k:=0; k<4; k++{
				for p:=0; p<4; p++{
					fmt.Println(s[i]+s[j]+s[k]+s[p])	
					}	
			}
			
		}
	
	}
	
}

/// abcd  abdc as
/// abcd aabc aaab aaaa
/// bacc bbcd bbbd 
// aaaa bbbb
// aaab aaac 
// aadc aadc
// acdc

/// aaaa aaab aaac aaad
/// aaba aabb aabc aabd
/// aaca aacb aacc aacd
//  aada aadb aadc aadd
//  abaa abab