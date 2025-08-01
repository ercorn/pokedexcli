package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type location_area struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func list_locations(pageURL *string) (location_area, error) {
	url := "https://pokeapi.co/api/v2/location-area/"

	if pageURL != nil {
		url = *pageURL
	}

	res, err := http.Get(url)
	if err != nil {
		return location_area{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return location_area{}, fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return location_area{}, err
	}

	current_location := location_area{}
	err = json.Unmarshal(body, &current_location)
	if err != nil {
		return location_area{}, err
	}

	return current_location, nil
}
