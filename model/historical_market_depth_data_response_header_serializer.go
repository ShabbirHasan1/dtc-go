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

func (m HistoricalMarketDepthDataResponseHeader) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":901,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoRecordsToReturn())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataResponseHeaderBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":901,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoRecordsToReturn())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataResponseHeaderPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":901,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoRecordsToReturn())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataResponseHeaderPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":901,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoRecordsToReturn())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataResponseHeader) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 901)
	w.Int32Field("RequestID", m.RequestID())
	w.BoolField("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("NoRecordsToReturn", m.NoRecordsToReturn())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataResponseHeaderBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 901)
	w.Int32Field("RequestID", m.RequestID())
	w.BoolField("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("NoRecordsToReturn", m.NoRecordsToReturn())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataResponseHeaderPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 901)
	w.Int32Field("RequestID", m.RequestID())
	w.BoolField("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("NoRecordsToReturn", m.NoRecordsToReturn())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataResponseHeaderPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 901)
	w.Int32Field("RequestID", m.RequestID())
	w.BoolField("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("NoRecordsToReturn", m.NoRecordsToReturn())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataResponseHeaderBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 901 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
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
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataResponseHeaderPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 901 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
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
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataResponseHeaderBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 901 {
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
		case "UseZLibCompression":
			m.SetUseZLibCompression(r.Bool())
		case "NoRecordsToReturn":
			m.SetNoRecordsToReturn(r.Uint8())
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

func (m *HistoricalMarketDepthDataResponseHeaderPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 901 {
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
		case "UseZLibCompression":
			m.SetUseZLibCompression(r.Bool())
		case "NoRecordsToReturn":
			m.SetNoRecordsToReturn(r.Uint8())
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

func (m *HistoricalMarketDepthDataResponseHeaderBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetUseZLibCompression(r.ReadBool())
		case 3:
			m.SetNoRecordsToReturn(r.ReadUint8())
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

func (m *HistoricalMarketDepthDataResponseHeaderPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetUseZLibCompression(r.ReadBool())
		case 3:
			m.SetNoRecordsToReturn(r.ReadUint8())
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

func (m HistoricalMarketDepthDataResponseHeaderBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 901)
	w.WriteVarint32(1, m.RequestID())
	w.WriteBool(2, m.UseZLibCompression())
	w.WriteUvarint8(3, m.NoRecordsToReturn())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataResponseHeaderPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 901)
	w.WriteVarint32(1, m.RequestID())
	w.WriteBool(2, m.UseZLibCompression())
	w.WriteUvarint8(3, m.NoRecordsToReturn())
	return w.Finish(), nil
}
