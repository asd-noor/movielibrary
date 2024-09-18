package domain

import (
	ws "movielibrary/internal/domain/enums/watchstatus"
	"movielibrary/internal/domain/models"
	"movielibrary/internal/domain/types"
)

// type ParserValidateFn func(obj interface{}) bool

type (
	movieType      = types.Movie
	movieTypeSlice = []types.Movie
)

type (
	movieModel      = models.Movie
	movieModelSlice = []models.Movie
)

type MovieLibraryUseCase interface {
	GetMovie(movieID int) types.Result[movieType]
	GetMovieList(filterBy ws.WatchStatus) types.Result[movieTypeSlice]
	GetAllMovies() types.Result[movieTypeSlice]

	AddToWatchList(movie movieType) types.Result[movieType]
	MarkAsWatching(movieID int) types.Result[movieType]
	MarkAsWatched(movieID int) types.Result[movieType]

	FetchMovieInfo(movie movieType) types.Result[movieType]
	RemoveMovie(movieID int) error
}

type MovieLibraryRepo interface {
	UpsertMovie(movie movieType) types.Result[movieModel]

	GetMovie(movieID int) types.Result[movieModel]
	GetMovies(filter ws.WatchStatus) types.Result[movieModelSlice]
	GetAllMovies() types.Result[movieModelSlice]

	RemoveMovie(movieID int) error
}
