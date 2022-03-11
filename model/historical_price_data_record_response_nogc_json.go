//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRecordResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":803,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Float64(m.BidVolume())
	w.Data = append(w.Data, ',')
	w.Float64(m.AskVolume())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRecordResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":803,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.LastPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Float64(m.BidVolume())
	w.Data = append(w.Data, ',')
	w.Float64(m.AskVolume())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRecordResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 803)
	w.Int32Field("RequestID", m.RequestID())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Float64Field("OpenPrice", m.OpenPrice())
	w.Float64Field("HighPrice", m.HighPrice())
	w.Float64Field("LowPrice", m.LowPrice())
	w.Float64Field("LastPrice", m.LastPrice())
	w.Float64Field("Volume", m.Volume())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Uint32Field("NumTrades", m.NumTrades())
	w.Float64Field("BidVolume", m.BidVolume())
	w.Float64Field("AskVolume", m.AskVolume())
	w.Uint8Field("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRecordResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 803)
	w.Int32Field("RequestID", m.RequestID())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Float64Field("OpenPrice", m.OpenPrice())
	w.Float64Field("HighPrice", m.HighPrice())
	w.Float64Field("LowPrice", m.LowPrice())
	w.Float64Field("LastPrice", m.LastPrice())
	w.Float64Field("Volume", m.Volume())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Uint32Field("NumTrades", m.NumTrades())
	w.Float64Field("BidVolume", m.BidVolume())
	w.Float64Field("AskVolume", m.AskVolume())
	w.Uint8Field("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRecordResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 803 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetHighPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLowPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenInterest(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidVolume(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskVolume(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalRecord(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRecordResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 803 {
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
		case "StartDateTime":
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "OpenPrice":
			m.SetOpenPrice(r.Float64())
		case "HighPrice":
			m.SetHighPrice(r.Float64())
		case "LowPrice":
			m.SetLowPrice(r.Float64())
		case "LastPrice":
			m.SetLastPrice(r.Float64())
		case "Volume":
			m.SetVolume(r.Float64())
		case "OpenInterest":
			m.SetOpenInterest(r.Uint32())
		case "NumTrades":
			m.SetNumTrades(r.Uint32())
		case "BidVolume":
			m.SetBidVolume(r.Float64())
		case "AskVolume":
			m.SetAskVolume(r.Float64())
		case "IsFinalRecord":
			m.SetIsFinalRecord(r.Uint8())
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
