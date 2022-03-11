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

func (m MarketDataRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":101,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

func (m MarketDataRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":101,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataRequestPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":101,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

func (m MarketDataRequestPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":101,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 101)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("IntervalForSnapshotUpdatesInMilliseconds", m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

func (m MarketDataRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 101)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("IntervalForSnapshotUpdatesInMilliseconds", m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataRequestPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 101)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("IntervalForSnapshotUpdatesInMilliseconds", m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

func (m MarketDataRequestPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 101)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("IntervalForSnapshotUpdatesInMilliseconds", m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 101 {
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
	m.SetIntervalForSnapshotUpdatesInMilliseconds(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataRequestPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 101 {
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
	m.SetIntervalForSnapshotUpdatesInMilliseconds(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 101 {
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
		case "IntervalForSnapshotUpdatesInMilliseconds":
			m.SetIntervalForSnapshotUpdatesInMilliseconds(r.Uint32())
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

func (m *MarketDataRequestPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 101 {
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
		case "IntervalForSnapshotUpdatesInMilliseconds":
			m.SetIntervalForSnapshotUpdatesInMilliseconds(r.Uint32())
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

func (m MarketDataRequestFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":101,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

func (m MarketDataRequestFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":101,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataRequestFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":101,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

func (m MarketDataRequestFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":101,\"F\":["...)
	w.Int32(int32(m.RequestAction()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Uint32(m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 101)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("IntervalForSnapshotUpdatesInMilliseconds", m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

func (m MarketDataRequestFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 101)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("IntervalForSnapshotUpdatesInMilliseconds", m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataRequestFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 101)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("IntervalForSnapshotUpdatesInMilliseconds", m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

func (m MarketDataRequestFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 101)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("IntervalForSnapshotUpdatesInMilliseconds", m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataRequestFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 101 {
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
	m.SetIntervalForSnapshotUpdatesInMilliseconds(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataRequestFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 101 {
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
	m.SetIntervalForSnapshotUpdatesInMilliseconds(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataRequestFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 101 {
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
		case "IntervalForSnapshotUpdatesInMilliseconds":
			m.SetIntervalForSnapshotUpdatesInMilliseconds(r.Uint32())
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

func (m *MarketDataRequestFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 101 {
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
		case "IntervalForSnapshotUpdatesInMilliseconds":
			m.SetIntervalForSnapshotUpdatesInMilliseconds(r.Uint32())
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

func (m *MarketDataRequestBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetIntervalForSnapshotUpdatesInMilliseconds(r.ReadUint32())
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

func (m *MarketDataRequestPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetIntervalForSnapshotUpdatesInMilliseconds(r.ReadUint32())
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

func (m MarketDataRequestBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 101)
	w.WriteVarint32(1, int32(m.RequestAction()))
	w.WriteUvarint32(2, m.SymbolID())
	w.WriteString(3, m.Symbol())
	w.WriteString(4, m.Exchange())
	w.WriteUvarint32(5, m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataRequestPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 101)
	w.WriteVarint32(1, int32(m.RequestAction()))
	w.WriteUvarint32(2, m.SymbolID())
	w.WriteString(3, m.Symbol())
	w.WriteString(4, m.Exchange())
	w.WriteUvarint32(5, m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataRequestFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetIntervalForSnapshotUpdatesInMilliseconds(r.ReadUint32())
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

func (m *MarketDataRequestFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetIntervalForSnapshotUpdatesInMilliseconds(r.ReadUint32())
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

func (m MarketDataRequestFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 101)
	w.WriteVarint32(1, int32(m.RequestAction()))
	w.WriteUvarint32(2, m.SymbolID())
	w.WriteString(3, m.Symbol())
	w.WriteString(4, m.Exchange())
	w.WriteUvarint32(5, m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataRequestFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 101)
	w.WriteVarint32(1, int32(m.RequestAction()))
	w.WriteUvarint32(2, m.SymbolID())
	w.WriteString(3, m.Symbol())
	w.WriteString(4, m.Exchange())
	w.WriteUvarint32(5, m.IntervalForSnapshotUpdatesInMilliseconds())
	return w.Finish(), nil
}
