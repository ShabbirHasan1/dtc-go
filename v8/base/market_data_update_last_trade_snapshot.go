// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketDataUpdateLastTradeSnapshotSize = 32

//     Size               uint16                    = MarketDataUpdateLastTradeSnapshotSize  (32)
//     Type               uint16                    = MARKET_DATA_UPDATE_LAST_TRADE_SNAPSHOT  (134)
//     SymbolID           uint32                    = 0
//     LastTradePrice     float64                   = 0.000000
//     LastTradeVolume    float64                   = 0.000000
//     LastTradeDateTime  DateTimeWithMilliseconds  = 0.000000
var _MarketDataUpdateLastTradeSnapshotDefault = []byte{32, 0, 134, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDataUpdateLastTradeSnapshot Sent by the Server to the Client to update the last trade price, volume
// and date-time fields under conditions when there is not a trade.
//
// This message type does not signify a trade has occurred. It should never
// be interpreted by the Client in that way.
type MarketDataUpdateLastTradeSnapshot struct {
	p message.Fixed
}

func NewMarketDataUpdateLastTradeSnapshotFrom(b []byte) MarketDataUpdateLastTradeSnapshot {
	return MarketDataUpdateLastTradeSnapshot{p: message.NewFixed(b)}
}

func WrapMarketDataUpdateLastTradeSnapshot(b []byte) MarketDataUpdateLastTradeSnapshot {
	return MarketDataUpdateLastTradeSnapshot{p: message.WrapFixed(b)}
}

func NewMarketDataUpdateLastTradeSnapshot() *MarketDataUpdateLastTradeSnapshot {
	return &MarketDataUpdateLastTradeSnapshot{p: message.NewFixed(_MarketDataUpdateLastTradeSnapshotDefault)}
}

func ParseMarketDataUpdateLastTradeSnapshot(b []byte) (MarketDataUpdateLastTradeSnapshot, error) {
	if len(b) < 4 {
		return MarketDataUpdateLastTradeSnapshot{}, message.ErrShortBuffer
	}
	m := WrapMarketDataUpdateLastTradeSnapshot(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketDataUpdateLastTradeSnapshot{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketDataUpdateLastTradeSnapshot{}, message.ErrBaseSizeOverflow
	}
	if size < 32 {
		clone := *NewMarketDataUpdateLastTradeSnapshot()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketDataUpdateLastTradeSnapshotDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m MarketDataUpdateLastTradeSnapshot) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m MarketDataUpdateLastTradeSnapshot) Type() uint16 {
	return m.p.Uint16LE(2)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDataUpdateLastTradeSnapshot) SymbolID() uint32 {
	return m.p.Uint32LE(4)
}

// LastTradePrice The most recent last trade price.
func (m MarketDataUpdateLastTradeSnapshot) LastTradePrice() float64 {
	return m.p.Float64LE(8)
}

// LastTradeVolume The quantity/volume of the most recent last trade.
func (m MarketDataUpdateLastTradeSnapshot) LastTradeVolume() float64 {
	return m.p.Float64LE(16)
}

// LastTradeDateTime The Date-Time of the last trade.
func (m MarketDataUpdateLastTradeSnapshot) LastTradeDateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(m.p.Float64LE(24))
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDataRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m *MarketDataUpdateLastTradeSnapshot) SetSymbolID(value uint32) *MarketDataUpdateLastTradeSnapshot {
	m.p.SetUint32LE(4, value)
	return m
}

// SetLastTradePrice The most recent last trade price.
func (m *MarketDataUpdateLastTradeSnapshot) SetLastTradePrice(value float64) *MarketDataUpdateLastTradeSnapshot {
	m.p.SetFloat64LE(8, value)
	return m
}

// SetLastTradeVolume The quantity/volume of the most recent last trade.
func (m *MarketDataUpdateLastTradeSnapshot) SetLastTradeVolume(value float64) *MarketDataUpdateLastTradeSnapshot {
	m.p.SetFloat64LE(16, value)
	return m
}

// SetLastTradeDateTime The Date-Time of the last trade.
func (m *MarketDataUpdateLastTradeSnapshot) SetLastTradeDateTime(value DateTimeWithMilliseconds) *MarketDataUpdateLastTradeSnapshot {
	m.p.SetFloat64LE(24, float64(value))
	return m
}

func (m *MarketDataUpdateLastTradeSnapshot) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *MarketDataUpdateLastTradeSnapshot) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *MarketDataUpdateLastTradeSnapshot) Clone() *MarketDataUpdateLastTradeSnapshot {
	return &MarketDataUpdateLastTradeSnapshot{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m MarketDataUpdateLastTradeSnapshot) Copy(to MarketDataUpdateLastTradeSnapshot) {
	to.SetSymbolID(m.SymbolID())
	to.SetLastTradePrice(m.LastTradePrice())
	to.SetLastTradeVolume(m.LastTradeVolume())
	to.SetLastTradeDateTime(m.LastTradeDateTime())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateLastTradeSnapshot) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *MarketDataUpdateLastTradeSnapshot) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 134)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Float64Field("LastTradePrice", m.LastTradePrice())
	w.Float64Field("LastTradeVolume", m.LastTradeVolume())
	w.Float64Field("LastTradeDateTime", float64(m.LastTradeDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDataUpdateLastTradeSnapshot) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *MarketDataUpdateLastTradeSnapshot) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 134 {
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
		case "LastTradePrice":
			m.SetLastTradePrice(r.Float64())
		case "LastTradeVolume":
			m.SetLastTradeVolume(r.Float64())
		case "LastTradeDateTime":
			m.SetLastTradeDateTime(DateTimeWithMilliseconds(r.Float64()))
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
