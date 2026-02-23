package pokeapi

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestGetPokemon(t *testing.T) {
	var buffer bytes.Buffer
	expected := "pikachu 112"

	pokemonInput := "pikachu"

	c := NewClient(5 * time.Second)

	pokemonGrowthRate, err := c.Get_Pokemon_Data(pokemonInput)

	if err != nil {
		t.Errorf("ExploreLocation returned unexpected error %v", err)
	}

	pokemonName := pokemonGrowthRate.Name
	pokemonBaseXp := pokemonGrowthRate.Base_Experience

	fmt.Fprintln(&buffer, pokemonName+" "+strconv.Itoa(pokemonBaseXp))

	actual := buffer.String()

	actual = strings.TrimSpace(actual)

	if actual != expected {
		t.Errorf("actual does not match expected\n actual: \n %v\n expected: \n %v", actual, expected)
	}

}

func TestGetPokemonHeight(t *testing.T) {
	var buffer bytes.Buffer
	expected := "Height: 4"

	pokemonInput := "pikachu"

	c := NewClient(5 * time.Second)

	PokemonData, err := c.Get_Pokemon_Data(pokemonInput)

	if err != nil {
		t.Errorf("ExploreLocation returned unexpected error %v", err)
	}

	pokemonHeight := PokemonData.Height

	fmt.Fprintln(&buffer, "Height: "+strconv.Itoa(pokemonHeight))

	actual := buffer.String()

	actual = strings.TrimSpace(actual)

	if actual != expected {
		t.Errorf("actual does not match expected\n actual: \n %v\n expected: \n %v", actual, expected)
	}

}
func TestGetPokemonWeight(t *testing.T) {
	var buffer bytes.Buffer
	expected := "Weight: 60"

	pokemonInput := "pikachu"

	c := NewClient(5 * time.Second)

	Pokemon_Data, err := c.Get_Pokemon_Data(pokemonInput)

	if err != nil {
		t.Errorf("ExploreLocation returned unexpected error %v", err)
	}

	pokemonWeight := Pokemon_Data.Weight

	fmt.Fprintln(&buffer, "Weight: "+strconv.Itoa(pokemonWeight))

	actual := buffer.String()

	actual = strings.TrimSpace(actual)

	if actual != expected {
		t.Errorf("actual does not match expected\n actual: \n %v\n expected: \n %v", actual, expected)
	}

}
func TestGetPokemonStats(t *testing.T) {
	var buffer bytes.Buffer
	expected := "pikachu\nhp: 35\nattack: 55\ndefense: 40\nspecial-attack: 50\nspecial-defense: 50\nspeed: 90"

	pokemonInput := "pikachu"

	c := NewClient(5 * time.Second)

	Pokemon_Data, err := c.Get_Pokemon_Data(pokemonInput)

	if err != nil {
		t.Errorf("ExploreLocation returned unexpected error %v", err)
	}

	pokemonName := Pokemon_Data.Name
	pokemonStats := Pokemon_Data.Stats

	statData := ""

	for _, stat := range pokemonStats {
		statData += stat.Stat.Name + ": " + strconv.Itoa(stat.Base_stat) + "\n"
	}

	fmt.Fprintln(&buffer, pokemonName+"\n"+statData)

	actual := buffer.String()

	actual = strings.TrimSpace(actual)

	if actual != expected {
		t.Errorf("actual does not match expected\n actual: \n %v\n expected: \n %v", actual, expected)
	}

}
func TestGetPokemonTypes(t *testing.T) {
	var buffer bytes.Buffer
	expected := "pikachu 112"

	pokemonInput := "pikachu"

	c := NewClient(5 * time.Second)

	Pokemon_Data, err := c.Get_Pokemon_Data(pokemonInput)

	if err != nil {
		t.Errorf("ExploreLocation returned unexpected error %v", err)
	}

	pokemonName := Pokemon_Data.Name
	pokemonBaseXp := Pokemon_Data.Base_Experience

	typeData := ""

	for _, aType := range Pokemon_Data.Types {
		typeData += "" + strconv.Itoa(aType.Slot) + " " + aType.Atype.Name + "\n"
	}

	fmt.Fprintln(&buffer, pokemonName+" "+strconv.Itoa(pokemonBaseXp))

	actual := buffer.String()

	actual = strings.TrimSpace(actual)

	if actual != expected {
		t.Errorf("actual does not match expected\n actual: \n %v\n expected: \n %v", actual, expected)
	}

}
func TestGetPokemonSpecies(t *testing.T) {
	var buffer bytes.Buffer
	expected := "pikachu 190"

	pokemonInput := "pikachu"

	c := NewClient(5 * time.Second)

	Pokemon_Species_Data, err := c.Get_Pokemon_Species_Data(pokemonInput)

	if err != nil {
		t.Errorf("Get Pokemon Species returned unexpected error %v", err)
	}

	pokemonName := Pokemon_Species_Data.Name
	pokemonCaptureRate := Pokemon_Species_Data.Capture_rate

	fmt.Fprintln(&buffer, pokemonName+" "+strconv.Itoa(pokemonCaptureRate))

	actual := buffer.String()

	actual = strings.TrimSpace(actual)

	if actual != expected {
		t.Errorf("actual does not match expected\n actual: \n %v\n expected: \n %v", actual, expected)
	}

}
