package pokeapi

// RespShallowLocations -
type RespPokemonGrowthRates struct {
	Name string `json:"name"`

	Levels []struct {
		levelData levelData `json:"pokemon"`
	} `json:"levels"`
}

type levelData struct {
	LEVEL      string `json:"level"`
	experience string `json:"experience"`
}
