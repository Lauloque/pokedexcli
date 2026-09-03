/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"fmt"
	"os"
)

func commandExit(_ *config, _ string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
