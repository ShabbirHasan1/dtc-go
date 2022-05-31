// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketDataUpdateTradeSize = 40

//     Size        uint16                    = MarketDataUpdateTradeSize  (40)
//     Type        uint16                    = MARKET_DATA_UPDATE_TRADE  (107)
//     SymbolID    uint32                    = 0
//     AtBidOrAsk  AtBidOrAskEnum            = BID_ASK_UNSET  (0)
//     Price       float64                   = 0.000000
//     Volume      float64                   = 0.000000
//     DateTime    DateTimeWithMilliseconds  = 0.000000
var _MarketDataUpdateTradeDefault = []byte{40, 0, 107, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataUpdateTrade The Server sends this market data feed message to the Client when a trade
// occurs.
type MarketDataUpdateTrade struct {
	p message.Fixed
}

func NewMarketDataUpdateTradeFrom(b []byte) MarketDataUpdateTrade {
	return MarketDataUpdateTrade{p: message.NewFixed(b)}
}

func WrapMarketDataUpdateTrade(b []byte) MarketDataUpdateTrade {
	return MarketDataUpdateTrade{p: message.WrapFixed(b)}
}

func NewMarketDataUpdateTrade() *MarketDataUpdateTrade {
	return &MarketDataUpdateTrade{p: message.NewFixed(_MarketDataUpdateTradeDefault)}
}

func ParseMarketDataUpdateTrade(b []byte) (MarketDataUpdateTrade, error) {
	if len(b) < 4 {
		return MarketDataUpdateTrade{}, message.ErrShortBuffer
	}
	m := WrapMarketDataUpdateTrade(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketDataUpdateTrade{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketDataUpdateTrade{}, message.ErrBaseSizeOverflow
	}
	if size < 40 {
		clone := *NewMarketDataUpdateTrade()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketDataUpdateTradeDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m MarketDataUpdateTrade) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m MarketDataUpdateTrade) Type() uint16 {
	return m.p.Uint16LE(2)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateTrade) SymbolID() uint32 {
	return m.p.Uint32LE(4)
}

// AtBidOrAsk Indicator whether the trade occurred at the bid or ask.
func (m MarketDataUpdateTrade) AtBidOrAsk() AtBidOrAskEnum {
	return AtBidOrAskEnum(m.p.Uint16LE(8))
}

// Price The price of the trade.
func (m MarketDataUpdateTrade) Price() float64 {
	return m.p.Float64LE(16)
}

// Volume The volume of the trade.
func (m MarketDataUpdateTrade) Volume() float64 {
	return m.p.Float64LE(24)
}

// DateTime The Date-Time of the trade.
func (m MarketDataUpdateTrade) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(m.p.Float64LE(32))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m *MarketDataUpdateTrade) SetSymbolID(value uint32) *MarketDataUpdateTrade {
	m.p.SetUint32LE(4, value)
	return m
}

// SetAtBidOrAsk Indicator whether the trade occurred at the bid or ask.
func (m *MarketDataUpdateTrade) SetAtBidOrAsk(value AtBidOrAskEnum) *MarketDataUpdateTrade {
	m.p.SetUint16LE(8, uint16(value))
	return m
}

// SetPrice The price of the trade.
func (m *MarketDataUpdateTrade) SetPrice(value float64) *MarketDataUpdateTrade {
	m.p.SetFloat64LE(16, value)
	return m
}

// SetVolume The volume of the trade.
func (m *MarketDataUpdateTrade) SetVolume(value float64) *MarketDataUpdateTrade {
	m.p.SetFloat64LE(24, value)
	return m
}

// SetDateTime The Date-Time of the trade.
func (m *MarketDataUpdateTrade) SetDateTime(value DateTimeWithMilliseconds) *MarketDataUpdateTrade {
	m.p.SetFloat64LE(32, float64(value))
	return m
}

func (m *MarketDataUpdateTrade) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *MarketDataUpdateTrade) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *MarketDataUpdateTrade) Clone() *MarketDataUpdateTrade {
	return &MarketDataUpdateTrade{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m MarketDataUpdateTrade) Copy(to MarketDataUpdateTrade) {
	to.SetSymbolID(m.SymbolID())
	to.SetAtBidOrAsk(m.AtBidOrAsk())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
	to.SetDateTime(m.DateTime())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTrade) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *MarketDataUpdateTrade) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 107)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("AtBidOrAsk", uint16(m.AtBidOrAsk()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Volume", m.Volume())
	w.Float64Field("DateTime", float64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateTrade) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *MarketDataUpdateTrade) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 107 {
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
			m.SetAtBidOrAsk(AtBidOrAskEnum(r.Uint16()))
		case "Price":
			m.SetPrice(r.Float64())
		case "Volume":
			m.SetVolume(r.Float64())
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
