package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(city string) (CityPokemon, error) {
	// if city == "" {
	// 	return CityPokemon{}, errors.New("no city provider. please provide a city")
	// }
	url := BaseURL + "/location-area/" + city

	if val, ok := c.cache.GetCache(url); ok {
		pokeResp := CityPokemon{}
		err := json.Unmarshal(val, &pokeResp)
		if err != nil {
			return CityPokemon{}, err
		}
		return pokeResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CityPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CityPokemon{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CityPokemon{}, err
	}
	pokeResp := CityPokemon{}
	err = json.Unmarshal(body, &pokeResp)
	if err != nil {
		fmt.Println("error unmarshalling", err)
		return CityPokemon{}, err
	}

	c.cache.AddCache(url, body)
	return pokeResp, nil
}
