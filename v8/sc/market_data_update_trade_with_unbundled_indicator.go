// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketDataUpdateTradeWithUnbundledIndicatorSize = 40

//     Size                     uint16                       = MarketDataUpdateTradeWithUnbundledIndicatorSize  (40)
//     Type                     uint16                       = MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR  (137)
//     SymbolID                 uint32                       = 0
//     AtBidOrAsk               AtBidOrAskEnum8              = BID_ASK_UNSET_8  (0)
//     UnbundledTradeIndicator  UnbundledTradeIndicatorEnum  = UNBUNDLED_TRADE_NONE  (0)
//     TradeCondition           uint8                        = 0
//     Reserve_1                uint8                        = 0
//     Reserve_2                uint32                       = 0
//     Price                    float64                      = 0.000000
//     Volume                   uint32                       = 0
//     Reserve_3                uint32                       = 0
//     DateTime                 DateTimeWithMilliseconds     = 0.000000
var _MarketDataUpdateTradeWithUnbundledIndicatorDefault = []byte{40, 0, 137, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketDataUpdateTradeWithUnbundledIndicator struct {
	p message.Fixed
}

func NewMarketDataUpdateTradeWithUnbundledIndicatorFrom(b []byte) MarketDataUpdateTradeWithUnbundledIndicator {
	return MarketDataUpdateTradeWithUnbundledIndicator{p: message.NewFixed(b)}
}

func WrapMarketDataUpdateTradeWithUnbundledIndicator(b []byte) MarketDataUpdateTradeWithUnbundledIndicator {
	return MarketDataUpdateTradeWithUnbundledIndicator{p: message.WrapFixed(b)}
}

func NewMarketDataUpdateTradeWithUnbundledIndicator() *MarketDataUpdateTradeWithUnbundledIndicator {
	return &MarketDataUpdateTradeWithUnbundledIndicator{p: message.NewFixed(_MarketDataUpdateTradeWithUnbundledIndicatorDefault)}
}

func ParseMarketDataUpdateTradeWithUnbundledIndicator(b []byte) (MarketDataUpdateTradeWithUnbundledIndicator, error) {
	if len(b) < 4 {
		return MarketDataUpdateTradeWithUnbundledIndicator{}, message.ErrShortBuffer
	}
	m := WrapMarketDataUpdateTradeWithUnbundledIndicator(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketDataUpdateTradeWithUnbundledIndicator{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketDataUpdateTradeWithUnbundledIndicator{}, message.ErrBaseSizeOverflow
	}
	if size < 40 {
		clone := *NewMarketDataUpdateTradeWithUnbundledIndicator()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketDataUpdateTradeWithUnbundledIndicatorDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size
func (m MarketDataUpdateTradeWithUnbundledIndicator) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m MarketDataUpdateTradeWithUnbundledIndicator) Type() uint16 {
	return m.p.Uint16LE(2)
}

// SymbolID
func (m MarketDataUpdateTradeWithUnbundledIndicator) SymbolID() uint32 {
	return m.p.Uint32LE(4)
}

// AtBidOrAsk
func (m MarketDataUpdateTradeWithUnbundledIndicator) AtBidOrAsk() AtBidOrAskEnum8 {
	return AtBidOrAskEnum8(m.p.Uint8(8))
}

// UnbundledTradeIndicator
func (m MarketDataUpdateTradeWithUnbundledIndicator) UnbundledTradeIndicator() UnbundledTradeIndicatorEnum {
	return UnbundledTradeIndicatorEnum(m.p.Int8(9))
}

// TradeCondition
func (m MarketDataUpdateTradeWithUnbundledIndicator) TradeCondition() uint8 {
	return m.p.Uint8(10)
}

// Reserve_1
func (m MarketDataUpdateTradeWithUnbundledIndicator) Reserve_1() uint8 {
	return m.p.Uint8(11)
}

// Reserve_2
func (m MarketDataUpdateTradeWithUnbundledIndicator) Reserve_2() uint32 {
	return m.p.Uint32LE(12)
}

// Price
func (m MarketDataUpdateTradeWithUnbundledIndicator) Price() float64 {
	return m.p.Float64LE(16)
}

// Volume
func (m MarketDataUpdateTradeWithUnbundledIndicator) Volume() uint32 {
	return m.p.Uint32LE(24)
}

// Reserve_3
func (m MarketDataUpdateTradeWithUnbundledIndicator) Reserve_3() uint32 {
	return m.p.Uint32LE(28)
}

// DateTime
func (m MarketDataUpdateTradeWithUnbundledIndicator) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(m.p.Float64LE(32))
}

// SetSymbolID
func (m *MarketDataUpdateTradeWithUnbundledIndicator) SetSymbolID(value uint32) *MarketDataUpdateTradeWithUnbundledIndicator {
	m.p.SetUint32LE(4, value)
	return m
}

// SetAtBidOrAsk
func (m *MarketDataUpdateTradeWithUnbundledIndicator) SetAtBidOrAsk(value AtBidOrAskEnum8) *MarketDataUpdateTradeWithUnbundledIndicator {
	m.p.SetUint8(8, uint8(value))
	return m
}

// SetUnbundledTradeIndicator
func (m *MarketDataUpdateTradeWithUnbundledIndicator) SetUnbundledTradeIndicator(value UnbundledTradeIndicatorEnum) *MarketDataUpdateTradeWithUnbundledIndicator {
	m.p.SetInt8(9, int8(value))
	return m
}

// SetTradeCondition
func (m *MarketDataUpdateTradeWithUnbundledIndicator) SetTradeCondition(value uint8) *MarketDataUpdateTradeWithUnbundledIndicator {
	m.p.SetUint8(10, value)
	return m
}

// SetReserve_1
func (m *MarketDataUpdateTradeWithUnbundledIndicator) SetReserve_1(value uint8) *MarketDataUpdateTradeWithUnbundledIndicator {
	m.p.SetUint8(11, value)
	return m
}

// SetReserve_2
func (m *MarketDataUpdateTradeWithUnbundledIndicator) SetReserve_2(value uint32) *MarketDataUpdateTradeWithUnbundledIndicator {
	m.p.SetUint32LE(12, value)
	return m
}

// SetPrice
func (m *MarketDataUpdateTradeWithUnbundledIndicator) SetPrice(value float64) *MarketDataUpdateTradeWithUnbundledIndicator {
	m.p.SetFloat64LE(16, value)
	return m
}

// SetVolume
func (m *MarketDataUpdateTradeWithUnbundledIndicator) SetVolume(value uint32) *MarketDataUpdateTradeWithUnbundledIndicator {
	m.p.SetUint32LE(24, value)
	return m
}

// SetReserve_3
func (m *MarketDataUpdateTradeWithUnbundledIndicator) SetReserve_3(value uint32) *MarketDataUpdateTradeWithUnbundledIndicator {
	m.p.SetUint32LE(28, value)
	return m
}

// SetDateTime
func (m *MarketDataUpdateTradeWithUnbundledIndicator) SetDateTime(value DateTimeWithMilliseconds) *MarketDataUpdateTradeWithUnbundledIndicator {
	m.p.SetFloat64LE(32, float64(value))
	return m
}

func (m MarketDataUpdateTradeWithUnbundledIndicator) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m MarketDataUpdateTradeWithUnbundledIndicator) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m MarketDataUpdateTradeWithUnbundledIndicator) Copy(to MarketDataUpdateTradeWithUnbundledIndicator) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetUnbundledTradeIndicator(m.UnbundledTradeIndicator())
	to.SetTradeCondition(m.TradeCondition())
	to.SetReserve_1(m.Reserve_1())
	to.SetReserve_2(m.Reserve_2())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetReserve_3(m.Reserve_3())
	to.SetDateTime(m.DateTime())
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

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTradeWithUnbundledIndicator) UnmarshalJSONDoc(r *json.Reader) error {
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

func (m *MarketDataUpdateTradeWithUnbundledIndicator) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
