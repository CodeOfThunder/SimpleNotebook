package models

import (
	"time"
)

type UserModel struct {
	Uid int64
	UserName string
	Password string
}

type NoteModel struct {
	Nid int64
	Title string
	Content string
	Author_id int64
	Book_id int64
	RecordTime time.Time
}

type BookModel struct {
	Bid int64
	BookName string
	Owner_id int64
}

type NotepageModel struct {
	Notes []NoteModel
	Books []BookModel
}