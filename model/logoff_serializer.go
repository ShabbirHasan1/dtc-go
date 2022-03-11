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

func (m Logoff) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":5,\"F\":["...)
	w.String(m.Reason())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotReconnect())
	return w.Finish(), nil
}

func (m LogoffBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":5,\"F\":["...)
	w.String(m.Reason())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogoffPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":5,\"F\":["...)
	w.String(m.Reason())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotReconnect())
	return w.Finish(), nil
}

func (m LogoffPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":5,\"F\":["...)
	w.String(m.Reason())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m Logoff) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 5)
	w.StringField("Reason", m.Reason())
	w.Uint8Field("DoNotReconnect", m.DoNotReconnect())
	return w.Finish(), nil
}

func (m LogoffBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 5)
	w.StringField("Reason", m.Reason())
	w.Uint8Field("DoNotReconnect", m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogoffPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 5)
	w.StringField("Reason", m.Reason())
	w.Uint8Field("DoNotReconnect", m.DoNotReconnect())
	return w.Finish(), nil
}

func (m LogoffPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 5)
	w.StringField("Reason", m.Reason())
	w.Uint8Field("DoNotReconnect", m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogoffBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 5 {
		return message.ErrWrongType
	}
	m.SetReason(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDoNotReconnect(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogoffPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 5 {
		return message.ErrWrongType
	}
	m.SetReason(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDoNotReconnect(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogoffBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 5 {
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
		case "Reason":
			m.SetReason(r.String())
		case "DoNotReconnect":
			m.SetDoNotReconnect(r.Uint8())
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

func (m *LogoffPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 5 {
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
		case "Reason":
			m.SetReason(r.String())
		case "DoNotReconnect":
			m.SetDoNotReconnect(r.Uint8())
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
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogoffFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":5,\"F\":["...)
	w.String(m.Reason())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotReconnect())
	return w.Finish(), nil
}

func (m LogoffFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":5,\"F\":["...)
	w.String(m.Reason())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogoffFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":5,\"F\":["...)
	w.String(m.Reason())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotReconnect())
	return w.Finish(), nil
}

func (m LogoffFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":5,\"F\":["...)
	w.String(m.Reason())
	w.Data = append(w.Data, ',')
	w.Uint8(m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogoffFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 5)
	w.StringField("Reason", m.Reason())
	w.Uint8Field("DoNotReconnect", m.DoNotReconnect())
	return w.Finish(), nil
}

func (m LogoffFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 5)
	w.StringField("Reason", m.Reason())
	w.Uint8Field("DoNotReconnect", m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogoffFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 5)
	w.StringField("Reason", m.Reason())
	w.Uint8Field("DoNotReconnect", m.DoNotReconnect())
	return w.Finish(), nil
}

func (m LogoffFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 5)
	w.StringField("Reason", m.Reason())
	w.Uint8Field("DoNotReconnect", m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogoffFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 5 {
		return message.ErrWrongType
	}
	m.SetReason(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDoNotReconnect(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogoffFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 5 {
		return message.ErrWrongType
	}
	m.SetReason(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDoNotReconnect(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogoffFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 5 {
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
		case "Reason":
			m.SetReason(r.String())
		case "DoNotReconnect":
			m.SetDoNotReconnect(r.Uint8())
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

func (m *LogoffFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 5 {
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
		case "Reason":
			m.SetReason(r.String())
		case "DoNotReconnect":
			m.SetDoNotReconnect(r.Uint8())
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

func (m *LogoffBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetReason(r.ReadString())
		case 2:
			m.SetDoNotReconnect(r.ReadUint8())
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

func (m *LogoffPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetReason(r.ReadString())
		case 2:
			m.SetDoNotReconnect(r.ReadUint8())
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

func (m LogoffBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 5)
	w.WriteString(1, m.Reason())
	w.WriteUvarint8(2, m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogoffPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 5)
	w.WriteString(1, m.Reason())
	w.WriteUvarint8(2, m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogoffFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetReason(r.ReadString())
		case 2:
			m.SetDoNotReconnect(r.ReadUint8())
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

func (m *LogoffFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetReason(r.ReadString())
		case 2:
			m.SetDoNotReconnect(r.ReadUint8())
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

func (m LogoffFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 5)
	w.WriteString(1, m.Reason())
	w.WriteUvarint8(2, m.DoNotReconnect())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogoffFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 5)
	w.WriteString(1, m.Reason())
	w.WriteUvarint8(2, m.DoNotReconnect())
	return w.Finish(), nil
}
