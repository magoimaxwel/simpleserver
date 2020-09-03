/*
	http supported by golang
	GET
	POST
	OPTION
	TRACE
	HEAD

	HTTP PROCTOL(consist of a new lines of the text in ther following order)
	--Request line
	--Zero or more request headers
	--An empty line
	--The message body

	{(typical HTTP request)
		GET /Protocols/rfc2616/rfc2616.html HTTP/1.1 -->request line
		Host: www.w3.org    -->header
		User-Agent: Mozilla/5.0 -->header

	}
GET-->tells the server to return the specified resource
HEAD --> same with GET ,except the server must not return a message body without the messge
POST -->tell the server that the data in the message body should be passed to the resource identified by the url
PUT -->
DELETE --> Tell the server to remove the resource identified by the URL
TRACE -->
OPTIONS -->Tells the server to return a list of HTTP method that the server supports
CONNECT -->Tells the server that the datan in the message body modified the resiource identifeid by the url

safe request --->request is safe if the request does not change the state of the server

idempotent request --->if the state of the server doesn`t change the second time the method is called with same data

Common HTTP request headers

Accept -->content type accepted by th client as part of the response
Accept-Charset -->tells the server which characterset it accpets
Authorization -->dend basic authenicatuion credential to the server
Cookie -->client send back cookie
Content-Length -->length of the request body in octets
Content-Type -->POST & PUT is set by default x-www-form-urlen-coded anf files as multipart/form-data
Host -->name of the server , along with then port
Referre -->adderes of the previous page that linked to the request page
User-Agent -->describes the calling client


-------HTTP Response

--A status line
--Zero or more ressponse headers
--An empty line
--The message body(optional)

example of the http response
{
	200 OK
	Date: Sat, 22 Nov 2014 12:58:58 GMT
	Server: Apache/2
	Last-Modified: Thu, 28 Aug 2014 21:01:33 GMT
	Content-Length: 33115
	Content-Type: text/html;
	charset=iso-8859-1

<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/ TR/xhtml1/DTD/xhtml1-strict.dtd"> <html xmlns='http://www.w3.org/1999/ xhtml'> <head><title>Hypertext Transfer Protocol -- HTTP/1.1</title></ head><body>…</body></html>

}

{(HTTP response status code)
	1xx -->tells the server received the request and is processing it
	2xx -->Success
	3xx -->Redirection
	4xx -->client error
	5xx -->Server error


}

{
	Allow --> tell the client which request n=methods are supported by the server
	Content-Length --> content of the response body on octet
	Content-Type -->content-type
	Date --. date
	Location --> header is used with redirection
	Server -->Domain name of the server
	Set-Cookie -->Sets a cookie at the client
	WWW-Authenticate -->
}

{
	URI --
}

Handling requests

Go/http library

[
	{
		(server struct)
		Addr string
		Handler handler
		ReadTimeout time.Duration
		WriteTimeout time.Duration
		MaxHeaderBytes int
		TLSConfig *tls.Config
		TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
		ConnState func(net.Conn, ConnState)
		ErrorLog *log.Logger
	}
]
*/

/*

package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/gorilla/mux"
)

func login(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<p>Ooooh shit</p>")
}

func homepage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<p>what the hell</p>")

}

func fileServer(w http.ResponseWriter, req *http.Request) {

	http.StripPrefix("/static/", http.FileServer(http.Dir("static")))

}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("[===] -Handler called --->" + name)
		h(w, req)
	}
}

func main() {

	/*
		r.PathPrefix()
		r.Methods("GET", "POST")
		r.Schemes("http")
		r.Headers("X-Requested-With", "XMLHttpRequest")
		r.Queries("key", "values  ")



	hammer := mux.NewRouter()
	rock := hammer.Host("localhost").Subrouter()

	rock.HandleFunc("/", log(login))
	rock.HandleFunc("/login", log(homepage))

	rock.HandleFunc("/static/", fileServer)
	http.Handle("/", hammer)

	fmt.Println("[===] -starting server on port8080")
	fmt.Println("[===]  -listening on the server")

	server := http.Server{
		Addr:         "localhost:7070",
		Handler:      hammer,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server.ListenAndServe()

}
*/