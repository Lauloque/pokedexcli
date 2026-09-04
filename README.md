# pokedexcli

This is a command-line Pokexed with a REPL, API calls, JSON parsing, and caching in Go.

Made in the context of the [Boot.dev class: Build a Pokedex in Go](https://www.boot.dev/courses/build-pokedex-cli-golang)

## Usage:

Run `./pokedexcli` from a terminal to start the command pokedex cli, then use any of these commands:

- `map`: List the next 20 location areas in the Pokemon world
- `mapb`: List the previous 20 location areas in the Pokemon world
- `explore`: List the pokemons found in a given area. e.g. `explore analave-city-area`
- `catch`: Try to catch a given Pokemon e.g. `catch pikachu`
- `inspect`: Inspect your Pokedex entry for a given Pokemon e.g. `inspect pikachu`
- `pokedex`: List the pokemons registered in your Pokedex
- `exit`: Exit the Pokedex
- `help`: Displays a help message
