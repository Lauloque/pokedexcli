/* SPDX-License-Identifier: GPL-3.0-or-later */
package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

type AreaList struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Areas    []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type AreaInfo struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetPokeapiData(url string) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Couldn't get location-area:")
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code %d:\n%s", res.StatusCode, body)
	}
	if err != nil {
		return nil, fmt.Errorf("Couldn't read response body: %w", err)
	}

	return body, nil
}
