package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                    = HistoricalOrderFillResponseFixedSize  (384)
//     Type                    = HISTORICAL_ORDER_FILL_RESPONSE  (304)
//     RequestID               = 0
//     TotalNumberMessages     = 0
//     MessageNumber           = 0
//     Symbol                  = ""
//     Exchange                = ""
//     ServerOrderID           = ""
//     BuySell                 = 0
//     Price                   = 0
//     DateTime                = 0
//     Quantity                = 0
//     UniqueExecutionID       = ""
//     TradeAccount            = ""
//     OpenClose               = 0
//     NoOrderFills            = 0
//     InfoText                = ""
//     HighPriceDuringPosition = 0
//     LowPriceDuringPosition  = 0
//     PositionQuantity        = math.MaxFloat64
// }
var _HistoricalOrderFillResponseFixedDefault = []byte{128, 1, 48, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127}

const HistoricalOrderFillResponseFixedSize = 384

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
//     BuySell                 = 0
//     Price                   = 0
//     DateTime                = 0
//     Quantity                = 0
//     UniqueExecutionID       = ""
//     TradeAccount            = ""
//     OpenClose               = 0
//     NoOrderFills            = 0
//     InfoText                = ""
//     HighPriceDuringPosition = 0
//     LowPriceDuringPosition  = 0
//     PositionQuantity        = math.MaxFloat64
// }
func (m HistoricalOrderFillResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalOrderFillResponseFixedDefault)
}

// ToBuilder
func (m HistoricalOrderFillResponseFixed) ToBuilder() HistoricalOrderFillResponseFixedBuilder {
	return HistoricalOrderFillResponseFixedBuilder{m.Fixed}
}

// Finish
func (m HistoricalOrderFillResponseFixedBuilder) Finish() HistoricalOrderFillResponseFixed {
	return HistoricalOrderFillResponseFixed{m.Fixed}
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

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 80, 16)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 80, 16)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponseFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 96, 80)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponseFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 96, 80)
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

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseFixed) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 132, 128))
}

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseFixedBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 132, 128))
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponseFixed) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 144, 136)
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 144, 136)
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseFixed) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 152, 144))
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) DateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 152, 144))
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseFixed) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 160, 152)
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 160, 152)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponseFixed) UniqueExecutionID() string {
	return message.StringFixed(m.Unsafe(), 224, 160)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) UniqueExecutionID() string {
	return message.StringFixed(m.Unsafe(), 224, 160)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponseFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 256, 224)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponseFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 256, 224)
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

// InfoText
func (m HistoricalOrderFillResponseFixed) InfoText() string {
	return message.StringFixed(m.Unsafe(), 357, 261)
}

// InfoText
func (m HistoricalOrderFillResponseFixedBuilder) InfoText() string {
	return message.StringFixed(m.Unsafe(), 357, 261)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponseFixed) HighPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Unsafe(), 368, 360)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponseFixedBuilder) HighPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Unsafe(), 368, 360)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponseFixed) LowPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Unsafe(), 376, 368)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponseFixedBuilder) LowPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Unsafe(), 376, 368)
}

// PositionQuantity
func (m HistoricalOrderFillResponseFixed) PositionQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 384, 376)
}

// PositionQuantity
func (m HistoricalOrderFillResponseFixedBuilder) PositionQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 384, 376)
}

// SetRequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponseFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetTotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponseFixedBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, value)
}

// SetMessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponseFixedBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 16, 12, value)
}

// SetSymbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 80, 16, value)
}

// SetExchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponseFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 96, 80, value)
}

// SetServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponseFixedBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 128, 96, value)
}

// SetBuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseFixedBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 132, 128, int32(value))
}

// SetPrice This is the price of the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 144, 136, value)
}

// SetDateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) SetDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 152, 144, int64(value))
}

// SetQuantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 160, 152, value)
}

// SetUniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponseFixedBuilder) SetUniqueExecutionID(value string) {
	message.SetStringFixed(m.Unsafe(), 224, 160, value)
}

// SetTradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponseFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 256, 224, value)
}

// SetOpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponseFixedBuilder) SetOpenClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 260, 256, int32(value))
}

// SetNoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponseFixedBuilder) SetNoOrderFills(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 261, 260, value)
}

// SetInfoText
func (m HistoricalOrderFillResponseFixedBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Unsafe(), 357, 261, value)
}

// SetHighPriceDuringPosition
func (m HistoricalOrderFillResponseFixedBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 368, 360, value)
}

// SetLowPriceDuringPosition
func (m HistoricalOrderFillResponseFixedBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 376, 368, value)
}

// SetPositionQuantity
func (m HistoricalOrderFillResponseFixedBuilder) SetPositionQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 384, 376, value)
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
