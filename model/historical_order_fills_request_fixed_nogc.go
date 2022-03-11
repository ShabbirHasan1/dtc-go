package model

import (
	"github.com/moontrade/dtc-go/message"
)

// HistoricalOrderFillsRequestFixedPointer This is a message from the Client to the Server to request order fills/executions
// for an order or orders.
type HistoricalOrderFillsRequestFixedPointer struct {
	message.FixedPointer
}

// HistoricalOrderFillsRequestFixedPointerBuilder This is a message from the Client to the Server to request order fills/executions
// for an order or orders.
type HistoricalOrderFillsRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocHistoricalOrderFillsRequestFixed() HistoricalOrderFillsRequestFixedPointerBuilder {
	a := HistoricalOrderFillsRequestFixedPointerBuilder{message.AllocFixed(88)}
	a.Ptr.SetBytes(0, _HistoricalOrderFillsRequestFixedDefault)
	return a
}

func AllocHistoricalOrderFillsRequestFixedFrom(b []byte) HistoricalOrderFillsRequestFixedPointer {
	return HistoricalOrderFillsRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
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
func (m HistoricalOrderFillsRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalOrderFillsRequestFixedDefault)
}

// ToBuilder
func (m HistoricalOrderFillsRequestFixedPointer) ToBuilder() HistoricalOrderFillsRequestFixedPointerBuilder {
	return HistoricalOrderFillsRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *HistoricalOrderFillsRequestFixedPointerBuilder) Finish() HistoricalOrderFillsRequestFixedPointer {
	return HistoricalOrderFillsRequestFixedPointer{m.FixedPointer}
}

// RequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// ServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m HistoricalOrderFillsRequestFixedPointer) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// ServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m HistoricalOrderFillsRequestFixedPointerBuilder) ServerOrderID() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// NumberOfDays The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The NumberOfDays field specifies to the Server to return order fills counting
// from the current day back by the specified number of days.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestFixedPointer) NumberOfDays() int32 {
	return message.Int32Fixed(m.Ptr, 44, 40)
}

// NumberOfDays The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The NumberOfDays field specifies to the Server to return order fills counting
// from the current day back by the specified number of days.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestFixedPointerBuilder) NumberOfDays() int32 {
	return message.Int32Fixed(m.Ptr, 44, 40)
}

// TradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m HistoricalOrderFillsRequestFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 76, 44)
}

// TradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m HistoricalOrderFillsRequestFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 76, 44)
}

// StartDateTime The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The StartDateTime field specifies to the Server to return order fills
// starting with date time specified.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestFixedPointer) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 88, 80))
}

// StartDateTime The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The StartDateTime field specifies to the Server to return order fills
// starting with date time specified.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestFixedPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 88, 80))
}

// SetRequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m HistoricalOrderFillsRequestFixedPointerBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// SetNumberOfDays The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The NumberOfDays field specifies to the Server to return order fills counting
// from the current day back by the specified number of days.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestFixedPointerBuilder) SetNumberOfDays(value int32) {
	message.SetInt32Fixed(m.Ptr, 44, 40, value)
}

// SetTradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m HistoricalOrderFillsRequestFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 76, 44, value)
}

// SetStartDateTime The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The StartDateTime field specifies to the Server to return order fills
// starting with date time specified.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestFixedPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 88, 80, int64(value))
}

// Copy
func (m HistoricalOrderFillsRequestFixedPointer) Copy(to HistoricalOrderFillsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalOrderFillsRequestFixedPointerBuilder) Copy(to HistoricalOrderFillsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalOrderFillsRequestFixedPointer) CopyPointer(to HistoricalOrderFillsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalOrderFillsRequestFixedPointerBuilder) CopyPointer(to HistoricalOrderFillsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalOrderFillsRequestFixedPointer) CopyTo(to HistoricalOrderFillsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalOrderFillsRequestFixedPointerBuilder) CopyTo(to HistoricalOrderFillsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalOrderFillsRequestFixedPointer) CopyToPointer(to HistoricalOrderFillsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalOrderFillsRequestFixedPointerBuilder) CopyToPointer(to HistoricalOrderFillsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}
