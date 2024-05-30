package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "poc_partition_table_postgres"
)

func insertLine(db *sql.DB, name string, email string, countryCode string) {
	insertSQL := `INSERT INTO users_partitioned (name, email, country_code) VALUES ($1, $2, $3) RETURNING id`
	var lastInsertId int
	err := db.QueryRow(insertSQL, name, email, countryCode).Scan(&lastInsertId)
	if err != nil {
		log.Fatalf("Unable to execute the query: %v\n", err)
	}
}

func insertBatch(db *sql.DB, valueStrings []string, valueArgs []interface{}) {
	// defer wg.Done()
	insertSQL := fmt.Sprintf("INSERT INTO users_partitioned (name, email, country_code) VALUES %s", strings.Join(valueStrings, ","))
	_, err := db.Exec(insertSQL, valueArgs...)
	if err != nil {
		log.Fatalf("Unable to execute batch insert: %v\n", err)
	}

}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(70) NOT NULL,
        email VARCHAR(120) NOT NULL,
		country_code VARCHAR(3)
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Unable to execute the query: %v\n", err)
	}

	valueStrings := make([]string, 0)
	valueArgs := make([]interface{}, 0)
	counter := 0
	for i := 0; i < 10000000; i++ {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", counter*3+1, counter*3+2, counter*3+3))
		valueArgs = append(valueArgs, fmt.Sprintf("%s%d", "Test", i))
		valueArgs = append(valueArgs, fmt.Sprintf("%s%d", "test@gmail.com", i))
		valueArgs = append(valueArgs, "USA")

		if len(valueStrings) == 21000 {
			insertBatch(db, valueStrings, valueArgs)
			valueStrings = make([]string, 0)
			valueArgs = make([]interface{}, 0)
			counter = 0
		} else {
			counter += 1
		}

	}

	if len(valueStrings) > 0 {
		insertBatch(db, valueStrings, valueArgs)
	}

	fmt.Println("Started script")
}
