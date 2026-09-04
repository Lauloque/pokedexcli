/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ string) error {
	if len(cfg.Pokedex) < 1 {
		fmt.Println("Your Pokedex is empty... For now! Go catch some Pokemons!")
		return nil
	}
	for pkmn, _ := range cfg.Pokedex {
		fmt.Println(" -", pkmn)
	}

	return nil
}
