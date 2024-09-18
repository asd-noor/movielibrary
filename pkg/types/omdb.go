package types

type OmdbResponse struct {
	TotalResults string      `json:"totalResults"`
	Response     string      `json:"Response"`
	Search       []OmdbMovie `json:"Search"`
}

type OmdbMovie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
	ImdbID string `json:"imdbID"`
}
