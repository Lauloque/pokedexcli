/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import "strings"

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
