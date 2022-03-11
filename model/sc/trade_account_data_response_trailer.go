package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const TradeAccountDataResponseTrailerSize = 17

// {
//     Size                  = TradeAccountDataResponseTrailerSize  (17)
//     Type                  = TRADE_ACCOUNT_DATA_RESPONSE_TRAILER  (10130)
//     BaseSize              = TradeAccountDataResponseTrailerSize  (17)
//     RequestID             = 0
//     IsSnapshot            = false
//     IsFirstMessageInBatch = false
//     IsLastMessageInBatch  = false
//     TradeAccount          = ""
// }
var _TradeAccountDataResponseTrailerDefault = []byte{17, 0, 146, 39, 17, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataResponseTrailer struct {
	message.VLS
}

type TradeAccountDataResponseTrailerBuilder struct {
	message.VLSBuilder
}

type TradeAccountDataResponseTrailerPointer struct {
	message.VLSPointer
}

type TradeAccountDataResponseTrailerPointerBuilder struct {
	message.VLSPointerBuilder
}

func NewTradeAccountDataResponseTrailerFrom(b []byte) TradeAccountDataResponseTrailer {
	return TradeAccountDataResponseTrailer{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataResponseTrailer(b []byte) TradeAccountDataResponseTrailer {
	return TradeAccountDataResponseTrailer{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataResponseTrailer() TradeAccountDataResponseTrailerBuilder {
	a := TradeAccountDataResponseTrailerBuilder{message.NewVLSBuilder(17)}
	a.Unsafe().SetBytes(0, _TradeAccountDataResponseTrailerDefault)
	return a
}

func AllocTradeAccountDataResponseTrailer() TradeAccountDataResponseTrailerPointerBuilder {
	a := TradeAccountDataResponseTrailerPointerBuilder{message.AllocVLSBuilder(17)}
	a.Ptr.SetBytes(0, _TradeAccountDataResponseTrailerDefault)
	return a
}

func AllocTradeAccountDataResponseTrailerFrom(b []byte) TradeAccountDataResponseTrailerPointer {
	return TradeAccountDataResponseTrailerPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                  = TradeAccountDataResponseTrailerSize  (17)
//     Type                  = TRADE_ACCOUNT_DATA_RESPONSE_TRAILER  (10130)
//     BaseSize              = TradeAccountDataResponseTrailerSize  (17)
//     RequestID             = 0
//     IsSnapshot            = false
//     IsFirstMessageInBatch = false
//     IsLastMessageInBatch  = false
//     TradeAccount          = ""
// }
func (m TradeAccountDataResponseTrailerBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataResponseTrailerDefault)
}

// Clear
// {
//     Size                  = TradeAccountDataResponseTrailerSize  (17)
//     Type                  = TRADE_ACCOUNT_DATA_RESPONSE_TRAILER  (10130)
//     BaseSize              = TradeAccountDataResponseTrailerSize  (17)
//     RequestID             = 0
//     IsSnapshot            = false
//     IsFirstMessageInBatch = false
//     IsLastMessageInBatch  = false
//     TradeAccount          = ""
// }
func (m TradeAccountDataResponseTrailerPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataResponseTrailerDefault)
}

// ToBuilder
func (m TradeAccountDataResponseTrailer) ToBuilder() TradeAccountDataResponseTrailerBuilder {
	return TradeAccountDataResponseTrailerBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m TradeAccountDataResponseTrailerPointer) ToBuilder() TradeAccountDataResponseTrailerPointerBuilder {
	return TradeAccountDataResponseTrailerPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m TradeAccountDataResponseTrailerBuilder) Finish() TradeAccountDataResponseTrailer {
	return TradeAccountDataResponseTrailer{m.VLSBuilder.Finish()}
}

// Finish
func (m *TradeAccountDataResponseTrailerPointerBuilder) Finish() TradeAccountDataResponseTrailerPointer {
	return TradeAccountDataResponseTrailerPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataResponseTrailer) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataResponseTrailerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataResponseTrailerPointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataResponseTrailerPointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// IsSnapshot
func (m TradeAccountDataResponseTrailer) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsSnapshot
func (m TradeAccountDataResponseTrailerBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 11, 10)
}

// IsSnapshot
func (m TradeAccountDataResponseTrailerPointer) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// IsSnapshot
func (m TradeAccountDataResponseTrailerPointerBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 11, 10)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponseTrailer) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 12, 11)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponseTrailerBuilder) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 12, 11)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponseTrailerPointer) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 12, 11)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponseTrailerPointerBuilder) IsFirstMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 12, 11)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponseTrailer) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 13, 12)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponseTrailerBuilder) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Unsafe(), 13, 12)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponseTrailerPointer) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 13, 12)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponseTrailerPointerBuilder) IsLastMessageInBatch() bool {
	return message.BoolVLS(m.Ptr, 13, 12)
}

// TradeAccount
func (m TradeAccountDataResponseTrailer) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 17, 13)
}

// TradeAccount
func (m TradeAccountDataResponseTrailerBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 17, 13)
}

// TradeAccount
func (m TradeAccountDataResponseTrailerPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 17, 13)
}

// TradeAccount
func (m TradeAccountDataResponseTrailerPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 17, 13)
}

// SetRequestID
func (m TradeAccountDataResponseTrailerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m TradeAccountDataResponseTrailerPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetIsSnapshot
func (m TradeAccountDataResponseTrailerBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Unsafe(), 11, 10, value)
}

// SetIsSnapshot
func (m TradeAccountDataResponseTrailerPointerBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Ptr, 11, 10, value)
}

// SetIsFirstMessageInBatch
func (m TradeAccountDataResponseTrailerBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolVLS(m.Unsafe(), 12, 11, value)
}

// SetIsFirstMessageInBatch
func (m TradeAccountDataResponseTrailerPointerBuilder) SetIsFirstMessageInBatch(value bool) {
	message.SetBoolVLS(m.Ptr, 12, 11, value)
}

// SetIsLastMessageInBatch
func (m TradeAccountDataResponseTrailerBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolVLS(m.Unsafe(), 13, 12, value)
}

// SetIsLastMessageInBatch
func (m TradeAccountDataResponseTrailerPointerBuilder) SetIsLastMessageInBatch(value bool) {
	message.SetBoolVLS(m.Ptr, 13, 12, value)
}

// SetTradeAccount
func (m *TradeAccountDataResponseTrailerBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 17, 13, value)
}

// SetTradeAccount
func (m *TradeAccountDataResponseTrailerPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 17, 13, value)
}

// Copy
func (m TradeAccountDataResponseTrailer) Copy(to TradeAccountDataResponseTrailerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m TradeAccountDataResponseTrailerBuilder) Copy(to TradeAccountDataResponseTrailerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataResponseTrailer) CopyPointer(to TradeAccountDataResponseTrailerPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataResponseTrailerBuilder) CopyPointer(to TradeAccountDataResponseTrailerPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m TradeAccountDataResponseTrailerPointer) Copy(to TradeAccountDataResponseTrailerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m TradeAccountDataResponseTrailerPointerBuilder) Copy(to TradeAccountDataResponseTrailerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataResponseTrailerPointer) CopyPointer(to TradeAccountDataResponseTrailerPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m TradeAccountDataResponseTrailerPointerBuilder) CopyPointer(to TradeAccountDataResponseTrailerPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetTradeAccount(m.TradeAccount())
}
