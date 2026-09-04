/* SPDX-License-Identifier: GPL-3.0-or-later */
package pokeapi

type AreaList struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Areas    []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
