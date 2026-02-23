package pokeapi

// RespShallowLocations -
type RespPokemon struct {
	Name            string `json:"name"`
	Base_Experience int    `json:"base_experience"`
	Height          int    `json:"height"`
	Weight          int    `json:"weight"`

	Stats []struct {
		Base_stat int  `json:"base_stat"`
		Effort    int  `json:"effort"`
		Stat      stat `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Slot  int   `json:"slot"`
		Atype aType `json:"type"`
	} `json:"types"`
}

type aType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type RespPokemon_species struct {
	Name string `json:"name"`

	Capture_rate int `json:"capture_rate"`
}
