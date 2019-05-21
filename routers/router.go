package routers

import (
	"github.com/UHERO/dbedt-rest-api/controllers"
	"github.com/UHERO/dbedt-rest-api/data"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func InitRoutes(cache *data.Cache) *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	apiRouter := mux.NewRouter().StrictSlash(false)
	apiRouter = SetRoutes(apiRouter)
	router.PathPrefix("/v1").Handler(negroni.New(
		negroni.HandlerFunc(controllers.CORSOptionsHandler),
		//negroni.HandlerFunc(controllers.ValidApiKey(applicationRepository)),
		negroni.HandlerFunc(controllers.CheckCache(cache)),
		negroni.Wrap(apiRouter),
	))
	return router
}
