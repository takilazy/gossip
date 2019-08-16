package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)



func init() {

	db, err := sql.Open("sqlite3", "./test.db")
	checkErr(err)

	createTable, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (uid INTEGER PRIMARY KEY AUTOINCREMENT, email VARCHAR(64), password VARCHAR(64), created DATE)")
	createTable.Exec()
	checkErr(err)

	createTable, err = db.Prepare("CREATE TABLE IF NOT EXISTS messages (mid INTEGER PRIMARY KEY AUTOINCREMENT, userfrom VARCHAR(64), userto VARCHAR(64) , body TEXT, created DATE)")
	createTable.Exec()
	checkErr(err)

	db.Close()

}

func (account *Account) InsertUser() {
	db, err := sql.Open("sqlite3", "./test.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO users(email, password, created) values(?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(account.Email, account.Password, time.Now())
	checkErr(err)

	db.Close()
}

func (account *Account) UpdateUser() {
	db, err := sql.Open("sqlite3", "./test.db")
	checkErr(err)

	stmt, err := db.Prepare("update users set password=? where email=?")
	checkErr(err)

	_, err = stmt.Exec(account.Password, account.Email)
	checkErr(err)

	db.Close()
}

func (account *Account) GetUser() (string, string){
	db, err := sql.Open("sqlite3", "./test.db")
	checkErr(err)

	stmt, err := db.Prepare("SELECT * FROM users where email=?")
	checkErr(err)

	var uid int
	var email, password string
	var created time.Time

	row, err := stmt.Query(account.Email)
	checkErr(err)
	for row.Next() {
		err = row.Scan(&uid, &email, &password, &created)
		checkErr(err)
		log.Println(uid, email, password, created)
	}

	db.Close()

	return email, password
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err)
	}
}