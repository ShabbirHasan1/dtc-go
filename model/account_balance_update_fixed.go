package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _AccountBalanceUpdateFixedDefault = []byte{112, 1, 88, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AccountBalanceUpdateFixedSize = 368

// AccountBalanceUpdateFixed This is an optional message from the Server to Client to provide Account
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
type AccountBalanceUpdateFixed struct {
	message.Fixed
}

// AccountBalanceUpdateFixedBuilder This is an optional message from the Server to Client to provide Account
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
type AccountBalanceUpdateFixedBuilder struct {
	message.Fixed
}

func NewAccountBalanceUpdateFixedFrom(b []byte) AccountBalanceUpdateFixed {
	return AccountBalanceUpdateFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapAccountBalanceUpdateFixed(b []byte) AccountBalanceUpdateFixed {
	return AccountBalanceUpdateFixed{Fixed: message.WrapFixed(b)}
}

func NewAccountBalanceUpdateFixed() AccountBalanceUpdateFixedBuilder {
	a := AccountBalanceUpdateFixedBuilder{message.NewFixed(368)}
	a.Unsafe().SetBytes(0, _AccountBalanceUpdateFixedDefault)
	return a
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
func (m AccountBalanceUpdateFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceUpdateFixedDefault)
}

// ToBuilder
func (m AccountBalanceUpdateFixed) ToBuilder() AccountBalanceUpdateFixedBuilder {
	return AccountBalanceUpdateFixedBuilder{m.Fixed}
}

// Finish
func (m AccountBalanceUpdateFixedBuilder) Finish() AccountBalanceUpdateFixed {
	return AccountBalanceUpdateFixed{m.Fixed}
}

// RequestID This is the RequestID which was set in the AccountBalanceRequest that
// this message is in response to.
//
// In the case when this is a periodic unsolicited Account Balance update,
// RequestID must be set to 0, the default.
func (m AccountBalanceUpdateFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID This is the RequestID which was set in the AccountBalanceRequest that
// this message is in response to.
//
// In the case when this is a periodic unsolicited Account Balance update,
// RequestID must be set to 0, the default.
func (m AccountBalanceUpdateFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// CashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdateFixed) CashBalance() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// CashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdateFixedBuilder) CashBalance() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// BalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixed) BalanceAvailableForNewPositions() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// BalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedBuilder) BalanceAvailableForNewPositions() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// AccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdateFixed) AccountCurrency() string {
	return message.StringFixed(m.Unsafe(), 32, 24)
}

// AccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdateFixedBuilder) AccountCurrency() string {
	return message.StringFixed(m.Unsafe(), 32, 24)
}

// TradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdateFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 64, 32)
}

// TradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdateFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 64, 32)
}

// SecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdateFixed) SecuritiesValue() float64 {
	return message.Float64Fixed(m.Unsafe(), 72, 64)
}

// SecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdateFixedBuilder) SecuritiesValue() float64 {
	return message.Float64Fixed(m.Unsafe(), 72, 64)
}

// MarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixed) MarginRequirement() float64 {
	return message.Float64Fixed(m.Unsafe(), 80, 72)
}

// MarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedBuilder) MarginRequirement() float64 {
	return message.Float64Fixed(m.Unsafe(), 80, 72)
}

// TotalNumberMessages This indicates the total number of Account Balance Update messages when
// a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateFixed) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Unsafe(), 84, 80)
}

// TotalNumberMessages This indicates the total number of Account Balance Update messages when
// a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateFixedBuilder) TotalNumberMessages() int32 {
	return message.Int32Fixed(m.Unsafe(), 84, 80)
}

// MessageNumber This indicates the 1-based index of the Account Balance Update message
// when a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateFixed) MessageNumber() int32 {
	return message.Int32Fixed(m.Unsafe(), 88, 84)
}

// MessageNumber This indicates the 1-based index of the Account Balance Update message
// when a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateFixedBuilder) MessageNumber() int32 {
	return message.Int32Fixed(m.Unsafe(), 88, 84)
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
func (m AccountBalanceUpdateFixed) NoAccountBalances() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 89, 88)
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
func (m AccountBalanceUpdateFixedBuilder) NoAccountBalances() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 89, 88)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdateFixed) Unsolicited() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 90, 89)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdateFixedBuilder) Unsolicited() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 90, 89)
}

// OpenPositionsProfitLoss
func (m AccountBalanceUpdateFixed) OpenPositionsProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 104, 96)
}

// OpenPositionsProfitLoss
func (m AccountBalanceUpdateFixedBuilder) OpenPositionsProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 104, 96)
}

// DailyProfitLoss
func (m AccountBalanceUpdateFixed) DailyProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 112, 104)
}

// DailyProfitLoss
func (m AccountBalanceUpdateFixedBuilder) DailyProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 112, 104)
}

// InfoText
func (m AccountBalanceUpdateFixed) InfoText() string {
	return message.StringFixed(m.Unsafe(), 208, 112)
}

// InfoText
func (m AccountBalanceUpdateFixedBuilder) InfoText() string {
	return message.StringFixed(m.Unsafe(), 208, 112)
}

// TransactionIdentifier
func (m AccountBalanceUpdateFixed) TransactionIdentifier() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 216, 208)
}

// TransactionIdentifier
func (m AccountBalanceUpdateFixedBuilder) TransactionIdentifier() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 216, 208)
}

// DailyNetLossLimit
func (m AccountBalanceUpdateFixed) DailyNetLossLimit() float64 {
	return message.Float64Fixed(m.Unsafe(), 224, 216)
}

// DailyNetLossLimit
func (m AccountBalanceUpdateFixedBuilder) DailyNetLossLimit() float64 {
	return message.Float64Fixed(m.Unsafe(), 224, 216)
}

// TrailingAccountValueToLimitPositions
func (m AccountBalanceUpdateFixed) TrailingAccountValueToLimitPositions() float64 {
	return message.Float64Fixed(m.Unsafe(), 232, 224)
}

// TrailingAccountValueToLimitPositions
func (m AccountBalanceUpdateFixedBuilder) TrailingAccountValueToLimitPositions() float64 {
	return message.Float64Fixed(m.Unsafe(), 232, 224)
}

// DailyNetLossLimitReached
func (m AccountBalanceUpdateFixed) DailyNetLossLimitReached() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 233, 232)
}

// DailyNetLossLimitReached
func (m AccountBalanceUpdateFixedBuilder) DailyNetLossLimitReached() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 233, 232)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdateFixed) IsUnderRequiredMargin() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 234, 233)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdateFixedBuilder) IsUnderRequiredMargin() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 234, 233)
}

// ClosePositionsAtEndOfDay
func (m AccountBalanceUpdateFixed) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 235, 234)
}

// ClosePositionsAtEndOfDay
func (m AccountBalanceUpdateFixedBuilder) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 235, 234)
}

// TradingIsDisabled
func (m AccountBalanceUpdateFixed) TradingIsDisabled() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 236, 235)
}

// TradingIsDisabled
func (m AccountBalanceUpdateFixedBuilder) TradingIsDisabled() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 236, 235)
}

// Description
func (m AccountBalanceUpdateFixed) Description() string {
	return message.StringFixed(m.Unsafe(), 332, 236)
}

// Description
func (m AccountBalanceUpdateFixedBuilder) Description() string {
	return message.StringFixed(m.Unsafe(), 332, 236)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdateFixed) IsUnderRequiredAccountValue() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 333, 332)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdateFixedBuilder) IsUnderRequiredAccountValue() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 333, 332)
}

// TransactionDateTime
func (m AccountBalanceUpdateFixed) TransactionDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 344, 336))
}

// TransactionDateTime
func (m AccountBalanceUpdateFixedBuilder) TransactionDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 344, 336))
}

// MarginRequirementFull
func (m AccountBalanceUpdateFixed) MarginRequirementFull() float64 {
	return message.Float64Fixed(m.Unsafe(), 352, 344)
}

// MarginRequirementFull
func (m AccountBalanceUpdateFixedBuilder) MarginRequirementFull() float64 {
	return message.Float64Fixed(m.Unsafe(), 352, 344)
}

// MarginRequirementFullPositionsOnly
func (m AccountBalanceUpdateFixed) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64Fixed(m.Unsafe(), 360, 352)
}

// MarginRequirementFullPositionsOnly
func (m AccountBalanceUpdateFixedBuilder) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64Fixed(m.Unsafe(), 360, 352)
}

// PeakMarginRequirement
func (m AccountBalanceUpdateFixed) PeakMarginRequirement() float64 {
	return message.Float64Fixed(m.Unsafe(), 368, 360)
}

// PeakMarginRequirement
func (m AccountBalanceUpdateFixedBuilder) PeakMarginRequirement() float64 {
	return message.Float64Fixed(m.Unsafe(), 368, 360)
}

// SetRequestID This is the RequestID which was set in the AccountBalanceRequest that
// this message is in response to.
//
// In the case when this is a periodic unsolicited Account Balance update,
// RequestID must be set to 0, the default.
func (m AccountBalanceUpdateFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetCashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdateFixedBuilder) SetCashBalance(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, value)
}

// SetBalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedBuilder) SetBalanceAvailableForNewPositions(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 24, 16, value)
}

// SetAccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdateFixedBuilder) SetAccountCurrency(value string) {
	message.SetStringFixed(m.Unsafe(), 32, 24, value)
}

// SetTradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdateFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 64, 32, value)
}

// SetSecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdateFixedBuilder) SetSecuritiesValue(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 72, 64, value)
}

// SetMarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 80, 72, value)
}

// SetTotalNumberMessages This indicates the total number of Account Balance Update messages when
// a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateFixedBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 84, 80, value)
}

// SetMessageNumber This indicates the 1-based index of the Account Balance Update message
// when a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateFixedBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 88, 84, value)
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
func (m AccountBalanceUpdateFixedBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 89, 88, value)
}

// SetUnsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdateFixedBuilder) SetUnsolicited(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 90, 89, value)
}

// SetOpenPositionsProfitLoss
func (m AccountBalanceUpdateFixedBuilder) SetOpenPositionsProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 104, 96, value)
}

// SetDailyProfitLoss
func (m AccountBalanceUpdateFixedBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 112, 104, value)
}

// SetInfoText
func (m AccountBalanceUpdateFixedBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Unsafe(), 208, 112, value)
}

// SetTransactionIdentifier
func (m AccountBalanceUpdateFixedBuilder) SetTransactionIdentifier(value uint64) {
	message.SetUint64Fixed(m.Unsafe(), 216, 208, value)
}

// SetDailyNetLossLimit
func (m AccountBalanceUpdateFixedBuilder) SetDailyNetLossLimit(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 224, 216, value)
}

// SetTrailingAccountValueToLimitPositions
func (m AccountBalanceUpdateFixedBuilder) SetTrailingAccountValueToLimitPositions(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 232, 224, value)
}

// SetDailyNetLossLimitReached
func (m AccountBalanceUpdateFixedBuilder) SetDailyNetLossLimitReached(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 233, 232, value)
}

// SetIsUnderRequiredMargin
func (m AccountBalanceUpdateFixedBuilder) SetIsUnderRequiredMargin(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 234, 233, value)
}

// SetClosePositionsAtEndOfDay
func (m AccountBalanceUpdateFixedBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 235, 234, value)
}

// SetTradingIsDisabled
func (m AccountBalanceUpdateFixedBuilder) SetTradingIsDisabled(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 236, 235, value)
}

// SetDescription
func (m AccountBalanceUpdateFixedBuilder) SetDescription(value string) {
	message.SetStringFixed(m.Unsafe(), 332, 236, value)
}

// SetIsUnderRequiredAccountValue
func (m AccountBalanceUpdateFixedBuilder) SetIsUnderRequiredAccountValue(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 333, 332, value)
}

// SetTransactionDateTime
func (m AccountBalanceUpdateFixedBuilder) SetTransactionDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 344, 336, int64(value))
}

// SetMarginRequirementFull
func (m AccountBalanceUpdateFixedBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 352, 344, value)
}

// SetMarginRequirementFullPositionsOnly
func (m AccountBalanceUpdateFixedBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 360, 352, value)
}

// SetPeakMarginRequirement
func (m AccountBalanceUpdateFixedBuilder) SetPeakMarginRequirement(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 368, 360, value)
}

// Copy
func (m AccountBalanceUpdateFixed) Copy(to AccountBalanceUpdateFixedBuilder) {
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
func (m AccountBalanceUpdateFixedBuilder) Copy(to AccountBalanceUpdateFixedBuilder) {
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
func (m AccountBalanceUpdateFixed) CopyPointer(to AccountBalanceUpdateFixedPointerBuilder) {
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
func (m AccountBalanceUpdateFixedBuilder) CopyPointer(to AccountBalanceUpdateFixedPointerBuilder) {
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
func (m AccountBalanceUpdateFixed) CopyToPointer(to AccountBalanceUpdatePointerBuilder) {
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
func (m AccountBalanceUpdateFixedBuilder) CopyToPointer(to AccountBalanceUpdatePointerBuilder) {
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
func (m AccountBalanceUpdateFixed) CopyTo(to AccountBalanceUpdateBuilder) {
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
func (m AccountBalanceUpdateFixedBuilder) CopyTo(to AccountBalanceUpdateBuilder) {
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
