package db

import (
	"fmt"

	"github.com/aywa/goNotify/auth"
	"github.com/aywa/goNotify/config"
	"github.com/aywa/goUtil"
)

var schema = `CREATE TABLE users (
  id SERIAL PRIMARY KEY,
	password  TEXT,
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL
);`

// User is the struct of the user schema
type User struct {
	ID        int    `db:"id" json:"id,omitempty"`
	Password  string `db:"password" json:"password,omitempty"`
	FirstName string `db:"first_name" json:"firstName,omitempty"`
	LastName  string `db:"last_name" json:"lastName,omitempty"`
	Email     string `db:"email" json:"email"`
}

func init() {
	if *config.GetFlag().ShouldInitDb {
		fmt.Println("initialize user table")
		_, err := db.Exec(schema)
		goUtil.ErrorPanic(err)
		CreateUser(&User{FirstName: "Marc", LastName: "Hurabielle", Email: "marc.hurabielle@gmail.com", Password: "secret"})
		CreateUser(&User{FirstName: "Cho", LastName: "Hura", Email: "cho.hurabielle@gmail.com", Password: "secret"})
		people := []User{}
		err = db.Select(&people, "SELECT * FROM users ORDER BY first_name ASC")
		goUtil.ErrorPanic(err)
		for _, p := range people {
			fmt.Println(p)
		}
	}
}

// CreateUser create a user in a db (salt the password too)
func CreateUser(u *User) (err error) {
	u.Password, err = auth.SaltPassword(u.Password)
	goUtil.ErrorHandler(err, goUtil.ErrorLogger)
	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO users (password, first_name, last_name, email) VALUES (:password, :first_name, :last_name, :email)", u)
	err = tx.Commit()
	goUtil.ErrorHandler(err, goUtil.ErrorLogger)
	return err
}

// CheckCredential take a user and check if the password is same as the bd
func CheckCredential(uToCheck User) (uFromBd User, err error) {
	err = db.Get(&uFromBd, "SELECT * FROM users WHERE email=$1", uToCheck.Email)
	goUtil.ErrorHandler(err, goUtil.ErrorLogger)
	return uFromBd, auth.CompareHashAndPassword([]byte(uFromBd.Password), []byte(uToCheck.Password))
}

// GetPrivateUser return a user with private info
func GetPrivateUser(email string) (uFromBd User, err error) {
	err = db.Get(&uFromBd, "SELECT first_name, last_name, email FROM users WHERE email=$1", email)
	return uFromBd, err
}
