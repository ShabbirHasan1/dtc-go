package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const ReplayChartDataSize = 88

const ReplayChartDataFixedSize = 426

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

type ReplayChartData struct {
	message.VLS
}

type ReplayChartDataBuilder struct {
	message.VLSBuilder
}

type ReplayChartDataFixed struct {
	message.Fixed
}

type ReplayChartDataFixedBuilder struct {
	message.Fixed
}

type ReplayChartDataPointer struct {
	message.VLSPointer
}

type ReplayChartDataPointerBuilder struct {
	message.VLSPointerBuilder
}

type ReplayChartDataFixedPointer struct {
	message.FixedPointer
}

type ReplayChartDataFixedPointerBuilder struct {
	message.FixedPointer
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

func AllocReplayChartData() ReplayChartDataPointerBuilder {
	a := ReplayChartDataPointerBuilder{message.AllocVLSBuilder(88)}
	a.Ptr.SetBytes(0, _ReplayChartDataDefault)
	return a
}

func AllocReplayChartDataFrom(b []byte) ReplayChartDataPointer {
	return ReplayChartDataPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m ReplayChartData) ToBuilder() ReplayChartDataBuilder {
	return ReplayChartDataBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m ReplayChartDataFixed) ToBuilder() ReplayChartDataFixedBuilder {
	return ReplayChartDataFixedBuilder{m.Fixed}
}

// ToBuilder
func (m ReplayChartDataPointer) ToBuilder() ReplayChartDataPointerBuilder {
	return ReplayChartDataPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m ReplayChartDataFixedPointer) ToBuilder() ReplayChartDataFixedPointerBuilder {
	return ReplayChartDataFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m ReplayChartDataBuilder) Finish() ReplayChartData {
	return ReplayChartData{m.VLSBuilder.Finish()}
}

// Finish
func (m ReplayChartDataFixedBuilder) Finish() ReplayChartDataFixed {
	return ReplayChartDataFixed{m.Fixed}
}

// Finish
func (m *ReplayChartDataPointerBuilder) Finish() ReplayChartDataPointer {
	return ReplayChartDataPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *ReplayChartDataFixedPointerBuilder) Finish() ReplayChartDataFixedPointer {
	return ReplayChartDataFixedPointer{m.FixedPointer}
}

// RequestID
func (m ReplayChartData) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m ReplayChartDataBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
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
func (m ReplayChartData) Symbol() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Symbol
func (m ReplayChartDataBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
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
func (m ReplayChartData) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount
func (m ReplayChartDataBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
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
func (m ReplayChartData) TimeZone() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
}

// TimeZone
func (m ReplayChartDataBuilder) TimeZone() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
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
func (m ReplayChartData) StartDateTimeForInitialData() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 30, 22))
}

// StartDateTimeForInitialData
func (m ReplayChartDataBuilder) StartDateTimeForInitialData() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 30, 22))
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
func (m ReplayChartData) StartDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 38, 30))
}

// StartDateTime
func (m ReplayChartDataBuilder) StartDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 38, 30))
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
func (m ReplayChartData) StopDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 46, 38))
}

// StopDateTime
func (m ReplayChartDataBuilder) StopDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64VLS(m.Unsafe(), 46, 38))
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
func (m ReplayChartData) SessionBeginTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Unsafe(), 48, 46)
}

// SessionBeginTimeInSeconds
func (m ReplayChartDataBuilder) SessionBeginTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Unsafe(), 48, 46)
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
func (m ReplayChartData) SessionEndTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Unsafe(), 50, 48)
}

// SessionEndTimeInSeconds
func (m ReplayChartDataBuilder) SessionEndTimeInSeconds() uint16 {
	return message.Uint16VLS(m.Unsafe(), 50, 48)
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
func (m ReplayChartData) ReplaySpeed() float32 {
	return message.Float32VLS(m.Unsafe(), 54, 50)
}

// ReplaySpeed
func (m ReplayChartDataBuilder) ReplaySpeed() float32 {
	return message.Float32VLS(m.Unsafe(), 54, 50)
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
func (m ReplayChartData) BarTimeInSeconds() int32 {
	return message.Int32VLS(m.Unsafe(), 58, 54)
}

// BarTimeInSeconds
func (m ReplayChartDataBuilder) BarTimeInSeconds() int32 {
	return message.Int32VLS(m.Unsafe(), 58, 54)
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
func (m ReplayChartData) PauseReplayAfterInitialDataSent() uint8 {
	return message.Uint8VLS(m.Unsafe(), 59, 58)
}

// PauseReplayAfterInitialDataSent
func (m ReplayChartDataBuilder) PauseReplayAfterInitialDataSent() uint8 {
	return message.Uint8VLS(m.Unsafe(), 59, 58)
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
func (m ReplayChartData) UseSavedPriorState() uint8 {
	return message.Uint8VLS(m.Unsafe(), 60, 59)
}

// UseSavedPriorState
func (m ReplayChartDataBuilder) UseSavedPriorState() uint8 {
	return message.Uint8VLS(m.Unsafe(), 60, 59)
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
func (m ReplayChartData) SymbolVolatility() float32 {
	return message.Float32VLS(m.Unsafe(), 64, 60)
}

// SymbolVolatility
func (m ReplayChartDataBuilder) SymbolVolatility() float32 {
	return message.Float32VLS(m.Unsafe(), 64, 60)
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
func (m ReplayChartData) InterestRate() float32 {
	return message.Float32VLS(m.Unsafe(), 68, 64)
}

// InterestRate
func (m ReplayChartDataBuilder) InterestRate() float32 {
	return message.Float32VLS(m.Unsafe(), 68, 64)
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
func (m ReplayChartData) NumberOfOrdersToTriggerFinishToStopDateTime() int32 {
	return message.Int32VLS(m.Unsafe(), 72, 68)
}

// NumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataBuilder) NumberOfOrdersToTriggerFinishToStopDateTime() int32 {
	return message.Int32VLS(m.Unsafe(), 72, 68)
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
func (m ReplayChartData) MaximumNumberOfOrdersPerReplaySession() int32 {
	return message.Int32VLS(m.Unsafe(), 76, 72)
}

// MaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataBuilder) MaximumNumberOfOrdersPerReplaySession() int32 {
	return message.Int32VLS(m.Unsafe(), 76, 72)
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
func (m ReplayChartData) NumberOfDaysForInitialDataFromBeforeLastSavedDateTime() int32 {
	return message.Int32VLS(m.Unsafe(), 80, 76)
}

// NumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataBuilder) NumberOfDaysForInitialDataFromBeforeLastSavedDateTime() int32 {
	return message.Int32VLS(m.Unsafe(), 80, 76)
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
func (m ReplayChartData) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 84, 80)
}

// SubAccountIdentifier
func (m ReplayChartDataBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32VLS(m.Unsafe(), 84, 80)
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
func (m ReplayChartData) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32VLS(m.Unsafe(), 88, 84)
}

// OptionsPriceSendIntervalInSeconds
func (m ReplayChartDataBuilder) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32VLS(m.Unsafe(), 88, 84)
}

// OptionsPriceSendIntervalInSeconds
func (m ReplayChartDataPointer) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32VLS(m.Ptr, 88, 84)
}

// OptionsPriceSendIntervalInSeconds
func (m ReplayChartDataPointerBuilder) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32VLS(m.Ptr, 88, 84)
}

// RequestID
func (m ReplayChartDataFixed) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m ReplayChartDataFixedBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
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
func (m ReplayChartDataFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Symbol
func (m ReplayChartDataFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
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
func (m ReplayChartDataFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 104, 72)
}

// TradeAccount
func (m ReplayChartDataFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 104, 72)
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
func (m ReplayChartDataFixed) TimeZone() string {
	return message.StringFixed(m.Unsafe(), 360, 104)
}

// TimeZone
func (m ReplayChartDataFixedBuilder) TimeZone() string {
	return message.StringFixed(m.Unsafe(), 360, 104)
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
func (m ReplayChartDataFixed) StartDateTimeForInitialData() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 368, 360))
}

// StartDateTimeForInitialData
func (m ReplayChartDataFixedBuilder) StartDateTimeForInitialData() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 368, 360))
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
func (m ReplayChartDataFixed) StartDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 376, 368))
}

// StartDateTime
func (m ReplayChartDataFixedBuilder) StartDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 376, 368))
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
func (m ReplayChartDataFixed) StopDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 384, 376))
}

// StopDateTime
func (m ReplayChartDataFixedBuilder) StopDateTime() DateTimeWithMillisecondsInt {
	return DateTimeWithMillisecondsInt(message.Int64Fixed(m.Unsafe(), 384, 376))
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
func (m ReplayChartDataFixed) SessionBeginTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 386, 384)
}

// SessionBeginTimeInSeconds
func (m ReplayChartDataFixedBuilder) SessionBeginTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 386, 384)
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
func (m ReplayChartDataFixed) SessionEndTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 388, 386)
}

// SessionEndTimeInSeconds
func (m ReplayChartDataFixedBuilder) SessionEndTimeInSeconds() uint16 {
	return message.Uint16Fixed(m.Unsafe(), 388, 386)
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
func (m ReplayChartDataFixed) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Unsafe(), 392, 388)
}

// ReplaySpeed
func (m ReplayChartDataFixedBuilder) ReplaySpeed() float32 {
	return message.Float32Fixed(m.Unsafe(), 392, 388)
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
func (m ReplayChartDataFixed) BarTimeInSeconds() int32 {
	return message.Int32Fixed(m.Unsafe(), 396, 392)
}

// BarTimeInSeconds
func (m ReplayChartDataFixedBuilder) BarTimeInSeconds() int32 {
	return message.Int32Fixed(m.Unsafe(), 396, 392)
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
func (m ReplayChartDataFixed) PauseReplayAfterInitialDataSent() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 397, 396)
}

// PauseReplayAfterInitialDataSent
func (m ReplayChartDataFixedBuilder) PauseReplayAfterInitialDataSent() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 397, 396)
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
func (m ReplayChartDataFixed) UseSavedPriorState() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 398, 397)
}

// UseSavedPriorState
func (m ReplayChartDataFixedBuilder) UseSavedPriorState() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 398, 397)
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
func (m ReplayChartDataFixed) SymbolVolatility() float32 {
	return message.Float32Fixed(m.Unsafe(), 402, 398)
}

// SymbolVolatility
func (m ReplayChartDataFixedBuilder) SymbolVolatility() float32 {
	return message.Float32Fixed(m.Unsafe(), 402, 398)
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
func (m ReplayChartDataFixed) InterestRate() float32 {
	return message.Float32Fixed(m.Unsafe(), 406, 402)
}

// InterestRate
func (m ReplayChartDataFixedBuilder) InterestRate() float32 {
	return message.Float32Fixed(m.Unsafe(), 406, 402)
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
func (m ReplayChartDataFixed) NumberOfOrdersToTriggerFinishToStopDateTime() int32 {
	return message.Int32Fixed(m.Unsafe(), 410, 406)
}

// NumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataFixedBuilder) NumberOfOrdersToTriggerFinishToStopDateTime() int32 {
	return message.Int32Fixed(m.Unsafe(), 410, 406)
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
func (m ReplayChartDataFixed) MaximumNumberOfOrdersPerReplaySession() int32 {
	return message.Int32Fixed(m.Unsafe(), 414, 410)
}

// MaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataFixedBuilder) MaximumNumberOfOrdersPerReplaySession() int32 {
	return message.Int32Fixed(m.Unsafe(), 414, 410)
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
func (m ReplayChartDataFixed) NumberOfDaysForInitialDataFromBeforeLastSavedDateTime() int32 {
	return message.Int32Fixed(m.Unsafe(), 418, 414)
}

// NumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataFixedBuilder) NumberOfDaysForInitialDataFromBeforeLastSavedDateTime() int32 {
	return message.Int32Fixed(m.Unsafe(), 418, 414)
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
func (m ReplayChartDataFixed) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 422, 418)
}

// SubAccountIdentifier
func (m ReplayChartDataFixedBuilder) SubAccountIdentifier() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 422, 418)
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
func (m ReplayChartDataFixed) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32Fixed(m.Unsafe(), 426, 422)
}

// OptionsPriceSendIntervalInSeconds
func (m ReplayChartDataFixedBuilder) OptionsPriceSendIntervalInSeconds() int32 {
	return message.Int32Fixed(m.Unsafe(), 426, 422)
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
func (m ReplayChartDataBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m ReplayChartDataPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetSymbol
func (m *ReplayChartDataBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetSymbol
func (m *ReplayChartDataPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetTradeAccount
func (m *ReplayChartDataBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetTradeAccount
func (m *ReplayChartDataPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetTimeZone
func (m *ReplayChartDataBuilder) SetTimeZone(value string) {
	message.SetStringVLS(&m.VLSBuilder, 22, 18, value)
}

// SetTimeZone
func (m *ReplayChartDataPointerBuilder) SetTimeZone(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 22, 18, value)
}

// SetStartDateTimeForInitialData
func (m ReplayChartDataBuilder) SetStartDateTimeForInitialData(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 30, 22, int64(value))
}

// SetStartDateTimeForInitialData
func (m ReplayChartDataPointerBuilder) SetStartDateTimeForInitialData(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Ptr, 30, 22, int64(value))
}

// SetStartDateTime
func (m ReplayChartDataBuilder) SetStartDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 38, 30, int64(value))
}

// SetStartDateTime
func (m ReplayChartDataPointerBuilder) SetStartDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Ptr, 38, 30, int64(value))
}

// SetStopDateTime
func (m ReplayChartDataBuilder) SetStopDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Unsafe(), 46, 38, int64(value))
}

// SetStopDateTime
func (m ReplayChartDataPointerBuilder) SetStopDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64VLS(m.Ptr, 46, 38, int64(value))
}

// SetSessionBeginTimeInSeconds
func (m ReplayChartDataBuilder) SetSessionBeginTimeInSeconds(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 48, 46, value)
}

// SetSessionBeginTimeInSeconds
func (m ReplayChartDataPointerBuilder) SetSessionBeginTimeInSeconds(value uint16) {
	message.SetUint16VLS(m.Ptr, 48, 46, value)
}

// SetSessionEndTimeInSeconds
func (m ReplayChartDataBuilder) SetSessionEndTimeInSeconds(value uint16) {
	message.SetUint16VLS(m.Unsafe(), 50, 48, value)
}

// SetSessionEndTimeInSeconds
func (m ReplayChartDataPointerBuilder) SetSessionEndTimeInSeconds(value uint16) {
	message.SetUint16VLS(m.Ptr, 50, 48, value)
}

// SetReplaySpeed
func (m ReplayChartDataBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 54, 50, value)
}

// SetReplaySpeed
func (m ReplayChartDataPointerBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32VLS(m.Ptr, 54, 50, value)
}

// SetBarTimeInSeconds
func (m ReplayChartDataBuilder) SetBarTimeInSeconds(value int32) {
	message.SetInt32VLS(m.Unsafe(), 58, 54, value)
}

// SetBarTimeInSeconds
func (m ReplayChartDataPointerBuilder) SetBarTimeInSeconds(value int32) {
	message.SetInt32VLS(m.Ptr, 58, 54, value)
}

// SetPauseReplayAfterInitialDataSent
func (m ReplayChartDataBuilder) SetPauseReplayAfterInitialDataSent(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 59, 58, value)
}

// SetPauseReplayAfterInitialDataSent
func (m ReplayChartDataPointerBuilder) SetPauseReplayAfterInitialDataSent(value uint8) {
	message.SetUint8VLS(m.Ptr, 59, 58, value)
}

// SetUseSavedPriorState
func (m ReplayChartDataBuilder) SetUseSavedPriorState(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 60, 59, value)
}

// SetUseSavedPriorState
func (m ReplayChartDataPointerBuilder) SetUseSavedPriorState(value uint8) {
	message.SetUint8VLS(m.Ptr, 60, 59, value)
}

// SetSymbolVolatility
func (m ReplayChartDataBuilder) SetSymbolVolatility(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 64, 60, value)
}

// SetSymbolVolatility
func (m ReplayChartDataPointerBuilder) SetSymbolVolatility(value float32) {
	message.SetFloat32VLS(m.Ptr, 64, 60, value)
}

// SetInterestRate
func (m ReplayChartDataBuilder) SetInterestRate(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 68, 64, value)
}

// SetInterestRate
func (m ReplayChartDataPointerBuilder) SetInterestRate(value float32) {
	message.SetFloat32VLS(m.Ptr, 68, 64, value)
}

// SetNumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataBuilder) SetNumberOfOrdersToTriggerFinishToStopDateTime(value int32) {
	message.SetInt32VLS(m.Unsafe(), 72, 68, value)
}

// SetNumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataPointerBuilder) SetNumberOfOrdersToTriggerFinishToStopDateTime(value int32) {
	message.SetInt32VLS(m.Ptr, 72, 68, value)
}

// SetMaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataBuilder) SetMaximumNumberOfOrdersPerReplaySession(value int32) {
	message.SetInt32VLS(m.Unsafe(), 76, 72, value)
}

// SetMaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataPointerBuilder) SetMaximumNumberOfOrdersPerReplaySession(value int32) {
	message.SetInt32VLS(m.Ptr, 76, 72, value)
}

// SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataBuilder) SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(value int32) {
	message.SetInt32VLS(m.Unsafe(), 80, 76, value)
}

// SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataPointerBuilder) SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(value int32) {
	message.SetInt32VLS(m.Ptr, 80, 76, value)
}

// SetSubAccountIdentifier
func (m ReplayChartDataBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 84, 80, value)
}

// SetSubAccountIdentifier
func (m ReplayChartDataPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32VLS(m.Ptr, 84, 80, value)
}

// SetOptionsPriceSendIntervalInSeconds
func (m ReplayChartDataBuilder) SetOptionsPriceSendIntervalInSeconds(value int32) {
	message.SetInt32VLS(m.Unsafe(), 88, 84, value)
}

// SetOptionsPriceSendIntervalInSeconds
func (m ReplayChartDataPointerBuilder) SetOptionsPriceSendIntervalInSeconds(value int32) {
	message.SetInt32VLS(m.Ptr, 88, 84, value)
}

// SetRequestID
func (m ReplayChartDataFixedBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m ReplayChartDataFixedPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetSymbol
func (m ReplayChartDataFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
}

// SetSymbol
func (m ReplayChartDataFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 72, 8, value)
}

// SetTradeAccount
func (m ReplayChartDataFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 72, value)
}

// SetTradeAccount
func (m ReplayChartDataFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 104, 72, value)
}

// SetTimeZone
func (m ReplayChartDataFixedBuilder) SetTimeZone(value string) {
	message.SetStringFixed(m.Unsafe(), 360, 104, value)
}

// SetTimeZone
func (m ReplayChartDataFixedPointerBuilder) SetTimeZone(value string) {
	message.SetStringFixed(m.Ptr, 360, 104, value)
}

// SetStartDateTimeForInitialData
func (m ReplayChartDataFixedBuilder) SetStartDateTimeForInitialData(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 368, 360, int64(value))
}

// SetStartDateTimeForInitialData
func (m ReplayChartDataFixedPointerBuilder) SetStartDateTimeForInitialData(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Ptr, 368, 360, int64(value))
}

// SetStartDateTime
func (m ReplayChartDataFixedBuilder) SetStartDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 376, 368, int64(value))
}

// SetStartDateTime
func (m ReplayChartDataFixedPointerBuilder) SetStartDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Ptr, 376, 368, int64(value))
}

// SetStopDateTime
func (m ReplayChartDataFixedBuilder) SetStopDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 384, 376, int64(value))
}

// SetStopDateTime
func (m ReplayChartDataFixedPointerBuilder) SetStopDateTime(value DateTimeWithMillisecondsInt) {
	message.SetInt64Fixed(m.Ptr, 384, 376, int64(value))
}

// SetSessionBeginTimeInSeconds
func (m ReplayChartDataFixedBuilder) SetSessionBeginTimeInSeconds(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 386, 384, value)
}

// SetSessionBeginTimeInSeconds
func (m ReplayChartDataFixedPointerBuilder) SetSessionBeginTimeInSeconds(value uint16) {
	message.SetUint16Fixed(m.Ptr, 386, 384, value)
}

// SetSessionEndTimeInSeconds
func (m ReplayChartDataFixedBuilder) SetSessionEndTimeInSeconds(value uint16) {
	message.SetUint16Fixed(m.Unsafe(), 388, 386, value)
}

// SetSessionEndTimeInSeconds
func (m ReplayChartDataFixedPointerBuilder) SetSessionEndTimeInSeconds(value uint16) {
	message.SetUint16Fixed(m.Ptr, 388, 386, value)
}

// SetReplaySpeed
func (m ReplayChartDataFixedBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 392, 388, value)
}

// SetReplaySpeed
func (m ReplayChartDataFixedPointerBuilder) SetReplaySpeed(value float32) {
	message.SetFloat32Fixed(m.Ptr, 392, 388, value)
}

// SetBarTimeInSeconds
func (m ReplayChartDataFixedBuilder) SetBarTimeInSeconds(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 396, 392, value)
}

// SetBarTimeInSeconds
func (m ReplayChartDataFixedPointerBuilder) SetBarTimeInSeconds(value int32) {
	message.SetInt32Fixed(m.Ptr, 396, 392, value)
}

// SetPauseReplayAfterInitialDataSent
func (m ReplayChartDataFixedBuilder) SetPauseReplayAfterInitialDataSent(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 397, 396, value)
}

// SetPauseReplayAfterInitialDataSent
func (m ReplayChartDataFixedPointerBuilder) SetPauseReplayAfterInitialDataSent(value uint8) {
	message.SetUint8Fixed(m.Ptr, 397, 396, value)
}

// SetUseSavedPriorState
func (m ReplayChartDataFixedBuilder) SetUseSavedPriorState(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 398, 397, value)
}

// SetUseSavedPriorState
func (m ReplayChartDataFixedPointerBuilder) SetUseSavedPriorState(value uint8) {
	message.SetUint8Fixed(m.Ptr, 398, 397, value)
}

// SetSymbolVolatility
func (m ReplayChartDataFixedBuilder) SetSymbolVolatility(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 402, 398, value)
}

// SetSymbolVolatility
func (m ReplayChartDataFixedPointerBuilder) SetSymbolVolatility(value float32) {
	message.SetFloat32Fixed(m.Ptr, 402, 398, value)
}

// SetInterestRate
func (m ReplayChartDataFixedBuilder) SetInterestRate(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 406, 402, value)
}

// SetInterestRate
func (m ReplayChartDataFixedPointerBuilder) SetInterestRate(value float32) {
	message.SetFloat32Fixed(m.Ptr, 406, 402, value)
}

// SetNumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataFixedBuilder) SetNumberOfOrdersToTriggerFinishToStopDateTime(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 410, 406, value)
}

// SetNumberOfOrdersToTriggerFinishToStopDateTime
func (m ReplayChartDataFixedPointerBuilder) SetNumberOfOrdersToTriggerFinishToStopDateTime(value int32) {
	message.SetInt32Fixed(m.Ptr, 410, 406, value)
}

// SetMaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataFixedBuilder) SetMaximumNumberOfOrdersPerReplaySession(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 414, 410, value)
}

// SetMaximumNumberOfOrdersPerReplaySession
func (m ReplayChartDataFixedPointerBuilder) SetMaximumNumberOfOrdersPerReplaySession(value int32) {
	message.SetInt32Fixed(m.Ptr, 414, 410, value)
}

// SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataFixedBuilder) SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 418, 414, value)
}

// SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime
func (m ReplayChartDataFixedPointerBuilder) SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(value int32) {
	message.SetInt32Fixed(m.Ptr, 418, 414, value)
}

// SetSubAccountIdentifier
func (m ReplayChartDataFixedBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 422, 418, value)
}

// SetSubAccountIdentifier
func (m ReplayChartDataFixedPointerBuilder) SetSubAccountIdentifier(value uint32) {
	message.SetUint32Fixed(m.Ptr, 422, 418, value)
}

// SetOptionsPriceSendIntervalInSeconds
func (m ReplayChartDataFixedBuilder) SetOptionsPriceSendIntervalInSeconds(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 426, 422, value)
}

// SetOptionsPriceSendIntervalInSeconds
func (m ReplayChartDataFixedPointerBuilder) SetOptionsPriceSendIntervalInSeconds(value int32) {
	message.SetInt32Fixed(m.Ptr, 426, 422, value)
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