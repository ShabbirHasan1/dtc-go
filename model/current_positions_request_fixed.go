package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = CurrentPositionsRequestFixedSize  (40)
//     Type         = CURRENT_POSITIONS_REQUEST  (305)
//     RequestID    = 0
//     TradeAccount = ""
// }
var _CurrentPositionsRequestFixedDefault = []byte{40, 0, 49, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const CurrentPositionsRequestFixedSize = 40

// CurrentPositionsRequestFixed This is a message from the Client to the Server to request the current
// open Trade Positions.
type CurrentPositionsRequestFixed struct {
	message.Fixed
}

// CurrentPositionsRequestFixedBuilder This is a message from the Client to the Server to request the current
// open Trade Positions.
type CurrentPositionsRequestFixedBuilder struct {
	message.Fixed
}

func NewCurrentPositionsRequestFixedFrom(b []byte) CurrentPositionsRequestFixed {
	return CurrentPositionsRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapCurrentPositionsRequestFixed(b []byte) CurrentPositionsRequestFixed {
	return CurrentPositionsRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewCurrentPositionsRequestFixed() CurrentPositionsRequestFixedBuilder {
	a := CurrentPositionsRequestFixedBuilder{message.NewFixed(40)}
	a.Unsafe().SetBytes(0, _CurrentPositionsRequestFixedDefault)
	return a
}

// Clear
// {
//     Size         = CurrentPositionsRequestFixedSize  (40)
//     Type         = CURRENT_POSITIONS_REQUEST  (305)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m CurrentPositionsRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CurrentPositionsRequestFixedDefault)
}

// ToBuilder
func (m CurrentPositionsRequestFixed) ToBuilder() CurrentPositionsRequestFixedBuilder {
	return CurrentPositionsRequestFixedBuilder{m.Fixed}
}

// Finish
func (m CurrentPositionsRequestFixedBuilder) Finish() CurrentPositionsRequestFixed {
	return CurrentPositionsRequestFixed{m.Fixed}
}

// RequestID The request identifier. The Server will send this back in the PositionUpdate
// response messages.
func (m CurrentPositionsRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The request identifier. The Server will send this back in the PositionUpdate
// response messages.
func (m CurrentPositionsRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m CurrentPositionsRequestFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// TradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m CurrentPositionsRequestFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// SetRequestID The request identifier. The Server will send this back in the PositionUpdate
// response messages.
func (m CurrentPositionsRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m CurrentPositionsRequestFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// Copy
func (m CurrentPositionsRequestFixed) Copy(to CurrentPositionsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m CurrentPositionsRequestFixedBuilder) Copy(to CurrentPositionsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CurrentPositionsRequestFixed) CopyPointer(to CurrentPositionsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m CurrentPositionsRequestFixedBuilder) CopyPointer(to CurrentPositionsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CurrentPositionsRequestFixed) CopyToPointer(to CurrentPositionsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyToPointer
func (m CurrentPositionsRequestFixedBuilder) CopyToPointer(to CurrentPositionsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CurrentPositionsRequestFixed) CopyTo(to CurrentPositionsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m CurrentPositionsRequestFixedBuilder) CopyTo(to CurrentPositionsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
}
