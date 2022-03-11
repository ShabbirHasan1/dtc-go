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

func (m HistoricalPriceDataTickRecordResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":804,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataTickRecordResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":804,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataTickRecordResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":804,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataTickRecordResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":804,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataTickRecordResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 804)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Volume", m.Volume())
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataTickRecordResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 804)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Volume", m.Volume())
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataTickRecordResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 804)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Volume", m.Volume())
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataTickRecordResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 804)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Volume", m.Volume())
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataTickRecordResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 804 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Float64())
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

func (m *HistoricalPriceDataTickRecordResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 804 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Float64())
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

func (m *HistoricalPriceDataTickRecordResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 804 {
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
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
		case "Price":
			m.SetPrice(r.Float64())
		case "Volume":
			m.SetVolume(r.Float64())
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

func (m *HistoricalPriceDataTickRecordResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 804 {
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
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
		case "Price":
			m.SetPrice(r.Float64())
		case "Volume":
			m.SetVolume(r.Float64())
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

func (m *HistoricalPriceDataTickRecordResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 3:
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.ReadUint16()))
		case 4:
			m.SetPrice(r.ReadFloat64())
		case 5:
			m.SetVolume(r.ReadFloat64())
		case 6:
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

func (m *HistoricalPriceDataTickRecordResponsePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 3:
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.ReadUint16()))
		case 4:
			m.SetPrice(r.ReadFloat64())
		case 5:
			m.SetVolume(r.ReadFloat64())
		case 6:
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

func (m HistoricalPriceDataTickRecordResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 804)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, float64(m.DateTime()))
	w.WriteUvarint16(3, uint16(m.AtBidOrAsk()))
	w.WriteFixed64Float64(4, m.Price())
	w.WriteFixed64Float64(5, m.Volume())
	w.WriteBool(6, m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataTickRecordResponsePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 804)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, float64(m.DateTime()))
	w.WriteUvarint16(3, uint16(m.AtBidOrAsk()))
	w.WriteFixed64Float64(4, m.Price())
	w.WriteFixed64Float64(5, m.Volume())
	w.WriteBool(6, m.IsFinalRecord())
	return w.Finish(), nil
}
