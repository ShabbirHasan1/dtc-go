package model

import (
	"github.com/moontrade/dtc-go/message"
)

type ReplayChartDataFixedPointer struct {
	message.FixedPointer
}

type ReplayChartDataFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocReplayChartDataFixed() ReplayChartDataFixedPointerBuilder {
	a := ReplayChartDataFixedPointerBuilder{message.AllocFixed(426)}
	a.Ptr.SetBytes(0, _ReplayChartDataFixedDefault)
	return a
}

func AllocReplayChartDataFixedFrom(b []byte) ReplayChartDataFixedPointer {
	return ReplayChartDataFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
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
func (m ReplayChartDataFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ReplayChartDataFixedDefault)
}

// ToBuilder
func (m ReplayChartDataFixedPointer) ToBuilder() ReplayChartDataFixedPointerBuilder {
	return ReplayChartDataFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *ReplayChartDataFixedPointerBuilder) Finish() ReplayChartDataFixedPointer {
	return ReplayChartDataFixedPointer{m.FixedPointer}
}

// RequestID
func (m ReplayChartDataFixedPointer) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m ReplayChartDataFixedPointerBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Symbol
func (m ReplayChartDataFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Symbol
func (m ReplayChartDataFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// TradeAccount
func (m ReplayChartDataFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 104, 72)
}

// TradeAccount
func (m ReplayChartDataFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 104, 72)
}

// TimeZone
func (m ReplayChartDataFixedPointer) TimeZone() string {
	return message.StringFixed(m.Ptr, 360, 104)
}

// TimeZone
func (m ReplayChartDataFixedPointerBuilder) TimeZone() string {
	return message.StringFixed(m.Ptr, 360, 104)
}

// StartDateTimeForInitialData
func (m ReplayChartDataFixedPointer) StartDateTimeForInitialData() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 368, 360))
}

// StartDateTimeForInitialData
func (m ReplayChartDataFixedPointerBuilder) StartDateTimeForInitialData() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 368, 360))
}

// StartDateTime
func (m ReplayChartDataFixedPointer) StartDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 376, 368))
}

// StartDateTime
func (m ReplayChartDataFixedPointerBuilder) StartDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 376, 368))
}

// StopDateTime
func (m ReplayChartDataFixedPointer) StopDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 384, 376))
}

// StopDateTime
func (m ReplayChartDataFixedPointerBuilder) StopDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Ptr, 384, 376))
}

// SessionBeginTimeInSeconds
func (m ReplayChartDataFixedPointer) SessionBeginTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Ptr, 386, 384)
}

// SessionBeginTimeInSeconds
func (m ReplayChartDataFixedPointerBuilder) SessionBeginTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Ptr, 386, 384)
}

// SessionEndTimeInSeconds
func (m ReplayChartDataFixedPointer) SessionEndTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Ptr, 388, 386)
}

// SessionEndTimeInSeconds
func (m ReplayChartDataFixedPointerBuilder) SessionEndTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Ptr, 388, 386)
}

// ReplaySpeed
func (m ReplayChartDataFixedPointer) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Ptr, 392, 388)
}

// ReplaySpeed
func (m ReplayChartDataFixedPointerBuilder) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Ptr, 392, 388)
}

// BarTimeInSeconds
func (m ReplayChartDataFixedPointer) BarTimeInSeconds() int32 {
	return message.Int32Fixed(m.Ptr, 396, 392)
}

// BarTimeInSeconds
func (m ReplayChartDataFixedPointerBuilder) BarTimeInSeconds() int32 {
	return message.Int32Fixed(m.Ptr, 396, 392)
}

// PauseReplayAfterInitialDataSent
func (m ReplayChartDataFixedPointer) PauseReplayAfterInitialDataSent() uint8 {
	return message.Uint8Fixed(m.Ptr, 397, 396)
}

// PauseReplayAfterInitialDataSent
func (m ReplayChartDataFixedPointerBuilder) PauseReplayAfterInitialDataSent() uint8 {
	return message.Uint8Fixed(m.Ptr, 397, 396)
}

// UseSavedPriorState
func (m ReplayChartDataFixedPointer) UseSavedPriorState() uint8 {
	return message.Uint8Fixed(m.Ptr, 398, 397)
}

// UseSavedPriorState
func (m ReplayChartDataFixedPointerBuilder) UseSavedPriorState() uint8 {
	return message.Uint8Fixed(m.Ptr, 398, 397)
}

// SymbolVolatility
func (m ReplayChartDataFixedPointer) SymbolVolatility() float32 {
	return message.Float32Fixed(m.Ptr, 402, 398)
}

// SymbolVolatility
func (m ReplayChartDataFixedPointerBuilder) SymbolVolatility() float32 {
	return message.Float32Fixed(m.Ptr, 402, 398)
}

// InterestRate
func (m ReplayChartDataFixedPointer) InterestRate() float32 {
	return message.Float32Fixed(m.Ptr, 406, 402)
}

// InterestRate
func (m ReplayChartDataFixedPointerBuilder) InterestRate() float32 {
	return message.Float32Fixed(m.Ptr, 406, 402)
}

// NumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataFixedPointer) NumberOfOrdersToTriggerFinishToStopDateTime() int32 {
	return message.Int32Fixed(m.Ptr, 410, 406)
}

// NumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataFixedPointerBuilder) NumberOfOrdersToTriggerFinishToStopDateTime() int32 {
	return message.Int32Fixed(m.Ptr, 410, 406)
}

// MaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataFixedPointer) MaximumNumberOfOrdersPerReplaySession() int32 {
	return message.Int32Fixed(m.Ptr, 414, 410)
}

// MaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataFixedPointerBuilder) MaximumNumberOfOrdersPerReplaySession() int32 {
	return message.Int32Fixed(m.Ptr, 414, 410)
}

// NumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataFixedPointer) NumberOfDaysForInitialDataFromBeforeLastSavedDateTime() int32 {
	return message.Int32Fixed(m.Ptr, 418, 414)
}

// NumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataFixedPointerBuilder) NumberOfDaysForInitialDataFromBeforeLastSavedDateTime() int32 {
	return message.Int32Fixed(m.Ptr, 418, 414)
}

// SubAccountIdentifier
func (m ReplayChartDataFixedPointer) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Ptr, 422, 418)
}

// SubAccountIdentifier
func (m ReplayChartDataFixedPointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Ptr, 422, 418)
}

// OptionsPriceSendIntervalInSeconds
func (m ReplayChartDataFixedPointer) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32Fixed(m.Ptr, 426, 422)
}

// OptionsPriceSendIntervalInSeconds
func (m ReplayChartDataFixedPointerBuilder) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32Fixed(m.Ptr, 426, 422)
}

// SetRequestID
func (m ReplayChartDataFixedPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetSymbol
func (m ReplayChartDataFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 72, 8, value)
}

// SetTradeAccount
func (m ReplayChartDataFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 104, 72, value)
}

// SetTimeZone
func (m ReplayChartDataFixedPointerBuilder) SetTimeZone(value string) {
	message.SetStringFixed(m.Ptr, 360, 104, value)
}

// SetStartDateTimeForInitialData
func (m ReplayChartDataFixedPointerBuilder) SetStartDateTimeForInitialData(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Ptr, 368, 360, int64(value))
}

// SetStartDateTime
func (m ReplayChartDataFixedPointerBuilder) SetStartDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Ptr, 376, 368, int64(value))
}

// SetStopDateTime
func (m ReplayChartDataFixedPointerBuilder) SetStopDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Ptr, 384, 376, int64(value))
}

// SetSessionBeginTimeInSeconds
func (m ReplayChartDataFixedPointerBuilder) SetSessionBeginTimeInSeconds(value uint16) {
	message.SetUint16Fixed(m.Ptr, 386, 384, value)
}

// SetSessionEndTimeInSeconds
func (m ReplayChartDataFixedPointerBuilder) SetSessionEndTimeInSeconds(value uint16) {
	message.SetUint16Fixed(m.Ptr, 388, 386, value)
}

// SetReplaySpeed
func (m ReplayChartDataFixedPointerBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32Fixed(m.Ptr, 392, 388, value)
}

// SetBarTimeInSeconds
func (m ReplayChartDataFixedPointerBuilder) SetBarTimeInSeconds(value int32) {
	message.SetInt32Fixed(m.Ptr, 396, 392, value)
}

// SetPauseReplayAfterInitialDataSent
func (m ReplayChartDataFixedPointerBuilder) SetPauseReplayAfterInitialDataSent(value uint8) {
	message.SetUint8Fixed(m.Ptr, 397, 396, value)
}

// SetUseSavedPriorState
func (m ReplayChartDataFixedPointerBuilder) SetUseSavedPriorState(value uint8) {
	message.SetUint8Fixed(m.Ptr, 398, 397, value)
}

// SetSymbolVolatility
func (m ReplayChartDataFixedPointerBuilder) SetSymbolVolatility(value float32) {
	message.SetFloat32Fixed(m.Ptr, 402, 398, value)
}

// SetInterestRate
func (m ReplayChartDataFixedPointerBuilder) SetInterestRate(value float32) {
	message.SetFloat32Fixed(m.Ptr, 406, 402, value)
}

// SetNumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataFixedPointerBuilder) SetNumberOfOrdersToTriggerFinishToStopDateTime(value int32) {
	message.SetInt32Fixed(m.Ptr, 410, 406, value)
}

// SetMaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataFixedPointerBuilder) SetMaximumNumberOfOrdersPerReplaySession(value int32) {
	message.SetInt32Fixed(m.Ptr, 414, 410, value)
}

// SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataFixedPointerBuilder) SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(value int32) {
	message.SetInt32Fixed(m.Ptr, 418, 414, value)
}

// SetSubAccountIdentifier
func (m ReplayChartDataFixedPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Ptr, 422, 418, value)
}

// SetOptionsPriceSendIntervalInSeconds
func (m ReplayChartDataFixedPointerBuilder) SetOptionsPriceSendIntervalInSeconds(value int32) {
	message.SetInt32Fixed(m.Ptr, 426, 422, value)
}

// Copy
func (m ReplayChartDataFixedPointer) Copy(to ReplayChartDataFixedBuilder) {
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
func (m ReplayChartDataFixedPointerBuilder) Copy(to ReplayChartDataFixedBuilder) {
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
func (m ReplayChartDataFixedPointer) CopyPointer(to ReplayChartDataFixedPointerBuilder) {
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
func (m ReplayChartDataFixedPointerBuilder) CopyPointer(to ReplayChartDataFixedPointerBuilder) {
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
func (m ReplayChartDataFixedPointer) CopyTo(to ReplayChartDataBuilder) {
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
func (m ReplayChartDataFixedPointerBuilder) CopyTo(to ReplayChartDataBuilder) {
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
func (m ReplayChartDataFixedPointer) CopyToPointer(to ReplayChartDataPointerBuilder) {
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
func (m ReplayChartDataFixedPointerBuilder) CopyToPointer(to ReplayChartDataPointerBuilder) {
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
