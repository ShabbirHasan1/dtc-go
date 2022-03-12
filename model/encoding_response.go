package model

import (
	"github.com/moontrade/dtc-go/message"
)

const EncodingResponseSize = 16

// {
//     Size            = EncodingResponseSize  (16)
//     Type            = ENCODING_RESPONSE  (7)
//     ProtocolVersion = CURRENT_VERSION  (8)
//     Encoding        = BINARY_ENCODING  (0)
//     ProtocolType    = "DTC"
// }
var _EncodingResponseDefault = []byte{16, 0, 7, 0, 8, 0, 0, 0, 0, 0, 0, 0, 68, 84, 67, 0}

// EncodingResponse Requirements: Required for Servers. Required for Clients if the Client
// needs to discover the encoding the Server uses.
//
// The EncodingResponse is a message from the Server to the Client, telling
// the Client what message encoding it must use to communicate with the Server.
// the Client what message encoding it must use to communicate with the Server.
//
// For the procedure to work with this message, refer to Encoding Request
// Sequence.
type EncodingResponse struct {
	message.Fixed
}

// EncodingResponseBuilder Requirements: Required for Servers. Required for Clients if the Client
// needs to discover the encoding the Server uses.
//
// The EncodingResponse is a message from the Server to the Client, telling
// the Client what message encoding it must use to communicate with the Server.
// the Client what message encoding it must use to communicate with the Server.
//
// For the procedure to work with this message, refer to Encoding Request
// Sequence.
type EncodingResponseBuilder struct {
	message.Fixed
}

// EncodingResponsePointer Requirements: Required for Servers. Required for Clients if the Client
// needs to discover the encoding the Server uses.
//
// The EncodingResponse is a message from the Server to the Client, telling
// the Client what message encoding it must use to communicate with the Server.
// the Client what message encoding it must use to communicate with the Server.
//
// For the procedure to work with this message, refer to Encoding Request
// Sequence.
type EncodingResponsePointer struct {
	message.FixedPointer
}

// EncodingResponsePointerBuilder Requirements: Required for Servers. Required for Clients if the Client
// needs to discover the encoding the Server uses.
//
// The EncodingResponse is a message from the Server to the Client, telling
// the Client what message encoding it must use to communicate with the Server.
// the Client what message encoding it must use to communicate with the Server.
//
// For the procedure to work with this message, refer to Encoding Request
// Sequence.
type EncodingResponsePointerBuilder struct {
	message.FixedPointer
}

func NewEncodingResponseFrom(b []byte) EncodingResponse {
	return EncodingResponse{Fixed: message.NewFixedFrom(b)}
}

func WrapEncodingResponse(b []byte) EncodingResponse {
	return EncodingResponse{Fixed: message.WrapFixed(b)}
}

func NewEncodingResponse() EncodingResponseBuilder {
	a := EncodingResponseBuilder{message.NewFixed(16)}
	a.Unsafe().SetBytes(0, _EncodingResponseDefault)
	return a
}

func AllocEncodingResponse() EncodingResponsePointerBuilder {
	a := EncodingResponsePointerBuilder{message.AllocFixed(16)}
	a.Ptr.SetBytes(0, _EncodingResponseDefault)
	return a
}

func AllocEncodingResponseFrom(b []byte) EncodingResponsePointer {
	return EncodingResponsePointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size            = EncodingResponseSize  (16)
//     Type            = ENCODING_RESPONSE  (7)
//     ProtocolVersion = CURRENT_VERSION  (8)
//     Encoding        = BINARY_ENCODING  (0)
//     ProtocolType    = "DTC"
// }
func (m EncodingResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _EncodingResponseDefault)
}

// Clear
// {
//     Size            = EncodingResponseSize  (16)
//     Type            = ENCODING_RESPONSE  (7)
//     ProtocolVersion = CURRENT_VERSION  (8)
//     Encoding        = BINARY_ENCODING  (0)
//     ProtocolType    = "DTC"
// }
func (m EncodingResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _EncodingResponseDefault)
}

// ToBuilder
func (m EncodingResponse) ToBuilder() EncodingResponseBuilder {
	return EncodingResponseBuilder{m.Fixed}
}

// ToBuilder
func (m EncodingResponsePointer) ToBuilder() EncodingResponsePointerBuilder {
	return EncodingResponsePointerBuilder{m.FixedPointer}
}

// Finish
func (m EncodingResponseBuilder) Finish() EncodingResponse {
	return EncodingResponse{m.Fixed}
}

// Finish
func (m *EncodingResponsePointerBuilder) Finish() EncodingResponsePointer {
	return EncodingResponsePointer{m.FixedPointer}
}

// ProtocolVersion This field is automatically set by the constructor.
func (m EncodingResponse) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// ProtocolVersion This field is automatically set by the constructor.
func (m EncodingResponseBuilder) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// ProtocolVersion This field is automatically set by the constructor.
func (m EncodingResponsePointer) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// ProtocolVersion This field is automatically set by the constructor.
func (m EncodingResponsePointerBuilder) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Encoding The DTC message encoding to be used.
//
// This value may be different from the requested DTC encoding if the Server
// does not support the requested encoding from the Client.
func (m EncodingResponse) Encoding() EncodingEnum {
	return EncodingEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Encoding The DTC message encoding to be used.
//
// This value may be different from the requested DTC encoding if the Server
// does not support the requested encoding from the Client.
func (m EncodingResponseBuilder) Encoding() EncodingEnum {
	return EncodingEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Encoding The DTC message encoding to be used.
//
// This value may be different from the requested DTC encoding if the Server
// does not support the requested encoding from the Client.
func (m EncodingResponsePointer) Encoding() EncodingEnum {
	return EncodingEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// Encoding The DTC message encoding to be used.
//
// This value may be different from the requested DTC encoding if the Server
// does not support the requested encoding from the Client.
func (m EncodingResponsePointerBuilder) Encoding() EncodingEnum {
	return EncodingEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// ProtocolType The ProtocolType field needs to be set to the text string "DTC".
//
// This field is automatically set with the binary encoding data structures.
// This field is automatically set with the binary encoding data structures.
//
// This field is used for the Client to know that it is communicating with
// a DTC compliant Server.
func (m EncodingResponse) ProtocolType() string {
	return message.StringFixed(m.Unsafe(), 16, 12)
}

// ProtocolType The ProtocolType field needs to be set to the text string "DTC".
//
// This field is automatically set with the binary encoding data structures.
// This field is automatically set with the binary encoding data structures.
//
// This field is used for the Client to know that it is communicating with
// a DTC compliant Server.
func (m EncodingResponseBuilder) ProtocolType() string {
	return message.StringFixed(m.Unsafe(), 16, 12)
}

// ProtocolType The ProtocolType field needs to be set to the text string "DTC".
//
// This field is automatically set with the binary encoding data structures.
// This field is automatically set with the binary encoding data structures.
//
// This field is used for the Client to know that it is communicating with
// a DTC compliant Server.
func (m EncodingResponsePointer) ProtocolType() string {
	return message.StringFixed(m.Ptr, 16, 12)
}

// ProtocolType The ProtocolType field needs to be set to the text string "DTC".
//
// This field is automatically set with the binary encoding data structures.
// This field is automatically set with the binary encoding data structures.
//
// This field is used for the Client to know that it is communicating with
// a DTC compliant Server.
func (m EncodingResponsePointerBuilder) ProtocolType() string {
	return message.StringFixed(m.Ptr, 16, 12)
}

// SetProtocolVersion This field is automatically set by the constructor.
func (m EncodingResponseBuilder) SetProtocolVersion(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetProtocolVersion This field is automatically set by the constructor.
func (m EncodingResponsePointerBuilder) SetProtocolVersion(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetEncoding The DTC message encoding to be used.
//
// This value may be different from the requested DTC encoding if the Server
// does not support the requested encoding from the Client.
func (m EncodingResponseBuilder) SetEncoding(value EncodingEnum) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, int32(value))
}

// SetEncoding The DTC message encoding to be used.
//
// This value may be different from the requested DTC encoding if the Server
// does not support the requested encoding from the Client.
func (m EncodingResponsePointerBuilder) SetEncoding(value EncodingEnum) {
	message.SetInt32Fixed(m.Ptr, 12, 8, int32(value))
}

// SetProtocolType The ProtocolType field needs to be set to the text string "DTC".
//
// This field is automatically set with the binary encoding data structures.
// This field is automatically set with the binary encoding data structures.
//
// This field is used for the Client to know that it is communicating with
// a DTC compliant Server.
func (m EncodingResponseBuilder) SetProtocolType(value string) {
	message.SetStringFixed(m.Unsafe(), 16, 12, value)
}

// SetProtocolType The ProtocolType field needs to be set to the text string "DTC".
//
// This field is automatically set with the binary encoding data structures.
// This field is automatically set with the binary encoding data structures.
//
// This field is used for the Client to know that it is communicating with
// a DTC compliant Server.
func (m EncodingResponsePointerBuilder) SetProtocolType(value string) {
	message.SetStringFixed(m.Ptr, 16, 12, value)
}

// Copy
func (m EncodingResponse) Copy(to EncodingResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
}

// Copy
func (m EncodingResponseBuilder) Copy(to EncodingResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
}

// CopyPointer
func (m EncodingResponse) CopyPointer(to EncodingResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
}

// CopyPointer
func (m EncodingResponseBuilder) CopyPointer(to EncodingResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
}

// Copy
func (m EncodingResponsePointer) Copy(to EncodingResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
}

// Copy
func (m EncodingResponsePointerBuilder) Copy(to EncodingResponseBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
}

// CopyPointer
func (m EncodingResponsePointer) CopyPointer(to EncodingResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
}

// CopyPointer
func (m EncodingResponsePointerBuilder) CopyPointer(to EncodingResponsePointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
}
