/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	// "encoding/json"
	"fmt"
	// "github.com/Lauloque/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, arg string) error {
	fmt.Printf("oy oi hoy!\nWe got a total of '%d' commands\n", len(cfg.commands))

	fmt.Printf("Received argument %s\n", arg)

	return nil
}
