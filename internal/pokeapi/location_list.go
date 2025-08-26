package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type location_area_shallow_resp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (location_area_shallow_resp, error) {
	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	if pageURL != nil {
		url = *pageURL
	}

	//use cache
	if val, exists := c.cache.Get(url); exists {
		current_location := location_area_shallow_resp{}
		err := json.Unmarshal(val, &current_location)

		if err != nil {
			return location_area_shallow_resp{}, err
		}
		return current_location, nil
	}
	//cache use block

	res, err := http.Get(url)
	if err != nil {
		return location_area_shallow_resp{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return location_area_shallow_resp{}, fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return location_area_shallow_resp{}, err
	}

	current_location := location_area_shallow_resp{}

	err = json.Unmarshal(body, &current_location)
	if err != nil {
		return location_area_shallow_resp{}, err
	}

	//add to cache
	c.cache.Add(url, body)

	return current_location, nil
}
