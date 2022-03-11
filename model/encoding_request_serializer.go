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

func (m EncodingRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":6,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Encoding()))
	w.Data = append(w.Data, ',')
	w.String(m.ProtocolType())
	return w.Finish(), nil
}

func (m EncodingRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":6,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Encoding()))
	w.Data = append(w.Data, ',')
	w.String(m.ProtocolType())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m EncodingRequestPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":6,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Encoding()))
	w.Data = append(w.Data, ',')
	w.String(m.ProtocolType())
	return w.Finish(), nil
}

func (m EncodingRequestPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":6,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Encoding()))
	w.Data = append(w.Data, ',')
	w.String(m.ProtocolType())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m EncodingRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 6)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Encoding", int32(m.Encoding()))
	w.StringField("ProtocolType", m.ProtocolType())
	return w.Finish(), nil
}

func (m EncodingRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 6)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Encoding", int32(m.Encoding()))
	w.StringField("ProtocolType", m.ProtocolType())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m EncodingRequestPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 6)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Encoding", int32(m.Encoding()))
	w.StringField("ProtocolType", m.ProtocolType())
	return w.Finish(), nil
}

func (m EncodingRequestPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 6)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Encoding", int32(m.Encoding()))
	w.StringField("ProtocolType", m.ProtocolType())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *EncodingRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 6 {
		return message.ErrWrongType
	}
	m.SetProtocolVersion(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetEncoding(EncodingEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetProtocolType(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *EncodingRequestPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 6 {
		return message.ErrWrongType
	}
	m.SetProtocolVersion(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetEncoding(EncodingEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetProtocolType(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *EncodingRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 6 {
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
		case "ProtocolVersion":
			m.SetProtocolVersion(r.Int32())
		case "Encoding":
			m.SetEncoding(EncodingEnum(r.Int32()))
		case "ProtocolType":
			m.SetProtocolType(r.String())
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

func (m *EncodingRequestPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 6 {
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
		case "ProtocolVersion":
			m.SetProtocolVersion(r.Int32())
		case "Encoding":
			m.SetEncoding(EncodingEnum(r.Int32()))
		case "ProtocolType":
			m.SetProtocolType(r.String())
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

func (m *EncodingRequestBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetProtocolVersion(r.ReadInt32())
		case 2:
			m.SetEncoding(EncodingEnum(r.ReadInt32()))
		case 3:
			m.SetProtocolType(r.ReadString())
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

func (m *EncodingRequestPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetProtocolVersion(r.ReadInt32())
		case 2:
			m.SetEncoding(EncodingEnum(r.ReadInt32()))
		case 3:
			m.SetProtocolType(r.ReadString())
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

func (m EncodingRequestBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 6)
	w.WriteVarint32(1, m.ProtocolVersion())
	w.WriteVarint32(2, int32(m.Encoding()))
	w.WriteString(3, m.ProtocolType())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m EncodingRequestPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 6)
	w.WriteVarint32(1, m.ProtocolVersion())
	w.WriteVarint32(2, int32(m.Encoding()))
	w.WriteString(3, m.ProtocolType())
	return w.Finish(), nil
}
