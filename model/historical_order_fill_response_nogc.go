package model

import (
	"github.com/moontrade/dtc-go/message"
)

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

func AllocHistoricalOrderFillResponse() HistoricalOrderFillResponsePointerBuilder {
	a := HistoricalOrderFillResponsePointerBuilder{message.AllocVLSBuilder(112)}
	a.Ptr.SetBytes(0, _HistoricalOrderFillResponseDefault)
	return a
}

func AllocHistoricalOrderFillResponseFrom(b []byte) HistoricalOrderFillResponsePointer {
	return HistoricalOrderFillResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
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

// ToBuilder
func (m HistoricalOrderFillResponsePointer) ToBuilder() HistoricalOrderFillResponsePointerBuilder {
	return HistoricalOrderFillResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *HistoricalOrderFillResponsePointerBuilder) Finish() HistoricalOrderFillResponsePointer {
	return HistoricalOrderFillResponsePointer{m.VLSPointerBuilder.Finish()}
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
func (m HistoricalOrderFillResponsePointer) Symbol() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponsePointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 24, 20)
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
func (m HistoricalOrderFillResponsePointer) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 32, 28)
}

// ServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponsePointerBuilder) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 32, 28)
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
func (m HistoricalOrderFillResponsePointer) Price() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponsePointerBuilder) Price() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
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
func (m HistoricalOrderFillResponsePointer) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 64, 56)
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponsePointerBuilder) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 64, 56)
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
func (m HistoricalOrderFillResponsePointer) InfoText() string {
	return message.StringVLS(m.Ptr, 82, 78)
}

// InfoText
func (m HistoricalOrderFillResponsePointerBuilder) InfoText() string {
	return message.StringVLS(m.Ptr, 82, 78)
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
func (m HistoricalOrderFillResponsePointer) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 104, 96)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponsePointerBuilder) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Ptr, 104, 96)
}

// PositionQuantity
func (m HistoricalOrderFillResponsePointer) PositionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 112, 104)
}

// PositionQuantity
func (m HistoricalOrderFillResponsePointerBuilder) PositionQuantity() float64 {
	return message.Float64VLS(m.Ptr, 112, 104)
}

// SetRequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponsePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
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
func (m HistoricalOrderFillResponsePointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Ptr, 20, 16, value)
}

// SetSymbol The symbol the order fill is for.
func (m *HistoricalOrderFillResponsePointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetExchange The optional exchange for the symbol.
func (m *HistoricalOrderFillResponsePointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 28, 24, value)
}

// SetServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m *HistoricalOrderFillResponsePointerBuilder) SetServerOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 32, 28, value)
}

// SetBuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponsePointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 36, 32, int32(value))
}

// SetPrice This is the price of the order fill.
func (m HistoricalOrderFillResponsePointerBuilder) SetPrice(value float64) {
	message.SetFloat64VLS(m.Ptr, 48, 40, value)
}

// SetDateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponsePointerBuilder) SetDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 56, 48, int64(value))
}

// SetQuantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponsePointerBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 64, 56, value)
}

// SetUniqueExecutionID This is the unique execution identifier for the order fill.
func (m *HistoricalOrderFillResponsePointerBuilder) SetUniqueExecutionID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 68, 64, value)
}

// SetTradeAccount This is the trade account that the order fill is associated with.
func (m *HistoricalOrderFillResponsePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 72, 68, value)
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
func (m HistoricalOrderFillResponsePointerBuilder) SetNoOrderFills(value uint8) {
	message.SetUint8VLS(m.Ptr, 77, 76, value)
}

// SetInfoText
func (m *HistoricalOrderFillResponsePointerBuilder) SetInfoText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 82, 78, value)
}

// SetHighPriceDuringPosition
func (m HistoricalOrderFillResponsePointerBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 96, 88, value)
}

// SetLowPriceDuringPosition
func (m HistoricalOrderFillResponsePointerBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Ptr, 104, 96, value)
}

// SetPositionQuantity
func (m HistoricalOrderFillResponsePointerBuilder) SetPositionQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 112, 104, value)
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
