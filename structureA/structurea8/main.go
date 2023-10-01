//Создайте структуру Playlist, представляющую плейлист песен.
//Напишите методы AddSong(song string) и Play(), которые соответственно добавляют песню в плейлист
//и выводят сообщение "Играет песня: <название песни>". Используйте слайс для хранения песен.

package main

import "fmt"

type Playlist struct {
	Playlist []string
	Current int
}

func (p *Playlist) AddSong(song string) {
	p.Playlist = append(p.Playlist, song)
}
func (p *Playlist) Play()  {
	fmt.Printf("Играет песня:%s",p.Playlist[p.Current])
	p.Current++
}
func main() {
	new := Playlist{}
	new.AddSong("nebo")
	 
}
