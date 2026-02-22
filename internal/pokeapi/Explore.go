package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ExploreLocation(locationName string) (RespPokemonEncounters, error) {

	if locationName == "" {
		return RespPokemonEncounters{}, errors.New("A location value is required. for example 'explore pastoria-city-area'")
	}

	url := baseURL + `/location-area/` + locationName

	//check if this url has a cached value
	data := c.cache.Get(url)

	if data == nil {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespPokemonEncounters{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespPokemonEncounters{}, err
		}
		defer resp.Body.Close()

		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespPokemonEncounters{}, err
		}

		exploreResp := RespPokemonEncounters{}
		err = json.Unmarshal(dat, &exploreResp)
		if err != nil {
			return RespPokemonEncounters{}, err
		}

		// cache the results
		c.cache.Add(url, dat)

		return exploreResp, nil
	} else {

		locationsResp := RespPokemonEncounters{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespPokemonEncounters{}, err
		}

		return locationsResp, nil
	}

}
