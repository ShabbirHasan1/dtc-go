package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _ReplayChartDataDefault = []byte{88, 0, 120, 39, 88, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const ReplayChartDataSize = 88

type ReplayChartData struct {
	message.VLS
}

type ReplayChartDataBuilder struct {
	message.VLSBuilder
}

func NewReplayChartDataFrom(b []byte) ReplayChartData {
	return ReplayChartData{VLS: message.NewVLSFrom(b)}
}

func WrapReplayChartData(b []byte) ReplayChartData {
	return ReplayChartData{VLS: message.WrapVLS(b)}
}

func NewReplayChartData() ReplayChartDataBuilder {
	a := ReplayChartDataBuilder{message.NewVLSBuilder(88)}
	a.Unsafe().SetBytes(0, _ReplayChartDataDefault)
	return a
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
func (m ReplayChartDataBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ReplayChartDataDefault)
}

// ToBuilder
func (m ReplayChartData) ToBuilder() ReplayChartDataBuilder {
	return ReplayChartDataBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m ReplayChartDataBuilder) Finish() ReplayChartData {
	return ReplayChartData{m.VLSBuilder.Finish()}
}

// RequestID
func (m ReplayChartData) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m ReplayChartDataBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// Symbol
func (m ReplayChartData) Symbol() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Symbol
func (m ReplayChartDataBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount
func (m ReplayChartData) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount
func (m ReplayChartDataBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TimeZone
func (m ReplayChartData) TimeZone() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// TimeZone
func (m ReplayChartDataBuilder) TimeZone() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// StartDateTimeForInitialData
func (m ReplayChartData) StartDateTimeForInitialData() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 30, 22))
}

// StartDateTimeForInitialData
func (m ReplayChartDataBuilder) StartDateTimeForInitialData() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 30, 22))
}

// StartDateTime
func (m ReplayChartData) StartDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 38, 30))
}

// StartDateTime
func (m ReplayChartDataBuilder) StartDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 38, 30))
}

// StopDateTime
func (m ReplayChartData) StopDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 46, 38))
}

// StopDateTime
func (m ReplayChartDataBuilder) StopDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 46, 38))
}

// SessionBeginTimeInSeconds
func (m ReplayChartData) SessionBeginTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Unsafe(), 48, 46)
}

// SessionBeginTimeInSeconds
func (m ReplayChartDataBuilder) SessionBeginTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Unsafe(), 48, 46)
}

// SessionEndTimeInSeconds
func (m ReplayChartData) SessionEndTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Unsafe(), 50, 48)
}

// SessionEndTimeInSeconds
func (m ReplayChartDataBuilder) SessionEndTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Unsafe(), 50, 48)
}

// ReplaySpeed
func (m ReplayChartData) ReplaySpeed() float32 {
	return message.Float32VLS(m.Unsafe(), 54, 50)
}

// ReplaySpeed
func (m ReplayChartDataBuilder) ReplaySpeed() float32 {
	return message.Float32VLS(m.Unsafe(), 54, 50)
}

// BarTimeInSeconds
func (m ReplayChartData) BarTimeInSeconds() int32 {
	return message.Int32VLS(m.Unsafe(), 58, 54)
}

// BarTimeInSeconds
func (m ReplayChartDataBuilder) BarTimeInSeconds() int32 {
	return message.Int32VLS(m.Unsafe(), 58, 54)
}

// PauseReplayAfterInitialDataSent
func (m ReplayChartData) PauseReplayAfterInitialDataSent() uint8 {
	return message.Uint8VLS(m.Unsafe(), 59, 58)
}

// PauseReplayAfterInitialDataSent
func (m ReplayChartDataBuilder) PauseReplayAfterInitialDataSent() uint8 {
	return message.Uint8VLS(m.Unsafe(), 59, 58)
}

// UseSavedPriorState
func (m ReplayChartData) UseSavedPriorState() uint8 {
	return message.Uint8VLS(m.Unsafe(), 60, 59)
}

// UseSavedPriorState
func (m ReplayChartDataBuilder) UseSavedPriorState() uint8 {
	return message.Uint8VLS(m.Unsafe(), 60, 59)
}

// SymbolVolatility
func (m ReplayChartData) SymbolVolatility() float32 {
	return message.Float32VLS(m.Unsafe(), 64, 60)
}

// SymbolVolatility
func (m ReplayChartDataBuilder) SymbolVolatility() float32 {
	return message.Float32VLS(m.Unsafe(), 64, 60)
}

// InterestRate
func (m ReplayChartData) InterestRate() float32 {
	return message.Float32VLS(m.Unsafe(), 68, 64)
}

// InterestRate
func (m ReplayChartDataBuilder) InterestRate() float32 {
	return message.Float32VLS(m.Unsafe(), 68, 64)
}

// NumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartData) NumberOfOrdersToTriggerFinishToStopDateTime() int32 {
	return message.Int32VLS(m.Unsafe(), 72, 68)
}

// NumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataBuilder) NumberOfOrdersToTriggerFinishToStopDateTime() int32 {
	return message.Int32VLS(m.Unsafe(), 72, 68)
}

// MaximumNumberOfOrdersPerReplaySession
func (m ReplayChartData) MaximumNumberOfOrdersPerReplaySession() int32 {
	return message.Int32VLS(m.Unsafe(), 76, 72)
}

// MaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataBuilder) MaximumNumberOfOrdersPerReplaySession() int32 {
	return message.Int32VLS(m.Unsafe(), 76, 72)
}

// NumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartData) NumberOfDaysForInitialDataFromBeforeLastSavedDateTime() int32 {
	return message.Int32VLS(m.Unsafe(), 80, 76)
}

// NumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataBuilder) NumberOfDaysForInitialDataFromBeforeLastSavedDateTime() int32 {
	return message.Int32VLS(m.Unsafe(), 80, 76)
}

// SubAccountIdentifier
func (m ReplayChartData) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 84, 80)
}

// SubAccountIdentifier
func (m ReplayChartDataBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 84, 80)
}

// OptionsPriceSendIntervalInSeconds
func (m ReplayChartData) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32VLS(m.Unsafe(), 88, 84)
}

// OptionsPriceSendIntervalInSeconds
func (m ReplayChartDataBuilder) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32VLS(m.Unsafe(), 88, 84)
}

// SetRequestID
func (m ReplayChartDataBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetSymbol
func (m *ReplayChartDataBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *ReplayChartDataBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetTimeZone
func (m *ReplayChartDataBuilder) SetTimeZone(value string) {
	message.SetStringVLS(&m.VLSBuilder, 22, 18, value)
}

// SetStartDateTimeForInitialData
func (m ReplayChartDataBuilder) SetStartDateTimeForInitialData(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 30, 22, int64(value))
}

// SetStartDateTime
func (m ReplayChartDataBuilder) SetStartDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 38, 30, int64(value))
}

// SetStopDateTime
func (m ReplayChartDataBuilder) SetStopDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 46, 38, int64(value))
}

// SetSessionBeginTimeInSeconds
func (m ReplayChartDataBuilder) SetSessionBeginTimeInSeconds(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 48, 46, value)
}

// SetSessionEndTimeInSeconds
func (m ReplayChartDataBuilder) SetSessionEndTimeInSeconds(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 50, 48, value)
}

// SetReplaySpeed
func (m ReplayChartDataBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 54, 50, value)
}

// SetBarTimeInSeconds
func (m ReplayChartDataBuilder) SetBarTimeInSeconds(value int32) {
	message.SetInt32VLS(m.Unsafe(), 58, 54, value)
}

// SetPauseReplayAfterInitialDataSent
func (m ReplayChartDataBuilder) SetPauseReplayAfterInitialDataSent(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 59, 58, value)
}

// SetUseSavedPriorState
func (m ReplayChartDataBuilder) SetUseSavedPriorState(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 60, 59, value)
}

// SetSymbolVolatility
func (m ReplayChartDataBuilder) SetSymbolVolatility(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 64, 60, value)
}

// SetInterestRate
func (m ReplayChartDataBuilder) SetInterestRate(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 68, 64, value)
}

// SetNumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataBuilder) SetNumberOfOrdersToTriggerFinishToStopDateTime(value int32) {
	message.SetInt32VLS(m.Unsafe(), 72, 68, value)
}

// SetMaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataBuilder) SetMaximumNumberOfOrdersPerReplaySession(value int32) {
	message.SetInt32VLS(m.Unsafe(), 76, 72, value)
}

// SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataBuilder) SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(value int32) {
	message.SetInt32VLS(m.Unsafe(), 80, 76, value)
}

// SetSubAccountIdentifier
func (m ReplayChartDataBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 84, 80, value)
}

// SetOptionsPriceSendIntervalInSeconds
func (m ReplayChartDataBuilder) SetOptionsPriceSendIntervalInSeconds(value int32) {
	message.SetInt32VLS(m.Unsafe(), 88, 84, value)
}

// Copy
func (m ReplayChartData) Copy(to ReplayChartDataBuilder) {
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
func (m ReplayChartDataBuilder) Copy(to ReplayChartDataBuilder) {
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
func (m ReplayChartData) CopyPointer(to ReplayChartDataPointerBuilder) {
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
func (m ReplayChartDataBuilder) CopyPointer(to ReplayChartDataPointerBuilder) {
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
func (m ReplayChartData) CopyToPointer(to ReplayChartDataFixedPointerBuilder) {
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
func (m ReplayChartDataBuilder) CopyToPointer(to ReplayChartDataFixedPointerBuilder) {
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
func (m ReplayChartData) CopyTo(to ReplayChartDataFixedBuilder) {
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
func (m ReplayChartDataBuilder) CopyTo(to ReplayChartDataFixedBuilder) {
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
