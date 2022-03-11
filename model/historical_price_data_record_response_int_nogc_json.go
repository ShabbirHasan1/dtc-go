//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRecordResponse_IntPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":805,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int32(m.OpenPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.HighPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.LowPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.LastPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidVolume())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskVolume())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":805,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int32(m.OpenPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.HighPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.LowPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.LastPrice())
	w.Data = append(w.Data, ',')
	w.Int32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.OpenInterest())
	w.Data = append(w.Data, ',')
	w.Int32(m.BidVolume())
	w.Data = append(w.Data, ',')
	w.Int32(m.AskVolume())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRecordResponse_IntPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 805)
	w.Int32Field("RequestID", m.RequestID())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int32Field("OpenPrice", m.OpenPrice())
	w.Int32Field("HighPrice", m.HighPrice())
	w.Int32Field("LowPrice", m.LowPrice())
	w.Int32Field("LastPrice", m.LastPrice())
	w.Int32Field("Volume", m.Volume())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Uint32Field("NumTrades", m.NumTrades())
	w.Int32Field("BidVolume", m.BidVolume())
	w.Int32Field("AskVolume", m.AskVolume())
	w.Uint8Field("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRecordResponse_IntPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 805)
	w.Int32Field("RequestID", m.RequestID())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int32Field("OpenPrice", m.OpenPrice())
	w.Int32Field("HighPrice", m.HighPrice())
	w.Int32Field("LowPrice", m.LowPrice())
	w.Int32Field("LastPrice", m.LastPrice())
	w.Int32Field("Volume", m.Volume())
	w.Uint32Field("OpenInterest", m.OpenInterest())
	w.Uint32Field("NumTrades", m.NumTrades())
	w.Int32Field("BidVolume", m.BidVolume())
	w.Int32Field("AskVolume", m.AskVolume())
	w.Uint8Field("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRecordResponse_IntPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 805 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetHighPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLowPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenInterest(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetBidVolume(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAskVolume(r.Int32())
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

func (m *HistoricalPriceDataRecordResponse_IntPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 805 {
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
			m.SetStartDateTime(DateTime(r.Int64()))
		case "OpenPrice":
			m.SetOpenPrice(r.Int32())
		case "HighPrice":
			m.SetHighPrice(r.Int32())
		case "LowPrice":
			m.SetLowPrice(r.Int32())
		case "LastPrice":
			m.SetLastPrice(r.Int32())
		case "Volume":
			m.SetVolume(r.Int32())
		case "OpenInterest":
			m.SetOpenInterest(r.Uint32())
		case "NumTrades":
			m.SetNumTrades(r.Uint32())
		case "BidVolume":
			m.SetBidVolume(r.Int32())
		case "AskVolume":
			m.SetAskVolume(r.Int32())
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
