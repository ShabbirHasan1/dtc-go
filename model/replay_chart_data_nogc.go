package model

import (
	"github.com/moontrade/dtc-go/message"
)

type ReplayChartDataPointer struct {
	message.VLSPointer
}

type ReplayChartDataPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocReplayChartData() ReplayChartDataPointerBuilder {
	a := ReplayChartDataPointerBuilder{message.AllocVLSBuilder(88)}
	a.Ptr.SetBytes(0, _ReplayChartDataDefault)
	return a
}

func AllocReplayChartDataFrom(b []byte) ReplayChartDataPointer {
	return ReplayChartDataPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                                                  = ReplayChartDataSize  (88)
//     Type                                                  = REPLAY_CHART_DATA  (10104)
//     BaseSize                                              = ReplayChartDataSize  (88)
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
func (m ReplayChartDataPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ReplayChartDataDefault)
}

// ToBuilder
func (m ReplayChartDataPointer) ToBuilder() ReplayChartDataPointerBuilder {
	return ReplayChartDataPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *ReplayChartDataPointerBuilder) Finish() ReplayChartDataPointer {
	return ReplayChartDataPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m ReplayChartDataPointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m ReplayChartDataPointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// Symbol
func (m ReplayChartDataPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Symbol
func (m ReplayChartDataPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount
func (m ReplayChartDataPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// TradeAccount
func (m ReplayChartDataPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// TimeZone
func (m ReplayChartDataPointer) TimeZone() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// TimeZone
func (m ReplayChartDataPointerBuilder) TimeZone() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// StartDateTimeForInitialData
func (m ReplayChartDataPointer) StartDateTimeForInitialData() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Ptr, 30, 22))
}

// StartDateTimeForInitialData
func (m ReplayChartDataPointerBuilder) StartDateTimeForInitialData() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Ptr, 30, 22))
}

// StartDateTime
func (m ReplayChartDataPointer) StartDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Ptr, 38, 30))
}

// StartDateTime
func (m ReplayChartDataPointerBuilder) StartDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Ptr, 38, 30))
}

// StopDateTime
func (m ReplayChartDataPointer) StopDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Ptr, 46, 38))
}

// StopDateTime
func (m ReplayChartDataPointerBuilder) StopDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Ptr, 46, 38))
}

// SessionBeginTimeInSeconds
func (m ReplayChartDataPointer) SessionBeginTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Ptr, 48, 46)
}

// SessionBeginTimeInSeconds
func (m ReplayChartDataPointerBuilder) SessionBeginTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Ptr, 48, 46)
}

// SessionEndTimeInSeconds
func (m ReplayChartDataPointer) SessionEndTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Ptr, 50, 48)
}

// SessionEndTimeInSeconds
func (m ReplayChartDataPointerBuilder) SessionEndTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Ptr, 50, 48)
}

// ReplaySpeed
func (m ReplayChartDataPointer) ReplaySpeed() float32 {
	return message.Float32VLS(m.Ptr, 54, 50)
}

// ReplaySpeed
func (m ReplayChartDataPointerBuilder) ReplaySpeed() float32 {
	return message.Float32VLS(m.Ptr, 54, 50)
}

// BarTimeInSeconds
func (m ReplayChartDataPointer) BarTimeInSeconds() int32 {
	return message.Int32VLS(m.Ptr, 58, 54)
}

// BarTimeInSeconds
func (m ReplayChartDataPointerBuilder) BarTimeInSeconds() int32 {
	return message.Int32VLS(m.Ptr, 58, 54)
}

// PauseReplayAfterInitialDataSent
func (m ReplayChartDataPointer) PauseReplayAfterInitialDataSent() uint8 {
	return message.Uint8VLS(m.Ptr, 59, 58)
}

// PauseReplayAfterInitialDataSent
func (m ReplayChartDataPointerBuilder) PauseReplayAfterInitialDataSent() uint8 {
	return message.Uint8VLS(m.Ptr, 59, 58)
}

// UseSavedPriorState
func (m ReplayChartDataPointer) UseSavedPriorState() uint8 {
	return message.Uint8VLS(m.Ptr, 60, 59)
}

// UseSavedPriorState
func (m ReplayChartDataPointerBuilder) UseSavedPriorState() uint8 {
	return message.Uint8VLS(m.Ptr, 60, 59)
}

// SymbolVolatility
func (m ReplayChartDataPointer) SymbolVolatility() float32 {
	return message.Float32VLS(m.Ptr, 64, 60)
}

// SymbolVolatility
func (m ReplayChartDataPointerBuilder) SymbolVolatility() float32 {
	return message.Float32VLS(m.Ptr, 64, 60)
}

// InterestRate
func (m ReplayChartDataPointer) InterestRate() float32 {
	return message.Float32VLS(m.Ptr, 68, 64)
}

// InterestRate
func (m ReplayChartDataPointerBuilder) InterestRate() float32 {
	return message.Float32VLS(m.Ptr, 68, 64)
}

// NumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataPointer) NumberOfOrdersToTriggerFinishToStopDateTime() int32 {
	return message.Int32VLS(m.Ptr, 72, 68)
}

// NumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataPointerBuilder) NumberOfOrdersToTriggerFinishToStopDateTime() int32 {
	return message.Int32VLS(m.Ptr, 72, 68)
}

// MaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataPointer) MaximumNumberOfOrdersPerReplaySession() int32 {
	return message.Int32VLS(m.Ptr, 76, 72)
}

// MaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataPointerBuilder) MaximumNumberOfOrdersPerReplaySession() int32 {
	return message.Int32VLS(m.Ptr, 76, 72)
}

// NumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataPointer) NumberOfDaysForInitialDataFromBeforeLastSavedDateTime() int32 {
	return message.Int32VLS(m.Ptr, 80, 76)
}

// NumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataPointerBuilder) NumberOfDaysForInitialDataFromBeforeLastSavedDateTime() int32 {
	return message.Int32VLS(m.Ptr, 80, 76)
}

// SubAccountIdentifier
func (m ReplayChartDataPointer) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 84, 80)
}

// SubAccountIdentifier
func (m ReplayChartDataPointerBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Ptr, 84, 80)
}

// OptionsPriceSendIntervalInSeconds
func (m ReplayChartDataPointer) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32VLS(m.Ptr, 88, 84)
}

// OptionsPriceSendIntervalInSeconds
func (m ReplayChartDataPointerBuilder) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32VLS(m.Ptr, 88, 84)
}

// SetRequestID
func (m ReplayChartDataPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetSymbol
func (m *ReplayChartDataPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *ReplayChartDataPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetTimeZone
func (m *ReplayChartDataPointerBuilder) SetTimeZone(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 22, 18, value)
}

// SetStartDateTimeForInitialData
func (m ReplayChartDataPointerBuilder) SetStartDateTimeForInitialData(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Ptr, 30, 22, int64(value))
}

// SetStartDateTime
func (m ReplayChartDataPointerBuilder) SetStartDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Ptr, 38, 30, int64(value))
}

// SetStopDateTime
func (m ReplayChartDataPointerBuilder) SetStopDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Ptr, 46, 38, int64(value))
}

// SetSessionBeginTimeInSeconds
func (m ReplayChartDataPointerBuilder) SetSessionBeginTimeInSeconds(value uint16) {
	message.SetUint16VLS(m.Ptr, 48, 46, value)
}

// SetSessionEndTimeInSeconds
func (m ReplayChartDataPointerBuilder) SetSessionEndTimeInSeconds(value uint16) {
	message.SetUint16VLS(m.Ptr, 50, 48, value)
}

// SetReplaySpeed
func (m ReplayChartDataPointerBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32VLS(m.Ptr, 54, 50, value)
}

// SetBarTimeInSeconds
func (m ReplayChartDataPointerBuilder) SetBarTimeInSeconds(value int32) {
	message.SetInt32VLS(m.Ptr, 58, 54, value)
}

// SetPauseReplayAfterInitialDataSent
func (m ReplayChartDataPointerBuilder) SetPauseReplayAfterInitialDataSent(value uint8) {
	message.SetUint8VLS(m.Ptr, 59, 58, value)
}

// SetUseSavedPriorState
func (m ReplayChartDataPointerBuilder) SetUseSavedPriorState(value uint8) {
	message.SetUint8VLS(m.Ptr, 60, 59, value)
}

// SetSymbolVolatility
func (m ReplayChartDataPointerBuilder) SetSymbolVolatility(value float32) {
	message.SetFloat32VLS(m.Ptr, 64, 60, value)
}

// SetInterestRate
func (m ReplayChartDataPointerBuilder) SetInterestRate(value float32) {
	message.SetFloat32VLS(m.Ptr, 68, 64, value)
}

// SetNumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataPointerBuilder) SetNumberOfOrdersToTriggerFinishToStopDateTime(value int32) {
	message.SetInt32VLS(m.Ptr, 72, 68, value)
}

// SetMaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataPointerBuilder) SetMaximumNumberOfOrdersPerReplaySession(value int32) {
	message.SetInt32VLS(m.Ptr, 76, 72, value)
}

// SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataPointerBuilder) SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(value int32) {
	message.SetInt32VLS(m.Ptr, 80, 76, value)
}

// SetSubAccountIdentifier
func (m ReplayChartDataPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Ptr, 84, 80, value)
}

// SetOptionsPriceSendIntervalInSeconds
func (m ReplayChartDataPointerBuilder) SetOptionsPriceSendIntervalInSeconds(value int32) {
	message.SetInt32VLS(m.Ptr, 88, 84, value)
}

// Copy
func (m ReplayChartDataPointer) Copy(to ReplayChartDataBuilder) {
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
func (m ReplayChartDataPointerBuilder) Copy(to ReplayChartDataBuilder) {
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
func (m ReplayChartDataPointer) CopyPointer(to ReplayChartDataPointerBuilder) {
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
func (m ReplayChartDataPointerBuilder) CopyPointer(to ReplayChartDataPointerBuilder) {
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
func (m ReplayChartDataPointer) CopyTo(to ReplayChartDataFixedBuilder) {
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
func (m ReplayChartDataPointerBuilder) CopyTo(to ReplayChartDataFixedBuilder) {
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
func (m ReplayChartDataPointer) CopyToPointer(to ReplayChartDataFixedPointerBuilder) {
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
func (m ReplayChartDataPointerBuilder) CopyToPointer(to ReplayChartDataFixedPointerBuilder) {
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
