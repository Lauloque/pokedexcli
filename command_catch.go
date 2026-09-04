/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"

	"github.com/Lauloque/pokedexcli/internal/pokeapi"
)

func getPkmnInfo(cfg *config, url string) (pokeapi.Pokemon, error) {
	pkmn := pokeapi.Pokemon{}
	var data []byte
	var ok bool

	// check if data is in cache, or cache it
	if data, ok = cfg.Cache.Get(url); !ok {
		var err error
		data, err = pokeapi.GetPokeapiData(url)
		if err != nil {
			return pkmn, err
		}
		cfg.Cache.Add(url, data)
	}

	err := json.Unmarshal(data, &pkmn)
	if err != nil {
		return pkmn, err
	}
	return pkmn, nil
}

func addPkmn2Pokedex(cfg *config, pkmn pokeapi.Pokemon) {
	_, exists := cfg.Pokedex[pkmn.Name]
	if !exists {
		cfg.Pokedex[pkmn.Name] = pkmn
	}
}

func tryCatch(cfg *config, pkmn pokeapi.Pokemon) {
	// current highest base_level pkmn ios Arceus at 360
	// https://pwo-wiki.info/index.php/Arceus
	// https://pwo-wiki.info/index.php/Base_Experience
	attempt := rand.IntN(360)
	fmt.Printf("%s catch rate is %d, attempting %d...\n", pkmn.Name, pkmn.CatchRate, attempt)
	if attempt >= pkmn.CatchRate {
		fmt.Println(pkmn.Name, "was caught!")
		addPkmn2Pokedex(cfg, pkmn)
	} else {
		fmt.Println(pkmn.Name, "escaped!")
	}
}

func commandCatch(cfg *config, arg string) error {
	url := fmt.Sprint("https://pokeapi.co/api/v2/pokemon/", arg)

	pkmn, err := getPkmnInfo(cfg, url)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pkmn.Name)
	tryCatch(cfg, pkmn)

	return nil
}
