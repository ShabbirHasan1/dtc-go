package model

import (
	"github.com/moontrade/dtc-go/message"
)

// HistoricalOrderFillsRequestPointer This is a message from the Client to the Server to request order fills/executions
// for an order or orders.
type HistoricalOrderFillsRequestPointer struct {
	message.VLSPointer
}

// HistoricalOrderFillsRequestPointerBuilder This is a message from the Client to the Server to request order fills/executions
// for an order or orders.
type HistoricalOrderFillsRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocHistoricalOrderFillsRequest() HistoricalOrderFillsRequestPointerBuilder {
	a := HistoricalOrderFillsRequestPointerBuilder{message.AllocVLSBuilder(32)}
	a.Ptr.SetBytes(0, _HistoricalOrderFillsRequestDefault)
	return a
}

func AllocHistoricalOrderFillsRequestFrom(b []byte) HistoricalOrderFillsRequestPointer {
	return HistoricalOrderFillsRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size          = HistoricalOrderFillsRequestSize  (32)
//     Type          = HISTORICAL_ORDER_FILLS_REQUEST  (303)
//     BaseSize      = HistoricalOrderFillsRequestSize  (32)
//     RequestID     = 0
//     ServerOrderID = ""
//     NumberOfDays  = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
func (m HistoricalOrderFillsRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalOrderFillsRequestDefault)
}

// ToBuilder
func (m HistoricalOrderFillsRequestPointer) ToBuilder() HistoricalOrderFillsRequestPointerBuilder {
	return HistoricalOrderFillsRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *HistoricalOrderFillsRequestPointerBuilder) Finish() HistoricalOrderFillsRequestPointer {
	return HistoricalOrderFillsRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// ServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m HistoricalOrderFillsRequestPointer) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// ServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m HistoricalOrderFillsRequestPointerBuilder) ServerOrderID() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// NumberOfDays The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The NumberOfDays field specifies to the Server to return order fills counting
// from the current day back by the specified number of days.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestPointer) NumberOfDays() int32 {
	return message.Int32VLS(m.Ptr, 20, 16)
}

// NumberOfDays The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The NumberOfDays field specifies to the Server to return order fills counting
// from the current day back by the specified number of days.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestPointerBuilder) NumberOfDays() int32 {
	return message.Int32VLS(m.Ptr, 20, 16)
}

// TradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m HistoricalOrderFillsRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// TradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m HistoricalOrderFillsRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 24, 20)
}

// StartDateTime The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The StartDateTime field specifies to the Server to return order fills
// starting with date time specified.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestPointer) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 32, 24))
}

// StartDateTime The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The StartDateTime field specifies to the Server to return order fills
// starting with date time specified.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 32, 24))
}

// SetRequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m *HistoricalOrderFillsRequestPointerBuilder) SetServerOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetNumberOfDays The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The NumberOfDays field specifies to the Server to return order fills counting
// from the current day back by the specified number of days.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestPointerBuilder) SetNumberOfDays(value int32) {
	message.SetInt32VLS(m.Ptr, 20, 16, value)
}

// SetTradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m *HistoricalOrderFillsRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 24, 20, value)
}

// SetStartDateTime The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The StartDateTime field specifies to the Server to return order fills
// starting with date time specified.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 32, 24, int64(value))
}

// Copy
func (m HistoricalOrderFillsRequestPointer) Copy(to HistoricalOrderFillsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalOrderFillsRequestPointerBuilder) Copy(to HistoricalOrderFillsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalOrderFillsRequestPointer) CopyPointer(to HistoricalOrderFillsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalOrderFillsRequestPointerBuilder) CopyPointer(to HistoricalOrderFillsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalOrderFillsRequestPointer) CopyTo(to HistoricalOrderFillsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalOrderFillsRequestPointerBuilder) CopyTo(to HistoricalOrderFillsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalOrderFillsRequestPointer) CopyToPointer(to HistoricalOrderFillsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalOrderFillsRequestPointerBuilder) CopyToPointer(to HistoricalOrderFillsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}
