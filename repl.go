/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
			break
		}
		input := scanner.Text()

		cleanInput := cleanInput(input)

		if len(cleanInput) == 0 {
			fmt.Print("Please type a command. Type 'help' to see commands list\n")
			continue
		}
		i := cleanInput[0]

		command, exists := cfg.commands[i]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("Unknown Command: '%s'\n", i)
			continue
		}
	}
}
