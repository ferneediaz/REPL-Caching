package main

import (
    "bufio"
    "fmt"
    "os"
)

func startRepl () {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("PLEASE ENTER SOME TEXT")

    scanner.Scan()
    text := scanner.Text()

    fmt.Println("echoing: ", text)
	} 
    
}