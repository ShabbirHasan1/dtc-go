package model

import (
	"github.com/moontrade/dtc-go/message"
)

const ExchangeListResponseSize = 24

const ExchangeListResponseFixedSize = 76

// {
//     Size           = ExchangeListResponseSize  (24)
//     Type           = EXCHANGE_LIST_RESPONSE  (501)
//     BaseSize       = ExchangeListResponseSize  (24)
//     RequestID      = 0
//     Exchange       = ""
//     IsFinalMessage = false
//     Description    = ""
// }
var _ExchangeListResponseDefault = []byte{24, 0, 245, 1, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size           = ExchangeListResponseFixedSize  (76)
//     Type           = EXCHANGE_LIST_RESPONSE  (501)
//     RequestID      = 0
//     Exchange       = ""
//     IsFinalMessage = false
//     Description    = ""
// }
var _ExchangeListResponseFixedDefault = []byte{76, 0, 245, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func AllocExchangeListResponse() ExchangeListResponsePointerBuilder {
	a := ExchangeListResponsePointerBuilder{message.AllocVLSBuilder(24)}
	a.Ptr.SetBytes(0, _ExchangeListResponseDefault)
	return a
}

func AllocExchangeListResponseFrom(b []byte) ExchangeListResponsePointer {
	return ExchangeListResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size           = ExchangeListResponseSize  (24)
//     Type           = EXCHANGE_LIST_RESPONSE  (501)
//     BaseSize       = ExchangeListResponseSize  (24)
//     RequestID      = 0
//     Exchange       = ""
//     IsFinalMessage = false
//     Description    = ""
// }
func (m ExchangeListResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ExchangeListResponseDefault)
}

// Clear
// {
//     Size           = ExchangeListResponseFixedSize  (76)
//     Type           = EXCHANGE_LIST_RESPONSE  (501)
//     RequestID      = 0
//     Exchange       = ""
//     IsFinalMessage = false
//     Description    = ""
// }
func (m ExchangeListResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ExchangeListResponseFixedDefault)
}

// Clear
// {
//     Size           = ExchangeListResponseSize  (24)
//     Type           = EXCHANGE_LIST_RESPONSE  (501)
//     BaseSize       = ExchangeListResponseSize  (24)
//     RequestID      = 0
//     Exchange       = ""
//     IsFinalMessage = false
//     Description    = ""
// }
func (m ExchangeListResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ExchangeListResponseDefault)
}

// Clear
// {
//     Size           = ExchangeListResponseFixedSize  (76)
//     Type           = EXCHANGE_LIST_RESPONSE  (501)
//     RequestID      = 0
//     Exchange       = ""
//     IsFinalMessage = false
//     Description    = ""
// }
func (m ExchangeListResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ExchangeListResponseFixedDefault)
}

// ToBuilder
func (m ExchangeListResponse) ToBuilder() ExchangeListResponseBuilder {
	return ExchangeListResponseBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m ExchangeListResponseFixed) ToBuilder() ExchangeListResponseFixedBuilder {
	return ExchangeListResponseFixedBuilder{m.Fixed}
}

// ToBuilder
func (m ExchangeListResponsePointer) ToBuilder() ExchangeListResponsePointerBuilder {
	return ExchangeListResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m ExchangeListResponseFixedPointer) ToBuilder() ExchangeListResponseFixedPointerBuilder {
	return ExchangeListResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m ExchangeListResponseBuilder) Finish() ExchangeListResponse {
	return ExchangeListResponse{m.VLSBuilder.Finish()}
}

// Finish
func (m ExchangeListResponseFixedBuilder) Finish() ExchangeListResponseFixed {
	return ExchangeListResponseFixed{m.Fixed}
}

// Finish
func (m *ExchangeListResponsePointerBuilder) Finish() ExchangeListResponsePointer {
	return ExchangeListResponsePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *ExchangeListResponseFixedPointerBuilder) Finish() ExchangeListResponseFixedPointer {
	return ExchangeListResponseFixedPointer{m.FixedPointer}
}

// RequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponse) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponseBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
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
func (m ExchangeListResponse) Exchange() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Exchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponseBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
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
func (m ExchangeListResponse) IsFinalMessage() bool {
	return message.BoolVLS(m.Unsafe(), 17, 16)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseBuilder) IsFinalMessage() bool {
	return message.BoolVLS(m.Unsafe(), 17, 16)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponsePointer) IsFinalMessage() bool {
	return message.BoolVLS(m.Ptr, 17, 16)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponsePointerBuilder) IsFinalMessage() bool {
	return message.BoolVLS(m.Ptr, 17, 16)
}

// Description The complete exchange description.
func (m ExchangeListResponse) Description() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// Description The complete exchange description.
func (m ExchangeListResponseBuilder) Description() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// Description The complete exchange description.
func (m ExchangeListResponsePointer) Description() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// Description The complete exchange description.
func (m ExchangeListResponsePointerBuilder) Description() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// RequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponseFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponseFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
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
func (m ExchangeListResponseFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 24, 8)
}

// Exchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponseFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 24, 8)
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
func (m ExchangeListResponseFixed) IsFinalMessage() bool {
	return message.BoolFixed(m.Unsafe(), 25, 24)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseFixedBuilder) IsFinalMessage() bool {
	return message.BoolFixed(m.Unsafe(), 25, 24)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseFixedPointer) IsFinalMessage() bool {
	return message.BoolFixed(m.Ptr, 25, 24)
}

// IsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseFixedPointerBuilder) IsFinalMessage() bool {
	return message.BoolFixed(m.Ptr, 25, 24)
}

// Description The complete exchange description.
func (m ExchangeListResponseFixed) Description() string {
	return message.StringFixed(m.Unsafe(), 73, 25)
}

// Description The complete exchange description.
func (m ExchangeListResponseFixedBuilder) Description() string {
	return message.StringFixed(m.Unsafe(), 73, 25)
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
func (m ExchangeListResponseBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetExchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m *ExchangeListResponseBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetExchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m *ExchangeListResponsePointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetIsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseBuilder) SetIsFinalMessage(value bool) {
	message.SetBoolVLS(m.Unsafe(), 17, 16, value)
}

// SetIsFinalMessage 1 = final message in batch.
func (m ExchangeListResponsePointerBuilder) SetIsFinalMessage(value bool) {
	message.SetBoolVLS(m.Ptr, 17, 16, value)
}

// SetDescription The complete exchange description.
func (m *ExchangeListResponseBuilder) SetDescription(value string) {
	message.SetStringVLS(&m.VLSBuilder, 22, 18, value)
}

// SetDescription The complete exchange description.
func (m *ExchangeListResponsePointerBuilder) SetDescription(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 22, 18, value)
}

// SetRequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponseFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID The RequestID sent in the request by the Client.
func (m ExchangeListResponseFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetExchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponseFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 24, 8, value)
}

// SetExchange The exchange identifier that the Server uses to identify a particular
// exchange.
func (m ExchangeListResponseFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 24, 8, value)
}

// SetIsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseFixedBuilder) SetIsFinalMessage(value bool) {
	message.SetBoolFixed(m.Unsafe(), 25, 24, value)
}

// SetIsFinalMessage 1 = final message in batch.
func (m ExchangeListResponseFixedPointerBuilder) SetIsFinalMessage(value bool) {
	message.SetBoolFixed(m.Ptr, 25, 24, value)
}

// SetDescription The complete exchange description.
func (m ExchangeListResponseFixedBuilder) SetDescription(value string) {
	message.SetStringFixed(m.Unsafe(), 73, 25, value)
}

// SetDescription The complete exchange description.
func (m ExchangeListResponseFixedPointerBuilder) SetDescription(value string) {
	message.SetStringFixed(m.Ptr, 73, 25, value)
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
