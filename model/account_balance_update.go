package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                                 = AccountBalanceUpdateSize  (168)
//     Type                                 = ACCOUNT_BALANCE_UPDATE  (600)
//     BaseSize                             = AccountBalanceUpdateSize  (168)
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
var _AccountBalanceUpdateDefault = []byte{168, 0, 88, 2, 168, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const AccountBalanceUpdateSize = 168

// AccountBalanceUpdate This is an optional message from the Server to Client to provide Account
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
type AccountBalanceUpdate struct {
	message.VLS
}

// AccountBalanceUpdateBuilder This is an optional message from the Server to Client to provide Account
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
type AccountBalanceUpdateBuilder struct {
	message.VLSBuilder
}

func NewAccountBalanceUpdateFrom(b []byte) AccountBalanceUpdate {
	return AccountBalanceUpdate{VLS: message.NewVLSFrom(b)}
}

func WrapAccountBalanceUpdate(b []byte) AccountBalanceUpdate {
	return AccountBalanceUpdate{VLS: message.WrapVLS(b)}
}

func NewAccountBalanceUpdate() AccountBalanceUpdateBuilder {
	a := AccountBalanceUpdateBuilder{message.NewVLSBuilder(168)}
	a.Unsafe().SetBytes(0, _AccountBalanceUpdateDefault)
	return a
}

// Clear
// {
//     Size                                 = AccountBalanceUpdateSize  (168)
//     Type                                 = ACCOUNT_BALANCE_UPDATE  (600)
//     BaseSize                             = AccountBalanceUpdateSize  (168)
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
func (m AccountBalanceUpdateBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceUpdateDefault)
}

// ToBuilder
func (m AccountBalanceUpdate) ToBuilder() AccountBalanceUpdateBuilder {
	return AccountBalanceUpdateBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m AccountBalanceUpdateBuilder) Finish() AccountBalanceUpdate {
	return AccountBalanceUpdate{m.VLSBuilder.Finish()}
}

// RequestID This is the RequestID which was set in the AccountBalanceRequest that
// this message is in response to.
//
// In the case when this is a periodic unsolicited Account Balance update,
// RequestID must be set to 0, the default.
func (m AccountBalanceUpdate) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID This is the RequestID which was set in the AccountBalanceRequest that
// this message is in response to.
//
// In the case when this is a periodic unsolicited Account Balance update,
// RequestID must be set to 0, the default.
func (m AccountBalanceUpdateBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// CashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdate) CashBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 24, 16)
}

// CashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdateBuilder) CashBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 24, 16)
}

// BalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdate) BalanceAvailableForNewPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
}

// BalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateBuilder) BalanceAvailableForNewPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
}

// AccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdate) AccountCurrency() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
}

// AccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdateBuilder) AccountCurrency() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
}

// TradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdate) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// TradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdateBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// SecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdate) SecuritiesValue() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// SecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdateBuilder) SecuritiesValue() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// MarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdate) MarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 56, 48)
}

// MarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateBuilder) MarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 56, 48)
}

// TotalNumberMessages This indicates the total number of Account Balance Update messages when
// a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdate) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Unsafe(), 60, 56)
}

// TotalNumberMessages This indicates the total number of Account Balance Update messages when
// a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateBuilder) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Unsafe(), 60, 56)
}

// MessageNumber This indicates the 1-based index of the Account Balance Update message
// when a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdate) MessageNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 64, 60)
}

// MessageNumber This indicates the 1-based index of the Account Balance Update message
// when a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateBuilder) MessageNumber() int32 {
	return message.Int32VLS(m.Unsafe(), 64, 60)
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
func (m AccountBalanceUpdate) NoAccountBalances() uint8 {
	return message.Uint8VLS(m.Unsafe(), 65, 64)
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
func (m AccountBalanceUpdateBuilder) NoAccountBalances() uint8 {
	return message.Uint8VLS(m.Unsafe(), 65, 64)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdate) Unsolicited() uint8 {
	return message.Uint8VLS(m.Unsafe(), 66, 65)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdateBuilder) Unsolicited() uint8 {
	return message.Uint8VLS(m.Unsafe(), 66, 65)
}

// OpenPositionsProfitLoss
func (m AccountBalanceUpdate) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 80, 72)
}

// OpenPositionsProfitLoss
func (m AccountBalanceUpdateBuilder) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 80, 72)
}

// DailyProfitLoss
func (m AccountBalanceUpdate) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 88, 80)
}

// DailyProfitLoss
func (m AccountBalanceUpdateBuilder) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 88, 80)
}

// InfoText
func (m AccountBalanceUpdate) InfoText() string {
	return message.StringVLS(m.Unsafe(), 92, 88)
}

// InfoText
func (m AccountBalanceUpdateBuilder) InfoText() string {
	return message.StringVLS(m.Unsafe(), 92, 88)
}

// TransactionIdentifier
func (m AccountBalanceUpdate) TransactionIdentifier() uint64 {
	return message.Uint64VLS(m.Unsafe(), 104, 96)
}

// TransactionIdentifier
func (m AccountBalanceUpdateBuilder) TransactionIdentifier() uint64 {
	return message.Uint64VLS(m.Unsafe(), 104, 96)
}

// DailyNetLossLimit
func (m AccountBalanceUpdate) DailyNetLossLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 112, 104)
}

// DailyNetLossLimit
func (m AccountBalanceUpdateBuilder) DailyNetLossLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 112, 104)
}

// TrailingAccountValueToLimitPositions
func (m AccountBalanceUpdate) TrailingAccountValueToLimitPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 120, 112)
}

// TrailingAccountValueToLimitPositions
func (m AccountBalanceUpdateBuilder) TrailingAccountValueToLimitPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 120, 112)
}

// DailyNetLossLimitReached
func (m AccountBalanceUpdate) DailyNetLossLimitReached() uint8 {
	return message.Uint8VLS(m.Unsafe(), 121, 120)
}

// DailyNetLossLimitReached
func (m AccountBalanceUpdateBuilder) DailyNetLossLimitReached() uint8 {
	return message.Uint8VLS(m.Unsafe(), 121, 120)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdate) IsUnderRequiredMargin() uint8 {
	return message.Uint8VLS(m.Unsafe(), 122, 121)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdateBuilder) IsUnderRequiredMargin() uint8 {
	return message.Uint8VLS(m.Unsafe(), 122, 121)
}

// ClosePositionsAtEndOfDay
func (m AccountBalanceUpdate) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 123, 122)
}

// ClosePositionsAtEndOfDay
func (m AccountBalanceUpdateBuilder) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 123, 122)
}

// TradingIsDisabled
func (m AccountBalanceUpdate) TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Unsafe(), 124, 123)
}

// TradingIsDisabled
func (m AccountBalanceUpdateBuilder) TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Unsafe(), 124, 123)
}

// Description
func (m AccountBalanceUpdate) Description() string {
	return message.StringVLS(m.Unsafe(), 128, 124)
}

// Description
func (m AccountBalanceUpdateBuilder) Description() string {
	return message.StringVLS(m.Unsafe(), 128, 124)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdate) IsUnderRequiredAccountValue() uint8 {
	return message.Uint8VLS(m.Unsafe(), 129, 128)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdateBuilder) IsUnderRequiredAccountValue() uint8 {
	return message.Uint8VLS(m.Unsafe(), 129, 128)
}

// TransactionDateTime
func (m AccountBalanceUpdate) TransactionDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Unsafe(), 144, 136))
}

// TransactionDateTime
func (m AccountBalanceUpdateBuilder) TransactionDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Unsafe(), 144, 136))
}

// MarginRequirementFull
func (m AccountBalanceUpdate) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Unsafe(), 152, 144)
}

// MarginRequirementFull
func (m AccountBalanceUpdateBuilder) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Unsafe(), 152, 144)
}

// MarginRequirementFullPositionsOnly
func (m AccountBalanceUpdate) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Unsafe(), 160, 152)
}

// MarginRequirementFullPositionsOnly
func (m AccountBalanceUpdateBuilder) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Unsafe(), 160, 152)
}

// PeakMarginRequirement
func (m AccountBalanceUpdate) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 168, 160)
}

// PeakMarginRequirement
func (m AccountBalanceUpdateBuilder) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 168, 160)
}

// SetRequestID This is the RequestID which was set in the AccountBalanceRequest that
// this message is in response to.
//
// In the case when this is a periodic unsolicited Account Balance update,
// RequestID must be set to 0, the default.
func (m AccountBalanceUpdateBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetCashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdateBuilder) SetCashBalance(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 24, 16, value)
}

// SetBalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateBuilder) SetBalanceAvailableForNewPositions(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 32, 24, value)
}

// SetAccountCurrency ISO Currency Code for the cash values in this message.
func (m *AccountBalanceUpdateBuilder) SetAccountCurrency(value string) {
	message.SetStringVLS(&m.VLSBuilder, 36, 32, value)
}

// SetTradeAccount The trade account identifier for the Account Balance information.
func (m *AccountBalanceUpdateBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 40, 36, value)
}

// SetSecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdateBuilder) SetSecuritiesValue(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 48, 40, value)
}

// SetMarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 56, 48, value)
}

// SetTotalNumberMessages This indicates the total number of Account Balance Update messages when
// a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32VLS(m.Unsafe(), 60, 56, value)
}

// SetMessageNumber This indicates the 1-based index of the Account Balance Update message
// when a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdateBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Unsafe(), 64, 60, value)
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
func (m AccountBalanceUpdateBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 65, 64, value)
}

// SetUnsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdateBuilder) SetUnsolicited(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 66, 65, value)
}

// SetOpenPositionsProfitLoss
func (m AccountBalanceUpdateBuilder) SetOpenPositionsProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 80, 72, value)
}

// SetDailyProfitLoss
func (m AccountBalanceUpdateBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 88, 80, value)
}

// SetInfoText
func (m *AccountBalanceUpdateBuilder) SetInfoText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 92, 88, value)
}

// SetTransactionIdentifier
func (m AccountBalanceUpdateBuilder) SetTransactionIdentifier(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 104, 96, value)
}

// SetDailyNetLossLimit
func (m AccountBalanceUpdateBuilder) SetDailyNetLossLimit(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 112, 104, value)
}

// SetTrailingAccountValueToLimitPositions
func (m AccountBalanceUpdateBuilder) SetTrailingAccountValueToLimitPositions(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 120, 112, value)
}

// SetDailyNetLossLimitReached
func (m AccountBalanceUpdateBuilder) SetDailyNetLossLimitReached(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 121, 120, value)
}

// SetIsUnderRequiredMargin
func (m AccountBalanceUpdateBuilder) SetIsUnderRequiredMargin(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 122, 121, value)
}

// SetClosePositionsAtEndOfDay
func (m AccountBalanceUpdateBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 123, 122, value)
}

// SetTradingIsDisabled
func (m AccountBalanceUpdateBuilder) SetTradingIsDisabled(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 124, 123, value)
}

// SetDescription
func (m *AccountBalanceUpdateBuilder) SetDescription(value string) {
	message.SetStringVLS(&m.VLSBuilder, 128, 124, value)
}

// SetIsUnderRequiredAccountValue
func (m AccountBalanceUpdateBuilder) SetIsUnderRequiredAccountValue(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 129, 128, value)
}

// SetTransactionDateTime
func (m AccountBalanceUpdateBuilder) SetTransactionDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 144, 136, int64(value))
}

// SetMarginRequirementFull
func (m AccountBalanceUpdateBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 152, 144, value)
}

// SetMarginRequirementFullPositionsOnly
func (m AccountBalanceUpdateBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 160, 152, value)
}

// SetPeakMarginRequirement
func (m AccountBalanceUpdateBuilder) SetPeakMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 168, 160, value)
}

// Copy
func (m AccountBalanceUpdate) Copy(to AccountBalanceUpdateBuilder) {
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
func (m AccountBalanceUpdateBuilder) Copy(to AccountBalanceUpdateBuilder) {
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
func (m AccountBalanceUpdate) CopyPointer(to AccountBalanceUpdatePointerBuilder) {
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
func (m AccountBalanceUpdateBuilder) CopyPointer(to AccountBalanceUpdatePointerBuilder) {
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
func (m AccountBalanceUpdate) CopyToPointer(to AccountBalanceUpdateFixedPointerBuilder) {
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
func (m AccountBalanceUpdateBuilder) CopyToPointer(to AccountBalanceUpdateFixedPointerBuilder) {
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
func (m AccountBalanceUpdate) CopyTo(to AccountBalanceUpdateFixedBuilder) {
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
func (m AccountBalanceUpdateBuilder) CopyTo(to AccountBalanceUpdateFixedBuilder) {
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
