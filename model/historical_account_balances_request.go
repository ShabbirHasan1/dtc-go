package model

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalAccountBalancesRequestSize = 24

const HistoricalAccountBalancesRequestFixedSize = 48

// {
//     Size          = HistoricalAccountBalancesRequestSize  (24)
//     Type          = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
//     BaseSize      = HistoricalAccountBalancesRequestSize  (24)
//     RequestID     = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
var _HistoricalAccountBalancesRequestDefault = []byte{24, 0, 91, 2, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size          = HistoricalAccountBalancesRequestFixedSize  (48)
//     Type          = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
//     RequestID     = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
var _HistoricalAccountBalancesRequestFixedDefault = []byte{48, 0, 91, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// HistoricalAccountBalancesRequest This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequest struct {
	message.VLS
}

// HistoricalAccountBalancesRequestBuilder This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestBuilder struct {
	message.VLSBuilder
}

// HistoricalAccountBalancesRequestFixed This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestFixed struct {
	message.Fixed
}

// HistoricalAccountBalancesRequestFixedBuilder This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestFixedBuilder struct {
	message.Fixed
}

// HistoricalAccountBalancesRequestPointer This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestPointer struct {
	message.VLSPointer
}

// HistoricalAccountBalancesRequestPointerBuilder This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

// HistoricalAccountBalancesRequestFixedPointer This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestFixedPointer struct {
	message.FixedPointer
}

// HistoricalAccountBalancesRequestFixedPointerBuilder This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalAccountBalancesRequestFrom(b []byte) HistoricalAccountBalancesRequest {
	return HistoricalAccountBalancesRequest{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalAccountBalancesRequest(b []byte) HistoricalAccountBalancesRequest {
	return HistoricalAccountBalancesRequest{VLS: message.WrapVLS(b)}
}

func NewHistoricalAccountBalancesRequest() HistoricalAccountBalancesRequestBuilder {
	a := HistoricalAccountBalancesRequestBuilder{message.NewVLSBuilder(24)}
	a.Unsafe().SetBytes(0, _HistoricalAccountBalancesRequestDefault)
	return a
}

func NewHistoricalAccountBalancesRequestFixedFrom(b []byte) HistoricalAccountBalancesRequestFixed {
	return HistoricalAccountBalancesRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalAccountBalancesRequestFixed(b []byte) HistoricalAccountBalancesRequestFixed {
	return HistoricalAccountBalancesRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalAccountBalancesRequestFixed() HistoricalAccountBalancesRequestFixedBuilder {
	a := HistoricalAccountBalancesRequestFixedBuilder{message.NewFixed(48)}
	a.Unsafe().SetBytes(0, _HistoricalAccountBalancesRequestFixedDefault)
	return a
}

func AllocHistoricalAccountBalancesRequest() HistoricalAccountBalancesRequestPointerBuilder {
	a := HistoricalAccountBalancesRequestPointerBuilder{message.AllocVLSBuilder(24)}
	a.Ptr.SetBytes(0, _HistoricalAccountBalancesRequestDefault)
	return a
}

func AllocHistoricalAccountBalancesRequestFrom(b []byte) HistoricalAccountBalancesRequestPointer {
	return HistoricalAccountBalancesRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocHistoricalAccountBalancesRequestFixed() HistoricalAccountBalancesRequestFixedPointerBuilder {
	a := HistoricalAccountBalancesRequestFixedPointerBuilder{message.AllocFixed(48)}
	a.Ptr.SetBytes(0, _HistoricalAccountBalancesRequestFixedDefault)
	return a
}

func AllocHistoricalAccountBalancesRequestFixedFrom(b []byte) HistoricalAccountBalancesRequestFixedPointer {
	return HistoricalAccountBalancesRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = HistoricalAccountBalancesRequestSize  (24)
//     Type          = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
//     BaseSize      = HistoricalAccountBalancesRequestSize  (24)
//     RequestID     = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
func (m HistoricalAccountBalancesRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalAccountBalancesRequestDefault)
}

// Clear
// {
//     Size          = HistoricalAccountBalancesRequestFixedSize  (48)
//     Type          = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
//     RequestID     = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
func (m HistoricalAccountBalancesRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalAccountBalancesRequestFixedDefault)
}

// Clear
// {
//     Size          = HistoricalAccountBalancesRequestSize  (24)
//     Type          = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
//     BaseSize      = HistoricalAccountBalancesRequestSize  (24)
//     RequestID     = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
func (m HistoricalAccountBalancesRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalAccountBalancesRequestDefault)
}

// Clear
// {
//     Size          = HistoricalAccountBalancesRequestFixedSize  (48)
//     Type          = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
//     RequestID     = 0
//     TradeAccount  = ""
//     StartDateTime = 0
// }
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalAccountBalancesRequestFixedDefault)
}

// ToBuilder
func (m HistoricalAccountBalancesRequest) ToBuilder() HistoricalAccountBalancesRequestBuilder {
	return HistoricalAccountBalancesRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m HistoricalAccountBalancesRequestFixed) ToBuilder() HistoricalAccountBalancesRequestFixedBuilder {
	return HistoricalAccountBalancesRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalAccountBalancesRequestPointer) ToBuilder() HistoricalAccountBalancesRequestPointerBuilder {
	return HistoricalAccountBalancesRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m HistoricalAccountBalancesRequestFixedPointer) ToBuilder() HistoricalAccountBalancesRequestFixedPointerBuilder {
	return HistoricalAccountBalancesRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalAccountBalancesRequestBuilder) Finish() HistoricalAccountBalancesRequest {
	return HistoricalAccountBalancesRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m HistoricalAccountBalancesRequestFixedBuilder) Finish() HistoricalAccountBalancesRequestFixed {
	return HistoricalAccountBalancesRequestFixed{m.Fixed}
}

// Finish
func (m *HistoricalAccountBalancesRequestPointerBuilder) Finish() HistoricalAccountBalancesRequestPointer {
	return HistoricalAccountBalancesRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *HistoricalAccountBalancesRequestFixedPointerBuilder) Finish() HistoricalAccountBalancesRequestFixedPointer {
	return HistoricalAccountBalancesRequestFixedPointer{m.FixedPointer}
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequest) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 24, 16))
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 24, 16))
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestPointer) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 24, 16))
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 24, 16))
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestFixed) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 48, 40))
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestFixedBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 48, 40))
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestFixedPointer) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 48, 40))
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 48, 40))
}

// SetRequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetTradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m *HistoricalAccountBalancesRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetTradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m *HistoricalAccountBalancesRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetStartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 24, 16, int64(value))
}

// SetStartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 24, 16, int64(value))
}

// SetRequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetTradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetTradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// SetStartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestFixedBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 48, 40, int64(value))
}

// SetStartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 48, 40, int64(value))
}

// Copy
func (m HistoricalAccountBalancesRequest) Copy(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalAccountBalancesRequestBuilder) Copy(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequest) CopyPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestBuilder) CopyPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequest) CopyToPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestBuilder) CopyToPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequest) CopyTo(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestBuilder) CopyTo(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalAccountBalancesRequestFixed) Copy(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalAccountBalancesRequestFixedBuilder) Copy(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestFixed) CopyPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestFixedBuilder) CopyPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestFixed) CopyToPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestFixedBuilder) CopyToPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestFixed) CopyTo(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestFixedBuilder) CopyTo(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalAccountBalancesRequestPointer) Copy(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalAccountBalancesRequestPointerBuilder) Copy(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestPointer) CopyPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestPointerBuilder) CopyPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestPointer) CopyTo(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestPointerBuilder) CopyTo(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestPointer) CopyToPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestPointerBuilder) CopyToPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalAccountBalancesRequestFixedPointer) Copy(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) Copy(to HistoricalAccountBalancesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestFixedPointer) CopyPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyPointer
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) CopyPointer(to HistoricalAccountBalancesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestFixedPointer) CopyTo(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) CopyTo(to HistoricalAccountBalancesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestFixedPointer) CopyToPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyToPointer
func (m HistoricalAccountBalancesRequestFixedPointerBuilder) CopyToPointer(to HistoricalAccountBalancesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}
