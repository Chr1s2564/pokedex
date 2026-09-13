package main

import (
	"fmt"
)

func displayMapb(cfg *config, arg string) error {
	if arg != "" {
		return fmt.Errorf("mapb command doesn't take any argument")
	}
	areas, err := cfg.client.GetPrevious(cfg.previousUrl)
	if err != nil {
		return fmt.Errorf("unable to retrieve previous locations: %v", err)
	}
	for _, result := range areas.Results {
		fmt.Println(result.Name)
	}
	return nil
}
