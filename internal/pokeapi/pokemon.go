package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Id                     int             `json:"id"`
	Name                   string          `json:"name"`
	IsDefault              bool            `json:"is_default"`
	Height                 int             `json:"height"`
	Weight                 int             `json:"weight"`
	BaseExperience         int             `json:"base_experience"`
	LocationAreaEncounters string          `json:"url"`
	Species                NamedResource   `json:"species"`
	Order                  int             `json:"order"`
	Forms                  []NamedResource `json:"forms"`
	Abilities              []struct {
		IsHidden bool          `json:"is_hidden"`
		Slot     int           `json:"slot"`
		Ability  NamedResource `json:"ability"`
	} `json:"abilities"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	PastAbilities []struct {
		Generation NamedResource `json:"generation"`
		Abilities  []struct {
			IsHidden bool          `json:"is_hidden"`
			Slot     int           `json:"slot"`
			Ability  NamedResource `json:"ability"`
		} `json:"abilities"`
	} `json:"past_abilities"`
	Types []struct {
		Slot int           `json:"slot"`
		Type NamedResource `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int           `json:"base_stat"`
		Effort   int           `json:"effort"`
		Stat     NamedResource `json:"stat"`
	} `json:"stats"`
	PastTypes []struct {
		Generation NamedResource `json:"generation"`
		Types      []struct {
			Slot int           `json:"slot"`
			Type NamedResource `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
	HeldItems []struct {
		Item           NamedResource `json:"item"`
		VersionDetails []struct {
			Rarity  int           `json:"rarity"`
			Version NamedResource `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	GameIndices []struct {
		GameIndex int           `json:"game_index"`
		Version   NamedResource `json:"version"`
	} `json:"game_indices"`
	Sprites struct {
		BackDefault      *string `json:"back_default"`
		BackFemale       *string `json:"back_female"`
		BackShiny        *string `json:"back_shiny"`
		BackShinyFemale  *string `json:"back_shiny_female"`
		FrontDefault     *string `json:"front_default"`
		FrontFemale      *string `json:"front_female"`
		FrontShiny       *string `json:"front_shiny"`
		FrontShinyFemale *string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault *string `json:"front_default"`
				FrontFemale  *string `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     *string `json:"front_default"`
				FrontFemale      *string `json:"front_female"`
				FrontShiny       *string `json:"front_shiny"`
				FrontShinyFemale *string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault *string `json:"front_default"`
				FrontShiny   *string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      *string `json:"back_default"`
				BackFemale       *string `json:"back_female"`
				BackShiny        *string `json:"back_shiny"`
				BackShinyFemale  *string `json:"back_shiny_female"`
				FrontDefault     *string `json:"front_default"`
				FrontFemale      *string `json:"front_female"`
				FrontShiny       *string `json:"front_shiny"`
				FrontShinyFemale *string `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      *string `json:"back_default"`
					BackGray         *string `json:"back_gray"`
					BackTransparent  *string `json:"back_transparent"`
					FrontDefault     *string `json:"front_default"`
					FrontGray        *string `json:"front_gray"`
					FrontTransparent *string `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      *string `json:"back_default"`
					BackGray         *string `json:"back_gray"`
					BackTransparent  *string `json:"back_transparent"`
					FrontDefault     *string `json:"front_default"`
					FrontGray        *string `json:"front_gray"`
					FrontTransparent *string `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationII struct {
				Crystal struct {
					BackDefault           *string `json:"back_default"`
					BackShiny             *string `json:"back_shiny"`
					BackShinyTransparent  *string `json:"back_shiny_transparent"`
					BackTransparent       *string `json:"back_transparent"`
					FrontDefault          *string `json:"front_default"`
					FrontShiny            *string `json:"front_shiny"`
					FrontShinyTransparent *string `json:"front_shiny_transparent"`
					FrontTransparent      *string `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault      *string `json:"back_default"`
					BackShiny        *string `json:"back_shiny"`
					FrontDefault     *string `json:"front_default"`
					FrontShiny       *string `json:"front_shiny"`
					FrontTransparent *string `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault      *string `json:"back_default"`
					BackShiny        *string `json:"back_shiny"`
					FrontDefault     *string `json:"front_default"`
					FrontShiny       *string `json:"front_shiny"`
					FrontTransparent *string `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIII struct {
				Emerald struct {
					FrontDefault *string `json:"front_default"`
					FrontShiny   *string `json:"front_shiny"`
				} `json:"emerald"`
				FireRedLeafGreen struct {
					BackDefault  *string `json:"back_default"`
					BackShiny    *string `json:"back_shiny"`
					FrontDefault *string `json:"front_default"`
					FrontShiny   *string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  *string `json:"back_default"`
					BackShiny    *string `json:"back_shiny"`
					FrontDefault *string `json:"front_default"`
					FrontShiny   *string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIV struct {
				DiamondPearl struct {
					BackDefault      *string `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        *string `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartGoldSoulSilver struct {
					BackDefault      *string `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        *string `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      *string `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        *string `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					BackDefault      *string `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        *string `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
					Animated         struct {
						BackDefault      *string `json:"back_default"`
						BackFemale       *string `json:"back_female"`
						BackShiny        *string `json:"back_shiny"`
						BackShinyFemale  *string `json:"back_shiny_female"`
						FrontDefault     *string `json:"front_default"`
						FrontFemale      *string `json:"front_female"`
						FrontShiny       *string `json:"front_shiny"`
						FrontShinyFemale *string `json:"front_shiny_female"`
					} `json:"animated"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVI struct {
				OmegaRubyAlphaSapphire struct {
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVII struct {
				Icons struct {
					FrontDefault *string `json:"front_default"`
					FrontFemale  *string `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationVIII struct {
				Icons struct {
					FrontDefault *string `json:"front_default"`
					FrontFemale  *string `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Moves []struct {
		Move                NamedResource `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int           `json:"level_learned_at"`
			Order           *int          `json:"order"`
			MoveLearnMethod NamedResource `json:"move_learn_method"`
			VersionGroup    NamedResource `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
}

func (c *Client) GetPokemon(pokemon_name string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon_name
	pokemon := Pokemon{}

	//use cache first
	if val, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(val, &pokemon)

		if err != nil {
			return pokemon, err
		}
		return pokemon, nil
	}

	//get .../pokemon/pokemon_name response
	res, err := http.Get(url)
	if err != nil {
		return pokemon, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return pokemon, fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return pokemon, err
	}

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return pokemon, err
	}

	//add to cache
	c.cache.Add(url, body)

	return pokemon, nil
}
