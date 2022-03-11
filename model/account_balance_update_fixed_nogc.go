package model

import (
	"github.com/moontrade/dtc-go/message"
)

// AccountBalanceUpdateFixedPointer This is an optional message from the Server to Client to provide Account
// Balance information for a particular Trade Account. The server needs to
// provide a separate message for each Trade Account associated with the
// logged in username if it supports Account Balance updates.
//
// The Server will respond with an AccountBalanceUpdate in response to a
// AccountBalanceRequest message. The Server will set the RequestID in the
// AccountBalanceUpdate message to match the RequestID in the AccountBalanceRequest
// message.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When an unsolicited AccountBalanceUpdate message is sent, the RequestID
// field will be 0.
//
// When the server is responding with one or more AccountBalanceUpdate messages
// in response to a AccountBalanceRequest message, it must not send any unsolicited
// AccountBalanceUpdate messages interleaved with the solicited AccountBalanceUpdate
// messages in response to the AccountBalanceRequest message.
type AccountBalanceUpdateFixedPointer struct {
	message.FixedPointer
}

// AccountBalanceUpdateFixedPointerBuilder This is an optional message from the Server to Client to provide Account
// Balance information for a particular Trade Account. The server needs to
// provide a separate message for each Trade Account associated with the
// logged in username if it supports Account Balance updates.
//
// The Server will respond with an AccountBalanceUpdate in response to a
// AccountBalanceRequest message. The Server will set the RequestID in the
// AccountBalanceUpdate message to match the RequestID in the AccountBalanceRequest
// message.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When an unsolicited AccountBalanceUpdate message is sent, the RequestID
// field will be 0.
//
// When the server is responding with one or more AccountBalanceUpdate messages
// in response to a AccountBalanceRequest message, it must not send any unsolicited
// AccountBalanceUpdate messages interleaved with the solicited AccountBalanceUpdate
// messages in response to the AccountBalanceRequest message.
type AccountBalanceUpdateFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocAccountBalanceUpdateFixed() AccountBalanceUpdateFixedPointerBuilder {
	a := AccountBalanceUpdateFixedPointerBuilder{message.AllocFixed(368)}
	a.Ptr.SetBytes(0, _AccountBalanceUpdateFixedDefault)
	return a
}

func AllocAccountBalanceUpdateFixedFrom(b []byte) AccountBalanceUpdateFixedPointer {
	return AccountBalanceUpdateFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                                 = AccountBalanceUpdateFixedSize  (368)
//     Type                                 = ACCOUNT_BALANCE_UPDATE  (600)
//     RequestID                            = 0
//     CashBalance                          = 0
//     BalanceAvailableForNewPositions      = 0
//     AccountCurrency                      = ""
//     TradeAccount                         = ""
//     SecuritiesValue                      = 0
//     MarginRequirement                    = 0
//     TotalNumberMessages                  = 0
//     MessageNumber                        = 0
//     NoAccountBalances                    = 0
//     Unsolicited                          = 0
//     OpenPositionsProfitLoss              = 0
//     DailyProfitLoss                      = 0
//     InfoText                             = ""
//     TransactionIdentifier                = 0
//     DailyNetLossLimit                    = 0
//     TrailingAccountValueToLimitPositions = 0
//     DailyNetLossLimitReached             = 0
//     IsUnderRequiredMargin                = 0
//     ClosePositionsAtEndOfDay             = 0
//     TradingIsDisabled                    = 0
//     Description                          = ""
//     IsUnderRequiredAccountValue          = 0
//     TransactionDateTime                  = 0
//     MarginRequirementFull                = 0
//     MarginRequirementFullPositionsOnly   = 0
//     PeakMarginRequirement                = 0
// }
func (m AccountBalanceUpdateFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceUpdateFixedDefault)
}

// ToBuilder
func (m AccountBalanceUpdateFixedPointer) ToBuilder() AccountBalanceUpdateFixedPointerBuilder {
	return AccountBalanceUpdateFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *AccountBalanceUpdateFixedPointerBuilder) Finish() AccountBalanceUpdateFixedPointer {
	return AccountBalanceUpdateFixedPointer{m.FixedPointer}
}

// RequestID This is the RequestID which was set in the AccountBalanceRequest that
// this message is in response to.
//
// In the case when this is a periodic unsolicited Account Balance update,
// RequestID must be set to 0, the default.
func (m AccountBalanceUpdateFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID This is the RequestID which was set in the AccountBalanceRequest that
// this message is in response to.
//
// In the case when this is a periodic unsolicited Account Balance update,
// RequestID must be set to 0, the default.
func (m AccountBalanceUpdateFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// CashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdateFixedPointer) CashBalance() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// CashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdateFixedPointerBuilder) CashBalance() float64 {
	return message.Float64Fixed(m.Ptr, 16, 8)
}

// BalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedPointer) BalanceAvailableForNewPositions() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// BalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedPointerBuilder) BalanceAvailableForNewPositions() float64 {
	return message.Float64Fixed(m.Ptr, 24, 16)
}

// AccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdateFixedPointer) AccountCurrency() string {
	return message.StringFixed(m.Ptr, 32, 24)
}

// AccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdateFixedPointerBuilder) AccountCurrency() string {
	return message.StringFixed(m.Ptr, 32, 24)
}

// TradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdateFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 64, 32)
}

// TradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdateFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 64, 32)
}

// SecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdateFixedPointer) SecuritiesValue() float64 {
	return message.Float64Fixed(m.Ptr, 72, 64)
}

// SecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdateFixedPointerBuilder) SecuritiesValue() float64 {
	return message.Float64Fixed(m.Ptr, 72, 64)
}

// MarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedPointer) MarginRequirement() float64 {
	return message.Float64Fixed(m.Ptr, 80, 72)
}

// MarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedPointerBuilder) MarginRequirement() float64 {
	return message.Float64Fixed(m.Ptr, 80, 72)
}

// TotalNumberMessages This indicates the total number of Account Balance Update messages when
// a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateFixedPointer) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Ptr, 84, 80)
}

// TotalNumberMessages This indicates the total number of Account Balance Update messages when
// a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateFixedPointerBuilder) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Ptr, 84, 80)
}

// MessageNumber This indicates the 1-based index of the Account Balance Update message
// when a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateFixedPointer) MessageNumber() int32 {
	return message.Int32Fixed(m.Ptr, 88, 84)
}

// MessageNumber This indicates the 1-based index of the Account Balance Update message
// when a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateFixedPointerBuilder) MessageNumber() int32 {
	return message.Int32Fixed(m.Ptr, 88, 84)
}

// NoAccountBalances Set to an integer value of 1 to indicate there are no Account Balances
// in response to an AccountBalanceRequest message.
//
// Otherwise, the Server must leave this at the default of 0.
//
// The Server is always expected to respond with a single AccountBalanceUpdate
// message when there is no AccountBalanceUpdate message for any Trade Account
// when the Account Balances have been requested by the Client with an AccountBalanceRequest
// message.
//
// When the Server is sending an AccountBalanceUpdate message to the Client
// and it is indicating that the balance related fields are all zero, then
// the NoAccountBalances field must be left at the default of 0. It is not
// used indicate the balance related fields are all zero.
//
// This is always set to the default of 0 for an unsolicited AccountBalanceUpdate
// .
func (m AccountBalanceUpdateFixedPointer) NoAccountBalances() uint8 {
	return message.Uint8Fixed(m.Ptr, 89, 88)
}

// NoAccountBalances Set to an integer value of 1 to indicate there are no Account Balances
// in response to an AccountBalanceRequest message.
//
// Otherwise, the Server must leave this at the default of 0.
//
// The Server is always expected to respond with a single AccountBalanceUpdate
// message when there is no AccountBalanceUpdate message for any Trade Account
// when the Account Balances have been requested by the Client with an AccountBalanceRequest
// message.
//
// When the Server is sending an AccountBalanceUpdate message to the Client
// and it is indicating that the balance related fields are all zero, then
// the NoAccountBalances field must be left at the default of 0. It is not
// used indicate the balance related fields are all zero.
//
// This is always set to the default of 0 for an unsolicited AccountBalanceUpdate
// .
func (m AccountBalanceUpdateFixedPointerBuilder) NoAccountBalances() uint8 {
	return message.Uint8Fixed(m.Ptr, 89, 88)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdateFixedPointer) Unsolicited() uint8 {
	return message.Uint8Fixed(m.Ptr, 90, 89)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdateFixedPointerBuilder) Unsolicited() uint8 {
	return message.Uint8Fixed(m.Ptr, 90, 89)
}

// OpenPositionsProfitLoss
func (m AccountBalanceUpdateFixedPointer) OpenPositionsProfitLoss() float64 {
	return message.Float64Fixed(m.Ptr, 104, 96)
}

// OpenPositionsProfitLoss
func (m AccountBalanceUpdateFixedPointerBuilder) OpenPositionsProfitLoss() float64 {
	return message.Float64Fixed(m.Ptr, 104, 96)
}

// DailyProfitLoss
func (m AccountBalanceUpdateFixedPointer) DailyProfitLoss() float64 {
	return message.Float64Fixed(m.Ptr, 112, 104)
}

// DailyProfitLoss
func (m AccountBalanceUpdateFixedPointerBuilder) DailyProfitLoss() float64 {
	return message.Float64Fixed(m.Ptr, 112, 104)
}

// InfoText
func (m AccountBalanceUpdateFixedPointer) InfoText() string {
	return message.StringFixed(m.Ptr, 208, 112)
}

// InfoText
func (m AccountBalanceUpdateFixedPointerBuilder) InfoText() string {
	return message.StringFixed(m.Ptr, 208, 112)
}

// TransactionIdentifier
func (m AccountBalanceUpdateFixedPointer) TransactionIdentifier() uint64 {
	return message.Uint64Fixed(m.Ptr, 216, 208)
}

// TransactionIdentifier
func (m AccountBalanceUpdateFixedPointerBuilder) TransactionIdentifier() uint64 {
	return message.Uint64Fixed(m.Ptr, 216, 208)
}

// DailyNetLossLimit
func (m AccountBalanceUpdateFixedPointer) DailyNetLossLimit() float64 {
	return message.Float64Fixed(m.Ptr, 224, 216)
}

// DailyNetLossLimit
func (m AccountBalanceUpdateFixedPointerBuilder) DailyNetLossLimit() float64 {
	return message.Float64Fixed(m.Ptr, 224, 216)
}

// TrailingAccountValueToLimitPositions
func (m AccountBalanceUpdateFixedPointer) TrailingAccountValueToLimitPositions() float64 {
	return message.Float64Fixed(m.Ptr, 232, 224)
}

// TrailingAccountValueToLimitPositions
func (m AccountBalanceUpdateFixedPointerBuilder) TrailingAccountValueToLimitPositions() float64 {
	return message.Float64Fixed(m.Ptr, 232, 224)
}

// DailyNetLossLimitReached
func (m AccountBalanceUpdateFixedPointer) DailyNetLossLimitReached() uint8 {
	return message.Uint8Fixed(m.Ptr, 233, 232)
}

// DailyNetLossLimitReached
func (m AccountBalanceUpdateFixedPointerBuilder) DailyNetLossLimitReached() uint8 {
	return message.Uint8Fixed(m.Ptr, 233, 232)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdateFixedPointer) IsUnderRequiredMargin() uint8 {
	return message.Uint8Fixed(m.Ptr, 234, 233)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdateFixedPointerBuilder) IsUnderRequiredMargin() uint8 {
	return message.Uint8Fixed(m.Ptr, 234, 233)
}

// ClosePositionsAtEndOfDay
func (m AccountBalanceUpdateFixedPointer) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8Fixed(m.Ptr, 235, 234)
}

// ClosePositionsAtEndOfDay
func (m AccountBalanceUpdateFixedPointerBuilder) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8Fixed(m.Ptr, 235, 234)
}

// TradingIsDisabled
func (m AccountBalanceUpdateFixedPointer) TradingIsDisabled() uint8 {
	return message.Uint8Fixed(m.Ptr, 236, 235)
}

// TradingIsDisabled
func (m AccountBalanceUpdateFixedPointerBuilder) TradingIsDisabled() uint8 {
	return message.Uint8Fixed(m.Ptr, 236, 235)
}

// Description
func (m AccountBalanceUpdateFixedPointer) Description() string {
	return message.StringFixed(m.Ptr, 332, 236)
}

// Description
func (m AccountBalanceUpdateFixedPointerBuilder) Description() string {
	return message.StringFixed(m.Ptr, 332, 236)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdateFixedPointer) IsUnderRequiredAccountValue() uint8 {
	return message.Uint8Fixed(m.Ptr, 333, 332)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdateFixedPointerBuilder) IsUnderRequiredAccountValue() uint8 {
	return message.Uint8Fixed(m.Ptr, 333, 332)
}

// TransactionDateTime
func (m AccountBalanceUpdateFixedPointer) TransactionDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 344, 336))
}

// TransactionDateTime
func (m AccountBalanceUpdateFixedPointerBuilder) TransactionDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 344, 336))
}

// MarginRequirementFull
func (m AccountBalanceUpdateFixedPointer) MarginRequirementFull() float64 {
	return message.Float64Fixed(m.Ptr, 352, 344)
}

// MarginRequirementFull
func (m AccountBalanceUpdateFixedPointerBuilder) MarginRequirementFull() float64 {
	return message.Float64Fixed(m.Ptr, 352, 344)
}

// MarginRequirementFullPositionsOnly
func (m AccountBalanceUpdateFixedPointer) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64Fixed(m.Ptr, 360, 352)
}

// MarginRequirementFullPositionsOnly
func (m AccountBalanceUpdateFixedPointerBuilder) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64Fixed(m.Ptr, 360, 352)
}

// PeakMarginRequirement
func (m AccountBalanceUpdateFixedPointer) PeakMarginRequirement() float64 {
	return message.Float64Fixed(m.Ptr, 368, 360)
}

// PeakMarginRequirement
func (m AccountBalanceUpdateFixedPointerBuilder) PeakMarginRequirement() float64 {
	return message.Float64Fixed(m.Ptr, 368, 360)
}

// SetRequestID This is the RequestID which was set in the AccountBalanceRequest that
// this message is in response to.
//
// In the case when this is a periodic unsolicited Account Balance update,
// RequestID must be set to 0, the default.
func (m AccountBalanceUpdateFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetCashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdateFixedPointerBuilder) SetCashBalance(value float64) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, value)
}

// SetBalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedPointerBuilder) SetBalanceAvailableForNewPositions(value float64) {
	message.SetFloat64Fixed(m.Ptr, 24, 16, value)
}

// SetAccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdateFixedPointerBuilder) SetAccountCurrency(value string) {
	message.SetStringFixed(m.Ptr, 32, 24, value)
}

// SetTradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdateFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 64, 32, value)
}

// SetSecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdateFixedPointerBuilder) SetSecuritiesValue(value float64) {
	message.SetFloat64Fixed(m.Ptr, 72, 64, value)
}

// SetMarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedPointerBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64Fixed(m.Ptr, 80, 72, value)
}

// SetTotalNumberMessages This indicates the total number of Account Balance Update messages when
// a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateFixedPointerBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32Fixed(m.Ptr, 84, 80, value)
}

// SetMessageNumber This indicates the 1-based index of the Account Balance Update message
// when a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateFixedPointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Ptr, 88, 84, value)
}

// SetNoAccountBalances Set to an integer value of 1 to indicate there are no Account Balances
// in response to an AccountBalanceRequest message.
//
// Otherwise, the Server must leave this at the default of 0.
//
// The Server is always expected to respond with a single AccountBalanceUpdate
// message when there is no AccountBalanceUpdate message for any Trade Account
// when the Account Balances have been requested by the Client with an AccountBalanceRequest
// message.
//
// When the Server is sending an AccountBalanceUpdate message to the Client
// and it is indicating that the balance related fields are all zero, then
// the NoAccountBalances field must be left at the default of 0. It is not
// used indicate the balance related fields are all zero.
//
// This is always set to the default of 0 for an unsolicited AccountBalanceUpdate
// .
func (m AccountBalanceUpdateFixedPointerBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8Fixed(m.Ptr, 89, 88, value)
}

// SetUnsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdateFixedPointerBuilder) SetUnsolicited(value uint8) {
	message.SetUint8Fixed(m.Ptr, 90, 89, value)
}

// SetOpenPositionsProfitLoss
func (m AccountBalanceUpdateFixedPointerBuilder) SetOpenPositionsProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Ptr, 104, 96, value)
}

// SetDailyProfitLoss
func (m AccountBalanceUpdateFixedPointerBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Ptr, 112, 104, value)
}

// SetInfoText
func (m AccountBalanceUpdateFixedPointerBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Ptr, 208, 112, value)
}

// SetTransactionIdentifier
func (m AccountBalanceUpdateFixedPointerBuilder) SetTransactionIdentifier(value uint64) {
	message.SetUint64Fixed(m.Ptr, 216, 208, value)
}

// SetDailyNetLossLimit
func (m AccountBalanceUpdateFixedPointerBuilder) SetDailyNetLossLimit(value float64) {
	message.SetFloat64Fixed(m.Ptr, 224, 216, value)
}

// SetTrailingAccountValueToLimitPositions
func (m AccountBalanceUpdateFixedPointerBuilder) SetTrailingAccountValueToLimitPositions(value float64) {
	message.SetFloat64Fixed(m.Ptr, 232, 224, value)
}

// SetDailyNetLossLimitReached
func (m AccountBalanceUpdateFixedPointerBuilder) SetDailyNetLossLimitReached(value uint8) {
	message.SetUint8Fixed(m.Ptr, 233, 232, value)
}

// SetIsUnderRequiredMargin
func (m AccountBalanceUpdateFixedPointerBuilder) SetIsUnderRequiredMargin(value uint8) {
	message.SetUint8Fixed(m.Ptr, 234, 233, value)
}

// SetClosePositionsAtEndOfDay
func (m AccountBalanceUpdateFixedPointerBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8Fixed(m.Ptr, 235, 234, value)
}

// SetTradingIsDisabled
func (m AccountBalanceUpdateFixedPointerBuilder) SetTradingIsDisabled(value uint8) {
	message.SetUint8Fixed(m.Ptr, 236, 235, value)
}

// SetDescription
func (m AccountBalanceUpdateFixedPointerBuilder) SetDescription(value string) {
	message.SetStringFixed(m.Ptr, 332, 236, value)
}

// SetIsUnderRequiredAccountValue
func (m AccountBalanceUpdateFixedPointerBuilder) SetIsUnderRequiredAccountValue(value uint8) {
	message.SetUint8Fixed(m.Ptr, 333, 332, value)
}

// SetTransactionDateTime
func (m AccountBalanceUpdateFixedPointerBuilder) SetTransactionDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 344, 336, int64(value))
}

// SetMarginRequirementFull
func (m AccountBalanceUpdateFixedPointerBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64Fixed(m.Ptr, 352, 344, value)
}

// SetMarginRequirementFullPositionsOnly
func (m AccountBalanceUpdateFixedPointerBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64Fixed(m.Ptr, 360, 352, value)
}

// SetPeakMarginRequirement
func (m AccountBalanceUpdateFixedPointerBuilder) SetPeakMarginRequirement(value float64) {
	message.SetFloat64Fixed(m.Ptr, 368, 360, value)
}

// Copy
func (m AccountBalanceUpdateFixedPointer) Copy(to AccountBalanceUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetCashBalance(m.CashBalance())
	to.SetBalanceAvailableForNewPositions(m.BalanceAvailableForNewPositions())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSecuritiesValue(m.SecuritiesValue())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetUnsolicited(m.Unsolicited())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetInfoText(m.InfoText())
	to.SetTransactionIdentifier(m.TransactionIdentifier())
	to.SetDailyNetLossLimit(m.DailyNetLossLimit())
	to.SetTrailingAccountValueToLimitPositions(m.TrailingAccountValueToLimitPositions())
	to.SetDailyNetLossLimitReached(m.DailyNetLossLimitReached())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescription(m.Description())
	to.SetIsUnderRequiredAccountValue(m.IsUnderRequiredAccountValue())
	to.SetTransactionDateTime(m.TransactionDateTime())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
}

// Copy
func (m AccountBalanceUpdateFixedPointerBuilder) Copy(to AccountBalanceUpdateFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetCashBalance(m.CashBalance())
	to.SetBalanceAvailableForNewPositions(m.BalanceAvailableForNewPositions())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSecuritiesValue(m.SecuritiesValue())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetUnsolicited(m.Unsolicited())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetInfoText(m.InfoText())
	to.SetTransactionIdentifier(m.TransactionIdentifier())
	to.SetDailyNetLossLimit(m.DailyNetLossLimit())
	to.SetTrailingAccountValueToLimitPositions(m.TrailingAccountValueToLimitPositions())
	to.SetDailyNetLossLimitReached(m.DailyNetLossLimitReached())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescription(m.Description())
	to.SetIsUnderRequiredAccountValue(m.IsUnderRequiredAccountValue())
	to.SetTransactionDateTime(m.TransactionDateTime())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
}

// CopyPointer
func (m AccountBalanceUpdateFixedPointer) CopyPointer(to AccountBalanceUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetCashBalance(m.CashBalance())
	to.SetBalanceAvailableForNewPositions(m.BalanceAvailableForNewPositions())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSecuritiesValue(m.SecuritiesValue())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetUnsolicited(m.Unsolicited())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetInfoText(m.InfoText())
	to.SetTransactionIdentifier(m.TransactionIdentifier())
	to.SetDailyNetLossLimit(m.DailyNetLossLimit())
	to.SetTrailingAccountValueToLimitPositions(m.TrailingAccountValueToLimitPositions())
	to.SetDailyNetLossLimitReached(m.DailyNetLossLimitReached())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescription(m.Description())
	to.SetIsUnderRequiredAccountValue(m.IsUnderRequiredAccountValue())
	to.SetTransactionDateTime(m.TransactionDateTime())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
}

// CopyPointer
func (m AccountBalanceUpdateFixedPointerBuilder) CopyPointer(to AccountBalanceUpdateFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetCashBalance(m.CashBalance())
	to.SetBalanceAvailableForNewPositions(m.BalanceAvailableForNewPositions())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSecuritiesValue(m.SecuritiesValue())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetUnsolicited(m.Unsolicited())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetInfoText(m.InfoText())
	to.SetTransactionIdentifier(m.TransactionIdentifier())
	to.SetDailyNetLossLimit(m.DailyNetLossLimit())
	to.SetTrailingAccountValueToLimitPositions(m.TrailingAccountValueToLimitPositions())
	to.SetDailyNetLossLimitReached(m.DailyNetLossLimitReached())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescription(m.Description())
	to.SetIsUnderRequiredAccountValue(m.IsUnderRequiredAccountValue())
	to.SetTransactionDateTime(m.TransactionDateTime())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
}

// CopyTo
func (m AccountBalanceUpdateFixedPointer) CopyTo(to AccountBalanceUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetCashBalance(m.CashBalance())
	to.SetBalanceAvailableForNewPositions(m.BalanceAvailableForNewPositions())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSecuritiesValue(m.SecuritiesValue())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetUnsolicited(m.Unsolicited())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetInfoText(m.InfoText())
	to.SetTransactionIdentifier(m.TransactionIdentifier())
	to.SetDailyNetLossLimit(m.DailyNetLossLimit())
	to.SetTrailingAccountValueToLimitPositions(m.TrailingAccountValueToLimitPositions())
	to.SetDailyNetLossLimitReached(m.DailyNetLossLimitReached())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescription(m.Description())
	to.SetIsUnderRequiredAccountValue(m.IsUnderRequiredAccountValue())
	to.SetTransactionDateTime(m.TransactionDateTime())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
}

// CopyTo
func (m AccountBalanceUpdateFixedPointerBuilder) CopyTo(to AccountBalanceUpdateBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetCashBalance(m.CashBalance())
	to.SetBalanceAvailableForNewPositions(m.BalanceAvailableForNewPositions())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSecuritiesValue(m.SecuritiesValue())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetUnsolicited(m.Unsolicited())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetInfoText(m.InfoText())
	to.SetTransactionIdentifier(m.TransactionIdentifier())
	to.SetDailyNetLossLimit(m.DailyNetLossLimit())
	to.SetTrailingAccountValueToLimitPositions(m.TrailingAccountValueToLimitPositions())
	to.SetDailyNetLossLimitReached(m.DailyNetLossLimitReached())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescription(m.Description())
	to.SetIsUnderRequiredAccountValue(m.IsUnderRequiredAccountValue())
	to.SetTransactionDateTime(m.TransactionDateTime())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
}

// CopyToPointer
func (m AccountBalanceUpdateFixedPointer) CopyToPointer(to AccountBalanceUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetCashBalance(m.CashBalance())
	to.SetBalanceAvailableForNewPositions(m.BalanceAvailableForNewPositions())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSecuritiesValue(m.SecuritiesValue())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetUnsolicited(m.Unsolicited())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetInfoText(m.InfoText())
	to.SetTransactionIdentifier(m.TransactionIdentifier())
	to.SetDailyNetLossLimit(m.DailyNetLossLimit())
	to.SetTrailingAccountValueToLimitPositions(m.TrailingAccountValueToLimitPositions())
	to.SetDailyNetLossLimitReached(m.DailyNetLossLimitReached())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescription(m.Description())
	to.SetIsUnderRequiredAccountValue(m.IsUnderRequiredAccountValue())
	to.SetTransactionDateTime(m.TransactionDateTime())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
}

// CopyToPointer
func (m AccountBalanceUpdateFixedPointerBuilder) CopyToPointer(to AccountBalanceUpdatePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetCashBalance(m.CashBalance())
	to.SetBalanceAvailableForNewPositions(m.BalanceAvailableForNewPositions())
	to.SetAccountCurrency(m.AccountCurrency())
	to.SetTradeAccount(m.TradeAccount())
	to.SetSecuritiesValue(m.SecuritiesValue())
	to.SetMarginRequirement(m.MarginRequirement())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetNoAccountBalances(m.NoAccountBalances())
	to.SetUnsolicited(m.Unsolicited())
	to.SetOpenPositionsProfitLoss(m.OpenPositionsProfitLoss())
	to.SetDailyProfitLoss(m.DailyProfitLoss())
	to.SetInfoText(m.InfoText())
	to.SetTransactionIdentifier(m.TransactionIdentifier())
	to.SetDailyNetLossLimit(m.DailyNetLossLimit())
	to.SetTrailingAccountValueToLimitPositions(m.TrailingAccountValueToLimitPositions())
	to.SetDailyNetLossLimitReached(m.DailyNetLossLimitReached())
	to.SetIsUnderRequiredMargin(m.IsUnderRequiredMargin())
	to.SetClosePositionsAtEndOfDay(m.ClosePositionsAtEndOfDay())
	to.SetTradingIsDisabled(m.TradingIsDisabled())
	to.SetDescription(m.Description())
	to.SetIsUnderRequiredAccountValue(m.IsUnderRequiredAccountValue())
	to.SetTransactionDateTime(m.TransactionDateTime())
	to.SetMarginRequirementFull(m.MarginRequirementFull())
	to.SetMarginRequirementFullPositionsOnly(m.MarginRequirementFullPositionsOnly())
	to.SetPeakMarginRequirement(m.PeakMarginRequirement())
}
