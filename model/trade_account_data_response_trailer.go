package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                  = TradeAccountDataResponseTrailerSize  (17)
//     Type                  = TRADE_ACCOUNT_DATA_RESPONSE_TRAILER  (10130)
//     BaseSize              = TradeAccountDataResponseTrailerSize  (17)
//     RequestID             = 0
//     IsSnapshot            = 0
//     IsFirstMessageInBatch = 0
//     IsLastMessageInBatch  = 0
//     TradeAccount          = ""
// }
var _TradeAccountDataResponseTrailerDefault = []byte{17, 0, 146, 39, 17, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const TradeAccountDataResponseTrailerSize = 17

type TradeAccountDataResponseTrailer struct {
	message.VLS
}

type TradeAccountDataResponseTrailerBuilder struct {
	message.VLSBuilder
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

// Clear
// {
//     Size                  = TradeAccountDataResponseTrailerSize  (17)
//     Type                  = TRADE_ACCOUNT_DATA_RESPONSE_TRAILER  (10130)
//     BaseSize              = TradeAccountDataResponseTrailerSize  (17)
//     RequestID             = 0
//     IsSnapshot            = 0
//     IsFirstMessageInBatch = 0
//     IsLastMessageInBatch  = 0
//     TradeAccount          = ""
// }
func (m TradeAccountDataResponseTrailerBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataResponseTrailerDefault)
}

// ToBuilder
func (m TradeAccountDataResponseTrailer) ToBuilder() TradeAccountDataResponseTrailerBuilder {
	return TradeAccountDataResponseTrailerBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m TradeAccountDataResponseTrailerBuilder) Finish() TradeAccountDataResponseTrailer {
	return TradeAccountDataResponseTrailer{m.VLSBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataResponseTrailer) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataResponseTrailerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// IsSnapshot
func (m TradeAccountDataResponseTrailer) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// IsSnapshot
func (m TradeAccountDataResponseTrailerBuilder) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Unsafe(), 11, 10)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponseTrailer) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 12, 11)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponseTrailerBuilder) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 12, 11)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponseTrailer) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 13, 12)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponseTrailerBuilder) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Unsafe(), 13, 12)
}

// TradeAccount
func (m TradeAccountDataResponseTrailer) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 17, 13)
}

// TradeAccount
func (m TradeAccountDataResponseTrailerBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 17, 13)
}

// SetRequestID
func (m TradeAccountDataResponseTrailerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetIsSnapshot
func (m TradeAccountDataResponseTrailerBuilder) SetIsSnapshot(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 11, 10, value)
}

// SetIsFirstMessageInBatch
func (m TradeAccountDataResponseTrailerBuilder) SetIsFirstMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 12, 11, value)
}

// SetIsLastMessageInBatch
func (m TradeAccountDataResponseTrailerBuilder) SetIsLastMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 13, 12, value)
}

// SetTradeAccount
func (m *TradeAccountDataResponseTrailerBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 17, 13, value)
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
