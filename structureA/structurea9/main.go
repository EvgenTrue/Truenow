//Создайте структуру Blog, имеющую поле Posts типа []string. Напишите метод AddPost(post string), 
//который добавляет новый пост в блог, и метод GetPosts() []string, 
//который возвращает все посты из блога.
package main

import "fmt"

type Blog struct{
	Post []string
}

func (b *Blog)AddPost(new string){
	b.Post=append(b.Post, new)
}
func (b *Blog)GetPosts()[]string{
	return b.Post
}
func main(){
	a:=Blog{}
	a.AddPost("eeee")
	a.AddPost("rfgrf")
	a.GetPosts()
 fmt.Println(a.GetPosts())

}