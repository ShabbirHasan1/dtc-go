package vls

import (
	dtc "github.com/moontrade/dtc-go"
	"github.com/moontrade/dtc-go/model/types"
	"github.com/moontrade/nogc"
)

var (
	_ types.EncodingRequest        = EncodingRequest{}
	_ types.EncodingRequest        = EncodingRequestPointer{}
	_ types.EncodingRequestBuilder = EncodingRequestBuilder{}
	_ types.EncodingRequestBuilder = EncodingRequestPointerBuilder{}
)

//////////////////////////////////////////////////////////////////////////////////////////
// Getters
//////////////////////////////////////////////////////////////////////////////////////////

type EncodingRequest struct {
	dtc.MessageVLS
}

type EncodingRequestBuilder struct {
	*dtc.MessageVLSBuilder
}

type EncodingRequestPointer struct {
	dtc.MessageVLSPointer
}

type EncodingRequestPointerBuilder struct {
	*dtc.MessageVLSPointerBuilder
}

func NewEncodingRequest() EncodingRequestBuilder {
	return NewEncodingRequestBuilderFrom(nil, 32, 32)
}

func NewEncodingRequestFromBytes(b []byte) EncodingRequest {
	return EncodingRequest{dtc.NewVLSFromBytesOfType(b, 6)}
}

func WrapEncodingRequest(b []byte) EncodingRequest {
	return EncodingRequest{dtc.WrapVLSFromBytesOfType(b, 6)}
}

func NewEncodingRequestBuilderFrom(b dtc.MessageBuffer, flex uintptr, growBy int) EncodingRequestBuilder {
	a := EncodingRequestBuilder{dtc.MessageVLSBuilderReset(b, nil, 16, flex, growBy)}
	clearEncodingRequest(a.AsPointer())
	return a
}

func AllocEncodingRequest() EncodingRequestPointerBuilder {
	return AllocEncodingRequestFrom(nil, 32, 32)
}

func AllocEncodingRequestFromBytes(b []byte) EncodingRequestPointer {
	return EncodingRequestPointer{dtc.AllocVLSFromBytesOfType(b, 6)}
}

func AllocEncodingRequestFrom(b dtc.MessageBuffer, flex uintptr, growBy int) EncodingRequestPointerBuilder {
	a := EncodingRequestPointerBuilder{dtc.MessageVLSPointerBuilderReset(b, nil, 16, flex, growBy)}
	clearEncodingRequest(a.AsPointer())
	return a
}

//////////////////////////////////////////////////////////////////////////////////////////
// ToBuilder
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequest) ToBuilder() types.EncodingRequestBuilder {
	return e.ToBuilderWithBuffer(nil, 16)
}
func (e EncodingRequestBuilder) ToBuilder() types.EncodingRequestBuilder {
	return e
}

func (e EncodingRequest) ToBuilderWithBuffer(b dtc.MessageBuffer, flex uintptr) types.EncodingRequestBuilder {
	return EncodingRequestBuilder{dtc.MessageVLSBuilderReset(b, &e.MessageVLS, 16, 16, 0)}
}

func (e EncodingRequestPointer) ToBuilder() types.EncodingRequestBuilder {
	return e.ToBuilderWithBuffer(nil, 16)
}
func (e EncodingRequestPointerBuilder) ToBuilder() types.EncodingRequestBuilder {
	return e
}

func (e EncodingRequestPointer) ToBuilderWithBuffer(b dtc.MessageBuffer, flex uintptr) types.EncodingRequestBuilder {
	return EncodingRequestPointerBuilder{dtc.MessageVLSPointerBuilderReset(b, &e.MessageVLSPointer, 16, flex, 0)}
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
	return dtc.StringVLS(e.AsPointer(), 16, 12)
}
func (e EncodingRequestPointer) ProtocolType() string {
	return dtc.StringVLS(e.AsPointer(), 16, 12)
}

func (e EncodingRequestBuilder) ProtocolVersion() int32 {
	return dtc.Int32(e.AsPointer(), e.BaseSize(), 8, 4)
}
func (e EncodingRequestPointerBuilder) ProtocolVersion() int32 {
	return dtc.Int32(e.AsPointer(), e.BaseSize(), 8, 4)
}
func (e EncodingRequestBuilder) Encoding() int32 {
	return dtc.Int32(e.AsPointer(), e.BaseSize(), 12, 8)
}
func (e EncodingRequestPointerBuilder) Encoding() int32 {
	return dtc.Int32(e.AsPointer(), e.BaseSize(), 12, 8)
}
func (e EncodingRequestBuilder) ProtocolType() string {
	return dtc.StringVLS(e.AsPointer(), 16, 12)
}
func (e EncodingRequestPointerBuilder) ProtocolType() string {
	return dtc.StringVLS(e.AsPointer(), 16, 12)
}

//////////////////////////////////////////////////////////////////////////////////////////
// Finish
//////////////////////////////////////////////////////////////////////////////////////////

func (e EncodingRequestBuilder) Finish() types.EncodingRequest {
	return EncodingRequest{e.MessageVLSBuilder.Finish()}
}
func (e EncodingRequestPointerBuilder) Finish() types.EncodingRequest {
	return EncodingRequestPointer{e.MessageVLSPointerBuilder.Finish()}
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
	dtc.SetInt32(e.AsPointer(), e.Size(), 12, 4, value)
}
func (e EncodingRequestPointerBuilder) SetProtocolVersion(value int32) {
	dtc.SetInt32(e.AsPointer(), e.Size(), 12, 4, value)
}
func (e EncodingRequestBuilder) SetEncoding(value int32) {
	dtc.SetInt32(e.AsPointer(), e.Size(), 12, 4, value)
}
func (e EncodingRequestPointerBuilder) SetEncoding(value int32) {
	dtc.SetInt32(e.AsPointer(), e.Size(), 12, 4, value)
}
func (e EncodingRequestBuilder) SetProtocolType(value string) {
	dtc.SetStringVLS(e.MessageVLSBuilder, 16, 16, 12, value)
}
func (e EncodingRequestPointerBuilder) SetProtocolType(value string) {
	dtc.SetStringVLSPointer(e.MessageVLSPointerBuilder, 16, 16, 12, value)
}
