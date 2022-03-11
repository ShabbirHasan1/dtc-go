package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size           = ExchangeListResponseSize  (24)
//     Type           = EXCHANGE_LIST_RESPONSE  (501)
//     BaseSize       = ExchangeListResponseSize  (24)
//     RequestID      = 0
//     Exchange       = ""
//     IsFinalMessage = 0
//     Description    = ""
// }
var _ExchangeListResponseDefault = []byte{24, 0, 245, 1, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const ExchangeListResponseSize = 24

// ExchangeListResponse The server will return this message for each supported exchange.
//
// If there are no exchanges to return in response to a request, send through
// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
// all other members in the default state and the Client will recognize there
// are no Exchanges.
type ExchangeListResponse struct {
	message.VLS
}

// ExchangeListResponseBuilder The server will return this message for each supported exchange.
//
// If there are no exchanges to return in response to a request, send through
// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
// all other members in the default state and the Client will recognize there
// are no Exchanges.
type ExchangeListResponseBuilder struct {
	message.VLSBuilder
}

func NewExchangeListResponseFrom(b []byte) ExchangeListResponse {
	return ExchangeListResponse{VLS: message.NewVLSFrom(b)}
}

func WrapExchangeListResponse(b []byte) ExchangeListResponse {
	return ExchangeListResponse{VLS: message.WrapVLS(b)}
}

func NewExchangeListResponse() ExchangeListResponseBuilder {
	a := ExchangeListResponseBuilder{message.NewVLSBuilder(24)}
	a.Unsafe().SetBytes(0, _ExchangeListResponseDefault)
	return a
}

// Clear
// {
//     Size           = ExchangeListResponseSize  (24)
//     Type           = EXCHANGE_LIST_RESPONSE  (501)
//     BaseSize       = ExchangeListResponseSize  (24)
//     RequestID      = 0
//     Exchange       = ""
//     IsFinalMessage = 0
//     Description    = ""
// }
func (m ExchangeListResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ExchangeListResponseDefault)
}

// ToBuilder
func (m ExchangeListResponse) ToBuilder() ExchangeListResponseBuilder {
	return ExchangeListResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m ExchangeListResponseBuilder) Finish() ExchangeListResponse {
	return ExchangeListResponse{m.VLSBuilder.Finish()}
}

// RequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponse) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponseBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// Exchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponse) Exchange() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Exchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponseBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponse) IsFinalMessage() uint8 {
	return message.Uint8VLS(m.Unsafe(), 17, 16)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseBuilder) IsFinalMessage() uint8 {
	return message.Uint8VLS(m.Unsafe(), 17, 16)
}

// Description The complete exchange description.
func (m ExchangeListResponse) Description() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// Description The complete exchange description.
func (m ExchangeListResponseBuilder) Description() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// SetRequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponseBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetExchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m *ExchangeListResponseBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetIsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseBuilder) SetIsFinalMessage(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 17, 16, value)
}

// SetDescription The complete exchange description.
func (m *ExchangeListResponseBuilder) SetDescription(value string) {
	message.SetStringVLS(&m.VLSBuilder, 22, 18, value)
}

// Copy
func (m ExchangeListResponse) Copy(to ExchangeListResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// Copy
func (m ExchangeListResponseBuilder) Copy(to ExchangeListResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyPointer
func (m ExchangeListResponse) CopyPointer(to ExchangeListResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyPointer
func (m ExchangeListResponseBuilder) CopyPointer(to ExchangeListResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyToPointer
func (m ExchangeListResponse) CopyToPointer(to ExchangeListResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyToPointer
func (m ExchangeListResponseBuilder) CopyToPointer(to ExchangeListResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyTo
func (m ExchangeListResponse) CopyTo(to ExchangeListResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyTo
func (m ExchangeListResponseBuilder) CopyTo(to ExchangeListResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}
