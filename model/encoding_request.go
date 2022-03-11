package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size            = EncodingRequestSize  (16)
//     Type            = ENCODING_REQUEST  (6)
//     ProtocolVersion = CURRENT_VERSION  (8)
//     Encoding        = BINARY_ENCODING  (0)
//     ProtocolType    = "DTC"
// }
var _EncodingRequestDefault = []byte{16, 0, 6, 0, 8, 0, 0, 0, 0, 0, 0, 0, 68, 84, 67, 0}

const EncodingRequestSize = 16

// EncodingRequest Requirements: Not required for Servers. Required for Clients if the Client
// needs to discover the encoding the Server uses.
//
// The EncodingRequest message is a message requesting to change the DTC
// encoding for messages.
//
// For the procedure to work with this message, refer to Encoding Request
// Sequence.
type EncodingRequest struct {
	message.Fixed
}

// EncodingRequestBuilder Requirements: Not required for Servers. Required for Clients if the Client
// needs to discover the encoding the Server uses.
//
// The EncodingRequest message is a message requesting to change the DTC
// encoding for messages.
//
// For the procedure to work with this message, refer to Encoding Request
// Sequence.
type EncodingRequestBuilder struct {
	message.Fixed
}

func NewEncodingRequestFrom(b []byte) EncodingRequest {
	return EncodingRequest{Fixed: message.NewFixedFrom(b)}
}

func WrapEncodingRequest(b []byte) EncodingRequest {
	return EncodingRequest{Fixed: message.WrapFixed(b)}
}

func NewEncodingRequest() EncodingRequestBuilder {
	a := EncodingRequestBuilder{message.NewFixed(16)}
	a.Unsafe().SetBytes(0, _EncodingRequestDefault)
	return a
}

// Clear
// {
//     Size            = EncodingRequestSize  (16)
//     Type            = ENCODING_REQUEST  (6)
//     ProtocolVersion = CURRENT_VERSION  (8)
//     Encoding        = BINARY_ENCODING  (0)
//     ProtocolType    = "DTC"
// }
func (m EncodingRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _EncodingRequestDefault)
}

// ToBuilder
func (m EncodingRequest) ToBuilder() EncodingRequestBuilder {
	return EncodingRequestBuilder{m.Fixed}
}

// Finish
func (m EncodingRequestBuilder) Finish() EncodingRequest {
	return EncodingRequest{m.Fixed}
}

// ProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m EncodingRequest) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// ProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m EncodingRequestBuilder) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// Encoding The DTC message encoding the Client is requesting the Server to use.
func (m EncodingRequest) Encoding() EncodingEnum {
	return EncodingEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Encoding The DTC message encoding the Client is requesting the Server to use.
func (m EncodingRequestBuilder) Encoding() EncodingEnum {
	return EncodingEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// ProtocolType The ProtocolType field needs to be set to the text string "DTC".
//
// This field is automatically set with the binary encoding data structures.
// This field is automatically set with the binary encoding data structures.
//
// This field is used for the Server to know that it is communicating with
// a DTC compliant Client.
func (m EncodingRequest) ProtocolType() string {
	return message.StringFixed(m.Unsafe(), 16, 12)
}

// ProtocolType The ProtocolType field needs to be set to the text string "DTC".
//
// This field is automatically set with the binary encoding data structures.
// This field is automatically set with the binary encoding data structures.
//
// This field is used for the Server to know that it is communicating with
// a DTC compliant Client.
func (m EncodingRequestBuilder) ProtocolType() string {
	return message.StringFixed(m.Unsafe(), 16, 12)
}

// SetProtocolVersion The protocol version supported by the Client. Automatically set by constructor.
// The protocol version supported by the Client. Automatically set by constructor.
func (m EncodingRequestBuilder) SetProtocolVersion(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetEncoding The DTC message encoding the Client is requesting the Server to use.
func (m EncodingRequestBuilder) SetEncoding(value EncodingEnum) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, int32(value))
}

// SetProtocolType The ProtocolType field needs to be set to the text string "DTC".
//
// This field is automatically set with the binary encoding data structures.
// This field is automatically set with the binary encoding data structures.
//
// This field is used for the Server to know that it is communicating with
// a DTC compliant Client.
func (m EncodingRequestBuilder) SetProtocolType(value string) {
	message.SetStringFixed(m.Unsafe(), 16, 12, value)
}

// Copy
func (m EncodingRequest) Copy(to EncodingRequestBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
}

// Copy
func (m EncodingRequestBuilder) Copy(to EncodingRequestBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
}

// CopyPointer
func (m EncodingRequest) CopyPointer(to EncodingRequestPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
}

// CopyPointer
func (m EncodingRequestBuilder) CopyPointer(to EncodingRequestPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
}
