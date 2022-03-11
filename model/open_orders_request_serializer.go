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

func (m OpenOrdersRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":300,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestAllOrders())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m OpenOrdersRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":300,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestAllOrders())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m OpenOrdersRequestPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":300,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestAllOrders())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m OpenOrdersRequestPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":300,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestAllOrders())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m OpenOrdersRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 300)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RequestAllOrders", m.RequestAllOrders())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m OpenOrdersRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 300)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RequestAllOrders", m.RequestAllOrders())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m OpenOrdersRequestPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 300)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RequestAllOrders", m.RequestAllOrders())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m OpenOrdersRequestPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 300)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RequestAllOrders", m.RequestAllOrders())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *OpenOrdersRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 300 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestAllOrders(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
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

func (m *OpenOrdersRequestPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 300 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestAllOrders(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
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

func (m *OpenOrdersRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 300 {
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
		case "RequestAllOrders":
			m.SetRequestAllOrders(r.Int32())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
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

func (m *OpenOrdersRequestPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 300 {
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
		case "RequestAllOrders":
			m.SetRequestAllOrders(r.Int32())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
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

func (m OpenOrdersRequestFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":300,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestAllOrders())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m OpenOrdersRequestFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":300,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestAllOrders())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m OpenOrdersRequestFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":300,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestAllOrders())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m OpenOrdersRequestFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":300,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.RequestAllOrders())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m OpenOrdersRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 300)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RequestAllOrders", m.RequestAllOrders())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m OpenOrdersRequestFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 300)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RequestAllOrders", m.RequestAllOrders())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m OpenOrdersRequestFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 300)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RequestAllOrders", m.RequestAllOrders())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m OpenOrdersRequestFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 300)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RequestAllOrders", m.RequestAllOrders())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *OpenOrdersRequestFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 300 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestAllOrders(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
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

func (m *OpenOrdersRequestFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 300 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRequestAllOrders(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
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

func (m *OpenOrdersRequestFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 300 {
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
		case "RequestAllOrders":
			m.SetRequestAllOrders(r.Int32())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
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

func (m *OpenOrdersRequestFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 300 {
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
		case "RequestAllOrders":
			m.SetRequestAllOrders(r.Int32())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
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

func (m *OpenOrdersRequestBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestAllOrders(r.ReadInt32())
		case 3:
			m.SetServerOrderID(r.ReadString())
		case 4:
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

func (m *OpenOrdersRequestPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestAllOrders(r.ReadInt32())
		case 3:
			m.SetServerOrderID(r.ReadString())
		case 4:
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

func (m OpenOrdersRequestBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 300)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.RequestAllOrders())
	w.WriteString(3, m.ServerOrderID())
	w.WriteString(4, m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m OpenOrdersRequestPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 300)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.RequestAllOrders())
	w.WriteString(3, m.ServerOrderID())
	w.WriteString(4, m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *OpenOrdersRequestFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestAllOrders(r.ReadInt32())
		case 3:
			m.SetServerOrderID(r.ReadString())
		case 4:
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

func (m *OpenOrdersRequestFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestAllOrders(r.ReadInt32())
		case 3:
			m.SetServerOrderID(r.ReadString())
		case 4:
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

func (m OpenOrdersRequestFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 300)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.RequestAllOrders())
	w.WriteString(3, m.ServerOrderID())
	w.WriteString(4, m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m OpenOrdersRequestFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 300)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.RequestAllOrders())
	w.WriteString(3, m.ServerOrderID())
	w.WriteString(4, m.TradeAccount())
	return w.Finish(), nil
}
