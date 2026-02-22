package main

import (
	"errors"
	"fmt"
	"io"
)

func commandHelp(w io.Writer, cfg *configuration, input string) error {

	fmt.Fprintln(w, "Welcome to the Pokedex!\nUsage:")
	for _, cmd := range getCommands() {

		outString := fmt.Sprintf("%s: %s\n", cmd.name, cmd.description)

		fmt.Fprintln(w, outString)
	}

	return nil
}

func commandMapForward(w io.Writer, cfg *configuration, input string) error {

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.Next)

	if err != nil {
		return err
	}

	cfg.Next = locationsResp.Next
	cfg.Previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {

		fmt.Fprintln(w, loc.Name)

	}
	return nil
}

func commandMapBack(w io.Writer, cfg *configuration, input string) error {

	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.Previous)

	if err != nil {
		return err
	}

	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Fprintln(w, loc.Name)
	}
	return nil
}

func commandExit(w io.Writer, cfg *configuration, input string) error {
	fmt.Fprintln(w, "Closing the Pokedex... Goodbye!")
	//os.Exit(0)

	return errors.New("ExitCode1")
}

func commandExplore(w io.Writer, cfg *configuration, location_name string) error {

	pokemonEncountersResp, err := cfg.pokeapiClient.ExploreLocation(location_name)

	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Exploring "+pokemonEncountersResp.Name+"...")
	fmt.Fprintln(w, "Found Pokemon:")
	for _, pokemon := range pokemonEncountersResp.Pokemon_encounters {
		fmt.Fprintln(w, " - "+pokemon.Pokemon.NAME)
	}

	return nil
}
