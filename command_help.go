/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import "fmt"

func commandHelp(cfg *config, _ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, c := range cfg.commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}
