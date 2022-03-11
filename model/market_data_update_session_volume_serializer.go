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

func (m MarketDataUpdateSessionVolume) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":113,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalSessionVolume())
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionVolumeBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":113,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalSessionVolume())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionVolumePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":113,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalSessionVolume())
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionVolumePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":113,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.TradingSessionDate()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalSessionVolume())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionVolume) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 113)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("Volume", m.Volume())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	w.BoolField("IsFinalSessionVolume", m.IsFinalSessionVolume())
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionVolumeBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 113)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("Volume", m.Volume())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	w.BoolField("IsFinalSessionVolume", m.IsFinalSessionVolume())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionVolumePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 113)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("Volume", m.Volume())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	w.BoolField("IsFinalSessionVolume", m.IsFinalSessionVolume())
	return w.Finish(), nil
}

func (m MarketDataUpdateSessionVolumePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 113)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("Volume", m.Volume())
	w.Uint32Field("TradingSessionDate", uint32(m.TradingSessionDate()))
	w.BoolField("IsFinalSessionVolume", m.IsFinalSessionVolume())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateSessionVolumeBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 113 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalSessionVolume(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateSessionVolumePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 113 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalSessionVolume(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateSessionVolumeBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 113 {
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
		case "Volume":
			m.SetVolume(r.Float64())
		case "TradingSessionDate":
			m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
		case "IsFinalSessionVolume":
			m.SetIsFinalSessionVolume(r.Bool())
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

func (m *MarketDataUpdateSessionVolumePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 113 {
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
		case "Volume":
			m.SetVolume(r.Float64())
		case "TradingSessionDate":
			m.SetTradingSessionDate(DateTime4Byte(r.Uint32()))
		case "IsFinalSessionVolume":
			m.SetIsFinalSessionVolume(r.Bool())
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

func (m *MarketDataUpdateSessionVolumeBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetVolume(r.ReadFloat64())
		case 3:
			m.SetTradingSessionDate(DateTime4Byte(r.ReadUint32()))
		case 4:
			m.SetIsFinalSessionVolume(r.ReadBool())
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

func (m *MarketDataUpdateSessionVolumePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetVolume(r.ReadFloat64())
		case 3:
			m.SetTradingSessionDate(DateTime4Byte(r.ReadUint32()))
		case 4:
			m.SetIsFinalSessionVolume(r.ReadBool())
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

func (m MarketDataUpdateSessionVolumeBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 113)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed64Float64(2, m.Volume())
	w.WriteUvarint32(3, uint32(m.TradingSessionDate()))
	w.WriteBool(4, m.IsFinalSessionVolume())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateSessionVolumePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 113)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed64Float64(2, m.Volume())
	w.WriteUvarint32(3, uint32(m.TradingSessionDate()))
	w.WriteBool(4, m.IsFinalSessionVolume())
	return w.Finish(), nil
}
