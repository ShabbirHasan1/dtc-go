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

func (m MarketOrdersRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":150,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

func (m MarketOrdersRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":150,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersRequestPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":150,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

func (m MarketOrdersRequestPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":150,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 150)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("SendQuantitiesGreaterOrEqualTo", m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

func (m MarketOrdersRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 150)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("SendQuantitiesGreaterOrEqualTo", m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersRequestPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 150)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("SendQuantitiesGreaterOrEqualTo", m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

func (m MarketOrdersRequestPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 150)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("SendQuantitiesGreaterOrEqualTo", m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 150 {
		return message.ErrWrongType
	}
	m.SetRequestAction(RequestActionEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolID(r.Uint32())
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
	m.SetSendQuantitiesGreaterOrEqualTo(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersRequestPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 150 {
		return message.ErrWrongType
	}
	m.SetRequestAction(RequestActionEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolID(r.Uint32())
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
	m.SetSendQuantitiesGreaterOrEqualTo(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 150 {
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
		case "RequestAction":
			m.SetRequestAction(RequestActionEnum(r.Int32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SendQuantitiesGreaterOrEqualTo":
			m.SetSendQuantitiesGreaterOrEqualTo(r.Uint32())
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

func (m *MarketOrdersRequestPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 150 {
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
		case "RequestAction":
			m.SetRequestAction(RequestActionEnum(r.Int32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SendQuantitiesGreaterOrEqualTo":
			m.SetSendQuantitiesGreaterOrEqualTo(r.Uint32())
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

func (m MarketOrdersRequestFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":150,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

func (m MarketOrdersRequestFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":150,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersRequestFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":150,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

func (m MarketOrdersRequestFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":150,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 150)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("SendQuantitiesGreaterOrEqualTo", m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

func (m MarketOrdersRequestFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 150)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("SendQuantitiesGreaterOrEqualTo", m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersRequestFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 150)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("SendQuantitiesGreaterOrEqualTo", m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

func (m MarketOrdersRequestFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 150)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("SendQuantitiesGreaterOrEqualTo", m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersRequestFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 150 {
		return message.ErrWrongType
	}
	m.SetRequestAction(RequestActionEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolID(r.Uint32())
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
	m.SetSendQuantitiesGreaterOrEqualTo(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersRequestFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 150 {
		return message.ErrWrongType
	}
	m.SetRequestAction(RequestActionEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbolID(r.Uint32())
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
	m.SetSendQuantitiesGreaterOrEqualTo(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersRequestFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 150 {
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
		case "RequestAction":
			m.SetRequestAction(RequestActionEnum(r.Int32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SendQuantitiesGreaterOrEqualTo":
			m.SetSendQuantitiesGreaterOrEqualTo(r.Uint32())
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

func (m *MarketOrdersRequestFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 150 {
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
		case "RequestAction":
			m.SetRequestAction(RequestActionEnum(r.Int32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SendQuantitiesGreaterOrEqualTo":
			m.SetSendQuantitiesGreaterOrEqualTo(r.Uint32())
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

func (m *MarketOrdersRequestBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestAction(RequestActionEnum(r.ReadInt32()))
		case 2:
			m.SetSymbolID(r.ReadUint32())
		case 3:
			m.SetSymbol(r.ReadString())
		case 4:
			m.SetExchange(r.ReadString())
		case 5:
			m.SetSendQuantitiesGreaterOrEqualTo(r.ReadUint32())
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

func (m *MarketOrdersRequestPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestAction(RequestActionEnum(r.ReadInt32()))
		case 2:
			m.SetSymbolID(r.ReadUint32())
		case 3:
			m.SetSymbol(r.ReadString())
		case 4:
			m.SetExchange(r.ReadString())
		case 5:
			m.SetSendQuantitiesGreaterOrEqualTo(r.ReadUint32())
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

func (m MarketOrdersRequestBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 150)
	w.WriteVarint32(1, int32(m.RequestAction()))
	w.WriteUvarint32(2, m.SymbolID())
	w.WriteString(3, m.Symbol())
	w.WriteString(4, m.Exchange())
	w.WriteUvarint32(5, m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersRequestPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 150)
	w.WriteVarint32(1, int32(m.RequestAction()))
	w.WriteUvarint32(2, m.SymbolID())
	w.WriteString(3, m.Symbol())
	w.WriteString(4, m.Exchange())
	w.WriteUvarint32(5, m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersRequestFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestAction(RequestActionEnum(r.ReadInt32()))
		case 2:
			m.SetSymbolID(r.ReadUint32())
		case 3:
			m.SetSymbol(r.ReadString())
		case 4:
			m.SetExchange(r.ReadString())
		case 5:
			m.SetSendQuantitiesGreaterOrEqualTo(r.ReadUint32())
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

func (m *MarketOrdersRequestFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestAction(RequestActionEnum(r.ReadInt32()))
		case 2:
			m.SetSymbolID(r.ReadUint32())
		case 3:
			m.SetSymbol(r.ReadString())
		case 4:
			m.SetExchange(r.ReadString())
		case 5:
			m.SetSendQuantitiesGreaterOrEqualTo(r.ReadUint32())
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

func (m MarketOrdersRequestFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 150)
	w.WriteVarint32(1, int32(m.RequestAction()))
	w.WriteUvarint32(2, m.SymbolID())
	w.WriteString(3, m.Symbol())
	w.WriteString(4, m.Exchange())
	w.WriteUvarint32(5, m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersRequestFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 150)
	w.WriteVarint32(1, int32(m.RequestAction()))
	w.WriteUvarint32(2, m.SymbolID())
	w.WriteString(3, m.Symbol())
	w.WriteString(4, m.Exchange())
	w.WriteUvarint32(5, m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}
