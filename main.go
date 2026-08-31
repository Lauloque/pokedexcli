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
		i := cleanInput[0]

		command, exists := getCommands()[i]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("Unknown Command: '%s'\n", i)
			continue
		}
	}
}
