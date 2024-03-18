package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Database map[string]int

func (db Database) List(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("server/templates/base.gohtml", "server/templates/nav.gohtml")
	if err != nil {
		//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	err = tmpl.Execute(w, db)
	if err != nil {
		//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
}

func (db Database) Price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "%d\n", price)

}

func (db Database) Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		item := r.FormValue("item")
		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		db[item] = price
		http.Redirect(w, r, "/list", http.StatusSeeOther)
		return
	}

	t, err := template.ParseFiles("server/templates/add.html", "server/templates/nav.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (db Database) Update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if item == "" {
		fmt.Fprint(w, "No such item")
		return
	}

	price, err := strconv.Atoi(r.URL.Query().Get("price"))
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	if price < 0 {
		fmt.Fprint(w, "Incorrent price")
		return
	}
	//http://localhost:8000/update?item=some&price=12
	db[item] = price
	fmt.Fprintf(w, "Add item %q with price %d", item, price)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Home page!")
}
