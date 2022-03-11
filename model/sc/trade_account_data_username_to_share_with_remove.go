package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const TradeAccountDataUsernameToShareWithRemoveSize = 19

// {
//     Size         = TradeAccountDataUsernameToShareWithRemoveSize  (19)
//     Type         = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_REMOVE  (10129)
//     BaseSize     = TradeAccountDataUsernameToShareWithRemoveSize  (19)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
//     IsReadOnly   = false
// }
var _TradeAccountDataUsernameToShareWithRemoveDefault = []byte{19, 0, 145, 39, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataUsernameToShareWithRemove struct {
	message.VLS
}

type TradeAccountDataUsernameToShareWithRemoveBuilder struct {
	message.VLSBuilder
}

type TradeAccountDataUsernameToShareWithRemovePointer struct {
	message.VLSPointer
}

type TradeAccountDataUsernameToShareWithRemovePointerBuilder struct {
	message.VLSPointerBuilder
}

func NewTradeAccountDataUsernameToShareWithRemoveFrom(b []byte) TradeAccountDataUsernameToShareWithRemove {
	return TradeAccountDataUsernameToShareWithRemove{VLS: message.NewVLSFrom(b)}
}

func WrapTradeAccountDataUsernameToShareWithRemove(b []byte) TradeAccountDataUsernameToShareWithRemove {
	return TradeAccountDataUsernameToShareWithRemove{VLS: message.WrapVLS(b)}
}

func NewTradeAccountDataUsernameToShareWithRemove() TradeAccountDataUsernameToShareWithRemoveBuilder {
	a := TradeAccountDataUsernameToShareWithRemoveBuilder{message.NewVLSBuilder(19)}
	a.Unsafe().SetBytes(0, _TradeAccountDataUsernameToShareWithRemoveDefault)
	return a
}

func AllocTradeAccountDataUsernameToShareWithRemove() TradeAccountDataUsernameToShareWithRemovePointerBuilder {
	a := TradeAccountDataUsernameToShareWithRemovePointerBuilder{message.AllocVLSBuilder(19)}
	a.Ptr.SetBytes(0, _TradeAccountDataUsernameToShareWithRemoveDefault)
	return a
}

func AllocTradeAccountDataUsernameToShareWithRemoveFrom(b []byte) TradeAccountDataUsernameToShareWithRemovePointer {
	return TradeAccountDataUsernameToShareWithRemovePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = TradeAccountDataUsernameToShareWithRemoveSize  (19)
//     Type         = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_REMOVE  (10129)
//     BaseSize     = TradeAccountDataUsernameToShareWithRemoveSize  (19)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
//     IsReadOnly   = false
// }
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) Clear() {
	m.Unsafe().SetBytes(0, _TradeAccountDataUsernameToShareWithRemoveDefault)
}

// Clear
// {
//     Size         = TradeAccountDataUsernameToShareWithRemoveSize  (19)
//     Type         = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_REMOVE  (10129)
//     BaseSize     = TradeAccountDataUsernameToShareWithRemoveSize  (19)
//     RequestID    = 0
//     TradeAccount = ""
//     Username     = ""
//     IsReadOnly   = false
// }
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _TradeAccountDataUsernameToShareWithRemoveDefault)
}

// ToBuilder
func (m TradeAccountDataUsernameToShareWithRemove) ToBuilder() TradeAccountDataUsernameToShareWithRemoveBuilder {
	return TradeAccountDataUsernameToShareWithRemoveBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m TradeAccountDataUsernameToShareWithRemovePointer) ToBuilder() TradeAccountDataUsernameToShareWithRemovePointerBuilder {
	return TradeAccountDataUsernameToShareWithRemovePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) Finish() TradeAccountDataUsernameToShareWithRemove {
	return TradeAccountDataUsernameToShareWithRemove{m.VLSBuilder.Finish()}
}

// Finish
func (m *TradeAccountDataUsernameToShareWithRemovePointerBuilder) Finish() TradeAccountDataUsernameToShareWithRemovePointer {
	return TradeAccountDataUsernameToShareWithRemovePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m TradeAccountDataUsernameToShareWithRemove) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m TradeAccountDataUsernameToShareWithRemovePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithRemove) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithRemovePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Username
func (m TradeAccountDataUsernameToShareWithRemove) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Username
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// Username
func (m TradeAccountDataUsernameToShareWithRemovePointer) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// Username
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithRemove) IsReadOnly() bool {
	return message.BoolVLS(m.Unsafe(), 19, 18)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) IsReadOnly() bool {
	return message.BoolVLS(m.Unsafe(), 19, 18)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithRemovePointer) IsReadOnly() bool {
	return message.BoolVLS(m.Ptr, 19, 18)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) IsReadOnly() bool {
	return message.BoolVLS(m.Ptr, 19, 18)
}

// SetRequestID
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetTradeAccount
func (m *TradeAccountDataUsernameToShareWithRemoveBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *TradeAccountDataUsernameToShareWithRemovePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetUsername
func (m *TradeAccountDataUsernameToShareWithRemoveBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetUsername
func (m *TradeAccountDataUsernameToShareWithRemovePointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetIsReadOnly
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) SetIsReadOnly(value bool) {
	message.SetBoolVLS(m.Unsafe(), 19, 18, value)
}

// SetIsReadOnly
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) SetIsReadOnly(value bool) {
	message.SetBoolVLS(m.Ptr, 19, 18, value)
}

// Copy
func (m TradeAccountDataUsernameToShareWithRemove) Copy(to TradeAccountDataUsernameToShareWithRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// Copy
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) Copy(to TradeAccountDataUsernameToShareWithRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithRemove) CopyPointer(to TradeAccountDataUsernameToShareWithRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithRemoveBuilder) CopyPointer(to TradeAccountDataUsernameToShareWithRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// Copy
func (m TradeAccountDataUsernameToShareWithRemovePointer) Copy(to TradeAccountDataUsernameToShareWithRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// Copy
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) Copy(to TradeAccountDataUsernameToShareWithRemoveBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithRemovePointer) CopyPointer(to TradeAccountDataUsernameToShareWithRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

// CopyPointer
func (m TradeAccountDataUsernameToShareWithRemovePointerBuilder) CopyPointer(to TradeAccountDataUsernameToShareWithRemovePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}