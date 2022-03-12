package model

import (
	"github.com/moontrade/dtc-go/message"
)

const TradingSymbolStatusSize = 12

// {
//     Size     = TradingSymbolStatusSize  (12)
//     Type     = TRADING_SYMBOL_STATUS  (138)
//     SymbolID = 0
//     Status   = TRADING_STATUS_UNKNOWN  (0)
// }
var _TradingSymbolStatusDefault = []byte{12, 0, 138, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// TradingSymbolStatus Sent by the Server to the Client to indicate the status of the symbol
// in regards to whether trading is open or closed or some other intermediate
// state.
type TradingSymbolStatus struct {
	message.Fixed
}

// TradingSymbolStatusBuilder Sent by the Server to the Client to indicate the status of the symbol
// in regards to whether trading is open or closed or some other intermediate
// state.
type TradingSymbolStatusBuilder struct {
	message.Fixed
}

// TradingSymbolStatusPointer Sent by the Server to the Client to indicate the status of the symbol
// in regards to whether trading is open or closed or some other intermediate
// state.
type TradingSymbolStatusPointer struct {
	message.FixedPointer
}

// TradingSymbolStatusPointerBuilder Sent by the Server to the Client to indicate the status of the symbol
// in regards to whether trading is open or closed or some other intermediate
// state.
type TradingSymbolStatusPointerBuilder struct {
	message.FixedPointer
}

func NewTradingSymbolStatusFrom(b []byte) TradingSymbolStatus {
	return TradingSymbolStatus{Fixed: message.NewFixedFrom(b)}
}

func WrapTradingSymbolStatus(b []byte) TradingSymbolStatus {
	return TradingSymbolStatus{Fixed: message.WrapFixed(b)}
}

func NewTradingSymbolStatus() TradingSymbolStatusBuilder {
	a := TradingSymbolStatusBuilder{message.NewFixed(12)}
	a.Unsafe().SetBytes(0, _TradingSymbolStatusDefault)
	return a
}

func AllocTradingSymbolStatus() TradingSymbolStatusPointerBuilder {
	a := TradingSymbolStatusPointerBuilder{message.AllocFixed(12)}
	a.Ptr.SetBytes(0, _TradingSymbolStatusDefault)
	return a
}

func AllocTradingSymbolStatusFrom(b []byte) TradingSymbolStatusPointer {
	return TradingSymbolStatusPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size     = TradingSymbolStatusSize  (12)
//     Type     = TRADING_SYMBOL_STATUS  (138)
//     SymbolID = 0
//     Status   = TRADING_STATUS_UNKNOWN  (0)
// }
func (m TradingSymbolStatusBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradingSymbolStatusDefault)
}

// Clear
// {
//     Size     = TradingSymbolStatusSize  (12)
//     Type     = TRADING_SYMBOL_STATUS  (138)
//     SymbolID = 0
//     Status   = TRADING_STATUS_UNKNOWN  (0)
// }
func (m TradingSymbolStatusPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradingSymbolStatusDefault)
}

// ToBuilder
func (m TradingSymbolStatus) ToBuilder() TradingSymbolStatusBuilder {
	return TradingSymbolStatusBuilder{m.Fixed}
}

// ToBuilder
func (m TradingSymbolStatusPointer) ToBuilder() TradingSymbolStatusPointerBuilder {
	return TradingSymbolStatusPointerBuilder{m.FixedPointer}
}

// Finish
func (m TradingSymbolStatusBuilder) Finish() TradingSymbolStatus {
	return TradingSymbolStatus{m.Fixed}
}

// Finish
func (m *TradingSymbolStatusPointerBuilder) Finish() TradingSymbolStatusPointer {
	return TradingSymbolStatusPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the status in this message
// is for.
func (m TradingSymbolStatus) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the status in this message
// is for.
func (m TradingSymbolStatusBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the status in this message
// is for.
func (m TradingSymbolStatusPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the status in this message
// is for.
func (m TradingSymbolStatusPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Status The current trading status for the symbol. Can be one of the following.
// The current trading status for the symbol. Can be one of the following.
//
// TRADING_STATUS_UNKNOWN = 0
// TRADING_STATUS_PRE_OPEN = 1
// TRADING_STATUS_OPEN = 2
// TRADING_STATUS_CLOSE = 3
// TRADING_STATUS_TRADING_HALT = 4
func (m TradingSymbolStatus) Status() TradingStatusEnum {
	return TradingStatusEnum(message.Int8Fixed(m.Unsafe(), 9, 8))
}

// Status The current trading status for the symbol. Can be one of the following.
// The current trading status for the symbol. Can be one of the following.
//
// TRADING_STATUS_UNKNOWN = 0
// TRADING_STATUS_PRE_OPEN = 1
// TRADING_STATUS_OPEN = 2
// TRADING_STATUS_CLOSE = 3
// TRADING_STATUS_TRADING_HALT = 4
func (m TradingSymbolStatusBuilder) Status() TradingStatusEnum {
	return TradingStatusEnum(message.Int8Fixed(m.Unsafe(), 9, 8))
}

// Status The current trading status for the symbol. Can be one of the following.
// The current trading status for the symbol. Can be one of the following.
//
// TRADING_STATUS_UNKNOWN = 0
// TRADING_STATUS_PRE_OPEN = 1
// TRADING_STATUS_OPEN = 2
// TRADING_STATUS_CLOSE = 3
// TRADING_STATUS_TRADING_HALT = 4
func (m TradingSymbolStatusPointer) Status() TradingStatusEnum {
	return TradingStatusEnum(message.Int8Fixed(m.Ptr, 9, 8))
}

// Status The current trading status for the symbol. Can be one of the following.
// The current trading status for the symbol. Can be one of the following.
//
// TRADING_STATUS_UNKNOWN = 0
// TRADING_STATUS_PRE_OPEN = 1
// TRADING_STATUS_OPEN = 2
// TRADING_STATUS_CLOSE = 3
// TRADING_STATUS_TRADING_HALT = 4
func (m TradingSymbolStatusPointerBuilder) Status() TradingStatusEnum {
	return TradingStatusEnum(message.Int8Fixed(m.Ptr, 9, 8))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the status in this message
// is for.
func (m TradingSymbolStatusBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the status in this message
// is for.
func (m TradingSymbolStatusPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetStatus The current trading status for the symbol. Can be one of the following.
// The current trading status for the symbol. Can be one of the following.
//
// TRADING_STATUS_UNKNOWN = 0
// TRADING_STATUS_PRE_OPEN = 1
// TRADING_STATUS_OPEN = 2
// TRADING_STATUS_CLOSE = 3
// TRADING_STATUS_TRADING_HALT = 4
func (m TradingSymbolStatusBuilder) SetStatus(value TradingStatusEnum) {
	message.SetInt8Fixed(m.Unsafe(), 9, 8, int8(value))
}

// SetStatus The current trading status for the symbol. Can be one of the following.
// The current trading status for the symbol. Can be one of the following.
//
// TRADING_STATUS_UNKNOWN = 0
// TRADING_STATUS_PRE_OPEN = 1
// TRADING_STATUS_OPEN = 2
// TRADING_STATUS_CLOSE = 3
// TRADING_STATUS_TRADING_HALT = 4
func (m TradingSymbolStatusPointerBuilder) SetStatus(value TradingStatusEnum) {
	message.SetInt8Fixed(m.Ptr, 9, 8, int8(value))
}

// Copy
func (m TradingSymbolStatus) Copy(to TradingSymbolStatusBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

// Copy
func (m TradingSymbolStatusBuilder) Copy(to TradingSymbolStatusBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

// CopyPointer
func (m TradingSymbolStatus) CopyPointer(to TradingSymbolStatusPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

// CopyPointer
func (m TradingSymbolStatusBuilder) CopyPointer(to TradingSymbolStatusPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

// Copy
func (m TradingSymbolStatusPointer) Copy(to TradingSymbolStatusBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

// Copy
func (m TradingSymbolStatusPointerBuilder) Copy(to TradingSymbolStatusBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

// CopyPointer
func (m TradingSymbolStatusPointer) CopyPointer(to TradingSymbolStatusPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}

// CopyPointer
func (m TradingSymbolStatusPointerBuilder) CopyPointer(to TradingSymbolStatusPointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetStatus(m.Status())
}
