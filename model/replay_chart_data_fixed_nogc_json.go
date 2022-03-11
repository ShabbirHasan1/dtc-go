//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ReplayChartDataFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10104,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.TimeZone())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTimeForInitialData()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StopDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.SessionBeginTimeInSeconds())
	w.Data = append(w.Data, ',')
	w.Uint16(m.SessionEndTimeInSeconds())
	w.Data = append(w.Data, ',')
	w.Float32(m.ReplaySpeed())
	w.Data = append(w.Data, ',')
	w.Int32(m.BarTimeInSeconds())
	w.Data = append(w.Data, ',')
	w.Uint8(m.PauseReplayAfterInitialDataSent())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseSavedPriorState())
	w.Data = append(w.Data, ',')
	w.Float32(m.SymbolVolatility())
	w.Data = append(w.Data, ',')
	w.Float32(m.InterestRate())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfOrdersToTriggerFinishToStopDateTime())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumNumberOfOrdersPerReplaySession())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfDaysForInitialDataFromBeforeLastSavedDateTime())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Int32(m.OptionsPriceSendIntervalInSeconds())
	return w.Finish(), nil
}

func (m ReplayChartDataFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10104,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.TimeZone())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTimeForInitialData()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StopDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.SessionBeginTimeInSeconds())
	w.Data = append(w.Data, ',')
	w.Uint16(m.SessionEndTimeInSeconds())
	w.Data = append(w.Data, ',')
	w.Float32(m.ReplaySpeed())
	w.Data = append(w.Data, ',')
	w.Int32(m.BarTimeInSeconds())
	w.Data = append(w.Data, ',')
	w.Uint8(m.PauseReplayAfterInitialDataSent())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseSavedPriorState())
	w.Data = append(w.Data, ',')
	w.Float32(m.SymbolVolatility())
	w.Data = append(w.Data, ',')
	w.Float32(m.InterestRate())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfOrdersToTriggerFinishToStopDateTime())
	w.Data = append(w.Data, ',')
	w.Int32(m.MaximumNumberOfOrdersPerReplaySession())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfDaysForInitialDataFromBeforeLastSavedDateTime())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	w.Data = append(w.Data, ',')
	w.Int32(m.OptionsPriceSendIntervalInSeconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ReplayChartDataFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10104)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("TimeZone", m.TimeZone())
	w.Int64Field("StartDateTimeForInitialData", int64(m.StartDateTimeForInitialData()))
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("StopDateTime", int64(m.StopDateTime()))
	w.Uint16Field("SessionBeginTimeInSeconds", m.SessionBeginTimeInSeconds())
	w.Uint16Field("SessionEndTimeInSeconds", m.SessionEndTimeInSeconds())
	w.Float32Field("ReplaySpeed", m.ReplaySpeed())
	w.Int32Field("BarTimeInSeconds", m.BarTimeInSeconds())
	w.Uint8Field("PauseReplayAfterInitialDataSent", m.PauseReplayAfterInitialDataSent())
	w.Uint8Field("UseSavedPriorState", m.UseSavedPriorState())
	w.Float32Field("SymbolVolatility", m.SymbolVolatility())
	w.Float32Field("InterestRate", m.InterestRate())
	w.Int32Field("NumberOfOrdersToTriggerFinishToStopDateTime", m.NumberOfOrdersToTriggerFinishToStopDateTime())
	w.Int32Field("MaximumNumberOfOrdersPerReplaySession", m.MaximumNumberOfOrdersPerReplaySession())
	w.Int32Field("NumberOfDaysForInitialDataFromBeforeLastSavedDateTime", m.NumberOfDaysForInitialDataFromBeforeLastSavedDateTime())
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	w.Int32Field("OptionsPriceSendIntervalInSeconds", m.OptionsPriceSendIntervalInSeconds())
	return w.Finish(), nil
}

func (m ReplayChartDataFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10104)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("TimeZone", m.TimeZone())
	w.Int64Field("StartDateTimeForInitialData", int64(m.StartDateTimeForInitialData()))
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("StopDateTime", int64(m.StopDateTime()))
	w.Uint16Field("SessionBeginTimeInSeconds", m.SessionBeginTimeInSeconds())
	w.Uint16Field("SessionEndTimeInSeconds", m.SessionEndTimeInSeconds())
	w.Float32Field("ReplaySpeed", m.ReplaySpeed())
	w.Int32Field("BarTimeInSeconds", m.BarTimeInSeconds())
	w.Uint8Field("PauseReplayAfterInitialDataSent", m.PauseReplayAfterInitialDataSent())
	w.Uint8Field("UseSavedPriorState", m.UseSavedPriorState())
	w.Float32Field("SymbolVolatility", m.SymbolVolatility())
	w.Float32Field("InterestRate", m.InterestRate())
	w.Int32Field("NumberOfOrdersToTriggerFinishToStopDateTime", m.NumberOfOrdersToTriggerFinishToStopDateTime())
	w.Int32Field("MaximumNumberOfOrdersPerReplaySession", m.MaximumNumberOfOrdersPerReplaySession())
	w.Int32Field("NumberOfDaysForInitialDataFromBeforeLastSavedDateTime", m.NumberOfDaysForInitialDataFromBeforeLastSavedDateTime())
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	w.Int32Field("OptionsPriceSendIntervalInSeconds", m.OptionsPriceSendIntervalInSeconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ReplayChartDataFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10104 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTimeZone(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTimeForInitialData(DateTimeWithMillisecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTimeWithMillisecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetStopDateTime(DateTimeWithMillisecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionBeginTimeInSeconds(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetSessionEndTimeInSeconds(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetReplaySpeed(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBarTimeInSeconds(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPauseReplayAfterInitialDataSent(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseSavedPriorState(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolVolatility(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetInterestRate(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumberOfOrdersToTriggerFinishToStopDateTime(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumNumberOfOrdersPerReplaySession(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSubAccountIdentifier(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOptionsPriceSendIntervalInSeconds(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ReplayChartDataFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10104 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "RequestID":
			m.SetRequestID(r.Uint32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "TimeZone":
			m.SetTimeZone(r.String())
		case "StartDateTimeForInitialData":
			m.SetStartDateTimeForInitialData(DateTimeWithMillisecondsInt(r.Int64()))
		case "StartDateTime":
			m.SetStartDateTime(DateTimeWithMillisecondsInt(r.Int64()))
		case "StopDateTime":
			m.SetStopDateTime(DateTimeWithMillisecondsInt(r.Int64()))
		case "SessionBeginTimeInSeconds":
			m.SetSessionBeginTimeInSeconds(r.Uint16())
		case "SessionEndTimeInSeconds":
			m.SetSessionEndTimeInSeconds(r.Uint16())
		case "ReplaySpeed":
			m.SetReplaySpeed(r.Float32())
		case "BarTimeInSeconds":
			m.SetBarTimeInSeconds(r.Int32())
		case "PauseReplayAfterInitialDataSent":
			m.SetPauseReplayAfterInitialDataSent(r.Uint8())
		case "UseSavedPriorState":
			m.SetUseSavedPriorState(r.Uint8())
		case "SymbolVolatility":
			m.SetSymbolVolatility(r.Float32())
		case "InterestRate":
			m.SetInterestRate(r.Float32())
		case "NumberOfOrdersToTriggerFinishToStopDateTime":
			m.SetNumberOfOrdersToTriggerFinishToStopDateTime(r.Int32())
		case "MaximumNumberOfOrdersPerReplaySession":
			m.SetMaximumNumberOfOrdersPerReplaySession(r.Int32())
		case "NumberOfDaysForInitialDataFromBeforeLastSavedDateTime":
			m.SetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime(r.Int32())
		case "SubAccountIdentifier":
			m.SetSubAccountIdentifier(r.Uint32())
		case "OptionsPriceSendIntervalInSeconds":
			m.SetOptionsPriceSendIntervalInSeconds(r.Int32())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}
