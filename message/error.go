package message

import "errors"

var (
	ErrJSONCompactNotAvailable = errors.New("json compact not available")
	ErrJSONCompactDetected     = errors.New("json compact detected")
	ErrWrongType               = errors.New("wrong type")
	ErrShortBuffer             = errors.New("short buffer")
	ErrOverflow                = errors.New("overflow")
	ErrBaseSizeOverflow        = errors.New("base size overflow")
	ErrOutOfMemory             = errors.New("out of memory")
	ErrProtobufUnavailable     = errors.New("protobuf unavailable")
	ErrShortWrite              = errors.New("short write")
	ErrFraming                 = errors.New("framing")
)
