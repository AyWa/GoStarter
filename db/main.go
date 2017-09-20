package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	// connector to postgre
	_ "github.com/lib/pq"
)

// docker run -d -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password123 --name db-my -p 5432:5432 --restart=always postgres
const (
	dbUser      = "user"
	dbPasseword = "password123"
	dbName      = "postgres"
)

//Db is the pointer to the db
var db *sqlx.DB

func init() {
	fmt.Println("initialize the connection to the db")
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPasseword, dbName)
	fmt.Println(dbinfo)
	db, err = sqlx.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
}
