package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "errors" // Import the errors package
)

// Define the Error type if it's missing
var Error = errors.New("an error occurred")

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("PLEASE ENTER SOME TEXT")

    scanner.Scan()
    text := scanner.Text()
    cleaned := cleanInput(text)
    if len(cleaned) == 0 {
        continue
    }
    commandName := cleaned[0]

    availableCommands := getCommands()

    command , ok := availableCommands[commandName]

    if !ok {
        fmt.Println("invalid command")
        continue
    }

    command.callback()

	} 
    
}
type cliCommand struct {
    name        string
    description string
    callback    func() error

}
func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "help": {
            name:        "help",
            description: "Prints the help menu",
            callback:    callbackHelp,
        },
        "exit": {
            name:        "exit",
            description: "Turns off the Pokedex",
            callback:    callbackExit,
        },
    }
}
func cleanInput(input string) []string {
    lowered := strings.ToLower(input)

    words := strings.Fields(lowered)
    
    return words
}

