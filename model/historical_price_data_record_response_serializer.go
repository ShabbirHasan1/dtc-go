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

func (m HistoricalPriceDataRecordResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRecordResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

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
	w.Bool(m.IsFinalRecord())
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
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRecordResponse) MarshalJSONTo(b []byte) ([]byte, error) {
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
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRecordResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
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
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
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
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
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
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRecordResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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
	m.SetIsFinalRecord(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
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
	m.SetIsFinalRecord(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRecordResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
			m.SetIsFinalRecord(r.Bool())
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
			m.SetIsFinalRecord(r.Bool())
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

func (m *HistoricalPriceDataRecordResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 3:
			m.SetOpenPrice(r.ReadFloat64())
		case 4:
			m.SetHighPrice(r.ReadFloat64())
		case 5:
			m.SetLowPrice(r.ReadFloat64())
		case 6:
			m.SetLastPrice(r.ReadFloat64())
		case 7:
			m.SetVolume(r.ReadFloat64())
		case 9:
			m.SetBidVolume(r.ReadFloat64())
		case 10:
			m.SetAskVolume(r.ReadFloat64())
		case 11:
			m.SetIsFinalRecord(r.ReadBool())
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

func (m *HistoricalPriceDataRecordResponsePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 3:
			m.SetOpenPrice(r.ReadFloat64())
		case 4:
			m.SetHighPrice(r.ReadFloat64())
		case 5:
			m.SetLowPrice(r.ReadFloat64())
		case 6:
			m.SetLastPrice(r.ReadFloat64())
		case 7:
			m.SetVolume(r.ReadFloat64())
		case 9:
			m.SetBidVolume(r.ReadFloat64())
		case 10:
			m.SetAskVolume(r.ReadFloat64())
		case 11:
			m.SetIsFinalRecord(r.ReadBool())
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

func (m HistoricalPriceDataRecordResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 803)
	w.WriteUvarint32(8, m.NumTrades())
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Int64(2, int64(m.StartDateTime()))
	w.WriteFixed64Float64(3, m.OpenPrice())
	w.WriteFixed64Float64(4, m.HighPrice())
	w.WriteFixed64Float64(5, m.LowPrice())
	w.WriteFixed64Float64(6, m.LastPrice())
	w.WriteFixed64Float64(7, m.Volume())
	w.WriteFixed64Float64(9, m.BidVolume())
	w.WriteFixed64Float64(10, m.AskVolume())
	w.WriteBool(11, m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRecordResponsePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 803)
	w.WriteUvarint32(8, m.NumTrades())
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Int64(2, int64(m.StartDateTime()))
	w.WriteFixed64Float64(3, m.OpenPrice())
	w.WriteFixed64Float64(4, m.HighPrice())
	w.WriteFixed64Float64(5, m.LowPrice())
	w.WriteFixed64Float64(6, m.LastPrice())
	w.WriteFixed64Float64(7, m.Volume())
	w.WriteFixed64Float64(9, m.BidVolume())
	w.WriteFixed64Float64(10, m.AskVolume())
	w.WriteBool(11, m.IsFinalRecord())
	return w.Finish(), nil
}
