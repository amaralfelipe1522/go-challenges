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
		myList.insertFinal(&music)
	}
	fmt.Println("Playlist 1.0:")
	myList.showPlaylist()

	newMusic := Music{
		title:  "Konohatron",
		author: "MC Maha",
	}

	afterMusic := listRandom[4] // "Baba Baby"

	myList.insertInto(&newMusic, afterMusic)
	fmt.Println("Playlist 1.1:")
	myList.showPlaylist()

	deleteMusic := Music{
		title:  "Konohatron",
		author: "MC Maha",
	}

	myList.removeMusic(&deleteMusic)
	fmt.Println("Playlist 1.2:")
	myList.showPlaylist()
}

func (pl *Playlist) removeMusic(deleteMusic *Music) {
	//Adiciona a nova musica no primeiro link caso o tamanho da playlist seja zero
	if pl.length == 0 {
		fmt.Println("Playlist já está vazia")
	} else {
		currentMusic := pl.start
		if currentMusic.title == deleteMusic.title {
			pl.start = currentMusic.next
		} else {
			for currentMusic.next != nil {
				beforeMusic := currentMusic
				currentMusic = currentMusic.next
				if currentMusic.title == deleteMusic.title {
					beforeMusic.next = currentMusic.next
				}
			}
		}
		pl.length--
	}
}

func (pl *Playlist) insertInto(newMusic *Music, afterMusic Music) {
	//Adiciona a nova musica no primeiro link caso o tamanho da playlist seja zero
	if pl.length == 0 {
		pl.start = newMusic
	} else {
		currentMusic := pl.start
		// Varre os links até encontrar um titulo de musica igual
		for currentMusic.title != afterMusic.title {
			currentMusic = currentMusic.next
		}
		// Caso o Next não seja nulo, a nova musica troca de lugar com o next atual
		if currentMusic.next != nil {
			newMusic.next = currentMusic.next
			currentMusic.next = newMusic
		} else {
			currentMusic.next = newMusic
		}
	}
	pl.length++
}

func (pl *Playlist) insertFinal(newMusic *Music) {
	//Adiciona a nova musica no primeiro link caso o tamanho da playlist seja zero
	if pl.length == 0 {
		pl.start = newMusic
	} else {
		// Varre os links até que encontre o elemento Next nulo, para poder armazenar a musica na playlist
		currentMusic := pl.start
		for currentMusic.next != nil {
			currentMusic = currentMusic.next
		}
		currentMusic.next = newMusic
	}
	pl.length++
}

func (pl *Playlist) showPlaylist() {
	// Inicia a busca pelo elemento Start
	list := pl.start
	//Exibe as musicas até que lista fique nula
	for list != nil {
		fmt.Printf("Music: %s - Author: %s;\n", list.title, list.author)
		list = list.next
	}
}
