package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _HistoricalOrderFillsRequestDefault = []byte{32, 0, 47, 1, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const HistoricalOrderFillsRequestSize = 32

// HistoricalOrderFillsRequest This is a message from the Client to the Server to request order fills/executions
// for an order or orders.
type HistoricalOrderFillsRequest struct {
	message.VLS
}

// HistoricalOrderFillsRequestBuilder This is a message from the Client to the Server to request order fills/executions
// for an order or orders.
type HistoricalOrderFillsRequestBuilder struct {
	message.VLSBuilder
}

func NewHistoricalOrderFillsRequestFrom(b []byte) HistoricalOrderFillsRequest {
	return HistoricalOrderFillsRequest{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalOrderFillsRequest(b []byte) HistoricalOrderFillsRequest {
	return HistoricalOrderFillsRequest{VLS: message.WrapVLS(b)}
}

func NewHistoricalOrderFillsRequest() HistoricalOrderFillsRequestBuilder {
	a := HistoricalOrderFillsRequestBuilder{message.NewVLSBuilder(32)}
	a.Unsafe().SetBytes(0, _HistoricalOrderFillsRequestDefault)
	return a
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
func (m HistoricalOrderFillsRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalOrderFillsRequestDefault)
}

// ToBuilder
func (m HistoricalOrderFillsRequest) ToBuilder() HistoricalOrderFillsRequestBuilder {
	return HistoricalOrderFillsRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m HistoricalOrderFillsRequestBuilder) Finish() HistoricalOrderFillsRequest {
	return HistoricalOrderFillsRequest{m.VLSBuilder.Finish()}
}

// RequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// ServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m HistoricalOrderFillsRequest) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// ServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m HistoricalOrderFillsRequestBuilder) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// NumberOfDays The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The NumberOfDays field specifies to the Server to return order fills counting
// from the current day back by the specified number of days.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequest) NumberOfDays() int32 {
	return message.Int32VLS(m.Unsafe(), 20, 16)
}

// NumberOfDays The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The NumberOfDays field specifies to the Server to return order fills counting
// from the current day back by the specified number of days.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestBuilder) NumberOfDays() int32 {
	return message.Int32VLS(m.Unsafe(), 20, 16)
}

// TradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m HistoricalOrderFillsRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// TradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m HistoricalOrderFillsRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// StartDateTime The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The StartDateTime field specifies to the Server to return order fills
// starting with date time specified.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequest) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 32, 24))
}

// StartDateTime The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The StartDateTime field specifies to the Server to return order fills
// starting with date time specified.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 32, 24))
}

// SetRequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m *HistoricalOrderFillsRequestBuilder) SetServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetNumberOfDays The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The NumberOfDays field specifies to the Server to return order fills counting
// from the current day back by the specified number of days.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestBuilder) SetNumberOfDays(value int32) {
	message.SetInt32VLS(m.Unsafe(), 20, 16, value)
}

// SetTradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m *HistoricalOrderFillsRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
}

// SetStartDateTime The NumberOfDays field is ignored by the Server when StartDateTime is
// set.
//
// The StartDateTime field specifies to the Server to return order fills
// starting with date time specified.
//
// If NumberOfDays and StartDateTime are both not set or 0, the Server will
// return all historical order fills available.
func (m HistoricalOrderFillsRequestBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 32, 24, int64(value))
}

// Copy
func (m HistoricalOrderFillsRequest) Copy(to HistoricalOrderFillsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalOrderFillsRequestBuilder) Copy(to HistoricalOrderFillsRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalOrderFillsRequest) CopyPointer(to HistoricalOrderFillsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalOrderFillsRequestBuilder) CopyPointer(to HistoricalOrderFillsRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalOrderFillsRequest) CopyToPointer(to HistoricalOrderFillsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalOrderFillsRequestBuilder) CopyToPointer(to HistoricalOrderFillsRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalOrderFillsRequest) CopyTo(to HistoricalOrderFillsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalOrderFillsRequestBuilder) CopyTo(to HistoricalOrderFillsRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetNumberOfDays(m.NumberOfDays())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}
