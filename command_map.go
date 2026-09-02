/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"encoding/json"
	"fmt"

	"github.com/Lauloque/pokedexcli/internal/pokeapi"
)

func derefOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func commandMap(cfg *config) error {
	nextUrl, previousUrl, err := pokeapi.PrintAreaList(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Previous = derefOrEmpty(previousUrl)
	cfg.Next = nextUrl

	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	nextUrl, previousUrl, err := pokeapi.PrintAreaList(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = nextUrl
	cfg.Previous = derefOrEmpty(previousUrl)

	return nil
}
