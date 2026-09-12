package pokeapi

import (
	"fmt"
)

func GetPrevious(previousUrl string) (LocationArea, error) {
	if previousUrl == "" {
		return LocationArea{}, fmt.Errorf("you're on the first page")
	}
	mapUrl := previousUrl
	areas, err := GetAreas(mapUrl)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error retrieving previous areas: %v", err)
	}
	return areas, nil
}
