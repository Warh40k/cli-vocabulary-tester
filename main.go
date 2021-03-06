package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"

	_ "github.com/mattn/go-sqlite3" // pass blank identifier the unused import error
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("Welcome to Go vocabulary test utility. Please, select what to do next:\n1 - add definitions to database;\n2 - show current definitions;\nempty line or 'q' - exit\n3 - test (from Foreign)\n4 - test (from native)")
	scanner.Scan()
	choice := scanner.Text()
	switch choice {
	case "1":
		FillDb()
	case "2":
		ShowDb()
	case "3":
		Test(true)
	case "4":
		Test(false)
	case "q", "":
		return
	default:
		main()
	}
	main()
}
func GetData(query string) *sql.Rows {
	db, err := sql.Open("sqlite3", "./words.db")
	if err != nil {
		panic(err)
	}
	res, err := db.Query(query)
	if err != nil {
		log.Panic(err)
	}
	db.Close()
	return res
}

func ShowDb() {
	res := GetData("SELECT * FROM vocabulary")
	fmt.Println("Id\tForeign\tTranslation")
	for res.Next() {
		var id int
		var foreign, translate string
		err := res.Scan(&id, &foreign, &translate)
		checkErr(err)
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
			checkErr(err)
			stmt, err := tx.Prepare("INSERT INTO vocabulary(FOREIGN_WORD, TRANSLATION) VALUES(?,?)")
			checkErr(err)
			for f, t := range vocabulary {
				_, err = stmt.Exec(f, t)
				checkErr(err)
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

func Test(foreign bool) {
	res := GetData("SELECT ID, FOREIGN_WORD, TRANSLATION FROM vocabulary")
	vocab := RowsToMap(res)
	keys := GetKeys(vocab)
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	for i := 0; i < len(keys); i++ {
		var item = [2]string{keys[i], vocab[keys[i]]}
		if foreign == false {
			item[0], item[1] = item[1], item[0]
		}
		fmt.Println(item[0])
		scanner.Scan()
		if scanner.Text() == "" {
			break
		}
		fmt.Println(item[1])
	}
}

func RowsToMap(rows *sql.Rows) map[string]string {
	vocab := make(map[string]string)
	for rows.Next() {
		var i int
		var f, t string
		rows.Scan(&i, &f, &t)
		vocab[f] = t
	}
	return vocab
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//func checkCount(string ) (count int) {
//	for rows.Next() {
//		err := rows.Scan(&count)
//		checkErr(err)
//	}
//	return count
//}
func GetKeys(dict map[string]string) (keys []string) {
	keys = make([]string, int(len(dict)))
	var i int = 0

	for f := range dict {
		keys[i] = f
		i++
	}
	return
}
