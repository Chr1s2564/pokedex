package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Stat struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Stat     Stat `json:"stat"`
}

type Type struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Types struct {
	PokeTypes Type `json:"type"`
}

type CatchPokemon struct {
	Name           string  `json:"name"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
}

func (c *Client) FindPokemon(pokemon string) (CatchPokemon, error) {
	pokemonUrl := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	data, ok := c.cache.Get(pokemonUrl)
	if ok {
		var caughtPokemon = CatchPokemon{}
		err := json.Unmarshal(data, &caughtPokemon)
		if err != nil {
			return CatchPokemon{}, fmt.Errorf("error retrieving pokemon data from cache: %v", err)
		}
		return caughtPokemon, nil
	}

	req, err := http.NewRequest("GET", pokemonUrl, nil)
	if err != nil {
		return CatchPokemon{}, fmt.Errorf("unable to create request: %v", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return CatchPokemon{}, fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	reader, err := io.ReadAll(res.Body)
	if err != nil {
		return CatchPokemon{}, fmt.Errorf("error reading response body:%v", err)
	}
	if res.StatusCode > 199 && res.StatusCode < 300 {
		c.cache.Add(pokemonUrl, reader)
	} else if res.StatusCode == 404 {
		return CatchPokemon{}, fmt.Errorf("pokemon does not exist")
	} else {
		return CatchPokemon{}, fmt.Errorf("status code error: %v", res.StatusCode)
	}

	var caughtPokemon = CatchPokemon{}
	err = json.Unmarshal(reader, &caughtPokemon)
	if err != nil {
		return CatchPokemon{}, fmt.Errorf("error unmarshaling data from request: %v", err)
	}
	return caughtPokemon, nil
}
