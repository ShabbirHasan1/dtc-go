package model

import (
	"github.com/moontrade/dtc-go/message"
)

// CurrentPositionsRequestFixedPointer This is a message from the Client to the Server to request the current
// open Trade Positions.
type CurrentPositionsRequestFixedPointer struct {
	message.FixedPointer
}

// CurrentPositionsRequestFixedPointerBuilder This is a message from the Client to the Server to request the current
// open Trade Positions.
type CurrentPositionsRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocCurrentPositionsRequestFixed() CurrentPositionsRequestFixedPointerBuilder {
	a := CurrentPositionsRequestFixedPointerBuilder{message.AllocFixed(40)}
	a.Ptr.SetBytes(0, _CurrentPositionsRequestFixedDefault)
	return a
}

func AllocCurrentPositionsRequestFixedFrom(b []byte) CurrentPositionsRequestFixedPointer {
	return CurrentPositionsRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size         = CurrentPositionsRequestFixedSize  (40)
//     Type         = CURRENT_POSITIONS_REQUEST  (305)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m CurrentPositionsRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CurrentPositionsRequestFixedDefault)
}

// ToBuilder
func (m CurrentPositionsRequestFixedPointer) ToBuilder() CurrentPositionsRequestFixedPointerBuilder {
	return CurrentPositionsRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *CurrentPositionsRequestFixedPointerBuilder) Finish() CurrentPositionsRequestFixedPointer {
	return CurrentPositionsRequestFixedPointer{m.FixedPointer}
}

// RequestID The request identifier. The Server will send this back in the PositionUpdate
// response messages.
func (m CurrentPositionsRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The request identifier. The Server will send this back in the PositionUpdate
// response messages.
func (m CurrentPositionsRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m CurrentPositionsRequestFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m CurrentPositionsRequestFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// SetRequestID The request identifier. The Server will send this back in the PositionUpdate
// response messages.
func (m CurrentPositionsRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m CurrentPositionsRequestFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// Copy
func (m CurrentPositionsRequestFixedPointer) Copy(to CurrentPositionsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m CurrentPositionsRequestFixedPointerBuilder) Copy(to CurrentPositionsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CurrentPositionsRequestFixedPointer) CopyPointer(to CurrentPositionsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CurrentPositionsRequestFixedPointerBuilder) CopyPointer(to CurrentPositionsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CurrentPositionsRequestFixedPointer) CopyTo(to CurrentPositionsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CurrentPositionsRequestFixedPointerBuilder) CopyTo(to CurrentPositionsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CurrentPositionsRequestFixedPointer) CopyToPointer(to CurrentPositionsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CurrentPositionsRequestFixedPointerBuilder) CopyToPointer(to CurrentPositionsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}
