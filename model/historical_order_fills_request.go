package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalOrderFillsRequestSize = 32

const HistoricalOrderFillsRequestFixedSize = 88

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

func AllocHistoricalOrderFillsRequest() HistoricalOrderFillsRequestPointerBuilder {
	a := HistoricalOrderFillsRequestPointerBuilder{message.AllocVLSBuilder(32)}
	a.Ptr.SetBytes(0, _HistoricalOrderFillsRequestDefault)
	return a
}

func AllocHistoricalOrderFillsRequestFrom(b []byte) HistoricalOrderFillsRequestPointer {
	return HistoricalOrderFillsRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m HistoricalOrderFillsRequest) ToBuilder() HistoricalOrderFillsRequestBuilder {
	return HistoricalOrderFillsRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m HistoricalOrderFillsRequestFixed) ToBuilder() HistoricalOrderFillsRequestFixedBuilder {
	return HistoricalOrderFillsRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalOrderFillsRequestPointer) ToBuilder() HistoricalOrderFillsRequestPointerBuilder {
	return HistoricalOrderFillsRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m HistoricalOrderFillsRequestFixedPointer) ToBuilder() HistoricalOrderFillsRequestFixedPointerBuilder {
	return HistoricalOrderFillsRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalOrderFillsRequestBuilder) Finish() HistoricalOrderFillsRequest {
	return HistoricalOrderFillsRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m HistoricalOrderFillsRequestFixedBuilder) Finish() HistoricalOrderFillsRequestFixed {
	return HistoricalOrderFillsRequestFixed{m.Fixed}
}

// Finish
func (m *HistoricalOrderFillsRequestPointerBuilder) Finish() HistoricalOrderFillsRequestPointer {
	return HistoricalOrderFillsRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *HistoricalOrderFillsRequestFixedPointerBuilder) Finish() HistoricalOrderFillsRequestFixedPointer {
	return HistoricalOrderFillsRequestFixedPointer{m.FixedPointer}
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
func (m HistoricalOrderFillsRequest) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// ServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m HistoricalOrderFillsRequestBuilder) ServerOrderID() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
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
func (m HistoricalOrderFillsRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
}

// TradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m HistoricalOrderFillsRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 24, 20)
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
func (m HistoricalOrderFillsRequestFixed) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// ServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m HistoricalOrderFillsRequestFixedBuilder) ServerOrderID() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
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
func (m HistoricalOrderFillsRequestFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 76, 44)
}

// TradeAccount This specifies the particular Trade Account to request order fills for.
// This specifies the particular Trade Account to request order fills for.
func (m HistoricalOrderFillsRequestFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 76, 44)
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
func (m HistoricalOrderFillsRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m *HistoricalOrderFillsRequestBuilder) SetServerOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
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
func (m HistoricalOrderFillsRequestBuilder) SetNumberOfDays(value int32) {
	message.SetInt32VLS(m.Unsafe(), 20, 16, value)
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
func (m *HistoricalOrderFillsRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 24, 20, value)
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
func (m HistoricalOrderFillsRequestBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 32, 24, int64(value))
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

// SetRequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID A unique request identifier. The Server will return the same identifier
// in the response.
func (m HistoricalOrderFillsRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetServerOrderID Leave empty if want all order fills. Otherwise, request order fills for
// given Server Order identifier.
func (m HistoricalOrderFillsRequestFixedBuilder) SetServerOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
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
func (m HistoricalOrderFillsRequestFixedBuilder) SetNumberOfDays(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 44, 40, value)
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
func (m HistoricalOrderFillsRequestFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 76, 44, value)
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
func (m HistoricalOrderFillsRequestFixedBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 88, 80, int64(value))
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