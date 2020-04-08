package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

/*
In order to successfully complete this challenge, your project will have to:

Create a simple REST API in Go that calls out to the PokeAPI and orders the height of these 5 pokemon:

	ditto
	charizard
	weedle
	mew
	bulbasaur

The PokeAPI documentation can be found here: https://pokeapi.co/
*/

//Pokemon armazena as informações de nome e altura consumidas da API
type Pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
}

func (pkm *Pokemon) getPokemon(name string, waitgroup *sync.WaitGroup) Pokemon {
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &pkm)
	fmt.Println(*pkm)
	waitgroup.Done()
	return *pkm
}

func main() {
	var pkm Pokemon
	var pkmList []Pokemon
	var waitgroup sync.WaitGroup
	waitgroup.Add(4)
	go func() {
		pkmList = append(pkmList, pkm.getPokemon("ditto", &waitgroup))
		waitgroup.Done()
	}()
	go func() {
		pkmList = append(pkmList, pkm.getPokemon("mew", &waitgroup))
		waitgroup.Done()
	}()
	waitgroup.Wait()

	// waitgroup.Add(1)
	// go pkm.getPokemon("mew", &waitgroup)
	// waitgroup.Wait()
	// pkmList = append(pkmList, pkm)

	fmt.Println(pkmList)
}
