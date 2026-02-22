package pokeapi

// RespShallowLocations -
type RespPokemon struct {
	Name            string `json:"name"`
	Base_Experience int    `json:"base_experience"`
}

type RespPokemon_species struct {
	Name string `json:"name"`

	Capture_rate int `json:"capture_rate"`
}
