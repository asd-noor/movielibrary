package types

import (
	it "movielibrary/internal/domain/types"
)

type MovieResp struct {
	Title  string `json:"title"`
	ImdbID string `json:"imdbID,omitempty"`
	Status int    `json:"status"`
	Year   int    `json:"year"`
	ID     int    `json:"id"`
}

type MovieListResp struct {
	Movies []it.Movie `json:"movies"`
	Count  int        `json:"count"`
}

type ErrResponse struct {
	Details string `json:"details,omitempty"`
	Error   string `json:"error"`
}
