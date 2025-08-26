package main

import (
	"time"

	"github.com/ercorn/pokedexcli/internal/pokeapi"
)

//

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	user_cfg := config{
		pokeapiClient: pokeClient,
		next:          nil,
		previous:      nil,
		pokemon_party: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&user_cfg)
}
