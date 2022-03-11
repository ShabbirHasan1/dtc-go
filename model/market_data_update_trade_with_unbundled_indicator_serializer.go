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

func (m MarketDataUpdateTradeWithUnbundledIndicator) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":137,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeCondition())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Reserve_1())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Reserve_2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Reserve_3())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":137,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeCondition())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Reserve_1())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Reserve_2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Reserve_3())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":137,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeCondition())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Reserve_1())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Reserve_2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Reserve_3())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":137,\"F\":["...)
	w.Uint32(m.SymbolID())
	w.Data = append(w.Data, ',')
	w.Uint8(uint8(m.AtBidOrAsk()))
	w.Data = append(w.Data, ',')
	w.Int8(int8(m.UnbundledTradeIndicator()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.TradeCondition())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Reserve_1())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Reserve_2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Volume())
	w.Data = append(w.Data, ',')
	w.Uint32(m.Reserve_3())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeWithUnbundledIndicator) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 137)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Uint8Field("TradeCondition", m.TradeCondition())
	w.Uint8Field("Reserve_1", m.Reserve_1())
	w.Uint32Field("Reserve_2", m.Reserve_2())
	w.Float64Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Uint32Field("Reserve_3", m.Reserve_3())
	w.Float64Field("DateTime", float64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 137)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Uint8Field("TradeCondition", m.TradeCondition())
	w.Uint8Field("Reserve_1", m.Reserve_1())
	w.Uint32Field("Reserve_2", m.Reserve_2())
	w.Float64Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Uint32Field("Reserve_3", m.Reserve_3())
	w.Float64Field("DateTime", float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeWithUnbundledIndicatorPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 137)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Uint8Field("TradeCondition", m.TradeCondition())
	w.Uint8Field("Reserve_1", m.Reserve_1())
	w.Uint32Field("Reserve_2", m.Reserve_2())
	w.Float64Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Uint32Field("Reserve_3", m.Reserve_3())
	w.Float64Field("DateTime", float64(m.DateTime()))
	return w.Finish(), nil
}

func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 137)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint8Field("AtBidOrAsk", uint8(m.AtBidOrAsk()))
	w.Int8Field("UnbundledTradeIndicator", int8(m.UnbundledTradeIndicator()))
	w.Uint8Field("TradeCondition", m.TradeCondition())
	w.Uint8Field("Reserve_1", m.Reserve_1())
	w.Uint32Field("Reserve_2", m.Reserve_2())
	w.Float64Field("Price", m.Price())
	w.Uint32Field("Volume", m.Volume())
	w.Uint32Field("Reserve_3", m.Reserve_3())
	w.Float64Field("DateTime", float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeWithUnbundledIndicatorBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 137 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
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
	m.SetTradeCondition(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetReserve_1(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetReserve_2(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetReserve_3(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 137 {
		return message.ErrWrongType
	}
	m.SetSymbolID(r.Uint32())
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
	m.SetTradeCondition(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetReserve_1(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetReserve_2(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetVolume(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetReserve_3(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeWithUnbundledIndicatorBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 137 {
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
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.Uint8()))
		case "UnbundledTradeIndicator":
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.Int8()))
		case "TradeCondition":
			m.SetTradeCondition(r.Uint8())
		case "Reserve_1":
			m.SetReserve_1(r.Uint8())
		case "Reserve_2":
			m.SetReserve_2(r.Uint32())
		case "Price":
			m.SetPrice(r.Float64())
		case "Volume":
			m.SetVolume(r.Uint32())
		case "Reserve_3":
			m.SetReserve_3(r.Uint32())
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
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

func (m *MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 137 {
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
		case "AtBidOrAsk":
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.Uint8()))
		case "UnbundledTradeIndicator":
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.Int8()))
		case "TradeCondition":
			m.SetTradeCondition(r.Uint8())
		case "Reserve_1":
			m.SetReserve_1(r.Uint8())
		case "Reserve_2":
			m.SetReserve_2(r.Uint32())
		case "Price":
			m.SetPrice(r.Float64())
		case "Volume":
			m.SetVolume(r.Uint32())
		case "Reserve_3":
			m.SetReserve_3(r.Uint32())
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
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

func (m *MarketDataUpdateTradeWithUnbundledIndicatorBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.ReadUint8()))
		case 3:
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.ReadInt8()))
		case 4:
			m.SetPrice(r.ReadFloat64())
		case 5:
			m.SetVolume(r.ReadUint32())
		case 6:
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 7:
			m.SetTradeCondition(r.ReadUint8())
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

func (m *MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetAtBidOrAsk(AtBidOrAskEnum8(r.ReadUint8()))
		case 3:
			m.SetUnbundledTradeIndicator(UnbundledTradeIndicatorEnum(r.ReadInt8()))
		case 4:
			m.SetPrice(r.ReadFloat64())
		case 5:
			m.SetVolume(r.ReadUint32())
		case 6:
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 7:
			m.SetTradeCondition(r.ReadUint8())
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

func (m MarketDataUpdateTradeWithUnbundledIndicatorBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 137)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint8(2, uint8(m.AtBidOrAsk()))
	w.WriteVarint8(3, int8(m.UnbundledTradeIndicator()))
	w.WriteFixed64Float64(4, m.Price())
	w.WriteUvarint32(5, m.Volume())
	w.WriteFixed64Float64(6, float64(m.DateTime()))
	w.WriteUvarint8(7, m.TradeCondition())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDataUpdateTradeWithUnbundledIndicatorPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 137)
	w.WriteUvarint32(1, m.SymbolID())
	w.WriteUvarint8(2, uint8(m.AtBidOrAsk()))
	w.WriteVarint8(3, int8(m.UnbundledTradeIndicator()))
	w.WriteFixed64Float64(4, m.Price())
	w.WriteUvarint32(5, m.Volume())
	w.WriteFixed64Float64(6, float64(m.DateTime()))
	w.WriteUvarint8(7, m.TradeCondition())
	return w.Finish(), nil
}
