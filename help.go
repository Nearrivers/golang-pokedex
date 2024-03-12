package main

import "fmt"

func commandHelp() error {
	commands := getCommands()
	fmt.Println("\nBienvenue dans le Pokédex!")
	fmt.Println("Usage:")

	for i := range commands {
		fmt.Printf("\n%s: %s", commands[i].name, commands[i].description)
	}

	return nil
}
