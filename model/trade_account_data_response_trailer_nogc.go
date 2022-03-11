package model

import (
	"github.com/moontrade/dtc-go/message"
)

type TradeAccountDataResponseTrailerPointer struct {
	message.VLSPointer
}

type TradeAccountDataResponseTrailerPointerBuilder struct {
	message.VLSPointerBuilder
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
//     IsSnapshot            = 0
//     IsFirstMessageInBatch = 0
//     IsLastMessageInBatch  = 0
//     TradeAccount          = ""
// }
func (m TradeAccountDataResponseTrailerPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataResponseTrailerDefault)
}

// ToBuilder
func (m TradeAccountDataResponseTrailerPointer) ToBuilder() TradeAccountDataResponseTrailerPointerBuilder {
	return TradeAccountDataResponseTrailerPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *TradeAccountDataResponseTrailerPointerBuilder) Finish() TradeAccountDataResponseTrailerPointer {
	return TradeAccountDataResponseTrailerPointer{m.VLSPointerBuilder.Finish()}
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
func (m TradeAccountDataResponseTrailerPointer) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// IsSnapshot
func (m TradeAccountDataResponseTrailerPointerBuilder) IsSnapshot() uint8 {
	return message.Uint8VLS(m.Ptr, 11, 10)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponseTrailerPointer) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 12, 11)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponseTrailerPointerBuilder) IsFirstMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 12, 11)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponseTrailerPointer) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 13, 12)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponseTrailerPointerBuilder) IsLastMessageInBatch() uint8 {
	return message.Uint8VLS(m.Ptr, 13, 12)
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
func (m TradeAccountDataResponseTrailerPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetIsSnapshot
func (m TradeAccountDataResponseTrailerPointerBuilder) SetIsSnapshot(value uint8) {
	message.SetUint8VLS(m.Ptr, 11, 10, value)
}

// SetIsFirstMessageInBatch
func (m TradeAccountDataResponseTrailerPointerBuilder) SetIsFirstMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Ptr, 12, 11, value)
}

// SetIsLastMessageInBatch
func (m TradeAccountDataResponseTrailerPointerBuilder) SetIsLastMessageInBatch(value uint8) {
	message.SetUint8VLS(m.Ptr, 13, 12, value)
}

// SetTradeAccount
func (m *TradeAccountDataResponseTrailerPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 17, 13, value)
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
