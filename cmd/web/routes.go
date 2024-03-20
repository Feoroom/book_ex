package web

import (
	"net/http"
)

func (app *Application) Routes() *http.ServeMux {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./server/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/list", app.List)
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/add", app.Add)
	mux.HandleFunc("/view", app.View)

	return mux
}
