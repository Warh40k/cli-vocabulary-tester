package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to easy vocabulary test utility. Please, append some words with translations below.")
	for true {
		fmt.Println("Foreign word")
		foreign, _, _ := reader.ReadLine()

		fmt.Println("Translation")
		transl, _, _ := reader.ReadLine()

		fmt.Printf("You have chosen foreign word %s and translation %s\n", foreign, transl)
	}
}
