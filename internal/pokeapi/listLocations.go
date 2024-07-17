package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (PokeLoc, error) {
	url := BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.GetCache(url); ok {
		locationsResp := PokeLoc{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return PokeLoc{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeLoc{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeLoc{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeLoc{}, err
	}

	locationResp := PokeLoc{}
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return PokeLoc{}, err
	}

	c.cache.AddCache(url, body)
	return locationResp, nil
}
