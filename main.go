package main

import (
    "fmt"
    "log"
    "github.com/ferneediaz/REPL-Caching/internal/pokeapi" // Corrected import path
)

func main() {
    pokeapiClient := pokeapi.NewClient() // Corrected the typo here

    resp, err := pokeapiClient.ListLocationAreas()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(resp)
    //startRepl()
}
