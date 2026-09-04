/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of the pokemons found in a given area. e.g. 'explore analave-city-area'",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch a given Pokemon e.g. 'catch pikachu'",
			callback:    commandCatch,
		},
	}
}
