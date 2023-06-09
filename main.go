package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	// password = "Musa the Fairy"
	password = "temporary"
	dbname   = "books"
)

var db *sql.DB

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(`SELECT "title", "author", "pages_num", "review" from books`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var title string
		var author string
		var pages_num int
		var review string
		// fmt.Println("oh no!")
		// var e error

		err = rows.Scan(&title, &author, &pages_num, &review)
		if err != nil {
			panic(err)
		}
		fmt.Println(title)
		fmt.Println(author)
		fmt.Println(pages_num)
		fmt.Println(review)
	}
}
