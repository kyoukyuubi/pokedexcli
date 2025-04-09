package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input.Scan()
		words := cleanInput(input.Text())
		fmt.Println("Your command was:", words[0])
	}
}
