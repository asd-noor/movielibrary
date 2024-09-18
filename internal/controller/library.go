package controller

import (
	"net/http"
	"movielibrary/internal/domain"
	it "movielibrary/internal/domain/types"
	"movielibrary/internal/utils"
	pt "movielibrary/pkg/types"

	"github.com/labstack/echo/v4"
)

type MovieLibraryController struct {
	svc domain.MovieLibraryUseCase
}

func NewMovieLibraryController(svc domain.MovieLibraryUseCase) MovieLibraryController {
	return MovieLibraryController{
		svc: svc,
	}
}

func (ctr MovieLibraryController) GetMovieList(c echo.Context) error {
	req := pt.GetMoviesReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.PrepareErrResponse(err, ""))
	}

	watchStatus, err := utils.ParseWatchStatusFromString(req.FilterBy)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.PrepareErrResponse(err, ""))
	}

	movies, err := ctr.svc.GetMovieList(watchStatus).Unwrap()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.PrepareErrResponse(err, ""))
	}

	return c.JSON(http.StatusOK, utils.PrepareMovieListResponse(movies))
}

func (ctr MovieLibraryController) AddToWatchList(c echo.Context) error {
	req := pt.AddMovieReq{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.PrepareErrResponse(err, ""))
	}

	m := it.Movie{
		Title: req.Title,
		Year:  req.Year,
	}

	if err := m.Validate(m.CheckTitle(), m.CheckYear()); err != nil {
		return c.JSON(http.StatusBadRequest, utils.PrepareErrResponse(err, ""))
	}

	movie, err := ctr.svc.AddToWatchList(m).Unwrap()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.PrepareErrResponse(err, ""))
	}

	return c.JSON(http.StatusCreated, utils.PrepareMovieResponse(movie))
}

func (ctr MovieLibraryController) RemoveMovie(c echo.Context) error {
	req := pt.DeleteMovieReq{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.PrepareErrResponse(err, ""))
	}

	if err := ctr.svc.RemoveMovie(req.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.PrepareErrResponse(err, ""))
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{"success": true})
}

func (ctr MovieLibraryController) MarkAsWatched(c echo.Context) error {
	req := pt.ModifyMovieReq{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.PrepareErrResponse(err, ""))
	}

	movie, err := ctr.svc.MarkAsWatched(req.ID).Unwrap()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.PrepareErrResponse(err, ""))
	}

	return c.JSON(http.StatusOK, utils.PrepareMovieResponse(movie))
}

func (ctr MovieLibraryController) FetchMovieInfo(c echo.Context) error {
	req := pt.FetchMovieInfoReq{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.PrepareErrResponse(err, ""))
	}

	m := it.Movie{
		Title: req.Title,
		Year:  req.Year,
	}

	if err := m.Validate(m.CheckTitle(), m.CheckYear()); err != nil {
		return c.JSON(http.StatusBadRequest, utils.PrepareErrResponse(err, ""))
	}

	movie, err := ctr.svc.FetchMovieInfo(m).Unwrap()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.PrepareErrResponse(err, ""))
	}

	return c.JSON(http.StatusOK, utils.PrepareMovieResponse(movie))
}
