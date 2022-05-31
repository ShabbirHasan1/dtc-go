// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketOrdersRejectSize = 16

const MarketOrdersRejectFixedSize = 104

//     Size        uint16  = MarketOrdersRejectSize  (16)
//     Type        uint16  = MARKET_ORDERS_REJECT  (151)
//     BaseSize    uint16  = MarketOrdersRejectSize  (16)
//     SymbolID    uint32  = 0
//     RejectText  string  = ""
var _MarketOrdersRejectDefault = []byte{16, 0, 151, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size        uint16      = MarketOrdersRejectFixedSize  (104)
//     Type        uint16      = MARKET_ORDERS_REJECT  (151)
//     SymbolID    uint32      = 0
//     RejectText  string[96]  = ""
var _MarketOrdersRejectFixedDefault = []byte{104, 0, 151, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketOrdersReject struct {
	p message.VLS
}

type MarketOrdersRejectFixed struct {
	p message.Fixed
}

func NewMarketOrdersRejectFrom(b []byte) MarketOrdersReject {
	return MarketOrdersReject{p: message.NewVLS(b)}
}

func WrapMarketOrdersReject(b []byte) MarketOrdersReject {
	return MarketOrdersReject{p: message.WrapVLS(b)}
}

func NewMarketOrdersReject() *MarketOrdersReject {
	return &MarketOrdersReject{p: message.NewVLS(_MarketOrdersRejectDefault)}
}

func ParseMarketOrdersReject(b []byte) (MarketOrdersReject, error) {
	if len(b) < 6 {
		return MarketOrdersReject{}, message.ErrShortBuffer
	}
	m := WrapMarketOrdersReject(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketOrdersReject{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return MarketOrdersReject{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 16 {
		newSize := len(b) + (16 - baseSize)
		if newSize > message.MaxSize {
			return MarketOrdersReject{}, message.ErrOverflow
		}
		clone := MarketOrdersReject{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _MarketOrdersRejectDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(16 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(12)
			if offset > 0 {
				clone.p.SetUint16LE(12, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewMarketOrdersRejectFixedFrom(b []byte) MarketOrdersRejectFixed {
	return MarketOrdersRejectFixed{p: message.NewFixed(b)}
}

func WrapMarketOrdersRejectFixed(b []byte) MarketOrdersRejectFixed {
	return MarketOrdersRejectFixed{p: message.WrapFixed(b)}
}

func NewMarketOrdersRejectFixed() *MarketOrdersRejectFixed {
	return &MarketOrdersRejectFixed{p: message.NewFixed(_MarketOrdersRejectFixedDefault)}
}

func ParseMarketOrdersRejectFixed(b []byte) (MarketOrdersRejectFixed, error) {
	if len(b) < 4 {
		return MarketOrdersRejectFixed{}, message.ErrShortBuffer
	}
	m := WrapMarketOrdersRejectFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketOrdersRejectFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketOrdersRejectFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 104 {
		clone := *NewMarketOrdersRejectFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketOrdersRejectFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size
func (m MarketOrdersReject) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m MarketOrdersReject) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m MarketOrdersReject) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// SymbolID
func (m MarketOrdersReject) SymbolID() uint32 {
	return m.p.Uint32LE(8)
}

// RejectText
func (m MarketOrdersReject) RejectText() string {
	return m.p.StringVLS(12)
}

// Size
func (m MarketOrdersRejectFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m MarketOrdersRejectFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// SymbolID
func (m MarketOrdersRejectFixed) SymbolID() uint32 {
	return m.p.Uint32LE(4)
}

// RejectText
func (m MarketOrdersRejectFixed) RejectText() string {
	return m.p.StringFixed(8, 96)
}

// SetSymbolID
func (m *MarketOrdersReject) SetSymbolID(value uint32) *MarketOrdersReject {
	m.p.SetUint32LE(8, value)
	return m
}

// SetRejectText
func (m *MarketOrdersReject) SetRejectText(value string) *MarketOrdersReject {
	m.p.SetStringVLS(12, value)
	return m
}

// SetSymbolID
func (m *MarketOrdersRejectFixed) SetSymbolID(value uint32) *MarketOrdersRejectFixed {
	m.p.SetUint32LE(4, value)
	return m
}

// SetRejectText
func (m *MarketOrdersRejectFixed) SetRejectText(value string) *MarketOrdersRejectFixed {
	m.p.SetStringFixed(8, 96, value)
	return m
}

func (m MarketOrdersReject) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m MarketOrdersReject) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m MarketOrdersRejectFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m MarketOrdersRejectFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m MarketOrdersReject) Copy(to MarketOrdersReject) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersReject) CopyTo(to MarketOrdersRejectFixed) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m MarketOrdersRejectFixed) Copy(to MarketOrdersRejectFixed) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m MarketOrdersRejectFixed) CopyTo(to MarketOrdersReject) {
	to.SetSymbolID(m.SymbolID())
	to.SetRejectText(m.RejectText())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersReject) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 151)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersReject) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 151 {
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
		case "RejectText":
			m.SetRejectText(r.String())
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

func (m *MarketOrdersReject) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersRejectFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 151)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersRejectFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 151 {
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
		case "RejectText":
			m.SetRejectText(r.String())
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

func (m *MarketOrdersRejectFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}