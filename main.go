package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()

		word := clearInput(scanner.Text())

		if word == nil {
			continue
		}
		fmt.Printf("Your command was: %v\n", strings.TrimRight(word[0], "\r\n"))

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
}
