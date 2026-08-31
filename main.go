/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()

		cleanInput := cleanInput(input)

		fmt.Printf("Your command was: %s\n", cleanInput[0])
	}
}
