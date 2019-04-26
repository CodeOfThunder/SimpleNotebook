package db

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/CodeOfThunder/SimpleNotebook/models"
)

func AddBook(book models.BookModel) {
	stmt, err := db.Prepare(`INSERT INTO bookinfo(bookname,owner_id) 
	VALUES(?,?,?) ; `)
	checkErr(err)
	_, err = stmt.Exec(book.BookName, book.Owner_id)
	checkErr(err)
}

func AllBooks() []models.BookModel {
	rows, err := db.Query("SELECT bid,bookname FROM bookinfo")
	checkErr(err)

	var books = make([]models.BookModel, 0)
	for rows.Next() {
		var book models.BookModel

		err = rows.Scan(&book.Bid, &book.BookName,)
		checkErr(err)
		books = append(books, book)
	}
	return books
}

func AllBooksOfUser(username string) []models.BookModel {
	rows, err := db.Query(`SELECT bid,bookname FROM bookinfo WHERE owner_Id in
							( SELECT uid FROM userinfo WHERE uname='` + username + "');")
	checkErr(err)

	var books = make([]models.BookModel, 0)
	for rows.Next() {
		var book models.BookModel

		err = rows.Scan(&book.Bid, &book.BookName)
		checkErr(err)
		books = append(books, book)
	}
	return books
}
