package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin) // corrected assignment
    fmt.Println("PLEASE ENTER SOME TEXT")

    scanner.Scan()
    text := scanner.Text()

    fmt.Println("echoing: ", text)
}
