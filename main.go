/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"time"

	"github.com/Lauloque/pokedexcli/internal/pokeapi"
	"github.com/Lauloque/pokedexcli/internal/pokecache"
)

func main() {

	cache := pokecache.NewCache(5 * time.Minute)

	cfg := &config{
		commands: getCommands(),
		Previous: "",
		Next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Cache:    cache,
		Pokedex:  map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
