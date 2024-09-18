package routes

import (
	"net/http"
	"movielibrary/api"
	"movielibrary/internal/controller"
)

type LibraryRouter struct {
	ctr controller.MovieLibraryController
}

func NewLibraryRouter(ctr controller.MovieLibraryController) api.Router {
	return LibraryRouter{ctr}
}

func (r LibraryRouter) GetRoutes() []api.Route {
	return []api.Route{
		// gets watchlist
		{
			Path:    "/watchlist",
			Method:  http.MethodGet,
			Handler: nil,
		},
		// adds movie to watchlist
		{
			Path:    "/watchlist",
			Method:  http.MethodPost,
			Handler: nil,
		},
		// gets currently watching movies
		{
			Path:    "/watching",
			Method:  http.MethodGet,
			Handler: nil,
		},
		// modifies the status currently watching movies
		{
			Path:    "/watching/:id",
			Method:  http.MethodPut,
			Handler: nil,
		},
		// gets the watched movies
		{
			Path:    "/watched",
			Method:  http.MethodGet,
			Handler: nil,
		},
		// modifies the status of watched movies
		{
			Path:    "/watched/:id",
			Method:  http.MethodPatch,
			Handler: nil,
		},
		// fetches movie info from omdb
		{
			Path:   "/movie-info",
			Method: http.MethodGet,
		},
		// removes movie from database
		{
			Path:   "/movies",
			Method: http.MethodGet,
		},
	}
}
