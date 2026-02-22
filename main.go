package main

import (
	"os"
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	//done := make(chan bool)

	interval := 5 * time.Second

	pokeClient := pokeapi.NewClient(interval)

	cfg := &configuration{
		pokeapiClient: pokeClient,
	}

	code := startRepl(os.Stdout, cfg)
	if code == 1 {

		os.Exit(0)

	}
}

// "https://pokeapi.co/api/v2/{endpoint}/" location areas
// https://pokeapi.co/api/v2/location-area/{id or name}/
