package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) Get_Pokemon_Data(pokemonName string) (RespPokemon, error) {

	if pokemonName == "" {
		return RespPokemon{}, errors.New("A Pokemon Name value is required.'")
	}

	url := baseURL + `/pokemon/` + pokemonName

	//check if this url has a cached value
	data := c.cache.Get(url)

	if data == nil {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespPokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespPokemon{}, err
		}
		defer resp.Body.Close()

		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespPokemon{}, err
		}

		exploreResp := RespPokemon{}
		err = json.Unmarshal(dat, &exploreResp)
		if err != nil {
			return RespPokemon{}, err
		}

		// cache the results
		c.cache.Add(url, dat)

		return exploreResp, nil
	} else {

		locationsResp := RespPokemon{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespPokemon{}, err
		}

		return locationsResp, nil
	}

}
