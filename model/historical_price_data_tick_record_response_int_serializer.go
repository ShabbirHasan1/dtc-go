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

func (m HistoricalPriceDataTickRecordResponse_Int) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":806,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Int32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataTickRecordResponse_IntBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":806,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Int32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataTickRecordResponse_IntPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":806,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Int32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":806,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Int32(m.Price())
	w.Data = append(w.Data, ',')
	w.Int32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint16(uint16(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataTickRecordResponse_Int) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 806)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Int32Field("Price", m.Price())
	w.Int32Field("Volume", m.Volume())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataTickRecordResponse_IntBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 806)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Int32Field("Price", m.Price())
	w.Int32Field("Volume", m.Volume())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataTickRecordResponse_IntPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 806)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Int32Field("Price", m.Price())
	w.Int32Field("Volume", m.Volume())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 806)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Int32Field("Price", m.Price())
	w.Int32Field("Volume", m.Volume())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataTickRecordResponse_IntBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 806 {
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
	m.SetPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
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

func (m *HistoricalPriceDataTickRecordResponse_IntPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 806 {
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
	m.SetPrice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
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

func (m *HistoricalPriceDataTickRecordResponse_IntBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 806 {
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
		case "Price":
			m.SetPrice(r.Int32())
		case "Volume":
			m.SetVolume(r.Int32())
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
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

func (m *HistoricalPriceDataTickRecordResponse_IntPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 806 {
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
		case "Price":
			m.SetPrice(r.Int32())
		case "Volume":
			m.SetVolume(r.Int32())
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
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

func (m *HistoricalPriceDataTickRecordResponse_IntBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadInt32())
		case 4:
			m.SetVolume(r.ReadInt32())
		case 5:
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.ReadUint16()))
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

func (m *HistoricalPriceDataTickRecordResponse_IntPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadInt32())
		case 4:
			m.SetVolume(r.ReadInt32())
		case 5:
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.ReadUint16()))
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

func (m HistoricalPriceDataTickRecordResponse_IntBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 806)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, float64(m.DateTime()))
	w.WriteVarint32(3, m.Price())
	w.WriteVarint32(4, m.Volume())
	w.WriteUvarint16(5, uint16(m.AtBidOrAsk()))
	w.WriteBool(6, m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataTickRecordResponse_IntPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 806)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, float64(m.DateTime()))
	w.WriteVarint32(3, m.Price())
	w.WriteVarint32(4, m.Volume())
	w.WriteUvarint16(5, uint16(m.AtBidOrAsk()))
	w.WriteBool(6, m.IsFinalRecord())
	return w.Finish(), nil
}
