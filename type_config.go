/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"github.com/Lauloque/pokedexcli/internal/pokeapi"
	"github.com/Lauloque/pokedexcli/internal/pokecache"
)

type config struct {
	commands map[string]cliCommand
	Next     string
	Previous string
	Cache    *pokecache.Cache
	Pokedex  map[string]pokeapi.Pokemon
}
