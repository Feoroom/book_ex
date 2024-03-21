package web

import (
	"book_ex/internal/models"
	"errors"
	"html/template"
	"net/http"
	"strconv"
)

//func (app *Application) Price(w http.ResponseWriter, r *http.Request) {
//	item := r.URL.Query().Get("item")
//
//	//price, ok := app.DB[item]
//	if !ok {
//		w.WriteHeader(http.StatusNotFound)
//		app.InfoLog.Printf("No such item: ", item)
//		return
//	}
//
//	fmt.Fprintf(w, "%d\n", price)
//
//}

func (app *Application) Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		item := r.FormValue("item")
		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		id, err := app.Items.Insert(item, price)
		if err != nil {
			app.serverError(w, err)
			return
		}
		app.InfoLog.Printf("Id of inserted item is %d", id)
		http.Redirect(w, r, "/list", http.StatusSeeOther)
		return
	}

	files := []string{
		"ui/templates/base.gohtml",
		"ui/templates/add.gohtml",
		"ui/templates/partials/nav.gohtml",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
		return
	}

}

func (app *Application) List(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"ui/templates/base.gohtml",
		"ui/templates/list.gohtml",
		"ui/templates/partials/nav.gohtml",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	items, err := app.Items.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = t.ExecuteTemplate(w, "base", items)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *Application) View(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		app.notFound(w)
		return
	}

	item, err := app.Items.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	files := []string{
		"ui/templates/base.gohtml",
		"ui/templates/view_item.gohtml",
		"ui/templates/partials/nav.gohtml",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = t.ExecuteTemplate(w, "base", item)
	if err != nil {
		app.serverError(w, err)
		return
	}

	return
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"ui/templates/base.gohtml",
		"ui/templates/home.gohtml",
		"ui/templates/partials/nav.gohtml",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}
