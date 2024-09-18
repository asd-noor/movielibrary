package utils

import (
	"errors"

	it "movielibrary/internal/domain/types"
	pt "movielibrary/pkg/types"

	ws "movielibrary/internal/domain/enums/watchstatus"
	pc "movielibrary/pkg/consts"
)

func ParseWatchStatusFromString(status string) (ws.WatchStatus, error) {
	switch status {
	case pc.MoviesToWatch:
		return ws.ToWatch, nil
	case pc.MoviesWatching:
		return ws.Watching, nil
	case pc.MoviesWatched:
		return ws.Watched, nil
	default:
		return ws.WatchStatus(0), errors.New("invalid watch status string")
	}
}

func PrepareMovieResponse(movie it.Movie) pt.MovieResp {
	return pt.MovieResp{
		ID:     movie.ID,
		Title:  movie.Title,
		Status: int(movie.Status),
		Year:   movie.Year,
		ImdbID: movie.ImdbID.GetOrDefault(""),
	}
}

func PrepareMovieListResponse(movies []it.Movie) pt.MovieListResp {
	return pt.MovieListResp{
		Movies: movies,
		Count:  len(movies),
	}
}

func PrepareErrResponse(err error, details string) pt.ErrResponse {
	return pt.ErrResponse{
		Details: details,
		Error:   err.Error(),
	}
}
