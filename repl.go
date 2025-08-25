package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ercorn/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient *pokeapi.Client
	next          *string
	previous      *string
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
	}
}

func cleanInput(text string) []string {
	output := strings.Fields(strings.ToLower(text))
	return output
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n\n")
	for key, value := range command_list {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandMap(cfg *config) error {
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

func commandMapB(cfg *config) error {
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

func startRepl(user_cfg *config) {
	init_commands()
	pokedex_scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		pokedex_scanner.Scan()
		user_input := cleanInput(pokedex_scanner.Text())
		if len(user_input) == 0 {
			continue
		}
		if _, exists := command_list[user_input[0]]; exists {
			err := command_list[user_input[0]].callback(user_cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
