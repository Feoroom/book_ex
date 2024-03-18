package server

import "net/http"

func Routes(mux *http.ServeMux, db Database) {

	fileServer := http.FileServer(http.Dir("./server/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/list", db.List)
	mux.HandleFunc("/price", db.Price)
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/add", db.Add)
}
