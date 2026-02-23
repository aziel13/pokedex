package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"strconv"
)

func commandHelp(w io.Writer, cfg *configuration, input string, pokedex pokedex) error {

	fmt.Fprintln(w, "Welcome to the Pokedex!\nUsage:")
	for _, cmd := range getCommands() {

		outString := fmt.Sprintf("%s: %s\n", cmd.name, cmd.description)

		fmt.Fprintln(w, outString)
	}

	return nil
}

func commandMapForward(w io.Writer, cfg *configuration, input string, pokedex pokedex) error {

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

func commandMapBack(w io.Writer, cfg *configuration, input string, pokedex pokedex) error {

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

func commandExit(w io.Writer, cfg *configuration, input string, pokedex pokedex) error {
	fmt.Fprintln(w, "Closing the Pokedex... Goodbye!")
	//os.Exit(0)

	return errors.New("ExitCode1")
}

func commandExplore(w io.Writer, cfg *configuration, location_name string, pokedex pokedex) error {

	pokemonEncountersResp, err := cfg.pokeapiClient.ExploreLocation(location_name)

	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Exploring "+pokemonEncountersResp.Name+"...")
	fmt.Fprintln(w, "Found Pokemon:")
	for _, pokemon := range pokemonEncountersResp.Pokemon_encounters {
		fmt.Fprintln(w, " - "+pokemon.Pokemon.NAME)

		pokedex.KnownPokemon[pokemon.Pokemon.NAME] = true

	}

	return nil
}

func commandCapture(w io.Writer, cfg *configuration, pokemon_name string, pokedex pokedex) error {

	RespPokemon, err := cfg.pokeapiClient.Get_Pokemon_Data(pokemon_name)

	if err != nil {
		return err
	}
	RespPokemon_Species, err2 := cfg.pokeapiClient.Get_Pokemon_Species_Data(pokemon_name)
	if err2 != nil {
		return err
	}
	fmt.Fprintln(w, "Throwing a Pokeball at "+RespPokemon.Name+"...")

	captureRate := RespPokemon_Species.Capture_rate
	baseXp := RespPokemon.Base_Experience
	percentCaptureRate := captureRate / 100

	randomNumber := rand.IntN(baseXp)

	if randomNumber > (percentCaptureRate / 2) {
		fmt.Fprintln(w, RespPokemon.Name+" was caught!")
		pokedex.capturedPokemon[RespPokemon.Name] += 1
	} else {
		fmt.Fprintln(w, RespPokemon.Name+" escaped!")
	}

	return nil
}

func commandInspect(w io.Writer, cfg *configuration, pokemon_name string, pokedex pokedex) error {

	_, ok := pokedex.capturedPokemon[pokemon_name]

	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	RespPokemon, err := cfg.pokeapiClient.Get_Pokemon_Data(pokemon_name)

	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Name: "+RespPokemon.Name)
	fmt.Fprintln(w, "Height: "+strconv.Itoa(RespPokemon.Height))
	fmt.Fprintln(w, "Weight: "+strconv.Itoa(RespPokemon.Weight))
	fmt.Fprintln(w, "Stats:")

	for _, stat := range RespPokemon.Stats {

		fmt.Fprintln(w, "  -"+stat.Stat.Name+": "+strconv.Itoa(stat.Base_stat))

	}
	fmt.Fprintln(w, "Types:")

	for _, typ := range RespPokemon.Types {
		fmt.Fprintln(w, "  - "+typ.Atype.Name)
	}

	return nil
}
