//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m EncodingRequestExtended) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":6,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Encoding()))
	w.Data = append(w.Data, ',')
	w.String(m.ProtocolType())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.UseCompression()))
	return w.Finish(), nil
}

func (m EncodingRequestExtendedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":6,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Encoding()))
	w.Data = append(w.Data, ',')
	w.String(m.ProtocolType())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.UseCompression()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m EncodingRequestExtendedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":6,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Encoding()))
	w.Data = append(w.Data, ',')
	w.String(m.ProtocolType())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.UseCompression()))
	return w.Finish(), nil
}

func (m EncodingRequestExtendedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":6,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.Encoding()))
	w.Data = append(w.Data, ',')
	w.String(m.ProtocolType())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.UseCompression()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m EncodingRequestExtended) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 6)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Encoding", int32(m.Encoding()))
	w.StringField("ProtocolType", m.ProtocolType())
	w.Int32Field("UseCompression", int32(m.UseCompression()))
	return w.Finish(), nil
}

func (m EncodingRequestExtendedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 6)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Encoding", int32(m.Encoding()))
	w.StringField("ProtocolType", m.ProtocolType())
	w.Int32Field("UseCompression", int32(m.UseCompression()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m EncodingRequestExtendedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 6)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Encoding", int32(m.Encoding()))
	w.StringField("ProtocolType", m.ProtocolType())
	w.Int32Field("UseCompression", int32(m.UseCompression()))
	return w.Finish(), nil
}

func (m EncodingRequestExtendedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 6)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Encoding", int32(m.Encoding()))
	w.StringField("ProtocolType", m.ProtocolType())
	w.Int32Field("UseCompression", int32(m.UseCompression()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *EncodingRequestExtendedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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
	m.SetUseCompression(UseCompressionEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *EncodingRequestExtendedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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
	m.SetUseCompression(UseCompressionEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *EncodingRequestExtendedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
		case "UseCompression":
			m.SetUseCompression(UseCompressionEnum(r.Int32()))
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

func (m *EncodingRequestExtendedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
		case "UseCompression":
			m.SetUseCompression(UseCompressionEnum(r.Int32()))
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