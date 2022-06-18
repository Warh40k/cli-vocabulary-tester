package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Go vocabulary test utility. Please, select what to do next.")
	translations := FillDb()
	fmt.Println(translations)
}

func FillDb() (dict map[string]string) {
	scanner := bufio.NewScanner(os.Stdin)
	translations := make(map[string]string)
	fmt.Println("Attempting to fill database with words")
	for {
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

	return translations
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
