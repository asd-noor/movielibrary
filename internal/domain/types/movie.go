package types

import (
	"encoding/json"
	"errors"
	ws "movielibrary/internal/domain/enums/watchstatus"
)

type Movie struct {
	Title  string         `json:"title"`
	ImdbID Option[string] `json:"ImdbID"`
	Status ws.WatchStatus `json:"status"`
	Year   int            `json:"year"`
	ID     int            `json:"id"`
	Liked  bool           `json:"liked"`
}

func (m Movie) Validate(validators ...Validator[Movie]) error {
	for _, v := range validators {
		if !v.Fn() {
			return errors.New(v.ErrMsg)
		}
	}

	return nil
}

func (m Movie) CheckYear() Validator[Movie] {
	return Validator[Movie]{
		Fn:     func() bool { return m.Year > 1900 && m.Year < 2100 },
		ErrMsg: "invalid year",
	}
}

func (m Movie) CheckTitle() Validator[Movie] {
	return Validator[Movie]{
		Fn:     func() bool { return m.Title != "" },
		ErrMsg: "title cannot be empty",
	}
}

func (m Movie) CheckImdbID() Validator[Movie] {
	return Validator[Movie]{
		Fn:     func() bool { return !m.ImdbID.IsNone() },
		ErrMsg: "imdb id cannot be empty",
	}
}

func (m Movie) MarshalJSON() ([]byte, error) {
	byteSlice, err := json.Marshal(m)
	if err != nil {
		return []byte{}, err
	}

	return byteSlice, nil
}

func (m *Movie) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, m); err != nil {
		return err
	}

	return nil
}
