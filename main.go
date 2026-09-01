/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

func main() {
	cfg := &config{
		commands: getCommands(),
	}

	startRepl(cfg)
}
