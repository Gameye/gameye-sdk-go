package clients

import (
	"errors"
)

/*
ErrUnexpectedStatus occurs when an unexpects http status is encountered
*/
var ErrUnexpectedStatus = errors.New("unexpected status")
