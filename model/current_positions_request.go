package model

import (
	"github.com/moontrade/dtc-go/message"
)

const CurrentPositionsRequestSize = 16

const CurrentPositionsRequestFixedSize = 40

// {
//     Size         = CurrentPositionsRequestSize  (16)
//     Type         = CURRENT_POSITIONS_REQUEST  (305)
//     BaseSize     = CurrentPositionsRequestSize  (16)
//     RequestID    = 0
//     TradeAccount = ""
// }
var _CurrentPositionsRequestDefault = []byte{16, 0, 49, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size         = CurrentPositionsRequestFixedSize  (40)
//     Type         = CURRENT_POSITIONS_REQUEST  (305)
//     RequestID    = 0
//     TradeAccount = ""
// }
var _CurrentPositionsRequestFixedDefault = []byte{40, 0, 49, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func AllocCurrentPositionsRequest() CurrentPositionsRequestPointerBuilder {
	a := CurrentPositionsRequestPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _CurrentPositionsRequestDefault)
	return a
}

func AllocCurrentPositionsRequestFrom(b []byte) CurrentPositionsRequestPointer {
	return CurrentPositionsRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size         = CurrentPositionsRequestSize  (16)
//     Type         = CURRENT_POSITIONS_REQUEST  (305)
//     BaseSize     = CurrentPositionsRequestSize  (16)
//     RequestID    = 0
//     TradeAccount = ""
// }
func (m CurrentPositionsRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CurrentPositionsRequestDefault)
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
func (m CurrentPositionsRequest) ToBuilder() CurrentPositionsRequestBuilder {
	return CurrentPositionsRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m CurrentPositionsRequestFixed) ToBuilder() CurrentPositionsRequestFixedBuilder {
	return CurrentPositionsRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m CurrentPositionsRequestPointer) ToBuilder() CurrentPositionsRequestPointerBuilder {
	return CurrentPositionsRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m CurrentPositionsRequestFixedPointer) ToBuilder() CurrentPositionsRequestFixedPointerBuilder {
	return CurrentPositionsRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m CurrentPositionsRequestBuilder) Finish() CurrentPositionsRequest {
	return CurrentPositionsRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m CurrentPositionsRequestFixedBuilder) Finish() CurrentPositionsRequestFixed {
	return CurrentPositionsRequestFixed{m.Fixed}
}

// Finish
func (m *CurrentPositionsRequestPointerBuilder) Finish() CurrentPositionsRequestPointer {
	return CurrentPositionsRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *CurrentPositionsRequestFixedPointerBuilder) Finish() CurrentPositionsRequestFixedPointer {
	return CurrentPositionsRequestFixedPointer{m.FixedPointer}
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
func (m CurrentPositionsRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
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
func (m *CurrentPositionsRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m *CurrentPositionsRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRequestID The request identifier. The Server will send this back in the PositionUpdate
// response messages.
func (m CurrentPositionsRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
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
func (m CurrentPositionsRequestFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetTradeAccount This is an optional field. Leave this empty to request the Server to return
// all current open Trade Positions for all trade accounts on the logged
// in Username. Otherwise, specify a particular trade account to request
// Trade Positions for.
func (m CurrentPositionsRequestFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
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
