// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketOrdersRequestSize = 28

const MarketOrdersRequestFixedSize = 96

//     Size                            uint16             = MarketOrdersRequestSize  (28)
//     Type                            uint16             = MARKET_ORDERS_REQUEST  (150)
//     BaseSize                        uint16             = MarketOrdersRequestSize  (28)
//     RequestAction                   RequestActionEnum  = SUBSCRIBE  (0)
//     SymbolID                        uint32             = 0
//     Symbol                          string             = ""
//     Exchange                        string             = ""
//     SendQuantitiesGreaterOrEqualTo  uint32             = 0
var _MarketOrdersRequestDefault = []byte{28, 0, 150, 0, 28, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size                            uint16             = MarketOrdersRequestFixedSize  (96)
//     Type                            uint16             = MARKET_ORDERS_REQUEST  (150)
//     RequestAction                   RequestActionEnum  = SUBSCRIBE  (1)
//     SymbolID                        uint32             = 0
//     Symbol                          string[64]         = ""
//     Exchange                        string[16]         = ""
//     SendQuantitiesGreaterOrEqualTo  uint32             = 0
var _MarketOrdersRequestFixedDefault = []byte{96, 0, 150, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketOrdersRequest struct {
	p message.VLS
}

type MarketOrdersRequestFixed struct {
	p message.Fixed
}

func NewMarketOrdersRequestFrom(b []byte) MarketOrdersRequest {
	return MarketOrdersRequest{p: message.NewVLS(b)}
}

func WrapMarketOrdersRequest(b []byte) MarketOrdersRequest {
	return MarketOrdersRequest{p: message.WrapVLS(b)}
}

func NewMarketOrdersRequest() *MarketOrdersRequest {
	return &MarketOrdersRequest{p: message.NewVLS(_MarketOrdersRequestDefault)}
}

func ParseMarketOrdersRequest(b []byte) (MarketOrdersRequest, error) {
	if len(b) < 6 {
		return MarketOrdersRequest{}, message.ErrShortBuffer
	}
	m := WrapMarketOrdersRequest(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketOrdersRequest{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return MarketOrdersRequest{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 28 {
		newSize := len(b) + (28 - baseSize)
		if newSize > message.MaxSize {
			return MarketOrdersRequest{}, message.ErrOverflow
		}
		clone := MarketOrdersRequest{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _MarketOrdersRequestDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(28 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(16)
			if offset > 0 {
				clone.p.SetUint16LE(16, offset+shift)
			}
			offset = clone.p.Uint16LE(20)
			if offset > 0 {
				clone.p.SetUint16LE(20, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewMarketOrdersRequestFixedFrom(b []byte) MarketOrdersRequestFixed {
	return MarketOrdersRequestFixed{p: message.NewFixed(b)}
}

func WrapMarketOrdersRequestFixed(b []byte) MarketOrdersRequestFixed {
	return MarketOrdersRequestFixed{p: message.WrapFixed(b)}
}

func NewMarketOrdersRequestFixed() *MarketOrdersRequestFixed {
	return &MarketOrdersRequestFixed{p: message.NewFixed(_MarketOrdersRequestFixedDefault)}
}

func ParseMarketOrdersRequestFixed(b []byte) (MarketOrdersRequestFixed, error) {
	if len(b) < 4 {
		return MarketOrdersRequestFixed{}, message.ErrShortBuffer
	}
	m := WrapMarketOrdersRequestFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketOrdersRequestFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketOrdersRequestFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 96 {
		clone := *NewMarketOrdersRequestFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketOrdersRequestFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size
func (m MarketOrdersRequest) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m MarketOrdersRequest) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m MarketOrdersRequest) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestAction
func (m MarketOrdersRequest) RequestAction() RequestActionEnum {
	return RequestActionEnum(m.p.Int32LE(8))
}

// SymbolID
func (m MarketOrdersRequest) SymbolID() uint32 {
	return m.p.Uint32LE(12)
}

// Symbol
func (m MarketOrdersRequest) Symbol() string {
	return m.p.StringVLS(16)
}

// Exchange
func (m MarketOrdersRequest) Exchange() string {
	return m.p.StringVLS(20)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequest) SendQuantitiesGreaterOrEqualTo() uint32 {
	return m.p.Uint32LE(24)
}

// Size
func (m MarketOrdersRequestFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m MarketOrdersRequestFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestAction
func (m MarketOrdersRequestFixed) RequestAction() RequestActionEnum {
	return RequestActionEnum(m.p.Int32LE(4))
}

// SymbolID
func (m MarketOrdersRequestFixed) SymbolID() uint32 {
	return m.p.Uint32LE(8)
}

// Symbol
func (m MarketOrdersRequestFixed) Symbol() string {
	return m.p.StringFixed(12, 64)
}

// Exchange
func (m MarketOrdersRequestFixed) Exchange() string {
	return m.p.StringFixed(76, 16)
}

// SendQuantitiesGreaterOrEqualTo
func (m MarketOrdersRequestFixed) SendQuantitiesGreaterOrEqualTo() uint32 {
	return m.p.Uint32LE(92)
}

// SetRequestAction
func (m *MarketOrdersRequest) SetRequestAction(value RequestActionEnum) *MarketOrdersRequest {
	m.p.SetInt32LE(8, int32(value))
	return m
}

// SetSymbolID
func (m *MarketOrdersRequest) SetSymbolID(value uint32) *MarketOrdersRequest {
	m.p.SetUint32LE(12, value)
	return m
}

// SetSymbol
func (m *MarketOrdersRequest) SetSymbol(value string) *MarketOrdersRequest {
	m.p.SetStringVLS(16, value)
	return m
}

// SetExchange
func (m *MarketOrdersRequest) SetExchange(value string) *MarketOrdersRequest {
	m.p.SetStringVLS(20, value)
	return m
}

// SetSendQuantitiesGreaterOrEqualTo
func (m *MarketOrdersRequest) SetSendQuantitiesGreaterOrEqualTo(value uint32) *MarketOrdersRequest {
	m.p.SetUint32LE(24, value)
	return m
}

// SetRequestAction
func (m *MarketOrdersRequestFixed) SetRequestAction(value RequestActionEnum) *MarketOrdersRequestFixed {
	m.p.SetInt32LE(4, int32(value))
	return m
}

// SetSymbolID
func (m *MarketOrdersRequestFixed) SetSymbolID(value uint32) *MarketOrdersRequestFixed {
	m.p.SetUint32LE(8, value)
	return m
}

// SetSymbol
func (m *MarketOrdersRequestFixed) SetSymbol(value string) *MarketOrdersRequestFixed {
	m.p.SetStringFixed(12, 64, value)
	return m
}

// SetExchange
func (m *MarketOrdersRequestFixed) SetExchange(value string) *MarketOrdersRequestFixed {
	m.p.SetStringFixed(76, 16, value)
	return m
}

// SetSendQuantitiesGreaterOrEqualTo
func (m *MarketOrdersRequestFixed) SetSendQuantitiesGreaterOrEqualTo(value uint32) *MarketOrdersRequestFixed {
	m.p.SetUint32LE(92, value)
	return m
}

func (m MarketOrdersRequest) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m MarketOrdersRequest) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m MarketOrdersRequestFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m MarketOrdersRequestFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m MarketOrdersRequest) Copy(to MarketOrdersRequest) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequest) CopyTo(to MarketOrdersRequestFixed) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// Copy
func (m MarketOrdersRequestFixed) Copy(to MarketOrdersRequestFixed) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

// CopyTo
func (m MarketOrdersRequestFixed) CopyTo(to MarketOrdersRequest) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetSendQuantitiesGreaterOrEqualTo(m.SendQuantitiesGreaterOrEqualTo())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 150)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("SendQuantitiesGreaterOrEqualTo", m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersRequest) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 150 {
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
		case "RequestAction":
			m.SetRequestAction(RequestActionEnum(r.Int32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SendQuantitiesGreaterOrEqualTo":
			m.SetSendQuantitiesGreaterOrEqualTo(r.Uint32())
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

func (m *MarketOrdersRequest) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketOrdersRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 150)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Uint32Field("SendQuantitiesGreaterOrEqualTo", m.SendQuantitiesGreaterOrEqualTo())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketOrdersRequestFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 150 {
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
		case "RequestAction":
			m.SetRequestAction(RequestActionEnum(r.Int32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SendQuantitiesGreaterOrEqualTo":
			m.SetSendQuantitiesGreaterOrEqualTo(r.Uint32())
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

func (m *MarketOrdersRequestFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
