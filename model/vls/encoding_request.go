package vls

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/model"
	"github.com/moontrade/dtc-go/model/serialize"
	"github.com/moontrade/nogc"
)

var (
	_ model.EncodingRequest        = EncodingRequest{}
	_ model.EncodingRequest        = EncodingRequestPointer{}
	_ model.EncodingRequestBuilder = EncodingRequestBuilder{}
	_ model.EncodingRequestBuilder = EncodingRequestPointerBuilder{}
)

//////////////////////////////////////////////////////////////////////////////////////////
// Getters
//////////////////////////////////////////////////////////////////////////////////////////

type EncodingRequest struct {
	message.VLS
}

type EncodingRequestBuilder struct {
	*message.VLSBuilder
}

type EncodingRequestPointer struct {
	message.VLSPointer
}

type EncodingRequestPointerBuilder struct {
	*message.VLSPointerBuilder
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
	return EncodingRequest{message.WrapVLSFromBytes(b)}
}

func (EncodingRequestFactoryImpl) NewCopy(b []byte) model.EncodingRequest {
	return EncodingRequest{message.NewVLSFromBytes(b)}
}

func (EncodingRequestFactoryImpl) AllocCopy(b []byte) model.EncodingRequest {
	return EncodingRequestPointer{dtc.AllocVLSCopyFrom(b)}
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
	clearEncodingRequest(to.AsPointer())
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

func NewEncodingRequest() EncodingRequestBuilder {
	return NewEncodingRequestBuilderFrom(nil, 32, 32)
}

//
//func NewEncodingRequestFromBytes(b []byte) EncodingRequest {
//	return EncodingRequest{dtc.NewVLSFromBytesOfType(b, 6)}
//}

func NewEncodingRequestBuilderFrom(b message.Buffer, flex uintptr, growBy int) EncodingRequestBuilder {
	a := EncodingRequestBuilder{message.VLSBuilderReset(b, nil, 16, flex, growBy)}
	clearEncodingRequest(a.AsPointer())
	return a
}

func AllocEncodingRequest() EncodingRequestPointerBuilder {
	return AllocEncodingRequestFrom(nil, 32, 32)
}

func AllocEncodingRequestFrom(b message.Buffer, flex uintptr, growBy int) EncodingRequestPointerBuilder {
	a := EncodingRequestPointerBuilder{message.VLSPointerBuilderReset(b, nil, 16, flex, growBy)}
	clearEncodingRequest(a.AsPointer())
	return a
}

//////////////////////////////////////////////////////////////////////////////////////////
// ToBuilder
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequest) ToBuilder() model.EncodingRequestBuilder {
	return e.ToBuilderWithBuffer(nil, 16)
}
func (e EncodingRequestBuilder) ToBuilder() model.EncodingRequestBuilder {
	return e
}
func (e EncodingRequest) ToBuilderWithBuffer(b message.Buffer, flex uintptr) model.EncodingRequestBuilder {
	return EncodingRequestBuilder{message.VLSBuilderReset(b, &e.VLS, 16, 16, 0)}
}
func (e EncodingRequestPointer) ToBuilder() model.EncodingRequestBuilder {
	return e.ToBuilderWithBuffer(nil, 16)
}
func (e EncodingRequestPointerBuilder) ToBuilder() model.EncodingRequestBuilder {
	return e
}
func (e EncodingRequestPointer) ToBuilderWithBuffer(b message.Buffer, flex uintptr) model.EncodingRequestBuilder {
	return EncodingRequestPointerBuilder{message.VLSPointerBuilderReset(b, &e.VLSPointer, 16, flex, 0)}
}

//////////////////////////////////////////////////////////////////////////////////////////
// Getters
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequest) ProtocolVersion() int32 {
	return message.Int32(e.AsPointer(), e.Size(), 8, 4)
}
func (e EncodingRequestPointer) ProtocolVersion() int32 {
	return message.Int32(e.AsPointer(), e.Size(), 8, 4)
}
func (e EncodingRequest) Encoding() int32 {
	return message.Int32(e.AsPointer(), e.Size(), 12, 8)
}
func (e EncodingRequestPointer) Encoding() int32 {
	return message.Int32(e.AsPointer(), e.Size(), 12, 8)
}
func (e EncodingRequest) ProtocolType() string {
	return message.StringVLS(e.AsPointer(), 16, 12)
}
func (e EncodingRequestPointer) ProtocolType() string {
	return message.StringVLS(e.AsPointer(), 16, 12)
}

func (e EncodingRequestBuilder) ProtocolVersion() int32 {
	return message.Int32(e.AsPointer(), e.BaseSize(), 8, 4)
}
func (e EncodingRequestPointerBuilder) ProtocolVersion() int32 {
	return message.Int32(e.AsPointer(), e.BaseSize(), 8, 4)
}
func (e EncodingRequestBuilder) Encoding() int32 {
	return message.Int32(e.AsPointer(), e.BaseSize(), 12, 8)
}
func (e EncodingRequestPointerBuilder) Encoding() int32 {
	return message.Int32(e.AsPointer(), e.BaseSize(), 12, 8)
}
func (e EncodingRequestBuilder) ProtocolType() string {
	return message.StringVLS(e.AsPointer(), 16, 12)
}
func (e EncodingRequestPointerBuilder) ProtocolType() string {
	return message.StringVLS(e.AsPointer(), 16, 12)
}

//////////////////////////////////////////////////////////////////////////////////////////
// Finish
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequestBuilder) Finish() model.EncodingRequest {
	return EncodingRequest{e.VLSBuilder.Finish()}
}
func (e EncodingRequestPointerBuilder) Finish() model.EncodingRequest {
	return EncodingRequestPointer{e.VLSPointerBuilder.Finish()}
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
	p.SetUInt16LE(2, 110)
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
	message.SetInt32(e.AsPointer(), e.Size(), 12, 4, value)
}
func (e EncodingRequestPointerBuilder) SetProtocolVersion(value int32) {
	message.SetInt32(e.AsPointer(), e.Size(), 12, 4, value)
}
func (e EncodingRequestBuilder) SetEncoding(value int32) {
	message.SetInt32(e.AsPointer(), e.Size(), 12, 4, value)
}
func (e EncodingRequestPointerBuilder) SetEncoding(value int32) {
	message.SetInt32(e.AsPointer(), e.Size(), 12, 4, value)
}
func (e EncodingRequestBuilder) SetProtocolType(value string) {
	message.SetStringVLS(e.VLSBuilder, 16, 16, 12, value)
}
func (e EncodingRequestPointerBuilder) SetProtocolType(value string) {
	message.SetStringVLSPointer(e.VLSPointerBuilder, 16, 16, 12, value)
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
