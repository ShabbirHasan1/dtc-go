package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
func (m HistoricalOrderFillResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalOrderFillResponseFixedDefault)
}

// ToBuilder
func (m HistoricalOrderFillResponseFixedPointer) ToBuilder() HistoricalOrderFillResponseFixedPointerBuilder {
	return HistoricalOrderFillResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalOrderFillResponseFixedPointerBuilder) Finish() HistoricalOrderFillResponseFixedPointer {
	return HistoricalOrderFillResponseFixedPointer{m.FixedPointer}
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
func (m HistoricalOrderFillResponseFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 80, 16)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 80, 16)
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
func (m HistoricalOrderFillResponseFixedPointer) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 128, 96)
}

// ServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponseFixedPointerBuilder) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 128, 96)
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
func (m HistoricalOrderFillResponseFixedPointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 144, 136)
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 144, 136)
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
func (m HistoricalOrderFillResponseFixedPointer) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 160, 152)
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 160, 152)
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
func (m HistoricalOrderFillResponseFixedPointer) InfoText() string {
	return message.StringFixed(m.Ptr, 357, 261)
}

// InfoText
func (m HistoricalOrderFillResponseFixedPointerBuilder) InfoText() string {
	return message.StringFixed(m.Ptr, 357, 261)
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
func (m HistoricalOrderFillResponseFixedPointer) LowPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Ptr, 376, 368)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponseFixedPointerBuilder) LowPriceDuringPosition() float64 {
	return message.Float64Fixed(m.Ptr, 376, 368)
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
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
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
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Ptr, 16, 12, value)
}

// SetSymbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 80, 16, value)
}

// SetExchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 96, 80, value)
}

// SetServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 128, 96, value)
}

// SetBuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 132, 128, int32(value))
}

// SetPrice This is the price of the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 144, 136, value)
}

// SetDateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 152, 144, int64(value))
}

// SetQuantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 160, 152, value)
}

// SetUniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetUniqueExecutionID(value string) {
	message.SetStringFixed(m.Ptr, 224, 160, value)
}

// SetTradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 256, 224, value)
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
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetNoOrderFills(value uint8) {
	message.SetUint8Fixed(m.Ptr, 261, 260, value)
}

// SetInfoText
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Ptr, 357, 261, value)
}

// SetHighPriceDuringPosition
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Ptr, 368, 360, value)
}

// SetLowPriceDuringPosition
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64Fixed(m.Ptr, 376, 368, value)
}

// SetPositionQuantity
func (m HistoricalOrderFillResponseFixedPointerBuilder) SetPositionQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 384, 376, value)
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
