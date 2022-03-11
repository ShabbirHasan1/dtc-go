package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = CurrentPositionsRequestSize  (16)
//     Type         = CURRENT_POSITIONS_REQUEST  (305)
//     BaseSize     = CurrentPositionsRequestSize  (16)
//     RequestID    = 0
//     TradeAccount = ""
// }
var _CurrentPositionsRequestDefault = []byte{16, 0, 49, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const CurrentPositionsRequestSize = 16

// CurrentPositionsRequest This is a message from the Client to the Server to request the current
// open Trade Positions.
type CurrentPositionsRequest struct {
	message.VLS
}

// CurrentPositionsRequestBuilder This is a message from the Client to the Server to request the current
// open Trade Positions.
type CurrentPositionsRequestBuilder struct {
	message.VLSBuilder
}

func NewCurrentPositionsRequestFrom(b []byte) CurrentPositionsRequest {
	return CurrentPositionsRequest{VLS: message.NewVLSFrom(b)}
}

func WrapCurrentPositionsRequest(b []byte) CurrentPositionsRequest {
	return CurrentPositionsRequest{VLS: message.WrapVLS(b)}
}

func NewCurrentPositionsRequest() CurrentPositionsRequestBuilder {
	a := CurrentPositionsRequestBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _CurrentPositionsRequestDefault)
	return a
}

// Clear
// {
//     Size         = CurrentPositionsRequestSize  (16)
//     Type         = CURRENT_POSITIONS_REQUEST  (305)
//     BaseSize     = CurrentPositionsRequestSize  (16)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m CurrentPositionsRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CurrentPositionsRequestDefault)
}

// ToBuilder
func (m CurrentPositionsRequest) ToBuilder() CurrentPositionsRequestBuilder {
	return CurrentPositionsRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m CurrentPositionsRequestBuilder) Finish() CurrentPositionsRequest {
	return CurrentPositionsRequest{m.VLSBuilder.Finish()}
}

// RequestID The request identifier. The Server will send this back in the PositionUpdate
// response messages.
func (m CurrentPositionsRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The request identifier. The Server will send this back in the PositionUpdate
// response messages.
func (m CurrentPositionsRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m CurrentPositionsRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m CurrentPositionsRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SetRequestID The request identifier. The Server will send this back in the PositionUpdate
// response messages.
func (m CurrentPositionsRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m *CurrentPositionsRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// Copy
func (m CurrentPositionsRequest) Copy(to CurrentPositionsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m CurrentPositionsRequestBuilder) Copy(to CurrentPositionsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CurrentPositionsRequest) CopyPointer(to CurrentPositionsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CurrentPositionsRequestBuilder) CopyPointer(to CurrentPositionsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CurrentPositionsRequest) CopyToPointer(to CurrentPositionsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CurrentPositionsRequestBuilder) CopyToPointer(to CurrentPositionsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CurrentPositionsRequest) CopyTo(to CurrentPositionsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CurrentPositionsRequestBuilder) CopyTo(to CurrentPositionsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}
