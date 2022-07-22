package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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
	//var id, age, salary int
	//var name, address string

	var id, age, salary, name, address string
	//这个类型不需要与数据库类型匹配

	rows, err := db.Query("select * from company") //不需要加“;"
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() { //
		err := rows.Scan(&id, &name, &age, &address, &salary) //
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id, name, age, address, salary)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
}

func query_one(db *sql.DB, id int) {
	var tid, name, age, address, salary string
	err := db.QueryRow("select * from company where id = $1", id).Scan(&tid, &name, &age, &address, &salary)
	if err != nil {
		log.Fatal(err) //
	}
	fmt.Println(name)
}

func insert(db *sql.DB) {
	temp, err := db.Prepare("insert into company (id,name,age,address,salary) values($1,$2,$3,$4,$5)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = temp.Exec(11, "Pablo", 22, "Tokyo", 100000)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("inserted into company succesfully")
	}
}

func update(db *sql.DB) {
	temp, err := db.Prepare("update company set salary = $1 where id = $2")
	if err != nil {
		log.Fatal(err)
	}

	_, err = temp.Exec(200000, 11)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("updated company succesfully")
	}
}

func delete(db *sql.DB) {
	_, err := db.Exec("delete from company where id = $1", 10)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("deleted company succesfully")
	}
}

func main() {
	db := connectDB()
	query(db)
	//insert(db)
	query_one(db, 11)
	update(db)
	query(db)
	delete(db)
	query(db)
}
