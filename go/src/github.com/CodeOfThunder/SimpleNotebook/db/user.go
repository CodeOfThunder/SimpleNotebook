package db

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"github.com/CodeOfThunder/SimpleNotebook/models"
)

var db *sql.DB
var err error

func checkErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
		//fmt.Println(err.Error())
	}
}

func InitDB() {
	db, err := sql.Open("sqlite3", "test.db")
	
	checkErr(err)
	create_userinfo := `CREATE TABLE IF NOT EXISTS userinfo (
		uid INTEGER PRIMARY KEY AUTOINCREMENT, 
		uname VARCHAR(40) NOT NULL,
		upwd  VARCHAR(40) NOT NULL ); `
	_, err = db.Exec(create_userinfo)
	checkErr(err)
	create_bookinfo := `CREATE TABLE IF NOT EXISTS bookinfo (
		bid INTEGER PRIMARY KEY AUTOINCREMENT,
		bookname VARCHAR(30) NOT NULL,
		owner_Id INTEGER NOT NULL,
		FOREIGN KEY (owner_id) REFERENCES userinfo(uid)
		); `
	_, err = db.Exec(create_bookinfo)
	checkErr(err)
	create_notes := `CREATE TABLE IF NOT EXISTS notes (
		nid INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(60) NOT NULL,
		content VARCHAR(300),
		author_id INTEGER NOT NULL,
		book_id INTERGER NOT NULL,
		record_time TEXT NOT NULL,
		FOREIGN KEY (author_id) REFERENCES userinfo(uid)
		FOREIGN KEY (book_id) REFERENCES bookinfo(bid)
		);  `
	_, err = db.Exec(create_notes)
	checkErr(err)
}

func ConnectDB() {
	db, err = sql.Open("sqlite3", "test.db")
	checkErr(err)
	InitDB()
}

func InsertUser(username string,password string) {
	// db, err = sql.Open("sqlite3", "test.db")
	// checkErr(err)
	stmt, err := db.Prepare("INSERT INTO userinfo(uname,upwd) VALUES(?,?)")
	checkErr(err)
	_, err = stmt.Exec(username, password)
	checkErr(err)
}

func AllUsers() []models.UserModel {
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	var users = make([]models.UserModel,0)
	for rows.Next() {
		var user models.UserModel
	
		err = rows.Scan(&user.Uid, &user.UserName, &user.Password)
		checkErr(err)
		users = append(users, user)
		//fmt.Println(uid, uname, upwd)
	}
	return users
}

func VerifyUserLogin(userName string,password string) bool {
	query_sql := "SELECT upwd FROM userinfo WHERE uname='" + userName + "'"
	fmt.Println(query_sql)
	rows, err := db.Query(query_sql)
	checkErr(err)
	if rows.Next() {
		var upwd string
		err = rows.Scan(&upwd)
		checkErr(err)
		if upwd == password {
			return true
		}
	}
	return false
}

func SelectUserId(userName string) int64 {
	rows, err := db.Query("SELECT uid FROM userinfo WHERE uname='" + userName + "';")
	checkErr(err)
	if rows.Next() {
		var UserId int64
		err = rows.Scan(&UserId)
		checkErr(err)
		return UserId
	}
	return -1
}
