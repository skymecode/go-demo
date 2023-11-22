package foo

import "errors"

var ErrCouldOpen = errors.New("Could not open!")

func Open() error {
	return ErrCouldOpen
}
