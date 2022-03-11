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

func (m HistoricalMarketDepthDataRecordResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":903,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.Command())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Flags())
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataRecordResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":903,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.Command())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Flags())
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataRecordResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":903,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.Command())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Flags())
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":903,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.Command())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Flags())
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumOrders())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataRecordResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 903)
	w.Int32Field("RequestID", m.RequestID())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Uint8Field("Command", m.Command())
	w.Uint8Field("Flags", m.Flags())
	w.Uint16Field("NumOrders", m.NumOrders())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Quantity", m.Quantity())
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataRecordResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 903)
	w.Int32Field("RequestID", m.RequestID())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Uint8Field("Command", m.Command())
	w.Uint8Field("Flags", m.Flags())
	w.Uint16Field("NumOrders", m.NumOrders())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Quantity", m.Quantity())
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataRecordResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 903)
	w.Int32Field("RequestID", m.RequestID())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Uint8Field("Command", m.Command())
	w.Uint8Field("Flags", m.Flags())
	w.Uint16Field("NumOrders", m.NumOrders())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Quantity", m.Quantity())
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 903)
	w.Int32Field("RequestID", m.RequestID())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Uint8Field("Command", m.Command())
	w.Uint8Field("Flags", m.Flags())
	w.Uint16Field("NumOrders", m.NumOrders())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Quantity", m.Quantity())
	w.BoolField("IsFinalRecord", m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataRecordResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 903 {
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
	m.SetCommand(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlags(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumOrders(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Uint32())
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

func (m *HistoricalMarketDepthDataRecordResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 903 {
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
	m.SetCommand(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetFlags(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumOrders(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Uint32())
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

func (m *HistoricalMarketDepthDataRecordResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 903 {
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
		case "Command":
			m.SetCommand(r.Uint8())
		case "Flags":
			m.SetFlags(r.Uint8())
		case "NumOrders":
			m.SetNumOrders(r.Uint16())
		case "Price":
			m.SetPrice(r.Float32())
		case "Quantity":
			m.SetQuantity(r.Uint32())
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

func (m *HistoricalMarketDepthDataRecordResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 903 {
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
		case "Command":
			m.SetCommand(r.Uint8())
		case "Flags":
			m.SetFlags(r.Uint8())
		case "NumOrders":
			m.SetNumOrders(r.Uint16())
		case "Price":
			m.SetPrice(r.Float32())
		case "Quantity":
			m.SetQuantity(r.Uint32())
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

func (m *HistoricalMarketDepthDataRecordResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetCommand(r.ReadUint8())
		case 4:
			m.SetFlags(r.ReadUint8())
		case 5:
			m.SetNumOrders(r.ReadUint16())
		case 6:
			m.SetPrice(r.ReadFloat32())
		case 7:
			m.SetQuantity(r.ReadUint32())
		case 8:
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

func (m *HistoricalMarketDepthDataRecordResponsePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetCommand(r.ReadUint8())
		case 4:
			m.SetFlags(r.ReadUint8())
		case 5:
			m.SetNumOrders(r.ReadUint16())
		case 6:
			m.SetPrice(r.ReadFloat32())
		case 7:
			m.SetQuantity(r.ReadUint32())
		case 8:
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

func (m HistoricalMarketDepthDataRecordResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 903)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint64(2, int64(m.StartDateTime()))
	w.WriteUvarint8(3, m.Command())
	w.WriteUvarint8(4, m.Flags())
	w.WriteUvarint16(5, m.NumOrders())
	w.WriteFixed32Float32(6, m.Price())
	w.WriteUvarint32(7, m.Quantity())
	w.WriteBool(8, m.IsFinalRecord())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataRecordResponsePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 903)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint64(2, int64(m.StartDateTime()))
	w.WriteUvarint8(3, m.Command())
	w.WriteUvarint8(4, m.Flags())
	w.WriteUvarint16(5, m.NumOrders())
	w.WriteFixed32Float32(6, m.Price())
	w.WriteUvarint32(7, m.Quantity())
	w.WriteBool(8, m.IsFinalRecord())
	return w.Finish(), nil
}
