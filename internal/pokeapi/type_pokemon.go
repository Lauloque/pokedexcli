/* SPDX-License-Identifier: GPL-3.0-or-later */
package pokeapi

type Pokemon struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CatchRate int    `json:"base_experience"`
	Height    int    `json:"height"`
	Types     struct {
		Name string `json:"name"`
	} `json:"types"`
}
