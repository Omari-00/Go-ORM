package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "pgadmintest"
)

func connectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func query(db *sql.DB) {
	var id, age, salary int
	var name, address string

	rows, err := db.Query("select * from company where id = 1")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() { //
		err := rows.Scan(&id, &name, &age, &address, &salary) //

		if err != nil {
			fmt.Println(err)
		}
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id, name, age, address, salary)
}

func main() {
	db := connectDB()
	query(db)//
}
