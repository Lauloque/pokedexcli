/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"encoding/json"
	"fmt"

	"github.com/Lauloque/pokedexcli/internal/pokeapi"
)

func getAreaInfo(cfg *config, url string) (pokeapi.AreaInfo, error) {
	areaInfo := pokeapi.AreaInfo{}
	var data []byte
	var ok bool

	// check if data is in cache, or cache it
	if data, ok = cfg.Cache.Get(url); !ok {
		var err error
		data, err = pokeapi.GetPokeapiData(url)
		if err != nil {
			return areaInfo, err
		}
		cfg.Cache.Add(url, data)
	}

	err := json.Unmarshal(data, &areaInfo)
	if err != nil {
		return areaInfo, err
	}
	return areaInfo, nil
}

func printEncounters(areaInfo pokeapi.AreaInfo) {
	for _, encounter := range areaInfo.PokemonEncounters {
		fmt.Println(" - ", encounter.Pokemon.Name)
	}
}

func commandExplore(cfg *config, arg string) error {
	url := fmt.Sprint("https://pokeapi.co/api/v2/location-area/", arg)

	areaInfo, err := getAreaInfo(cfg, url)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", arg)
	fmt.Println("Found Pokemon:")
	printEncounters(areaInfo)

	return nil
}
