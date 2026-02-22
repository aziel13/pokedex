package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type configuration struct {
	pokeapiClient    pokeapi.Client
	Next             *string
	Previous         *string
	explore_location *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(io.Writer, *configuration, string) error
}

func startRepl(w io.Writer, cfg *configuration) int {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex >")
		scanner.Scan()

		line := cleanInput(scanner.Text())
		if len(line) == 0 {
			continue
		}

		commandName := line[0]

		inputString := ""

		if len(line) > 1 {
			inputString = line[1]

		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(w, cfg, inputString)
			if err != nil {

				errText := err.Error()

				if len(errText) > 0 {

					if errText == "ExitCode1" {
						return 1
					}

					fmt.Println(errText)
				}
			}

			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

	return 0
}

func cleanInput(text string) []string {

	cleaned := []string{}

	split_list := strings.Split(text, " ")

	for _, v := range split_list {
		if strings.TrimSpace(v) != "" {
			cleaned = append(cleaned, strings.ToLower(strings.TrimSpace(v)))
		}

	}

	return cleaned
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"quit": {
			name:        "quit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Help about any command",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display Pokemon World Locations",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Display Pokemon World Locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "explore a world location. explore <location>",
			callback:    commandExplore,
		},
	}
}
