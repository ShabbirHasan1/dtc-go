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

func (m TradeAccountResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":401,\"F\":["...)
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestID())
	return w.Finish(), nil
}

func (m TradeAccountResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":401,\"F\":["...)
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":401,\"F\":["...)
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestID())
	return w.Finish(), nil
}

func (m TradeAccountResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":401,\"F\":["...)
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 401)
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("RequestID", m.RequestID())
	return w.Finish(), nil
}

func (m TradeAccountResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 401)
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("RequestID", m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 401)
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("RequestID", m.RequestID())
	return w.Finish(), nil
}

func (m TradeAccountResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 401)
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("RequestID", m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 401 {
		return message.ErrWrongType
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 401 {
		return message.ErrWrongType
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 401 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "RequestID":
			m.SetRequestID(r.Int32())
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

func (m *TradeAccountResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 401 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "RequestID":
			m.SetRequestID(r.Int32())
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

func (m TradeAccountResponseFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":401,\"F\":["...)
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestID())
	return w.Finish(), nil
}

func (m TradeAccountResponseFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":401,\"F\":["...)
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountResponseFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":401,\"F\":["...)
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestID())
	return w.Finish(), nil
}

func (m TradeAccountResponseFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":401,\"F\":["...)
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 401)
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("RequestID", m.RequestID())
	return w.Finish(), nil
}

func (m TradeAccountResponseFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 401)
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("RequestID", m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountResponseFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 401)
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("RequestID", m.RequestID())
	return w.Finish(), nil
}

func (m TradeAccountResponseFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 401)
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("RequestID", m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountResponseFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 401 {
		return message.ErrWrongType
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountResponseFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 401 {
		return message.ErrWrongType
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountResponseFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 401 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "RequestID":
			m.SetRequestID(r.Int32())
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

func (m *TradeAccountResponseFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 401 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "RequestID":
			m.SetRequestID(r.Int32())
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

func (m *TradeAccountResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetTotalNumberMessages(r.ReadInt32())
		case 2:
			m.SetMessageNumber(r.ReadInt32())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetRequestID(r.ReadInt32())
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

func (m *TradeAccountResponsePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetTotalNumberMessages(r.ReadInt32())
		case 2:
			m.SetMessageNumber(r.ReadInt32())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetRequestID(r.ReadInt32())
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

func (m TradeAccountResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 401)
	w.WriteVarint32(1, m.TotalNumberMessages())
	w.WriteVarint32(2, m.MessageNumber())
	w.WriteString(3, m.TradeAccount())
	w.WriteVarint32(4, m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountResponsePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 401)
	w.WriteVarint32(1, m.TotalNumberMessages())
	w.WriteVarint32(2, m.MessageNumber())
	w.WriteString(3, m.TradeAccount())
	w.WriteVarint32(4, m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountResponseFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetTotalNumberMessages(r.ReadInt32())
		case 2:
			m.SetMessageNumber(r.ReadInt32())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetRequestID(r.ReadInt32())
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

func (m *TradeAccountResponseFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetTotalNumberMessages(r.ReadInt32())
		case 2:
			m.SetMessageNumber(r.ReadInt32())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetRequestID(r.ReadInt32())
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

func (m TradeAccountResponseFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 401)
	w.WriteVarint32(1, m.TotalNumberMessages())
	w.WriteVarint32(2, m.MessageNumber())
	w.WriteString(3, m.TradeAccount())
	w.WriteVarint32(4, m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountResponseFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 401)
	w.WriteVarint32(1, m.TotalNumberMessages())
	w.WriteVarint32(2, m.MessageNumber())
	w.WriteString(3, m.TradeAccount())
	w.WriteVarint32(4, m.RequestID())
	return w.Finish(), nil
}
