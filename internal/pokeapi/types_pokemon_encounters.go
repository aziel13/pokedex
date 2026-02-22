package pokeapi

// RespShallowLocations -
type RespPokemonEncounters struct {
	Name string `json:"name"`

	Pokemon_encounters []struct {
		Pokemon pokemon `json:"pokemon"`
	} `json:"Pokemon_encounters"`
}

type pokemon struct {
	NAME string `json:"name"`
	URL  string `json:"url"`
}
