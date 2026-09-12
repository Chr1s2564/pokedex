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

func (c *Client) GetAreas(mapUrl string) (LocationArea, error) {
	data, ok := c.cache.Get(mapUrl)
	if ok {
		var results = LocationArea{}
		err := json.Unmarshal(data, &results)
		if err != nil {
			return LocationArea{}, err
		}
		return results, nil
	}

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
	if res.StatusCode > 199 && res.StatusCode < 300 {
		c.cache.Add(mapUrl, reader)
	} else {
		return LocationArea{}, fmt.Errorf("error status code: %v", res.StatusCode)
	}

	var results = LocationArea{}
	err = json.Unmarshal(reader, &results)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error unmarshaling reader: %v", err)
	}

	return results, nil
}
