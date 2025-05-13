package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.Fields(strings.ToLower(text))
	return output
}

func startRepl() {
	pokedex_scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		pokedex_scanner.Scan()
		user_input := cleanInput(pokedex_scanner.Text())
		if len(user_input) == 0 {
			continue
		}
		fmt.Printf("Your command was: %v\n", user_input[0])
	}
}
