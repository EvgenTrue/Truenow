//Расширьте структуру "Person" методом "ChangeName", который принимает новое имя в качестве аргумента
//и изменяет поле "Name" в структуре "Person".

package main
import "fmt"


type Person struct{
	Name string
	Age int
}

func (p Person) ChangeName(nickname string){
	p.Name=nickname
	  
}
func main(){
  fmt.Println()
}