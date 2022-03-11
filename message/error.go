package message

import "errors"

var (
	ErrJSONCompactNotAvailable = errors.New("json compact not available")
	ErrJSONCompactDetected     = errors.New("json compact detected")
	ErrWrongType               = errors.New("wrong type")
	ErrOutOfMemory             = errors.New("out of memory")
	ErrProtobufUnavailable     = errors.New("protobuf unavailable")
)
