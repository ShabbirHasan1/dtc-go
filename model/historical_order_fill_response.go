package model

import (
	"github.com/moontrade/dtc-go/message"
)

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

const HistoricalOrderFillResponseSize = 112

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

// ToBuilder
func (m HistoricalOrderFillResponse) ToBuilder() HistoricalOrderFillResponseBuilder {
	return HistoricalOrderFillResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m HistoricalOrderFillResponseBuilder) Finish() HistoricalOrderFillResponse {
	return HistoricalOrderFillResponse{m.VLSBuilder.Finish()}
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

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponse) Symbol() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponse) Exchange() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponseBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 28, 24)
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

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponse) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 36, 32))
}

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 36, 32))
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponse) Price() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponseBuilder) Price() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponse) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 56, 48))
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseBuilder) DateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 56, 48))
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponse) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 64, 56)
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseBuilder) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 64, 56)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponse) UniqueExecutionID() string {
	return message.StringVLS(m.Unsafe(), 68, 64)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponseBuilder) UniqueExecutionID() string {
	return message.StringVLS(m.Unsafe(), 68, 64)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponse) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 72, 68)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponseBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 72, 68)
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

// InfoText
func (m HistoricalOrderFillResponse) InfoText() string {
	return message.StringVLS(m.Unsafe(), 82, 78)
}

// InfoText
func (m HistoricalOrderFillResponseBuilder) InfoText() string {
	return message.StringVLS(m.Unsafe(), 82, 78)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponse) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 96, 88)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponseBuilder) HighPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 96, 88)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponse) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 104, 96)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponseBuilder) LowPriceDuringPosition() float64 {
	return message.Float64VLS(m.Unsafe(), 104, 96)
}

// PositionQuantity
func (m HistoricalOrderFillResponse) PositionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 112, 104)
}

// PositionQuantity
func (m HistoricalOrderFillResponseBuilder) PositionQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 112, 104)
}

// SetRequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponseBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetTotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponseBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32VLS(m.Unsafe(), 16, 12, value)
}

// SetMessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponseBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Unsafe(), 20, 16, value)
}

// SetSymbol The symbol the order fill is for.
func (m *HistoricalOrderFillResponseBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetExchange The optional exchange for the symbol.
func (m *HistoricalOrderFillResponseBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 28, 24, value)
}

// SetServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m *HistoricalOrderFillResponseBuilder) SetServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 32, 28, value)
}

// SetBuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 36, 32, int32(value))
}

// SetPrice This is the price of the order fill.
func (m HistoricalOrderFillResponseBuilder) SetPrice(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 48, 40, value)
}

// SetDateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseBuilder) SetDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 56, 48, int64(value))
}

// SetQuantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 64, 56, value)
}

// SetUniqueExecutionID This is the unique execution identifier for the order fill.
func (m *HistoricalOrderFillResponseBuilder) SetUniqueExecutionID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 68, 64, value)
}

// SetTradeAccount This is the trade account that the order fill is associated with.
func (m *HistoricalOrderFillResponseBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 72, 68, value)
}

// SetOpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponseBuilder) SetOpenClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Unsafe(), 76, 72, int32(value))
}

// SetNoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponseBuilder) SetNoOrderFills(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 77, 76, value)
}

// SetInfoText
func (m *HistoricalOrderFillResponseBuilder) SetInfoText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 82, 78, value)
}

// SetHighPriceDuringPosition
func (m HistoricalOrderFillResponseBuilder) SetHighPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 96, 88, value)
}

// SetLowPriceDuringPosition
func (m HistoricalOrderFillResponseBuilder) SetLowPriceDuringPosition(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 104, 96, value)
}

// SetPositionQuantity
func (m HistoricalOrderFillResponseBuilder) SetPositionQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 112, 104, value)
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
