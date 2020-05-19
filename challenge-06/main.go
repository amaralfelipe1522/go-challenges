package main

import (
	"fmt"
)

/*
You will need to create a linked list Playlist data structure as well as a struct that will represent a Song in your Playlist.
Each Song should have a pointer to the next Song in the Playlist.

Acceptance Criteria
You must implement this Playlist without using the standard library linked list implementation!!
You must implement a function or method that allows you to remove a Song from your Playlist
You must implement a function or method that allows you to insert a Song into your Playlist at any given point
You must implement a function or method that allows you to add a Song to the end of your Playlist
You must implement a function or method that displays the entire playlist in order.
*/

//Music é a estrutura básica da música armazenada
type Music struct {
	title  string
	author string
	next   *Music
}

//Playlist armazena o nome da lista e o ponteiro para a primeira música
type Playlist struct {
	length int
	name   string
	start  *Music
}

func main() {
	var listRandom [6]Music
	listRandom[0].author = "Nickelback"
	listRandom[0].title = "Photograph"
	listRandom[1].author = "Linkin Park"
	listRandom[1].title = "Faint"
	listRandom[2].author = "RHCP"
	listRandom[2].title = "Dani California"
	listRandom[3].author = "Gorillaz"
	listRandom[3].title = "Feel Good Inc"
	listRandom[4].author = "Kelly Key"
	listRandom[4].title = "Baba Baby"
	listRandom[5].author = "Latino"
	listRandom[5].title = "Festa no Apê"

	myList := &Playlist{}

	for i := 0; i < len(listRandom); i++ {
		music := Music{
			title:  listRandom[i].title,
			author: listRandom[i].author,
		}
		myList.insertMusic(&music)
	}
	myList.showPlaylist()
}

func (pl *Playlist) insertMusic(newMusic *Music) {
	if pl.length == 0 {
		pl.start = newMusic
	} else {
		currentMusic := pl.start
		for currentMusic.next != nil {
			currentMusic = currentMusic.next
		}
		currentMusic.next = newMusic
	}
	pl.length++
}

func (pl *Playlist) showPlaylist() {
	list := pl.start
	for list != nil {
		fmt.Printf("Music: %s - Author: %s;\n", list.title, list.author)
		list = list.next
	}
}
