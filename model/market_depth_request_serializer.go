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

func (m MarketDepthRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":102,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumLevels())
	return w.Finish(), nil
}

func (m MarketDepthRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":102,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumLevels())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthRequestPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":102,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumLevels())
	return w.Finish(), nil
}

func (m MarketDepthRequestPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":102,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumLevels())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 102)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("NumLevels", m.NumLevels())
	return w.Finish(), nil
}

func (m MarketDepthRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 102)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("NumLevels", m.NumLevels())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthRequestPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 102)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("NumLevels", m.NumLevels())
	return w.Finish(), nil
}

func (m MarketDepthRequestPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 102)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("NumLevels", m.NumLevels())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 102 {
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
	m.SetNumLevels(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthRequestPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 102 {
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
	m.SetNumLevels(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 102 {
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
		case "NumLevels":
			m.SetNumLevels(r.Int32())
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

func (m *MarketDepthRequestPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 102 {
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
		case "NumLevels":
			m.SetNumLevels(r.Int32())
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

func (m MarketDepthRequestFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":102,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumLevels())
	return w.Finish(), nil
}

func (m MarketDepthRequestFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":102,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumLevels())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthRequestFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":102,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumLevels())
	return w.Finish(), nil
}

func (m MarketDepthRequestFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":102,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumLevels())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 102)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("NumLevels", m.NumLevels())
	return w.Finish(), nil
}

func (m MarketDepthRequestFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 102)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("NumLevels", m.NumLevels())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthRequestFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 102)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("NumLevels", m.NumLevels())
	return w.Finish(), nil
}

func (m MarketDepthRequestFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 102)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("NumLevels", m.NumLevels())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthRequestFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 102 {
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
	m.SetNumLevels(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthRequestFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 102 {
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
	m.SetNumLevels(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthRequestFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 102 {
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
		case "NumLevels":
			m.SetNumLevels(r.Int32())
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

func (m *MarketDepthRequestFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 102 {
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
		case "NumLevels":
			m.SetNumLevels(r.Int32())
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

func (m *MarketDepthRequestBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetNumLevels(r.ReadInt32())
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

func (m *MarketDepthRequestPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetNumLevels(r.ReadInt32())
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

func (m MarketDepthRequestBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 102)
	w.WriteVarint32(1, int32(m.RequestAction()))
	w.WriteUvarint32(2, m.SymbolID())
	w.WriteString(3, m.Symbol())
	w.WriteString(4, m.Exchange())
	w.WriteVarint32(5, m.NumLevels())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthRequestPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 102)
	w.WriteVarint32(1, int32(m.RequestAction()))
	w.WriteUvarint32(2, m.SymbolID())
	w.WriteString(3, m.Symbol())
	w.WriteString(4, m.Exchange())
	w.WriteVarint32(5, m.NumLevels())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthRequestFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetNumLevels(r.ReadInt32())
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

func (m *MarketDepthRequestFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetNumLevels(r.ReadInt32())
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

func (m MarketDepthRequestFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 102)
	w.WriteVarint32(1, int32(m.RequestAction()))
	w.WriteUvarint32(2, m.SymbolID())
	w.WriteString(3, m.Symbol())
	w.WriteString(4, m.Exchange())
	w.WriteVarint32(5, m.NumLevels())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthRequestFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 102)
	w.WriteVarint32(1, int32(m.RequestAction()))
	w.WriteUvarint32(2, m.SymbolID())
	w.WriteString(3, m.Symbol())
	w.WriteString(4, m.Exchange())
	w.WriteVarint32(5, m.NumLevels())
	return w.Finish(), nil
}
