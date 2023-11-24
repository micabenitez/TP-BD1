package lib

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	bolt "go.etcd.io/bbolt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mypassword"
	dbnameTP = "tp"
)

func CreateDatabase() {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "mypassword"
	dbname := "postgres"

	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname = %s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("* Creando database '" + dbnameTP + "' *")
	_, err = db.Exec("create database " + dbnameTP + ";")
	if err != nil {
		log.Fatal(err)
	}
}

func DropDatabase() {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "mypassword"
	dbname := "postgres"

	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname = %s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("* Borrando database '" + dbnameTP + "' *")
	_, err = db.Exec("drop database " + dbnameTP + ";")
	if err != nil {
		log.Fatal(err)
	}
}

func Connection() *sql.DB {
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname = %s sslmode=disable",
		host, port, user, password, dbnameTP)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func ConnectionBolt() *bolt.DB {
	db, err := bolt.Open("tpbolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
