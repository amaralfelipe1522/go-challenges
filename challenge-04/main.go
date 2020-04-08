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

func (pkm *Pokemon) getPokemon(waitgroup *sync.WaitGroup, name ...string) []Pokemon {
	var pkmSlice []Pokemon
	for i := 0; i < len(name); i++ {
		url := "https://pokeapi.co/api/v2/pokemon/" + name[i]
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
		pkmSlice = append(pkmSlice, *pkm)
	}

	waitgroup.Done()
	return pkmSlice
}

func main() {
	var pkm Pokemon
	var pkmList []Pokemon
	var waitgroup sync.WaitGroup

	waitgroup.Add(2)
	go func() {
		pkmList = pkm.getPokemon(&waitgroup, "ditto", "mew", "pikachu")
		//pkmList = append()
		waitgroup.Done()
		fmt.Println(pkmList)
	}()
	waitgroup.Wait()
	//começou of FOR concluido
	fmt.Println(pkmList)
}
