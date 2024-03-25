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

	router.HandlerFunc(http.MethodGet, "/", app.Home)
	router.HandlerFunc(http.MethodGet, "/list", app.List)
	router.HandlerFunc(http.MethodGet, "/add", app.Add)
	router.HandlerFunc(http.MethodGet, "/list/:id", app.View)
	router.HandlerFunc(http.MethodPost, "/add", app.AddPost)

	return alice.New(app.recoverPanic, app.logRequest, secureHeaders).Then(router)
}
