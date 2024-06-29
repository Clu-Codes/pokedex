/*
	Steps I need to understand to complete this part:

1. How do I create a mechanism for user input?
2. Once I have that input, should I pass it through a channel using a go routine?
3. How does NewScanner increment through tokens and how does that work with the infinite for loop?
4. Create commandHelp function to display the help command's output
5. Create `commandExit` function to end the REPL environment and close the loop
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Printf(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex

`)
	return nil
}

func commandExit() error {
	return errors.New("exit")
}

func main() {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
	}

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !reader.Scan() {
			break
		}
		input := reader.Text()

		if cmd, exists := commands[input]; exists {
			err := cmd.callback()
			if err != nil {
				if err.Error() == "exit" {
					break
				}
			}
		} else {
			fmt.Println("Unknown command:", input)
		}
	}

	// 	go func readInput(c cliCommand) {
	// 		reader := bufio.NewScanner(os.Stdin)
	// 		fmt.Print("Pokedex > ")
	// 		for reader.Scan() {
	// 			fmt.Println(reader.Text())
	// 			ch <- reader.Text()
	// 		}
	// 	}()

	// loop:
	// 	for {
	// 		input := <-ch
	// 		switch {
	// 		case input == "help":
	// 			commandHelp()
	// 		case input == "exit":
	// 			commandExit()
	// 		default:
	// 			continue
	// 		}
	// 	}

}
