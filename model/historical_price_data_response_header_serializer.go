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

func (m HistoricalPriceDataResponseHeader) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":801,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RecordInterval()))
	w.Data = append(w.Data, ',')
	w.Bool(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoRecordsToReturn())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatPriceDivisor())
	return w.Finish(), nil
}

func (m HistoricalPriceDataResponseHeaderBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":801,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RecordInterval()))
	w.Data = append(w.Data, ',')
	w.Bool(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoRecordsToReturn())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatPriceDivisor())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataResponseHeaderPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":801,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RecordInterval()))
	w.Data = append(w.Data, ',')
	w.Bool(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoRecordsToReturn())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatPriceDivisor())
	return w.Finish(), nil
}

func (m HistoricalPriceDataResponseHeaderPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":801,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.RecordInterval()))
	w.Data = append(w.Data, ',')
	w.Bool(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoRecordsToReturn())
	w.Data = append(w.Data, ',')
	w.Float32(m.IntToFloatPriceDivisor())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataResponseHeader) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 801)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RecordInterval", int32(m.RecordInterval()))
	w.BoolField("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("NoRecordsToReturn", m.NoRecordsToReturn())
	w.Float32Field("IntToFloatPriceDivisor", m.IntToFloatPriceDivisor())
	return w.Finish(), nil
}

func (m HistoricalPriceDataResponseHeaderBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 801)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RecordInterval", int32(m.RecordInterval()))
	w.BoolField("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("NoRecordsToReturn", m.NoRecordsToReturn())
	w.Float32Field("IntToFloatPriceDivisor", m.IntToFloatPriceDivisor())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataResponseHeaderPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 801)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RecordInterval", int32(m.RecordInterval()))
	w.BoolField("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("NoRecordsToReturn", m.NoRecordsToReturn())
	w.Float32Field("IntToFloatPriceDivisor", m.IntToFloatPriceDivisor())
	return w.Finish(), nil
}

func (m HistoricalPriceDataResponseHeaderPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 801)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RecordInterval", int32(m.RecordInterval()))
	w.BoolField("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("NoRecordsToReturn", m.NoRecordsToReturn())
	w.Float32Field("IntToFloatPriceDivisor", m.IntToFloatPriceDivisor())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataResponseHeaderBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 801 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRecordInterval(HistoricalDataIntervalEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUseZLibCompression(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoRecordsToReturn(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIntToFloatPriceDivisor(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataResponseHeaderPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 801 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRecordInterval(HistoricalDataIntervalEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUseZLibCompression(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoRecordsToReturn(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIntToFloatPriceDivisor(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataResponseHeaderBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 801 {
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
		case "RecordInterval":
			m.SetRecordInterval(HistoricalDataIntervalEnum(r.Int32()))
		case "UseZLibCompression":
			m.SetUseZLibCompression(r.Bool())
		case "NoRecordsToReturn":
			m.SetNoRecordsToReturn(r.Uint8())
		case "IntToFloatPriceDivisor":
			m.SetIntToFloatPriceDivisor(r.Float32())
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

func (m *HistoricalPriceDataResponseHeaderPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 801 {
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
		case "RecordInterval":
			m.SetRecordInterval(HistoricalDataIntervalEnum(r.Int32()))
		case "UseZLibCompression":
			m.SetUseZLibCompression(r.Bool())
		case "NoRecordsToReturn":
			m.SetNoRecordsToReturn(r.Uint8())
		case "IntToFloatPriceDivisor":
			m.SetIntToFloatPriceDivisor(r.Float32())
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

func (m *HistoricalPriceDataResponseHeaderBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRecordInterval(HistoricalDataIntervalEnum(r.ReadInt32()))
		case 3:
			m.SetUseZLibCompression(r.ReadBool())
		case 4:
			m.SetNoRecordsToReturn(r.ReadUint8())
		case 5:
			m.SetIntToFloatPriceDivisor(r.ReadFloat32())
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

func (m *HistoricalPriceDataResponseHeaderPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRecordInterval(HistoricalDataIntervalEnum(r.ReadInt32()))
		case 3:
			m.SetUseZLibCompression(r.ReadBool())
		case 4:
			m.SetNoRecordsToReturn(r.ReadUint8())
		case 5:
			m.SetIntToFloatPriceDivisor(r.ReadFloat32())
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

func (m HistoricalPriceDataResponseHeaderBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 801)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, int32(m.RecordInterval()))
	w.WriteBool(3, m.UseZLibCompression())
	w.WriteUvarint8(4, m.NoRecordsToReturn())
	w.WriteFixed32Float32(5, m.IntToFloatPriceDivisor())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataResponseHeaderPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 801)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, int32(m.RecordInterval()))
	w.WriteBool(3, m.UseZLibCompression())
	w.WriteUvarint8(4, m.NoRecordsToReturn())
	w.WriteFixed32Float32(5, m.IntToFloatPriceDivisor())
	return w.Finish(), nil
}
