package main

// mysql: github.com/go-sql-driver/mysql
// sqlite3: github.com/mattn/go-sqlite3
// postgres: github.com/lib/pq

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// alternatively
	// db, err := sql.Open("postgres",
	//     "postgres://{user}:{password}@{hostname}:{port}/{database-name}?sslmode=disable")
	db, err := sql.Open("postgres",
		"user=user password=pass dbname=db port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Ping successful!")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Unexpected Error!")
		}
	}(db)
}
