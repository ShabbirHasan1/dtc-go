package fixed

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/model"
	"github.com/moontrade/dtc-go/model/serialize"
	"github.com/moontrade/nogc"
)

var (
	_                      model.EncodingRequest        = EncodingRequest{}
	_                      model.EncodingRequest        = EncodingRequestPointer{}
	_                      model.EncodingRequestBuilder = EncodingRequestBuilder{}
	_                      model.EncodingRequestBuilder = EncodingRequestPointerBuilder{}
	EncodingRequestFactory model.EncodingRequestFactory = EncodingRequestFactoryImpl{}
)

//////////////////////////////////////////////////////////////////////////////////////////
// Types
//////////////////////////////////////////////////////////////////////////////////////////

type EncodingRequest struct {
	message.Fixed
}

type EncodingRequestBuilder struct {
	message.Fixed
}

type EncodingRequestPointer struct {
	message.FixedPointer
}

type EncodingRequestPointerBuilder struct {
	message.FixedPointer
}

//////////////////////////////////////////////////////////////////////////////////////////
// Constructors
//////////////////////////////////////////////////////////////////////////////////////////

func NewEncodingRequest() EncodingRequestBuilder {
	a := EncodingRequestBuilder{message.NewFixed(16)}
	a.Clear()
	return a
}

func AllocEncodingRequest() EncodingRequestPointerBuilder {
	a := EncodingRequestPointerBuilder{message.AllocFixed(16)}
	a.Clear()
	return a
}

//////////////////////////////////////////////////////////////////////////////////////////
// Factory
//////////////////////////////////////////////////////////////////////////////////////////

type EncodingRequestFactoryImpl struct{}

func (EncodingRequestFactoryImpl) New() model.EncodingRequestBuilder {
	return NewEncodingRequest()
}

func (EncodingRequestFactoryImpl) NewEx(flex int) model.EncodingRequestBuilder {
	return AllocEncodingRequest()
}

func (EncodingRequestFactoryImpl) Alloc() model.EncodingRequestBuilder {
	return AllocEncodingRequest()
}

func (EncodingRequestFactoryImpl) AllocEx(flex int) model.EncodingRequestBuilder {
	return AllocEncodingRequest()
}

func (EncodingRequestFactoryImpl) Wrap(b []byte) model.EncodingRequest {
	return EncodingRequest{message.WrapFixedFromBytes(b)}
}

func (EncodingRequestFactoryImpl) NewCopy(b []byte) model.EncodingRequest {
	return EncodingRequest{message.NewFixedFromBytes(b)}
}

func (EncodingRequestFactoryImpl) AllocCopy(b []byte) model.EncodingRequest {
	return EncodingRequestPointer{message.AllocFixedFrom(b)}
}

func (f EncodingRequestFactoryImpl) Clone(of model.EncodingRequest) model.EncodingRequest {
	if of.IsGC() {
		return f.NewClone(of)
	}
	return f.AllocClone(of)
}

func (EncodingRequestFactoryImpl) NewClone(of model.EncodingRequest) model.EncodingRequest {
	a := NewEncodingRequest()
	a.CopyFrom(of)
	return a
}

func (EncodingRequestFactoryImpl) AllocClone(of model.EncodingRequest) model.EncodingRequest {
	a := AllocEncodingRequest()
	a.CopyFrom(of)
	return a
}

//////////////////////////////////////////////////////////////////////////////////////////
// CopyFrom
//////////////////////////////////////////////////////////////////////////////////////////

func encodingRequestCopyFrom(from model.EncodingRequest, to model.EncodingRequestBuilder) {
	clearEncodingRequest(to.Unsafe())
	to.SetEncoding(from.Encoding())
	to.SetProtocolVersion(from.ProtocolVersion())
	to.SetProtocolType(from.ProtocolType())
}

func (e EncodingRequestBuilder) CopyFrom(from model.EncodingRequest) {
	encodingRequestCopyFrom(from, e)
}
func (e EncodingRequestPointerBuilder) CopyFrom(from model.EncodingRequest) {
	encodingRequestCopyFrom(from, e)
}

//////////////////////////////////////////////////////////////////////////////////////////
// ToBuilder
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequest) ToBuilder() model.EncodingRequestBuilder {
	return EncodingRequestBuilder{Fixed: e.Fixed}
}
func (e EncodingRequestBuilder) ToBuilder() model.EncodingRequestBuilder {
	return e
}
func (e EncodingRequestPointer) ToBuilder() model.EncodingRequestBuilder {
	return EncodingRequestPointerBuilder{e.FixedPointer}
}
func (e EncodingRequestPointerBuilder) ToBuilder() model.EncodingRequestBuilder {
	return e
}

//////////////////////////////////////////////////////////////////////////////////////////
// Getters
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequest) ProtocolVersion() int32 {
	return message.Int32(e.Unsafe(), e.Size(), 8, 4)
}
func (e EncodingRequestPointer) ProtocolVersion() int32 {
	return message.Int32(e.Unsafe(), e.Size(), 8, 4)
}
func (e EncodingRequest) Encoding() int32 {
	return message.Int32(e.Unsafe(), e.Size(), 12, 8)
}
func (e EncodingRequestPointer) Encoding() int32 {
	return message.Int32(e.Unsafe(), e.Size(), 12, 8)
}
func (e EncodingRequest) ProtocolType() string {
	return message.StringFixed(e.Unsafe(), e.Size(), 16, 12, 4)
}
func (e EncodingRequestPointer) ProtocolType() string {
	return message.StringFixed(e.Unsafe(), e.Size(), 16, 12, 4)
}

func (e EncodingRequestBuilder) ProtocolVersion() int32 {
	return message.Int32(e.Unsafe(), e.Size(), 8, 4)
}
func (e EncodingRequestPointerBuilder) ProtocolVersion() int32 {
	return message.Int32(e.Unsafe(), e.Size(), 8, 4)
}
func (e EncodingRequestBuilder) Encoding() int32 {
	return message.Int32(e.Unsafe(), e.Size(), 12, 8)
}
func (e EncodingRequestPointerBuilder) Encoding() int32 {
	return message.Int32(e.Unsafe(), e.Size(), 12, 8)
}
func (e EncodingRequestBuilder) ProtocolType() string {
	return message.StringFixed(e.Unsafe(), e.Size(), 16, 12, 4)
}
func (e EncodingRequestPointerBuilder) ProtocolType() string {
	return message.StringFixed(e.Unsafe(), e.Size(), 16, 12, 4)
}

//////////////////////////////////////////////////////////////////////////////////////////
// Finish
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequestBuilder) Finish() model.EncodingRequest {
	return EncodingRequest{e.Fixed}
}
func (e EncodingRequestPointerBuilder) Finish() model.EncodingRequest {
	return EncodingRequestPointer{e.FixedPointer}
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
	clearEncodingRequest(e.Unsafe())
}
func (e EncodingRequestPointerBuilder) Clear() {
	clearEncodingRequest(e.Unsafe())
}

//////////////////////////////////////////////////////////////////////////////////////////
// Setters
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequestBuilder) SetProtocolVersion(value int32) {
	message.SetInt32(e.Unsafe(), e.Size(), 8, 4, value)
}
func (e EncodingRequestPointerBuilder) SetProtocolVersion(value int32) {
	message.SetInt32(e.Unsafe(), e.Size(), 8, 4, value)
}
func (e EncodingRequestBuilder) SetEncoding(value int32) {
	message.SetInt32(e.Unsafe(), e.Size(), 12, 8, value)
}
func (e EncodingRequestPointerBuilder) SetEncoding(value int32) {
	message.SetInt32(e.Unsafe(), e.Size(), 12, 8, value)
}
func (e EncodingRequestBuilder) SetProtocolType(value string) {
	message.SetStringFixed(e.Unsafe(), e.Size(), 16, 12, 4, value)
}
func (e EncodingRequestPointerBuilder) SetProtocolType(value string) {
	message.SetStringFixed(e.Unsafe(), e.Size(), 16, 12, 4, value)
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
