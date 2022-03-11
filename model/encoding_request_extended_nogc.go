package model

import (
	"github.com/moontrade/dtc-go/message"
)

type EncodingRequestExtendedPointer struct {
	message.FixedPointer
}

type EncodingRequestExtendedPointerBuilder struct {
	message.FixedPointer
}

func AllocEncodingRequestExtended() EncodingRequestExtendedPointerBuilder {
	a := EncodingRequestExtendedPointerBuilder{message.AllocFixed(20)}
	a.Ptr.SetBytes(0, _EncodingRequestExtendedDefault)
	return a
}

func AllocEncodingRequestExtendedFrom(b []byte) EncodingRequestExtendedPointer {
	return EncodingRequestExtendedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size            = EncodingRequestExtendedSize  (20)
//     Type            = ENCODING_REQUEST  (6)
//     ProtocolVersion = CURRENT_VERSION  (8)
//     Encoding        = BINARY_ENCODING  (0)
//     ProtocolType    = ""
//     UseCompression  = USE_COMPRESSION_DISABLED  (0)
// }
func (m EncodingRequestExtendedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _EncodingRequestExtendedDefault)
}

// ToBuilder
func (m EncodingRequestExtendedPointer) ToBuilder() EncodingRequestExtendedPointerBuilder {
	return EncodingRequestExtendedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *EncodingRequestExtendedPointerBuilder) Finish() EncodingRequestExtendedPointer {
	return EncodingRequestExtendedPointer{m.FixedPointer}
}

// ProtocolVersion
func (m EncodingRequestExtendedPointer) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// ProtocolVersion
func (m EncodingRequestExtendedPointerBuilder) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Encoding
func (m EncodingRequestExtendedPointer) Encoding() EncodingEnum {
	return EncodingEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// Encoding
func (m EncodingRequestExtendedPointerBuilder) Encoding() EncodingEnum {
	return EncodingEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// ProtocolType
func (m EncodingRequestExtendedPointer) ProtocolType() string {
	return message.StringFixed(m.Ptr, 16, 12)
}

// ProtocolType
func (m EncodingRequestExtendedPointerBuilder) ProtocolType() string {
	return message.StringFixed(m.Ptr, 16, 12)
}

// UseCompression
func (m EncodingRequestExtendedPointer) UseCompression() UseCompressionEnum {
	return UseCompressionEnum(message.Int32Fixed(m.Ptr, 20, 16))
}

// UseCompression
func (m EncodingRequestExtendedPointerBuilder) UseCompression() UseCompressionEnum {
	return UseCompressionEnum(message.Int32Fixed(m.Ptr, 20, 16))
}

// SetProtocolVersion
func (m EncodingRequestExtendedPointerBuilder) SetProtocolVersion(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetEncoding
func (m EncodingRequestExtendedPointerBuilder) SetEncoding(value EncodingEnum) {
	message.SetInt32Fixed(m.Ptr, 12, 8, int32(value))
}

// SetProtocolType
func (m EncodingRequestExtendedPointerBuilder) SetProtocolType(value string) {
	message.SetStringFixed(m.Ptr, 16, 12, value)
}

// SetUseCompression
func (m EncodingRequestExtendedPointerBuilder) SetUseCompression(value UseCompressionEnum) {
	message.SetInt32Fixed(m.Ptr, 20, 16, int32(value))
}

// Copy
func (m EncodingRequestExtendedPointer) Copy(to EncodingRequestExtendedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
	to.SetUseCompression(m.UseCompression())
}

// Copy
func (m EncodingRequestExtendedPointerBuilder) Copy(to EncodingRequestExtendedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
	to.SetUseCompression(m.UseCompression())
}

// CopyPointer
func (m EncodingRequestExtendedPointer) CopyPointer(to EncodingRequestExtendedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
	to.SetUseCompression(m.UseCompression())
}

// CopyPointer
func (m EncodingRequestExtendedPointerBuilder) CopyPointer(to EncodingRequestExtendedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
	to.SetUseCompression(m.UseCompression())
}
