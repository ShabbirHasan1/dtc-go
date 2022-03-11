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

func (m CancelOrder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":203,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m CancelOrderBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":203,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelOrderPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":203,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m CancelOrderPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":203,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelOrder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 203)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m CancelOrderBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 203)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelOrderPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 203)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m CancelOrderPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 203)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelOrderBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 203 {
		return message.ErrWrongType
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelOrderPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 203 {
		return message.ErrWrongType
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelOrderBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 203 {
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
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
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

func (m *CancelOrderPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 203 {
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
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
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

func (m CancelOrderFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":203,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m CancelOrderFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":203,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelOrderFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":203,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m CancelOrderFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":203,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelOrderFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 203)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m CancelOrderFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 203)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelOrderFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 203)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m CancelOrderFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 203)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelOrderFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 203 {
		return message.ErrWrongType
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelOrderFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 203 {
		return message.ErrWrongType
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelOrderFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 203 {
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
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
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

func (m *CancelOrderFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 203 {
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
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
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

func (m *CancelOrderBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 2:
			m.SetClientOrderID(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
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

func (m *CancelOrderPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 2:
			m.SetClientOrderID(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
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

func (m CancelOrderBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 203)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteString(3, m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelOrderPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 203)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteString(3, m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelOrderFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 2:
			m.SetClientOrderID(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
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

func (m *CancelOrderFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 2:
			m.SetClientOrderID(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
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

func (m CancelOrderFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 203)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteString(3, m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelOrderFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 203)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteString(3, m.TradeAccount())
	return w.Finish(), nil
}
