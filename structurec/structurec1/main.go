//go:generate mockgen -source ${GOFILE} -destination mocks_test.go -package ${GOPACKAGE}_test
package main

import "fmt"

// Вот тебе задачка на интерфейсы, просто на интерес: Структура Player, 
//с методом CreatePlaylist (можно сделать как мапу), который на вход принимает Song, 
//Podcast, VideoClip, AudioBook, каждая из этих структур обладает 
//своими полями но точно у каждой есть название и длительность 
//структура Player должна также иметь метод переключения на следующую запись в плейлисте Next(),
// и метод который считает всю длительность плейлиста

type Item interface{
 Getlen()int
 Getname()string
}
type Song struct{
  Name string
  Time int
}
type Podcast struct{
  Name string
  Time int
}

func (s *Song)Getlen()int{
  return s.Time
}
func (s *Song)Getname()string{
  return s.Name
}
func (s *Podcast)Getlen()int{
  return s.Time
}
func (s *Podcast)Getname()string{
  return s.Name
}


type Player struct {
  Playlist map[string]Item
}

func (p *Player)CreatePlaylist(list ...Item){
  for _,item:=range list{
    p.Playlist[item.Getname()]=item
  }
}
func (p *Player)Timelen()int{
  sum:=0
  for _,item:=range p.Playlist{
   sum+=item.Getlen()
   
  }
  return sum
}
func main(){
  a:=&Song{Name: "edede", Time: 12}
  b:=&Podcast{Name: "ewrwe", Time: 542}
  p:=Player{Playlist: make(map[string]Item)}
  p.CreatePlaylist(a,b)
	 fmt.Println(p.Timelen())

}