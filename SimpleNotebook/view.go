package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"
	"github.com/CodeOfThunder/SimpleNotebook/db"
	"github.com/CodeOfThunder/SimpleNotebook/models"
	"time"
	"strconv"
)

func htmlRender(w http.ResponseWriter,temp_name string, temp_path string) {
	t, err := template.New(temp_name).ParseFiles(temp_path)
	if err != nil {
		log.Fatal("cannot create template: ",temp_path)
	}
	t.Execute(w,"");
}

func home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}

func note(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if !checkAlreadyLogin(r) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		t,err := template.ParseFiles("static/note.html", "static/bar.html")
		if err != nil {
			log.Fatal(err.Error())
		}
		userName,_ := getUserCookie(r)

		var NotepageData models.NotepageModel

		NotepageData.Notes = db.AllNotes()
		NotepageData.Books = db.AllBooksOfUser(userName)

		err = t.Execute(w, NotepageData)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		var new_note models.NoteModel
		new_note.Book_id,_ = strconv.ParseInt(r.Form["book"][0], 10, 64)
		new_note.Author_id = 1
		new_note.Title = r.Form["title"][0]
		new_note.Content = r.Form["content"][0]
		new_note.RecordTime = time.Now()
		db.AddNote(new_note)
		http.Redirect(w, r, "/note", http.StatusFound)
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hehe,world!\n"))
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if checkAlreadyLogin(r) {
			http.Redirect(w, r, "/note", http.StatusFound)
			return
		}
		t,err := template.ParseFiles("static/register.html")
		if err != nil {
			log.Fatal(err.Error())
		}
		users := db.AllUsers()
		err = t.Execute(w, users)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		db.InsertUser(r.Form["username"][0], r.Form["password"][0])
		http.Redirect(w,r,"/login",http.StatusFound)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if checkAlreadyLogin(r) {
			http.Redirect(w,r,"/note",http.StatusFound)
			return
		}
		t, err := template.ParseFiles("static/login.html","static/bar.html")
		if err != nil {
			log.Fatal(err.Error())
		}
		getUserCookie(r)
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		if db.VerifyUserLogin(r.Form["username"][0],r.Form["password"][0]) {
			fmt.Println("login success")
			addUserCookie(w, r.Form["username"][0])
			http.Redirect(w,r,"/note",http.StatusFound)
		} else {
			fmt.Println("login failed")
		}
	}
}

func books(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path[1:])
	if r.Method == "GET" {
		if !checkAlreadyLogin(r) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return 
		}
		t, err := template.ParseFiles("static/books.html", "static/bar.html",)
		if err != nil {
			log.Fatal(err.Error())
		}
		books := db.AllBooks()
		err = t.Execute(w, books)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else if r.Method == "POST" {
		r.ParseForm()

		userName,_ := getUserCookie(r)
		userId := db.SelectUserId(userName)

		book := models.BookModel {
			BookName: r.Form["bookname"][0],
			Owner_id: userId,
			}
		db.AddBook(book)
		http.Redirect(w,r,"/books",http.StatusFound)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("static/bar.html")
		if err != nil {
			log.Fatal(err.Error())
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		deleteUserCookie(w, r)
		http.Redirect(w, r, "/home", http.StatusFound)
	}
}

func checkAlreadyLogin(r *http.Request) (bool) {
	_, alreadyLogin := getUserCookie(r)
	return alreadyLogin	
}

func initViews() {
	http.HandleFunc("/", home)
	http.HandleFunc("/note", note)
	http.HandleFunc("/register",register)
	http.HandleFunc("/login",login)
	http.HandleFunc("/books",books)
	http.HandleFunc("/logout", logout)
}
