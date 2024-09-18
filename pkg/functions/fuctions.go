package functions

import (
	"encoding/json"
)

func StructToStruct(in interface{}, out interface{}) error {
	if b, err := json.Marshal(in); err == nil {
		if err := json.Unmarshal(b, &out); err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

func InSlice[T comparable](needle T, haystack []T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
