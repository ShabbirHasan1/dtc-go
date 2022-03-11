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

func (m ExchangeListResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":501,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	return w.Finish(), nil
}

func (m ExchangeListResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":501,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ExchangeListResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":501,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	return w.Finish(), nil
}

func (m ExchangeListResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":501,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ExchangeListResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 501)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.StringField("Description", m.Description())
	return w.Finish(), nil
}

func (m ExchangeListResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 501)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.StringField("Description", m.Description())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ExchangeListResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 501)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.StringField("Description", m.Description())
	return w.Finish(), nil
}

func (m ExchangeListResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 501)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.StringField("Description", m.Description())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ExchangeListResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 501 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalMessage(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ExchangeListResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 501 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalMessage(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ExchangeListResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 501 {
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
		case "Exchange":
			m.SetExchange(r.String())
		case "IsFinalMessage":
			m.SetIsFinalMessage(r.Bool())
		case "Description":
			m.SetDescription(r.String())
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

func (m *ExchangeListResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 501 {
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
		case "Exchange":
			m.SetExchange(r.String())
		case "IsFinalMessage":
			m.SetIsFinalMessage(r.Bool())
		case "Description":
			m.SetDescription(r.String())
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

func (m ExchangeListResponseFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":501,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	return w.Finish(), nil
}

func (m ExchangeListResponseFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":501,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ExchangeListResponseFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":501,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	return w.Finish(), nil
}

func (m ExchangeListResponseFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":501,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.String(m.Description())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ExchangeListResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 501)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.StringField("Description", m.Description())
	return w.Finish(), nil
}

func (m ExchangeListResponseFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 501)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.StringField("Description", m.Description())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ExchangeListResponseFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 501)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.StringField("Description", m.Description())
	return w.Finish(), nil
}

func (m ExchangeListResponseFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 501)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.BoolField("IsFinalMessage", m.IsFinalMessage())
	w.StringField("Description", m.Description())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ExchangeListResponseFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 501 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalMessage(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ExchangeListResponseFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 501 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalMessage(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetDescription(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ExchangeListResponseFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 501 {
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
		case "Exchange":
			m.SetExchange(r.String())
		case "IsFinalMessage":
			m.SetIsFinalMessage(r.Bool())
		case "Description":
			m.SetDescription(r.String())
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

func (m *ExchangeListResponseFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 501 {
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
		case "Exchange":
			m.SetExchange(r.String())
		case "IsFinalMessage":
			m.SetIsFinalMessage(r.Bool())
		case "Description":
			m.SetDescription(r.String())
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

func (m *ExchangeListResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetExchange(r.ReadString())
		case 3:
			m.SetIsFinalMessage(r.ReadBool())
		case 4:
			m.SetDescription(r.ReadString())
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

func (m *ExchangeListResponsePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetExchange(r.ReadString())
		case 3:
			m.SetIsFinalMessage(r.ReadBool())
		case 4:
			m.SetDescription(r.ReadString())
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

func (m ExchangeListResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 501)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Exchange())
	w.WriteBool(3, m.IsFinalMessage())
	w.WriteString(4, m.Description())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ExchangeListResponsePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 501)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Exchange())
	w.WriteBool(3, m.IsFinalMessage())
	w.WriteString(4, m.Description())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ExchangeListResponseFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetExchange(r.ReadString())
		case 3:
			m.SetIsFinalMessage(r.ReadBool())
		case 4:
			m.SetDescription(r.ReadString())
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

func (m *ExchangeListResponseFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetExchange(r.ReadString())
		case 3:
			m.SetIsFinalMessage(r.ReadBool())
		case 4:
			m.SetDescription(r.ReadString())
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

func (m ExchangeListResponseFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 501)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Exchange())
	w.WriteBool(3, m.IsFinalMessage())
	w.WriteString(4, m.Description())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ExchangeListResponseFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 501)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Exchange())
	w.WriteBool(3, m.IsFinalMessage())
	w.WriteString(4, m.Description())
	return w.Finish(), nil
}
