package main

import (
	"bytes"
	"errors"
	"pokedexcli/internal/pokeapi"
	"strings"
	"testing"
	"time"
)

func TestExit(t *testing.T) {

	var buffer bytes.Buffer
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &configuration{
		pokeapiClient: pokeClient,
	}

	expectedReturn := errors.New("ExitCode1")

	errorActual := commandExit(&buffer, cfg, "")

	if errorActual.Error() != expectedReturn.Error() {
		t.Errorf("actual does not match expected\nactual: %v\nexpected: %v", errorActual.Error(), expectedReturn.Error())
	}

}

func TestExitMessage(t *testing.T) {

	var buffer bytes.Buffer
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &configuration{
		pokeapiClient: pokeClient,
	}

	expectedMessage := "Closing the Pokedex... Goodbye!\n"

	commandExit(&buffer, cfg, "")

	actual := buffer.String()

	if actual != expectedMessage {
		t.Errorf("actual does not match expected\nactual: %v\nexpected: %v \nactual length: %v\nexpected length: %v", actual, expectedMessage, len(actual), len(expectedMessage))
	}

}

func TestHelp(t *testing.T) {

	// Without reproducing exactly what the help command does its difficult to test.
	// best I can think to do is test that its longer than the base welcome message.

	var buffer bytes.Buffer
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &configuration{
		pokeapiClient: pokeClient,
	}

	expectedMessage := "Welcome to the Pokedex!\nUsage:"

	commandHelp(&buffer, cfg, "")

	actual := buffer.String()

	if len(actual) < len(expectedMessage) {
		t.Errorf("actual does not match expected\nactual: %v\nexpected: %v \nactual length: %v\nexpected length: %v", actual, expectedMessage, len(actual), len(expectedMessage))
	}
}

func TestMap(t *testing.T) {

	// Testing the base map command is difficult. It should have pulled the first 20 locations from the pokedex api.

	var buffer bytes.Buffer
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &configuration{
		pokeapiClient: pokeClient,
	}

	expected := [...]string{
		"canalave-city-area",
		"eterna-city-area",
		"pastoria-city-area",
		"sunyshore-city-area",
		"sinnoh-pokemon-league-area",
		"oreburgh-mine-1f",
		"oreburgh-mine-b1f",
		"valley-windworks-area",
		"eterna-forest-area",
		"fuego-ironworks-area",
		"mt-coronet-1f-route-207",
		"mt-coronet-2f",
		"mt-coronet-3f",
		"mt-coronet-exterior-snowfall",
		"mt-coronet-exterior-blizzard",
		"mt-coronet-4f",
		"mt-coronet-4f-small-room",
		"mt-coronet-5f",
		"mt-coronet-6f",
		"mt-coronet-1f-from-exterior",
	}

	commandMapForward(&buffer, cfg, "")

	actual := []string{}
	stringArray := strings.Split(buffer.String(), "\n")

	for i := range stringArray {
		line := strings.TrimSpace(stringArray[i])

		if line != "" {
			actual = append(actual, line)
		}
	}

	//
	/*
		pastoria-city-area
		sunyshore-city-area
		sinnoh-pokemon-league-area
		oreburgh-mine-1f
		oreburgh-mine-b1f
		valley-windworks-area
		eterna-forest-area
		fuego-ironworks-area
		mt-coronet-1f-route-207
		mt-coronet-2f
		mt-coronet-3f
		mt-coronet-exterior-snowfall
		mt-coronet-exterior-blizzard
		mt-coronet-4f
		mt-coronet-4f-small-room
		mt-coronet-5f
		mt-coronet-6f
		mt-coronet-1f-from-exterior*/

	if len(actual) != len(expected) {
		t.Errorf("actual does not match expected\nactual: %v\nexpected: %v \n", len(actual), len(expected))
	}

}

func TestMapBackOnFirstPage(t *testing.T) {

	// this should cause the your on you're on the first page error

	var buffer bytes.Buffer
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &configuration{
		pokeapiClient: pokeClient,
	}

	expectedMessage := errors.New("you're on the first page")
	actual := commandMapBack(&buffer, cfg, "")

	if actual.Error() != expectedMessage.Error() {
		t.Errorf("actual does not match expected\nactual: %v\nexpected: %v", actual, expectedMessage)
	}

}

func TestMapBackAfterFirstPage(t *testing.T) {

	// the result of mapb should be the same as the result for the first call of the forward command...

	var buffer bytes.Buffer
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &configuration{
		pokeapiClient: pokeClient,
	}

	commandMapForward(&buffer, cfg, "")

	expectedMessage := buffer.String()

	commandMapForward(&buffer, cfg, "")

	buffer.Reset()

	commandMapBack(&buffer, cfg, "")

	actual := buffer.String()

	if actual != expectedMessage {
		t.Errorf("actual does not match expected\nactual: %v\nexpected: %v \nactual length: %v\nexpected length: %v", actual, expectedMessage, len(actual), len(expectedMessage))
	}

}

func TestExploreCommand(t *testing.T) {
	var buffer bytes.Buffer
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &configuration{
		pokeapiClient: pokeClient,
	}

	exploreAreaString := `pastoria-city-area`
	expectedMessage := "Exploring pastoria-city-area...\nFound Pokemon:\n - tentacool\n - tentacruel\n " +
		"- magikarp\n - gyarados\n - remoraid\n - octillery\n - wingull\n - pelipper\n - shellos\n" +
		" - gastrodon\n"

	commandExplore(&buffer, cfg, exploreAreaString)

	actual := buffer.String()

	if actual != expectedMessage {
		t.Errorf("actual does not match expected\nactual: %v\nexpected: %v", actual, expectedMessage)
	}

}

func TestExploreNoLocationCommand(t *testing.T) {
	var buffer bytes.Buffer
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &configuration{
		pokeapiClient: pokeClient,
	}

	exploreAreaString := ``
	expectedMessage := "A location value is required. for example 'explore pastoria-city-area'"

	err := commandExplore(&buffer, cfg, exploreAreaString)

	actual := err.Error()

	if actual != expectedMessage {
		t.Errorf("actual does not match expected\nactual: %v\nexpected: %v", actual, expectedMessage)
	}

}

func TestCaptureCommand(t *testing.T) {
	var buffer bytes.Buffer
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &configuration{
		pokeapiClient: pokeClient,
	}

	exploreAreaString := `pikachu`
	expectedMessage := "Throwing a Pokeball at pikachu...\npikachu escaped!"
	expectedMessage2 := "Throwing a Pokeball at pikachu...\npikachu was caught!"

	err := commandCapture(&buffer, cfg, exploreAreaString)

	if err != nil {
		t.Error(err)
	}

	actual := buffer.String()

	if !(actual == expectedMessage || actual == expectedMessage2) {
		t.Errorf("actual does not match expected\nactual: %v\nexpected: %v or %v", actual, expectedMessage, expectedMessage2)
	}

}

/*
func TestCaptureNoInputCommand(t *testing.T) {
	var buffer bytes.Buffer
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &configuration{
		pokeapiClient: pokeClient,
	}

	exploreAreaString := ``
	expectedMessage := "A pokemon name value is required. for example 'capture pikachu'"

	err := commandCapture(&buffer, cfg, exploreAreaString)

	actual := err.Error()

	if actual != expectedMessage {
		t.Errorf("actual does not match expected\nactual: %v\nexpected: %v", actual, expectedMessage)
	}

}*/
