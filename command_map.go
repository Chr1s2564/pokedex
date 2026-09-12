package main

import (
	"fmt"
)

func displayMap(cfg *config) error {
	areas, err := cfg.client.GetAreas(cfg.nextUrl)
	if err != nil {
		return fmt.Errorf("unable to retrieve areas: %v", err)
	}
	for _, result := range areas.Results {
		fmt.Println(result.Name)
	}
	cfg.nextUrl = areas.Next
	cfg.previousUrl = areas.Previous
	return nil
}
