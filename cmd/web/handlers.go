package web

import (
	"book_ex/internal/models"
	"book_ex/internal/validator"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type ReviewCreateForm struct {
	Title               string `form:"title"`
	Text                string `form:"text"`
	validator.Validator `form:"-"`
}

func (app *Application) Add(w http.ResponseWriter, r *http.Request) {
	data := app.NewTemplateData(r)
	data.Form = ReviewCreateForm{}
	app.render(w, http.StatusOK, "add.gohtml", data)
}

func (app *Application) AddPost(w http.ResponseWriter, r *http.Request) {

	var form ReviewCreateForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	//Validation check
	form.CheckFields(validator.NotBlank(form.Title), "title", "Это поле не может быть пустым")
	form.CheckFields(validator.MaxChars(form.Title, 30), "title", "Длина этого поля должны быть не больше 30 символов")
	form.CheckFields(validator.NotBlank(form.Text), "text", "Это поле не может быть пустым")

	if !form.Valid() {
		app.InfoLog.Println(form.FieldErrors)
		data := app.NewTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "add.gohtml", data)
		return
	}

	id, err := app.Reviews.Insert(form.Title, form.Text)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.SessionManager.Put(r.Context(), "add", "Рецензия успешно создана!")

	app.InfoLog.Printf("Id of inserted item is %d", id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) List(w http.ResponseWriter, r *http.Request) {

	items, err := app.Items.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.NewTemplateData(r)
	data.Items = items
	app.render(w, http.StatusOK, "list.gohtml", data)
}

func (app *Application) View(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	app.InfoLog.Println(id)
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

	data := app.NewTemplateData(r)
	data.Item = item

	app.render(w, http.StatusOK, "view_item.gohtml", data)

	return
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {

	reviews, err := app.Reviews.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.NewTemplateData(r)
	data.Reviews = reviews
	//data.Flash = app.SessionManager.PopString(r.Context(), "add")
	app.render(w, http.StatusOK, "home.gohtml", data)
}

func (app *Application) Review(w http.ResponseWriter, r *http.Request) {

}
