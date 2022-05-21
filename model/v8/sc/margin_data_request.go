// generated by github.com/moontrade/dtc-go/codegen/go at 2022-03-14 08:04:12.929995 +0800 WITA m=+0.027582959

package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const MarginDataRequestSize = 18

//     Size          uint16  = MarginDataRequestSize  (18)
//     Type          uint16  = MARGIN_DATA_REQUEST  (10141)
//     BaseSize      uint16  = MarginDataRequestSize  (18)
//     RequestID     int32   = 0
//     TradeAccount  string  = ""
//     Symbol        string  = ""
var _MarginDataRequestDefault = []byte{18, 0, 157, 39, 18, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarginDataRequest struct {
	message.VLS
}

type MarginDataRequestBuilder struct {
	message.VLSBuilder
}

type MarginDataRequestPointer struct {
	message.VLSPointer
}

type MarginDataRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func NewMarginDataRequestFrom(b []byte) MarginDataRequest {
	return MarginDataRequest{VLS: message.NewVLS(b)}
}

func WrapMarginDataRequest(b []byte) MarginDataRequest {
	return MarginDataRequest{VLS: message.WrapVLS(b)}
}

func NewMarginDataRequest() MarginDataRequestBuilder {
	return MarginDataRequestBuilder{message.NewVLSBuilder(_MarginDataRequestDefault)}
}

func AllocMarginDataRequest() MarginDataRequestPointerBuilder {
	return MarginDataRequestPointerBuilder{message.AllocVLSBuilder(_MarginDataRequestDefault)}
}

func AllocMarginDataRequestFrom(b []byte) MarginDataRequestPointer {
	return MarginDataRequestPointer{VLSPointer: message.AllocVLS(b)}
}

// Clear
//     Size          uint16  = MarginDataRequestSize  (18)
//     Type          uint16  = MARGIN_DATA_REQUEST  (10141)
//     BaseSize      uint16  = MarginDataRequestSize  (18)
//     RequestID     int32   = 0
//     TradeAccount  string  = ""
//     Symbol        string  = ""
func (m MarginDataRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarginDataRequestDefault)
}

// Clear
//     Size          uint16  = MarginDataRequestSize  (18)
//     Type          uint16  = MARGIN_DATA_REQUEST  (10141)
//     BaseSize      uint16  = MarginDataRequestSize  (18)
//     RequestID     int32   = 0
//     TradeAccount  string  = ""
//     Symbol        string  = ""
func (m MarginDataRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarginDataRequestDefault)
}

// ToBuilder
func (m MarginDataRequest) ToBuilder() MarginDataRequestBuilder {
	return MarginDataRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m MarginDataRequestPointer) ToBuilder() MarginDataRequestPointerBuilder {
	return MarginDataRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m MarginDataRequestBuilder) Finish() MarginDataRequest {
	return MarginDataRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m *MarginDataRequestPointerBuilder) Finish() MarginDataRequestPointer {
	return MarginDataRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m MarginDataRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m MarginDataRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m MarginDataRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m MarginDataRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m MarginDataRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m MarginDataRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m MarginDataRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m MarginDataRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Symbol
func (m MarginDataRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Symbol
func (m MarginDataRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Symbol
func (m MarginDataRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// Symbol
func (m MarginDataRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// SetRequestID
func (m MarginDataRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m MarginDataRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *MarginDataRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *MarginDataRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetSymbol
func (m *MarginDataRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetSymbol
func (m *MarginDataRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}