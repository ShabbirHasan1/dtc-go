package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                                                  = ReplayChartDataFixedSize  (426)
//     Type                                                  = REPLAY_CHART_DATA  (10104)
//     RequestID                                             = 0
//     Symbol                                                = ""
//     TradeAccount                                          = ""
//     TimeZone                                              = ""
//     StartDateTimeForInitialData                           = 0
//     StartDateTime                                         = 0
//     StopDateTime                                          = 0
//     SessionBeginTimeInSeconds                             = 0
//     SessionEndTimeInSeconds                               = 0
//     ReplaySpeed                                           = 1.000000
//     BarTimeInSeconds                                      = 0
//     PauseReplayAfterInitialDataSent                       = 0
//     UseSavedPriorState                                    = 0
//     SymbolVolatility                                      = 0
//     InterestRate                                          = 0
//     NumberOfOrdersToTriggerFinishToStopDateTime           = 0
//     MaximumNumberOfOrdersPerReplaySession                 = 0
//     NumberOfDaysForInitialDataFromBeforeLastSavedDateTime = 0
//     SubAccountIdentifier                                  = 0
//     OptionsPriceSendIntervalInSeconds                     = 0
// }
var _ReplayChartDataFixedDefault = []byte{170, 1, 120, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const ReplayChartDataFixedSize = 426

type ReplayChartDataFixed struct {
	message.Fixed
}

type ReplayChartDataFixedBuilder struct {
	message.Fixed
}

func NewReplayChartDataFixedFrom(b []byte) ReplayChartDataFixed {
	return ReplayChartDataFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapReplayChartDataFixed(b []byte) ReplayChartDataFixed {
	return ReplayChartDataFixed{Fixed: message.WrapFixed(b)}
}

func NewReplayChartDataFixed() ReplayChartDataFixedBuilder {
	a := ReplayChartDataFixedBuilder{message.NewFixed(426)}
	a.Unsafe().SetBytes(0, _ReplayChartDataFixedDefault)
	return a
}

// Clear
// {
//     Size                                                  = ReplayChartDataFixedSize  (426)
//     Type                                                  = REPLAY_CHART_DATA  (10104)
//     RequestID                                             = 0
//     Symbol                                                = ""
//     TradeAccount                                          = ""
//     TimeZone                                              = ""
//     StartDateTimeForInitialData                           = 0
//     StartDateTime                                         = 0
//     StopDateTime                                          = 0
//     SessionBeginTimeInSeconds                             = 0
//     SessionEndTimeInSeconds                               = 0
//     ReplaySpeed                                           = 1.000000
//     BarTimeInSeconds                                      = 0
//     PauseReplayAfterInitialDataSent                       = 0
//     UseSavedPriorState                                    = 0
//     SymbolVolatility                                      = 0
//     InterestRate                                          = 0
//     NumberOfOrdersToTriggerFinishToStopDateTime           = 0
//     MaximumNumberOfOrdersPerReplaySession                 = 0
//     NumberOfDaysForInitialDataFromBeforeLastSavedDateTime = 0
//     SubAccountIdentifier                                  = 0
//     OptionsPriceSendIntervalInSeconds                     = 0
// }
func (m ReplayChartDataFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ReplayChartDataFixedDefault)
}

// ToBuilder
func (m ReplayChartDataFixed) ToBuilder() ReplayChartDataFixedBuilder {
	return ReplayChartDataFixedBuilder{m.Fixed}
}

// Finish
func (m ReplayChartDataFixedBuilder) Finish() ReplayChartDataFixed {
	return ReplayChartDataFixed{m.Fixed}
}

// RequestID
func (m ReplayChartDataFixed) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m ReplayChartDataFixedBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Symbol
func (m ReplayChartDataFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol
func (m ReplayChartDataFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// TradeAccount
func (m ReplayChartDataFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 104, 72)
}

// TradeAccount
func (m ReplayChartDataFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 104, 72)
}

// TimeZone
func (m ReplayChartDataFixed) TimeZone() string {
	return message.StringFixed(m.Unsafe(), 360, 104)
}

// TimeZone
func (m ReplayChartDataFixedBuilder) TimeZone() string {
	return message.StringFixed(m.Unsafe(), 360, 104)
}

// StartDateTimeForInitialData
func (m ReplayChartDataFixed) StartDateTimeForInitialData() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 368, 360))
}

// StartDateTimeForInitialData
func (m ReplayChartDataFixedBuilder) StartDateTimeForInitialData() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 368, 360))
}

// StartDateTime
func (m ReplayChartDataFixed) StartDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 376, 368))
}

// StartDateTime
func (m ReplayChartDataFixedBuilder) StartDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 376, 368))
}

// StopDateTime
func (m ReplayChartDataFixed) StopDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 384, 376))
}

// StopDateTime
func (m ReplayChartDataFixedBuilder) StopDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 384, 376))
}

// SessionBeginTimeInSeconds
func (m ReplayChartDataFixed) SessionBeginTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 386, 384)
}

// SessionBeginTimeInSeconds
func (m ReplayChartDataFixedBuilder) SessionBeginTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 386, 384)
}

// SessionEndTimeInSeconds
func (m ReplayChartDataFixed) SessionEndTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 388, 386)
}

// SessionEndTimeInSeconds
func (m ReplayChartDataFixedBuilder) SessionEndTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 388, 386)
}

// ReplaySpeed
func (m ReplayChartDataFixed) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Unsafe(), 392, 388)
}

// ReplaySpeed
func (m ReplayChartDataFixedBuilder) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Unsafe(), 392, 388)
}

// BarTimeInSeconds
func (m ReplayChartDataFixed) BarTimeInSeconds() int32 {
	return message.Int32Fixed(m.Unsafe(), 396, 392)
}

// BarTimeInSeconds
func (m ReplayChartDataFixedBuilder) BarTimeInSeconds() int32 {
	return message.Int32Fixed(m.Unsafe(), 396, 392)
}

// PauseReplayAfterInitialDataSent
func (m ReplayChartDataFixed) PauseReplayAfterInitialDataSent() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 397, 396)
}

// PauseReplayAfterInitialDataSent
func (m ReplayChartDataFixedBuilder) PauseReplayAfterInitialDataSent() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 397, 396)
}

// UseSavedPriorState
func (m ReplayChartDataFixed) UseSavedPriorState() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 398, 397)
}

// UseSavedPriorState
func (m ReplayChartDataFixedBuilder) UseSavedPriorState() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 398, 397)
}

// SymbolVolatility
func (m ReplayChartDataFixed) SymbolVolatility() float32 {
	return message.Float32Fixed(m.Unsafe(), 402, 398)
}

// SymbolVolatility
func (m ReplayChartDataFixedBuilder) SymbolVolatility() float32 {
	return message.Float32Fixed(m.Unsafe(), 402, 398)
}

// InterestRate
func (m ReplayChartDataFixed) InterestRate() float32 {
	return message.Float32Fixed(m.Unsafe(), 406, 402)
}

// InterestRate
func (m ReplayChartDataFixedBuilder) InterestRate() float32 {
	return message.Float32Fixed(m.Unsafe(), 406, 402)
}

// NumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataFixed) NumberOfOrdersToTriggerFinishToStopDateTime() int32 {
	return message.Int32Fixed(m.Unsafe(), 410, 406)
}

// NumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataFixedBuilder) NumberOfOrdersToTriggerFinishToStopDateTime() int32 {
	return message.Int32Fixed(m.Unsafe(), 410, 406)
}

// MaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataFixed) MaximumNumberOfOrdersPerReplaySession() int32 {
	return message.Int32Fixed(m.Unsafe(), 414, 410)
}

// MaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataFixedBuilder) MaximumNumberOfOrdersPerReplaySession() int32 {
	return message.Int32Fixed(m.Unsafe(), 414, 410)
}

// NumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataFixed) NumberOfDaysForInitialDataFromBeforeLastSavedDateTime() int32 {
	return message.Int32Fixed(m.Unsafe(), 418, 414)
}

// NumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataFixedBuilder) NumberOfDaysForInitialDataFromBeforeLastSavedDateTime() int32 {
	return message.Int32Fixed(m.Unsafe(), 418, 414)
}

// SubAccountIdentifier
func (m ReplayChartDataFixed) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 422, 418)
}

// SubAccountIdentifier
func (m ReplayChartDataFixedBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 422, 418)
}

// OptionsPriceSendIntervalInSeconds
func (m ReplayChartDataFixed) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32Fixed(m.Unsafe(), 426, 422)
}

// OptionsPriceSendIntervalInSeconds
func (m ReplayChartDataFixedBuilder) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32Fixed(m.Unsafe(), 426, 422)
}

// SetRequestID
func (m ReplayChartDataFixedBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbol
func (m ReplayChartDataFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
}

// SetTradeAccount
func (m ReplayChartDataFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 72, value)
}

// SetTimeZone
func (m ReplayChartDataFixedBuilder) SetTimeZone(value string) {
	message.SetStringFixed(m.Unsafe(), 360, 104, value)
}

// SetStartDateTimeForInitialData
func (m ReplayChartDataFixedBuilder) SetStartDateTimeForInitialData(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 368, 360, int64(value))
}

// SetStartDateTime
func (m ReplayChartDataFixedBuilder) SetStartDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 376, 368, int64(value))
}

// SetStopDateTime
func (m ReplayChartDataFixedBuilder) SetStopDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 384, 376, int64(value))
}

// SetSessionBeginTimeInSeconds
func (m ReplayChartDataFixedBuilder) SetSessionBeginTimeInSeconds(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 386, 384, value)
}

// SetSessionEndTimeInSeconds
func (m ReplayChartDataFixedBuilder) SetSessionEndTimeInSeconds(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 388, 386, value)
}

// SetReplaySpeed
func (m ReplayChartDataFixedBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 392, 388, value)
}

// SetBarTimeInSeconds
func (m ReplayChartDataFixedBuilder) SetBarTimeInSeconds(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 396, 392, value)
}

// SetPauseReplayAfterInitialDataSent
func (m ReplayChartDataFixedBuilder) SetPauseReplayAfterInitialDataSent(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 397, 396, value)
}

// SetUseSavedPriorState
func (m ReplayChartDataFixedBuilder) SetUseSavedPriorState(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 398, 397, value)
}

// SetSymbolVolatility
func (m ReplayChartDataFixedBuilder) SetSymbolVolatility(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 402, 398, value)
}

// SetInterestRate
func (m ReplayChartDataFixedBuilder) SetInterestRate(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 406, 402, value)
}

// SetNumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataFixedBuilder) SetNumberOfOrdersToTriggerFinishToStopDateTime(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 410, 406, value)
}

// SetMaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataFixedBuilder) SetMaximumNumberOfOrdersPerReplaySession(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 414, 410, value)
}

// SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataFixedBuilder) SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 418, 414, value)
}

// SetSubAccountIdentifier
func (m ReplayChartDataFixedBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 422, 418, value)
}

// SetOptionsPriceSendIntervalInSeconds
func (m ReplayChartDataFixedBuilder) SetOptionsPriceSendIntervalInSeconds(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 426, 422, value)
}

// Copy
func (m ReplayChartDataFixed) Copy(to ReplayChartDataFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetTimeZone(m.TimeZone())
	to.SetStartDateTimeForInitialData(m.StartDateTimeForInitialData())
	to.SetStartDateTime(m.StartDateTime())
	to.SetStopDateTime(m.StopDateTime())
	to.SetSessionBeginTimeInSeconds(m.SessionBeginTimeInSeconds())
	to.SetSessionEndTimeInSeconds(m.SessionEndTimeInSeconds())
	to.SetReplaySpeed(m.ReplaySpeed())
	to.SetBarTimeInSeconds(m.BarTimeInSeconds())
	to.SetPauseReplayAfterInitialDataSent(m.PauseReplayAfterInitialDataSent())
	to.SetUseSavedPriorState(m.UseSavedPriorState())
	to.SetSymbolVolatility(m.SymbolVolatility())
	to.SetInterestRate(m.InterestRate())
	to.SetNumberOfOrdersToTriggerFinishToStopDateTime(m.NumberOfOrdersToTriggerFinishToStopDateTime())
	to.SetMaximumNumberOfOrdersPerReplaySession(m.MaximumNumberOfOrdersPerReplaySession())
	to.SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(m.NumberOfDaysForInitialDataFromBeforeLastSavedDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetOptionsPriceSendIntervalInSeconds(m.OptionsPriceSendIntervalInSeconds())
}

// Copy
func (m ReplayChartDataFixedBuilder) Copy(to ReplayChartDataFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetTimeZone(m.TimeZone())
	to.SetStartDateTimeForInitialData(m.StartDateTimeForInitialData())
	to.SetStartDateTime(m.StartDateTime())
	to.SetStopDateTime(m.StopDateTime())
	to.SetSessionBeginTimeInSeconds(m.SessionBeginTimeInSeconds())
	to.SetSessionEndTimeInSeconds(m.SessionEndTimeInSeconds())
	to.SetReplaySpeed(m.ReplaySpeed())
	to.SetBarTimeInSeconds(m.BarTimeInSeconds())
	to.SetPauseReplayAfterInitialDataSent(m.PauseReplayAfterInitialDataSent())
	to.SetUseSavedPriorState(m.UseSavedPriorState())
	to.SetSymbolVolatility(m.SymbolVolatility())
	to.SetInterestRate(m.InterestRate())
	to.SetNumberOfOrdersToTriggerFinishToStopDateTime(m.NumberOfOrdersToTriggerFinishToStopDateTime())
	to.SetMaximumNumberOfOrdersPerReplaySession(m.MaximumNumberOfOrdersPerReplaySession())
	to.SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(m.NumberOfDaysForInitialDataFromBeforeLastSavedDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetOptionsPriceSendIntervalInSeconds(m.OptionsPriceSendIntervalInSeconds())
}

// CopyPointer
func (m ReplayChartDataFixed) CopyPointer(to ReplayChartDataFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetTimeZone(m.TimeZone())
	to.SetStartDateTimeForInitialData(m.StartDateTimeForInitialData())
	to.SetStartDateTime(m.StartDateTime())
	to.SetStopDateTime(m.StopDateTime())
	to.SetSessionBeginTimeInSeconds(m.SessionBeginTimeInSeconds())
	to.SetSessionEndTimeInSeconds(m.SessionEndTimeInSeconds())
	to.SetReplaySpeed(m.ReplaySpeed())
	to.SetBarTimeInSeconds(m.BarTimeInSeconds())
	to.SetPauseReplayAfterInitialDataSent(m.PauseReplayAfterInitialDataSent())
	to.SetUseSavedPriorState(m.UseSavedPriorState())
	to.SetSymbolVolatility(m.SymbolVolatility())
	to.SetInterestRate(m.InterestRate())
	to.SetNumberOfOrdersToTriggerFinishToStopDateTime(m.NumberOfOrdersToTriggerFinishToStopDateTime())
	to.SetMaximumNumberOfOrdersPerReplaySession(m.MaximumNumberOfOrdersPerReplaySession())
	to.SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(m.NumberOfDaysForInitialDataFromBeforeLastSavedDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetOptionsPriceSendIntervalInSeconds(m.OptionsPriceSendIntervalInSeconds())
}

// CopyPointer
func (m ReplayChartDataFixedBuilder) CopyPointer(to ReplayChartDataFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetTimeZone(m.TimeZone())
	to.SetStartDateTimeForInitialData(m.StartDateTimeForInitialData())
	to.SetStartDateTime(m.StartDateTime())
	to.SetStopDateTime(m.StopDateTime())
	to.SetSessionBeginTimeInSeconds(m.SessionBeginTimeInSeconds())
	to.SetSessionEndTimeInSeconds(m.SessionEndTimeInSeconds())
	to.SetReplaySpeed(m.ReplaySpeed())
	to.SetBarTimeInSeconds(m.BarTimeInSeconds())
	to.SetPauseReplayAfterInitialDataSent(m.PauseReplayAfterInitialDataSent())
	to.SetUseSavedPriorState(m.UseSavedPriorState())
	to.SetSymbolVolatility(m.SymbolVolatility())
	to.SetInterestRate(m.InterestRate())
	to.SetNumberOfOrdersToTriggerFinishToStopDateTime(m.NumberOfOrdersToTriggerFinishToStopDateTime())
	to.SetMaximumNumberOfOrdersPerReplaySession(m.MaximumNumberOfOrdersPerReplaySession())
	to.SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(m.NumberOfDaysForInitialDataFromBeforeLastSavedDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetOptionsPriceSendIntervalInSeconds(m.OptionsPriceSendIntervalInSeconds())
}

// CopyToPointer
func (m ReplayChartDataFixed) CopyToPointer(to ReplayChartDataPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetTimeZone(m.TimeZone())
	to.SetStartDateTimeForInitialData(m.StartDateTimeForInitialData())
	to.SetStartDateTime(m.StartDateTime())
	to.SetStopDateTime(m.StopDateTime())
	to.SetSessionBeginTimeInSeconds(m.SessionBeginTimeInSeconds())
	to.SetSessionEndTimeInSeconds(m.SessionEndTimeInSeconds())
	to.SetReplaySpeed(m.ReplaySpeed())
	to.SetBarTimeInSeconds(m.BarTimeInSeconds())
	to.SetPauseReplayAfterInitialDataSent(m.PauseReplayAfterInitialDataSent())
	to.SetUseSavedPriorState(m.UseSavedPriorState())
	to.SetSymbolVolatility(m.SymbolVolatility())
	to.SetInterestRate(m.InterestRate())
	to.SetNumberOfOrdersToTriggerFinishToStopDateTime(m.NumberOfOrdersToTriggerFinishToStopDateTime())
	to.SetMaximumNumberOfOrdersPerReplaySession(m.MaximumNumberOfOrdersPerReplaySession())
	to.SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(m.NumberOfDaysForInitialDataFromBeforeLastSavedDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetOptionsPriceSendIntervalInSeconds(m.OptionsPriceSendIntervalInSeconds())
}

// CopyToPointer
func (m ReplayChartDataFixedBuilder) CopyToPointer(to ReplayChartDataPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetTimeZone(m.TimeZone())
	to.SetStartDateTimeForInitialData(m.StartDateTimeForInitialData())
	to.SetStartDateTime(m.StartDateTime())
	to.SetStopDateTime(m.StopDateTime())
	to.SetSessionBeginTimeInSeconds(m.SessionBeginTimeInSeconds())
	to.SetSessionEndTimeInSeconds(m.SessionEndTimeInSeconds())
	to.SetReplaySpeed(m.ReplaySpeed())
	to.SetBarTimeInSeconds(m.BarTimeInSeconds())
	to.SetPauseReplayAfterInitialDataSent(m.PauseReplayAfterInitialDataSent())
	to.SetUseSavedPriorState(m.UseSavedPriorState())
	to.SetSymbolVolatility(m.SymbolVolatility())
	to.SetInterestRate(m.InterestRate())
	to.SetNumberOfOrdersToTriggerFinishToStopDateTime(m.NumberOfOrdersToTriggerFinishToStopDateTime())
	to.SetMaximumNumberOfOrdersPerReplaySession(m.MaximumNumberOfOrdersPerReplaySession())
	to.SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(m.NumberOfDaysForInitialDataFromBeforeLastSavedDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetOptionsPriceSendIntervalInSeconds(m.OptionsPriceSendIntervalInSeconds())
}

// CopyTo
func (m ReplayChartDataFixed) CopyTo(to ReplayChartDataBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetTimeZone(m.TimeZone())
	to.SetStartDateTimeForInitialData(m.StartDateTimeForInitialData())
	to.SetStartDateTime(m.StartDateTime())
	to.SetStopDateTime(m.StopDateTime())
	to.SetSessionBeginTimeInSeconds(m.SessionBeginTimeInSeconds())
	to.SetSessionEndTimeInSeconds(m.SessionEndTimeInSeconds())
	to.SetReplaySpeed(m.ReplaySpeed())
	to.SetBarTimeInSeconds(m.BarTimeInSeconds())
	to.SetPauseReplayAfterInitialDataSent(m.PauseReplayAfterInitialDataSent())
	to.SetUseSavedPriorState(m.UseSavedPriorState())
	to.SetSymbolVolatility(m.SymbolVolatility())
	to.SetInterestRate(m.InterestRate())
	to.SetNumberOfOrdersToTriggerFinishToStopDateTime(m.NumberOfOrdersToTriggerFinishToStopDateTime())
	to.SetMaximumNumberOfOrdersPerReplaySession(m.MaximumNumberOfOrdersPerReplaySession())
	to.SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(m.NumberOfDaysForInitialDataFromBeforeLastSavedDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetOptionsPriceSendIntervalInSeconds(m.OptionsPriceSendIntervalInSeconds())
}

// CopyTo
func (m ReplayChartDataFixedBuilder) CopyTo(to ReplayChartDataBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetTradeAccount(m.TradeAccount())
	to.SetTimeZone(m.TimeZone())
	to.SetStartDateTimeForInitialData(m.StartDateTimeForInitialData())
	to.SetStartDateTime(m.StartDateTime())
	to.SetStopDateTime(m.StopDateTime())
	to.SetSessionBeginTimeInSeconds(m.SessionBeginTimeInSeconds())
	to.SetSessionEndTimeInSeconds(m.SessionEndTimeInSeconds())
	to.SetReplaySpeed(m.ReplaySpeed())
	to.SetBarTimeInSeconds(m.BarTimeInSeconds())
	to.SetPauseReplayAfterInitialDataSent(m.PauseReplayAfterInitialDataSent())
	to.SetUseSavedPriorState(m.UseSavedPriorState())
	to.SetSymbolVolatility(m.SymbolVolatility())
	to.SetInterestRate(m.InterestRate())
	to.SetNumberOfOrdersToTriggerFinishToStopDateTime(m.NumberOfOrdersToTriggerFinishToStopDateTime())
	to.SetMaximumNumberOfOrdersPerReplaySession(m.MaximumNumberOfOrdersPerReplaySession())
	to.SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(m.NumberOfDaysForInitialDataFromBeforeLastSavedDateTime())
	to.SetSubAccountIdentifier(m.SubAccountIdentifier())
	to.SetOptionsPriceSendIntervalInSeconds(m.OptionsPriceSendIntervalInSeconds())
}
