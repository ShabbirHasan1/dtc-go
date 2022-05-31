// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketOrdersRemoveSize = 40

//     Size      uint16                       = MarketOrdersRemoveSize  (40)
//     Type      uint16                       = MARKET_ORDERS_REMOVE  (154)
//     SymbolID  uint32                       = 0
//     Quantity  uint32                       = 0
//     Price     float64                      = 0.000000
//     OrderID   uint64                       = 0
//     DateTime  DateTimeWithMicrosecondsInt  = 0
//     Side      BuySellEnum                  = BUY_SELL_UNSET  (0)
var _MarketOrdersRemoveDefault = []byte{40, 0, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketOrdersRemove struct {
	p message.Fixed
}

func NewMarketOrdersRemoveFrom(b []byte) MarketOrdersRemove {
	return MarketOrdersRemove{p: message.NewFixed(b)}
}

func WrapMarketOrdersRemove(b []byte) MarketOrdersRemove {
	return MarketOrdersRemove{p: message.WrapFixed(b)}
}

func NewMarketOrdersRemove() *MarketOrdersRemove {
	return &MarketOrdersRemove{p: message.NewFixed(_MarketOrdersRemoveDefault)}
}

func ParseMarketOrdersRemove(b []byte) (MarketOrdersRemove, error) {
	if len(b) < 4 {
		return MarketOrdersRemove{}, message.ErrShortBuffer
	}
	m := WrapMarketOrdersRemove(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketOrdersRemove{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketOrdersRemove{}, message.ErrBaseSizeOverflow
	}
	if size < 40 {
		clone := *NewMarketOrdersRemove()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketOrdersRemoveDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size
func (m MarketOrdersRemove) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m MarketOrdersRemove) Type() uint16 {
	return m.p.Uint16LE(2)
}

// SymbolID
func (m MarketOrdersRemove) SymbolID() uint32 {
	return m.p.Uint32LE(4)
}

// Quantity
func (m MarketOrdersRemove) Quantity() uint32 {
	return m.p.Uint32LE(8)
}

// Price
func (m MarketOrdersRemove) Price() float64 {
	return m.p.Float64LE(12)
}

// OrderID
func (m MarketOrdersRemove) OrderID() uint64 {
	return m.p.Uint64LE(20)
}

// DateTime
func (m MarketOrdersRemove) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(m.p.Int64LE(28))
}

// Side
func (m MarketOrdersRemove) Side() BuySellEnum {
	return BuySellEnum(m.p.Int32LE(36))
}

// SetSymbolID
func (m *MarketOrdersRemove) SetSymbolID(value uint32) *MarketOrdersRemove {
	m.p.SetUint32LE(4, value)
	return m
}

// SetQuantity
func (m *MarketOrdersRemove) SetQuantity(value uint32) *MarketOrdersRemove {
	m.p.SetUint32LE(8, value)
	return m
}

// SetPrice
func (m *MarketOrdersRemove) SetPrice(value float64) *MarketOrdersRemove {
	m.p.SetFloat64LE(12, value)
	return m
}

// SetOrderID
func (m *MarketOrdersRemove) SetOrderID(value uint64) *MarketOrdersRemove {
	m.p.SetUint64LE(20, value)
	return m
}

// SetDateTime
func (m *MarketOrdersRemove) SetDateTime(value DateTimeWithMicrosecondsInt) *MarketOrdersRemove {
	m.p.SetInt64LE(28, int64(value))
	return m
}

// SetSide
func (m *MarketOrdersRemove) SetSide(value BuySellEnum) *MarketOrdersRemove {
	m.p.SetInt32LE(36, int32(value))
	return m
}

func (m MarketOrdersRemove) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m MarketOrdersRemove) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m MarketOrdersRemove) Copy(to MarketOrdersRemove) {
	to.SetSymbolID(m.SymbolID())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
	to.SetSide(m.Side())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersRemove) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 154)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint32Field("Quantity", m.Quantity())
	w.Float64Field("Price", m.Price())
	w.Uint64Field("OrderID", m.OrderID())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Int32Field("Side", int32(m.Side()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersRemove) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 154 {
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
		case "Quantity":
			m.SetQuantity(r.Uint32())
		case "Price":
			m.SetPrice(r.Float64())
		case "OrderID":
			m.SetOrderID(r.Uint64())
		case "DateTime":
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "Side":
			m.SetSide(BuySellEnum(r.Int32()))
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

func (m *MarketOrdersRemove) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
