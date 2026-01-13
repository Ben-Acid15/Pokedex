package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; 1 > 0; i++ {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		firstWord := cleanedInput[0]
		msg := fmt.Sprintf("Your command was: %s", firstWord)
		fmt.Println(msg)
	}
}
