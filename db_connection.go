package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var secret_serverCode string = "vnknfxjlhv243Jmvc,mTlklx;8"

var sessions_store = sessions.NewCookieStore([]byte(secret_serverCode))

func initialisingDataBase() *sql.DB {

	db_driver := "mysql"
	db_username := "lambda"
	db_password := "deepspace"
	db_name := "goBlog"

	db_connection, err := sql.Open(db_driver, db_username+":"+db_password+"@/"+db_name)

	if err != nil {

		panic(err.Error())

	}

	//important setting section
	db_connection.SetConnMaxLifetime(time.Minute * 3)
	db_connection.SetMaxOpenConns(10)
	db_connection.SetMaxIdleConns(10)

	return db_connection

}

func PasswordHashing(password_submit string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(password_submit), 10)

	if err != nil {
		panic(err.Error())

	}

	return string(hash)

}

func check_password(username string, password string) bool {

	password_connection := initialisingDataBase()
	db_row, err := password_connection.Query("SELECT password FROM user WHERE  username = ?", username)

	if err != nil {

		panic(err.Error())
	}

	defer db_row.Close()
	var db_hashedPasswords string

	for db_row.Next() {

		err := db_row.Scan(&db_hashedPasswords)

		if err != nil {

			panic(err.Error())
		}
	}
	/*
		var db_hashedPasswords string
		scan_error := sel_dbPasswords.Scan.Next(&db_hashedPasswords)

		if scan_error != nil {

			panic(scan_error.Error())
		}
	*/

	correct_match := bcrypt.CompareHashAndPassword([]byte(db_hashedPasswords), []byte(password))

	var correct_password bool

	if correct_match != nil {

		correct_password = false
	} else {

		correct_password = true
	}

	return correct_password

}

func signIn(wrt http.ResponseWriter, req *http.Request) {

	parsed_template, err := template.ParseFiles("singin.html")

	if err != nil {

		panic(err.Error())
	}
	parsed_template.Execute(wrt, nil)
}

func registration(wrt http.ResponseWriter, req *http.Request) {

	req.ParseMultipartForm(1024)

	usernameReg := req.FormValue("username")
	passwordReg := req.FormValue("password")

	hashed_password := PasswordHashing(passwordReg)

	insert_crediential := initialisingDataBase()

	database, err := insert_crediential.Prepare("INSERT INTO user(username, password) values(?, ?)")

	if err != nil {

		panic(err.Error())
	}

	defer database.Close()

	res, err := database.Exec(usernameReg, hashed_password)

	if err != nil {

		panic(err.Error())
	}

	fmt.Println(res)

	http.Redirect(wrt, req, "https://localhost:8080", 301)

}

func login(wrt http.ResponseWriter, req *http.Request) {

	parsed_template, err := template.ParseFiles("login.html")

	if err != nil {

		panic(err.Error())
	}

	parsed_template.Execute(wrt, nil)

}

func login_request(wrt http.ResponseWriter, req *http.Request) {

	req.ParseMultipartForm(1024)

	username := req.FormValue("username")
	password := req.FormValue("password")

	login_status := check_password(username, password)

	if login_status {
		fmt.Println("[===] --correct password")
		session, err := sessions_store.Get(req, "session_cookie")

		if err != nil {

			panic(err.Error())
		}
		session.Values["username"] = username

		session.Save(req, wrt)

		http.Redirect(wrt, req, "https://"+username+"@localhost:8080/homepage", 301)
	}

}

func homepage(wrt http.ResponseWriter, req *http.Request) {

	parsed_template, err := template.ParseFiles("homepage.html")

	if err != nil {

		panic(err.Error())
	}

	username := req.URL.User
	parsed_template.Execute(wrt, username)

}

func log(h http.HandlerFunc) http.HandlerFunc {

	return func(wrt http.ResponseWriter, req *http.Request) {

		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("[===] -Handler called --->" + name)
		h(wrt, req)
	}
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", log(login))
	router.HandleFunc("/signIn", log(signIn))
	router.HandleFunc("/registration", log(registration))
	router.HandleFunc("/login_request", login_request)
	router.HandleFunc("/homepage", log(homepage))
	http.Handle("/", router)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
		//ReadTimeout:  15 * time.Second,
		//WriteTimeout: 15 * time.Second,
	}

	fmt.Println("[===] --starting server")
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
