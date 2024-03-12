package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	os.Exit(0)
	return nil
}

// Ce programme est un REPL (Read Eval Print Loop).
func main() {
	for {
		input := ""
		fmt.Print("\nPokédex > ")
		fmt.Scanln(&input)

		cmd, ok := getCommands()[input]
		if ok {
			err := cmd.callback()

			if err != nil {
				panic(err)
			}
		}
	}
}
