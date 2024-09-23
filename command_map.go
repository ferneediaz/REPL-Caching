package main

import (
	"fmt"

	"github.com/ferneediaz/REPL-Caching/internal/pokeapi"
)

func CallbackMap(cfg *config) error {
	var err error
	var resp pokeapi.LocationAreasResp

	if cfg.nextLocationAreaURL != nil {
		resp, err = cfg.pokeapiClient.ListLocationAreas(*cfg.nextLocationAreaURL)
	} else {
		resp, err = cfg.pokeapiClient.ListLocationAreas("")
	}

	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}

func CallbackMapB(cfg *config) error {
	if cfg.previousLocationAreaURL == nil {
		fmt.Println("You're on the first page.")
		return nil
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(*cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}
