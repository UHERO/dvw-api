package routers

import (
	"github.com/UHERO/dvw-api/controllers"
	"github.com/UHERO/dvw-api/data"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func InitRoutes(cache *data.Cache) *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	apiRouter := mux.NewRouter().StrictSlash(false)
	apiRouter.SetRoutes()
	router.PathPrefix("/dvw").Handler(negroni.New(
		negroni.HandlerFunc(controllers.CORSOptionsHandler),
		//negroni.HandlerFunc(controllers.ValidApiKey(applicationRepository)),
		negroni.HandlerFunc(controllers.CheckCache(cache)),
		negroni.Wrap(apiRouter),
	))
	return router
}

func (r *mux.Router) SetRoutes() {
	return
}
