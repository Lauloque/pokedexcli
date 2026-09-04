/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"fmt"
)

func commandInspect(cfg *config, arg string) error {

	pkmn, exists := cfg.Pokedex[arg]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", pkmn.Name)
	fmt.Println("Height:", pkmn.Height)
	fmt.Println("Weight:", pkmn.Weight)
	fmt.Println("Stats:")
	for _, s := range pkmn.PokeStats {
		fmt.Printf("  - %s: %d\n", s.Stat.Name, s.Level)
	}
	fmt.Println("Types:")
	for _, t := range pkmn.Types {
		fmt.Println("  -", t.Type.Name)
	}

	return nil
}
