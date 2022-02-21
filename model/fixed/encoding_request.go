package fixed

import (
	dtc "github.com/moontrade/dtc-go"
	"github.com/moontrade/dtc-go/json"
	"github.com/moontrade/dtc-go/model/types"
	"github.com/moontrade/nogc"
)

var (
	_ types.EncodingRequest        = EncodingRequest{}
	_ types.EncodingRequest        = EncodingRequestPointer{}
	_ types.EncodingRequestBuilder = EncodingRequestBuilder{}
	_ types.EncodingRequestBuilder = EncodingRequestPointerBuilder{}
	_ types.EncodingRequestFactory = EncodingRequestFactory{}
)

type EncodingRequestFactory struct {
}

func (EncodingRequestFactory) New() types.EncodingRequestBuilder {
	return NewEncodingRequest()
}

func (EncodingRequestFactory) Alloc() types.EncodingRequestBuilder {
	return AllocEncodingRequest()
}

func (EncodingRequestFactory) FromJSON(b []byte, into types.EncodingRequestBuilder) error {
	lexer := json.Lexer{
		Data:              b,
		UseMultipleErrors: false,
	}
	_ = lexer
	return nil
}

func (EncodingRequestFactory) FromJSONCompact(b []byte, to types.EncodingRequestBuilder) (types.EncodingRequestBuilder, error) {
	return NewEncodingRequestFromJSONCompact(&json.Lexer{Data: b}, to)
}

//////////////////////////////////////////////////////////////////////////////////////////
// Getters
//////////////////////////////////////////////////////////////////////////////////////////

type EncodingRequest struct {
	dtc.MessageFixed
}

type EncodingRequestBuilder struct {
	dtc.MessageFixed
}

type EncodingRequestPointer struct {
	dtc.MessageFixedPointer
}

type EncodingRequestPointerBuilder struct {
	dtc.MessageFixedPointer
}

func WrapEncodingRequest(b []byte) EncodingRequest {
	return EncodingRequest{}
}

func (e EncodingRequestBuilder) CopyFrom(from types.EncodingRequest) {
	clearEncodingRequest(e.AsPointer())
	e.SetEncoding(from.Encoding())
	e.SetProtocolVersion(from.ProtocolVersion())
	e.SetProtocolType(from.ProtocolType())
}

func NewEncodingRequest() EncodingRequestBuilder {
	e := EncodingRequestBuilder{dtc.NewFixed(16)}
	clearEncodingRequest(e.AsPointer())
	return e
}

func AllocEncodingRequest() EncodingRequestPointerBuilder {
	a := EncodingRequestPointerBuilder{dtc.AllocFixed(16)}
	clearEncodingRequest(a.AsPointer())
	return a
}

//////////////////////////////////////////////////////////////////////////////////////////
// ToBuilder
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequest) ToBuilder() types.EncodingRequestBuilder {
	return EncodingRequestBuilder{MessageFixed: e.MessageFixed}
}
func (e EncodingRequestBuilder) ToBuilder() types.EncodingRequestBuilder {
	return e
}
func (e EncodingRequestPointer) ToBuilder() types.EncodingRequestBuilder {
	return EncodingRequestPointerBuilder{e.MessageFixedPointer}
}
func (e EncodingRequestPointerBuilder) ToBuilder() types.EncodingRequestBuilder {
	return e
}

//////////////////////////////////////////////////////////////////////////////////////////
// Getters
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequest) ProtocolVersion() int32 {
	return dtc.Int32(e.AsPointer(), e.Size(), 8, 4)
}
func (e EncodingRequestPointer) ProtocolVersion() int32 {
	return dtc.Int32(e.AsPointer(), e.Size(), 8, 4)
}
func (e EncodingRequest) Encoding() int32 {
	return dtc.Int32(e.AsPointer(), e.Size(), 12, 8)
}
func (e EncodingRequestPointer) Encoding() int32 {
	return dtc.Int32(e.AsPointer(), e.Size(), 12, 8)
}
func (e EncodingRequest) ProtocolType() string {
	return dtc.StringFixed(e.AsPointer(), e.Size(), 16, 12, 4)
}
func (e EncodingRequestPointer) ProtocolType() string {
	return dtc.StringFixed(e.AsPointer(), e.Size(), 16, 12, 4)
}

func (e EncodingRequestBuilder) ProtocolVersion() int32 {
	return dtc.Int32(e.AsPointer(), e.Size(), 8, 4)
}
func (e EncodingRequestPointerBuilder) ProtocolVersion() int32 {
	return dtc.Int32(e.AsPointer(), e.Size(), 8, 4)
}
func (e EncodingRequestBuilder) Encoding() int32 {
	return dtc.Int32(e.AsPointer(), e.Size(), 12, 8)
}
func (e EncodingRequestPointerBuilder) Encoding() int32 {
	return dtc.Int32(e.AsPointer(), e.Size(), 12, 8)
}
func (e EncodingRequestBuilder) ProtocolType() string {
	return dtc.StringFixed(e.AsPointer(), e.Size(), 16, 12, 4)
}
func (e EncodingRequestPointerBuilder) ProtocolType() string {
	return dtc.StringFixed(e.AsPointer(), e.Size(), 16, 12, 4)
}

//////////////////////////////////////////////////////////////////////////////////////////
// Finish
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequestBuilder) Finish() types.EncodingRequest {
	return EncodingRequest{e.MessageFixed}
}
func (e EncodingRequestPointerBuilder) Finish() types.EncodingRequest {
	return EncodingRequestPointer{e.TakePointer()}
}

//////////////////////////////////////////////////////////////////////////////////////////
// Clear
//////////////////////////////////////////////////////////////////////////////////////////

func clearEncodingRequest(p nogc.Pointer) {
	if p == 0 {
		return
	}
	nogc.Zero(p.Unsafe(), 16)
	p.SetUInt16LE(0, 16)
	p.SetUInt16LE(2, 6)
	p.SetInt32LE(4, 8)
}
func (e EncodingRequestBuilder) Clear() {
	clearEncodingRequest(e.AsPointer())
}
func (e EncodingRequestPointer) Clear() {
	clearEncodingRequest(e.AsPointer())
}

//////////////////////////////////////////////////////////////////////////////////////////
// Setters
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequestBuilder) SetProtocolVersion(value int32) {
	dtc.SetInt32(e.AsPointer(), e.Size(), 8, 4, value)
}
func (e EncodingRequestPointerBuilder) SetProtocolVersion(value int32) {
	dtc.SetInt32(e.AsPointer(), e.Size(), 8, 4, value)
}
func (e EncodingRequestBuilder) SetEncoding(value int32) {
	dtc.SetInt32(e.AsPointer(), e.Size(), 12, 8, value)
}
func (e EncodingRequestPointerBuilder) SetEncoding(value int32) {
	dtc.SetInt32(e.AsPointer(), e.Size(), 12, 8, value)
}
func (e EncodingRequestBuilder) SetProtocolType(value string) {
	dtc.SetStringFixed(e.AsPointer(), e.Size(), 16, 12, 4, value)
}
func (e EncodingRequestPointerBuilder) SetProtocolType(value string) {
	dtc.SetStringFixed(e.AsPointer(), e.Size(), 16, 12, 4, value)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact
//////////////////////////////////////////////////////////////////////////////////////////

func NewEncodingRequestFromJSONCompact(r *json.Lexer, to types.EncodingRequestBuilder) (types.EncodingRequestBuilder, error) {
	if to == nil {
		to = NewEncodingRequest()
	}
	return encodingRequestFromJSONCompact(r, to)
}

func AllocEncodingRequestFromJSONCompact(r *json.Lexer, to types.EncodingRequestBuilder) (types.EncodingRequestBuilder, error) {
	if to == nil {
		to = AllocEncodingRequest()
		result, err := encodingRequestFromJSONCompact(r, to)
		if err != nil {
			// free memory
			to.Finish()
			return nil, err
		}
		return result, nil
	}
	return encodingRequestFromJSONCompact(r, to)
}

func encodingRequestFromJSONCompact(r *json.Lexer, to types.EncodingRequestBuilder) (types.EncodingRequestBuilder, error) {
	return to, nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact
//////////////////////////////////////////////////////////////////////////////////////////

func encodingRequestFromJSON(r *json.Lexer, to types.EncodingRequestBuilder) (types.EncodingRequestBuilder, error) {
	return to, nil
}

func encodingRequestWriteJSONCompact(m types.EncodingRequest, b []byte) ([]byte, error) {
	w := json.NewCompactWriterI32(b, 6, m.ProtocolVersion())
	w.Int32Compact(m.Encoding()).StringCompact(m.ProtocolType())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON
//////////////////////////////////////////////////////////////////////////////////////////

func encodingRequestWriteJSON(m types.EncodingRequest, b []byte) ([]byte, error) {
	w := json.NewWriter(b, m.Type())
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Encoding", m.Encoding())
	w.StringField("ProtocolType", m.ProtocolType())
	return w.Finish(), nil
}
