package main

import (
	"net/http"
)

var cookieName_user string = "notebook-username"

func addUserCookie(w http.ResponseWriter, userName string) {
	c := http.Cookie {
		Name:cookieName_user,
		Value:userName, }
	http.SetCookie(w, &c)
}

func getUserCookie(r *http.Request) (string,bool) {
	c, err := r.Cookie(cookieName_user)
	ret := false
	if err != nil {
		return "", ret
	}
	if len(c.Name) > 0 {
		ret = true;
		//fmt.Println("getUserCookie: " + c.Value)
	}
	return c.Value, ret
}

func deleteUserCookie(w http.ResponseWriter, r* http.Request) {
	c, err := r.Cookie(cookieName_user)
	if err != nil {
		return 
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
}