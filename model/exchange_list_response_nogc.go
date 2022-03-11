package model

import (
	"github.com/moontrade/dtc-go/message"
)

// ExchangeListResponsePointer The server will return this message for each supported exchange.
//
// If there are no exchanges to return in response to a request, send through
// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
// all other members in the default state and the Client will recognize there
// are no Exchanges.
type ExchangeListResponsePointer struct {
	message.VLSPointer
}

// ExchangeListResponsePointerBuilder The server will return this message for each supported exchange.
//
// If there are no exchanges to return in response to a request, send through
// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
// all other members in the default state and the Client will recognize there
// are no Exchanges.
type ExchangeListResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocExchangeListResponse() ExchangeListResponsePointerBuilder {
	a := ExchangeListResponsePointerBuilder{message.AllocVLSBuilder(24)}
	a.Ptr.SetBytes(0, _ExchangeListResponseDefault)
	return a
}

func AllocExchangeListResponseFrom(b []byte) ExchangeListResponsePointer {
	return ExchangeListResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m ExchangeListResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ExchangeListResponseDefault)
}

// ToBuilder
func (m ExchangeListResponsePointer) ToBuilder() ExchangeListResponsePointerBuilder {
	return ExchangeListResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *ExchangeListResponsePointerBuilder) Finish() ExchangeListResponsePointer {
	return ExchangeListResponsePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponsePointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponsePointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// Exchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponsePointer) Exchange() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Exchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponsePointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponsePointer) IsFinalMessage() uint8 {
	return message.Uint8VLS(m.Ptr, 17, 16)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponsePointerBuilder) IsFinalMessage() uint8 {
	return message.Uint8VLS(m.Ptr, 17, 16)
}

// Description The complete exchange description.
func (m ExchangeListResponsePointer) Description() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// Description The complete exchange description.
func (m ExchangeListResponsePointerBuilder) Description() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// SetRequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetExchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m *ExchangeListResponsePointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetIsFinalMessage 1 = final message in batch.
func (m ExchangeListResponsePointerBuilder) SetIsFinalMessage(value uint8) {
	message.SetUint8VLS(m.Ptr, 17, 16, value)
}

// SetDescription The complete exchange description.
func (m *ExchangeListResponsePointerBuilder) SetDescription(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 22, 18, value)
}

// Copy
func (m ExchangeListResponsePointer) Copy(to ExchangeListResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// Copy
func (m ExchangeListResponsePointerBuilder) Copy(to ExchangeListResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyPointer
func (m ExchangeListResponsePointer) CopyPointer(to ExchangeListResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyPointer
func (m ExchangeListResponsePointerBuilder) CopyPointer(to ExchangeListResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyTo
func (m ExchangeListResponsePointer) CopyTo(to ExchangeListResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyTo
func (m ExchangeListResponsePointerBuilder) CopyTo(to ExchangeListResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyToPointer
func (m ExchangeListResponsePointer) CopyToPointer(to ExchangeListResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyToPointer
func (m ExchangeListResponsePointerBuilder) CopyToPointer(to ExchangeListResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}
