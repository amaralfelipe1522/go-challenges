package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
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

// Pokemon armazena as informações de nome e altura consumidas da API
type Pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
}

// Recupera da API um ou mais pokémons passados por parametro, retornando um slice do tipo []Pokemon
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
		// Converte []bytes para JSON e armazena no endereço de pkm
		json.Unmarshal(responseData, &pkm)
		//fmt.Println(*pkm)
		pkmSlice = append(pkmSlice, *pkm)
	}

	waitgroup.Done()
	return pkmSlice
}

// Varre o slice de pokemons ordenando em altura e atualizando o slice
func byHeight(pkmList []Pokemon) []Pokemon {

	sort.Slice(pkmList, func(i, j int) bool {
		return pkmList[i].Height < pkmList[j].Height
	})

	return pkmList
}

func main() {
	var pkm Pokemon
	var pkmList []Pokemon
	var waitgroup sync.WaitGroup
	//Aguarda todos os waitgroup.Done() encerrarem, garantindo assim que todas as chamadas de API foram encerradas
	waitgroup.Add(2)
	go func() {
		//Armazena o pkmSlice retornado para a a struct pkmList
		pkmList = pkm.getPokemon(&waitgroup, "ditto", "charizard", "weedle", "mew", "bulbasaur")
		waitgroup.Done()
	}()
	waitgroup.Wait()

	fmt.Println("Lista de Pokémons:", pkmList)
	fmt.Println("Pokémons por altura:", byHeight(pkmList))
}
