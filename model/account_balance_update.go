package model

import (
	"github.com/moontrade/dtc-go/message"
)

const AccountBalanceUpdateSize = 168

const AccountBalanceUpdateFixedSize = 368

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
//     IsUnderRequiredMargin                = false
//     ClosePositionsAtEndOfDay             = 0
//     TradingIsDisabled                    = false
//     Description                          = ""
//     IsUnderRequiredAccountValue          = false
//     TransactionDateTime                  = 0
//     MarginRequirementFull                = 0
//     MarginRequirementFullPositionsOnly   = 0
//     PeakMarginRequirement                = 0
// }
var _AccountBalanceUpdateDefault = []byte{168, 0, 88, 2, 168, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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
//     IsUnderRequiredMargin                = false
//     ClosePositionsAtEndOfDay             = 0
//     TradingIsDisabled                    = false
//     Description                          = ""
//     IsUnderRequiredAccountValue          = false
//     TransactionDateTime                  = 0
//     MarginRequirementFull                = 0
//     MarginRequirementFullPositionsOnly   = 0
//     PeakMarginRequirement                = 0
// }
var _AccountBalanceUpdateFixedDefault = []byte{112, 1, 88, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func AllocAccountBalanceUpdate() AccountBalanceUpdatePointerBuilder {
	a := AccountBalanceUpdatePointerBuilder{message.AllocVLSBuilder(168)}
	a.Ptr.SetBytes(0, _AccountBalanceUpdateDefault)
	return a
}

func AllocAccountBalanceUpdateFrom(b []byte) AccountBalanceUpdatePointer {
	return AccountBalanceUpdatePointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     IsUnderRequiredMargin                = false
//     ClosePositionsAtEndOfDay             = 0
//     TradingIsDisabled                    = false
//     Description                          = ""
//     IsUnderRequiredAccountValue          = false
//     TransactionDateTime                  = 0
//     MarginRequirementFull                = 0
//     MarginRequirementFullPositionsOnly   = 0
//     PeakMarginRequirement                = 0
// }
func (m AccountBalanceUpdateBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceUpdateDefault)
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
//     IsUnderRequiredMargin                = false
//     ClosePositionsAtEndOfDay             = 0
//     TradingIsDisabled                    = false
//     Description                          = ""
//     IsUnderRequiredAccountValue          = false
//     TransactionDateTime                  = 0
//     MarginRequirementFull                = 0
//     MarginRequirementFullPositionsOnly   = 0
//     PeakMarginRequirement                = 0
// }
func (m AccountBalanceUpdateFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _AccountBalanceUpdateFixedDefault)
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
//     IsUnderRequiredMargin                = false
//     ClosePositionsAtEndOfDay             = 0
//     TradingIsDisabled                    = false
//     Description                          = ""
//     IsUnderRequiredAccountValue          = false
//     TransactionDateTime                  = 0
//     MarginRequirementFull                = 0
//     MarginRequirementFullPositionsOnly   = 0
//     PeakMarginRequirement                = 0
// }
func (m AccountBalanceUpdatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceUpdateDefault)
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
//     IsUnderRequiredMargin                = false
//     ClosePositionsAtEndOfDay             = 0
//     TradingIsDisabled                    = false
//     Description                          = ""
//     IsUnderRequiredAccountValue          = false
//     TransactionDateTime                  = 0
//     MarginRequirementFull                = 0
//     MarginRequirementFullPositionsOnly   = 0
//     PeakMarginRequirement                = 0
// }
func (m AccountBalanceUpdateFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _AccountBalanceUpdateFixedDefault)
}

// ToBuilder
func (m AccountBalanceUpdate) ToBuilder() AccountBalanceUpdateBuilder {
	return AccountBalanceUpdateBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m AccountBalanceUpdateFixed) ToBuilder() AccountBalanceUpdateFixedBuilder {
	return AccountBalanceUpdateFixedBuilder{m.Fixed}
}

// ToBuilder
func (m AccountBalanceUpdatePointer) ToBuilder() AccountBalanceUpdatePointerBuilder {
	return AccountBalanceUpdatePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m AccountBalanceUpdateFixedPointer) ToBuilder() AccountBalanceUpdateFixedPointerBuilder {
	return AccountBalanceUpdateFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m AccountBalanceUpdateBuilder) Finish() AccountBalanceUpdate {
	return AccountBalanceUpdate{m.VLSBuilder.Finish()}
}

// Finish
func (m AccountBalanceUpdateFixedBuilder) Finish() AccountBalanceUpdateFixed {
	return AccountBalanceUpdateFixed{m.Fixed}
}

// Finish
func (m *AccountBalanceUpdatePointerBuilder) Finish() AccountBalanceUpdatePointer {
	return AccountBalanceUpdatePointer{m.VLSPointerBuilder.Finish()}
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
func (m AccountBalanceUpdate) CashBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 24, 16)
}

// CashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdateBuilder) CashBalance() float64 {
	return message.Float64VLS(m.Unsafe(), 24, 16)
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
func (m AccountBalanceUpdate) BalanceAvailableForNewPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
}

// BalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateBuilder) BalanceAvailableForNewPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 32, 24)
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
func (m AccountBalanceUpdate) AccountCurrency() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
}

// AccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdateBuilder) AccountCurrency() string {
	return message.StringVLS(m.Unsafe(), 36, 32)
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
func (m AccountBalanceUpdate) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
}

// TradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdateBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 40, 36)
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
func (m AccountBalanceUpdate) SecuritiesValue() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// SecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdateBuilder) SecuritiesValue() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
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
func (m AccountBalanceUpdate) MarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 56, 48)
}

// MarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateBuilder) MarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 56, 48)
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
func (m AccountBalanceUpdate) Unsolicited() uint8 {
	return message.Uint8VLS(m.Unsafe(), 66, 65)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdateBuilder) Unsolicited() uint8 {
	return message.Uint8VLS(m.Unsafe(), 66, 65)
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
func (m AccountBalanceUpdate) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 80, 72)
}

// OpenPositionsProfitLoss
func (m AccountBalanceUpdateBuilder) OpenPositionsProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 80, 72)
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
func (m AccountBalanceUpdate) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 88, 80)
}

// DailyProfitLoss
func (m AccountBalanceUpdateBuilder) DailyProfitLoss() float64 {
	return message.Float64VLS(m.Unsafe(), 88, 80)
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
func (m AccountBalanceUpdate) InfoText() string {
	return message.StringVLS(m.Unsafe(), 92, 88)
}

// InfoText
func (m AccountBalanceUpdateBuilder) InfoText() string {
	return message.StringVLS(m.Unsafe(), 92, 88)
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
func (m AccountBalanceUpdate) TransactionIdentifier() uint64 {
	return message.Uint64VLS(m.Unsafe(), 104, 96)
}

// TransactionIdentifier
func (m AccountBalanceUpdateBuilder) TransactionIdentifier() uint64 {
	return message.Uint64VLS(m.Unsafe(), 104, 96)
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
func (m AccountBalanceUpdate) DailyNetLossLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 112, 104)
}

// DailyNetLossLimit
func (m AccountBalanceUpdateBuilder) DailyNetLossLimit() float64 {
	return message.Float64VLS(m.Unsafe(), 112, 104)
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
func (m AccountBalanceUpdate) TrailingAccountValueToLimitPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 120, 112)
}

// TrailingAccountValueToLimitPositions
func (m AccountBalanceUpdateBuilder) TrailingAccountValueToLimitPositions() float64 {
	return message.Float64VLS(m.Unsafe(), 120, 112)
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
func (m AccountBalanceUpdate) DailyNetLossLimitReached() uint8 {
	return message.Uint8VLS(m.Unsafe(), 121, 120)
}

// DailyNetLossLimitReached
func (m AccountBalanceUpdateBuilder) DailyNetLossLimitReached() uint8 {
	return message.Uint8VLS(m.Unsafe(), 121, 120)
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
func (m AccountBalanceUpdate) IsUnderRequiredMargin() bool {
	return message.BoolVLS(m.Unsafe(), 122, 121)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdateBuilder) IsUnderRequiredMargin() bool {
	return message.BoolVLS(m.Unsafe(), 122, 121)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdatePointer) IsUnderRequiredMargin() bool {
	return message.BoolVLS(m.Ptr, 122, 121)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdatePointerBuilder) IsUnderRequiredMargin() bool {
	return message.BoolVLS(m.Ptr, 122, 121)
}

// ClosePositionsAtEndOfDay
func (m AccountBalanceUpdate) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 123, 122)
}

// ClosePositionsAtEndOfDay
func (m AccountBalanceUpdateBuilder) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8VLS(m.Unsafe(), 123, 122)
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
func (m AccountBalanceUpdate) TradingIsDisabled() bool {
	return message.BoolVLS(m.Unsafe(), 124, 123)
}

// TradingIsDisabled
func (m AccountBalanceUpdateBuilder) TradingIsDisabled() bool {
	return message.BoolVLS(m.Unsafe(), 124, 123)
}

// TradingIsDisabled
func (m AccountBalanceUpdatePointer) TradingIsDisabled() bool {
	return message.BoolVLS(m.Ptr, 124, 123)
}

// TradingIsDisabled
func (m AccountBalanceUpdatePointerBuilder) TradingIsDisabled() bool {
	return message.BoolVLS(m.Ptr, 124, 123)
}

// Description
func (m AccountBalanceUpdate) Description() string {
	return message.StringVLS(m.Unsafe(), 128, 124)
}

// Description
func (m AccountBalanceUpdateBuilder) Description() string {
	return message.StringVLS(m.Unsafe(), 128, 124)
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
func (m AccountBalanceUpdate) IsUnderRequiredAccountValue() bool {
	return message.BoolVLS(m.Unsafe(), 129, 128)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdateBuilder) IsUnderRequiredAccountValue() bool {
	return message.BoolVLS(m.Unsafe(), 129, 128)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdatePointer) IsUnderRequiredAccountValue() bool {
	return message.BoolVLS(m.Ptr, 129, 128)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdatePointerBuilder) IsUnderRequiredAccountValue() bool {
	return message.BoolVLS(m.Ptr, 129, 128)
}

// TransactionDateTime
func (m AccountBalanceUpdate) TransactionDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Unsafe(), 144, 136))
}

// TransactionDateTime
func (m AccountBalanceUpdateBuilder) TransactionDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64VLS(m.Unsafe(), 144, 136))
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
func (m AccountBalanceUpdate) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Unsafe(), 152, 144)
}

// MarginRequirementFull
func (m AccountBalanceUpdateBuilder) MarginRequirementFull() float64 {
	return message.Float64VLS(m.Unsafe(), 152, 144)
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
func (m AccountBalanceUpdate) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Unsafe(), 160, 152)
}

// MarginRequirementFullPositionsOnly
func (m AccountBalanceUpdateBuilder) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64VLS(m.Unsafe(), 160, 152)
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
func (m AccountBalanceUpdate) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 168, 160)
}

// PeakMarginRequirement
func (m AccountBalanceUpdateBuilder) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Unsafe(), 168, 160)
}

// PeakMarginRequirement
func (m AccountBalanceUpdatePointer) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 168, 160)
}

// PeakMarginRequirement
func (m AccountBalanceUpdatePointerBuilder) PeakMarginRequirement() float64 {
	return message.Float64VLS(m.Ptr, 168, 160)
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
func (m AccountBalanceUpdateFixed) CashBalance() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
}

// CashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdateFixedBuilder) CashBalance() float64 {
	return message.Float64Fixed(m.Unsafe(), 16, 8)
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
func (m AccountBalanceUpdateFixed) BalanceAvailableForNewPositions() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
}

// BalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedBuilder) BalanceAvailableForNewPositions() float64 {
	return message.Float64Fixed(m.Unsafe(), 24, 16)
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
func (m AccountBalanceUpdateFixed) AccountCurrency() string {
	return message.StringFixed(m.Unsafe(), 32, 24)
}

// AccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdateFixedBuilder) AccountCurrency() string {
	return message.StringFixed(m.Unsafe(), 32, 24)
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
func (m AccountBalanceUpdateFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 64, 32)
}

// TradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdateFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 64, 32)
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
func (m AccountBalanceUpdateFixed) SecuritiesValue() float64 {
	return message.Float64Fixed(m.Unsafe(), 72, 64)
}

// SecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdateFixedBuilder) SecuritiesValue() float64 {
	return message.Float64Fixed(m.Unsafe(), 72, 64)
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
func (m AccountBalanceUpdateFixed) MarginRequirement() float64 {
	return message.Float64Fixed(m.Unsafe(), 80, 72)
}

// MarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedBuilder) MarginRequirement() float64 {
	return message.Float64Fixed(m.Unsafe(), 80, 72)
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
func (m AccountBalanceUpdateFixed) Unsolicited() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 90, 89)
}

// Unsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdateFixedBuilder) Unsolicited() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 90, 89)
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
func (m AccountBalanceUpdateFixed) OpenPositionsProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 104, 96)
}

// OpenPositionsProfitLoss
func (m AccountBalanceUpdateFixedBuilder) OpenPositionsProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 104, 96)
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
func (m AccountBalanceUpdateFixed) DailyProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 112, 104)
}

// DailyProfitLoss
func (m AccountBalanceUpdateFixedBuilder) DailyProfitLoss() float64 {
	return message.Float64Fixed(m.Unsafe(), 112, 104)
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
func (m AccountBalanceUpdateFixed) InfoText() string {
	return message.StringFixed(m.Unsafe(), 208, 112)
}

// InfoText
func (m AccountBalanceUpdateFixedBuilder) InfoText() string {
	return message.StringFixed(m.Unsafe(), 208, 112)
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
func (m AccountBalanceUpdateFixed) TransactionIdentifier() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 216, 208)
}

// TransactionIdentifier
func (m AccountBalanceUpdateFixedBuilder) TransactionIdentifier() uint64 {
	return message.Uint64Fixed(m.Unsafe(), 216, 208)
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
func (m AccountBalanceUpdateFixed) DailyNetLossLimit() float64 {
	return message.Float64Fixed(m.Unsafe(), 224, 216)
}

// DailyNetLossLimit
func (m AccountBalanceUpdateFixedBuilder) DailyNetLossLimit() float64 {
	return message.Float64Fixed(m.Unsafe(), 224, 216)
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
func (m AccountBalanceUpdateFixed) TrailingAccountValueToLimitPositions() float64 {
	return message.Float64Fixed(m.Unsafe(), 232, 224)
}

// TrailingAccountValueToLimitPositions
func (m AccountBalanceUpdateFixedBuilder) TrailingAccountValueToLimitPositions() float64 {
	return message.Float64Fixed(m.Unsafe(), 232, 224)
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
func (m AccountBalanceUpdateFixed) DailyNetLossLimitReached() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 233, 232)
}

// DailyNetLossLimitReached
func (m AccountBalanceUpdateFixedBuilder) DailyNetLossLimitReached() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 233, 232)
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
func (m AccountBalanceUpdateFixed) IsUnderRequiredMargin() bool {
	return message.BoolFixed(m.Unsafe(), 234, 233)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdateFixedBuilder) IsUnderRequiredMargin() bool {
	return message.BoolFixed(m.Unsafe(), 234, 233)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdateFixedPointer) IsUnderRequiredMargin() bool {
	return message.BoolFixed(m.Ptr, 234, 233)
}

// IsUnderRequiredMargin
func (m AccountBalanceUpdateFixedPointerBuilder) IsUnderRequiredMargin() bool {
	return message.BoolFixed(m.Ptr, 234, 233)
}

// ClosePositionsAtEndOfDay
func (m AccountBalanceUpdateFixed) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 235, 234)
}

// ClosePositionsAtEndOfDay
func (m AccountBalanceUpdateFixedBuilder) ClosePositionsAtEndOfDay() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 235, 234)
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
func (m AccountBalanceUpdateFixed) TradingIsDisabled() bool {
	return message.BoolFixed(m.Unsafe(), 236, 235)
}

// TradingIsDisabled
func (m AccountBalanceUpdateFixedBuilder) TradingIsDisabled() bool {
	return message.BoolFixed(m.Unsafe(), 236, 235)
}

// TradingIsDisabled
func (m AccountBalanceUpdateFixedPointer) TradingIsDisabled() bool {
	return message.BoolFixed(m.Ptr, 236, 235)
}

// TradingIsDisabled
func (m AccountBalanceUpdateFixedPointerBuilder) TradingIsDisabled() bool {
	return message.BoolFixed(m.Ptr, 236, 235)
}

// Description
func (m AccountBalanceUpdateFixed) Description() string {
	return message.StringFixed(m.Unsafe(), 332, 236)
}

// Description
func (m AccountBalanceUpdateFixedBuilder) Description() string {
	return message.StringFixed(m.Unsafe(), 332, 236)
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
func (m AccountBalanceUpdateFixed) IsUnderRequiredAccountValue() bool {
	return message.BoolFixed(m.Unsafe(), 333, 332)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdateFixedBuilder) IsUnderRequiredAccountValue() bool {
	return message.BoolFixed(m.Unsafe(), 333, 332)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdateFixedPointer) IsUnderRequiredAccountValue() bool {
	return message.BoolFixed(m.Ptr, 333, 332)
}

// IsUnderRequiredAccountValue
func (m AccountBalanceUpdateFixedPointerBuilder) IsUnderRequiredAccountValue() bool {
	return message.BoolFixed(m.Ptr, 333, 332)
}

// TransactionDateTime
func (m AccountBalanceUpdateFixed) TransactionDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 344, 336))
}

// TransactionDateTime
func (m AccountBalanceUpdateFixedBuilder) TransactionDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 344, 336))
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
func (m AccountBalanceUpdateFixed) MarginRequirementFull() float64 {
	return message.Float64Fixed(m.Unsafe(), 352, 344)
}

// MarginRequirementFull
func (m AccountBalanceUpdateFixedBuilder) MarginRequirementFull() float64 {
	return message.Float64Fixed(m.Unsafe(), 352, 344)
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
func (m AccountBalanceUpdateFixed) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64Fixed(m.Unsafe(), 360, 352)
}

// MarginRequirementFullPositionsOnly
func (m AccountBalanceUpdateFixedBuilder) MarginRequirementFullPositionsOnly() float64 {
	return message.Float64Fixed(m.Unsafe(), 360, 352)
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
func (m AccountBalanceUpdateFixed) PeakMarginRequirement() float64 {
	return message.Float64Fixed(m.Unsafe(), 368, 360)
}

// PeakMarginRequirement
func (m AccountBalanceUpdateFixedBuilder) PeakMarginRequirement() float64 {
	return message.Float64Fixed(m.Unsafe(), 368, 360)
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
func (m AccountBalanceUpdateBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
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
func (m AccountBalanceUpdateBuilder) SetCashBalance(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 24, 16, value)
}

// SetCashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdatePointerBuilder) SetCashBalance(value float64) {
	message.SetFloat64VLS(m.Ptr, 24, 16, value)
}

// SetBalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateBuilder) SetBalanceAvailableForNewPositions(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 32, 24, value)
}

// SetBalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdatePointerBuilder) SetBalanceAvailableForNewPositions(value float64) {
	message.SetFloat64VLS(m.Ptr, 32, 24, value)
}

// SetAccountCurrency ISO Currency Code for the cash values in this message.
func (m *AccountBalanceUpdateBuilder) SetAccountCurrency(value string) {
	message.SetStringVLS(&m.VLSBuilder, 36, 32, value)
}

// SetAccountCurrency ISO Currency Code for the cash values in this message.
func (m *AccountBalanceUpdatePointerBuilder) SetAccountCurrency(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 36, 32, value)
}

// SetTradeAccount The trade account identifier for the Account Balance information.
func (m *AccountBalanceUpdateBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 40, 36, value)
}

// SetTradeAccount The trade account identifier for the Account Balance information.
func (m *AccountBalanceUpdatePointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 40, 36, value)
}

// SetSecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdateBuilder) SetSecuritiesValue(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 48, 40, value)
}

// SetSecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdatePointerBuilder) SetSecuritiesValue(value float64) {
	message.SetFloat64VLS(m.Ptr, 48, 40, value)
}

// SetMarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 56, 48, value)
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
func (m AccountBalanceUpdateBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32VLS(m.Unsafe(), 60, 56, value)
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
func (m AccountBalanceUpdateBuilder) SetMessageNumber(value int32) {
	message.SetInt32VLS(m.Unsafe(), 64, 60, value)
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
func (m AccountBalanceUpdateBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 65, 64, value)
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
func (m AccountBalanceUpdateBuilder) SetUnsolicited(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 66, 65, value)
}

// SetUnsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdatePointerBuilder) SetUnsolicited(value uint8) {
	message.SetUint8VLS(m.Ptr, 66, 65, value)
}

// SetOpenPositionsProfitLoss
func (m AccountBalanceUpdateBuilder) SetOpenPositionsProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 80, 72, value)
}

// SetOpenPositionsProfitLoss
func (m AccountBalanceUpdatePointerBuilder) SetOpenPositionsProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 80, 72, value)
}

// SetDailyProfitLoss
func (m AccountBalanceUpdateBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 88, 80, value)
}

// SetDailyProfitLoss
func (m AccountBalanceUpdatePointerBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64VLS(m.Ptr, 88, 80, value)
}

// SetInfoText
func (m *AccountBalanceUpdateBuilder) SetInfoText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 92, 88, value)
}

// SetInfoText
func (m *AccountBalanceUpdatePointerBuilder) SetInfoText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 92, 88, value)
}

// SetTransactionIdentifier
func (m AccountBalanceUpdateBuilder) SetTransactionIdentifier(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 104, 96, value)
}

// SetTransactionIdentifier
func (m AccountBalanceUpdatePointerBuilder) SetTransactionIdentifier(value uint64) {
	message.SetUint64VLS(m.Ptr, 104, 96, value)
}

// SetDailyNetLossLimit
func (m AccountBalanceUpdateBuilder) SetDailyNetLossLimit(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 112, 104, value)
}

// SetDailyNetLossLimit
func (m AccountBalanceUpdatePointerBuilder) SetDailyNetLossLimit(value float64) {
	message.SetFloat64VLS(m.Ptr, 112, 104, value)
}

// SetTrailingAccountValueToLimitPositions
func (m AccountBalanceUpdateBuilder) SetTrailingAccountValueToLimitPositions(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 120, 112, value)
}

// SetTrailingAccountValueToLimitPositions
func (m AccountBalanceUpdatePointerBuilder) SetTrailingAccountValueToLimitPositions(value float64) {
	message.SetFloat64VLS(m.Ptr, 120, 112, value)
}

// SetDailyNetLossLimitReached
func (m AccountBalanceUpdateBuilder) SetDailyNetLossLimitReached(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 121, 120, value)
}

// SetDailyNetLossLimitReached
func (m AccountBalanceUpdatePointerBuilder) SetDailyNetLossLimitReached(value uint8) {
	message.SetUint8VLS(m.Ptr, 121, 120, value)
}

// SetIsUnderRequiredMargin
func (m AccountBalanceUpdateBuilder) SetIsUnderRequiredMargin(value bool) {
	message.SetBoolVLS(m.Unsafe(), 122, 121, value)
}

// SetIsUnderRequiredMargin
func (m AccountBalanceUpdatePointerBuilder) SetIsUnderRequiredMargin(value bool) {
	message.SetBoolVLS(m.Ptr, 122, 121, value)
}

// SetClosePositionsAtEndOfDay
func (m AccountBalanceUpdateBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 123, 122, value)
}

// SetClosePositionsAtEndOfDay
func (m AccountBalanceUpdatePointerBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8VLS(m.Ptr, 123, 122, value)
}

// SetTradingIsDisabled
func (m AccountBalanceUpdateBuilder) SetTradingIsDisabled(value bool) {
	message.SetBoolVLS(m.Unsafe(), 124, 123, value)
}

// SetTradingIsDisabled
func (m AccountBalanceUpdatePointerBuilder) SetTradingIsDisabled(value bool) {
	message.SetBoolVLS(m.Ptr, 124, 123, value)
}

// SetDescription
func (m *AccountBalanceUpdateBuilder) SetDescription(value string) {
	message.SetStringVLS(&m.VLSBuilder, 128, 124, value)
}

// SetDescription
func (m *AccountBalanceUpdatePointerBuilder) SetDescription(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 128, 124, value)
}

// SetIsUnderRequiredAccountValue
func (m AccountBalanceUpdateBuilder) SetIsUnderRequiredAccountValue(value bool) {
	message.SetBoolVLS(m.Unsafe(), 129, 128, value)
}

// SetIsUnderRequiredAccountValue
func (m AccountBalanceUpdatePointerBuilder) SetIsUnderRequiredAccountValue(value bool) {
	message.SetBoolVLS(m.Ptr, 129, 128, value)
}

// SetTransactionDateTime
func (m AccountBalanceUpdateBuilder) SetTransactionDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 144, 136, int64(value))
}

// SetTransactionDateTime
func (m AccountBalanceUpdatePointerBuilder) SetTransactionDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64VLS(m.Ptr, 144, 136, int64(value))
}

// SetMarginRequirementFull
func (m AccountBalanceUpdateBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 152, 144, value)
}

// SetMarginRequirementFull
func (m AccountBalanceUpdatePointerBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64VLS(m.Ptr, 152, 144, value)
}

// SetMarginRequirementFullPositionsOnly
func (m AccountBalanceUpdateBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 160, 152, value)
}

// SetMarginRequirementFullPositionsOnly
func (m AccountBalanceUpdatePointerBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64VLS(m.Ptr, 160, 152, value)
}

// SetPeakMarginRequirement
func (m AccountBalanceUpdateBuilder) SetPeakMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 168, 160, value)
}

// SetPeakMarginRequirement
func (m AccountBalanceUpdatePointerBuilder) SetPeakMarginRequirement(value float64) {
	message.SetFloat64VLS(m.Ptr, 168, 160, value)
}

// SetRequestID This is the RequestID which was set in the AccountBalanceRequest that
// this message is in response to.
//
// In the case when this is a periodic unsolicited Account Balance update,
// RequestID must be set to 0, the default.
func (m AccountBalanceUpdateFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
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
func (m AccountBalanceUpdateFixedBuilder) SetCashBalance(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 16, 8, value)
}

// SetCashBalance The current cash balance for the account in the currency specified by
// the AccountCurrency field.
func (m AccountBalanceUpdateFixedPointerBuilder) SetCashBalance(value float64) {
	message.SetFloat64Fixed(m.Ptr, 16, 8, value)
}

// SetBalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedBuilder) SetBalanceAvailableForNewPositions(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 24, 16, value)
}

// SetBalanceAvailableForNewPositions The CashBalance minus the cash required to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedPointerBuilder) SetBalanceAvailableForNewPositions(value float64) {
	message.SetFloat64Fixed(m.Ptr, 24, 16, value)
}

// SetAccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdateFixedBuilder) SetAccountCurrency(value string) {
	message.SetStringFixed(m.Unsafe(), 32, 24, value)
}

// SetAccountCurrency ISO Currency Code for the cash values in this message.
func (m AccountBalanceUpdateFixedPointerBuilder) SetAccountCurrency(value string) {
	message.SetStringFixed(m.Ptr, 32, 24, value)
}

// SetTradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdateFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 64, 32, value)
}

// SetTradeAccount The trade account identifier for the Account Balance information.
func (m AccountBalanceUpdateFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 64, 32, value)
}

// SetSecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdateFixedBuilder) SetSecuritiesValue(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 72, 64, value)
}

// SetSecuritiesValue This is the value of all cash and securities as a cash value.
func (m AccountBalanceUpdateFixedPointerBuilder) SetSecuritiesValue(value float64) {
	message.SetFloat64Fixed(m.Ptr, 72, 64, value)
}

// SetMarginRequirement This is the current cash requirement to maintain securities on margin
// the Trade Account currently has.
func (m AccountBalanceUpdateFixedBuilder) SetMarginRequirement(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 80, 72, value)
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
func (m AccountBalanceUpdateFixedBuilder) SetTotalNumberMessages(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 84, 80, value)
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
func (m AccountBalanceUpdateFixedBuilder) SetMessageNumber(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 88, 84, value)
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
func (m AccountBalanceUpdateFixedBuilder) SetNoAccountBalances(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 89, 88, value)
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
func (m AccountBalanceUpdateFixedBuilder) SetUnsolicited(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 90, 89, value)
}

// SetUnsolicited Set to 1 to indicate this is an unsolicited Account Balance Update message.
// In other words, it is a real-time Account Balance Update message which
// is not an initial response to an AccountBalanceRequest message.
func (m AccountBalanceUpdateFixedPointerBuilder) SetUnsolicited(value uint8) {
	message.SetUint8Fixed(m.Ptr, 90, 89, value)
}

// SetOpenPositionsProfitLoss
func (m AccountBalanceUpdateFixedBuilder) SetOpenPositionsProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 104, 96, value)
}

// SetOpenPositionsProfitLoss
func (m AccountBalanceUpdateFixedPointerBuilder) SetOpenPositionsProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Ptr, 104, 96, value)
}

// SetDailyProfitLoss
func (m AccountBalanceUpdateFixedBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 112, 104, value)
}

// SetDailyProfitLoss
func (m AccountBalanceUpdateFixedPointerBuilder) SetDailyProfitLoss(value float64) {
	message.SetFloat64Fixed(m.Ptr, 112, 104, value)
}

// SetInfoText
func (m AccountBalanceUpdateFixedBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Unsafe(), 208, 112, value)
}

// SetInfoText
func (m AccountBalanceUpdateFixedPointerBuilder) SetInfoText(value string) {
	message.SetStringFixed(m.Ptr, 208, 112, value)
}

// SetTransactionIdentifier
func (m AccountBalanceUpdateFixedBuilder) SetTransactionIdentifier(value uint64) {
	message.SetUint64Fixed(m.Unsafe(), 216, 208, value)
}

// SetTransactionIdentifier
func (m AccountBalanceUpdateFixedPointerBuilder) SetTransactionIdentifier(value uint64) {
	message.SetUint64Fixed(m.Ptr, 216, 208, value)
}

// SetDailyNetLossLimit
func (m AccountBalanceUpdateFixedBuilder) SetDailyNetLossLimit(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 224, 216, value)
}

// SetDailyNetLossLimit
func (m AccountBalanceUpdateFixedPointerBuilder) SetDailyNetLossLimit(value float64) {
	message.SetFloat64Fixed(m.Ptr, 224, 216, value)
}

// SetTrailingAccountValueToLimitPositions
func (m AccountBalanceUpdateFixedBuilder) SetTrailingAccountValueToLimitPositions(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 232, 224, value)
}

// SetTrailingAccountValueToLimitPositions
func (m AccountBalanceUpdateFixedPointerBuilder) SetTrailingAccountValueToLimitPositions(value float64) {
	message.SetFloat64Fixed(m.Ptr, 232, 224, value)
}

// SetDailyNetLossLimitReached
func (m AccountBalanceUpdateFixedBuilder) SetDailyNetLossLimitReached(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 233, 232, value)
}

// SetDailyNetLossLimitReached
func (m AccountBalanceUpdateFixedPointerBuilder) SetDailyNetLossLimitReached(value uint8) {
	message.SetUint8Fixed(m.Ptr, 233, 232, value)
}

// SetIsUnderRequiredMargin
func (m AccountBalanceUpdateFixedBuilder) SetIsUnderRequiredMargin(value bool) {
	message.SetBoolFixed(m.Unsafe(), 234, 233, value)
}

// SetIsUnderRequiredMargin
func (m AccountBalanceUpdateFixedPointerBuilder) SetIsUnderRequiredMargin(value bool) {
	message.SetBoolFixed(m.Ptr, 234, 233, value)
}

// SetClosePositionsAtEndOfDay
func (m AccountBalanceUpdateFixedBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 235, 234, value)
}

// SetClosePositionsAtEndOfDay
func (m AccountBalanceUpdateFixedPointerBuilder) SetClosePositionsAtEndOfDay(value uint8) {
	message.SetUint8Fixed(m.Ptr, 235, 234, value)
}

// SetTradingIsDisabled
func (m AccountBalanceUpdateFixedBuilder) SetTradingIsDisabled(value bool) {
	message.SetBoolFixed(m.Unsafe(), 236, 235, value)
}

// SetTradingIsDisabled
func (m AccountBalanceUpdateFixedPointerBuilder) SetTradingIsDisabled(value bool) {
	message.SetBoolFixed(m.Ptr, 236, 235, value)
}

// SetDescription
func (m AccountBalanceUpdateFixedBuilder) SetDescription(value string) {
	message.SetStringFixed(m.Unsafe(), 332, 236, value)
}

// SetDescription
func (m AccountBalanceUpdateFixedPointerBuilder) SetDescription(value string) {
	message.SetStringFixed(m.Ptr, 332, 236, value)
}

// SetIsUnderRequiredAccountValue
func (m AccountBalanceUpdateFixedBuilder) SetIsUnderRequiredAccountValue(value bool) {
	message.SetBoolFixed(m.Unsafe(), 333, 332, value)
}

// SetIsUnderRequiredAccountValue
func (m AccountBalanceUpdateFixedPointerBuilder) SetIsUnderRequiredAccountValue(value bool) {
	message.SetBoolFixed(m.Ptr, 333, 332, value)
}

// SetTransactionDateTime
func (m AccountBalanceUpdateFixedBuilder) SetTransactionDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 344, 336, int64(value))
}

// SetTransactionDateTime
func (m AccountBalanceUpdateFixedPointerBuilder) SetTransactionDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 344, 336, int64(value))
}

// SetMarginRequirementFull
func (m AccountBalanceUpdateFixedBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 352, 344, value)
}

// SetMarginRequirementFull
func (m AccountBalanceUpdateFixedPointerBuilder) SetMarginRequirementFull(value float64) {
	message.SetFloat64Fixed(m.Ptr, 352, 344, value)
}

// SetMarginRequirementFullPositionsOnly
func (m AccountBalanceUpdateFixedBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 360, 352, value)
}

// SetMarginRequirementFullPositionsOnly
func (m AccountBalanceUpdateFixedPointerBuilder) SetMarginRequirementFullPositionsOnly(value float64) {
	message.SetFloat64Fixed(m.Ptr, 360, 352, value)
}

// SetPeakMarginRequirement
func (m AccountBalanceUpdateFixedBuilder) SetPeakMarginRequirement(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 368, 360, value)
}

// SetPeakMarginRequirement
func (m AccountBalanceUpdateFixedPointerBuilder) SetPeakMarginRequirement(value float64) {
	message.SetFloat64Fixed(m.Ptr, 368, 360, value)
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
