package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     *string
	previous *string
}

type location_area struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
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
	res, err := http.Get(*c.next)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	current_location := location_area{}
	err = json.Unmarshal(body, &current_location)
	if err != nil {
		return err
	}

	c.next = current_location.Next
	c.previous = current_location.Previous

	//print areas
	for _, result := range current_location.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapB(c *config) error {
	if c.previous == nil {
		return fmt.Errorf("you're on the first page")
	}

	res, err := http.Get(*c.previous)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	current_location := location_area{}
	err = json.Unmarshal(body, &current_location)
	if err != nil {
		return err
	}

	c.next = current_location.Next
	c.previous = current_location.Previous

	//print areas
	for _, result := range current_location.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func startRepl() {
	init_commands()
	next_str := "https://pokeapi.co/api/v2/location-area/"
	user_config := config{
		next:     &next_str,
		previous: nil,
	}
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
			err := command_list[user_input[0]].callback(&user_config)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
