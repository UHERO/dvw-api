package routers

import (
	"github.com/UHERO/dvw-api/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
func CreateRouter(apiName string) *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	prefix := "/" + apiName

	apiRouter := mux.NewRouter().StrictSlash(false).PathPrefix(prefix).Subrouter()
	SetRoutes(apiRouter)
	router.PathPrefix(prefix).Handler(negroni.New(
		negroni.HandlerFunc(controllers.CORSHandler),
		//negroni.HandlerFunc(controllers.ValidApiKey(applicationRepository)),
		controllers.CheckCache(),
		negroni.Wrap(apiRouter),
	))
	return router
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func SetRoutes(r *mux.Router) {
	r.HandleFunc("/dimensions/{module:[A-Za-z]+}",					   controllers.GetModuleDimensions()).Methods("GET")
	r.HandleFunc("/{dimension:[a-z]+}/all/{module:[A-Za-z]+}",         controllers.GetDimensionAll()).Methods("GET")
	//   Following two endpoint URLs assume handles are unique per dimension, i.e., across modules. At present this is
	//   true, but it would be good to ensure before implementing them for real.
	r.HandleFunc("/{dimension:[a-z]+}/handle/{handle:[A-Za-z0-9]+}",   controllers.GetDimensionByHandle()).Methods("GET")
	r.HandleFunc("/{dimension:[a-z]+}/{handle:[A-Za-z0-9]+}/children", controllers.GetDimensionKidsByHandle()).Methods("GET")

	r.HandleFunc("/series/ade", controllers.GetAdeAirseatData("trend")).Methods("GET").Queries(
		"f", "{frequency:[AaQqMm]}",
		"i", "{i_list:[A-Za-z0-9,]+}",
		"m", "{m_list:[A-Za-z0-9,]+}",
		"d", "{d_list:[A-Za-z0-9,]+}",
	)
	r.HandleFunc("/series/ade", controllers.GetAdeAirseatData("trend")).Methods("GET").Queries(
		"f", "{frequency:[AaQqMm]}",
		"i", "{i_list:[A-Za-z0-9,]+}",
		"m", "{m_list:[A-Za-z0-9,]+}",
	)
	r.HandleFunc("/series/ade", controllers.GetAdeAirseatData("trend")).Methods("GET").Queries(
		"f", "{frequency:[AaQqMm]}",
		"i", "{i_list:[A-Za-z0-9,]+}",
		"d", "{d_list:[A-Za-z0-9,]+}",
	)

	r.HandleFunc("/series/airseat", controllers.GetAdeAirseatData("airseat")).Methods("GET").Queries(
		"f", "{frequency:[AaQqMm]}",
		"i", "{i_list:[A-Za-z0-9,]+}",
		"m", "{m_list:[A-Za-z0-9,]+}",
		"d", "{d_list:[A-Za-z0-9,]+}",
	)
	r.HandleFunc("/series/airseat", controllers.GetAdeAirseatData("airseat")).Methods("GET").Queries(
		"f", "{frequency:[AaQqMm]}",
		"i", "{i_list:[A-Za-z0-9,]+}",
		"m", "{m_list:[A-Za-z0-9,]+}",
	)
	r.HandleFunc("/series/airseat", controllers.GetAdeAirseatData("airseat")).Methods("GET").Queries(
		"f", "{frequency:[AaQqMm]}",
		"i", "{i_list:[A-Za-z0-9,]+}",
		"d", "{d_list:[A-Za-z0-9,]+}",
	)

	r.HandleFunc("/series/hotel", controllers.GetHotelData()).Methods("GET").Queries(
		"f", "{frequency:[AaQqMm]}",
		"i", "{i_list:[A-Za-z0-9,]+}",
		"c", "{c_list:[A-Za-z0-9,]+}",
	)

	r.HandleFunc("/series/char", controllers.GetCharacteristicData()).Methods("GET").Queries(
		"f", "{frequency:[AaQqMm]}",
		"i", "{i_list:[A-Za-z0-9,]+}",
		"g", "{g_list:[A-Za-z0-9,]+}",
	)

	r.HandleFunc("/series/exp", controllers.GetExpenditureData()).Methods("GET").Queries(
		"f", "{frequency:[AaQqMm]}",
		"i", "{i_list:[A-Za-z0-9,]+}",
		"g", "{g_list:[A-Za-z0-9,]+}",
		"c", "{c_list:[A-Za-z0-9,]+}",
	)
	r.HandleFunc("/series/exp", controllers.GetExpenditureData()).Methods("GET").Queries(
		"f", "{frequency:[AaQqMm]}",
		"i", "{i_list:[A-Za-z0-9,]+}",
		"g", "{g_list:[A-Za-z0-9,]+}",
	)
	r.HandleFunc("/series/exp", controllers.GetExpenditureData()).Methods("GET").Queries(
		"f", "{frequency:[AaQqMm]}",
		"i", "{i_list:[A-Za-z0-9,]+}",
		"c", "{c_list:[A-Za-z0-9,]+}",
	)

}
