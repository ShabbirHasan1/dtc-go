package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = HistoricalOrderFillsRequestFixedSize  (88)
//     Type          = HISTORICAL_ORDER_FILLS_REQUEST  (303)
//     RequestID     = 0
//     ServerOrderID = ""
//     NumberOfDays  = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
var _HistoricalOrderFillsRequestFixedDefault = []byte{88, 0, 47, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalOrderFillsRequestFixedSize = 88

// HistoricalOrderFillsRequestFixed This is a message from the Client to the Server to request order fills/executions
// for an order or orders.
type HistoricalOrderFillsRequestFixed struct {
	message.Fixed
}

// HistoricalOrderFillsRequestFixedBuilder This is a message from the Client to the Server to request order fills/executions
// for an order or orders.
type HistoricalOrderFillsRequestFixedBuilder struct {
	message.Fixed
}

func NewHistoricalOrderFillsRequestFixedFrom(b []byte) HistoricalOrderFillsRequestFixed {
	return HistoricalOrderFillsRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalOrderFillsRequestFixed(b []byte) HistoricalOrderFillsRequestFixed {
	return HistoricalOrderFillsRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalOrderFillsRequestFixed() HistoricalOrderFillsRequestFixedBuilder {
	a := HistoricalOrderFillsRequestFixedBuilder{message.NewFixed(88)}
	a.Unsafe().SetBytes(0, _HistoricalOrderFillsRequestFixedDefault)
	return a
}

// Clear
// {
//     Size          = HistoricalOrderFillsRequestFixedSize  (88)
//     Type          = HISTORICAL_ORDER_FILLS_REQUEST  (303)
//     RequestID     = 0
//     ServerOrderID = ""
//     NumberOfDays  = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
func (m HistoricalOrderFillsRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalOrderFillsRequestFixedDefault)
}

// ToBuilder
func (m HistoricalOrderFillsRequestFixed) ToBuilder() HistoricalOrderFillsRequestFixedBuilder {
	return HistoricalOrderFillsRequestFixedBuilder{m.Fixed}
}

// Finish
func (m HistoricalOrderFillsRequestFixedBuilder) Finish() HistoricalOrderFillsRequestFixed {
	return HistoricalOrderFillsRequestFixed{m.Fixed}
}

// RequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// ServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m HistoricalOrderFillsRequestFixed) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// ServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m HistoricalOrderFillsRequestFixedBuilder) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// NumberOfDays The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The NumberOfDays field specifies to the Server to return order fills counting
// from the current day back by the specified number of days.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestFixed) NumberOfDays() int32 {
	return message.Int32Fixed(m.Unsafe(), 44, 40)
}

// NumberOfDays The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The NumberOfDays field specifies to the Server to return order fills counting
// from the current day back by the specified number of days.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestFixedBuilder) NumberOfDays() int32 {
	return message.Int32Fixed(m.Unsafe(), 44, 40)
}

// TradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m HistoricalOrderFillsRequestFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 76, 44)
}

// TradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m HistoricalOrderFillsRequestFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 76, 44)
}

// StartDateTime The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The StartDateTime field specifies to the Server to return order fills
// starting with date time specified.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestFixed) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 88, 80))
}

// StartDateTime The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The StartDateTime field specifies to the Server to return order fills
// starting with date time specified.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestFixedBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 88, 80))
}

// SetRequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m HistoricalOrderFillsRequestFixedBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetNumberOfDays The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The NumberOfDays field specifies to the Server to return order fills counting
// from the current day back by the specified number of days.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestFixedBuilder) SetNumberOfDays(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 44, 40, value)
}

// SetTradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m HistoricalOrderFillsRequestFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 76, 44, value)
}

// SetStartDateTime The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The StartDateTime field specifies to the Server to return order fills
// starting with date time specified.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestFixedBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 88, 80, int64(value))
}

// Copy
func (m HistoricalOrderFillsRequestFixed) Copy(to HistoricalOrderFillsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalOrderFillsRequestFixedBuilder) Copy(to HistoricalOrderFillsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalOrderFillsRequestFixed) CopyPointer(to HistoricalOrderFillsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalOrderFillsRequestFixedBuilder) CopyPointer(to HistoricalOrderFillsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalOrderFillsRequestFixed) CopyToPointer(to HistoricalOrderFillsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalOrderFillsRequestFixedBuilder) CopyToPointer(to HistoricalOrderFillsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalOrderFillsRequestFixed) CopyTo(to HistoricalOrderFillsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalOrderFillsRequestFixedBuilder) CopyTo(to HistoricalOrderFillsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}
