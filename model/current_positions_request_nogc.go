package model

import (
	"github.com/moontrade/dtc-go/message"
)

// CurrentPositionsRequestPointer This is a message from the Client to the Server to request the current
// open Trade Positions.
type CurrentPositionsRequestPointer struct {
	message.VLSPointer
}

// CurrentPositionsRequestPointerBuilder This is a message from the Client to the Server to request the current
// open Trade Positions.
type CurrentPositionsRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocCurrentPositionsRequest() CurrentPositionsRequestPointerBuilder {
	a := CurrentPositionsRequestPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _CurrentPositionsRequestDefault)
	return a
}

func AllocCurrentPositionsRequestFrom(b []byte) CurrentPositionsRequestPointer {
	return CurrentPositionsRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = CurrentPositionsRequestSize  (16)
//     Type         = CURRENT_POSITIONS_REQUEST  (305)
//     BaseSize     = CurrentPositionsRequestSize  (16)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m CurrentPositionsRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CurrentPositionsRequestDefault)
}

// ToBuilder
func (m CurrentPositionsRequestPointer) ToBuilder() CurrentPositionsRequestPointerBuilder {
	return CurrentPositionsRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *CurrentPositionsRequestPointerBuilder) Finish() CurrentPositionsRequestPointer {
	return CurrentPositionsRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID The request identifier. The Server will send this back in the PositionUpdate
// response messages.
func (m CurrentPositionsRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The request identifier. The Server will send this back in the PositionUpdate
// response messages.
func (m CurrentPositionsRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m CurrentPositionsRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m CurrentPositionsRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SetRequestID The request identifier. The Server will send this back in the PositionUpdate
// response messages.
func (m CurrentPositionsRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m *CurrentPositionsRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// Copy
func (m CurrentPositionsRequestPointer) Copy(to CurrentPositionsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m CurrentPositionsRequestPointerBuilder) Copy(to CurrentPositionsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CurrentPositionsRequestPointer) CopyPointer(to CurrentPositionsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CurrentPositionsRequestPointerBuilder) CopyPointer(to CurrentPositionsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CurrentPositionsRequestPointer) CopyTo(to CurrentPositionsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CurrentPositionsRequestPointerBuilder) CopyTo(to CurrentPositionsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CurrentPositionsRequestPointer) CopyToPointer(to CurrentPositionsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CurrentPositionsRequestPointerBuilder) CopyToPointer(to CurrentPositionsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}
