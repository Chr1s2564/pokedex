package main

import (
	"fmt"
)

func displayMapb(cfg *config) error {
	areas, err := cfg.client.GetPrevious(cfg.previousUrl)
	if err != nil {
		return fmt.Errorf("unable to retrieve previous locations: %v", err)
	}
	for _, result := range areas.Results {
		fmt.Println(result.Name)
	}
	return nil
}
