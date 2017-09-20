package db

import (
	"fmt"

	"github.com/aywa/goNotify/auth"
	"github.com/aywa/goNotify/config"
)

var schema = `CREATE TABLE users (
  id SERIAL PRIMARY KEY,
	password TEXT,
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL
);`

// User is the struct of the user schema
type User struct {
	ID        int    `db:"id"`
	Password  string `db:"password"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

func init() {
	if *config.GetFlag().ShouldInitDb {
		fmt.Println("initialize user table")
		_, err := db.Exec(schema)
		if err != nil {
			panic(err)
		}
		tx := db.MustBegin()
		tx.MustExec("INSERT INTO users (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
		tx.NamedExec("INSERT INTO users (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &User{FirstName: "Jane", LastName: "Citizen", Email: "jane.citzen@example.com"})
		tx.Commit()
		CreateUser(&User{FirstName: "Marc", LastName: "Hurabielle", Email: "marc.hurabielle@gmail.com", Password: "secret"})
		people := []User{}
		db.Select(&people, "SELECT * FROM users")
		fmt.Println(people)
		for _, p := range people {
			fmt.Println(p)
		}
	}
}

// CreateUser create a user in a db (salt the password too)
func CreateUser(u *User) (err error) {
	u.Password, err = auth.SaltPassword(u.Password)
	if err != nil {
		println(err)
		return err
	}
	fmt.Println(*u)
	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO users (password, first_name, last_name, email) VALUES (:password, :first_name, :last_name, :email)", u)
	err = tx.Commit()
	if err != nil {
		println(err)
	}
	return err
}
