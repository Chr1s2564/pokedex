package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResult struct {
	Name string
	Url  string
}

type LocationArea struct {
	Count    int
	Next     string
	Previous string
	Results  []LocationAreaResult
}

func GetAreas() (LocationArea, error) {
	const mapUrl = "https://pokeapi.co/api/v2/location-area/"
	req, err := http.NewRequest("GET", mapUrl, nil)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error creating request: %v", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error sending request: %v", err)
	}
	defer res.Body.Close()

	reader, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error reading body: %v", err)
	}

	var results = LocationArea{}
	err = json.Unmarshal(reader, &results)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error unmarshaling reader: %v", err)
	}

	return results, nil
}
