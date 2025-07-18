package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next string
	prev string
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
	}
}

func cleanInput(text string) []string {
	output := strings.Fields(strings.ToLower(text))
	return output
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n\n")
	for key, value := range command_list {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandMap(c *config) error {
	return nil
}

func startRepl() {
	init_commands()
	pokedex_scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		pokedex_scanner.Scan()
		user_input := cleanInput(pokedex_scanner.Text())
		if len(user_input) == 0 {
			continue
		}
		//fmt.Printf("Your command was: %v\n", user_input[0])
		if _, exists := command_list[user_input[0]]; exists {
			err := command_list[user_input[0]].callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
