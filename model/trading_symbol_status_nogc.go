package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
//     Status   = 0
// }
func (m TradingSymbolStatusPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradingSymbolStatusDefault)
}

// ToBuilder
func (m TradingSymbolStatusPointer) ToBuilder() TradingSymbolStatusPointerBuilder {
	return TradingSymbolStatusPointerBuilder{m.FixedPointer}
}

// Finish
func (m *TradingSymbolStatusPointerBuilder) Finish() TradingSymbolStatusPointer {
	return TradingSymbolStatusPointer{m.FixedPointer}
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
func (m TradingSymbolStatusPointerBuilder) SetStatus(value TradingStatusEnum) {
	message.SetInt8Fixed(m.Ptr, 9, 8, int8(value))
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
