package server

import "net/http"

func Routes(mux *http.ServeMux, db Database) {

	mux.HandleFunc("/list", db.List)
	mux.HandleFunc("/price", db.Price)
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/update", db.Update)
	mux.HandleFunc("/add", db.Add)
}
