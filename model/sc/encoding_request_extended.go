package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const EncodingRequestExtendedSize = 20

// {
//     Size            = EncodingRequestExtendedSize  (20)
//     Type            = ENCODING_REQUEST  (6)
//     ProtocolVersion = CURRENT_VERSION  (8)
//     Encoding        = BINARY_ENCODING  (0)
//     ProtocolType    = ""
//     UseCompression  = USE_COMPRESSION_DISABLED  (0)
// }
var _EncodingRequestExtendedDefault = []byte{20, 0, 6, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type EncodingRequestExtended struct {
	message.Fixed
}

type EncodingRequestExtendedBuilder struct {
	message.Fixed
}

type EncodingRequestExtendedPointer struct {
	message.FixedPointer
}

type EncodingRequestExtendedPointerBuilder struct {
	message.FixedPointer
}

func NewEncodingRequestExtendedFrom(b []byte) EncodingRequestExtended {
	return EncodingRequestExtended{Fixed: message.NewFixedFrom(b)}
}

func WrapEncodingRequestExtended(b []byte) EncodingRequestExtended {
	return EncodingRequestExtended{Fixed: message.WrapFixed(b)}
}

func NewEncodingRequestExtended() EncodingRequestExtendedBuilder {
	a := EncodingRequestExtendedBuilder{message.NewFixed(20)}
	a.Unsafe().SetBytes(0, _EncodingRequestExtendedDefault)
	return a
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
func (m EncodingRequestExtendedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _EncodingRequestExtendedDefault)
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
func (m EncodingRequestExtended) ToBuilder() EncodingRequestExtendedBuilder {
	return EncodingRequestExtendedBuilder{m.Fixed}
}

// ToBuilder
func (m EncodingRequestExtendedPointer) ToBuilder() EncodingRequestExtendedPointerBuilder {
	return EncodingRequestExtendedPointerBuilder{m.FixedPointer}
}

// Finish
func (m EncodingRequestExtendedBuilder) Finish() EncodingRequestExtended {
	return EncodingRequestExtended{m.Fixed}
}

// Finish
func (m *EncodingRequestExtendedPointerBuilder) Finish() EncodingRequestExtendedPointer {
	return EncodingRequestExtendedPointer{m.FixedPointer}
}

// ProtocolVersion
func (m EncodingRequestExtended) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// ProtocolVersion
func (m EncodingRequestExtendedBuilder) ProtocolVersion() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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
func (m EncodingRequestExtended) Encoding() EncodingEnum {
	return EncodingEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// Encoding
func (m EncodingRequestExtendedBuilder) Encoding() EncodingEnum {
	return EncodingEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
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
func (m EncodingRequestExtended) ProtocolType() string {
	return message.StringFixed(m.Unsafe(), 16, 12)
}

// ProtocolType
func (m EncodingRequestExtendedBuilder) ProtocolType() string {
	return message.StringFixed(m.Unsafe(), 16, 12)
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
func (m EncodingRequestExtended) UseCompression() UseCompressionEnum {
	return UseCompressionEnum(message.Int32Fixed(m.Unsafe(), 20, 16))
}

// UseCompression
func (m EncodingRequestExtendedBuilder) UseCompression() UseCompressionEnum {
	return UseCompressionEnum(message.Int32Fixed(m.Unsafe(), 20, 16))
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
func (m EncodingRequestExtendedBuilder) SetProtocolVersion(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetProtocolVersion
func (m EncodingRequestExtendedPointerBuilder) SetProtocolVersion(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetEncoding
func (m EncodingRequestExtendedBuilder) SetEncoding(value EncodingEnum) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, int32(value))
}

// SetEncoding
func (m EncodingRequestExtendedPointerBuilder) SetEncoding(value EncodingEnum) {
	message.SetInt32Fixed(m.Ptr, 12, 8, int32(value))
}

// SetProtocolType
func (m EncodingRequestExtendedBuilder) SetProtocolType(value string) {
	message.SetStringFixed(m.Unsafe(), 16, 12, value)
}

// SetProtocolType
func (m EncodingRequestExtendedPointerBuilder) SetProtocolType(value string) {
	message.SetStringFixed(m.Ptr, 16, 12, value)
}

// SetUseCompression
func (m EncodingRequestExtendedBuilder) SetUseCompression(value UseCompressionEnum) {
	message.SetInt32Fixed(m.Unsafe(), 20, 16, int32(value))
}

// SetUseCompression
func (m EncodingRequestExtendedPointerBuilder) SetUseCompression(value UseCompressionEnum) {
	message.SetInt32Fixed(m.Ptr, 20, 16, int32(value))
}

// Copy
func (m EncodingRequestExtended) Copy(to EncodingRequestExtendedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
	to.SetUseCompression(m.UseCompression())
}

// Copy
func (m EncodingRequestExtendedBuilder) Copy(to EncodingRequestExtendedBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
	to.SetUseCompression(m.UseCompression())
}

// CopyPointer
func (m EncodingRequestExtended) CopyPointer(to EncodingRequestExtendedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
	to.SetUseCompression(m.UseCompression())
}

// CopyPointer
func (m EncodingRequestExtendedBuilder) CopyPointer(to EncodingRequestExtendedPointerBuilder) {
	to.SetProtocolVersion(m.ProtocolVersion())
	to.SetEncoding(m.Encoding())
	to.SetProtocolType(m.ProtocolType())
	to.SetUseCompression(m.UseCompression())
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
