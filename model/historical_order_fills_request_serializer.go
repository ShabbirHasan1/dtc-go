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

func (m HistoricalOrderFillsRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":303,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfDays())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	return w.Finish(), nil
}

func (m HistoricalOrderFillsRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":303,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfDays())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillsRequestPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":303,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfDays())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	return w.Finish(), nil
}

func (m HistoricalOrderFillsRequestPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":303,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfDays())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillsRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 303)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("NumberOfDays", m.NumberOfDays())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	return w.Finish(), nil
}

func (m HistoricalOrderFillsRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 303)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("NumberOfDays", m.NumberOfDays())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillsRequestPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 303)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("NumberOfDays", m.NumberOfDays())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	return w.Finish(), nil
}

func (m HistoricalOrderFillsRequestPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 303)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("NumberOfDays", m.NumberOfDays())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillsRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 303 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumberOfDays(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillsRequestPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 303 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumberOfDays(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillsRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 303 {
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
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "NumberOfDays":
			m.SetNumberOfDays(r.Int32())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTime(r.Int64()))
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

func (m *HistoricalOrderFillsRequestPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 303 {
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
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "NumberOfDays":
			m.SetNumberOfDays(r.Int32())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTime(r.Int64()))
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

func (m HistoricalOrderFillsRequestFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":303,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfDays())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	return w.Finish(), nil
}

func (m HistoricalOrderFillsRequestFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":303,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfDays())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillsRequestFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":303,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfDays())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	return w.Finish(), nil
}

func (m HistoricalOrderFillsRequestFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":303,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumberOfDays())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillsRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 303)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("NumberOfDays", m.NumberOfDays())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	return w.Finish(), nil
}

func (m HistoricalOrderFillsRequestFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 303)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("NumberOfDays", m.NumberOfDays())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillsRequestFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 303)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("NumberOfDays", m.NumberOfDays())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	return w.Finish(), nil
}

func (m HistoricalOrderFillsRequestFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 303)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("NumberOfDays", m.NumberOfDays())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillsRequestFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 303 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumberOfDays(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillsRequestFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 303 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumberOfDays(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillsRequestFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 303 {
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
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "NumberOfDays":
			m.SetNumberOfDays(r.Int32())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTime(r.Int64()))
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

func (m *HistoricalOrderFillsRequestFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 303 {
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
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "NumberOfDays":
			m.SetNumberOfDays(r.Int32())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTime(r.Int64()))
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

func (m *HistoricalOrderFillsRequestBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 3:
			m.SetNumberOfDays(r.ReadInt32())
		case 4:
			m.SetTradeAccount(r.ReadString())
		case 5:
			m.SetStartDateTime(DateTime(r.ReadInt64()))
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

func (m *HistoricalOrderFillsRequestPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 3:
			m.SetNumberOfDays(r.ReadInt32())
		case 4:
			m.SetTradeAccount(r.ReadString())
		case 5:
			m.SetStartDateTime(DateTime(r.ReadInt64()))
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

func (m HistoricalOrderFillsRequestBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 303)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.ServerOrderID())
	w.WriteVarint32(3, m.NumberOfDays())
	w.WriteString(4, m.TradeAccount())
	w.WriteFixed64Int64(5, int64(m.StartDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillsRequestPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 303)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.ServerOrderID())
	w.WriteVarint32(3, m.NumberOfDays())
	w.WriteString(4, m.TradeAccount())
	w.WriteFixed64Int64(5, int64(m.StartDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillsRequestFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 3:
			m.SetNumberOfDays(r.ReadInt32())
		case 4:
			m.SetTradeAccount(r.ReadString())
		case 5:
			m.SetStartDateTime(DateTime(r.ReadInt64()))
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

func (m *HistoricalOrderFillsRequestFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 3:
			m.SetNumberOfDays(r.ReadInt32())
		case 4:
			m.SetTradeAccount(r.ReadString())
		case 5:
			m.SetStartDateTime(DateTime(r.ReadInt64()))
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

func (m HistoricalOrderFillsRequestFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 303)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.ServerOrderID())
	w.WriteVarint32(3, m.NumberOfDays())
	w.WriteString(4, m.TradeAccount())
	w.WriteFixed64Int64(5, int64(m.StartDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillsRequestFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 303)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.ServerOrderID())
	w.WriteVarint32(3, m.NumberOfDays())
	w.WriteString(4, m.TradeAccount())
	w.WriteFixed64Int64(5, int64(m.StartDateTime()))
	return w.Finish(), nil
}