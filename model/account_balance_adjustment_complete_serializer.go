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

func (m AccountBalanceAdjustmentComplete) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":609,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(m.TransactionID())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentCompleteBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":609,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(m.TransactionID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentCompletePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":609,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(m.TransactionID())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentCompletePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":609,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int64(m.TransactionID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentComplete) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 609)
	w.Int32Field("RequestID", m.RequestID())
	w.Int64Field("TransactionID", m.TransactionID())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentCompleteBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 609)
	w.Int32Field("RequestID", m.RequestID())
	w.Int64Field("TransactionID", m.TransactionID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentCompletePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 609)
	w.Int32Field("RequestID", m.RequestID())
	w.Int64Field("TransactionID", m.TransactionID())
	return w.Finish(), nil
}

func (m AccountBalanceAdjustmentCompletePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 609)
	w.Int32Field("RequestID", m.RequestID())
	w.Int64Field("TransactionID", m.TransactionID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentCompleteBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 609 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionID(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentCompletePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 609 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionID(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentCompleteBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 609 {
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
		case "TransactionID":
			m.SetTransactionID(r.Int64())
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

func (m *AccountBalanceAdjustmentCompletePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 609 {
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
		case "TransactionID":
			m.SetTransactionID(r.Int64())
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

func (m *AccountBalanceAdjustmentCompleteBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetTransactionID(r.ReadInt64())
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

func (m *AccountBalanceAdjustmentCompletePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetTransactionID(r.ReadInt64())
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

func (m AccountBalanceAdjustmentCompleteBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 609)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint64(2, m.TransactionID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentCompletePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 609)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint64(2, m.TransactionID())
	return w.Finish(), nil
}
