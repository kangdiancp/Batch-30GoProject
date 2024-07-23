package router

import (
	"day13/HTTPServer03/internal/api"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string //url pattern
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"findAllCategoriesHandler", "GET",
		"/categories", api.FindAllCategoriesHandler,
	},
	Route{
		"findCategoryHandler", "GET",
		"/category/{id}", api.FindCategoryHandler,
	},
	Route{
		"AddCategoryHandler", "POST",
		"/categoryAdd", api.AddCategoryHandler,
	},
	Route{
		"DeleteCategoryHandler", "DELETE",
		"/category/{id}", api.DeleteCategoryHandler,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
