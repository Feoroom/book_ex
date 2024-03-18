package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

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

	files := []string{
		"server/templates/base.gohtml",
		"server/templates/add.gohtml",
		"server/templates/partials/nav.gohtml",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (db Database) List(w http.ResponseWriter, r *http.Request) {

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

	files := []string{
		"server/templates/base.gohtml",
		"server/templates/list.gohtml",
		"server/templates/partials/nav.gohtml",
	}

	item := r.URL.Query().Get("item")
	if item != "" {

		files = []string{
			"server/templates/base.gohtml",
			"server/templates/item.gohtml",
			"server/templates/partials/nav.gohtml",
		}

		t, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = t.ExecuteTemplate(w, "base", item)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		return
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "base", db)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	//item := r.URL.Query().Get("item")
	//if item == "" {
	//	fmt.Fprint(w, "No such item")
	//	return
	//}
	//
	//price, err := strconv.Atoi(r.URL.Query().Get("price"))
	//if err != nil {
	//	fmt.Fprint(w, err)
	//	return
	//}
	//if price < 0 {
	//	fmt.Fprint(w, "Incorrent price")
	//	return
	//}
	////http://localhost:8000/update?item=some&price=12
	//db[item] = price
	//fmt.Fprintf(w, "Add item %q with price %d", item, price)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"server/templates/base.gohtml",
		"server/templates/home.gohtml",
		"server/templates/partials/nav.gohtml",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
