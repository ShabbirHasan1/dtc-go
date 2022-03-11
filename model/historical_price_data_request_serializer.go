//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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

func (m HistoricalPriceDataRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRequestPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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

func (m HistoricalPriceDataRequestPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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

func (m HistoricalPriceDataRequest) MarshalJSONTo(b []byte) ([]byte, error) {
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

func (m HistoricalPriceDataRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
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
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRequestPointer) MarshalJSONTo(b []byte) ([]byte, error) {
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

func (m HistoricalPriceDataRequestPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
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

func (m *HistoricalPriceDataRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRequestPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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

func (m *HistoricalPriceDataRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRequestPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRequestFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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

func (m HistoricalPriceDataRequestFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRequestFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
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

func (m HistoricalPriceDataRequestFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
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
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRequestFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRequestFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRequestBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetSymbol(r.ReadString())
		case 3:
			m.SetExchange(r.ReadString())
		case 4:
			m.SetRecordInterval(HistoricalDataIntervalEnum(r.ReadInt32()))
		case 5:
			m.SetStartDateTime(DateTime(r.ReadInt64()))
		case 6:
			m.SetEndDateTime(DateTime(r.ReadInt64()))
		case 7:
			m.SetMaxDaysToReturn(r.ReadUint32())
		case 8:
			m.SetUseZLibCompression(r.ReadUint8())
		case 9:
			m.SetRequestDividendAdjustedStockData(r.ReadUint8())
		case 10:
			m.SetInteger_1(r.ReadUint16())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRequestPointerBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetSymbol(r.ReadString())
		case 3:
			m.SetExchange(r.ReadString())
		case 4:
			m.SetRecordInterval(HistoricalDataIntervalEnum(r.ReadInt32()))
		case 5:
			m.SetStartDateTime(DateTime(r.ReadInt64()))
		case 6:
			m.SetEndDateTime(DateTime(r.ReadInt64()))
		case 7:
			m.SetMaxDaysToReturn(r.ReadUint32())
		case 8:
			m.SetUseZLibCompression(r.ReadUint8())
		case 9:
			m.SetRequestDividendAdjustedStockData(r.ReadUint8())
		case 10:
			m.SetInteger_1(r.ReadUint16())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRequestBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 800)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Symbol())
	w.WriteString(3, m.Exchange())
	w.WriteVarint32(4, int32(m.RecordInterval()))
	w.WriteFixed64Int64(5, int64(m.StartDateTime()))
	w.WriteFixed64Int64(6, int64(m.EndDateTime()))
	w.WriteUvarint32(7, m.MaxDaysToReturn())
	w.WriteUvarint8(8, m.UseZLibCompression())
	w.WriteUvarint8(9, m.RequestDividendAdjustedStockData())
	w.WriteUvarint16(10, m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRequestPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 800)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Symbol())
	w.WriteString(3, m.Exchange())
	w.WriteVarint32(4, int32(m.RecordInterval()))
	w.WriteFixed64Int64(5, int64(m.StartDateTime()))
	w.WriteFixed64Int64(6, int64(m.EndDateTime()))
	w.WriteUvarint32(7, m.MaxDaysToReturn())
	w.WriteUvarint8(8, m.UseZLibCompression())
	w.WriteUvarint8(9, m.RequestDividendAdjustedStockData())
	w.WriteUvarint16(10, m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRequestFixedBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetSymbol(r.ReadString())
		case 3:
			m.SetExchange(r.ReadString())
		case 4:
			m.SetRecordInterval(HistoricalDataIntervalEnum(r.ReadInt32()))
		case 5:
			m.SetStartDateTime(DateTime(r.ReadInt64()))
		case 6:
			m.SetEndDateTime(DateTime(r.ReadInt64()))
		case 7:
			m.SetMaxDaysToReturn(r.ReadUint32())
		case 8:
			m.SetUseZLibCompression(r.ReadUint8())
		case 9:
			m.SetRequestDividendAdjustedStockData(r.ReadUint8())
		case 10:
			m.SetInteger_1(r.ReadUint16())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRequestFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetSymbol(r.ReadString())
		case 3:
			m.SetExchange(r.ReadString())
		case 4:
			m.SetRecordInterval(HistoricalDataIntervalEnum(r.ReadInt32()))
		case 5:
			m.SetStartDateTime(DateTime(r.ReadInt64()))
		case 6:
			m.SetEndDateTime(DateTime(r.ReadInt64()))
		case 7:
			m.SetMaxDaysToReturn(r.ReadUint32())
		case 8:
			m.SetUseZLibCompression(r.ReadUint8())
		case 9:
			m.SetRequestDividendAdjustedStockData(r.ReadUint8())
		case 10:
			m.SetInteger_1(r.ReadUint16())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRequestFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 800)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Symbol())
	w.WriteString(3, m.Exchange())
	w.WriteVarint32(4, int32(m.RecordInterval()))
	w.WriteFixed64Int64(5, int64(m.StartDateTime()))
	w.WriteFixed64Int64(6, int64(m.EndDateTime()))
	w.WriteUvarint32(7, m.MaxDaysToReturn())
	w.WriteUvarint8(8, m.UseZLibCompression())
	w.WriteUvarint8(9, m.RequestDividendAdjustedStockData())
	w.WriteUvarint16(10, m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRequestFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 800)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Symbol())
	w.WriteString(3, m.Exchange())
	w.WriteVarint32(4, int32(m.RecordInterval()))
	w.WriteFixed64Int64(5, int64(m.StartDateTime()))
	w.WriteFixed64Int64(6, int64(m.EndDateTime()))
	w.WriteUvarint32(7, m.MaxDaysToReturn())
	w.WriteUvarint8(8, m.UseZLibCompression())
	w.WriteUvarint8(9, m.RequestDividendAdjustedStockData())
	w.WriteUvarint16(10, m.Integer_1())
	return w.Finish(), nil
}
