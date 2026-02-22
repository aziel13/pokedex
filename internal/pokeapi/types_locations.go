package pokeapi

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Name     *string `json:"name"`

	Results []struct {
		Name               string `json:"name"`
		pokemon_encounters string `json:"pokemon_encounters"`
		URL                string `json:"url"`
	} `json:"results"`
}
