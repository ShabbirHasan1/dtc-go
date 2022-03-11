//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersSnapshotMessageBoundary) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":155,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.MessageBoundary()))
	return w.Finish(), nil
}

func (m MarketOrdersSnapshotMessageBoundaryBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":155,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.MessageBoundary()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersSnapshotMessageBoundary) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 155)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint8Field("MessageBoundary", uint8(m.MessageBoundary()))
	return w.Finish(), nil
}

func (m MarketOrdersSnapshotMessageBoundaryBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 155)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint8Field("MessageBoundary", uint8(m.MessageBoundary()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersSnapshotMessageBoundaryBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 155 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageBoundary(MessageSetBoundaryEnum(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersSnapshotMessageBoundaryBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 155 {
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
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "MessageBoundary":
			m.SetMessageBoundary(MessageSetBoundaryEnum(r.Uint8()))
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
