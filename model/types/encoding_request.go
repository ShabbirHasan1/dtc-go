package types

import (
	dtc "github.com/moontrade/dtc-go"
)

type EncodingRequestFactory interface {
	New() EncodingRequestBuilder

	Alloc() EncodingRequestBuilder
}

type EncodingRequest interface {
	dtc.Message

	ToBuilder() EncodingRequestBuilder

	ProtocolVersion() int32

	Encoding() int32

	ProtocolType() string
}

type EncodingRequestBuilder interface {
	dtc.MessageBuilder
	EncodingRequest

	Finish() EncodingRequest

	SetProtocolVersion(value int32)

	SetEncoding(value int32)

	SetProtocolType(value string)
}
