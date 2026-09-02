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

func printAreaList(areaList pokeapi.AreaList) {
	for _, area := range areaList.Areas {
		fmt.Println(area.Name)
	}
}

func getAreaList(cfg *config, url string) (pokeapi.AreaList, error) {
	areaList := pokeapi.AreaList{}
	var data []byte
	var ok bool
	if data, ok = cfg.Cache.Get(url); !ok {
		var err error
		data, err = pokeapi.GetAreaListData(url)
		if err != nil {
			return areaList, err
		}
		cfg.Cache.Add(url, data)
	}
	err := json.Unmarshal(data, &areaList)
	if err != nil {
		return areaList, err
	}
	return areaList, nil
}

func commandMap(cfg *config) error {
	areaList, err := getAreaList(cfg, cfg.Next)
	if err != nil {
		return err
	}

	printAreaList(areaList)

	cfg.Previous = derefOrEmpty(areaList.Previous)
	cfg.Next = areaList.Next

	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	areaList, err := getAreaList(cfg, cfg.Previous)
	if err != nil {
		return err
	}

	printAreaList(areaList)

	cfg.Next = areaList.Next
	cfg.Previous = derefOrEmpty(areaList.Previous)

	return nil
}
