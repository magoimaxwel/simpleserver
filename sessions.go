package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
	"log"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secrectpassword"))
var email string = "magoimaxwel@gmail.com"

func secret(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	
	log.Println(req.Form)
	fmt.Println(req.FormValue("email"))



	fmt.Fprintf(w, "what the hell")
	fmt.Println(req.Header["Cookie"])
	fmt.Println("[==============================================]")
	fmt.Println(req.Cookie("session_cookie"))

	fmt.Println("[=============================]")

	length := req.ContentLength
	body_Content := make([]byte, length)

	req.Body.Read(body_Content)
	fmt.Println(string(body_Content))

	session, _ := store.Get(req, "session_cookie")

	aemail := session.Values["email"]
	fmt.Println(aemail)

	if aemail == email {
		fmt.Fprintf(w, "authenitcated")
	}

	else {
		fmt.Fprintf(w, "404")
	}

}

func login(w http.ResponseWriter, req *http.Request) {

	session, _ := store.Get(req, "session_cookie")

	session.Values["email"] = email
	session.Save(req, w)

	io.WriteString(w, `
	<!DOCTYPE html>
	<html>
		<body>
			<form action = "https://localhost:7076/secret" method = "post">
				<input type = "email" name = "email">
				<input type = "text" name = "username">
				<input type = "submit">
			
			</form>
		
		</body>

	
	`)
}

func main() {

	fmt.Println("[===] starting server")
	r := mux.NewRouter()
	r.HandleFunc("/", login)
	r.HandleFunc("/secret", secret)
	http.Handle("/", r)

	server := http.Server{
		Addr:         "localhost:7076",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server.ListenAndServeTLS("cert.pem", "key.pem")
}
