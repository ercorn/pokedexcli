package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/ercorn/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	pokeapiClient *pokeapi.Client
	next          *string
	previous      *string
	pokemon_party map[string]pokeapi.Pokemon
}

// commands
var command_list map[string]cliCommand

func init_commands() {
	command_list = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Display the previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "List all Pokemon in <area_name>. Usage: 'explore <area_name>' , (ex. 'explore pastoria-city-area')",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch <pokemon_name>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "See the detail of a Pokemon you have caught",
			callback:    commandInspect,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.Fields(strings.ToLower(text))
	return output
}

func commandExit(cfg *config, parameter string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, parameter string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n\n")
	for key, value := range command_list {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandMap(cfg *config, parameter string) error {
	current_location, err := cfg.pokeapiClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = current_location.Next
	cfg.previous = current_location.Previous

	//print areas
	for _, result := range current_location.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapB(cfg *config, parameter string) error {
	if cfg.previous == nil {
		return fmt.Errorf("you're on the first page")
	}

	current_location, err := cfg.pokeapiClient.ListLocations(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = current_location.Next
	cfg.previous = current_location.Previous

	//print names of the previous 20 locations
	for _, result := range current_location.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandExplore(cfg *config, parameter string) error {
	//Add an explore command. It takes the name of a location area as an argument
	if parameter == "" {
		return fmt.Errorf("no <area_name> provided")
	}

	current_location_area, err := cfg.pokeapiClient.Explore(parameter)
	if err != nil {
		return err
	}

	//parse list of pokemon from location_area endpoint response
	pokemon_list := current_location_area.PokemonEncounters
	//print list of pokemon found in this area
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemon_list {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}
	return nil
}

func catchChance(base_experience int) float32 {
	chance := 100.0 / float32(base_experience+50.0)
	if chance < 0.05 {
		chance = 0.05
	} else if chance > 0.95 {
		chance = 0.95
	}
	return chance
}

func tryCatch(base_experience int) bool {
	chance := catchChance(base_experience)
	r := float32(rand.IntN(base_experience)) / float32(base_experience)
	fmt.Println("Catch chance is: ", chance, " and r is: ", r, " and base exp is: ", base_experience)
	return r < chance
}

func commandCatch(cfg *config, parameter string) error {
	if parameter == "" {
		return fmt.Errorf("no <pokemon_name> provided")
	}

	//catch a pokemon
	pokemon, err := cfg.pokeapiClient.GetPokemon(parameter)

	if err != nil {
		return err
	}

	//fmt.Println(pokemon)
	fmt.Println("Throwing a Pokeball at", pokemon.Name+"...")

	if tryCatch(pokemon.BaseExperience) {
		//success
		cfg.pokemon_party[pokemon.Name] = pokemon
		fmt.Println(pokemon.Name + " was caught!")
	} else {
		fmt.Println(pokemon.Name + " escaped!")
	}
	return nil
}

func commandInspect(cfg *config, parameter string) error {
	if pokemon, exists := cfg.pokemon_party[parameter]; exists {
		//list stats
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pkmn_type := range pokemon.Types {
			fmt.Println(" -", pkmn_type.Type.Name)
		}
		return nil
	}
	return fmt.Errorf("you have not caught that pokemon")
}

func startRepl(user_cfg *config) {
	init_commands()
	pokedex_scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		pokedex_scanner.Scan()
		user_input := cleanInput(pokedex_scanner.Text())
		input_length := len(user_input)
		if input_length == 0 {
			continue
		}
		if _, exists := command_list[user_input[0]]; exists {
			if input_length == 1 {
				err := command_list[user_input[0]].callback(user_cfg, "")
				if err != nil {
					fmt.Println(err)
				}
				continue
			}
			if input_length == 2 {
				err := command_list[user_input[0]].callback(user_cfg, user_input[1])
				if err != nil {
					fmt.Println(err)
				}
				continue
			}

			fmt.Println("ERROR: Too many parameters. Use the 'help' command for proper usage guidance.")
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
