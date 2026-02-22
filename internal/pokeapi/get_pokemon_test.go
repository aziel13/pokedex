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
