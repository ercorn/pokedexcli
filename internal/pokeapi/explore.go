package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NamedResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type location_area struct {
	Id        int    `json:"id"`
	GameIndex int    `json:"game_index"`
	Name      string `json:"name"`
	Names     []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"location"`
	PokemonEncounters []struct {
		//fill in shape from response
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
			EncounterDetails []struct {
				//
				Chance          int `json:"chance"`
				MaxLevel        int `json:"max_level"`
				MinLevel        int `json:"min_level"`
				ConditionValues []struct {
					//
					Id        int           `json:"id"`
					Name      string        `json:"name"`
					Condition NamedResource `json:"condition"`
					Names     []struct {
						Name     string `json:"name"`
						Language struct {
							Name string `json:"name"`
							Url  string `json:"url"`
						} `json:"language"`
					} `json:"names"`
				} `json:"condition_values"`
				Method NamedResource `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) Explore(area_name string) (location_area, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + area_name

	//use cache first
	if val, exists := c.cache.Get(url); exists {
		current_location := location_area{}
		err := json.Unmarshal(val, &current_location)

		if err != nil {
			return location_area{}, err
		}
		return current_location, nil
	}

	//get .../location-area/area_name response
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

	//add to cache
	c.cache.Add(url, body)

	return current_location, nil
}
