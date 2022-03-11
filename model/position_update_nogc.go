package model

import (
	"github.com/moontrade/dtc-go/message"
)

// PositionUpdatePointer This is a message from the Server to the Client to report a Trade Position
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
type PositionUpdatePointer struct {
	message.VLSPointer
}

// PositionUpdatePointerBuilder This is a message from the Server to the Client to report a Trade Position
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
type PositionUpdatePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocPositionUpdate() PositionUpdatePointerBuilder {
	a := PositionUpdatePointerBuilder{message.AllocVLSBuilder(120)}
	a.Ptr.SetBytes(0, _PositionUpdateDefault)
	return a
}

func AllocPositionUpdateFrom(b []byte) PositionUpdatePointer {
	return PositionUpdatePointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m PositionUpdatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _PositionUpdateDefault)
}

// ToBuilder
func (m PositionUpdatePointer) ToBuilder() PositionUpdatePointerBuilder {
	return PositionUpdatePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *PositionUpdatePointerBuilder) Finish() PositionUpdatePointer {
	return PositionUpdatePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID The Server sets this to 0, the default, if this is a real-time Trade Position
// update.
//
// Otherwise, when the Server is sending Trade Positions in response to a
// CurrentPositionsRequest message, it must set this to the RequestID given
// in the CurrentPositionsRequest message
func (m PositionUpdatePointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The Server sets this to 0, the default, if this is a real-time Trade Position
// update.
//
// Otherwise, when the Server is sending Trade Positions in response to a
// CurrentPositionsRequest message, it must set this to the RequestID given
// in the CurrentPositionsRequest message
func (m PositionUpdatePointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// TotalNumberMessages This indicates the total number of Position Update messages when a batch
// of messages is being sent. If there is only one Position Update message
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdatePointer) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Ptr, 16, 12)
}

// TotalNumberMessages This indicates the total number of Position Update messages when a batch
// of messages is being sent. If there is only one Position Update message
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdatePointerBuilder) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Ptr, 16, 12)
}

// MessageNumber This indicates the 1-based index of the Position Update message when a
// batch of messages is being sent. If there is only one Position Update
// message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdatePointer) MessageNumber() int32 {
	return message.Int32VLS(m.Ptr, 20, 16)
}

// MessageNumber This indicates the 1-based index of the Position Update message when a
// batch of messages is being sent. If there is only one Position Update
// message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdatePointerBuilder) MessageNumber() int32 {
	return message.Int32VLS(m.Ptr, 20, 16)
}

// Symbol The symbol for the Position.
func (m PositionUpdatePointer) Symbol() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Symbol The symbol for the Position.
func (m PositionUpdatePointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Exchange The optional exchange for the symbol.
func (m PositionUpdatePointer) Exchange() string {
	return message.StringVLS(m.Ptr, 28, 24)
}

// Exchange The optional exchange for the symbol.
func (m PositionUpdatePointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 28, 24)
}

// Quantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdatePointer) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 40, 32)
}

// Quantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdatePointerBuilder) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 40, 32)
}

// AveragePrice The average position price.
func (m PositionUpdatePointer) AveragePrice() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
}

// AveragePrice The average position price.
func (m PositionUpdatePointerBuilder) AveragePrice() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
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
func (m PositionUpdatePointer) PositionIdentifier() string {
	return message.StringVLS(m.Ptr, 52, 48)
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
func (m PositionUpdatePointerBuilder) PositionIdentifier() string {
	return message.StringVLS(m.Ptr, 52, 48)
}

// TradeAccount The trade account the Trade Position is in for the Symbol.
func (m PositionUpdatePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 56, 52)
}

// TradeAccount The trade account the Trade Position is in for the Symbol.
func (m PositionUpdatePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 56, 52)
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
func (m PositionUpdatePointer) NoPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 57, 56)
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
func (m PositionUpdatePointerBuilder) NoPositions() uint8 {
	return message.Uint8VLS(m.Ptr, 57, 56)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdatePointer) Unsolicited() uint8 {
	return message.Uint8VLS(m.Ptr, 58, 57)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdatePointerBuilder) Unsolicited() uint8 {
	return message.Uint8VLS(m.Ptr, 58, 57)
}

// MarginRequirement is the required margin as a currency value for the current
// trade Position Quantity and any working orders for the Trade Account.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointer) MarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 72, 64)
}

// MarginRequirement is the required margin as a currency value for the current
// trade Position Quantity and any working orders for the Trade Account.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointerBuilder) MarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 72, 64)
}

// EntryDateTime is the Date-Time of the initial entry of the Trade Position.
// It is in the UTC time zone.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointer) EntryDateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Ptr, 76, 72))
}

// EntryDateTime is the Date-Time of the initial entry of the Trade Position.
// It is in the UTC time zone.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointerBuilder) EntryDateTime() DateTime4Byte {
	return DateTime4Byte(message.Uint32VLS(m.Ptr, 76, 72))
}

// OpenProfitLoss is the current open Trade Position profit or loss as a
// currency value.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointer) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 88, 80)
}

// OpenProfitLoss is the current open Trade Position profit or loss as a
// currency value.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointerBuilder) OpenProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 88, 80)
}

// HighPriceDuringPosition is the highest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointer) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 96, 88)
}

// HighPriceDuringPosition is the highest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointerBuilder) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 96, 88)
}

// LowPriceDuringPosition is the lowest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointer) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 104, 96)
}

// LowPriceDuringPosition is the lowest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointerBuilder) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 104, 96)
}

// QuantityLimit This is the limit or maximum Trade Position Quantity possible for a Trade
// Position for the Trade Account and Symbol. This applies equally to a long
// or short position.
//
// This is only an informational value provided from the risk management
// system.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointer) QuantityLimit() float64 {
	return message.Float64VLS(m.Ptr, 112, 104)
}

// QuantityLimit This is the limit or maximum Trade Position Quantity possible for a Trade
// Position for the Trade Account and Symbol. This applies equally to a long
// or short position.
//
// This is only an informational value provided from the risk management
// system.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointerBuilder) QuantityLimit() float64 {
	return message.Float64VLS(m.Ptr, 112, 104)
}

// MaxPotentialPostionQuantity
func (m PositionUpdatePointer) MaxPotentialPostionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 120, 112)
}

// MaxPotentialPostionQuantity
func (m PositionUpdatePointerBuilder) MaxPotentialPostionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 120, 112)
}

// SetRequestID The Server sets this to 0, the default, if this is a real-time Trade Position
// update.
//
// Otherwise, when the Server is sending Trade Positions in response to a
// CurrentPositionsRequest message, it must set this to the RequestID given
// in the CurrentPositionsRequest message
func (m PositionUpdatePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetTotalNumberMessages This indicates the total number of Position Update messages when a batch
// of messages is being sent. If there is only one Position Update message
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
// being sent, this will be 1. Use a value of 1 for an unsolicited message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdatePointerBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32VLS(m.Ptr, 16, 12, value)
}

// SetMessageNumber This indicates the 1-based index of the Position Update message when a
// batch of messages is being sent. If there is only one Position Update
// message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m PositionUpdatePointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Ptr, 20, 16, value)
}

// SetSymbol The symbol for the Position.
func (m *PositionUpdatePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetExchange The optional exchange for the symbol.
func (m *PositionUpdatePointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 28, 24, value)
}

// SetQuantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdatePointerBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 40, 32, value)
}

// SetAveragePrice The average position price.
func (m PositionUpdatePointerBuilder) SetAveragePrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 48, 40, value)
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
func (m *PositionUpdatePointerBuilder) SetPositionIdentifier(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 52, 48, value)
}

// SetTradeAccount The trade account the Trade Position is in for the Symbol.
func (m *PositionUpdatePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 56, 52, value)
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
func (m PositionUpdatePointerBuilder) SetNoPositions(value uint8) {
	message.SetUint8VLS(m.Ptr, 57, 56, value)
}

// SetUnsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdatePointerBuilder) SetUnsolicited(value uint8) {
	message.SetUint8VLS(m.Ptr, 58, 57, value)
}

// SetMarginRequirement MarginRequirement is the required margin as a currency value for the current
// trade Position Quantity and any working orders for the Trade Account.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointerBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Ptr, 72, 64, value)
}

// SetEntryDateTime EntryDateTime is the Date-Time of the initial entry of the Trade Position.
// It is in the UTC time zone.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointerBuilder) SetEntryDateTime(value DateTime4Byte) {
	message.SetUint32VLS(m.Ptr, 76, 72, uint32(value))
}

// SetOpenProfitLoss OpenProfitLoss is the current open Trade Position profit or loss as a
// currency value.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointerBuilder) SetOpenProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 88, 80, value)
}

// SetHighPriceDuringPosition HighPriceDuringPosition is the highest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointerBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 96, 88, value)
}

// SetLowPriceDuringPosition LowPriceDuringPosition is the lowest price the symbol traded at during
// the life of the Trade Position. This value can be reset by the Server
// when there is a change in Trade Position Quantity.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointerBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 104, 96, value)
}

// SetQuantityLimit This is the limit or maximum Trade Position Quantity possible for a Trade
// Position for the Trade Account and Symbol. This applies equally to a long
// or short position.
//
// This is only an informational value provided from the risk management
// system.
//
// This is an optional field for the Server to provide.
func (m PositionUpdatePointerBuilder) SetQuantityLimit(value float64) {
	message.SetFloat64VLS(m.Ptr, 112, 104, value)
}

// SetMaxPotentialPostionQuantity
func (m PositionUpdatePointerBuilder) SetMaxPotentialPostionQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 120, 112, value)
}

// Copy
func (m PositionUpdatePointer) Copy(to PositionUpdateBuilder) {
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
func (m PositionUpdatePointerBuilder) Copy(to PositionUpdateBuilder) {
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
func (m PositionUpdatePointer) CopyPointer(to PositionUpdatePointerBuilder) {
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
func (m PositionUpdatePointerBuilder) CopyPointer(to PositionUpdatePointerBuilder) {
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
func (m PositionUpdatePointer) CopyTo(to PositionUpdateFixedBuilder) {
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
func (m PositionUpdatePointerBuilder) CopyTo(to PositionUpdateFixedBuilder) {
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
func (m PositionUpdatePointer) CopyToPointer(to PositionUpdateFixedPointerBuilder) {
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
func (m PositionUpdatePointerBuilder) CopyToPointer(to PositionUpdateFixedPointerBuilder) {
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
