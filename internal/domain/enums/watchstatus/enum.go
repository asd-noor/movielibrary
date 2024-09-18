package watchstatus

import (
	"errors"
)

type WatchStatus int

const invalidString string = "invalid WatchStatus value"

const (
	ToWatch WatchStatus = iota + 1
	Watching
	Watched
)

var ws = map[WatchStatus]string{
	ToWatch:  "ToWatch",
	Watching: "Watching",
	Watched:  "Watched",
}

func (w WatchStatus) String() string {
	v, ok := ws[w]

	if !ok {
		return invalidString
	}

	return v
}

func (w WatchStatus) Check() (string, error) {
	v := w.String()
	if v == invalidString {
		return v, errors.New(invalidString)
	}
	return v, nil
}
