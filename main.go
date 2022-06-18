package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Welcome to Go vocabulary test utility. Please, select what to do next.")
	translations := FillDb()
	fmt.Println(translations)
}

func FillDb() (dict map[string]string) {
	//connecting database
	vocabulary := make(map[string]string)
	db, err := sql.Open("sqlite3", "./words.db")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Attempting to fill database with words")
	for {
		fmt.Println("Foreign word")
		scanner.Scan()
		var foreign string = scanner.Text()

		fmt.Println("Translation")
		scanner.Scan()
		var transl string = scanner.Text()
		if foreign == "" && transl == "" {
			fmt.Println("Empty input, stop")

			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			stmt, err := tx.Prepare("INSERT INTO vocabulary(FOREIGN_WORD, TRANSLATION) VALUES(?,?)")
			if err != nil {
				log.Fatal(err)
			}
			for f, t := range vocabulary {
				_, err = stmt.Exec(f, t)
				if err != nil {
					log.Fatal(err)
				}
			}
			tx.Commit()
			db.Close()

			for f, t := range vocabulary {
				fmt.Printf("Foreign language %s with translation %s\n", f, t)
			}
			break
		} else if foreign == "" || transl == "" {
			fmt.Println("One field is empty, ignore")
			continue
		}
		vocabulary[foreign] = transl
	}

	return vocabulary
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
