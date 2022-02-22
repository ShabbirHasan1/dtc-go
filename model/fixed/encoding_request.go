package fixed

import (
	"github.com/moontrade/dtc-go"
	"github.com/moontrade/dtc-go/json"
	"github.com/moontrade/dtc-go/model/serialize"
	"github.com/moontrade/dtc-go/model/types"
	"github.com/moontrade/nogc"
)

var (
	_                      types.EncodingRequest        = EncodingRequest{}
	_                      types.EncodingRequest        = EncodingRequestPointer{}
	_                      types.EncodingRequestBuilder = EncodingRequestBuilder{}
	_                      types.EncodingRequestBuilder = EncodingRequestPointerBuilder{}
	EncodingRequestFactory types.EncodingRequestFactory = EncodingRequestFactoryImpl{}
)

type EncodingRequestFactoryImpl struct {
}

func (EncodingRequestFactoryImpl) New() types.EncodingRequestBuilder {
	return NewEncodingRequest()
}

func (EncodingRequestFactoryImpl) NewEx(flex int) types.EncodingRequestBuilder {
	return AllocEncodingRequest()
}

func (EncodingRequestFactoryImpl) Alloc() types.EncodingRequestBuilder {
	return AllocEncodingRequest()
}

func (EncodingRequestFactoryImpl) AllocEx(flex int) types.EncodingRequestBuilder {
	return AllocEncodingRequest()
}

func (EncodingRequestFactoryImpl) Wrap(b []byte) types.EncodingRequest {
	return EncodingRequest{dtc.WrapFixedFromBytes(b)}
}

func (EncodingRequestFactoryImpl) NewCopy(b []byte) types.EncodingRequest {
	return EncodingRequest{dtc.NewFixedFromBytes(b)}
}

func (EncodingRequestFactoryImpl) AllocCopy(b []byte) types.EncodingRequest {
	return EncodingRequestPointer{dtc.AllocCopyFrom(b)}
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
	return EncodingRequestPointer{e.MessageFixedPointer}
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
func (e EncodingRequestPointerBuilder) Clear() {
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
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequestBuilder) UnmarshalJSON(r *json.Reader) error {
	if r.IsCompact {
		return serialize.EncodingRequestUnmarshalJSONCompact(r, e)
	}
	return serialize.EncodingRequestFromJSON(r, e)
}
func (e EncodingRequestPointerBuilder) UnmarshalJSON(r *json.Reader) error {
	if r.IsCompact {
		return serialize.EncodingRequestUnmarshalJSONCompact(r, e)
	}
	return serialize.EncodingRequestFromJSON(r, e)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequest) MarshalJSONCompact(b []byte) ([]byte, error) {
	return serialize.EncodingRequestMarshalJSONCompact(e, b)
}
func (e EncodingRequestBuilder) MarshalJSONCompact(b []byte) ([]byte, error) {
	return serialize.EncodingRequestMarshalJSONCompact(e, b)
}
func (e EncodingRequestPointer) MarshalJSONCompact(b []byte) ([]byte, error) {
	return serialize.EncodingRequestMarshalJSONCompact(e, b)
}
func (e EncodingRequestPointerBuilder) MarshalJSONCompact(b []byte) ([]byte, error) {
	return serialize.EncodingRequestMarshalJSONCompact(e, b)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequest) MarshalJSON(b []byte) ([]byte, error) {
	return serialize.EncodingRequestMarshalJSON(e, b)
}
func (e EncodingRequestBuilder) MarshalJSON(b []byte) ([]byte, error) {
	return serialize.EncodingRequestMarshalJSON(e, b)
}
func (e EncodingRequestPointer) MarshalJSON(b []byte) ([]byte, error) {
	return serialize.EncodingRequestMarshalJSON(e, b)
}
func (e EncodingRequestPointerBuilder) MarshalJSON(b []byte) ([]byte, error) {
	return serialize.EncodingRequestMarshalJSON(e, b)
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequestBuilder) UnmarshalProtobuf(b []byte) error {
	return serialize.EncodingRequestUnmarshalProtobuf(b, e)
}
func (e EncodingRequestPointerBuilder) UnmarshalProtobuf(b []byte) error {
	return serialize.EncodingRequestUnmarshalProtobuf(b, e)
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequest) MarshalProtobuf(b []byte) ([]byte, error) {
	return serialize.EncodingRequestMarshalProtobuf(e, b)
}
func (e EncodingRequestBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	return serialize.EncodingRequestMarshalProtobuf(e, b)
}
func (e EncodingRequestPointer) MarshalProtobuf(b []byte) ([]byte, error) {
	return serialize.EncodingRequestMarshalProtobuf(e, b)
}
func (e EncodingRequestPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	return serialize.EncodingRequestMarshalProtobuf(e, b)
}
