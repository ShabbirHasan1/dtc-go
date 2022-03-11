package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size           = ExchangeListResponseFixedSize  (76)
//     Type           = EXCHANGE_LIST_RESPONSE  (501)
//     RequestID      = 0
//     Exchange       = ""
//     IsFinalMessage = 0
//     Description    = ""
// }
var _ExchangeListResponseFixedDefault = []byte{76, 0, 245, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const ExchangeListResponseFixedSize = 76

// ExchangeListResponseFixed The server will return this message for each supported exchange.
//
// If there are no exchanges to return in response to a request, send through
// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
// all other members in the default state and the Client will recognize there
// are no Exchanges.
type ExchangeListResponseFixed struct {
	message.Fixed
}

// ExchangeListResponseFixedBuilder The server will return this message for each supported exchange.
//
// If there are no exchanges to return in response to a request, send through
// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
// all other members in the default state and the Client will recognize there
// are no Exchanges.
type ExchangeListResponseFixedBuilder struct {
	message.Fixed
}

func NewExchangeListResponseFixedFrom(b []byte) ExchangeListResponseFixed {
	return ExchangeListResponseFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapExchangeListResponseFixed(b []byte) ExchangeListResponseFixed {
	return ExchangeListResponseFixed{Fixed: message.WrapFixed(b)}
}

func NewExchangeListResponseFixed() ExchangeListResponseFixedBuilder {
	a := ExchangeListResponseFixedBuilder{message.NewFixed(76)}
	a.Unsafe().SetBytes(0, _ExchangeListResponseFixedDefault)
	return a
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
func (m ExchangeListResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ExchangeListResponseFixedDefault)
}

// ToBuilder
func (m ExchangeListResponseFixed) ToBuilder() ExchangeListResponseFixedBuilder {
	return ExchangeListResponseFixedBuilder{m.Fixed}
}

// Finish
func (m ExchangeListResponseFixedBuilder) Finish() ExchangeListResponseFixed {
	return ExchangeListResponseFixed{m.Fixed}
}

// RequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponseFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponseFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// Exchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponseFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 24, 8)
}

// Exchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponseFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 24, 8)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseFixed) IsFinalMessage() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 25, 24)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseFixedBuilder) IsFinalMessage() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 25, 24)
}

// Description The complete exchange description.
func (m ExchangeListResponseFixed) Description() string {
	return message.StringFixed(m.Unsafe(), 73, 25)
}

// Description The complete exchange description.
func (m ExchangeListResponseFixedBuilder) Description() string {
	return message.StringFixed(m.Unsafe(), 73, 25)
}

// SetRequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponseFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetExchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponseFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 24, 8, value)
}

// SetIsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseFixedBuilder) SetIsFinalMessage(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 25, 24, value)
}

// SetDescription The complete exchange description.
func (m ExchangeListResponseFixedBuilder) SetDescription(value string) {
	message.SetStringFixed(m.Unsafe(), 73, 25, value)
}

// Copy
func (m ExchangeListResponseFixed) Copy(to ExchangeListResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// Copy
func (m ExchangeListResponseFixedBuilder) Copy(to ExchangeListResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyPointer
func (m ExchangeListResponseFixed) CopyPointer(to ExchangeListResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyPointer
func (m ExchangeListResponseFixedBuilder) CopyPointer(to ExchangeListResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyToPointer
func (m ExchangeListResponseFixed) CopyToPointer(to ExchangeListResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyToPointer
func (m ExchangeListResponseFixedBuilder) CopyToPointer(to ExchangeListResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyTo
func (m ExchangeListResponseFixed) CopyTo(to ExchangeListResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}

// CopyTo
func (m ExchangeListResponseFixedBuilder) CopyTo(to ExchangeListResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetExchange(m.Exchange())
	to.SetIsFinalMessage(m.IsFinalMessage())
	to.SetDescription(m.Description())
}
