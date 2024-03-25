package web

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func (app *Application) Routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.SessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/",
		dynamic.ThenFunc(app.Home))

	router.Handler(http.MethodGet, "/list",
		dynamic.ThenFunc(app.List))

	router.Handler(http.MethodGet, "/add",
		dynamic.ThenFunc(app.Add))

	router.Handler(http.MethodGet, "/list/:id",
		dynamic.ThenFunc(app.View))

	router.Handler(http.MethodPost, "/add",
		dynamic.ThenFunc(app.AddPost))

	//router.HandlerFunc(http.MethodGet, "/", app.Home)
	//router.HandlerFunc(http.MethodGet, "/list", app.List)
	//router.HandlerFunc(http.MethodGet, "/add", app.Add)
	//router.HandlerFunc(http.MethodGet, "/list/:id", app.View)
	//router.HandlerFunc(http.MethodPost, "/add", app.AddPost)

	return alice.New(app.recoverPanic, app.logRequest, secureHeaders).Then(router)
}
