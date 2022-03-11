package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalOrderFillResponseSize = 112

const HistoricalOrderFillResponseFixedSize = 384

// {
//     Size                    = HistoricalOrderFillResponseSize  (112)
//     Type                    = HISTORICAL_ORDER_FILL_RESPONSE  (304)
//     BaseSize                = HistoricalOrderFillResponseSize  (112)
//     RequestID               = 0
//     TotalNumberMessages     = 0
//     MessageNumber           = 0
//     Symbol                  = ""
//     Exchange                = ""
//     ServerOrderID           = ""
//     BuySell                 = BUY_SELL_UNSET  (0)
//     Price                   = 0
//     DateTime                = 0
//     Quantity                = 0
//     UniqueExecutionID       = ""
//     TradeAccount            = ""
//     OpenClose               = TRADE_UNSET  (0)
//     NoOrderFills            = 0
//     InfoText                = ""
//     HighPriceDuringPosition = 0
//     LowPriceDuringPosition  = 0
//     PositionQuantity        = math.MaxFloat64
// }
var _HistoricalOrderFillResponseDefault = []byte{112, 0, 48, 1, 112, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127}

// {
//     Size                    = HistoricalOrderFillResponseFixedSize  (384)
//     Type                    = HISTORICAL_ORDER_FILL_RESPONSE  (304)
//     RequestID               = 0
//     TotalNumberMessages     = 0
//     MessageNumber           = 0
//     Symbol                  = ""
//     Exchange                = ""
//     ServerOrderID           = ""
//     BuySell                 = BUY_SELL_UNSET  (0)
//     Price                   = 0
//     DateTime                = 0
//     Quantity                = 0
//     UniqueExecutionID       = ""
//     TradeAccount            = ""
//     OpenClose               = TRADE_UNSET  (0)
//     NoOrderFills            = 0
//     InfoText                = ""
//     HighPriceDuringPosition = 0
//     LowPriceDuringPosition  = 0
//     PositionQuantity        = math.MaxFloat64
// }
var _HistoricalOrderFillResponseFixedDefault = []byte{128, 1, 48, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127}

// HistoricalOrderFillResponse This is a message from the Server to the Client providing an individual
// historical order fill in response to a HistoricalOrderFillsRequest message.
// historical order fill in response to a HistoricalOrderFillsRequest message.
//
// The Server is expected to send this message to the Client in response
// to a HistoricalOrderFillsRequest message even when there are no order
// fills to return. If there are no order fills to return, it needs to set
// the NoOrderFills field to 1.
type HistoricalOrderFillResponse struct {
	message.VLS
}

// HistoricalOrderFillResponseBuilder This is a message from the Server to the Client providing an individual
// historical order fill in response to a HistoricalOrderFillsRequest message.
// historical order fill in response to a HistoricalOrderFillsRequest message.
//
// The Server is expected to send this message to the Client in response
// to a HistoricalOrderFillsRequest message even when there are no order
// fills to return. If there are no order fills to return, it needs to set
// the NoOrderFills field to 1.
type HistoricalOrderFillResponseBuilder struct {
	message.VLSBuilder
}

// HistoricalOrderFillResponseFixed This is a message from the Server to the Client providing an individual
// historical order fill in response to a HistoricalOrderFillsRequest message.
// historical order fill in response to a HistoricalOrderFillsRequest message.
//
// The Server is expected to send this message to the Client in response
// to a HistoricalOrderFillsRequest message even when there are no order
// fills to return. If there are no order fills to return, it needs to set
// the NoOrderFills field to 1.
type HistoricalOrderFillResponseFixed struct {
	message.Fixed
}

// HistoricalOrderFillResponseFixedBuilder This is a message from the Server to the Client providing an individual
// historical order fill in response to a HistoricalOrderFillsRequest message.
// historical order fill in response to a HistoricalOrderFillsRequest message.
//
// The Server is expected to send this message to the Client in response
// to a HistoricalOrderFillsRequest message even when there are no order
// fills to return. If there are no order fills to return, it needs to set
// the NoOrderFills field to 1.
type HistoricalOrderFillResponseFixedBuilder struct {
	message.Fixed
}

// HistoricalOrderFillResponsePointer This is a message from the Server to the Client providing an individual
// historical order fill in response to a HistoricalOrderFillsRequest message.
// historical order fill in response to a HistoricalOrderFillsRequest message.
//
// The Server is expected to send this message to the Client in response
// to a HistoricalOrderFillsRequest message even when there are no order
// fills to return. If there are no order fills to return, it needs to set
// the NoOrderFills field to 1.
type HistoricalOrderFillResponsePointer struct {
	message.VLSPointer
}

// HistoricalOrderFillResponsePointerBuilder This is a message from the Server to the Client providing an individual
// historical order fill in response to a HistoricalOrderFillsRequest message.
// historical order fill in response to a HistoricalOrderFillsRequest message.
//
// The Server is expected to send this message to the Client in response
// to a HistoricalOrderFillsRequest message even when there are no order
// fills to return. If there are no order fills to return, it needs to set
// the NoOrderFills field to 1.
type HistoricalOrderFillResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

// HistoricalOrderFillResponseFixedPointer This is a message from the Server to the Client providing an individual
// historical order fill in response to a HistoricalOrderFillsRequest message.
// historical order fill in response to a HistoricalOrderFillsRequest message.
//
// The Server is expected to send this message to the Client in response
// to a HistoricalOrderFillsRequest message even when there are no order
// fills to return. If there are no order fills to return, it needs to set
// the NoOrderFills field to 1.
type HistoricalOrderFillResponseFixedPointer struct {
	message.FixedPointer
}

// HistoricalOrderFillResponseFixedPointerBuilder This is a message from the Server to the Client providing an individual
// historical order fill in response to a HistoricalOrderFillsRequest message.
// historical order fill in response to a HistoricalOrderFillsRequest message.
//
// The Server is expected to send this message to the Client in response
// to a HistoricalOrderFillsRequest message even when there are no order
// fills to return. If there are no order fills to return, it needs to set
// the NoOrderFills field to 1.
type HistoricalOrderFillResponseFixedPointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalOrderFillResponseFrom(b []byte) HistoricalOrderFillResponse {
	return HistoricalOrderFillResponse{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalOrderFillResponse(b []byte) HistoricalOrderFillResponse {
	return HistoricalOrderFillResponse{VLS: message.WrapVLS(b)}
}

func NewHistoricalOrderFillResponse() HistoricalOrderFillResponseBuilder {
	a := HistoricalOrderFillResponseBuilder{message.NewVLSBuilder(112)}
	a.Unsafe().SetBytes(0, _HistoricalOrderFillResponseDefault)
	return a
}

func NewHistoricalOrderFillResponseFixedFrom(b []byte) HistoricalOrderFillResponseFixed {
	return HistoricalOrderFillResponseFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalOrderFillResponseFixed(b []byte) HistoricalOrderFillResponseFixed {
	return HistoricalOrderFillResponseFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalOrderFillResponseFixed() HistoricalOrderFillResponseFixedBuilder {
	a := HistoricalOrderFillResponseFixedBuilder{message.NewFixed(384)}
	a.Unsafe().SetBytes(0, _HistoricalOrderFillResponseFixedDefault)
	return a
}

func AllocHistoricalOrderFillResponse() HistoricalOrderFillResponsePointerBuilder {
	a := HistoricalOrderFillResponsePointerBuilder{message.AllocVLSBuilder(112)}
	a.Ptr.SetBytes(0, _HistoricalOrderFillResponseDefault)
	return a
}

func AllocHistoricalOrderFillResponseFrom(b []byte) HistoricalOrderFillResponsePointer {
	return HistoricalOrderFillResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocHistoricalOrderFillResponseFixed() HistoricalOrderFillResponseFixedPointerBuilder {
	a := HistoricalOrderFillResponseFixedPointerBuilder{message.AllocFixed(384)}
	a.Ptr.SetBytes(0, _HistoricalOrderFillResponseFixedDefault)
	return a
}

func AllocHistoricalOrderFillResponseFixedFrom(b []byte) HistoricalOrderFillResponseFixedPointer {
	return HistoricalOrderFillResponseFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                    = HistoricalOrderFillResponseSize  (112)
//     Type                    = HISTORICAL_ORDER_FILL_RESPONSE  (304)
//     BaseSize                = HistoricalOrderFillResponseSize  (112)
//     RequestID               = 0
//     TotalNumberMessages     = 0
//     MessageNumber           = 0
//     Symbol                  = ""
//     Exchange                = ""
//     ServerOrderID           = ""
//     BuySell                 = BUY_SELL_UNSET  (0)
//     Price                   = 0
//     DateTime                = 0
//     Quantity                = 0
//     UniqueExecutionID       = ""
//     TradeAccount            = ""
//     OpenClose               = TRADE_UNSET  (0)
//     NoOrderFills            = 0
//     InfoText                = ""
//     HighPriceDuringPosition = 0
//     LowPriceDuringPosition  = 0
//     PositionQuantity        = math.MaxFloat64
// }
func (m HistoricalOrderFillResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalOrderFillResponseDefault)
}

// Clear
// {
//     Size                    = HistoricalOrderFillResponseFixedSize  (384)
//     Type                    = HISTORICAL_ORDER_FILL_RESPONSE  (304)
//     RequestID               = 0
//     TotalNumberMessages     = 0
//     MessageNumber           = 0
//     Symbol                  = ""
//     Exchange                = ""
//     ServerOrderID           = ""
//     BuySell                 = BUY_SELL_UNSET  (0)
//     Price                   = 0
//     DateTime                = 0
//     Quantity                = 0
//     UniqueExecutionID       = ""
//     TradeAccount            = ""
//     OpenClose               = TRADE_UNSET  (0)
//     NoOrderFills            = 0
//     InfoText                = ""
//     HighPriceDuringPosition = 0
//     LowPriceDuringPosition  = 0
//     PositionQuantity        = math.MaxFloat64
// }
func (m HistoricalOrderFillResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalOrderFillResponseFixedDefault)
}

// Clear
// {
//     Size                    = HistoricalOrderFillResponseSize  (112)
//     Type                    = HISTORICAL_ORDER_FILL_RESPONSE  (304)
//     BaseSize                = HistoricalOrderFillResponseSize  (112)
//     RequestID               = 0
//     TotalNumberMessages     = 0
//     MessageNumber           = 0
//     Symbol                  = ""
//     Exchange                = ""
//     ServerOrderID           = ""
//     BuySell                 = BUY_SELL_UNSET  (0)
//     Price                   = 0
//     DateTime                = 0
//     Quantity                = 0
//     UniqueExecutionID       = ""
//     TradeAccount            = ""
//     OpenClose               = TRADE_UNSET  (0)
//     NoOrderFills            = 0
//     InfoText                = ""
//     HighPriceDuringPosition = 0
//     LowPriceDuringPosition  = 0
//     PositionQuantity        = math.MaxFloat64
// }
func (m HistoricalOrderFillResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalOrderFillResponseDefault)
}

// Clear
// {
//     Size                    = HistoricalOrderFillResponseFixedSize  (384)
//     Type                    = HISTORICAL_ORDER_FILL_RESPONSE  (304)
//     RequestID               = 0
//     TotalNumberMessages     = 0
//     MessageNumber           = 0
//     Symbol                  = ""
//     Exchange                = ""
//     ServerOrderID           = ""
//     BuySell                 = BUY_SELL_UNSET  (0)
//     Price                   = 0
//     DateTime                = 0
//     Quantity                = 0
//     UniqueExecutionID       = ""
//     TradeAccount            = ""
//     OpenClose               = TRADE_UNSET  (0)
//     NoOrderFills            = 0
//     InfoText                = ""
//     HighPriceDuringPosition = 0
//     LowPriceDuringPosition  = 0
//     PositionQuantity        = math.MaxFloat64
// }
func (m HistoricalOrderFillResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalOrderFillResponseFixedDefault)
}

// ToBuilder
func (m HistoricalOrderFillResponse) ToBuilder() HistoricalOrderFillResponseBuilder {
	return HistoricalOrderFillResponseBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m HistoricalOrderFillResponseFixed) ToBuilder() HistoricalOrderFillResponseFixedBuilder {
	return HistoricalOrderFillResponseFixedBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalOrderFillResponsePointer) ToBuilder() HistoricalOrderFillResponsePointerBuilder {
	return HistoricalOrderFillResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m HistoricalOrderFillResponseFixedPointer) ToBuilder() HistoricalOrderFillResponseFixedPointerBuilder {
	return HistoricalOrderFillResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalOrderFillResponseBuilder) Finish() HistoricalOrderFillResponse {
	return HistoricalOrderFillResponse{m.VLSBuilder.Finish()}
}

// Finish
func (m HistoricalOrderFillResponseFixedBuilder) Finish() HistoricalOrderFillResponseFixed {
	return HistoricalOrderFillResponseFixed{m.Fixed}
}

// Finish
func (m *HistoricalOrderFillResponsePointerBuilder) Finish() HistoricalOrderFillResponsePointer {
	return HistoricalOrderFillResponsePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *HistoricalOrderFillResponseFixedPointerBuilder) Finish() HistoricalOrderFillResponseFixedPointer {
	return HistoricalOrderFillResponseFixedPointer{m.FixedPointer}
}

// RequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponse) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponseBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponsePointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponsePointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// TotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponse) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Unsafe(), 16, 12)
}

// TotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponseBuilder) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Unsafe(), 16, 12)
}

// TotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponsePointer) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Ptr, 16, 12)
}

// TotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponsePointerBuilder) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Ptr, 16, 12)
}

// MessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponse) MessageNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 20, 16)
}

// MessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponseBuilder) MessageNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 20, 16)
}

// MessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponsePointer) MessageNumber() int32 {
	return message.Int32VLS(m.Ptr, 20, 16)
}

// MessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponsePointerBuilder) MessageNumber() int32 {
	return message.Int32VLS(m.Ptr, 20, 16)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponse) Symbol() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponsePointer) Symbol() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponsePointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponse) Exchange() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponseBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponsePointer) Exchange() string {
	return message.StringVLS(m.Ptr, 28, 24)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponsePointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 28, 24)
}

// ServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponse) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 32, 28)
}

// ServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponseBuilder) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 32, 28)
}

// ServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponsePointer) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 32, 28)
}

// ServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponsePointerBuilder) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 32, 28)
}

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponse) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 36, 32))
}

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 36, 32))
}

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponsePointer) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 36, 32))
}

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponsePointerBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 36, 32))
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponse) Price() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponseBuilder) Price() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponsePointer) Price() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponsePointerBuilder) Price() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponse) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 56, 48))
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseBuilder) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 56, 48))
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponsePointer) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 56, 48))
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponsePointerBuilder) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 56, 48))
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponse) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 64, 56)
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseBuilder) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 64, 56)
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponsePointer) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 64, 56)
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponsePointerBuilder) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 64, 56)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponse) UniqueExecutionID() string {
	return message.StringVLS(m.Unsafe(), 68, 64)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponseBuilder) UniqueExecutionID() string {
	return message.StringVLS(m.Unsafe(), 68, 64)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponsePointer) UniqueExecutionID() string {
	return message.StringVLS(m.Ptr, 68, 64)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponsePointerBuilder) UniqueExecutionID() string {
	return message.StringVLS(m.Ptr, 68, 64)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 72, 68)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 72, 68)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponsePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 72, 68)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponsePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 72, 68)
}

// OpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponse) OpenClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 76, 72))
}

// OpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponseBuilder) OpenClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 76, 72))
}

// OpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponsePointer) OpenClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Ptr, 76, 72))
}

// OpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponsePointerBuilder) OpenClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Ptr, 76, 72))
}

// NoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponse) NoOrderFills() uint8 {
	return message.Uint8VLS(m.Unsafe(), 77, 76)
}

// NoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponseBuilder) NoOrderFills() uint8 {
	return message.Uint8VLS(m.Unsafe(), 77, 76)
}

// NoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponsePointer) NoOrderFills() uint8 {
	return message.Uint8VLS(m.Ptr, 77, 76)
}

// NoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponsePointerBuilder) NoOrderFills() uint8 {
	return message.Uint8VLS(m.Ptr, 77, 76)
}

// InfoText
func (m HistoricalOrderFillResponse) InfoText() string {
	return message.StringVLS(m.Unsafe(), 82, 78)
}

// InfoText
func (m HistoricalOrderFillResponseBuilder) InfoText() string {
	return message.StringVLS(m.Unsafe(), 82, 78)
}

// InfoText
func (m HistoricalOrderFillResponsePointer) InfoText() string {
	return message.StringVLS(m.Ptr, 82, 78)
}

// InfoText
func (m HistoricalOrderFillResponsePointerBuilder) InfoText() string {
	return message.StringVLS(m.Ptr, 82, 78)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponse) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 96, 88)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponseBuilder) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 96, 88)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponsePointer) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 96, 88)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponsePointerBuilder) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 96, 88)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponse) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 104, 96)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponseBuilder) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 104, 96)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponsePointer) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 104, 96)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponsePointerBuilder) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 104, 96)
}

// PositionQuantity
func (m HistoricalOrderFillResponse) PositionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 112, 104)
}

// PositionQuantity
func (m HistoricalOrderFillResponseBuilder) PositionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 112, 104)
}

// PositionQuantity
func (m HistoricalOrderFillResponsePointer) PositionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 112, 104)
}

// PositionQuantity
func (m HistoricalOrderFillResponsePointerBuilder) PositionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 112, 104)
}

// RequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponseFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponseFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponseFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponseFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// TotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponseFixed) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// TotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponseFixedBuilder) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Unsafe(), 12, 8)
}

// TotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponseFixedPointer) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// TotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponseFixedPointerBuilder) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Ptr, 12, 8)
}

// MessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponseFixed) MessageNumber() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// MessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponseFixedBuilder) MessageNumber() int32 {
	return message.Int32Fixed(m.Unsafe(), 16, 12)
}

// MessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponseFixedPointer) MessageNumber() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// MessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponseFixedPointerBuilder) MessageNumber() int32 {
	return message.Int32Fixed(m.Ptr, 16, 12)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 80, 16)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 80, 16)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 80, 16)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 80, 16)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponseFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 96, 80)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponseFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 96, 80)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponseFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 96, 80)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponseFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 96, 80)
}

// ServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponseFixed) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 128, 96)
}

// ServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponseFixedBuilder) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 128, 96)
}

// ServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponseFixedPointer) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 128, 96)
}

// ServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponseFixedPointerBuilder) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 128, 96)
}

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseFixed) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 132, 128))
}

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseFixedBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 132, 128))
}

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseFixedPointer) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 132, 128))
}

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseFixedPointerBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 132, 128))
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponseFixed) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 144, 136)
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 144, 136)
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponseFixedPointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 144, 136)
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 144, 136)
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseFixed) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 152, 144))
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 152, 144))
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseFixedPointer) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 152, 144))
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 152, 144))
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseFixed) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 160, 152)
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 160, 152)
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseFixedPointer) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 160, 152)
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 160, 152)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponseFixed) UniqueExecutionID() string {
	return message.StringFixed(m.Unsafe(), 224, 160)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) UniqueExecutionID() string {
	return message.StringFixed(m.Unsafe(), 224, 160)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponseFixedPointer) UniqueExecutionID() string {
	return message.StringFixed(m.Ptr, 224, 160)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) UniqueExecutionID() string {
	return message.StringFixed(m.Ptr, 224, 160)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponseFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 256, 224)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponseFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 256, 224)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponseFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 256, 224)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponseFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 256, 224)
}

// OpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponseFixed) OpenClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Unsafe(), 260, 256))
}

// OpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponseFixedBuilder) OpenClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Unsafe(), 260, 256))
}

// OpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponseFixedPointer) OpenClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 260, 256))
}

// OpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponseFixedPointerBuilder) OpenClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 260, 256))
}

// NoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponseFixed) NoOrderFills() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 261, 260)
}

// NoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponseFixedBuilder) NoOrderFills() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 261, 260)
}

// NoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponseFixedPointer) NoOrderFills() uint8 {
	return message.Uint8Fixed(m.Ptr, 261, 260)
}

// NoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponseFixedPointerBuilder) NoOrderFills() uint8 {
	return message.Uint8Fixed(m.Ptr, 261, 260)
}

// InfoText
func (m HistoricalOrderFillResponseFixed) InfoText() string {
	return message.StringFixed(m.Unsafe(), 357, 261)
}

// InfoText
func (m HistoricalOrderFillResponseFixedBuilder) InfoText() string {
	return message.StringFixed(m.Unsafe(), 357, 261)
}

// InfoText
func (m HistoricalOrderFillResponseFixedPointer) InfoText() string {
	return message.StringFixed(m.Ptr, 357, 261)
}

// InfoText
func (m HistoricalOrderFillResponseFixedPointerBuilder) InfoText() string {
	return message.StringFixed(m.Ptr, 357, 261)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponseFixed) HighPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Unsafe(), 368, 360)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponseFixedBuilder) HighPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Unsafe(), 368, 360)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponseFixedPointer) HighPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Ptr, 368, 360)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponseFixedPointerBuilder) HighPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Ptr, 368, 360)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponseFixed) LowPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Unsafe(), 376, 368)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponseFixedBuilder) LowPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Unsafe(), 376, 368)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponseFixedPointer) LowPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Ptr, 376, 368)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponseFixedPointerBuilder) LowPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Ptr, 376, 368)
}

// PositionQuantity
func (m HistoricalOrderFillResponseFixed) PositionQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 384, 376)
}

// PositionQuantity
func (m HistoricalOrderFillResponseFixedBuilder) PositionQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 384, 376)
}

// PositionQuantity
func (m HistoricalOrderFillResponseFixedPointer) PositionQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 384, 376)
}

// PositionQuantity
func (m HistoricalOrderFillResponseFixedPointerBuilder) PositionQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 384, 376)
}

// SetRequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponseBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetTotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponseBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32VLS(m.Unsafe(), 16, 12, value)
}

// SetTotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponsePointerBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32VLS(m.Ptr, 16, 12, value)
}

// SetMessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponseBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Unsafe(), 20, 16, value)
}

// SetMessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponsePointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Ptr, 20, 16, value)
}

// SetSymbol The symbol the order fill is for.
func (m *HistoricalOrderFillResponseBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetSymbol The symbol the order fill is for.
func (m *HistoricalOrderFillResponsePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetExchange The optional exchange for the symbol.
func (m *HistoricalOrderFillResponseBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 28, 24, value)
}

// SetExchange The optional exchange for the symbol.
func (m *HistoricalOrderFillResponsePointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 28, 24, value)
}

// SetServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m *HistoricalOrderFillResponseBuilder) SetServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 32, 28, value)
}

// SetServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m *HistoricalOrderFillResponsePointerBuilder) SetServerOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 32, 28, value)
}

// SetBuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 36, 32, int32(value))
}

// SetBuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponsePointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 36, 32, int32(value))
}

// SetPrice This is the price of the order fill.
func (m HistoricalOrderFillResponseBuilder) SetPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 48, 40, value)
}

// SetPrice This is the price of the order fill.
func (m HistoricalOrderFillResponsePointerBuilder) SetPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 48, 40, value)
}

// SetDateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseBuilder) SetDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 56, 48, int64(value))
}

// SetDateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponsePointerBuilder) SetDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 56, 48, int64(value))
}

// SetQuantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 64, 56, value)
}

// SetQuantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponsePointerBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 64, 56, value)
}

// SetUniqueExecutionID This is the unique execution identifier for the order fill.
func (m *HistoricalOrderFillResponseBuilder) SetUniqueExecutionID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 68, 64, value)
}

// SetUniqueExecutionID This is the unique execution identifier for the order fill.
func (m *HistoricalOrderFillResponsePointerBuilder) SetUniqueExecutionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 68, 64, value)
}

// SetTradeAccount This is the trade account that the order fill is associated with.
func (m *HistoricalOrderFillResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 72, 68, value)
}

// SetTradeAccount This is the trade account that the order fill is associated with.
func (m *HistoricalOrderFillResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 72, 68, value)
}

// SetOpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponseBuilder) SetOpenClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Unsafe(), 76, 72, int32(value))
}

// SetOpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponsePointerBuilder) SetOpenClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Ptr, 76, 72, int32(value))
}

// SetNoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponseBuilder) SetNoOrderFills(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 77, 76, value)
}

// SetNoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponsePointerBuilder) SetNoOrderFills(value uint8) {
	message.SetUint8VLS(m.Ptr, 77, 76, value)
}

// SetInfoText
func (m *HistoricalOrderFillResponseBuilder) SetInfoText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 82, 78, value)
}

// SetInfoText
func (m *HistoricalOrderFillResponsePointerBuilder) SetInfoText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 82, 78, value)
}

// SetHighPriceDuringPosition
func (m HistoricalOrderFillResponseBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 96, 88, value)
}

// SetHighPriceDuringPosition
func (m HistoricalOrderFillResponsePointerBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 96, 88, value)
}

// SetLowPriceDuringPosition
func (m HistoricalOrderFillResponseBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 104, 96, value)
}

// SetLowPriceDuringPosition
func (m HistoricalOrderFillResponsePointerBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 104, 96, value)
}

// SetPositionQuantity
func (m HistoricalOrderFillResponseBuilder) SetPositionQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 112, 104, value)
}

// SetPositionQuantity
func (m HistoricalOrderFillResponsePointerBuilder) SetPositionQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 112, 104, value)
}

// SetRequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponseFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetTotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponseFixedBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetTotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32Fixed(m.Ptr, 12, 8, value)
}

// SetMessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponseFixedBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 16, 12, value)
}

// SetMessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Ptr, 16, 12, value)
}

// SetSymbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 80, 16, value)
}

// SetSymbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 80, 16, value)
}

// SetExchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponseFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 96, 80, value)
}

// SetExchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 96, 80, value)
}

// SetServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponseFixedBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 128, 96, value)
}

// SetServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 128, 96, value)
}

// SetBuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseFixedBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 132, 128, int32(value))
}

// SetBuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 132, 128, int32(value))
}

// SetPrice This is the price of the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 144, 136, value)
}

// SetPrice This is the price of the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 144, 136, value)
}

// SetDateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) SetDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 152, 144, int64(value))
}

// SetDateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 152, 144, int64(value))
}

// SetQuantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 160, 152, value)
}

// SetQuantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 160, 152, value)
}

// SetUniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) SetUniqueExecutionID(value string) {
	message.SetStringFixed(m.Unsafe(), 224, 160, value)
}

// SetUniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetUniqueExecutionID(value string) {
	message.SetStringFixed(m.Ptr, 224, 160, value)
}

// SetTradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponseFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 256, 224, value)
}

// SetTradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 256, 224, value)
}

// SetOpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponseFixedBuilder) SetOpenClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 260, 256, int32(value))
}

// SetOpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetOpenClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Ptr, 260, 256, int32(value))
}

// SetNoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponseFixedBuilder) SetNoOrderFills(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 261, 260, value)
}

// SetNoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetNoOrderFills(value uint8) {
	message.SetUint8Fixed(m.Ptr, 261, 260, value)
}

// SetInfoText
func (m HistoricalOrderFillResponseFixedBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Unsafe(), 357, 261, value)
}

// SetInfoText
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Ptr, 357, 261, value)
}

// SetHighPriceDuringPosition
func (m HistoricalOrderFillResponseFixedBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 368, 360, value)
}

// SetHighPriceDuringPosition
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Ptr, 368, 360, value)
}

// SetLowPriceDuringPosition
func (m HistoricalOrderFillResponseFixedBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 376, 368, value)
}

// SetLowPriceDuringPosition
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Ptr, 376, 368, value)
}

// SetPositionQuantity
func (m HistoricalOrderFillResponseFixedBuilder) SetPositionQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 384, 376, value)
}

// SetPositionQuantity
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetPositionQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 384, 376, value)
}

// Copy
func (m HistoricalOrderFillResponse) Copy(to HistoricalOrderFillResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// Copy
func (m HistoricalOrderFillResponseBuilder) Copy(to HistoricalOrderFillResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyPointer
func (m HistoricalOrderFillResponse) CopyPointer(to HistoricalOrderFillResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyPointer
func (m HistoricalOrderFillResponseBuilder) CopyPointer(to HistoricalOrderFillResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyToPointer
func (m HistoricalOrderFillResponse) CopyToPointer(to HistoricalOrderFillResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyToPointer
func (m HistoricalOrderFillResponseBuilder) CopyToPointer(to HistoricalOrderFillResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyTo
func (m HistoricalOrderFillResponse) CopyTo(to HistoricalOrderFillResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyTo
func (m HistoricalOrderFillResponseBuilder) CopyTo(to HistoricalOrderFillResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// Copy
func (m HistoricalOrderFillResponseFixed) Copy(to HistoricalOrderFillResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// Copy
func (m HistoricalOrderFillResponseFixedBuilder) Copy(to HistoricalOrderFillResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyPointer
func (m HistoricalOrderFillResponseFixed) CopyPointer(to HistoricalOrderFillResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyPointer
func (m HistoricalOrderFillResponseFixedBuilder) CopyPointer(to HistoricalOrderFillResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyToPointer
func (m HistoricalOrderFillResponseFixed) CopyToPointer(to HistoricalOrderFillResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyToPointer
func (m HistoricalOrderFillResponseFixedBuilder) CopyToPointer(to HistoricalOrderFillResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyTo
func (m HistoricalOrderFillResponseFixed) CopyTo(to HistoricalOrderFillResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyTo
func (m HistoricalOrderFillResponseFixedBuilder) CopyTo(to HistoricalOrderFillResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// Copy
func (m HistoricalOrderFillResponsePointer) Copy(to HistoricalOrderFillResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// Copy
func (m HistoricalOrderFillResponsePointerBuilder) Copy(to HistoricalOrderFillResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyPointer
func (m HistoricalOrderFillResponsePointer) CopyPointer(to HistoricalOrderFillResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyPointer
func (m HistoricalOrderFillResponsePointerBuilder) CopyPointer(to HistoricalOrderFillResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyTo
func (m HistoricalOrderFillResponsePointer) CopyTo(to HistoricalOrderFillResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyTo
func (m HistoricalOrderFillResponsePointerBuilder) CopyTo(to HistoricalOrderFillResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyToPointer
func (m HistoricalOrderFillResponsePointer) CopyToPointer(to HistoricalOrderFillResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyToPointer
func (m HistoricalOrderFillResponsePointerBuilder) CopyToPointer(to HistoricalOrderFillResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// Copy
func (m HistoricalOrderFillResponseFixedPointer) Copy(to HistoricalOrderFillResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// Copy
func (m HistoricalOrderFillResponseFixedPointerBuilder) Copy(to HistoricalOrderFillResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyPointer
func (m HistoricalOrderFillResponseFixedPointer) CopyPointer(to HistoricalOrderFillResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyPointer
func (m HistoricalOrderFillResponseFixedPointerBuilder) CopyPointer(to HistoricalOrderFillResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyTo
func (m HistoricalOrderFillResponseFixedPointer) CopyTo(to HistoricalOrderFillResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyTo
func (m HistoricalOrderFillResponseFixedPointerBuilder) CopyTo(to HistoricalOrderFillResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyToPointer
func (m HistoricalOrderFillResponseFixedPointer) CopyToPointer(to HistoricalOrderFillResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyToPointer
func (m HistoricalOrderFillResponseFixedPointerBuilder) CopyToPointer(to HistoricalOrderFillResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}
