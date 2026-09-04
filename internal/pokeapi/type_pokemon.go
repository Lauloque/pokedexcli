/* SPDX-License-Identifier: GPL-3.0-or-later */
package pokeapi

type Pokemon struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CatchRate int    `json:"base_experience"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
	PokeStats []struct {
		Level int `json:"base_stat"`
		Stat  struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
