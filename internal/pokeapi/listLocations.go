package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func (c *Client) ListLocations(pageURL *string) (PokeLoc, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
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
	fmt.Printf("Body type: %v", reflect.TypeOf(body))
	if err != nil {
		return PokeLoc{}, err
	}

	locationResp := PokeLoc{}
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return PokeLoc{}, err
	}

	return locationResp, nil
}
