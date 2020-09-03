/*

	url general form
		scheme://[userinfo@]host/path[?query][#fragment]


		r := req.Header["Accept_Encoding"]

		request body



*/


package main

import(
	"fmt"
	"log"
)

func search_Page(w http.ResponseWriter, req *http.Request){

	req.PostForm()
	//parse file

	fmt.Println()


}