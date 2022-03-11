package model

import (
	"github.com/moontrade/dtc-go/message"
)

// AccountBalanceUpdatePointer This is an optional message from the Server to Client to provide Account
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
type AccountBalanceUpdatePointer struct {
	message.VLSPointer
}

// AccountBalanceUpdatePointerBuilder This is an optional message from the Server to Client to provide Account
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
type AccountBalanceUpdatePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocAccountBalanceUpdate() AccountBalanceUpdatePointerBuilder {
	a := AccountBalanceUpdatePointerBuilder{message.AllocVLSBuilder(168)}
	a.Ptr.SetBytes(0, _AccountBalanceUpdateDefault)
	return a
}

func AllocAccountBalanceUpdateFrom(b []byte) AccountBalanceUpdatePointer {
	return AccountBalanceUpdatePointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m AccountBalanceUpdatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceUpdateDefault)
}

// ToBuilder
func (m AccountBalanceUpdatePointer) ToBuilder() AccountBalanceUpdatePointerBuilder {
	return AccountBalanceUpdatePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *AccountBalanceUpdatePointerBuilder) Finish() AccountBalanceUpdatePointer {
	return AccountBalanceUpdatePointer{m.VLSPointerBuilder.Finish()}
}

// RequestID This is the RequestID which was set in the AccountBalanceRequest that
// this message is in response to.
//
// In the case when this is a periodic unsolicited Account Balance update,
// RequestID must be set to 0, the default.
func (m AccountBalanceUpdatePointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID This is the RequestID which was set in the AccountBalanceRequest that
// this message is in response to.
//
// In the case when this is a periodic unsolicited Account Balance update,
// RequestID must be set to 0, the default.
func (m AccountBalanceUpdatePointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// CashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdatePointer) CashBalance() float64 {
	return message.Float64VLS(m.Ptr, 24, 16)
}

// CashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdatePointerBuilder) CashBalance() float64 {
	return message.Float64VLS(m.Ptr, 24, 16)
}

// BalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdatePointer) BalanceAvailableForNewPositions() float64 {
	return message.Float64VLS(m.Ptr, 32, 24)
}

// BalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdatePointerBuilder) BalanceAvailableForNewPositions() float64 {
	return message.Float64VLS(m.Ptr, 32, 24)
}

// AccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdatePointer) AccountCurrency() string {
	return message.StringVLS(m.Ptr, 36, 32)
}

// AccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdatePointerBuilder) AccountCurrency() string {
	return message.StringVLS(m.Ptr, 36, 32)
}

// TradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdatePointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 40, 36)
}

// TradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdatePointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 40, 36)
}

// SecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdatePointer) SecuritiesValue() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
}

// SecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdatePointerBuilder) SecuritiesValue() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
}

// MarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdatePointer) MarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 56, 48)
}

// MarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdatePointerBuilder) MarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 56, 48)
}

// TotalNumberMessages This indicates the total number of Account Balance Update messages when
// a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdatePointer) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Ptr, 60, 56)
}

// TotalNumberMessages This indicates the total number of Account Balance Update messages when
// a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdatePointerBuilder) TotalNumberMessages() int32 {
	return message.Int32VLS(m.Ptr, 60, 56)
}

// MessageNumber This indicates the 1-based index of the Account Balance Update message
// when a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdatePointer) MessageNumber() int32 {
	return message.Int32VLS(m.Ptr, 64, 60)
}

// MessageNumber This indicates the 1-based index of the Account Balance Update message
// when a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdatePointerBuilder) MessageNumber() int32 {
	return message.Int32VLS(m.Ptr, 64, 60)
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
func (m AccountBalanceUpdatePointer) NoAccountBalances() uint8 {
	return message.Uint8VLS(m.Ptr, 65, 64)
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
func (m AccountBalanceUpdatePointerBuilder) NoAccountBalances() uint8 {
	return message.Uint8VLS(m.Ptr, 65, 64)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdatePointer) Unsolicited() uint8 {
	return message.Uint8VLS(m.Ptr, 66, 65)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdatePointerBuilder) Unsolicited() uint8 {
	return message.Uint8VLS(m.Ptr, 66, 65)
}

// OpenPositionsProfitLoss
func (m AccountBalanceUpdatePointer) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 80, 72)
}

// OpenPositionsProfitLoss
func (m AccountBalanceUpdatePointerBuilder) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 80, 72)
}

// DailyProfitLoss
func (m AccountBalanceUpdatePointer) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 88, 80)
}

// DailyProfitLoss
func (m AccountBalanceUpdatePointerBuilder) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Ptr, 88, 80)
}

// InfoText
func (m AccountBalanceUpdatePointer) InfoText() string {
	return message.StringVLS(m.Ptr, 92, 88)
}

// InfoText
func (m AccountBalanceUpdatePointerBuilder) InfoText() string {
	return message.StringVLS(m.Ptr, 92, 88)
}

// TransactionIdentifier
func (m AccountBalanceUpdatePointer) TransactionIdentifier() uint64 {
	return message.Uint64VLS(m.Ptr, 104, 96)
}

// TransactionIdentifier
func (m AccountBalanceUpdatePointerBuilder) TransactionIdentifier() uint64 {
	return message.Uint64VLS(m.Ptr, 104, 96)
}

// DailyNetLossLimit
func (m AccountBalanceUpdatePointer) DailyNetLossLimit() float64 {
	return message.Float64VLS(m.Ptr, 112, 104)
}

// DailyNetLossLimit
func (m AccountBalanceUpdatePointerBuilder) DailyNetLossLimit() float64 {
	return message.Float64VLS(m.Ptr, 112, 104)
}

// TrailingAccountValueToLimitPositions
func (m AccountBalanceUpdatePointer) TrailingAccountValueToLimitPositions() float64 {
	return message.Float64VLS(m.Ptr, 120, 112)
}

// TrailingAccountValueToLimitPositions
func (m AccountBalanceUpdatePointerBuilder) TrailingAccountValueToLimitPositions() float64 {
	return message.Float64VLS(m.Ptr, 120, 112)
}

// DailyNetLossLimitReached
func (m AccountBalanceUpdatePointer) DailyNetLossLimitReached() uint8 {
	return message.Uint8VLS(m.Ptr, 121, 120)
}

// DailyNetLossLimitReached
func (m AccountBalanceUpdatePointerBuilder) DailyNetLossLimitReached() uint8 {
	return message.Uint8VLS(m.Ptr, 121, 120)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdatePointer) IsUnderRequiredMargin() uint8 {
	return message.Uint8VLS(m.Ptr, 122, 121)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdatePointerBuilder) IsUnderRequiredMargin() uint8 {
	return message.Uint8VLS(m.Ptr, 122, 121)
}

// ClosePositionsAtEndOfDay
func (m AccountBalanceUpdatePointer) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 123, 122)
}

// ClosePositionsAtEndOfDay
func (m AccountBalanceUpdatePointerBuilder) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Ptr, 123, 122)
}

// TradingIsDisabled
func (m AccountBalanceUpdatePointer) TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Ptr, 124, 123)
}

// TradingIsDisabled
func (m AccountBalanceUpdatePointerBuilder) TradingIsDisabled() uint8 {
	return message.Uint8VLS(m.Ptr, 124, 123)
}

// Description
func (m AccountBalanceUpdatePointer) Description() string {
	return message.StringVLS(m.Ptr, 128, 124)
}

// Description
func (m AccountBalanceUpdatePointerBuilder) Description() string {
	return message.StringVLS(m.Ptr, 128, 124)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdatePointer) IsUnderRequiredAccountValue() uint8 {
	return message.Uint8VLS(m.Ptr, 129, 128)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdatePointerBuilder) IsUnderRequiredAccountValue() uint8 {
	return message.Uint8VLS(m.Ptr, 129, 128)
}

// TransactionDateTime
func (m AccountBalanceUpdatePointer) TransactionDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Ptr, 144, 136))
}

// TransactionDateTime
func (m AccountBalanceUpdatePointerBuilder) TransactionDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Ptr, 144, 136))
}

// MarginRequirementFull
func (m AccountBalanceUpdatePointer) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Ptr, 152, 144)
}

// MarginRequirementFull
func (m AccountBalanceUpdatePointerBuilder) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Ptr, 152, 144)
}

// MarginRequirementFullPositionsOnly
func (m AccountBalanceUpdatePointer) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Ptr, 160, 152)
}

// MarginRequirementFullPositionsOnly
func (m AccountBalanceUpdatePointerBuilder) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Ptr, 160, 152)
}

// PeakMarginRequirement
func (m AccountBalanceUpdatePointer) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 168, 160)
}

// PeakMarginRequirement
func (m AccountBalanceUpdatePointerBuilder) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 168, 160)
}

// SetRequestID This is the RequestID which was set in the AccountBalanceRequest that
// this message is in response to.
//
// In the case when this is a periodic unsolicited Account Balance update,
// RequestID must be set to 0, the default.
func (m AccountBalanceUpdatePointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetCashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdatePointerBuilder) SetCashBalance(value float64) {
	message.SetFloat64VLS(m.Ptr, 24, 16, value)
}

// SetBalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdatePointerBuilder) SetBalanceAvailableForNewPositions(value float64) {
	message.SetFloat64VLS(m.Ptr, 32, 24, value)
}

// SetAccountCurrency ISO Currency Code for the cash values in this message.
func (m *AccountBalanceUpdatePointerBuilder) SetAccountCurrency(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 36, 32, value)
}

// SetTradeAccount The trade account identifier for the Account Balance information.
func (m *AccountBalanceUpdatePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 40, 36, value)
}

// SetSecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdatePointerBuilder) SetSecuritiesValue(value float64) {
	message.SetFloat64VLS(m.Ptr, 48, 40, value)
}

// SetMarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdatePointerBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Ptr, 56, 48, value)
}

// SetTotalNumberMessages This indicates the total number of Account Balance Update messages when
// a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdatePointerBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32VLS(m.Ptr, 60, 56, value)
}

// SetMessageNumber This indicates the 1-based index of the Account Balance Update message
// when a batch of messages is being sent. If there is only one Account Balance
// Update message being sent, this will be 1. Use a value of 1 for an unsolicited
// message.
//
// A Client should not rely on this field for an unsolicited message. This
// is required to be set.
func (m AccountBalanceUpdatePointerBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Ptr, 64, 60, value)
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
func (m AccountBalanceUpdatePointerBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8VLS(m.Ptr, 65, 64, value)
}

// SetUnsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdatePointerBuilder) SetUnsolicited(value uint8) {
	message.SetUint8VLS(m.Ptr, 66, 65, value)
}

// SetOpenPositionsProfitLoss
func (m AccountBalanceUpdatePointerBuilder) SetOpenPositionsProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 80, 72, value)
}

// SetDailyProfitLoss
func (m AccountBalanceUpdatePointerBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 88, 80, value)
}

// SetInfoText
func (m *AccountBalanceUpdatePointerBuilder) SetInfoText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 92, 88, value)
}

// SetTransactionIdentifier
func (m AccountBalanceUpdatePointerBuilder) SetTransactionIdentifier(value uint64) {
	message.SetUint64VLS(m.Ptr, 104, 96, value)
}

// SetDailyNetLossLimit
func (m AccountBalanceUpdatePointerBuilder) SetDailyNetLossLimit(value float64) {
	message.SetFloat64VLS(m.Ptr, 112, 104, value)
}

// SetTrailingAccountValueToLimitPositions
func (m AccountBalanceUpdatePointerBuilder) SetTrailingAccountValueToLimitPositions(value float64) {
	message.SetFloat64VLS(m.Ptr, 120, 112, value)
}

// SetDailyNetLossLimitReached
func (m AccountBalanceUpdatePointerBuilder) SetDailyNetLossLimitReached(value uint8) {
	message.SetUint8VLS(m.Ptr, 121, 120, value)
}

// SetIsUnderRequiredMargin
func (m AccountBalanceUpdatePointerBuilder) SetIsUnderRequiredMargin(value uint8) {
	message.SetUint8VLS(m.Ptr, 122, 121, value)
}

// SetClosePositionsAtEndOfDay
func (m AccountBalanceUpdatePointerBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 123, 122, value)
}

// SetTradingIsDisabled
func (m AccountBalanceUpdatePointerBuilder) SetTradingIsDisabled(value uint8) {
	message.SetUint8VLS(m.Ptr, 124, 123, value)
}

// SetDescription
func (m *AccountBalanceUpdatePointerBuilder) SetDescription(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 128, 124, value)
}

// SetIsUnderRequiredAccountValue
func (m AccountBalanceUpdatePointerBuilder) SetIsUnderRequiredAccountValue(value uint8) {
	message.SetUint8VLS(m.Ptr, 129, 128, value)
}

// SetTransactionDateTime
func (m AccountBalanceUpdatePointerBuilder) SetTransactionDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64VLS(m.Ptr, 144, 136, int64(value))
}

// SetMarginRequirementFull
func (m AccountBalanceUpdatePointerBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64VLS(m.Ptr, 152, 144, value)
}

// SetMarginRequirementFullPositionsOnly
func (m AccountBalanceUpdatePointerBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64VLS(m.Ptr, 160, 152, value)
}

// SetPeakMarginRequirement
func (m AccountBalanceUpdatePointerBuilder) SetPeakMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Ptr, 168, 160, value)
}

// Copy
func (m AccountBalanceUpdatePointer) Copy(to AccountBalanceUpdateBuilder) {
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
func (m AccountBalanceUpdatePointerBuilder) Copy(to AccountBalanceUpdateBuilder) {
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
func (m AccountBalanceUpdatePointer) CopyPointer(to AccountBalanceUpdatePointerBuilder) {
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
func (m AccountBalanceUpdatePointerBuilder) CopyPointer(to AccountBalanceUpdatePointerBuilder) {
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
func (m AccountBalanceUpdatePointer) CopyTo(to AccountBalanceUpdateFixedBuilder) {
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
func (m AccountBalanceUpdatePointerBuilder) CopyTo(to AccountBalanceUpdateFixedBuilder) {
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
func (m AccountBalanceUpdatePointer) CopyToPointer(to AccountBalanceUpdateFixedPointerBuilder) {
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
func (m AccountBalanceUpdatePointerBuilder) CopyToPointer(to AccountBalanceUpdateFixedPointerBuilder) {
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
