package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                        = PositionUpdateSize  (120)
//     Type                        = POSITION_UPDATE  (306)
//     BaseSize                    = PositionUpdateSize  (120)
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
var _PositionUpdateDefault = []byte{120, 0, 50, 1, 120, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const PositionUpdateSize = 120

// PositionUpdate This is a message from the Server to the Client to report a Trade Position
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
type PositionUpdate struct {
	message.VLS
}

// PositionUpdateBuilder This is a message from the Server to the Client to report a Trade Position
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
type PositionUpdateBuilder struct {
	message.VLSBuilder
}

func NewPositionUpdateFrom(b []byte) PositionUpdate {
	return PositionUpdate{VLS: message.NewVLSFrom(b)}
}

func WrapPositionUpdate(b []byte) PositionUpdate {
	return PositionUpdate{VLS: message.WrapVLS(b)}
}

func NewPositionUpdate() PositionUpdateBuilder {
	a := PositionUpdateBuilder{message.NewVLSBuilder(120)}
	a.Unsafe().SetBytes(0, _PositionUpdateDefault)
	return a
}

// Clear
// {
//     Size                        = PositionUpdateSize  (120)
//     Type                        = POSITION_UPDATE  (306)
//     BaseSize                    = PositionUpdateSize  (120)
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
func (m PositionUpdateBuilder) Clear() {
	m.Unsafe().SetBytes(0, _PositionUpdateDefault)
}

// ToBuilder
func (m PositionUpdate) ToBuilder() PositionUpdateBuilder {
	return PositionUpdateBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m PositionUpdateBuilder) Finish() PositionUpdate {
	return PositionUpdate{m.VLSBuilder.Finish()}
}

// RequestID The Server sets this to 0, the default, if this is a real-time Trade Position
// update.
//
// Otherwise, when the Server is sending Trade Positions in response to a
// CurrentPositionsRequest message, it must set this to the RequestID given
// in the CurrentPositionsRequest message
func (m PositionUpdate) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The Server sets this to 0, the default, if this is a real-time Trade Position
// update.
//
// Otherwise, when the Server is sending Trade Positions in response to a
// CurrentPositionsRequest message, it must set this to the RequestID given
// in the CurrentPositionsRequest message
func (m PositionUpdateBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// TotalNumberMessages This indicates the total number of Position Update messages when a batch
// of messages is being sent. If there is only one Position Update message
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdate) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Unsafe(), 16, 12)
}

// TotalNumberMessages This indicates the total number of Position Update messages when a batch
// of messages is being sent. If there is only one Position Update message
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateBuilder) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Unsafe(), 16, 12)
}

// MessageNumber This indicates the 1-based index of the Position Update message when a
// batch of messages is being sent. If there is only one Position Update
// message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdate) MessageNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 20, 16)
}

// MessageNumber This indicates the 1-based index of the Position Update message when a
// batch of messages is being sent. If there is only one Position Update
// message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateBuilder) MessageNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 20, 16)
}

// Symbol The symbol for the Position.
func (m PositionUpdate) Symbol() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Symbol The symbol for the Position.
func (m PositionUpdateBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Exchange The optional exchange for the symbol.
func (m PositionUpdate) Exchange() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// Exchange The optional exchange for the symbol.
func (m PositionUpdateBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// Quantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdate) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 40, 32)
}

// Quantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdateBuilder) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 40, 32)
}

// AveragePrice The average position price.
func (m PositionUpdate) AveragePrice() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// AveragePrice The average position price.
func (m PositionUpdateBuilder) AveragePrice() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
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
func (m PositionUpdate) PositionIdentifier() string {
	return message.StringVLS(m.Unsafe(), 52, 48)
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
func (m PositionUpdateBuilder) PositionIdentifier() string {
	return message.StringVLS(m.Unsafe(), 52, 48)
}

// TradeAccount The trade account the Trade Position is in for the Symbol.
func (m PositionUpdate) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 56, 52)
}

// TradeAccount The trade account the Trade Position is in for the Symbol.
func (m PositionUpdateBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 56, 52)
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
func (m PositionUpdate) NoPositions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 57, 56)
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
func (m PositionUpdateBuilder) NoPositions() uint8 {
	return message.Uint8VLS(m.Unsafe(), 57, 56)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdate) Unsolicited() uint8 {
	return message.Uint8VLS(m.Unsafe(), 58, 57)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdateBuilder) Unsolicited() uint8 {
	return message.Uint8VLS(m.Unsafe(), 58, 57)
}

// MarginRequirement is the required margin as a currency value for the current
// trade Position Quantity and any working orders for the Trade Account.
//
// This is an optional field for the Server to provide.
func (m PositionUpdate) MarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 72, 64)
}

// MarginRequirement is the required margin as a currency value for the current
// trade Position Quantity and any working orders for the Trade Account.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateBuilder) MarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 72, 64)
}

// EntryDateTime is the Date-Time of the initial entry of the Trade Position.
// It is in the UTC time zone.
//
// This is an optional field for the Server to provide.
func (m PositionUpdate) EntryDateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Unsafe(), 76, 72))
}

// EntryDateTime is the Date-Time of the initial entry of the Trade Position.
// It is in the UTC time zone.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateBuilder) EntryDateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Unsafe(), 76, 72))
}

// OpenProfitLoss is the current open Trade Position profit or loss as a
// currency value.
//
// This is an optional field for the Server to provide.
func (m PositionUpdate) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 88, 80)
}

// OpenProfitLoss is the current open Trade Position profit or loss as a
// currency value.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateBuilder) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 88, 80)
}

// HighPriceDuringPosition is the highest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdate) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 96, 88)
}

// HighPriceDuringPosition is the highest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateBuilder) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 96, 88)
}

// LowPriceDuringPosition is the lowest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdate) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 104, 96)
}

// LowPriceDuringPosition is the lowest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateBuilder) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 104, 96)
}

// QuantityLimit This is the limit or maximum Trade Position Quantity possible for a Trade
// Position for the Trade Account and Symbol. This applies equally to a long
// or short position.
//
// This is only an informational value provided from the risk management
// system.
//
// This is an optional field for the Server to provide.
func (m PositionUpdate) QuantityLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 112, 104)
}

// QuantityLimit This is the limit or maximum Trade Position Quantity possible for a Trade
// Position for the Trade Account and Symbol. This applies equally to a long
// or short position.
//
// This is only an informational value provided from the risk management
// system.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateBuilder) QuantityLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 112, 104)
}

// MaxPotentialPostionQuantity
func (m PositionUpdate) MaxPotentialPostionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 120, 112)
}

// MaxPotentialPostionQuantity
func (m PositionUpdateBuilder) MaxPotentialPostionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 120, 112)
}

// SetRequestID The Server sets this to 0, the default, if this is a real-time Trade Position
// update.
//
// Otherwise, when the Server is sending Trade Positions in response to a
// CurrentPositionsRequest message, it must set this to the RequestID given
// in the CurrentPositionsRequest message
func (m PositionUpdateBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetTotalNumberMessages This indicates the total number of Position Update messages when a batch
// of messages is being sent. If there is only one Position Update message
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32VLS(m.Unsafe(), 16, 12, value)
}

// SetMessageNumber This indicates the 1-based index of the Position Update message when a
// batch of messages is being sent. If there is only one Position Update
// message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdateBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Unsafe(), 20, 16, value)
}

// SetSymbol The symbol for the Position.
func (m *PositionUpdateBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetExchange The optional exchange for the symbol.
func (m *PositionUpdateBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 28, 24, value)
}

// SetQuantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdateBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 40, 32, value)
}

// SetAveragePrice The average position price.
func (m PositionUpdateBuilder) SetAveragePrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 48, 40, value)
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
func (m *PositionUpdateBuilder) SetPositionIdentifier(value string) {
	message.SetStringVLS(&m.VLSBuilder, 52, 48, value)
}

// SetTradeAccount The trade account the Trade Position is in for the Symbol.
func (m *PositionUpdateBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 56, 52, value)
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
func (m PositionUpdateBuilder) SetNoPositions(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 57, 56, value)
}

// SetUnsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdateBuilder) SetUnsolicited(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 58, 57, value)
}

// SetMarginRequirement MarginRequirement is the required margin as a currency value for the current
// trade Position Quantity and any working orders for the Trade Account.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 72, 64, value)
}

// SetEntryDateTime EntryDateTime is the Date-Time of the initial entry of the Trade Position.
// It is in the UTC time zone.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateBuilder) SetEntryDateTime(value DateTime4Byte) {
	message.SetUint32VLS(m.Unsafe(), 76, 72, uint32(value))
}

// SetOpenProfitLoss OpenProfitLoss is the current open Trade Position profit or loss as a
// currency value.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateBuilder) SetOpenProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 88, 80, value)
}

// SetHighPriceDuringPosition HighPriceDuringPosition is the highest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 96, 88, value)
}

// SetLowPriceDuringPosition LowPriceDuringPosition is the lowest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 104, 96, value)
}

// SetQuantityLimit This is the limit or maximum Trade Position Quantity possible for a Trade
// Position for the Trade Account and Symbol. This applies equally to a long
// or short position.
//
// This is only an informational value provided from the risk management
// system.
//
// This is an optional field for the Server to provide.
func (m PositionUpdateBuilder) SetQuantityLimit(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 112, 104, value)
}

// SetMaxPotentialPostionQuantity
func (m PositionUpdateBuilder) SetMaxPotentialPostionQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 120, 112, value)
}

// Copy
func (m PositionUpdate) Copy(to PositionUpdateBuilder) {
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
func (m PositionUpdateBuilder) Copy(to PositionUpdateBuilder) {
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
func (m PositionUpdate) CopyPointer(to PositionUpdatePointerBuilder) {
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
func (m PositionUpdateBuilder) CopyPointer(to PositionUpdatePointerBuilder) {
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
func (m PositionUpdate) CopyToPointer(to PositionUpdateFixedPointerBuilder) {
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
func (m PositionUpdateBuilder) CopyToPointer(to PositionUpdateFixedPointerBuilder) {
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
func (m PositionUpdate) CopyTo(to PositionUpdateFixedBuilder) {
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
func (m PositionUpdateBuilder) CopyTo(to PositionUpdateFixedBuilder) {
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
