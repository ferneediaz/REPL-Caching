package pokeapi

import (
    "fmt"
    "log"
)

func CallbackMap() error {
	pokeapiClient := NewClient() // NewClient should be defined in this package

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
}