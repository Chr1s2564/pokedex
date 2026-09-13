package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocationAreaDetails struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func (c *Client) GetPokemons(areaName string) (LocationAreaDetails, error) {
	locationUrl := "https://pokeapi.co/api/v2/location-area/" + areaName
	data, ok := c.cache.Get(locationUrl)
	if ok {
		var pokemons = LocationAreaDetails{}
		err := json.Unmarshal(data, &pokemons)
		if err != nil {
			return LocationAreaDetails{}, fmt.Errorf("error unmarshaling data from cache: %v", err)
		}
		return pokemons, nil
	}

	req, err := http.NewRequest("GET", locationUrl, nil)
	if err != nil {
		return LocationAreaDetails{}, fmt.Errorf("error creating request: %v", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return LocationAreaDetails{}, fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	reader, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetails{}, fmt.Errorf("error reading response: %v", err)
	}
	if res.StatusCode > 199 && res.StatusCode < 300 {
		c.cache.Add(locationUrl, reader)
	} else if res.StatusCode == 404 {
		return LocationAreaDetails{}, fmt.Errorf("location does not exist")
	} else {
		return LocationAreaDetails{}, fmt.Errorf("status code error: %v", res.StatusCode)
	}

	var pokemons = LocationAreaDetails{}
	err = json.Unmarshal(reader, &pokemons)
	if err != nil {
		return LocationAreaDetails{}, fmt.Errorf("error unmarshaling data from request: %v", err)
	}
	return pokemons, nil
}
