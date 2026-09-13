package main

import (
	"fmt"
)

func displayMap(cfg *config, arg string) error {
	if arg != "" {
		return fmt.Errorf("map command doesn't take any argument")
	}
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
