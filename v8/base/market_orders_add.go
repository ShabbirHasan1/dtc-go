// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketOrdersAddSize = 40

//     Size      uint16                       = MarketOrdersAddSize  (40)
//     Type      uint16                       = MARKET_ORDERS_ADD  (152)
//     SymbolID  uint32                       = 0
//     Side      BuySellEnum                  = BUY_SELL_UNSET  (0)
//     Quantity  uint32                       = 0
//     Price     float64                      = 0.000000
//     OrderID   uint64                       = 0
//     DateTime  DateTimeWithMicrosecondsInt  = 0
var _MarketOrdersAddDefault = []byte{40, 0, 152, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketOrdersAdd struct {
	p message.Fixed
}

func NewMarketOrdersAddFrom(b []byte) MarketOrdersAdd {
	return MarketOrdersAdd{p: message.NewFixed(b)}
}

func WrapMarketOrdersAdd(b []byte) MarketOrdersAdd {
	return MarketOrdersAdd{p: message.WrapFixed(b)}
}

func NewMarketOrdersAdd() *MarketOrdersAdd {
	return &MarketOrdersAdd{p: message.NewFixed(_MarketOrdersAddDefault)}
}

func ParseMarketOrdersAdd(b []byte) (MarketOrdersAdd, error) {
	if len(b) < 4 {
		return MarketOrdersAdd{}, message.ErrShortBuffer
	}
	m := WrapMarketOrdersAdd(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketOrdersAdd{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketOrdersAdd{}, message.ErrBaseSizeOverflow
	}
	if size < 40 {
		clone := *NewMarketOrdersAdd()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketOrdersAddDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size
func (m MarketOrdersAdd) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m MarketOrdersAdd) Type() uint16 {
	return m.p.Uint16LE(2)
}

// SymbolID
func (m MarketOrdersAdd) SymbolID() uint32 {
	return m.p.Uint32LE(4)
}

// Side
func (m MarketOrdersAdd) Side() BuySellEnum {
	return BuySellEnum(m.p.Int32LE(8))
}

// Quantity
func (m MarketOrdersAdd) Quantity() uint32 {
	return m.p.Uint32LE(12)
}

// Price
func (m MarketOrdersAdd) Price() float64 {
	return m.p.Float64LE(16)
}

// OrderID
func (m MarketOrdersAdd) OrderID() uint64 {
	return m.p.Uint64LE(24)
}

// DateTime
func (m MarketOrdersAdd) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(m.p.Int64LE(32))
}

// SetSymbolID
func (m *MarketOrdersAdd) SetSymbolID(value uint32) *MarketOrdersAdd {
	m.p.SetUint32LE(4, value)
	return m
}

// SetSide
func (m *MarketOrdersAdd) SetSide(value BuySellEnum) *MarketOrdersAdd {
	m.p.SetInt32LE(8, int32(value))
	return m
}

// SetQuantity
func (m *MarketOrdersAdd) SetQuantity(value uint32) *MarketOrdersAdd {
	m.p.SetUint32LE(12, value)
	return m
}

// SetPrice
func (m *MarketOrdersAdd) SetPrice(value float64) *MarketOrdersAdd {
	m.p.SetFloat64LE(16, value)
	return m
}

// SetOrderID
func (m *MarketOrdersAdd) SetOrderID(value uint64) *MarketOrdersAdd {
	m.p.SetUint64LE(24, value)
	return m
}

// SetDateTime
func (m *MarketOrdersAdd) SetDateTime(value DateTimeWithMicrosecondsInt) *MarketOrdersAdd {
	m.p.SetInt64LE(32, int64(value))
	return m
}

func (m *MarketOrdersAdd) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *MarketOrdersAdd) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *MarketOrdersAdd) Clone() *MarketOrdersAdd {
	return &MarketOrdersAdd{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m MarketOrdersAdd) Copy(to MarketOrdersAdd) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetQuantity(m.Quantity())
	to.SetPrice(m.Price())
	to.SetOrderID(m.OrderID())
	to.SetDateTime(m.DateTime())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersAdd) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *MarketOrdersAdd) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 152)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Int32Field("Side", int32(m.Side()))
	w.Uint32Field("Quantity", m.Quantity())
	w.Float64Field("Price", m.Price())
	w.Uint64Field("OrderID", m.OrderID())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersAdd) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *MarketOrdersAdd) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 152 {
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
		case "Side":
			m.SetSide(BuySellEnum(r.Int32()))
		case "Quantity":
			m.SetQuantity(r.Uint32())
		case "Price":
			m.SetPrice(r.Float64())
		case "OrderID":
			m.SetOrderID(r.Uint64())
		case "DateTime":
			m.SetDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
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
