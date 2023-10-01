//Рассмотрим структуру Playlist, представляющую музыкальный плейлист, со следующими полями:
//Songs (слайс) - хранит названия песен в плейлисте.
//Задача: Напишите метод RemoveSong(song string), который удаляет указанную песню из плейлиста, если она присутствует.


package main

import "fmt"

type Playlist struct{
	Song []string
}
func (p *Playlist)RemoveSong(song string){
a:=make([]string,0,len(p.Song))
	for _,v:= range p.Song{
if v==song{
	continue
}
a=append(a, v)
}
p.Song=a

}
func main(){
	q:= Playlist{}
	p.RemoveSong("dfdf")  // как добавить конкретную песню чтобы работалоо
	fmt.Println(q.Song)
}

 
