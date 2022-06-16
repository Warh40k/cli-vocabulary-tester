package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	translations := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to easy vocabulary test utility. Please, append some words with translations below.")
	for true {
		fmt.Println("Foreign word")
		scanner.Scan()
		var foreign string = scanner.Text()

		fmt.Println("Translation")
		scanner.Scan()
		var transl string = scanner.Text()

		if foreign == "" || transl == "" {
			fmt.Println("Empty input, ignored")
			continue
		} else if foreign == "stop" || transl == "stop" {
			for f, t := range translations {
				fmt.Printf("Foreign language %s with translation %s\n", f, t)
			}
			break
		}
		translations[foreign] = transl
	}

	fmt.Println(GetKeys(translations))

}

func GetKeys(dict map[string]string) (keys []string) {
	keys = make([]string, int(len(dict)))
	var i int = 0

	for f := range dict {
		keys[i] = f
		i++
	}

	return
}
