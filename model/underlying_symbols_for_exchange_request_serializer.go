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

func (m UnderlyingSymbolsForExchangeRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":503,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	return w.Finish(), nil
}

func (m UnderlyingSymbolsForExchangeRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":503,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UnderlyingSymbolsForExchangeRequestPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":503,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	return w.Finish(), nil
}

func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":503,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UnderlyingSymbolsForExchangeRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 503)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	return w.Finish(), nil
}

func (m UnderlyingSymbolsForExchangeRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 503)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UnderlyingSymbolsForExchangeRequestPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 503)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	return w.Finish(), nil
}

func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 503)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UnderlyingSymbolsForExchangeRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 503 {
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
	m.SetSecurityType(SecurityTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UnderlyingSymbolsForExchangeRequestPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 503 {
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
	m.SetSecurityType(SecurityTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UnderlyingSymbolsForExchangeRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 503 {
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
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
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

func (m *UnderlyingSymbolsForExchangeRequestPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 503 {
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
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
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

func (m UnderlyingSymbolsForExchangeRequestFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":503,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	return w.Finish(), nil
}

func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":503,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UnderlyingSymbolsForExchangeRequestFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":503,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	return w.Finish(), nil
}

func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":503,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.SecurityType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UnderlyingSymbolsForExchangeRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 503)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	return w.Finish(), nil
}

func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 503)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UnderlyingSymbolsForExchangeRequestFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 503)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	return w.Finish(), nil
}

func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 503)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UnderlyingSymbolsForExchangeRequestFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 503 {
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
	m.SetSecurityType(SecurityTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 503 {
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
	m.SetSecurityType(SecurityTypeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UnderlyingSymbolsForExchangeRequestFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 503 {
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
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
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

func (m *UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 503 {
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
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
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

func (m *UnderlyingSymbolsForExchangeRequestBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSecurityType(SecurityTypeEnum(r.ReadInt32()))
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

func (m *UnderlyingSymbolsForExchangeRequestPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSecurityType(SecurityTypeEnum(r.ReadInt32()))
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

func (m UnderlyingSymbolsForExchangeRequestBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 503)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Exchange())
	w.WriteVarint32(3, int32(m.SecurityType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UnderlyingSymbolsForExchangeRequestPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 503)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Exchange())
	w.WriteVarint32(3, int32(m.SecurityType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UnderlyingSymbolsForExchangeRequestFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSecurityType(SecurityTypeEnum(r.ReadInt32()))
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

func (m *UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSecurityType(SecurityTypeEnum(r.ReadInt32()))
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

func (m UnderlyingSymbolsForExchangeRequestFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 503)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Exchange())
	w.WriteVarint32(3, int32(m.SecurityType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UnderlyingSymbolsForExchangeRequestFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 503)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Exchange())
	w.WriteVarint32(3, int32(m.SecurityType()))
	return w.Finish(), nil
}
