package model

import (
	"github.com/moontrade/dtc-go/message"
)

type EncodingRequestFactory interface {
	New() EncodingRequestBuilder

	Alloc() EncodingRequestBuilder
}

type EncodingRequest interface {
	message.Message

	ToBuilder() EncodingRequestBuilder

	ProtocolVersion() int32

	Encoding() int32

	ProtocolType() string
}

type EncodingRequestBuilder interface {
	message.Builder

	EncodingRequest

	Finish() EncodingRequest

	SetProtocolVersion(value int32)

	SetEncoding(value int32)

	SetProtocolType(value string)
}
