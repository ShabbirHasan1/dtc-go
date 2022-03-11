package model

import (
	"github.com/moontrade/dtc-go/message"
)

// ExchangeListResponseFixedPointer The server will return this message for each supported exchange.
//
// If there are no exchanges to return in response to a request, send through
// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
// all other members in the default state and the Client will recognize there
// are no Exchanges.
type ExchangeListResponseFixedPointer struct {
	message.FixedPointer
}

// ExchangeListResponseFixedPointerBuilder The server will return this message for each supported exchange.
//
// If there are no exchanges to return in response to a request, send through
// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
// all other members in the default state and the Client will recognize there
// are no Exchanges.
type ExchangeListResponseFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocExchangeListResponseFixed() ExchangeListResponseFixedPointerBuilder {
	a := ExchangeListResponseFixedPointerBuilder{message.AllocFixed(76)}
	a.Ptr.SetBytes(0, _ExchangeListResponseFixedDefault)
	return a
}

func AllocExchangeListResponseFixedFrom(b []byte) ExchangeListResponseFixedPointer {
	return ExchangeListResponseFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size           = ExchangeListResponseFixedSize  (76)
//     Type           = EXCHANGE_LIST_RESPONSE  (501)
//     RequestID      = 0
//     Exchange       = ""
//     IsFinalMessage = 0
//     Description    = ""
// }
func (m ExchangeListResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ExchangeListResponseFixedDefault)
}

// ToBuilder
func (m ExchangeListResponseFixedPointer) ToBuilder() ExchangeListResponseFixedPointerBuilder {
	return ExchangeListResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *ExchangeListResponseFixedPointerBuilder) Finish() ExchangeListResponseFixedPointer {
	return ExchangeListResponseFixedPointer{m.FixedPointer}
}

// RequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponseFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponseFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Exchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponseFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 24, 8)
}

// Exchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponseFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 24, 8)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseFixedPointer) IsFinalMessage() uint8 {
	return message.Uint8Fixed(m.Ptr, 25, 24)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseFixedPointerBuilder) IsFinalMessage() uint8 {
	return message.Uint8Fixed(m.Ptr, 25, 24)
}

// Description The complete exchange description.
func (m ExchangeListResponseFixedPointer) Description() string {
	return message.StringFixed(m.Ptr, 73, 25)
}

// Description The complete exchange description.
func (m ExchangeListResponseFixedPointerBuilder) Description() string {
	return message.StringFixed(m.Ptr, 73, 25)
}

// SetRequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponseFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetExchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponseFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 24, 8, value)
}

// SetIsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseFixedPointerBuilder) SetIsFinalMessage(value uint8) {
	message.SetUint8Fixed(m.Ptr, 25, 24, value)
}

// SetDescription The complete exchange description.
func (m ExchangeListResponseFixedPointerBuilder) SetDescription(value string) {
	message.SetStringFixed(m.Ptr, 73, 25, value)
}

// Copy
func (m ExchangeListResponseFixedPointer) Copy(to ExchangeListResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// Copy
func (m ExchangeListResponseFixedPointerBuilder) Copy(to ExchangeListResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyPointer
func (m ExchangeListResponseFixedPointer) CopyPointer(to ExchangeListResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyPointer
func (m ExchangeListResponseFixedPointerBuilder) CopyPointer(to ExchangeListResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyTo
func (m ExchangeListResponseFixedPointer) CopyTo(to ExchangeListResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyTo
func (m ExchangeListResponseFixedPointerBuilder) CopyTo(to ExchangeListResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyToPointer
func (m ExchangeListResponseFixedPointer) CopyToPointer(to ExchangeListResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyToPointer
func (m ExchangeListResponseFixedPointerBuilder) CopyToPointer(to ExchangeListResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}
