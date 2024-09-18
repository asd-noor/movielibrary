package service

import (
	"net/http"
	"strconv"
	"movielibrary/client"
	"movielibrary/internal/domain"
	de "movielibrary/internal/domain/errors"
	"movielibrary/internal/domain/models"
	"movielibrary/internal/domain/types"
	pt "movielibrary/pkg/types"

	ws "movielibrary/internal/domain/enums/watchstatus"
)

type MovieLibraryService struct {
	repo  domain.MovieLibraryRepo
	omdbc client.OmdbClient
}

func NewMovieLibraryService(r domain.MovieLibraryRepo, o client.OmdbClient) domain.MovieLibraryUseCase {
	return MovieLibraryService{
		omdbc: o,
		repo:  r,
	}
}

func (s MovieLibraryService) GetMovie(movieID int) types.Result[types.Movie] {
	result := s.repo.GetMovie(movieID)
	m, e := result.Unwrap()
	if e != nil {
		return *types.NewResult(types.Movie{}).SetErr(e)
	}

	return *types.NewResult(movieModelToType(m))
}

func (s MovieLibraryService) GetMovieList(filterBy ws.WatchStatus) types.Result[[]types.Movie] {
	result := s.repo.GetMovies(filterBy)
	ms, e := result.Unwrap()
	if e != nil {
		return *types.NewResult([]types.Movie{}).SetErr(e)
	}

	movies := make([]types.Movie, len(ms))
	for i := range ms {
		movies = append(movies, movieModelToType(ms[i]))
	}

	return *types.NewResult(movies)
}

func (s MovieLibraryService) GetAllMovies() types.Result[[]types.Movie] {
	result := s.repo.GetAllMovies()
	ms, e := result.Unwrap()
	if e != nil {
		return *types.NewResult([]types.Movie{}).SetErr(e)
	}

	movies := make([]types.Movie, len(ms))
	for i := range ms {
		movies = append(movies, movieModelToType(ms[i]))
	}

	return *types.NewResult(movies)
}

func (s MovieLibraryService) AddToWatchList(movie types.Movie) types.Result[types.Movie] {
	result := s.repo.UpsertMovie(movie)
	m, e := result.Unwrap()
	if e != nil {
		return *types.NewResult(types.Movie{}).SetErr(e)
	}

	return *types.NewResult(movieModelToType(m))
}

func (s MovieLibraryService) MarkAsWatching(movieID int) types.Result[types.Movie] {
	retval := types.Movie{}

	result := s.repo.UpsertMovie(types.Movie{ID: movieID, Status: ws.Watching})
	m, e := result.Unwrap()
	if e != nil {
		return *types.NewResult(retval).SetErr(e)
	}

	return *types.NewResult(movieModelToType(m))
}

func (s MovieLibraryService) MarkAsWatched(movieID int) types.Result[types.Movie] {
	result := s.repo.UpsertMovie(types.Movie{ID: movieID, Status: ws.Watched})
	m, e := result.Unwrap()
	if e != nil {
		return *types.NewResult(types.Movie{}).SetErr(e)
	}

	return *types.NewResult(movieModelToType(m))
}

func (s MovieLibraryService) FetchMovieInfo(movie types.Movie) types.Result[types.Movie] {
	req := pt.HttpRequest{
		Method:   http.MethodGet,
		Endpoint: "",
		QueryParams: map[string]string{
			"s":    movie.Title,
			"year": strconv.Itoa(movie.Year),
			"type": "movie",
		},
	}

	res, err := s.omdbc.Request(req)
	if err != nil {
		return *types.NewResult(movie).SetErr(err)
	}

	if res.StatusCode() != http.StatusOK {
		return *types.NewResult(movie).SetErr(de.ErrExpectationFailed)
	}

	omdbResult := res.Result().(pt.OmdbResponse)
	for _, v := range omdbResult.Search {
		if v.Year == strconv.Itoa(movie.Year) {
			movie.Title = v.Title

			if v.ImdbID != "" {
				movie.ImdbID = *types.Some(v.ImdbID)
			} else {
				movie.ImdbID = *types.None[string]()
			}

			break
		}
	}

	return *types.NewResult(movie)
}

func (s MovieLibraryService) RemoveMovie(movieID int) error {
	return s.repo.RemoveMovie(movieID)
}

func movieModelToType(mm models.Movie) (tm types.Movie) {
	tm.ID = mm.ID
	tm.Title = mm.Title
	tm.Year = mm.Year
	tm.Status = ws.WatchStatus(mm.Status)

	switch mm.ImdbID {
	case "":
		tm.ImdbID = *types.None[string]()
	default:
		tm.ImdbID = *types.Some(mm.ImdbID)
	}

	return
}
