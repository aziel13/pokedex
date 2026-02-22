package pokeapi

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestExplore(t *testing.T) {
	var buffer bytes.Buffer

	expected := "pastoria-city-area tentacooltentacruelmagikarpgyaradosremoraidoctillerywingullpelippershellosgastrodon"

	c := NewClient(5 * time.Second)

	exploreAreaString := `pastoria-city-area`
	//exploreArea := &exploreAreaString

	RespPokemonEncounters, err2 := c.ExploreLocation(exploreAreaString)

	if err2 != nil {

		t.Error(err2)
	}

	RespPokemonEncountersName := RespPokemonEncounters.Name

	pkmn := ""
	for _, pokemon := range RespPokemonEncounters.Pokemon_encounters {

		pkmn += pokemon.Pokemon.NAME
	}

	fmt.Fprintln(&buffer, RespPokemonEncountersName+" "+pkmn)

	if err2 != nil {
		t.Errorf("ExploreLocation returned unexpected error %v", err2)
	}

	actual := buffer.String()

	actual = strings.TrimSpace(actual)

	if actual != expected {
		t.Errorf("actual does not match expected\n actual: \n %v\n expected: \n %v", actual, expected)
	}
}
