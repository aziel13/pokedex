package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) Get_Pokemon_Species_Data(pokemonName string) (RespPokemon_species, error) {

	if pokemonName == "" {
		return RespPokemon_species{}, errors.New("Pokemon species name is required. for example 'pikachu'")
	}

	url := baseURL + `/pokemon-species/` + pokemonName

	//check if this url has a cached value
	data := c.cache.Get(url)

	if data == nil {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespPokemon_species{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespPokemon_species{}, err
		}
		defer resp.Body.Close()

		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespPokemon_species{}, err
		}

		exploreResp := RespPokemon_species{}
		err = json.Unmarshal(dat, &exploreResp)
		if err != nil {
			return RespPokemon_species{}, err
		}

		// cache the results
		c.cache.Add(url, dat)

		return exploreResp, nil
	} else {

		locationsResp := RespPokemon_species{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespPokemon_species{}, err
		}

		return locationsResp, nil
	}

}
