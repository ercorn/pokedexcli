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
	}

	startRepl(&user_cfg)
}
