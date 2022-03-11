package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const HistoricalTradesRequestSize = 31

const HistoricalTradesRequestFixedSize = 117

// {
//     Size                   = HistoricalTradesRequestSize  (31)
//     Type                   = HISTORICAL_TRADES_REQUEST  (10100)
//     BaseSize               = HistoricalTradesRequestSize  (31)
//     RequestID              = 0
//     Symbol                 = ""
//     TradeAccount           = ""
//     StartDateTime          = 0
//     SubAccountIdentifier   = 0
//     CreateFlatToFlatTrades = 0
// }
var _HistoricalTradesRequestDefault = []byte{31, 0, 116, 39, 31, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size                   = HistoricalTradesRequestFixedSize  (117)
//     Type                   = HISTORICAL_TRADES_REQUEST  (10100)
//     RequestID              = 0
//     Symbol                 = ""
//     TradeAccount           = ""
//     StartDateTime          = 0
//     SubAccountIdentifier   = 0
//     CreateFlatToFlatTrades = 0
// }
var _HistoricalTradesRequestFixedDefault = []byte{117, 0, 116, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HistoricalTradesRequest struct {
	message.VLS
}

type HistoricalTradesRequestBuilder struct {
	message.VLSBuilder
}

type HistoricalTradesRequestFixed struct {
	message.Fixed
}

type HistoricalTradesRequestFixedBuilder struct {
	message.Fixed
}

type HistoricalTradesRequestPointer struct {
	message.VLSPointer
}

type HistoricalTradesRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

type HistoricalTradesRequestFixedPointer struct {
	message.FixedPointer
}

type HistoricalTradesRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func NewHistoricalTradesRequestFrom(b []byte) HistoricalTradesRequest {
	return HistoricalTradesRequest{VLS: message.NewVLSFrom(b)}
}

func WrapHistoricalTradesRequest(b []byte) HistoricalTradesRequest {
	return HistoricalTradesRequest{VLS: message.WrapVLS(b)}
}

func NewHistoricalTradesRequest() HistoricalTradesRequestBuilder {
	a := HistoricalTradesRequestBuilder{message.NewVLSBuilder(31)}
	a.Unsafe().SetBytes(0, _HistoricalTradesRequestDefault)
	return a
}

func NewHistoricalTradesRequestFixedFrom(b []byte) HistoricalTradesRequestFixed {
	return HistoricalTradesRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapHistoricalTradesRequestFixed(b []byte) HistoricalTradesRequestFixed {
	return HistoricalTradesRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewHistoricalTradesRequestFixed() HistoricalTradesRequestFixedBuilder {
	a := HistoricalTradesRequestFixedBuilder{message.NewFixed(117)}
	a.Unsafe().SetBytes(0, _HistoricalTradesRequestFixedDefault)
	return a
}

func AllocHistoricalTradesRequest() HistoricalTradesRequestPointerBuilder {
	a := HistoricalTradesRequestPointerBuilder{message.AllocVLSBuilder(31)}
	a.Ptr.SetBytes(0, _HistoricalTradesRequestDefault)
	return a
}

func AllocHistoricalTradesRequestFrom(b []byte) HistoricalTradesRequestPointer {
	return HistoricalTradesRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocHistoricalTradesRequestFixed() HistoricalTradesRequestFixedPointerBuilder {
	a := HistoricalTradesRequestFixedPointerBuilder{message.AllocFixed(117)}
	a.Ptr.SetBytes(0, _HistoricalTradesRequestFixedDefault)
	return a
}

func AllocHistoricalTradesRequestFixedFrom(b []byte) HistoricalTradesRequestFixedPointer {
	return HistoricalTradesRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                   = HistoricalTradesRequestSize  (31)
//     Type                   = HISTORICAL_TRADES_REQUEST  (10100)
//     BaseSize               = HistoricalTradesRequestSize  (31)
//     RequestID              = 0
//     Symbol                 = ""
//     TradeAccount           = ""
//     StartDateTime          = 0
//     SubAccountIdentifier   = 0
//     CreateFlatToFlatTrades = 0
// }
func (m HistoricalTradesRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalTradesRequestDefault)
}

// Clear
// {
//     Size                   = HistoricalTradesRequestFixedSize  (117)
//     Type                   = HISTORICAL_TRADES_REQUEST  (10100)
//     RequestID              = 0
//     Symbol                 = ""
//     TradeAccount           = ""
//     StartDateTime          = 0
//     SubAccountIdentifier   = 0
//     CreateFlatToFlatTrades = 0
// }
func (m HistoricalTradesRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _HistoricalTradesRequestFixedDefault)
}

// Clear
// {
//     Size                   = HistoricalTradesRequestSize  (31)
//     Type                   = HISTORICAL_TRADES_REQUEST  (10100)
//     BaseSize               = HistoricalTradesRequestSize  (31)
//     RequestID              = 0
//     Symbol                 = ""
//     TradeAccount           = ""
//     StartDateTime          = 0
//     SubAccountIdentifier   = 0
//     CreateFlatToFlatTrades = 0
// }
func (m HistoricalTradesRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalTradesRequestDefault)
}

// Clear
// {
//     Size                   = HistoricalTradesRequestFixedSize  (117)
//     Type                   = HISTORICAL_TRADES_REQUEST  (10100)
//     RequestID              = 0
//     Symbol                 = ""
//     TradeAccount           = ""
//     StartDateTime          = 0
//     SubAccountIdentifier   = 0
//     CreateFlatToFlatTrades = 0
// }
func (m HistoricalTradesRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _HistoricalTradesRequestFixedDefault)
}

// ToBuilder
func (m HistoricalTradesRequest) ToBuilder() HistoricalTradesRequestBuilder {
	return HistoricalTradesRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m HistoricalTradesRequestFixed) ToBuilder() HistoricalTradesRequestFixedBuilder {
	return HistoricalTradesRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m HistoricalTradesRequestPointer) ToBuilder() HistoricalTradesRequestPointerBuilder {
	return HistoricalTradesRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m HistoricalTradesRequestFixedPointer) ToBuilder() HistoricalTradesRequestFixedPointerBuilder {
	return HistoricalTradesRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m HistoricalTradesRequestBuilder) Finish() HistoricalTradesRequest {
	return HistoricalTradesRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m HistoricalTradesRequestFixedBuilder) Finish() HistoricalTradesRequestFixed {
	return HistoricalTradesRequestFixed{m.Fixed}
}

// Finish
func (m *HistoricalTradesRequestPointerBuilder) Finish() HistoricalTradesRequestPointer {
	return HistoricalTradesRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *HistoricalTradesRequestFixedPointerBuilder) Finish() HistoricalTradesRequestFixedPointer {
	return HistoricalTradesRequestFixedPointer{m.FixedPointer}
}

// RequestID
func (m HistoricalTradesRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m HistoricalTradesRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m HistoricalTradesRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m HistoricalTradesRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 10, 6)
}

// Symbol
func (m HistoricalTradesRequest) Symbol() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Symbol
func (m HistoricalTradesRequestBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Symbol
func (m HistoricalTradesRequestPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Symbol
func (m HistoricalTradesRequestPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m HistoricalTradesRequest) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount
func (m HistoricalTradesRequestBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount
func (m HistoricalTradesRequestPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// TradeAccount
func (m HistoricalTradesRequestPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// StartDateTime
func (m HistoricalTradesRequest) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 26, 18))
}

// StartDateTime
func (m HistoricalTradesRequestBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 26, 18))
}

// StartDateTime
func (m HistoricalTradesRequestPointer) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 26, 18))
}

// StartDateTime
func (m HistoricalTradesRequestPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 26, 18))
}

// SubAccountIdentifier
func (m HistoricalTradesRequest) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 30, 26)
}

// SubAccountIdentifier
func (m HistoricalTradesRequestBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 30, 26)
}

// SubAccountIdentifier
func (m HistoricalTradesRequestPointer) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 30, 26)
}

// SubAccountIdentifier
func (m HistoricalTradesRequestPointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 30, 26)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequest) CreateFlatToFlatTrades() uint8 {
	return message.Uint8VLS(m.Unsafe(), 31, 30)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestBuilder) CreateFlatToFlatTrades() uint8 {
	return message.Uint8VLS(m.Unsafe(), 31, 30)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestPointer) CreateFlatToFlatTrades() uint8 {
	return message.Uint8VLS(m.Ptr, 31, 30)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestPointerBuilder) CreateFlatToFlatTrades() uint8 {
	return message.Uint8VLS(m.Ptr, 31, 30)
}

// RequestID
func (m HistoricalTradesRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalTradesRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m HistoricalTradesRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m HistoricalTradesRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// Symbol
func (m HistoricalTradesRequestFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol
func (m HistoricalTradesRequestFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol
func (m HistoricalTradesRequestFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Symbol
func (m HistoricalTradesRequestFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// TradeAccount
func (m HistoricalTradesRequestFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 104, 72)
}

// TradeAccount
func (m HistoricalTradesRequestFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 104, 72)
}

// TradeAccount
func (m HistoricalTradesRequestFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 104, 72)
}

// TradeAccount
func (m HistoricalTradesRequestFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 104, 72)
}

// StartDateTime
func (m HistoricalTradesRequestFixed) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 112, 104))
}

// StartDateTime
func (m HistoricalTradesRequestFixedBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 112, 104))
}

// StartDateTime
func (m HistoricalTradesRequestFixedPointer) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 112, 104))
}

// StartDateTime
func (m HistoricalTradesRequestFixedPointerBuilder) StartDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 112, 104))
}

// SubAccountIdentifier
func (m HistoricalTradesRequestFixed) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 116, 112)
}

// SubAccountIdentifier
func (m HistoricalTradesRequestFixedBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 116, 112)
}

// SubAccountIdentifier
func (m HistoricalTradesRequestFixedPointer) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Ptr, 116, 112)
}

// SubAccountIdentifier
func (m HistoricalTradesRequestFixedPointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Ptr, 116, 112)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestFixed) CreateFlatToFlatTrades() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 117, 116)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestFixedBuilder) CreateFlatToFlatTrades() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 117, 116)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestFixedPointer) CreateFlatToFlatTrades() uint8 {
	return message.Uint8Fixed(m.Ptr, 117, 116)
}

// CreateFlatToFlatTrades
func (m HistoricalTradesRequestFixedPointerBuilder) CreateFlatToFlatTrades() uint8 {
	return message.Uint8Fixed(m.Ptr, 117, 116)
}

// SetRequestID
func (m HistoricalTradesRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m HistoricalTradesRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 10, 6, value)
}

// SetSymbol
func (m *HistoricalTradesRequestBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetSymbol
func (m *HistoricalTradesRequestPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *HistoricalTradesRequestBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetTradeAccount
func (m *HistoricalTradesRequestPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetStartDateTime
func (m HistoricalTradesRequestBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 26, 18, int64(value))
}

// SetStartDateTime
func (m HistoricalTradesRequestPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 26, 18, int64(value))
}

// SetSubAccountIdentifier
func (m HistoricalTradesRequestBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 30, 26, value)
}

// SetSubAccountIdentifier
func (m HistoricalTradesRequestPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Ptr, 30, 26, value)
}

// SetCreateFlatToFlatTrades
func (m HistoricalTradesRequestBuilder) SetCreateFlatToFlatTrades(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 31, 30, value)
}

// SetCreateFlatToFlatTrades
func (m HistoricalTradesRequestPointerBuilder) SetCreateFlatToFlatTrades(value uint8) {
	message.SetUint8VLS(m.Ptr, 31, 30, value)
}

// SetRequestID
func (m HistoricalTradesRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m HistoricalTradesRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetSymbol
func (m HistoricalTradesRequestFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
}

// SetSymbol
func (m HistoricalTradesRequestFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 72, 8, value)
}

// SetTradeAccount
func (m HistoricalTradesRequestFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 72, value)
}

// SetTradeAccount
func (m HistoricalTradesRequestFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 104, 72, value)
}

// SetStartDateTime
func (m HistoricalTradesRequestFixedBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 112, 104, int64(value))
}

// SetStartDateTime
func (m HistoricalTradesRequestFixedPointerBuilder) SetStartDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 112, 104, int64(value))
}

// SetSubAccountIdentifier
func (m HistoricalTradesRequestFixedBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 116, 112, value)
}

// SetSubAccountIdentifier
func (m HistoricalTradesRequestFixedPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Ptr, 116, 112, value)
}

// SetCreateFlatToFlatTrades
func (m HistoricalTradesRequestFixedBuilder) SetCreateFlatToFlatTrades(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 117, 116, value)
}

// SetCreateFlatToFlatTrades
func (m HistoricalTradesRequestFixedPointerBuilder) SetCreateFlatToFlatTrades(value uint8) {
	message.SetUint8Fixed(m.Ptr, 117, 116, value)
}

// Copy
func (m HistoricalTradesRequest) Copy(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// Copy
func (m HistoricalTradesRequestBuilder) Copy(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequest) CopyPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestBuilder) CopyPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequest) CopyToPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestBuilder) CopyToPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequest) CopyTo(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestBuilder) CopyTo(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// Copy
func (m HistoricalTradesRequestFixed) Copy(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// Copy
func (m HistoricalTradesRequestFixedBuilder) Copy(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestFixed) CopyPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestFixedBuilder) CopyPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestFixed) CopyToPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestFixedBuilder) CopyToPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestFixed) CopyTo(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestFixedBuilder) CopyTo(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// Copy
func (m HistoricalTradesRequestPointer) Copy(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// Copy
func (m HistoricalTradesRequestPointerBuilder) Copy(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestPointer) CopyPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestPointerBuilder) CopyPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestPointer) CopyTo(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestPointerBuilder) CopyTo(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestPointer) CopyToPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestPointerBuilder) CopyToPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// Copy
func (m HistoricalTradesRequestFixedPointer) Copy(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// Copy
func (m HistoricalTradesRequestFixedPointerBuilder) Copy(to HistoricalTradesRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestFixedPointer) CopyPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyPointer
func (m HistoricalTradesRequestFixedPointerBuilder) CopyPointer(to HistoricalTradesRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestFixedPointer) CopyTo(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyTo
func (m HistoricalTradesRequestFixedPointerBuilder) CopyTo(to HistoricalTradesRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestFixedPointer) CopyToPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}

// CopyToPointer
func (m HistoricalTradesRequestFixedPointerBuilder) CopyToPointer(to HistoricalTradesRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetCreateFlatToFlatTrades(m.CreateFlatToFlatTrades())
}
