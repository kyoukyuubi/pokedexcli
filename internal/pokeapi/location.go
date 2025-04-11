package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) Location(name string) (RespLocation, error){
	url := baseURL + "/location-area/" + name

	//  Use the cache data if found
	data, found := c.cache.Get(url)
	if found {
		cacheResp := RespLocation{}
		err := json.Unmarshal(data, &cacheResp)
		if err != nil {
			return RespLocation{}, err
		}
		return cacheResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	// add to the cache
	c.cache.Add(url, dat)

	locationResp := RespLocation{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocation{}, err
	}
	return locationResp, nil
}