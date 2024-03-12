package main

import "github.com/nearrivers/go-cli-pokedex/pokeapi"

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Affiche un message d'aide",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Quitte le pokédex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Affiche les zones du monde",
			callback:    pokeapi.CommandMap,
		},
	}
}
