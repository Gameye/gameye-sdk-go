package models

import "errors"

/*
ErrUnknownType occurs when an unknown type is encounterd during
decoding of a model
*/
var ErrUnknownType = errors.New("invalid type")
