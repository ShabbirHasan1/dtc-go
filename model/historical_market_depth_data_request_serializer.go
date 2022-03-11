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

func (m HistoricalMarketDepthDataRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":900,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.EndDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Integer_1())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":900,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.EndDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataRequestPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":900,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.EndDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Integer_1())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataRequestPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":900,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.EndDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 900)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("EndDateTime", int64(m.EndDateTime()))
	w.Uint8Field("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("Integer_1", m.Integer_1())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 900)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("EndDateTime", int64(m.EndDateTime()))
	w.Uint8Field("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("Integer_1", m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataRequestPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 900)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("EndDateTime", int64(m.EndDateTime()))
	w.Uint8Field("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("Integer_1", m.Integer_1())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataRequestPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 900)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("EndDateTime", int64(m.EndDateTime()))
	w.Uint8Field("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("Integer_1", m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 900 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUseZLibCompression(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInteger_1(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataRequestPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 900 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUseZLibCompression(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInteger_1(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 900 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "EndDateTime":
			m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "UseZLibCompression":
			m.SetUseZLibCompression(r.Uint8())
		case "Integer_1":
			m.SetInteger_1(r.Uint8())
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

func (m *HistoricalMarketDepthDataRequestPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 900 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "EndDateTime":
			m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "UseZLibCompression":
			m.SetUseZLibCompression(r.Uint8())
		case "Integer_1":
			m.SetInteger_1(r.Uint8())
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

func (m HistoricalMarketDepthDataRequestFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":900,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.EndDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Integer_1())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataRequestFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":900,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.EndDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataRequestFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":900,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.EndDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Integer_1())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":900,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.StartDateTime()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.EndDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.UseZLibCompression())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 900)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("EndDateTime", int64(m.EndDateTime()))
	w.Uint8Field("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("Integer_1", m.Integer_1())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataRequestFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 900)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("EndDateTime", int64(m.EndDateTime()))
	w.Uint8Field("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("Integer_1", m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataRequestFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 900)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("EndDateTime", int64(m.EndDateTime()))
	w.Uint8Field("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("Integer_1", m.Integer_1())
	return w.Finish(), nil
}

func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 900)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("EndDateTime", int64(m.EndDateTime()))
	w.Uint8Field("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("Integer_1", m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataRequestFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 900 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUseZLibCompression(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInteger_1(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataRequestFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 900 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUseZLibCompression(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInteger_1(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataRequestFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 900 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "EndDateTime":
			m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "UseZLibCompression":
			m.SetUseZLibCompression(r.Uint8())
		case "Integer_1":
			m.SetInteger_1(r.Uint8())
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

func (m *HistoricalMarketDepthDataRequestFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 900 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "EndDateTime":
			m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "UseZLibCompression":
			m.SetUseZLibCompression(r.Uint8())
		case "Integer_1":
			m.SetInteger_1(r.Uint8())
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

func (m *HistoricalMarketDepthDataRequestBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 3:
			m.SetExchange(r.ReadString())
		case 4:
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 5:
			m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 6:
			m.SetUseZLibCompression(r.ReadUint8())
		case 7:
			m.SetInteger_1(r.ReadUint8())
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

func (m *HistoricalMarketDepthDataRequestPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 3:
			m.SetExchange(r.ReadString())
		case 4:
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 5:
			m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 6:
			m.SetUseZLibCompression(r.ReadUint8())
		case 7:
			m.SetInteger_1(r.ReadUint8())
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

func (m HistoricalMarketDepthDataRequestBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 900)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Symbol())
	w.WriteString(3, m.Exchange())
	w.WriteVarint64(4, int64(m.StartDateTime()))
	w.WriteVarint64(5, int64(m.EndDateTime()))
	w.WriteUvarint8(6, m.UseZLibCompression())
	w.WriteUvarint8(7, m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataRequestPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 900)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Symbol())
	w.WriteString(3, m.Exchange())
	w.WriteVarint64(4, int64(m.StartDateTime()))
	w.WriteVarint64(5, int64(m.EndDateTime()))
	w.WriteUvarint8(6, m.UseZLibCompression())
	w.WriteUvarint8(7, m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataRequestFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 3:
			m.SetExchange(r.ReadString())
		case 4:
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 5:
			m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 6:
			m.SetUseZLibCompression(r.ReadUint8())
		case 7:
			m.SetInteger_1(r.ReadUint8())
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

func (m *HistoricalMarketDepthDataRequestFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 3:
			m.SetExchange(r.ReadString())
		case 4:
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 5:
			m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.ReadInt64()))
		case 6:
			m.SetUseZLibCompression(r.ReadUint8())
		case 7:
			m.SetInteger_1(r.ReadUint8())
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

func (m HistoricalMarketDepthDataRequestFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 900)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Symbol())
	w.WriteString(3, m.Exchange())
	w.WriteVarint64(4, int64(m.StartDateTime()))
	w.WriteVarint64(5, int64(m.EndDateTime()))
	w.WriteUvarint8(6, m.UseZLibCompression())
	w.WriteUvarint8(7, m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalMarketDepthDataRequestFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 900)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.Symbol())
	w.WriteString(3, m.Exchange())
	w.WriteVarint64(4, int64(m.StartDateTime()))
	w.WriteVarint64(5, int64(m.EndDateTime()))
	w.WriteUvarint8(6, m.UseZLibCompression())
	w.WriteUvarint8(7, m.Integer_1())
	return w.Finish(), nil
}
