package web

import (
	"book_ex/internal/models"
	"book_ex/internal/validator"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *Application) Create(w http.ResponseWriter, r *http.Request) {
	data := app.NewTemplateData(r)
	data.Form = ReviewCreateForm{}
	app.render(w, http.StatusOK, "create.gohtml", data)
}

func (app *Application) CreatePost(w http.ResponseWriter, r *http.Request) {

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
		app.render(w, http.StatusUnprocessableEntity, "create.gohtml", data)
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

func (app *Application) userSignup(w http.ResponseWriter, r *http.Request) {

	data := app.NewTemplateData(r)
	data.Form = UserSignupForm{}
	app.render(w, http.StatusOK, "signup.gohtml", data)
}

func (app *Application) userSignupPost(w http.ResponseWriter, r *http.Request) {
	var form UserSignupForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	app.InfoLog.Println(form.Name, form.Password, form.Email)

	form.CheckFields(validator.NotBlank(form.Name), "name", "Это поле не может быть пустым")
	form.CheckFields(validator.NotBlank(form.Email), "email", "Это поле не может быть пустым")
	form.CheckFields(validator.Matches(form.Email, validator.EmailRX), "email", "Введен некорректный e-mail")
	form.CheckFields(validator.NotBlank(form.Password), "password", "Это поле не может быть пустым")
	form.CheckFields(validator.MinChars(form.Password, 8), "password", "Пароль должен быть не меньше 8 символов")

	if !form.Valid() {
		data := app.NewTemplateData(r)
		data.Form = form
		app.InfoLog.Println(form.FieldErrors)
		app.render(w, http.StatusUnprocessableEntity, "signup.gohtml", data)
		return
	}

	err = app.Users.Insert(form.Name, form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.AddFieldError("email", "Этот e-mail уже используется")

			data := app.NewTemplateData(r)
			data.Form = form
			app.render(w, http.StatusUnprocessableEntity, "signup.gohtml", data)
		} else {
			app.serverError(w, err)
		}

		return
	}

	app.SessionManager.Put(r.Context(), "flash", "Регистрация завершена")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *Application) userLogin(w http.ResponseWriter, r *http.Request) {

	data := app.NewTemplateData(r)
	data.Form = UserLoginForm{}
	app.render(w, http.StatusOK, "login.gohtml", data)
}

func (app *Application) userLoginPost(w http.ResponseWriter, r *http.Request) {
}

func (app *Application) userLogoutPost(w http.ResponseWriter, r *http.Request) {
}
