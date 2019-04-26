package main

import (
	"net/http"
	"github.com/CodeOfThunder/SimpleNotebook/db"
)

func main() {
	initViews()
	db.ConnectDB()
	http.ListenAndServe(":8000",nil)
}