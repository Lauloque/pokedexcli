/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

func main() {
	cfg := &config{
		commands: getCommands(),
		Previous: "",
		Next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	}

	startRepl(cfg)
}
