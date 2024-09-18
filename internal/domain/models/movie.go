package models

import "encoding/json"

type Movie struct {
	Liked  *bool
	Title  string
	ImdbID string
	Status int
	Year   int
	ID     int `gorm:"primaryKey"`
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
