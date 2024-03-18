package main

import (
	"book_ex/server"
	"net/http"
)

func main() {

	db := server.Database{
		"socks": 1,
		"some":  2,
		"hello": 3,
		"haha":  4,
	}
	mux := http.NewServeMux()

	server.Routes(mux, db)

	http.ListenAndServe(":8000", mux)

}
