package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error){
	url := baseURL + "/pokemon/" + name

	//  Use the cache data if found
	data, found := c.cache.Get(url)
	if found {
		cacheResp := Pokemon{}
		err := json.Unmarshal(data, &cacheResp)
		if err != nil {
			return Pokemon{}, err
		}
		return cacheResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	// Add to cache
	c.cache.Add(url, dat)

	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemonResp, nil
}