package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // pass blank identifier the unused import error
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("Welcome to Go vocabulary test utility. Please, select what to do next:\n1) Add definitions to database;\n2) Show current definitions.")
	scanner.Scan()
	choice := scanner.Text()
	switch choice {
	case "1":
		FillDb()
	case "2":
		ShowDb()
	default:
		main()
	}
}

func ShowDb() {
	db, err := sql.Open("sqlite3", "./words.db")
	if err != nil {
		panic(err)
	}
	res, err := db.Query("SELECT * FROM vocabulary")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Id\tForeign\tTranslation")
	for res.Next() {
		var id int
		var foreign, translate string

		err = res.Scan(&id, &foreign, &translate)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d\t%s\t%s\n", id, foreign, translate)
	}
}

func FillDb() (dict map[string]string) {
	//connecting database
	vocabulary := make(map[string]string)
	db, err := sql.Open("sqlite3", "./words.db")
	if err != nil {
		panic(err)
	}
	//input words
	fmt.Println("Attempting to fill database with words. To stop input, press CTRL-D or pass both empty values")
	for {
		fmt.Println("Foreign word")
		scanner.Scan()
		var foreign string = scanner.Text()

		fmt.Println("Translation")
		scanner.Scan()
		var transl string = scanner.Text()
		// inserting words to db
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
			break
		} else if foreign == "" || transl == "" {
			fmt.Println("One of the input fields is empty, ignore")
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
