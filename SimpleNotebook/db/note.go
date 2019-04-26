package db

import (
	"time"
	_ "github.com/mattn/go-sqlite3"
	"github.com/CodeOfThunder/SimpleNotebook/models"
)

func AddNote(note models.NoteModel) {
	stmt, err := db.Prepare(`INSERT INTO notes(title,content,author_id,book_id,record_time) 
							VALUES(?,?,?,?,?)`)
	checkErr(err)
	datetime := note.RecordTime.Format(time.RFC3339)
	_, err = stmt.Exec(note.Title, note.Content, note.Author_id, note.Book_id, datetime)
	checkErr(err)
}

func AllNotes() []models.NoteModel {
	notes := make([]models.NoteModel, 0)
	rows, err := db.Query("SELECT title,content,book_id,record_time FROM notes")
	checkErr(err)

	for rows.Next() {
		var note models.NoteModel
		var datetime string
		err = rows.Scan(&note.Title, &note.Content, &note.Book_id, &datetime)
		checkErr(err)
		note.RecordTime, err = time.Parse(time.RFC3339, datetime)
		checkErr(err)
		notes = append(notes,note)
	}

	return notes
}
