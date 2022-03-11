package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const TradeAccountDataAuthorizedUsernameRemoveSize = 18

// {
//     Size         = TradeAccountDataAuthorizedUsernameRemoveSize  (18)
//     Type         = TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE  (10126)
//     BaseSize     = TradeAccountDataAuthorizedUsernameRemoveSize  (18)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
// }
var _TradeAccountDataAuthorizedUsernameRemoveDefault = []byte{18, 0, 142, 39, 18, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataAuthorizedUsernameRemove struct {
	message.VLS
}

type TradeAccountDataAuthorizedUsernameRemoveBuilder struct {
	message.VLSBuilder
}

type TradeAccountDataAuthorizedUsernameRemovePointer struct {
	message.VLSPointer
}

type TradeAccountDataAuthorizedUsernameRemovePointerBuilder struct {
	message.VLSPointerBuilder
}

func NewTradeAccountDataAuthorizedUsernameRemoveFrom(b []byte) TradeAccountDataAuthorizedUsernameRemove {
	return TradeAccountDataAuthorizedUsernameRemove{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataAuthorizedUsernameRemove(b []byte) TradeAccountDataAuthorizedUsernameRemove {
	return TradeAccountDataAuthorizedUsernameRemove{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataAuthorizedUsernameRemove() TradeAccountDataAuthorizedUsernameRemoveBuilder {
	a := TradeAccountDataAuthorizedUsernameRemoveBuilder{message.NewVLSBuilder(18)}
	a.Unsafe().SetBytes(0, _TradeAccountDataAuthorizedUsernameRemoveDefault)
	return a
}

func AllocTradeAccountDataAuthorizedUsernameRemove() TradeAccountDataAuthorizedUsernameRemovePointerBuilder {
	a := TradeAccountDataAuthorizedUsernameRemovePointerBuilder{message.AllocVLSBuilder(18)}
	a.Ptr.SetBytes(0, _TradeAccountDataAuthorizedUsernameRemoveDefault)
	return a
}

func AllocTradeAccountDataAuthorizedUsernameRemoveFrom(b []byte) TradeAccountDataAuthorizedUsernameRemovePointer {
	return TradeAccountDataAuthorizedUsernameRemovePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = TradeAccountDataAuthorizedUsernameRemoveSize  (18)
//     Type         = TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE  (10126)
//     BaseSize     = TradeAccountDataAuthorizedUsernameRemoveSize  (18)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
// }
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataAuthorizedUsernameRemoveDefault)
}

// Clear
// {
//     Size         = TradeAccountDataAuthorizedUsernameRemoveSize  (18)
//     Type         = TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE  (10126)
//     BaseSize     = TradeAccountDataAuthorizedUsernameRemoveSize  (18)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
// }
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataAuthorizedUsernameRemoveDefault)
}

// ToBuilder
func (m TradeAccountDataAuthorizedUsernameRemove) ToBuilder() TradeAccountDataAuthorizedUsernameRemoveBuilder {
	return TradeAccountDataAuthorizedUsernameRemoveBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m TradeAccountDataAuthorizedUsernameRemovePointer) ToBuilder() TradeAccountDataAuthorizedUsernameRemovePointerBuilder {
	return TradeAccountDataAuthorizedUsernameRemovePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) Finish() TradeAccountDataAuthorizedUsernameRemove {
	return TradeAccountDataAuthorizedUsernameRemove{m.VLSBuilder.Finish()}
}

// Finish
func (m *TradeAccountDataAuthorizedUsernameRemovePointerBuilder) Finish() TradeAccountDataAuthorizedUsernameRemovePointer {
	return TradeAccountDataAuthorizedUsernameRemovePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameRemove) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameRemovePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameRemove) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameRemovePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Username
func (m TradeAccountDataAuthorizedUsernameRemove) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Username
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Username
func (m TradeAccountDataAuthorizedUsernameRemovePointer) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// Username
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// SetRequestID
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataAuthorizedUsernameRemoveBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *TradeAccountDataAuthorizedUsernameRemovePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataAuthorizedUsernameRemoveBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetUsername
func (m *TradeAccountDataAuthorizedUsernameRemovePointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// Copy
func (m TradeAccountDataAuthorizedUsernameRemove) Copy(to TradeAccountDataAuthorizedUsernameRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// Copy
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) Copy(to TradeAccountDataAuthorizedUsernameRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameRemove) CopyPointer(to TradeAccountDataAuthorizedUsernameRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameRemoveBuilder) CopyPointer(to TradeAccountDataAuthorizedUsernameRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// Copy
func (m TradeAccountDataAuthorizedUsernameRemovePointer) Copy(to TradeAccountDataAuthorizedUsernameRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// Copy
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) Copy(to TradeAccountDataAuthorizedUsernameRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameRemovePointer) CopyPointer(to TradeAccountDataAuthorizedUsernameRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}

// CopyPointer
func (m TradeAccountDataAuthorizedUsernameRemovePointerBuilder) CopyPointer(to TradeAccountDataAuthorizedUsernameRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
}
