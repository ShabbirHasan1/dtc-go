package model

import (
	"github.com/moontrade/dtc-go/message"
)

// PositionUpdateFixedPointer This is a message from the Server to the Client to report a Trade Position
// for a symbol in any Trade Account for the logged in Username.
//
// The Position Update message can either be solicited, in response to CurrentPositionsRequest.
// Or unsolicited as a Trade Position for a symbol changes during the connection
// to the Server. Each Trade Position is contained within a single message.
// to the Server. Each Trade Position is contained within a single message.
//
// When the server is responding with one or more PositionUpdate messages
// in response to a CurrentPositionsRequest message, it must not send any
// unsolicited PositionUpdate messages interleaved with the solicited PositionUpdate
// messages in response to the CurrentPositionsRequest message.
type PositionUpdateFixedPointer struct {
	message.FixedPointer
}

// PositionUpdateFixedPointerBuilder This is a message from the Server to the Client to report a Trade Position
// for a symbol in any Trade Account for the logged in Username.
//
// The Position Update message can either be solicited, in response to CurrentPositionsRequest.
// Or unsolicited as a Trade Position for a symbol changes during the connection
// to the Server. Each Trade Position is contained within a single message.
// to the Server. Each Trade Position is contained within a single message.
//
// When the server is responding with one or more PositionUpdate messages
// in response to a CurrentPositionsRequest message, it must not send any
// unsolicited PositionUpdate messages interleaved with the solicited PositionUpdate
// messages in response to the CurrentPositionsRequest message.
type PositionUpdateFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocPositionUpdateFixed() PositionUpdateFixedPointerBuilder {
	a := PositionUpdateFixedPointerBuilder{message.AllocFixed(240)}
	a.Ptr.SetBytes(0, _PositionUpdateFixedDefault)
	return a
}

func AllocPositionUpdateFixedFrom(b []byte) PositionUpdateFixedPointer {
	return PositionUpdateFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                        = PositionUpdateFixedSize  (240)
//     Type                        = POSITION_UPDATE  (306)
//     RequestID                   = 0
//     TotalNumberMessages         = 0
//     MessageNumber               = 0
//     Symbol                      = ""
//     Exchange                    = ""
//     Quantity                    = 0
//     AveragePrice                = 0
//     PositionIdentifier          = ""
//     TradeAccount                = ""
//     NoPositions                 = 0
//     Unsolicited                 = 0
//     MarginRequirement           = 0
//     EntryDateTime               = 0
//     OpenProfitLoss              = 0
//     HighPriceDuringPosition     = 0
//     LowPriceDuringPosition      = 0
//     QuantityLimit               = 0
//     MaxPotentialPostionQuantity = 0
// }
func (m PositionUpdateFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _PositionUpdateFixedDefault)
}

// ToBuilder
func (m PositionUpdateFixedPointer) ToBuilder() PositionUpdateFixedPointerBuilder {
	return PositionUpdateFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *PositionUpdateFixedPointerBuilder) Finish() PositionUpdateFixedPointer {
	return PositionUpdateFixedPointer{m.FixedPointer}
}

// RequestID The Server sets this to 0, the default, if this is a real-time Trade Position
// update.
//
// Otherwise, when the Server is sending Trade Positions in response to a
// CurrentPositionsRequest message, it must set this to the RequestID given
// in the CurrentPositionsRequest message
func (m PositionUpdateFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The Server sets this to 0, the default, if this is a real-time Trade Position
// update.
//
// Otherwise, when the Server is sending Trade Positions in response to a
// CurrentPositionsRequest message, it must set this to the RequestID given
// in the CurrentPositionsRequest message
func (m PositionUpdateFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// TotalNumberMessages This indicates the total number of Position Update messages when a batch
// of messages is being sent. If there is only one Position Update message
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateFixedPointer) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// TotalNumberMessages This indicates the total number of Position Update messages when a batch
// of messages is being sent. If there is only one Position Update message
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateFixedPointerBuilder) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// MessageNumber This indicates the 1-based index of the Position Update message when a
// batch of messages is being sent. If there is only one Position Update
// message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateFixedPointer) MessageNumber() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// MessageNumber This indicates the 1-based index of the Position Update message when a
// batch of messages is being sent. If there is only one Position Update
// message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateFixedPointerBuilder) MessageNumber() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// Symbol The symbol for the Position.
func (m PositionUpdateFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 80, 16)
}

// Symbol The symbol for the Position.
func (m PositionUpdateFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 80, 16)
}

// Exchange The optional exchange for the symbol.
func (m PositionUpdateFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 96, 80)
}

// Exchange The optional exchange for the symbol.
func (m PositionUpdateFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 96, 80)
}

// Quantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdateFixedPointer) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 104, 96)
}

// Quantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdateFixedPointerBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 104, 96)
}

// AveragePrice The average position price.
func (m PositionUpdateFixedPointer) AveragePrice() float64 {
	return message.Float64Fixed(m.Ptr, 112, 104)
}

// AveragePrice The average position price.
func (m PositionUpdateFixedPointerBuilder) AveragePrice() float64 {
	return message.Float64Fixed(m.Ptr, 112, 104)
}

// PositionIdentifier When the Server sets LogonResponse::UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1, then it must set PositionIdentifier to a unique identifier to identify
// the Trade Position for the particular Symbol and Trade Account for which
// it is for.
//
// Otherwise, this field is optional and it is recommended for a Server not
// to use it if it does not provide multiple positions for a Symbol and Trade
// Account.
//
// This field identifies an individual Trade Position in the case of where
// there are multiple Positions for a particular Symbol and Trade Account.
// there are multiple Positions for a particular Symbol and Trade Account.
//
// A Client must implement support for PositionIdentifier when LogonResponse::UsesMultiplePositionsPerSymbolAndTradeAccount
// is set to 1.
func (m PositionUpdateFixedPointer) PositionIdentifier() string {
	return message.StringFixed(m.Ptr, 144, 112)
}

// PositionIdentifier When the Server sets LogonResponse::UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1, then it must set PositionIdentifier to a unique identifier to identify
// the Trade Position for the particular Symbol and Trade Account for which
// it is for.
//
// Otherwise, this field is optional and it is recommended for a Server not
// to use it if it does not provide multiple positions for a Symbol and Trade
// Account.
//
// This field identifies an individual Trade Position in the case of where
// there are multiple Positions for a particular Symbol and Trade Account.
// there are multiple Positions for a particular Symbol and Trade Account.
//
// A Client must implement support for PositionIdentifier when LogonResponse::UsesMultiplePositionsPerSymbolAndTradeAccount
// is set to 1.
func (m PositionUpdateFixedPointerBuilder) PositionIdentifier() string {
	return message.StringFixed(m.Ptr, 144, 112)
}

// TradeAccount The trade account the Trade Position is in for the Symbol.
func (m PositionUpdateFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 176, 144)
}

// TradeAccount The trade account the Trade Position is in for the Symbol.
func (m PositionUpdateFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 176, 144)
}

// NoPositions Set to an integer value of 1 to indicate there are no Trade Positions
// in response to a CurrentPositionsRequest message.
//
// Otherwise, the Server must leave this at the default of 0. The Server
// is always expected to respond with a single PositionUpdate message when
// there are no Trade Positions for any Symbol when current Trade Positions
// have been requested by the Client with CurrentPositionsRequest.
//
// When the Server is sending a PositionUpdate message to the Client and
// it is indicating that the Quantity field is 0, then the NoPositions field
// must be left at the default of 0. It is not used to indicate a Quantity
// of 0 for a particular Symbol and TradeAccount.
//
// This is always set to the default of 0 for an unsolicited Trade Position
// Update.
func (m PositionUpdateFixedPointer) NoPositions() uint8 {
	return message.Uint8Fixed(m.Ptr, 177, 176)
}

// NoPositions Set to an integer value of 1 to indicate there are no Trade Positions
// in response to a CurrentPositionsRequest message.
//
// Otherwise, the Server must leave this at the default of 0. The Server
// is always expected to respond with a single PositionUpdate message when
// there are no Trade Positions for any Symbol when current Trade Positions
// have been requested by the Client with CurrentPositionsRequest.
//
// When the Server is sending a PositionUpdate message to the Client and
// it is indicating that the Quantity field is 0, then the NoPositions field
// must be left at the default of 0. It is not used to indicate a Quantity
// of 0 for a particular Symbol and TradeAccount.
//
// This is always set to the default of 0 for an unsolicited Trade Position
// Update.
func (m PositionUpdateFixedPointerBuilder) NoPositions() uint8 {
	return message.Uint8Fixed(m.Ptr, 177, 176)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdateFixedPointer) Unsolicited() uint8 {
	return message.Uint8Fixed(m.Ptr, 178, 177)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdateFixedPointerBuilder) Unsolicited() uint8 {
	return message.Uint8Fixed(m.Ptr, 178, 177)
}

// MarginRequirement is the required margin as a currency value for the current
// trade Position Quantity and any working orders for the Trade Account.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointer) MarginRequirement() float64 {
	return message.Float64Fixed(m.Ptr, 192, 184)
}

// MarginRequirement is the required margin as a currency value for the current
// trade Position Quantity and any working orders for the Trade Account.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointerBuilder) MarginRequirement() float64 {
	return message.Float64Fixed(m.Ptr, 192, 184)
}

// EntryDateTime is the Date-Time of the initial entry of the Trade Position.
// It is in the UTC time zone.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointer) EntryDateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 196, 192))
}

// EntryDateTime is the Date-Time of the initial entry of the Trade Position.
// It is in the UTC time zone.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointerBuilder) EntryDateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Ptr, 196, 192))
}

// OpenProfitLoss is the current open Trade Position profit or loss as a
// currency value.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointer) OpenProfitLoss() float64 {
	return message.Float64Fixed(m.Ptr, 208, 200)
}

// OpenProfitLoss is the current open Trade Position profit or loss as a
// currency value.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointerBuilder) OpenProfitLoss() float64 {
	return message.Float64Fixed(m.Ptr, 208, 200)
}

// HighPriceDuringPosition is the highest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointer) HighPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Ptr, 216, 208)
}

// HighPriceDuringPosition is the highest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointerBuilder) HighPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Ptr, 216, 208)
}

// LowPriceDuringPosition is the lowest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointer) LowPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Ptr, 224, 216)
}

// LowPriceDuringPosition is the lowest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointerBuilder) LowPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Ptr, 224, 216)
}

// QuantityLimit This is the limit or maximum Trade Position Quantity possible for a Trade
// Position for the Trade Account and Symbol. This applies equally to a long
// or short position.
//
// This is only an informational value provided from the risk management
// system.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointer) QuantityLimit() float64 {
	return message.Float64Fixed(m.Ptr, 232, 224)
}

// QuantityLimit This is the limit or maximum Trade Position Quantity possible for a Trade
// Position for the Trade Account and Symbol. This applies equally to a long
// or short position.
//
// This is only an informational value provided from the risk management
// system.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointerBuilder) QuantityLimit() float64 {
	return message.Float64Fixed(m.Ptr, 232, 224)
}

// MaxPotentialPostionQuantity
func (m PositionUpdateFixedPointer) MaxPotentialPostionQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 240, 232)
}

// MaxPotentialPostionQuantity
func (m PositionUpdateFixedPointerBuilder) MaxPotentialPostionQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 240, 232)
}

// SetRequestID The Server sets this to 0, the default, if this is a real-time Trade Position
// update.
//
// Otherwise, when the Server is sending Trade Positions in response to a
// CurrentPositionsRequest message, it must set this to the RequestID given
// in the CurrentPositionsRequest message
func (m PositionUpdateFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetTotalNumberMessages This indicates the total number of Position Update messages when a batch
// of messages is being sent. If there is only one Position Update message
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateFixedPointerBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32Fixed(m.Ptr, 12, 8, value)
}

// SetMessageNumber This indicates the 1-based index of the Position Update message when a
// batch of messages is being sent. If there is only one Position Update
// message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateFixedPointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Ptr, 16, 12, value)
}

// SetSymbol The symbol for the Position.
func (m PositionUpdateFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 80, 16, value)
}

// SetExchange The optional exchange for the symbol.
func (m PositionUpdateFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 96, 80, value)
}

// SetQuantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdateFixedPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 104, 96, value)
}

// SetAveragePrice The average position price.
func (m PositionUpdateFixedPointerBuilder) SetAveragePrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 112, 104, value)
}

// SetPositionIdentifier When the Server sets LogonResponse::UsesMultiplePositionsPerSymbolAndTradeAccount
// to 1, then it must set PositionIdentifier to a unique identifier to identify
// the Trade Position for the particular Symbol and Trade Account for which
// it is for.
//
// Otherwise, this field is optional and it is recommended for a Server not
// to use it if it does not provide multiple positions for a Symbol and Trade
// Account.
//
// This field identifies an individual Trade Position in the case of where
// there are multiple Positions for a particular Symbol and Trade Account.
// there are multiple Positions for a particular Symbol and Trade Account.
//
// A Client must implement support for PositionIdentifier when LogonResponse::UsesMultiplePositionsPerSymbolAndTradeAccount
// is set to 1.
func (m PositionUpdateFixedPointerBuilder) SetPositionIdentifier(value string) {
	message.SetStringFixed(m.Ptr, 144, 112, value)
}

// SetTradeAccount The trade account the Trade Position is in for the Symbol.
func (m PositionUpdateFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 176, 144, value)
}

// SetNoPositions Set to an integer value of 1 to indicate there are no Trade Positions
// in response to a CurrentPositionsRequest message.
//
// Otherwise, the Server must leave this at the default of 0. The Server
// is always expected to respond with a single PositionUpdate message when
// there are no Trade Positions for any Symbol when current Trade Positions
// have been requested by the Client with CurrentPositionsRequest.
//
// When the Server is sending a PositionUpdate message to the Client and
// it is indicating that the Quantity field is 0, then the NoPositions field
// must be left at the default of 0. It is not used to indicate a Quantity
// of 0 for a particular Symbol and TradeAccount.
//
// This is always set to the default of 0 for an unsolicited Trade Position
// Update.
func (m PositionUpdateFixedPointerBuilder) SetNoPositions(value uint8) {
	message.SetUint8Fixed(m.Ptr, 177, 176, value)
}

// SetUnsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdateFixedPointerBuilder) SetUnsolicited(value uint8) {
	message.SetUint8Fixed(m.Ptr, 178, 177, value)
}

// SetMarginRequirement MarginRequirement is the required margin as a currency value for the current
// trade Position Quantity and any working orders for the Trade Account.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointerBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64Fixed(m.Ptr, 192, 184, value)
}

// SetEntryDateTime EntryDateTime is the Date-Time of the initial entry of the Trade Position.
// It is in the UTC time zone.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointerBuilder) SetEntryDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Ptr, 196, 192, uint32(value))
}

// SetOpenProfitLoss OpenProfitLoss is the current open Trade Position profit or loss as a
// currency value.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointerBuilder) SetOpenProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Ptr, 208, 200, value)
}

// SetHighPriceDuringPosition HighPriceDuringPosition is the highest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointerBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Ptr, 216, 208, value)
}

// SetLowPriceDuringPosition LowPriceDuringPosition is the lowest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointerBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Ptr, 224, 216, value)
}

// SetQuantityLimit This is the limit or maximum Trade Position Quantity possible for a Trade
// Position for the Trade Account and Symbol. This applies equally to a long
// or short position.
//
// This is only an informational value provided from the risk management
// system.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedPointerBuilder) SetQuantityLimit(value float64) {
	message.SetFloat64Fixed(m.Ptr, 232, 224, value)
}

// SetMaxPotentialPostionQuantity
func (m PositionUpdateFixedPointerBuilder) SetMaxPotentialPostionQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 240, 232, value)
}

// Copy
func (m PositionUpdateFixedPointer) Copy(to PositionUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetPositionIdentifier(m.PositionIdentifier())
	to.SetTradeAccount(m.TradeAccount())
	to.SetNoPositions(m.NoPositions())
	to.SetUnsolicited(m.Unsolicited())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetQuantityLimit(m.QuantityLimit())
	to.SetMaxPotentialPostionQuantity(m.MaxPotentialPostionQuantity())
}

// Copy
func (m PositionUpdateFixedPointerBuilder) Copy(to PositionUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetPositionIdentifier(m.PositionIdentifier())
	to.SetTradeAccount(m.TradeAccount())
	to.SetNoPositions(m.NoPositions())
	to.SetUnsolicited(m.Unsolicited())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetQuantityLimit(m.QuantityLimit())
	to.SetMaxPotentialPostionQuantity(m.MaxPotentialPostionQuantity())
}

// CopyPointer
func (m PositionUpdateFixedPointer) CopyPointer(to PositionUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetPositionIdentifier(m.PositionIdentifier())
	to.SetTradeAccount(m.TradeAccount())
	to.SetNoPositions(m.NoPositions())
	to.SetUnsolicited(m.Unsolicited())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetQuantityLimit(m.QuantityLimit())
	to.SetMaxPotentialPostionQuantity(m.MaxPotentialPostionQuantity())
}

// CopyPointer
func (m PositionUpdateFixedPointerBuilder) CopyPointer(to PositionUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetPositionIdentifier(m.PositionIdentifier())
	to.SetTradeAccount(m.TradeAccount())
	to.SetNoPositions(m.NoPositions())
	to.SetUnsolicited(m.Unsolicited())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetQuantityLimit(m.QuantityLimit())
	to.SetMaxPotentialPostionQuantity(m.MaxPotentialPostionQuantity())
}

// CopyTo
func (m PositionUpdateFixedPointer) CopyTo(to PositionUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetPositionIdentifier(m.PositionIdentifier())
	to.SetTradeAccount(m.TradeAccount())
	to.SetNoPositions(m.NoPositions())
	to.SetUnsolicited(m.Unsolicited())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetQuantityLimit(m.QuantityLimit())
	to.SetMaxPotentialPostionQuantity(m.MaxPotentialPostionQuantity())
}

// CopyTo
func (m PositionUpdateFixedPointerBuilder) CopyTo(to PositionUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetPositionIdentifier(m.PositionIdentifier())
	to.SetTradeAccount(m.TradeAccount())
	to.SetNoPositions(m.NoPositions())
	to.SetUnsolicited(m.Unsolicited())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetQuantityLimit(m.QuantityLimit())
	to.SetMaxPotentialPostionQuantity(m.MaxPotentialPostionQuantity())
}

// CopyToPointer
func (m PositionUpdateFixedPointer) CopyToPointer(to PositionUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetPositionIdentifier(m.PositionIdentifier())
	to.SetTradeAccount(m.TradeAccount())
	to.SetNoPositions(m.NoPositions())
	to.SetUnsolicited(m.Unsolicited())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetQuantityLimit(m.QuantityLimit())
	to.SetMaxPotentialPostionQuantity(m.MaxPotentialPostionQuantity())
}

// CopyToPointer
func (m PositionUpdateFixedPointerBuilder) CopyToPointer(to PositionUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetQuantity(m.Quantity())
	to.SetAveragePrice(m.AveragePrice())
	to.SetPositionIdentifier(m.PositionIdentifier())
	to.SetTradeAccount(m.TradeAccount())
	to.SetNoPositions(m.NoPositions())
	to.SetUnsolicited(m.Unsolicited())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetEntryDateTime(m.EntryDateTime())
	to.SetOpenProfitLoss(m.OpenProfitLoss())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetQuantityLimit(m.QuantityLimit())
	to.SetMaxPotentialPostionQuantity(m.MaxPotentialPostionQuantity())
}
