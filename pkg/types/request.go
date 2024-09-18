package types

type GetMoviesReq struct {
	FilterBy string `query:"filter"`
}

type AddMovieReq struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

type ModifyMovieReq struct {
	ID     int `json:"id" param:"id"`
	Status int `json:"status"`
}

type DeleteMovieReq struct {
	ID int `param:"id"`
}

type FetchMovieInfoReq struct {
	Title string `query:"s"`
	Year  int    `query:"year"`
}
