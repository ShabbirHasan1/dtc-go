package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _PositionUpdateFixedDefault = []byte{240, 0, 50, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const PositionUpdateFixedSize = 240

// PositionUpdateFixed This is a message from the Server to the Client to report a Trade Position
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
type PositionUpdateFixed struct {
	message.Fixed
}

// PositionUpdateFixedBuilder This is a message from the Server to the Client to report a Trade Position
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
type PositionUpdateFixedBuilder struct {
	message.Fixed
}

func NewPositionUpdateFixedFrom(b []byte) PositionUpdateFixed {
	return PositionUpdateFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapPositionUpdateFixed(b []byte) PositionUpdateFixed {
	return PositionUpdateFixed{Fixed: message.WrapFixed(b)}
}

func NewPositionUpdateFixed() PositionUpdateFixedBuilder {
	a := PositionUpdateFixedBuilder{message.NewFixed(240)}
	a.Unsafe().SetBytes(0, _PositionUpdateFixedDefault)
	return a
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
func (m PositionUpdateFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _PositionUpdateFixedDefault)
}

// ToBuilder
func (m PositionUpdateFixed) ToBuilder() PositionUpdateFixedBuilder {
	return PositionUpdateFixedBuilder{m.Fixed}
}

// Finish
func (m PositionUpdateFixedBuilder) Finish() PositionUpdateFixed {
	return PositionUpdateFixed{m.Fixed}
}

// RequestID The Server sets this to 0, the default, if this is a real-time Trade Position
// update.
//
// Otherwise, when the Server is sending Trade Positions in response to a
// CurrentPositionsRequest message, it must set this to the RequestID given
// in the CurrentPositionsRequest message
func (m PositionUpdateFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The Server sets this to 0, the default, if this is a real-time Trade Position
// update.
//
// Otherwise, when the Server is sending Trade Positions in response to a
// CurrentPositionsRequest message, it must set this to the RequestID given
// in the CurrentPositionsRequest message
func (m PositionUpdateFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// TotalNumberMessages This indicates the total number of Position Update messages when a batch
// of messages is being sent. If there is only one Position Update message
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateFixed) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// TotalNumberMessages This indicates the total number of Position Update messages when a batch
// of messages is being sent. If there is only one Position Update message
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateFixedBuilder) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// MessageNumber This indicates the 1-based index of the Position Update message when a
// batch of messages is being sent. If there is only one Position Update
// message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateFixed) MessageNumber() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// MessageNumber This indicates the 1-based index of the Position Update message when a
// batch of messages is being sent. If there is only one Position Update
// message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateFixedBuilder) MessageNumber() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// Symbol The symbol for the Position.
func (m PositionUpdateFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 80, 16)
}

// Symbol The symbol for the Position.
func (m PositionUpdateFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 80, 16)
}

// Exchange The optional exchange for the symbol.
func (m PositionUpdateFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 96, 80)
}

// Exchange The optional exchange for the symbol.
func (m PositionUpdateFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 96, 80)
}

// Quantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdateFixed) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 104, 96)
}

// Quantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdateFixedBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 104, 96)
}

// AveragePrice The average position price.
func (m PositionUpdateFixed) AveragePrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 112, 104)
}

// AveragePrice The average position price.
func (m PositionUpdateFixedBuilder) AveragePrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 112, 104)
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
func (m PositionUpdateFixed) PositionIdentifier() string {
	return message.StringFixed(m.Unsafe(), 144, 112)
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
func (m PositionUpdateFixedBuilder) PositionIdentifier() string {
	return message.StringFixed(m.Unsafe(), 144, 112)
}

// TradeAccount The trade account the Trade Position is in for the Symbol.
func (m PositionUpdateFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 176, 144)
}

// TradeAccount The trade account the Trade Position is in for the Symbol.
func (m PositionUpdateFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 176, 144)
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
func (m PositionUpdateFixed) NoPositions() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 177, 176)
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
func (m PositionUpdateFixedBuilder) NoPositions() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 177, 176)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdateFixed) Unsolicited() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 178, 177)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdateFixedBuilder) Unsolicited() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 178, 177)
}

// MarginRequirement is the required margin as a currency value for the current
// trade Position Quantity and any working orders for the Trade Account.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixed) MarginRequirement() float64 {
	return message.Float64Fixed(m.Unsafe(), 192, 184)
}

// MarginRequirement is the required margin as a currency value for the current
// trade Position Quantity and any working orders for the Trade Account.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedBuilder) MarginRequirement() float64 {
	return message.Float64Fixed(m.Unsafe(), 192, 184)
}

// EntryDateTime is the Date-Time of the initial entry of the Trade Position.
// It is in the UTC time zone.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixed) EntryDateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 196, 192))
}

// EntryDateTime is the Date-Time of the initial entry of the Trade Position.
// It is in the UTC time zone.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedBuilder) EntryDateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32Fixed(m.Unsafe(), 196, 192))
}

// OpenProfitLoss is the current open Trade Position profit or loss as a
// currency value.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixed) OpenProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 208, 200)
}

// OpenProfitLoss is the current open Trade Position profit or loss as a
// currency value.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedBuilder) OpenProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 208, 200)
}

// HighPriceDuringPosition is the highest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixed) HighPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Unsafe(), 216, 208)
}

// HighPriceDuringPosition is the highest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedBuilder) HighPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Unsafe(), 216, 208)
}

// LowPriceDuringPosition is the lowest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixed) LowPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Unsafe(), 224, 216)
}

// LowPriceDuringPosition is the lowest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedBuilder) LowPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Unsafe(), 224, 216)
}

// QuantityLimit This is the limit or maximum Trade Position Quantity possible for a Trade
// Position for the Trade Account and Symbol. This applies equally to a long
// or short position.
//
// This is only an informational value provided from the risk management
// system.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixed) QuantityLimit() float64 {
	return message.Float64Fixed(m.Unsafe(), 232, 224)
}

// QuantityLimit This is the limit or maximum Trade Position Quantity possible for a Trade
// Position for the Trade Account and Symbol. This applies equally to a long
// or short position.
//
// This is only an informational value provided from the risk management
// system.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedBuilder) QuantityLimit() float64 {
	return message.Float64Fixed(m.Unsafe(), 232, 224)
}

// MaxPotentialPostionQuantity
func (m PositionUpdateFixed) MaxPotentialPostionQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 240, 232)
}

// MaxPotentialPostionQuantity
func (m PositionUpdateFixedBuilder) MaxPotentialPostionQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 240, 232)
}

// SetRequestID The Server sets this to 0, the default, if this is a real-time Trade Position
// update.
//
// Otherwise, when the Server is sending Trade Positions in response to a
// CurrentPositionsRequest message, it must set this to the RequestID given
// in the CurrentPositionsRequest message
func (m PositionUpdateFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetTotalNumberMessages This indicates the total number of Position Update messages when a batch
// of messages is being sent. If there is only one Position Update message
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateFixedBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetMessageNumber This indicates the 1-based index of the Position Update message when a
// batch of messages is being sent. If there is only one Position Update
// message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateFixedBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 16, 12, value)
}

// SetSymbol The symbol for the Position.
func (m PositionUpdateFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 80, 16, value)
}

// SetExchange The optional exchange for the symbol.
func (m PositionUpdateFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 96, 80, value)
}

// SetQuantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdateFixedBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 104, 96, value)
}

// SetAveragePrice The average position price.
func (m PositionUpdateFixedBuilder) SetAveragePrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 112, 104, value)
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
func (m PositionUpdateFixedBuilder) SetPositionIdentifier(value string) {
	message.SetStringFixed(m.Unsafe(), 144, 112, value)
}

// SetTradeAccount The trade account the Trade Position is in for the Symbol.
func (m PositionUpdateFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 176, 144, value)
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
func (m PositionUpdateFixedBuilder) SetNoPositions(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 177, 176, value)
}

// SetUnsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdateFixedBuilder) SetUnsolicited(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 178, 177, value)
}

// SetMarginRequirement MarginRequirement is the required margin as a currency value for the current
// trade Position Quantity and any working orders for the Trade Account.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 192, 184, value)
}

// SetEntryDateTime EntryDateTime is the Date-Time of the initial entry of the Trade Position.
// It is in the UTC time zone.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedBuilder) SetEntryDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 196, 192, uint32(value))
}

// SetOpenProfitLoss OpenProfitLoss is the current open Trade Position profit or loss as a
// currency value.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedBuilder) SetOpenProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 208, 200, value)
}

// SetHighPriceDuringPosition HighPriceDuringPosition is the highest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 216, 208, value)
}

// SetLowPriceDuringPosition LowPriceDuringPosition is the lowest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 224, 216, value)
}

// SetQuantityLimit This is the limit or maximum Trade Position Quantity possible for a Trade
// Position for the Trade Account and Symbol. This applies equally to a long
// or short position.
//
// This is only an informational value provided from the risk management
// system.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateFixedBuilder) SetQuantityLimit(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 232, 224, value)
}

// SetMaxPotentialPostionQuantity
func (m PositionUpdateFixedBuilder) SetMaxPotentialPostionQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 240, 232, value)
}

// Copy
func (m PositionUpdateFixed) Copy(to PositionUpdateFixedBuilder) {
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
func (m PositionUpdateFixedBuilder) Copy(to PositionUpdateFixedBuilder) {
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
func (m PositionUpdateFixed) CopyPointer(to PositionUpdateFixedPointerBuilder) {
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
func (m PositionUpdateFixedBuilder) CopyPointer(to PositionUpdateFixedPointerBuilder) {
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
func (m PositionUpdateFixed) CopyToPointer(to PositionUpdatePointerBuilder) {
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
func (m PositionUpdateFixedBuilder) CopyToPointer(to PositionUpdatePointerBuilder) {
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
func (m PositionUpdateFixed) CopyTo(to PositionUpdateBuilder) {
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
func (m PositionUpdateFixedBuilder) CopyTo(to PositionUpdateBuilder) {
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
