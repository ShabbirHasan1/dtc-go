//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRequestFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":800,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RecordInterval()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.EndDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.MaxDaysToReturn())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequestDividendAdjustedStockData())
	w.Data = append(w.Data, ',')
	w.Uint16(m.Integer_1())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRequestFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":800,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RecordInterval()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.EndDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.MaxDaysToReturn())
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.RequestDividendAdjustedStockData())
	w.Data = append(w.Data, ',')
	w.Uint16(m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 800)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("RecordInterval", int32(m.RecordInterval()))
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("EndDateTime", int64(m.EndDateTime()))
	w.Uint32Field("MaxDaysToReturn", m.MaxDaysToReturn())
	w.Uint8Field("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("RequestDividendAdjustedStockData", m.RequestDividendAdjustedStockData())
	w.Uint16Field("Integer_1", m.Integer_1())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRequestFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 800)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("RecordInterval", int32(m.RecordInterval()))
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("EndDateTime", int64(m.EndDateTime()))
	w.Uint32Field("MaxDaysToReturn", m.MaxDaysToReturn())
	w.Uint8Field("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("RequestDividendAdjustedStockData", m.RequestDividendAdjustedStockData())
	w.Uint16Field("Integer_1", m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRequestFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 800 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetRecordInterval(HistoricalDataIntervalEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetEndDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetMaxDaysToReturn(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUseZLibCompression(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestDividendAdjustedStockData(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInteger_1(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRequestFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 800 {
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
			m.SetRequestID(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "RecordInterval":
			m.SetRecordInterval(HistoricalDataIntervalEnum(r.Int32()))
		case "StartDateTime":
			m.SetStartDateTime(DateTime(r.Int64()))
		case "EndDateTime":
			m.SetEndDateTime(DateTime(r.Int64()))
		case "MaxDaysToReturn":
			m.SetMaxDaysToReturn(r.Uint32())
		case "UseZLibCompression":
			m.SetUseZLibCompression(r.Uint8())
		case "RequestDividendAdjustedStockData":
			m.SetRequestDividendAdjustedStockData(r.Uint8())
		case "Integer_1":
			m.SetInteger_1(r.Uint16())
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
