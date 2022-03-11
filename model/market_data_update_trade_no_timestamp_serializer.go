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

func (m MarketDataUpdateTradeNoTimestamp) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":142,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradeCondition()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeNoTimestampBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":142,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradeCondition()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeNoTimestampPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":142,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradeCondition()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeNoTimestampPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":142,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Float32(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.TradeCondition()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeNoTimestamp) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 142)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Int8Field("TradeCondition", int8(m.TradeCondition()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeNoTimestampBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 142)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Int8Field("TradeCondition", int8(m.TradeCondition()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeNoTimestampPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 142)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Int8Field("TradeCondition", int8(m.TradeCondition()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeNoTimestampPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 142)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float32Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Int8Field("TradeCondition", int8(m.TradeCondition()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeNoTimestampBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 142 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum8(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeCondition(TradeConditionEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeNoTimestampPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 142 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetAtBidOrAsk(AtBidOrAskEnum8(r.Uint8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeCondition(TradeConditionEnum(r.Int8()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeNoTimestampBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 142 {
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
		case "Price":
			m.SetPrice(r.Float32())
		case "Volume":
			m.SetVolume(r.Uint32())
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.Uint8()))
		case "UnbundledTradeIndicator":
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.Int8()))
		case "TradeCondition":
			m.SetTradeCondition(TradeConditionEnum(r.Int8()))
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

func (m *MarketDataUpdateTradeNoTimestampPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 142 {
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
		case "Price":
			m.SetPrice(r.Float32())
		case "Volume":
			m.SetVolume(r.Uint32())
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.Uint8()))
		case "UnbundledTradeIndicator":
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.Int8()))
		case "TradeCondition":
			m.SetTradeCondition(TradeConditionEnum(r.Int8()))
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

func (m *MarketDataUpdateTradeNoTimestampBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat32())
		case 3:
			m.SetVolume(r.ReadUint32())
		case 4:
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.ReadUint8()))
		case 5:
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.ReadInt8()))
		case 6:
			m.SetTradeCondition(TradeConditionEnum(r.ReadInt8()))
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

func (m *MarketDataUpdateTradeNoTimestampPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice(r.ReadFloat32())
		case 3:
			m.SetVolume(r.ReadUint32())
		case 4:
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.ReadUint8()))
		case 5:
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.ReadInt8()))
		case 6:
			m.SetTradeCondition(TradeConditionEnum(r.ReadInt8()))
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

func (m MarketDataUpdateTradeNoTimestampBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 142)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.Price())
	w.WriteUvarint32(3, m.Volume())
	w.WriteUvarint8(4, uint8(m.AtBidOrAsk()))
	w.WriteVarint8(5, int8(m.UnbundledTradeIndicator()))
	w.WriteVarint8(6, int8(m.TradeCondition()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeNoTimestampPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 142)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteFixed32Float32(2, m.Price())
	w.WriteUvarint32(3, m.Volume())
	w.WriteUvarint8(4, uint8(m.AtBidOrAsk()))
	w.WriteVarint8(5, int8(m.UnbundledTradeIndicator()))
	w.WriteVarint8(6, int8(m.TradeCondition()))
	return w.Finish(), nil
}
