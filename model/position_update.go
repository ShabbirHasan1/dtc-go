package model

import (
	"github.com/moontrade/dtc-go/message"
)

const PositionUpdateSize = 120

const PositionUpdateFixedSize = 240

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

func AllocPositionUpdate() PositionUpdatePointerBuilder {
	a := PositionUpdatePointerBuilder{message.AllocVLSBuilder(120)}
	a.Ptr.SetBytes(0, _PositionUpdateDefault)
	return a
}

func AllocPositionUpdateFrom(b []byte) PositionUpdatePointer {
	return PositionUpdatePointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m PositionUpdate) ToBuilder() PositionUpdateBuilder {
	return PositionUpdateBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m PositionUpdateFixed) ToBuilder() PositionUpdateFixedBuilder {
	return PositionUpdateFixedBuilder{m.Fixed}
}

// ToBuilder
func (m PositionUpdatePointer) ToBuilder() PositionUpdatePointerBuilder {
	return PositionUpdatePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m PositionUpdateFixedPointer) ToBuilder() PositionUpdateFixedPointerBuilder {
	return PositionUpdateFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m PositionUpdateBuilder) Finish() PositionUpdate {
	return PositionUpdate{m.VLSBuilder.Finish()}
}

// Finish
func (m PositionUpdateFixedBuilder) Finish() PositionUpdateFixed {
	return PositionUpdateFixed{m.Fixed}
}

// Finish
func (m *PositionUpdatePointerBuilder) Finish() PositionUpdatePointer {
	return PositionUpdatePointer{m.VLSPointerBuilder.Finish()}
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
func (m PositionUpdate) Symbol() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Symbol The symbol for the Position.
func (m PositionUpdateBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
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
func (m PositionUpdate) Exchange() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// Exchange The optional exchange for the symbol.
func (m PositionUpdateBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
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
func (m PositionUpdate) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 40, 32)
}

// Quantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdateBuilder) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 40, 32)
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
func (m PositionUpdate) AveragePrice() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// AveragePrice The average position price.
func (m PositionUpdateBuilder) AveragePrice() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
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
func (m PositionUpdate) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 56, 52)
}

// TradeAccount The trade account the Trade Position is in for the Symbol.
func (m PositionUpdateBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 56, 52)
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
func (m PositionUpdate) Unsolicited() uint8 {
	return message.Uint8VLS(m.Unsafe(), 58, 57)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdateBuilder) Unsolicited() uint8 {
	return message.Uint8VLS(m.Unsafe(), 58, 57)
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
func (m PositionUpdate) MaxPotentialPostionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 120, 112)
}

// MaxPotentialPostionQuantity
func (m PositionUpdateBuilder) MaxPotentialPostionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 120, 112)
}

// MaxPotentialPostionQuantity
func (m PositionUpdatePointer) MaxPotentialPostionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 120, 112)
}

// MaxPotentialPostionQuantity
func (m PositionUpdatePointerBuilder) MaxPotentialPostionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 120, 112)
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
func (m PositionUpdateFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 80, 16)
}

// Symbol The symbol for the Position.
func (m PositionUpdateFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 80, 16)
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
func (m PositionUpdateFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 96, 80)
}

// Exchange The optional exchange for the symbol.
func (m PositionUpdateFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 96, 80)
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
func (m PositionUpdateFixed) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 104, 96)
}

// Quantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdateFixedBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 104, 96)
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
func (m PositionUpdateFixed) AveragePrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 112, 104)
}

// AveragePrice The average position price.
func (m PositionUpdateFixedBuilder) AveragePrice() float64 {
	return message.Float64Fixed(m.Unsafe(), 112, 104)
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
func (m PositionUpdateFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 176, 144)
}

// TradeAccount The trade account the Trade Position is in for the Symbol.
func (m PositionUpdateFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 176, 144)
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
func (m PositionUpdateFixed) Unsolicited() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 178, 177)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Position Update message. In
// other words, it is a real-time Trade Position Update message which is
// not an initial response to a CurrentPositionsRequest message.
func (m PositionUpdateFixedBuilder) Unsolicited() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 178, 177)
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
func (m PositionUpdateFixed) MaxPotentialPostionQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 240, 232)
}

// MaxPotentialPostionQuantity
func (m PositionUpdateFixedBuilder) MaxPotentialPostionQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 240, 232)
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
func (m PositionUpdateBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
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
func (m PositionUpdateBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32VLS(m.Unsafe(), 16, 12, value)
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
func (m PositionUpdateBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Unsafe(), 20, 16, value)
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
func (m *PositionUpdateBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetSymbol The symbol for the Position.
func (m *PositionUpdatePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetExchange The optional exchange for the symbol.
func (m *PositionUpdateBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 28, 24, value)
}

// SetExchange The optional exchange for the symbol.
func (m *PositionUpdatePointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 28, 24, value)
}

// SetQuantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdateBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 40, 32, value)
}

// SetQuantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdatePointerBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 40, 32, value)
}

// SetAveragePrice The average position price.
func (m PositionUpdateBuilder) SetAveragePrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 48, 40, value)
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
func (m *PositionUpdateBuilder) SetPositionIdentifier(value string) {
	message.SetStringVLS(&m.VLSBuilder, 52, 48, value)
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
func (m *PositionUpdateBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 56, 52, value)
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
func (m PositionUpdateBuilder) SetNoPositions(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 57, 56, value)
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
func (m PositionUpdateBuilder) SetUnsolicited(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 58, 57, value)
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
func (m PositionUpdateBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 72, 64, value)
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
func (m PositionUpdateBuilder) SetEntryDateTime(value DateTime4Byte) {
	message.SetUint32VLS(m.Unsafe(), 76, 72, uint32(value))
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
func (m PositionUpdateBuilder) SetOpenProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 88, 80, value)
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
func (m PositionUpdateBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 96, 88, value)
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
func (m PositionUpdateBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 104, 96, value)
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
func (m PositionUpdateBuilder) SetQuantityLimit(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 112, 104, value)
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
func (m PositionUpdateBuilder) SetMaxPotentialPostionQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 120, 112, value)
}

// SetMaxPotentialPostionQuantity
func (m PositionUpdatePointerBuilder) SetMaxPotentialPostionQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 120, 112, value)
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
func (m PositionUpdateFixedBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
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
func (m PositionUpdateFixedBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 16, 12, value)
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
func (m PositionUpdateFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 80, 16, value)
}

// SetSymbol The symbol for the Position.
func (m PositionUpdateFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 80, 16, value)
}

// SetExchange The optional exchange for the symbol.
func (m PositionUpdateFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 96, 80, value)
}

// SetExchange The optional exchange for the symbol.
func (m PositionUpdateFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 96, 80, value)
}

// SetQuantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdateFixedBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 104, 96, value)
}

// SetQuantity The quantity of the current Position for the symbol. A positive number
// is for a long Position and a negative number is for a short Position.
func (m PositionUpdateFixedPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 104, 96, value)
}

// SetAveragePrice The average position price.
func (m PositionUpdateFixedBuilder) SetAveragePrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 112, 104, value)
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
func (m PositionUpdateFixedBuilder) SetPositionIdentifier(value string) {
	message.SetStringFixed(m.Unsafe(), 144, 112, value)
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
func (m PositionUpdateFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 176, 144, value)
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
func (m PositionUpdateFixedBuilder) SetNoPositions(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 177, 176, value)
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
func (m PositionUpdateFixedBuilder) SetUnsolicited(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 178, 177, value)
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
func (m PositionUpdateFixedBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 192, 184, value)
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
func (m PositionUpdateFixedBuilder) SetEntryDateTime(value DateTime4Byte) {
	message.SetUint32Fixed(m.Unsafe(), 196, 192, uint32(value))
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
func (m PositionUpdateFixedBuilder) SetOpenProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 208, 200, value)
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
func (m PositionUpdateFixedBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 216, 208, value)
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
func (m PositionUpdateFixedBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 224, 216, value)
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
func (m PositionUpdateFixedBuilder) SetQuantityLimit(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 232, 224, value)
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
func (m PositionUpdateFixedBuilder) SetMaxPotentialPostionQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 240, 232, value)
}

// SetMaxPotentialPostionQuantity
func (m PositionUpdateFixedPointerBuilder) SetMaxPotentialPostionQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 240, 232, value)
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
