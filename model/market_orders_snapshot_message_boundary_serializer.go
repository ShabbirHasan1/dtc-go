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
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersSnapshotMessageBoundaryPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":155,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.MessageBoundary()))
	return w.Finish(), nil
}

func (m MarketOrdersSnapshotMessageBoundaryPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersSnapshotMessageBoundaryPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 155)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint8Field("MessageBoundary", uint8(m.MessageBoundary()))
	return w.Finish(), nil
}

func (m MarketOrdersSnapshotMessageBoundaryPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
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
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersSnapshotMessageBoundaryPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersSnapshotMessageBoundaryPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersSnapshotMessageBoundaryBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbolID(r.ReadUint32())
		case 2:
			m.SetMessageBoundary(MessageSetBoundaryEnum(r.ReadUint8()))
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

func (m *MarketOrdersSnapshotMessageBoundaryPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbolID(r.ReadUint32())
		case 2:
			m.SetMessageBoundary(MessageSetBoundaryEnum(r.ReadUint8()))
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

func (m MarketOrdersSnapshotMessageBoundaryBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 155)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint8(2, uint8(m.MessageBoundary()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersSnapshotMessageBoundaryPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 155)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint8(2, uint8(m.MessageBoundary()))
	return w.Finish(), nil
}
